// Code generated by go-bindata.
// sources:
// asset/templates/thread.html
// DO NOT EDIT!

package asset

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetTemplatesThreadHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x5b\x5b\x73\xdb\xb8\x92\x7e\x9f\x5f\x81\xc3\x54\xc5\x99\xb3\xba\xf8\x9e\xf8\xa6\x29\xdb\x89\x33\xd9\x8d\xe3\x54\xec\x9d\xec\xd4\xd6\x94\x0b\x24\x21\x11\x11\x49\xf0\x00\xa0\x64\xc5\xa5\xff\xbe\xdd\xe0\x0d\xa4\x28\x8b\xce\x4e\x1e\xce\x43\x22\x12\x40\x77\x03\x8d\xbe\x7c\x0d\xc2\x8f\x8f\x3e\x1b\xf3\x98\x11\x47\x07\x92\x51\xdf\x59\x2e\x7f\x39\xfd\xc7\xdb\x9b\xcb\xbb\x3f\x3f\xbf\x23\x81\x8e\xc2\xd1\x2f\xa7\xd9\x0f\x21\xa7\x01\x8c\xc0\x07\x78\x8c\x98\xa6\xc4\x0b\xa8\x54\x4c\x9f\x39\xa9\x1e\xf7\xdf\x38\x76\x57\x4c\x23\x76\xe6\x48\xe1\x0a\xad\x1c\xe2\x89\x58\xb3\x18\x06\xc6\x82\x4a\x2f\xe0\x33\xd6\x32\xd8\x67\xca\x93\x3c\xd1\x5c\xc4\x16\xc5\x55\xea\x4d\x09\xf4\x4b\xee\xd1\x58\xf5\x08\x9b\xb1\x98\x50\x49\x5d\x05\xff\x33\xa2\xd2\x04\xba\x84\x04\x82\x28\x81\x06\x9f\x68\x41\x74\xc0\x88\xcf\x23\x32\xe7\x5a\x43\xcb\x9c\x2a\xcd\x88\x18\x13\x16\x33\x39\x59\xf4\xc8\x24\x8d\x49\x22\x64\x4c\x42\x31\xe3\xf1\xa4\x37\x18\x0c\x48\x9f\xbc\xfc\x57\x2a\xf4\xc9\xd0\x1d\xc2\xf3\x17\x1a\xfb\x22\xca\x5a\x08\x57\x86\xa3\xcb\xa5\x0e\x92\x90\x7a\x86\xd7\x79\x2c\xe2\x45\x24\x52\x98\x12\x8c\x25\xf3\x80\xc1\x6c\x12\x26\x92\x90\x91\x89\xc0\x59\xf8\x5c\x79\xa9\x52\x44\x1a\x5e\xd0\x92\x70\x4f\x99\xc1\x1e\xa8\x1a\x66\x14\xb1\x88\x29\x22\x62\xb2\x0f\x8a\x8c\x07\x2d\x2a\x99\xb2\xc5\x5c\x48\xdf\xd6\x20\x8f\xe8\x84\xb9\xa0\x46\xbf\x07\x5d\x53\x45\xc7\xac\x47\xcb\xb9\x64\xb2\x7a\x86\x73\x0f\x55\xd7\xab\x54\x87\x9a\xeb\x19\xcd\xf5\x0a\xad\xf5\x0a\xad\xf5\x32\x55\xf5\x8c\xaa\x7a\xb9\x9e\x50\x45\xbd\x4c\x45\x6d\x7b\xcb\xc6\x4c\x4a\x26\xad\xb9\x09\xc9\x27\x3c\x6e\x19\x3b\xe3\x6c\x0e\xdc\xb4\x35\x76\xce\x7d\x1d\x9c\xf9\x6c\xc6\x3d\xd6\x37\x2f\x3d\x1e\x73\xcd\x69\xd8\x57\x1e\x0d\xd9\xd9\x4e\xc1\x27\xe4\xf1\x94\x48\x16\x9e\x39\x2a\x00\x1e\x5e\xaa\x09\xf7\xd0\x46\xc0\x62\xc7\x67\xce\x70\xa8\x06\xfb\x9e\x1f\x0f\x84\x9c\x0c\x8d\x76\x86\x63\x3a\xc3\x11\x03\xf8\xaf\x85\x89\x5e\x84\x4c\x05\x8c\xc1\x6c\x34\xd7\x20\xc9\x51\xb0\x7a\x2f\x68\x65\xe8\x29\x35\x5c\x80\x0d\xa7\x2e\x8d\xd9\x7c\x70\xb8\x7b\x34\x80\xa6\x55\xae\x34\xd4\x4c\xc6\xb8\xab\x36\x7f\xf3\x7c\xe6\x68\xf6\xa0\x91\xd3\x73\x24\x14\x73\xfb\x33\xeb\x22\x9f\xd8\xfc\xe7\x48\x0d\x53\xf6\x94\xdc\x8b\x9f\x21\x79\x9c\x6a\xda\xbe\xdc\x2b\xd3\xf3\x33\x64\xba\x29\xf8\x01\xf8\x59\x9b\xd4\x8b\xbc\xef\x67\xc8\x4d\x02\xa1\xc1\x16\x9b\x22\x3f\x9b\xe6\xbf\x5b\x98\x16\x91\x90\x52\xac\xae\xf0\x2e\xef\x78\xda\x1f\x9e\x36\xcf\x48\xb8\x3c\x64\xeb\x7d\xc0\xc3\x38\x04\xb1\x26\x2c\x18\x05\x5a\x27\xc7\xc3\xa1\x89\x56\xc8\x12\xa3\x1c\xf2\x74\x87\x59\xb6\x19\x1e\xee\x1f\xed\xbe\x39\x3c\x7c\xb3\x0f\x16\xe1\x4d\xfb\x65\xac\xea\x63\xac\xea\x9b\x58\x05\xff\xb3\x7e\x11\xaf\xfa\x45\xbc\x7a\x42\x6f\xe5\x92\xbf\xdc\xde\x92\x31\x83\xb1\xc5\xb2\xdc\x21\x8f\x7d\xf6\x30\x90\x46\x2f\x8b\x04\xc6\xd0\x24\x09\x41\x22\xe6\x9c\x21\x34\xff\xc7\x43\x14\x16\xac\x0d\x97\x51\x96\x0e\x9e\x9f\x84\xca\x1c\x02\x0f\x66\xe1\x19\xd3\x61\xc6\x35\x7b\xc9\xd2\x5d\x3e\x13\xb3\xb5\xdf\xe8\x8c\x66\xad\xce\x68\x46\x25\xc9\xf6\xfd\x7e\x22\x45\x9a\x90\x33\xe2\xc4\x73\x75\x6f\x9a\x9c\x1e\x21\xb0\x07\x7f\x30\xa9\x60\xea\xd0\x05\x9b\x02\x4d\xdf\xac\x96\x9d\xed\x9d\x5d\x1c\x25\xa2\x90\xe1\xfb\xee\xf6\xf6\x36\xbc\x47\xf4\xe1\x0a\x76\x51\xf1\xef\x0c\x1a\xf7\x77\x8e\xf6\xf7\xb6\xf7\xb3\xf6\x8f\x00\x04\x14\x34\x1e\xe0\xb8\x31\x0c\xba\xd7\x42\xdc\xbb\x7c\x82\xa2\xdf\x81\xf5\xc8\x63\x72\x4d\x1f\x78\x94\x46\xa6\x9b\x18\x26\x34\x0c\xc5\x1c\x96\x0c\x59\x72\x7f\xfb\xe8\x90\xfc\xd7\xc5\xc0\xcc\x0e\xf4\x3a\xa5\x2e\x30\xe1\x3e\x32\xdd\x31\x73\x11\xa1\x2f\xe6\x31\xbe\x3f\x16\x88\xe3\xf8\x70\xbb\x07\xc9\x24\x09\x17\xce\xf1\xce\x41\x2f\x4b\x70\xd9\xa3\x69\xbd\xe7\xb1\x96\xd4\xea\xab\x1a\x96\x27\xa8\xa4\x34\xe6\xff\x4a\xa1\x35\x41\xb6\xbb\xa6\x09\x16\xf3\x95\xb9\x91\xb5\xd0\xbd\x9d\xfd\x83\xd7\xbb\x6f\x4c\xaf\x17\x30\x6f\x7a\x3f\x16\xf2\xde\x0d\x05\x6c\x2c\x4c\xee\x84\x98\xc7\xcf\x21\xa3\xf0\xba\x85\xbf\x8a\x9d\x02\xc2\x09\xa9\x52\x68\x26\x61\x3c\xad\xdc\x63\x3e\x9f\x5b\xd6\x0c\x01\x45\xfd\x06\x5a\x78\xb1\xb3\x7d\x04\x76\x45\xe5\x04\x21\x11\xf0\xa6\x40\x33\x02\xdb\xc0\xcc\x97\x19\xc1\xe9\x90\x8e\x88\xbb\x40\x68\x00\xaa\x81\xbc\x4a\x16\x22\x95\x84\xfa\x99\x78\x26\x11\x0d\xfc\xd3\xe2\xfd\xcf\xde\x9a\x49\xa0\x6b\xa9\xe3\xe6\x54\xa8\x3f\x63\x52\x73\xc5\x7e\x53\x2c\x1c\x2b\x26\x67\x6c\x75\x46\x49\x0a\x10\x8c\x2a\x94\x4e\x09\x8e\xeb\x9b\x81\x30\x0b\x9c\x5e\x8f\x08\xf9\x2c\x99\x09\x35\xee\xd4\x10\xe2\xa6\x8b\x4c\x80\x19\x48\x3e\xc3\x20\xe4\x3e\xd8\x3a\xc9\x1d\x21\x33\xf3\x2e\x9e\x40\x7c\xaa\x69\xdf\x1b\x53\xb5\x88\xbd\x33\x67\x4c\x43\x05\xab\x52\xd2\x6b\x44\xaa\x6f\x6a\xe8\x09\xc9\x06\x11\x8f\x07\x47\x47\x07\x83\x6f\x65\x88\xfa\x39\xc2\x80\x92\xc5\xe8\x6c\x46\xe2\xce\xf6\xee\xee\x8f\x89\x1c\x21\xee\x39\xf7\xd5\xab\x2d\x70\xef\xad\x1e\xd9\x72\xb7\x7e\x6d\xd5\x52\x2c\xea\x4c\x31\x10\xd8\x3c\x4d\x50\x7e\x91\x08\xa5\xaf\x84\x8c\x1e\xc1\xc8\x00\xad\x2e\x8e\x35\x7a\xe1\x3f\x78\x84\x66\x48\x63\x7d\xb2\x7c\x31\xe9\x4b\xe6\xd1\x04\x30\x0f\x2d\x47\x41\xe8\x66\x27\xcb\x42\x2c\xb2\x2e\x56\x62\x8b\x3d\x1d\x16\x35\xc0\xa9\x2b\xfc\x45\x61\x26\x26\xc2\xdf\xbb\xc5\xe2\x55\x02\x5b\xce\x7d\x00\xab\xfe\xbd\x95\x29\x86\xd8\x9e\x3f\xfb\x7c\x66\x46\x18\xca\x4f\x74\xf6\x96\xa9\x29\x20\x64\xa7\xe0\xe8\xe7\xef\xd9\xf0\x82\xa7\x2d\xee\x23\x57\xa0\xbb\xff\xcd\xfb\x61\x04\x2d\xfc\x93\x0e\xcb\x1c\x70\x1e\xf3\x88\x91\x97\x10\xb6\xe2\x09\x75\x46\xb4\x1a\x8d\x9e\x38\x6c\x21\x76\x2b\xe2\x2c\x80\x83\x2d\x6f\xa6\xf2\x1a\x22\x87\x97\x29\x64\xa2\x91\xb7\x99\xd2\xaf\x28\x7f\x07\x60\x4c\xf9\xf0\x3c\xcf\x64\xa6\x4a\xf2\x37\x73\x60\x15\x87\x77\x1e\x14\x57\xce\x88\x6d\x26\x1a\x57\x44\x57\xa0\xd4\xc0\x19\x8d\x37\x13\x4d\x2a\xa2\x3b\xe6\x05\xb1\x08\xc5\x64\xe1\x8c\x26\x1d\x28\xb9\x25\xf0\xdc\x4f\x43\x4d\xde\x7f\xb8\x02\x52\xde\x41\x6c\xd0\x54\x91\x33\x0a\x3a\x50\x49\x8b\x8c\x4f\x02\xf2\x85\x29\x11\xa6\xa6\xbc\x1c\x05\x72\x33\x83\x69\x45\xff\x95\xd1\x44\xc4\x60\xc8\xd3\xcd\x64\x51\x45\x76\x0d\x4a\x02\xab\x8b\x36\x13\x09\x4b\x3b\xa9\x86\xaa\x45\x6c\xa6\x49\x86\x75\x30\x09\x91\x7d\x33\x91\xa5\x94\x2f\x0c\xd2\x26\xfa\x50\x07\x5d\xa8\x8a\xec\x96\x3d\x2c\xc8\x05\xa3\xa0\xc9\x71\x1a\x92\xaf\x22\x62\xa0\x50\xb5\x99\x87\xb6\xac\x07\xb0\x28\xec\x24\x28\x54\x6f\xa6\x4b\x2b\xba\x3f\x01\xa7\x3b\xa3\x74\x33\xcd\xac\xa2\xf9\x83\xfb\x4c\x90\xf7\x80\xde\x40\xdc\xac\x03\xe9\xa4\x8d\x96\xbc\xc7\xa2\x18\xb2\x01\xf0\xe8\x60\xef\xb3\x9a\x9e\xb5\xac\x26\xd0\x41\xd9\xf3\x66\x30\xf9\x0a\xe8\x22\xa1\x80\x31\x81\xc1\xbc\x03\xbd\xb5\x82\x8a\x74\x98\xaf\x00\x58\xd4\x17\xf0\x17\x69\x0b\xa2\xbc\xe2\x71\xc3\xa6\x74\x0a\x7a\xe7\x9b\x45\x73\x3b\x10\x4a\x8d\x27\x14\xc3\x4b\x09\x65\x3d\x98\x1a\x30\xf0\x3a\x08\x96\x47\x96\xe3\x7d\xb9\xb9\xb8\xb9\x3b\xda\xde\xde\x01\x2b\x3d\x9a\x76\xa0\x56\xfb\xb6\xa9\x06\x3c\x07\x5e\xe4\x96\x2e\x40\x79\xd0\xdb\x81\x87\x67\xb9\x30\xc6\xf1\xe1\x35\x0d\x31\x98\x77\xf0\xe3\xc0\x22\xfd\x1d\xd2\x87\x02\xef\x20\xd7\xe8\x20\x41\x07\xea\x70\xe2\x5a\x4e\xf2\xf1\xfd\xc5\x9d\x33\xc2\xb6\xcd\x94\x0b\xcb\x47\xa8\x80\xbd\x5a\x74\x58\xe7\x5e\x45\xb4\xf7\xf6\xf2\xbd\x33\xda\xdb\x2c\x88\x7a\xa2\x19\xc8\x2f\xa9\x84\x3a\x01\xa3\x23\x74\x76\xe0\xe0\xcf\x6c\x0e\x78\xfe\x03\x94\x7e\x07\xcf\xa4\x71\xdd\x31\xc0\x1b\x21\xb5\x7f\xa2\x3a\x95\xc8\x22\xee\xc0\x41\x59\x41\xd3\xca\xb3\xe4\x16\xc1\x11\xae\x40\x75\x88\xa1\x2e\xff\x3e\xb4\x8e\x0e\x14\x96\x4d\x38\x93\x2b\x1e\xd3\x18\x57\x03\x03\x3a\x40\x86\x49\x68\x99\x99\x30\x28\x0c\x98\xbc\x7b\xff\x11\x4c\x6d\x12\x76\x60\x60\xf9\xc9\x95\x10\x3e\x10\x5f\x0a\x31\x35\xa7\x75\x5e\x87\x44\x65\xef\xe4\xa5\x88\xf0\x70\xf2\xa5\xb5\x99\x5d\xf6\xd2\xe7\x96\xe1\xbd\x15\xe4\x83\x26\x7f\x42\x41\x83\x45\x05\x40\x17\xbe\xe8\x80\x43\x2c\xb0\x76\x05\x38\xc4\x64\xe7\x71\x07\x9c\x36\xe6\x96\xab\x5c\x71\x8d\x7b\x00\x94\xbc\x83\xb3\x4c\x2c\xcc\xf5\x5e\xd2\x24\xe0\x1e\x01\xfc\xc9\x27\x20\x7a\xd2\x01\x70\x05\x9e\xed\xe4\xd2\xc7\xaa\x03\x1c\xbc\x03\xd8\x83\xca\xb5\x22\xfd\x10\xe7\xf6\x27\x62\x8c\xca\xd0\xb7\x99\xc1\x37\xcb\x7e\x6f\x34\x9d\xa6\xe4\x12\x5c\xd0\xd8\xff\xb7\x0e\x96\x1b\xda\x4a\xfb\xc8\x41\x7e\xee\x3c\x61\x17\xbd\x45\xa1\x0d\x39\x44\x0c\xd0\x0f\x9a\x3a\xd0\x59\x19\xfc\x1a\xdc\xc5\x03\xba\x0e\x39\xdc\xf2\xf6\x3b\x49\x63\x65\xea\x97\x0c\xbe\x75\x70\x75\x91\x5a\x4b\xbd\x49\xb5\x2f\x04\xa6\x4f\x68\xee\x80\xac\x2c\xcf\xf8\x8c\xc9\xd3\x93\x74\xac\xc1\x3b\x6e\x24\x9f\xd0\x08\xc2\x6b\xd2\x05\x9f\x89\xd0\x56\x17\xa8\x18\xcf\xc6\xc2\x05\xf9\x10\x7b\x88\x7f\x3c\x8d\x7c\x3a\x38\xba\xf2\xac\x64\x7c\xeb\x71\x06\x41\xc6\xd4\x34\x1a\x60\x3b\x74\x76\xe0\x20\x2c\x93\xbd\xa4\x11\x7a\xfa\x35\x63\x3a\x4d\x30\x29\x8a\x0e\x96\x6b\x07\xce\x22\x58\x76\x89\x95\x7a\x52\xdb\x45\x9f\x67\xe6\x5e\xc0\x21\xdd\x01\x4b\x69\xb1\xb0\x91\x23\xe6\x71\x68\xea\x40\x27\x67\x35\xd9\x33\x06\x5e\x06\x8d\x1d\x28\x6d\x42\x16\xb2\x19\x37\x67\x6a\x18\xe1\x43\xa8\x07\x75\x17\x14\x59\x73\x95\xe9\x4b\x46\x3d\x40\x12\x27\x11\x5a\xef\xac\x83\xde\xe6\xca\x06\x71\xf9\x87\x9e\xac\x66\x82\xae\xcd\xf4\x0f\xb6\x05\x83\xf7\x08\x19\x61\x90\x79\xa8\xc3\x82\x5f\x8a\xe7\xaa\x32\x27\x56\xf9\x1e\xd3\x19\xd4\xe0\x60\xf3\x41\x7b\xa5\x5d\x1d\x62\x1c\xcf\x04\xf7\x5f\x6d\xff\x7a\xe2\x18\x42\xc5\xb4\x86\x3c\xa4\xbe\x72\xa8\xa0\xe7\x1f\x39\x9e\x08\xdd\xe6\x6d\x1d\x80\x49\xe3\x84\xc9\x3a\x5d\x32\x47\x02\xbf\x03\xa8\xda\xb8\x8e\x53\x48\x4e\xb3\x35\x87\x0d\xd7\xe6\x20\xbb\x3c\x6b\xc8\xce\xb5\xab\xa3\x06\x1c\x6e\x9f\x34\xdc\x82\x09\xa0\xb3\x56\x32\x15\x20\xfa\x78\x32\xba\xc0\x5e\x6b\x26\x79\xb3\x35\xce\x50\x56\xb2\x33\x4e\xd7\x75\x79\x86\xd2\x74\x94\x33\xa8\xe6\x5e\x9f\x4e\x42\x27\xec\x3f\xd3\x28\xb1\x69\x0b\xad\xbd\x70\x85\xd6\x78\x5c\xf1\xf2\xc5\xd1\xe1\xe1\xf6\x49\x4d\x43\xff\xaf\xbd\x2b\xe6\xdb\xba\x83\x1b\xb7\x6f\xe3\xee\xad\x2e\xba\xb9\x77\xf6\x66\x5c\xe0\x17\x1c\x59\xdf\x2c\xa3\x5f\xd3\x7e\x19\xeb\x72\x5b\x8d\xf9\x93\xe2\x20\x29\x3b\xd8\x33\x07\x79\x3b\x07\x7b\x83\xa4\xfc\xa6\xf9\x84\xba\x8d\xc0\x3b\x64\xe3\x8c\xec\xef\xc2\x9d\xe8\x6e\x53\x57\x67\xa4\x77\x01\x7e\xcc\x11\x92\xb3\xec\xf3\x2f\x8f\xc7\xe8\x8e\x18\x09\x09\x9e\xd6\x31\x9f\x98\xaf\xc7\xd4\xfc\xd3\x5c\x41\xa2\x20\xe6\xe3\x2e\x7e\x68\x1e\x73\xcf\x8c\x44\x4a\x73\x1c\x19\x00\xd6\x1b\x54\x1a\x74\xe5\xe8\x26\x86\xac\x42\xc9\x58\x88\x10\xe8\xd2\xd0\x07\x95\x4f\x81\x57\xbc\xd0\x01\x9e\xc2\xd6\x84\x28\xe0\xe2\xe9\x41\x67\xad\x43\x10\x40\x23\x50\xa4\xe9\x26\xf6\x89\x5c\xd6\xc7\x5d\xe2\xa6\xba\xfa\xac\x55\x37\x8b\x7d\x73\xb6\x46\x3d\x0f\x90\xda\x94\x2d\xce\x1c\xea\x8c\xa0\x36\x4e\x65\xdc\x6e\x0e\xab\x31\xa9\xbb\xb0\xd2\x19\x2e\xcc\xef\x4f\x10\x00\x36\x75\x2f\x33\x8f\x81\x06\xc9\x54\x90\x99\xf7\x97\xec\x65\x93\xc4\xa6\xb2\xb3\x0f\x7c\x5b\xb0\x53\x26\x43\x1e\x4b\x16\x9a\xca\x64\x6b\x65\x78\x7d\x77\x22\xee\xfb\x21\xa3\x3e\xf1\x18\xc2\x49\x42\xfd\xbe\x17\xd7\x62\x55\xe1\x20\xf4\xfb\xf4\xe8\xcd\x9b\x37\xaf\x6b\x71\xc7\xe6\x5a\x7b\x09\x64\xf9\x58\x2e\xb9\xf4\xac\x4c\x2d\x44\x32\xff\x22\x7b\xca\x14\x46\x02\x98\x0c\xd4\xba\x46\xdc\x24\x14\x2e\x0d\xef\xc4\x64\x82\x36\xf3\x07\x67\x73\x72\x1e\xc7\x22\x05\xec\x12\xb1\x12\xea\xda\xda\xb1\x56\x95\x11\x5f\x83\xa1\x40\xc0\x43\xb6\xac\x08\xd9\x15\xeb\xbc\x37\x77\xec\x54\xa3\x63\xef\xef\x1f\xee\x1c\xec\xbf\x3e\xdc\x76\x46\x50\x87\x90\x08\x33\x17\x09\x39\xf8\x02\xd7\x83\xd5\x4d\xb4\x3f\x68\x2c\x00\x19\xa6\x2e\x1b\x78\x22\x1a\xce\xa9\xf6\x82\xdf\x66\x67\xb3\xff\x91\xfd\xdd\x60\x7e\x37\x3d\x78\x53\x46\xb1\xec\x03\xc7\xc8\x41\xa7\xce\x10\x3c\xd3\x5b\x8a\xdc\xcc\x63\x72\x21\x16\xc7\x04\xdb\x6f\xc1\xd9\x17\xe6\x92\x08\x85\x84\x40\x6e\xe7\xe0\xd5\xdf\x9d\x0d\x41\xcf\x28\xbd\x79\xfc\x6d\xd6\x6b\x1f\x83\xb7\x39\x66\x79\x56\xde\x96\x4b\x3b\x7a\x5d\x6b\x22\x7e\xca\x8b\xfe\xda\x64\x36\x18\xe3\xca\x9b\x3d\x21\xbe\x65\xab\x29\x5f\xa8\x89\x6b\xd5\x2e\xa8\x45\xfd\x1b\x31\x8f\x26\x26\x90\x0e\x92\x00\x14\x11\x31\x1d\x08\x20\xc7\x58\xd6\x34\x6f\x3b\xec\x5a\x7d\xf5\xde\xfc\x73\xa3\x99\x83\x7e\x7c\x1c\x7c\x12\xcb\x65\x6d\x30\x21\x8f\x8f\x80\x94\xc0\xe0\x06\x9f\x41\x88\x5a\x2e\x6b\x9d\x8d\x4c\x0c\x23\x2e\x05\x1e\x46\x43\xd6\x01\x42\x3e\x26\x83\x9b\x64\xb9\x14\x49\xd9\xfa\xf8\xc8\x20\x5a\x2f\x97\xe6\x2b\xa6\xdd\x1a\xfb\x20\xd9\xcc\x23\xf1\xda\x27\xd2\x94\xa6\xc0\x03\xce\xf1\xf3\xbd\xca\x73\x34\x2d\xe9\x5e\x4e\xf4\x09\xfe\xab\xed\x43\x83\x8b\x11\x55\x50\xd8\x2b\xa8\x4d\xbc\x36\xdf\x62\x9a\xab\x1c\x57\xf5\xf0\x01\xf2\xd9\x75\x91\x1e\x32\x61\x3c\x32\x9f\xf6\x5f\xef\x6e\x1f\xb4\xb2\x68\x04\x5c\x34\x93\x0b\xfc\x02\xba\x66\x70\xcb\x70\x67\x84\x0b\x82\x87\xe5\x32\x8f\xab\x98\x09\xdb\x45\xd5\x22\xfd\x13\x7c\x21\x92\xb0\x3b\xfc\x60\x84\xab\xfa\x94\x46\x76\x6c\x01\x69\xd8\x85\x2a\x31\x9a\x9c\x83\xdc\xd2\x4f\x2a\xe5\x96\xb5\x75\x3c\xcd\x6e\xa2\x71\x45\x32\x93\xfd\x24\x06\xe8\x37\x6d\x00\x0c\xaf\x99\xb1\x57\x5b\x39\x93\x2d\x44\x62\xe5\xd9\x31\xec\x45\x83\x51\x3e\xcc\x30\x5b\xbb\xb2\x35\xe6\xd0\xbe\x79\xa4\x16\x6e\x12\xbe\xde\x28\x0d\x07\x1e\x27\x69\xf1\x31\xd3\x7c\x4b\x77\xc5\x83\x93\xbb\x7a\xa9\x88\x19\x0d\xd3\xcc\xf5\x99\x66\xcf\xb4\x81\x2e\x9b\xfd\xdc\x3d\xdd\xb4\x97\x5d\xf9\xe5\xa6\x51\x05\xdd\x7f\x4b\x2b\xc8\xfc\x1e\xef\x47\xc4\x46\xab\x1b\x0c\x65\x5c\x7a\xf6\x78\x83\x6d\x34\x88\xee\xd8\x83\xce\x09\xef\x20\x22\xec\x01\xf4\xde\x07\x04\x82\x82\x8f\x6b\x87\x79\x21\x1b\x66\xbb\xb2\x5c\xc2\xef\xbb\x07\x6d\xd4\xd8\xb8\x54\x00\x3d\xd5\x94\x8b\x61\xa6\x04\x7e\x85\x5d\xb7\xfc\x3b\xfb\x2c\xa1\x48\x59\x2c\x97\x3d\x58\xe2\xe0\xeb\x72\xf9\x00\x3f\xbf\x2f\x97\xbf\xae\x09\x0c\x6b\xd4\x63\x72\xa0\xbd\x8c\x20\x8d\x5c\xe7\x99\xb3\x5d\x17\xc9\x20\xaf\xe5\x17\x0a\xd6\x30\xa2\xa1\x36\x36\x5a\x5b\x50\x6e\xbf\x91\x7f\x60\xfa\xae\xfd\x03\x6c\xcb\xef\x85\x05\x0c\x91\xce\x31\xae\xf9\x2e\x86\xe5\x26\x0f\x27\xc4\xdc\xa9\xcc\x9b\xbe\x9a\xa6\xf5\x53\xc2\x5d\x33\xec\x35\x4f\xce\x1c\xa7\x7c\xee\x7b\x2e\x40\xcc\xdb\x40\xcc\xaf\xd2\x30\xac\x0a\x66\xdc\x05\x13\x38\x8a\x9a\xa0\x39\x5b\x52\xac\x66\x9d\xc4\x27\x34\x6f\x17\x95\x1d\x28\xf2\x4c\xd5\x46\x62\x2e\xd5\x18\xb7\xb2\xbd\xb7\x04\x8e\x68\x96\x51\x69\xcf\xf0\x70\x29\x22\xb4\xa7\x8a\xac\x25\x9f\xb6\x4d\xa3\xa5\x71\x75\x56\x2b\x83\x2c\xb0\x64\x5e\xdb\xa0\x5d\xf1\x00\x00\xac\x0d\xe6\x19\xba\xe7\x40\xbd\x42\xc1\x36\xdc\xab\x31\x79\x61\x24\xdc\x89\x64\x95\x64\xc3\x5a\xd6\xdd\xdc\x58\x19\xdd\xa5\xb8\x34\xe3\xba\x95\x64\x3f\xac\x83\xda\x7a\x9a\x11\xf4\x47\x84\xaf\xd7\xdd\xdf\x2f\x2a\x43\xe6\xcd\x0a\xb4\xc0\xeb\xcd\x22\xb4\xc3\x3c\x9a\x9b\x54\x6d\xe8\xca\xd6\xd4\x31\x37\x4a\x7c\xa2\xfe\xac\x57\xa0\x07\x7b\x7b\xaf\x8f\xea\xbd\x0d\xb9\xeb\x6d\x05\x58\x27\x78\x3f\x2f\xff\xed\x17\x8b\xb5\x2d\xb9\xb5\xb4\x7b\xde\xfd\xb8\x22\xe9\xe6\xdf\x8d\xf1\x7e\xdc\xad\xb9\x1f\x77\xee\x2b\x67\x74\x5e\x70\x28\xef\xed\xff\x52\x57\xf1\x5f\x4f\x68\xf4\x29\x15\x5e\x6a\x19\x92\x55\xdf\x69\xa2\x19\x03\xa7\x4c\xfd\xd4\xb0\x0c\x1b\x94\x15\x75\x78\x06\xc9\x22\xe1\xb3\x12\x8f\xa5\x4a\xfa\x78\x0e\xfe\xd6\x30\x22\x58\xe8\x1c\x77\x65\x04\x26\x55\xf2\x29\x6f\xee\x3a\xa3\x66\x28\x79\x0a\x1e\x8a\x38\x5c\x40\xfe\xc3\x29\x14\x9c\xd0\xc8\x31\x9f\x10\x3c\xc5\xfa\x6b\xf3\x5c\xf2\x1a\x12\x6f\x14\xe2\xdf\x46\x14\x9c\x93\xb9\xff\x94\x4a\x54\xea\x46\x5c\x97\x42\xdf\xb6\xc1\xd2\x9c\x20\x3b\xad\xc5\x3d\x01\xdc\x25\xa4\xbe\xd0\x71\x71\x5d\x38\x77\xca\x82\x4b\xd6\xef\x3c\xc3\xbd\x4d\xbe\x46\xab\x99\xe0\x19\xe6\x2d\xbe\x35\xd4\x6f\x1d\x19\x9b\xc1\xd9\x91\xb1\x90\x2b\xc9\xfb\x54\x98\x3f\x95\x29\xe6\x52\xfb\x3b\x81\xfc\x65\x25\x4d\x65\x24\x1d\x39\x5d\xd4\x78\x91\x8b\x1f\xe3\x66\x5f\xe8\xcf\x9e\x7f\x8c\x4f\xfd\x8a\x7e\xf1\xf6\x63\xbc\xaa\xcb\xf0\xc5\xd3\x8f\xf1\x29\xee\xf0\x67\xbf\x9d\x78\x34\x8f\xfa\xf3\xb6\xf5\xa1\xf8\x74\x88\xde\xbe\x72\xdc\x4d\x5d\x05\x36\x5a\x22\x32\x78\x05\x88\x60\xd0\x76\x7b\x90\x59\x9b\x5f\x8d\x3b\x99\xeb\xc8\xac\x5f\x0c\x5a\x83\x0d\xac\x5a\x25\xa7\xc8\x4e\xe5\x5e\x41\xa9\x32\x7a\x9b\xb5\x90\xac\x89\x98\x03\xbf\x21\xf9\x6f\x08\x94\xf9\xd5\x4e\x72\xcb\x35\x7b\x12\x55\x10\xb2\x52\xc2\x37\x1a\xd6\x79\x17\x2e\x02\xca\x82\xe7\xad\x21\x23\xb0\x96\xf0\x2e\x5e\xb3\x82\xbc\xe5\x27\x2c\xa0\x99\x23\xec\x58\x91\xed\x29\xe8\x15\xde\xa1\x5a\x85\x08\xf0\x56\xe0\x1f\x38\x40\xe1\xe4\x8e\xbe\x9e\x7f\xfc\xf8\xee\x0e\xb0\x6a\xed\x2b\x52\x9d\x77\x61\x2b\x63\x01\x50\x56\xf6\xf1\x6f\x25\x20\x83\xbd\x74\x01\xca\x9f\x74\xfa\xda\x96\x17\x32\x78\xa4\xf9\x85\xd1\x90\xbc\x7a\x97\x04\x2c\xc2\x4b\x64\xbf\x92\xfd\xcb\xe2\x3e\x7b\xc6\xf0\x89\xb3\xe4\xea\xf3\x5b\x96\xb0\x57\x07\xc1\x3a\x84\xbf\xc0\xe7\xd3\x61\xf6\xd7\x86\x39\x7e\xfe\xbf\x00\x00\x00\xff\xff\x16\x77\xab\xea\xa0\x38\x00\x00")

func assetTemplatesThreadHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetTemplatesThreadHtml,
		"asset/templates/thread.html",
	)
}

func assetTemplatesThreadHtml() (*asset, error) {
	bytes, err := assetTemplatesThreadHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "asset/templates/thread.html", size: 14496, mode: os.FileMode(420), modTime: time.Unix(1446268623, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"asset/templates/thread.html": assetTemplatesThreadHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"asset": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"thread.html": &bintree{assetTemplatesThreadHtml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

