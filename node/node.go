package node

import (
	"encoding/json"
  "golang.org/x/net/context"
  etcd "github.com/coreos/etcd/client"
	"github.com/llparse/streamingchan/fourchan"
	"github.com/satori/go.uuid"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	"flag"
	"net/http"
	"fmt"
	"errors"
)

type Node struct {
	Id               string
	Stats            *NodeStats
	Boards           *fourchan.Boards
	Storage          *fourchan.Storage
	Config           *NodeConfig
	Keys             etcd.KeysAPI
	LMDelay          time.Duration
	CBoard           chan *fourchan.Board
	CThread          chan *fourchan.ThreadInfo
	CPost            chan *fourchan.Post
	CFile            chan *fourchan.File
	BoardStop        []chan bool
	Shutdown         bool
	Closed           bool
	OwnedBoards      []string
	LastNodeIdx      int
	NodeCount        int
	DivideMutex      sync.Mutex
	Info             *NodeInfo
	stop             chan<- bool
}

type NodeInfo struct {
	Id        string `json:"id"`
	Hostname  string `json:"hostname"`
}

type NodeConfig struct {
	EtcdEndpoints []string
	CassEndpoints []string
	CassKeyspace  string
	CassNumConns  int
	Hostname      string
	BindIp        string
	CmdLine       string
	OnlyBoards    []string
	ExcludeBoards []string
	ClusterName   string
	ThreadWorkers int
	HttpPort      int
	NoFiles       bool
}

const (
	boardLock = "/%s/board-lock/%s"
	boardLM = "/%s/board-lm/%s"
	nodeLockTTL = 30 * time.Second
	boardLockTTL = 10 * time.Second
)

func (n *Node) Close() {
	n.Closed = true
	close(n.CBoard)
	close(n.CThread)
	close(n.CPost)
	close(n.CFile)
	n.Storage.Close()
}

func parseFlags() *NodeConfig {
	c := new(NodeConfig)
	hostname, _ := os.Hostname()
	flag.Bool(                         "node",          false,                   "Node : Enable node process.")
	var etcdEndpoints = flag.String(   "etcd",          "http://127.0.0.1:2379", "Node : Etcd addresses (comma-delimited)")
	var cassEndpoints = flag.String(   "cassandra",     "127.0.0.1",             "Node : Cassandra addresses (comma-delimited)")
	var onlyBoards    = flag.String(   "onlyboards",    "",                      "Node : Boards (a,b,sp) to process. Comma seperated.")
	var excludeBoards = flag.String(   "excludeboards", "",                      "Node : Boards (a,b,sp) to exclude. Comma seperated.")
	flag.IntVar(   &(c.ThreadWorkers), "tw",            10,                      "Node : Number of concurrent thread downloaders.")
	flag.StringVar(&(c.Hostname),      "hostname",      hostname,                "Node : Hostname or ip, visible from other machines on the network. ")
	flag.StringVar(&(c.BindIp),        "bindip",        "127.0.0.1",             "Node : Address to bind to.")
	flag.StringVar(&(c.CassKeyspace),  "keyspace",      "chan",                  "Node : Cassandra keyspace")
	flag.StringVar(&(c.ClusterName),   "clustername",   "streamingchan",         "Node : Cluster name")
	flag.IntVar(   &(c.HttpPort),      "httpport",      3333,                    "Node : Host for HTTP Server for serving stats. 0 for disabled.")
	flag.BoolVar(  &(c.NoFiles),       "nofiles",       false,                   "Node : Don't download files")
	flag.Parse()

	c.EtcdEndpoints = strings.Split(*etcdEndpoints, ",")
	c.CassEndpoints = strings.Split(*cassEndpoints, ",")
	if len(*onlyBoards) > 0 {
		c.OnlyBoards    = strings.Split(*onlyBoards,    ",")
	}
	if len(*excludeBoards) > 0 {
		c.ExcludeBoards = strings.Split(*excludeBoards, ",")
	}
	c.CmdLine = strings.Join(os.Args, " ")
	return c
}

func NewNode(stop chan<- bool) *Node {
	n := new(Node)
	n.stop = stop
	n.Stats = NewNodeStats()
	n.Config = parseFlags()
	n.Storage = fourchan.NewStorage(n.Config.CassKeyspace, n.Config.CassEndpoints...)
	cfg := etcd.Config {
		Endpoints:               n.Config.EtcdEndpoints,
		Transport:               etcd.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := etcd.New(cfg)
	if err != nil {
		log.Fatal("Failed to connected to etcd: ", err)
	}
	n.Keys = etcd.NewKeysAPI(c)
	n.Shutdown = false
	n.Closed = false
	return n
}

func (n *Node) refreshNode() {
	log.Print("Refreshing node ", n.Info.Id)
	_, err := n.Keys.Set(
		context.Background(),
		fmt.Sprintf("/%s/nodes/%s", n.Config.ClusterName, n.Info.Id),
		"",
		&etcd.SetOptions{TTL: nodeLockTTL, Refresh: true})

	if err != nil {
		log.Fatalf("Couldn't refresh node %s", n.Info.Id)
	}
}

func (n *Node) refreshNodeLoop() {
	ticker := time.NewTicker(nodeLockTTL / 2)
	for {
		select {
			case <-ticker.C:
				n.refreshNode()
		}
	}
}

func (n *Node) Bootstrap() error {
	log.Print("Bootstrap routine started")

	n.Id = uuid.NewV4().String()
	n.Info = &NodeInfo{n.Id, n.Config.Hostname}
	nodeInfoJson, _ := json.Marshal(n.Info)
	log.Printf("Creating node %s\n", n.Info.Id)

	_, err := n.Keys.Set(
		context.Background(),
		fmt.Sprintf("/%s/nodes/%s", n.Config.ClusterName, n.Info.Id),
		string(nodeInfoJson),
		&etcd.SetOptions{TTL: nodeLockTTL})

	if err != nil {
		return err
	}

	log.Print("Bootstrap routine completed")

	go n.refreshNodeLoop()
	go n.topologyWatcher()

	log.Print("Downloading boards list...")
	n.Boards, err = fourchan.DownloadBoards(n.Config.OnlyBoards, n.Config.ExcludeBoards)
	log.Printf("%+v", n.Boards)
	if err != nil {
		return err
	}

	go n.statusServer()

	n.LMDelay = 2 * time.Second / time.Duration(len(n.Boards.Boards))
	n.CBoard = make(chan *fourchan.Board, len(n.Boards.Boards) + 1)
	n.CThread = make(chan *fourchan.ThreadInfo, 4)
	n.CPost = make(chan *fourchan.Post, 8)
	n.CFile = make(chan *fourchan.File, 4)

	for _, board := range n.Boards.Boards {
		n.Storage.PersistBoard(board)
		n.CBoard <- board
	}
	log.Print("Finished bootstrapping node.")

	log.Print("Starting processor routines...")
	go n.boardProcessor(n.CBoard, n.CThread)
	for i := 0; i < 2; i++ {
		go n.threadProcessor(n.CThread, n.CPost)
	}
	for i := 0; i < 4; i++ {
		go n.postProcessor(n.CPost, n.CFile)
	}
	for i := 0; i < 2; i++ {
		go n.fileProcessor(n.CFile)
	}
	return nil
}

func (n *Node) topologyWatcher() {
	log.Println("Watching node topology for changes")

	watcher := n.Keys.Watcher(
		fmt.Sprintf("/%s/nodes", n.Config.ClusterName), 
		&etcd.WatcherOptions{Recursive: true})

	for {
		resp, err := watcher.Next(context.Background())
		if err != nil {
			log.Print("watcher failed:", err)
			continue
		}
		log.Printf("node topology updated (%s)\n", resp.Action)
	}
	// coordination stuff
	/*nodesKey := fmt.Sprintf("/%s/nodes", n.Config.ClusterName)
	resp, err := n.Keys.Get(
		context.Background(),
		nodesKey,
		&etcd.GetOptions{Recursive: true, Quorum: true})
	}*/
	// end coordination stuff
}

func (n *Node) acquireBoardLock(board string, ttl time.Duration) error {
	path := fmt.Sprintf(boardLock, n.Config.ClusterName, board)
	setopt := &etcd.SetOptions{TTL: ttl}
	resp, err := n.Keys.Get(context.Background(), path, nil)
	// Acquire lock
	if err != nil && (err.(etcd.Error)).Code == etcd.ErrorCodeKeyNotFound {
		_, err = n.Keys.Set(context.Background(), path, n.Id, setopt)
	// Refresh held lock
	} else if resp.Node.Value == n.Id {
		_, err = n.Keys.Set(context.Background(), path, n.Id, setopt)
	// Lock held by other node
	} else {
		err = errors.New(fmt.Sprintf("lock already exists on board %s", board))
	}
	return err
}

func (n *Node) releaseBoardLock(board string) {
	path := fmt.Sprintf(boardLock, n.Config.ClusterName, board)
	if _, err := n.Keys.Delete(context.Background(), path, nil); err != nil {
		log.Printf("couldn't release lock on board %s", board)
		log.Println(err)
	}
}

func (n *Node) getBoardLM(board string) int {
	var resp *etcd.Response
	var err error
	lm := 0

	path := fmt.Sprintf(boardLM, n.Config.ClusterName, board)
	resp, err = n.Keys.Get(context.Background(), path, nil)
	if err != nil {
		if (err.(etcd.Error)).Code != etcd.ErrorCodeKeyNotFound {
			log.Fatal("error getting lastModified: ", err)
		} else {
			log.Printf("no lastModified set for board %s\n", board)
		}
		return 0
	} else {
		lm, err = strconv.Atoi(resp.Node.Value)
		if err != nil {
			log.Print("error parsing lastModified: ", err)
			return 0
		}		
	}
	return lm
}

func (n *Node) setBoardLM(board string, lastModified int) {
	path := fmt.Sprintf(boardLM, n.Config.ClusterName, board)
	if _, err := n.Keys.Set(context.Background(), path, strconv.Itoa(lastModified), nil); err != nil {
		log.Printf("Error setting lastModified for board %s", board)
	}
}

func (n *Node) boardProcessor(boards chan *fourchan.Board, threads chan<- *fourchan.ThreadInfo) {
	lockTTL := 10 * time.Second
	for board := range boards {
		// eliminate drift by publishing to channel immediately
		boards <- board
		if err := n.acquireBoardLock(board.Board, lockTTL); err != nil {
			log.Print(err)
			continue
		}
		ticker := time.NewTicker(lockTTL / 2)
		go func() {
			for _ = range ticker.C {
				if err := n.acquireBoardLock(board.Board, lockTTL); err != nil {
					log.Fatal("couldn't refresh lock on board ", board.Board, ": ", err)
				}
			}
		}()
		log.Printf("processing /%s/", board.Board)
		board.LM = n.getBoardLM(board.Board)
		var lastModifiedHeader time.Time
		if t, statusCode, lastModifiedStr, e := fourchan.DownloadBoard(board.Board, lastModifiedHeader); e == nil {
			n.Stats.Incr(METRIC_BOARD_REQUESTS, 1)
			lastModifiedHeader, _ = time.Parse(http.TimeFormat, lastModifiedStr)
			//log.Printf("lm %d: %s", board.LM, board.Board)
			newLastModified := board.LM
			for _, page := range t {
				for _, thread := range page.Threads {
					if thread.LastModified > board.LM {
						thread.Board = page.Board
						thread.MinPost = board.LM
						thread.OwnerId = n.Id
						threads <- thread
					}
					if thread.LastModified > newLastModified {
						newLastModified = thread.LastModified
					}
				}
			}
			if board.LM == newLastModified {
				time.Sleep(n.LMDelay)
				//log.Printf("board %s didn't change", board.Board)
			} else {
				board.LM = newLastModified
				n.setBoardLM(board.Board, board.LM)
			}
		} else if statusCode != 304 {
			log.Print("Error downloading board ", board.Board, " ", e)
		}
    ticker.Stop()
		n.releaseBoardLock(board.Board)
	}
}

func (n *Node) threadProcessor(threads <-chan *fourchan.ThreadInfo, posts chan<- *fourchan.Post) {
	for threadInfo := range threads {
		//log.Printf("processing /%s/thread/%d", threadInfo.Board, threadInfo.No)
		n.Storage.PersistThread(threadInfo)
		if thread, err := fourchan.DownloadThread(threadInfo.Board, threadInfo.No); err == nil {
			n.Stats.Incr(METRIC_THREADS, 1)
			var postNos []int
			for _, post := range thread.Posts {
				if post.Time >= threadInfo.MinPost {
					postNos = append(postNos, post.No)
					posts <- post
					n.Stats.Incr(METRIC_POSTS, 1)
				}
			}
			n.Storage.PersistThreadPosts(threadInfo, postNos)
		} else {
			//log.Print("Error downloading thread: ", err)
		}
	}
}

func (n *Node) postProcessor(posts <-chan *fourchan.Post, files chan<- *fourchan.File) {
	for post := range posts {
		//log.Printf("processing /%s/post/%d", post.Board, post.No)
		n.Storage.PersistPost(post)
		if post.Md5 != "" && post.Ext != "" && !n.Closed {
			file := &fourchan.File{
				Board: post.Board,
				Tim: post.Tim,
				Md5: post.Md5,
				Ext: post.Ext,
				FSize: post.FSize,
			}
			files <- file
		}
	}
}

func (n *Node) fileProcessor(files <-chan *fourchan.File) {
	for file := range files {
		if n.Config.NoFiles {
			continue
		}
		//log.Printf("processing /%s/file/%d", file.Board, file.Tim)
		if !n.Storage.FileExists(file) {
			data, err := fourchan.DownloadFile(file.Board, file.Tim, file.Ext)
			if err == nil {
				file.Data = data
				n.Storage.WriteFile(file)
			} else {
				log.Printf("Error downloading file %+v: %+v", file, err)
			}
		} else {
			//log.Printf("File exists: %+v", file)
		}
	}
}

func (n *Node) CleanShutdown() {
	if n.Shutdown {
		return
	}
	n.Shutdown = true
	log.Print("Removing node from cluster.")
	// TODO remove node from cluster
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(120 * time.Second)
		if false {
			log.Print("Timeout waiting for sockets.")
			stack := make([]byte, 262144)
			runtime.Stack(stack, true)
			log.Print("----------- DUMP STACK CALLED ----------------")
			log.Print("\n", string(stack))
		}
		timeout <- true
	}()
	go func() {
		log.Print("Closing...")
		n.Close()
		log.Print("Wait for routines to finish...")
		// TODO: wait group(s)?
		log.Print("Shut down node.")
		timeout <- true
	}()
	<-timeout
}
