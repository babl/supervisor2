// Code generated by go-bindata.
// sources:
// data/server.crt
// data/server.key
// data/server.pem
// DO NOT EDIT!

package main

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

var _dataServerCrt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x96\x49\xd3\xa2\xcc\xb2\xc7\xf7\x7c\x8a\xbb\x37\x6e\x88\x28\x08\x8b\xb3\xa8\x82\x62\x10\x0b\x65\x50\x84\x1d\x93\xcc\xa2\xa2\x14\xf0\xe9\x6f\xe9\xdb\x6f\x9f\xbe\xdd\x27\xde\xf3\x44\x3c\x0b\x93\xac\x21\x33\xff\x99\xbf\xfa\xdf\xcf\x1f\x44\x9a\x61\xfd\x8f\x8c\x1c\xcf\x50\x0d\x19\x78\xe8\x6b\x65\xb0\x61\x20\xb1\x92\x65\x20\x3d\x72\x40\x0c\x08\x72\xc3\x05\xa3\xdf\x3c\xa5\x6e\x53\x35\x27\xbf\x2b\x5b\xef\xfe\x40\xed\x36\xbd\x5e\x31\x60\x35\xd9\x7d\x68\xae\x11\xaf\x15\x1b\x41\x99\x9c\x00\x83\x51\x37\xca\x33\xd8\xc1\xdc\x3a\x43\x10\x78\xa0\x39\x7b\xd8\x09\x88\x0a\x02\xe5\x6c\xdb\x26\x22\x2b\x1c\x5e\xec\x5b\x32\x41\x35\x6e\xad\x29\xbb\x40\x16\xbb\x98\x18\xf6\xf7\xbb\xc2\xa0\xf1\xfe\x87\x83\x81\xd4\x55\xaa\x15\x43\xd2\x36\x6c\xe6\xc2\x00\xcf\x20\x53\x09\x3b\x5a\x15\x98\xf1\x8c\x26\xac\x9c\xd6\x58\x81\xd1\xc7\xc6\x50\x23\x87\xab\xfa\xa7\x11\x43\x63\x44\x0a\x38\xfc\x75\x23\xec\x41\x6e\x57\x04\x2d\x79\x27\x5c\x4e\xf2\x1c\x95\xbf\x87\xc1\x00\xfb\x04\xc0\xc6\x80\x0a\x01\x1f\x07\x13\x74\x34\x0f\xb6\x4c\x7c\xe3\x1e\xc8\xcb\xf8\xd1\x1d\xca\x9b\xb2\x03\x78\x98\x47\x23\x1c\x14\xf6\xfd\x32\x94\xa1\x6a\xdc\x07\x2b\xfb\x99\x14\x12\xe6\xe4\xb4\x3e\x5f\xdd\x97\x2a\xeb\x5e\xe6\xb2\x39\x4a\x6b\x74\xde\x26\xee\xb5\x2a\xcf\xbb\xea\x74\x92\xf1\xdd\xda\xf6\xd2\x98\x3e\x8c\x9d\x7a\x7e\x20\x58\x04\x64\x8b\x22\x51\xed\x39\x03\xbc\xef\xa6\x4e\x73\xb0\x47\x1c\x9f\x25\x9a\x99\x93\xad\xbd\xea\x37\x1c\x3c\x6c\x03\xf9\x6e\x8c\x72\x8c\xa6\xe7\x6a\xa1\xca\x2a\xf2\x4a\x81\x55\x9f\x50\x43\xe7\xc5\xb1\xc0\x57\xf1\xe6\x0d\xb2\xb2\x74\x83\xe8\x2c\x31\xa2\x5d\xf4\xe3\x51\xb7\x6f\xd0\x7e\x9c\x51\x91\xbe\x70\xa8\xa5\x33\x3f\x18\x1e\xcb\xae\xaa\x63\x15\x2e\x16\x1b\x47\xba\xf3\xea\x13\xd8\x5b\xf1\x0d\x96\x65\xdc\x3c\xab\xe8\x3d\x6a\x4a\x77\x7c\x91\x81\xc1\xf2\xb3\x6e\x96\xda\xea\xb0\xde\x82\xcd\xbb\x36\xb3\x4c\x58\xf0\x92\xad\x4f\xcd\x7d\xb5\x68\xf9\xd5\x3c\xa5\xc9\xf6\x36\x7b\xbe\xc3\xa1\x55\x79\xdc\x09\xdb\x28\x5a\x3e\xcc\x62\xde\xdf\xa3\xcb\x48\x16\x8c\x55\xf6\x03\x1b\x44\x26\x8c\x9e\xac\x3e\x97\xd7\xa8\x58\xdb\x5b\x70\x5d\xcf\x69\x33\xae\x87\x52\x6e\x96\x7c\x33\x5d\xf7\xf9\xe1\xc4\xc7\x20\xc7\x10\x00\xad\xca\x73\xc3\xa0\x02\x94\x21\xad\x16\xf3\x29\x97\x6e\x8b\x10\x5c\x45\x44\xcb\x26\xc3\x08\x10\xfd\x23\x12\x87\x6d\x20\xfc\x0a\xca\x30\x49\x00\xa1\x7d\xd2\x01\x41\x9a\x2c\xf7\x1a\xad\x9c\x0a\x09\x96\x31\x20\x1a\x03\x56\xa7\x14\x11\x04\x97\xc4\xa6\x86\x9f\x8b\x0f\x74\x31\x52\xf7\xbe\x9a\xa0\x47\x7f\xdc\x16\xee\xb6\x40\x47\xd7\x1d\xb0\x72\x5c\xd2\x8a\xf1\x0f\x0c\x45\xed\xb3\x96\x31\x88\x1d\x60\x7a\xac\x6a\x16\x66\xd4\x62\x74\x4d\x39\x81\x7f\x21\xfe\x25\x84\x6a\x98\x2d\xe7\xc7\x61\x2a\xb0\x0e\x7e\x39\x19\x41\xa8\xd9\x24\xa8\xc0\x00\xf3\xfc\x49\x43\x40\x2a\xb4\x13\x02\x82\xa0\x8a\x74\x87\x4d\x94\x6e\xd8\x73\x52\x95\xac\xc1\x3b\xf2\x79\x76\x7f\xc9\xe7\x7d\x3b\x36\xa9\x6e\x35\x7f\x4b\x7d\xdf\x4a\x53\x38\x89\x64\x4f\x02\x83\xf9\x3b\x3e\x0c\x4c\xcd\xe0\x0a\x36\xd5\x81\xb0\x9f\xa4\x2a\xbc\xec\xa8\x63\xf3\x4e\xe5\xd5\x06\x4f\x7c\x1f\x5e\x9c\x39\xa4\x92\x4b\x6e\x0d\x49\x65\x9e\xf6\x49\x32\x50\xd9\xff\xc8\x81\x6d\xef\x31\x68\xe5\xff\xa7\xfb\xeb\xe6\x1b\xa3\x01\x5c\x28\x55\x32\x14\x67\x60\xc0\x3c\xbc\xe5\x36\xa1\x2d\x8f\x48\x9e\x05\x9a\x3c\xf5\x74\x03\xdb\x86\xf9\x4b\xf4\x00\x6d\x0c\x6c\xe8\x3e\x96\x83\x5f\xe2\x35\xa0\x5a\xdc\xbb\x54\x77\xc8\xa1\x14\x87\x60\x0d\xff\x08\x87\xf9\xc6\x33\xcb\xf0\x41\x7e\x29\x57\x6e\x90\x3c\xdc\xe0\x3c\x7c\x9d\x22\xad\x99\x0d\x64\x35\xc9\xcd\xb9\x87\x6d\x53\x05\x17\xa7\x31\xb4\x55\x41\xdb\x7c\x88\xdb\x91\x67\x0c\x6d\xd7\x18\xfa\xae\x89\xb5\xa6\x09\x65\xb8\x4a\x34\xe9\x4d\x6d\xbc\xa1\x52\x9b\xde\xdc\xe3\x36\xc9\x4f\x9a\x3a\xa5\xf4\x3b\x9d\x19\x45\xdc\xda\x79\xcc\xf1\x3d\xdd\xe0\x1e\x97\xb0\x60\x02\xce\xa2\x09\x71\xe8\x07\x8b\x6e\x94\xde\x53\x2d\xcf\xe9\x30\x69\xfe\x38\x55\xa5\x27\x6a\x4d\x45\x17\xb6\xf1\xfa\xfc\xa6\x87\x15\xa9\x0c\x3b\xe6\x13\x5f\x32\x7f\x4a\xf7\x9f\x4a\x45\xff\x2f\x70\x48\xb8\x86\x8d\xd7\x3b\x7e\x3f\x03\x0b\xe6\xf5\xa3\xa8\x4b\x4d\x22\x2c\x9d\x1b\xbd\xca\x00\x70\x90\x69\x02\x81\x91\xbb\x6e\x50\xc4\x71\xff\xaa\xb9\xa4\x69\xde\xc7\x41\xce\x9f\x85\x17\x56\x23\x9a\x1c\xff\x1d\x6d\x76\x4b\x61\xe8\x2e\x78\x5a\xc8\xfb\x08\x9c\x0c\x49\x78\x2c\xc2\x5c\xe1\x19\xb3\xf1\x97\xd5\xe0\xdc\x8c\x4e\x76\xa4\x34\xb1\x7b\x7f\xc8\xf0\xcb\x21\x0a\x7e\x84\x8f\x6c\x1b\x16\x21\xb1\x05\x76\xe5\x8e\xaf\xf8\x51\xa0\x80\x6f\x35\xc5\x54\x94\x57\x00\xe3\xa5\x89\xab\x2b\x60\x8c\xe1\xa8\x8e\xa9\x7c\x77\x8b\x26\x18\x97\x86\x62\xbd\xea\xe1\x95\xcf\x2b\x9b\xac\xfc\x40\xcb\x96\x8a\xb4\x31\xe4\xfa\xf1\x38\x6c\xde\x6b\x21\x34\x7b\x24\x1d\xa0\xb4\xf7\x9a\xe2\x5a\x5c\x42\x64\x9d\x98\x9b\xe8\x19\xb5\x1c\x94\xd7\x4b\x15\x10\xb9\xb2\x4d\x11\xa7\xbb\x72\x9c\x9f\xc2\x8a\x48\x46\x4a\x12\xd8\x7b\x3c\x99\x40\x26\xd5\xf3\xab\xcb\x2b\x2d\x5d\xde\x76\x0b\xa2\xbe\x9c\x32\x6b\xfd\x63\xc4\xe4\xcf\x06\xe6\x4f\x2b\x8b\xa3\xe5\xa2\x51\xab\x03\xa7\x65\xe5\x3b\xe2\xe5\xd7\x6d\x71\xd3\xcc\x63\xfc\xc0\xe1\x74\x6f\xa3\x03\x3c\x72\xe6\x8d\x97\x51\xed\x58\xcd\x35\xeb\xc1\xe6\x70\x55\x38\x22\x36\xcc\x79\x25\x77\xea\xba\x8d\x83\x55\xe8\xae\x9e\x52\xbb\x51\x0c\xcf\x30\xe6\x0e\x3c\x9d\x32\x04\xff\xfa\x17\xf3\x05\x18\xb2\x94\x3f\xa1\xf6\xdf\x80\x57\x7f\x80\xb7\xfe\x09\x3c\x5b\xce\x69\xdb\xe6\x00\x80\xf3\x41\x4d\xb8\x6e\x5f\x64\xd3\x4d\x56\xfe\x2c\x2c\x50\xc0\x92\xc1\xae\x4d\x8c\xfc\x07\xdc\xc6\x17\x8a\xfc\x8f\xc4\xd4\xde\x50\xad\x7b\xc8\xf1\x45\xaa\x9f\xa7\xd0\x85\xa7\xe4\x76\x9e\xa9\x9c\x94\x78\xda\x8c\xea\x0c\xce\x7f\xe3\x88\x51\x6a\xc7\x3b\xcb\xd0\x8d\x39\x89\x82\xce\x82\x86\x9a\xcf\x18\x6e\x2e\x8a\x87\x38\xac\xe0\xd1\x9a\x11\x67\x29\x80\xb5\x9a\x8e\xda\x8c\xf1\x37\x1b\x61\xdc\x0a\xd1\x06\xaf\x3f\x0d\x4d\x59\x82\xe5\xf3\x19\x8f\x6a\x05\x4e\x7f\x9d\xd0\x79\xca\x89\xea\x56\x4e\x69\x9b\x9d\xdf\x01\xd5\x68\xa2\xdb\xa3\x31\x83\xe2\xe7\x0d\xb4\xfa\x0f\x87\xdc\xbe\x9c\xd9\x48\x93\xa6\xe8\xe2\xf0\xdf\x1b\x19\x06\x34\xaa\xdf\x73\x80\xd4\x8f\xb6\xe9\x94\x10\xc1\xc7\x41\xce\xcd\xaf\xd0\x6f\x16\xc6\xa2\xfa\x6c\xf6\x75\xb6\x4e\x1a\x76\x9d\x6f\xad\x2e\x98\x95\xc7\x6a\x3e\xb5\x9a\x7b\x29\x86\x78\xb3\x12\x2f\xb2\xbb\xdf\x66\x1b\x97\x45\x2a\xf3\x10\xda\xcc\xb2\x8b\x60\xbb\x47\x8f\x51\x2b\x75\x59\x38\x56\x69\xe6\xb5\xa2\x90\x96\x49\x7c\xe7\x73\x1f\x5c\x57\xbc\x16\xdd\x96\x47\x3b\xd3\xd2\x71\xd2\xea\x43\x13\xea\xc7\xe5\x3b\x0a\x05\x1f\x88\x8c\x8b\xc7\xc5\x54\xaf\xd6\xa8\x74\x53\x67\x7c\x45\xc2\xf6\xd6\xeb\x55\x02\xf4\xdd\xd4\x67\x42\xa2\x0a\x3d\x6f\x0a\xdb\x15\xe4\xbd\x28\x78\x27\x83\x14\x7b\x93\x1f\x59\x62\x65\xd6\xa6\xad\x18\x21\xcb\x84\x62\xb1\xbc\x87\x8f\xcd\xa9\x45\x27\x34\x4b\x8d\x10\x98\xfa\x24\x0d\x82\xd2\xc4\x5c\xd1\xdd\xe6\xc2\x5b\x5c\x8a\xc7\x82\xac\xe1\x73\x88\x08\x77\x56\x6f\x6b\x64\x0a\xb0\xe9\xef\x35\xb2\x6e\x3e\x60\x22\x61\x34\xc5\xf1\x6d\xbb\x97\x7c\xe8\xee\xe1\xd1\x2c\x41\x63\xda\x9e\x96\x62\xc5\xc6\x09\x77\xc4\x5e\x79\x56\x9f\x8f\x0e\x6f\x0b\x45\x8c\xd1\x95\xcc\x70\xd9\xdd\xea\x11\xcd\xec\xcb\x1a\xaa\x8a\x59\x1e\x8d\x39\x7a\xd6\x3c\x4e\xfc\x61\x34\x58\x4b\xf7\x6d\x1f\x0b\x4f\xa1\x90\x5b\x6e\x05\x06\xc0\xe9\x6b\xa5\x26\x86\x02\x6c\x00\x3b\xfa\x3c\xb9\x7a\x54\xb5\x97\x9a\xa0\xfc\xcb\x35\x8f\x56\x41\x5f\x52\xfd\x12\x8a\x1d\x0a\x3d\x2a\xe5\x1f\x8f\x9e\x5f\x29\x4a\x51\x44\xae\xbf\x8e\x62\x1b\xa1\x64\x86\x54\x54\x94\x16\x3f\x07\xfb\x0c\x60\x51\x86\x3f\x07\x7b\x74\xb1\xe8\xa0\xe3\xd9\xaf\x84\xb5\xa6\xa6\x83\xef\x07\xc3\x9c\x26\xbe\x39\x53\x7a\xb1\xe8\x30\xb4\x06\x26\xf6\xc0\xf6\x43\xbe\x9f\xe0\xeb\x82\xe1\xdf\xe0\x53\x49\xa2\xe3\x3f\x16\xc5\xae\x34\x7d\xa4\x9f\x4c\x52\xcd\x24\x6b\xe7\xfb\x23\xe0\xd4\x0f\xd8\x88\xc5\x61\xa2\x93\x6f\x7c\x15\x84\x39\x51\x3b\x70\x1a\xcd\x6c\xbc\xeb\x7d\x9f\x5c\x9f\xf1\xe6\xf4\xb6\xd3\xeb\x12\xa9\xbe\xac\x96\x0e\x20\xcc\x19\x7c\x9d\x73\x88\x58\xe2\xfe\x06\xb5\x23\x09\xf6\xdf\xb8\x29\xb2\xd7\x23\xfe\xc8\x97\x60\x05\xbc\x7f\xb9\x31\x95\xb2\x13\x94\xff\xbe\xb2\x45\x92\x89\xff\xde\x88\xb2\x7a\xfc\x8f\xac\x9e\x81\xf8\x4d\xb2\x23\x22\xcb\x03\x33\x83\x15\x2d\xdf\x0b\xe0\x55\x94\xaf\x5f\xc8\xb8\xeb\x29\xaf\x6b\xca\xe8\xbf\x92\x28\xf3\x55\xcc\xb1\x83\xa3\x5a\xa7\x53\x2d\x1d\xcf\xb4\xed\x7d\xc5\x52\x4e\x35\x79\x33\x1f\x67\x0c\xd9\x2f\x9f\x95\xdc\xf6\x21\x74\x3b\xb7\xbd\x57\x50\x5f\xa7\xef\x77\xec\x1c\xe2\xac\x75\xfc\xcb\x20\x0a\x55\xdf\x79\xff\x80\x9b\xb5\x77\xbc\xa0\xab\x55\xf9\x4a\x95\x6a\xf0\xb2\x95\xcf\xfe\x22\x6d\x22\x3e\x41\x65\x13\xd1\x79\x94\x89\x74\xa4\xef\xf6\xa3\x5f\x48\x26\x2a\xeb\xf5\x4e\x77\x1c\x5d\xdb\x75\xcc\x1b\x73\xe7\x44\xbb\x36\x92\xe0\x8a\x5e\x59\xd0\xb7\xc5\xd0\x3d\xbb\x2c\x15\x5e\xa5\xe0\xd3\xf7\x6a\xfb\x9a\xc9\xda\xef\xd2\xe8\x95\x2f\xce\xd3\x21\xbb\x17\x1b\x14\xdc\x9f\xcb\x15\xb9\x98\xaf\x51\x5c\x32\x04\xdc\x8d\x61\xe7\x92\x57\x7b\x2e\x37\x58\x3d\xf1\x11\x7e\x3c\x5d\x05\x09\x59\xb4\x5d\xe3\x8a\x7b\x25\x78\xea\xf8\x0a\xa7\x42\xd5\x66\xfe\x49\x37\xc5\xbe\x5b\x56\x9d\x7f\xea\xf4\xc3\x29\xa7\x39\xb8\x6c\x8e\xdd\xca\x0e\xe6\xc5\x3a\xed\xe7\x5a\x79\xe0\xfb\xe6\x5a\x37\x23\x24\x17\xa7\xf7\x57\xac\x79\x99\x69\x3f\x85\x8b\xfe\x70\x04\x43\x36\x8d\xe5\x2d\x6d\xab\xda\x17\x1b\x6d\x5a\xd8\xbd\xd3\x68\xcc\xf1\x1a\x2e\x34\x21\x14\x8a\x6d\x5b\x65\x2d\x1b\x2c\x4a\xbf\xa9\x83\xe4\xbc\x39\x1a\xfe\x7e\x55\x12\x58\x8a\x7d\x24\xc7\x9a\xcb\x57\x16\x77\x17\xf1\xe2\xb2\xb0\xb7\x27\xcb\xa4\xdc\x39\xc4\x6b\x4b\x60\xcc\xc3\xa3\x7e\xb4\xfc\xd6\xd3\x69\xdf\x65\xf4\xb9\x5f\xbb\xb7\x42\x58\x2a\x96\xfa\x66\xed\xfc\x9f\x70\xf3\x7f\x01\x00\x00\xff\xff\x2a\xb1\x7c\xf3\x5f\x0d\x00\x00")

func dataServerCrtBytes() ([]byte, error) {
	return bindataRead(
		_dataServerCrt,
		"data/server.crt",
	)
}

func dataServerCrt() (*asset, error) {
	bytes, err := dataServerCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/server.crt", size: 3423, mode: os.FileMode(420), modTime: time.Unix(1465410894, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataServerKey = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xd3\xb7\xb6\xab\x46\x00\x85\xe1\x9e\xa7\x38\xbd\x96\x97\x24\x40\x81\xe2\x16\x33\xe4\x9c\x41\xd0\x91\x73\x90\x10\x0c\xc3\xd3\xdb\x3e\xf5\xdd\xed\x6e\xfe\xe6\xfb\xe7\xff\x41\x5e\x94\x8d\x1f\xc7\x05\x3f\x96\x23\x07\xc0\xe3\x7f\x54\x3e\xfa\x7d\x08\x5d\x96\xd9\x10\xc9\x10\x00\x15\x56\x36\x57\xa3\xdd\x1f\x55\x16\xf5\x6e\x12\x02\x74\x98\x90\xd4\x5c\x29\x28\x2a\x45\xda\x1e\x1c\x17\x0e\xcd\xe8\x55\x34\x3a\x34\x6d\x29\x4f\x9a\xfd\x1c\x63\x22\x9d\xdb\x72\xbc\x6d\x95\x7c\xec\xec\xea\xd4\xb4\x33\x33\x76\xa4\xdf\x84\xba\xfc\x28\x2f\x06\x65\x51\x82\xac\xb3\xe4\xa5\xca\x63\x7e\x68\x81\x6d\x92\x76\x04\xac\x53\xd2\xe9\x9e\xb4\x57\xea\xaa\x13\xf0\x08\x82\xd3\x89\x42\xa1\x37\xa9\x65\x10\xfb\xad\xf0\x7a\x8e\xa5\x57\x8a\xba\x88\x42\x20\x29\xdc\xd8\xf3\xe2\xe8\xd3\xdf\x9e\xf1\xd7\x76\x62\x03\x9a\x37\x94\xef\x21\x4c\xb6\xcc\x01\x1b\x40\x02\x4c\x22\x50\x4e\xf7\x7a\x3e\x8c\x0f\x49\x63\xcf\x8e\xbf\x42\x68\xcf\xdc\xcc\x63\x61\xee\xf3\x5c\x55\x74\x73\xe7\x70\x42\xb9\xcc\x3c\xab\xd4\xe6\x85\x1f\x59\xf6\x02\x72\x37\xb2\x35\x43\x0f\x99\xc8\x0a\x2f\xcd\xf1\x47\x5c\x70\xbb\xf8\x97\xb3\x94\xb3\x99\x5c\x32\x78\x20\xdb\x72\x10\x35\x1f\x0d\xb8\xee\xbf\x6a\x80\x2e\x76\x94\x09\xf0\xf2\xd2\xba\xec\x32\xca\xb7\x68\xe3\xa3\xa9\x08\xb8\x8a\x98\x86\x58\x7e\x8d\x57\x9e\xe2\xc3\xd3\xe2\xca\xa1\x9b\x76\xfa\xca\xa4\x11\xd9\x2d\x5f\xf5\xe5\x90\x7e\x3c\xe8\x55\xc8\x7d\x07\xc8\x5b\x7a\xc2\xb3\xb6\xcd\xdd\xb1\x47\x7b\x20\xd6\x6f\xb5\x08\x53\x22\x6b\xa0\x5a\xe9\xbe\x75\xb7\x90\x59\x5b\xb5\xa9\x5b\xf2\xd6\xb4\xe6\xed\x02\x9a\xf5\xde\xac\x6c\xf0\x7c\xf6\x57\x5b\xc6\xd4\xb3\x0a\x83\x7a\x6f\x8d\xcf\xfb\x6e\x51\xf4\xbd\xa5\x65\x58\x9d\x3a\x48\x30\x49\x8f\x66\xd6\xe4\x40\xc7\x83\xdb\xe8\x8e\x03\xd3\xdd\x71\x17\x15\x76\x68\xb8\x97\xd2\x08\x1b\xa7\x19\xd9\xbc\x27\x29\xf0\x58\x17\x4e\x74\x57\x55\xed\x87\xcb\x95\x69\xa6\x41\x39\x3b\x55\x4e\xe8\x2a\x67\xce\x97\xf7\x99\x34\xa7\xe5\x9b\x7e\x0b\x33\xd4\x49\xdd\x11\x4a\xba\xd5\x25\x0a\xe1\x80\x45\x0a\x28\x41\xde\x2a\xcf\xe5\x98\x0c\xd5\x73\xdc\xa4\x72\xd3\x9a\xd9\x42\x5c\x8d\x90\xdc\x08\x19\xe2\xec\xde\xd3\x9e\xb7\xc6\xb6\xe2\x88\x07\xbb\x15\x49\x79\x30\xfd\xb4\xad\x90\x0a\xa7\x9a\x05\x30\xb7\x9d\x9c\x59\x3b\xf6\xa5\x91\xec\xbc\xf0\xf3\xfb\xe8\x92\xd2\xb4\x15\xa0\xf8\x6d\x46\xf8\x6e\x91\x73\xfd\x9b\x3a\xc4\x18\x45\xfa\x35\x42\xcf\xfc\xe9\xac\xcd\x07\xfa\x42\x6c\xbc\x95\xa2\x8f\x92\xe6\x34\x3a\x11\xdb\x73\x9d\x13\x08\x55\x7a\xc3\xdd\x32\x45\x38\x4b\xdf\x37\x6f\x17\x27\x22\x28\xde\xaa\xb9\x55\xd6\xdc\xd2\x4e\x68\x49\xbd\xf6\x5f\xb2\xe8\xeb\x1d\x95\xbe\x3d\xe6\xba\x43\x9f\x1d\x35\x67\x39\x6f\xe5\xc4\xce\xd2\x5c\xdd\x0b\x8c\x6c\x0f\x72\x20\xb8\xef\x1d\x3e\xe8\x84\x70\x24\x4a\x7e\x9c\x4f\xb8\xa5\xe2\xdd\x21\x95\x29\x94\x2a\x1f\xc5\x8f\x3e\xee\xae\xc1\x58\x97\xa5\x80\x8b\x87\x0b\x5f\x38\xa9\xfe\xfc\x21\x7e\xf9\xf0\x06\xf7\x77\x56\xff\x06\x00\x00\xff\xff\x8c\x8d\x56\x78\x77\x03\x00\x00")

func dataServerKeyBytes() ([]byte, error) {
	return bindataRead(
		_dataServerKey,
		"data/server.key",
	)
}

func dataServerKey() (*asset, error) {
	bytes, err := dataServerKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/server.key", size: 887, mode: os.FileMode(420), modTime: time.Unix(1465410894, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataServerPem = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x93\x49\x93\xb3\x36\x10\x86\xef\xfc\x8a\xdc\xa7\x52\x16\x78\x19\x7c\xc8\x41\x42\x62\x33\xc2\x23\x2c\xc0\x70\x63\xb7\x59\x8c\x6d\x66\x2c\xf0\xaf\x8f\x99\xc9\x97\xaa\xe4\x53\x95\x2e\x6f\xb7\x5a\x5d\xcf\xdb\xfd\xe7\x7c\x10\x31\x2c\xf7\x0f\x8d\x78\xdc\xd2\x2d\x0d\x72\xf2\xad\x4a\xd4\xb2\xb4\xee\xa9\x69\xb0\x3a\x54\x50\x58\x08\x56\xaf\xfb\x84\x2e\xaa\x9a\xdb\xa9\x39\x1b\x5b\x01\x10\x64\xbe\x0e\x31\x0a\x29\x1b\x84\xc6\x22\x1c\x30\x66\x10\x61\xa3\x80\x13\x2e\xd1\x57\x61\x28\xfb\x44\x83\x62\xe7\x2b\xdb\xcf\xf8\x20\xf3\xdc\xd0\x41\xcc\xc9\x89\x22\xf5\x27\x56\x89\xe8\x10\xae\x41\x7c\xb4\xbf\xe2\x23\xab\x02\xa5\x6d\x62\xa5\x05\xd9\x84\x98\x94\x9b\x4d\xc5\x4d\xaf\xa1\x4c\x15\xf8\xa7\x3a\xc6\x30\x7e\x25\xbb\x20\x52\x88\x30\x4f\x99\x4b\x39\x13\xee\xd3\x9a\x28\x8e\x04\xc5\xfe\x32\x9c\xb5\x7a\xd6\xc8\x5a\xfa\x57\xac\x51\xf3\xbf\x16\x03\xff\x49\x3c\x0a\xff\xe9\x02\x51\xeb\x10\x8e\x43\x12\xae\x1f\xc9\x91\x8e\x04\xc3\x3d\xaa\xdc\x40\x42\x30\xe3\x08\xb8\x7d\x12\xba\xa7\x58\x51\x47\x1d\xc3\xc3\x1c\x40\xb0\xe7\x1a\xc8\x1f\xa9\x92\x0f\xf1\x01\xd9\x69\x47\xbf\xa8\xd7\x0b\x03\xfe\x74\xa9\x23\x72\x73\x2e\x5e\x2b\x65\x4b\xf6\x15\x2b\xdb\x47\xac\x8c\xad\xd3\xb9\x8f\x94\x6b\xe8\xf2\x1b\x44\xa2\xc3\xd7\x87\x35\x83\xa2\x8a\x1a\xad\x8a\x08\x5c\x65\x34\xb0\xa7\x6a\x90\x6c\xbf\x6b\x2b\x4a\x9f\x55\x7e\x06\x27\xf9\xd8\x6b\xde\xfb\x9b\x10\xf2\xb5\xbf\x82\xd5\x9e\xd2\x69\x7a\x37\x17\x67\xf0\x61\x2b\xe1\x26\x5a\xae\xdf\xd2\x95\x46\x55\x76\x6f\xa2\x82\x24\xa5\x6f\xe0\xbd\xe4\x45\xc7\xcd\x14\x2c\x32\xd3\x30\x06\xbc\x18\x81\x32\x15\x9b\x62\x92\x7d\xfc\xd9\x18\x10\x2f\xba\x2b\x26\xa3\x1a\x69\xf7\x1a\x66\x72\x50\x3e\x3e\x55\xbd\xd9\x68\x17\x39\x0c\xac\x31\x58\xd8\x4b\x20\x2d\xc7\x1a\x0d\x7a\x85\x26\xb6\x5e\x7b\xe8\x2a\xef\x9d\xb8\x0c\x9c\x7e\x03\xdb\x02\xe1\x43\x9a\x8d\xc9\xec\x30\x81\x30\x71\xef\xd4\x68\x7e\x20\x7b\x80\x23\x68\x09\x88\xa1\x23\xcd\xc0\x4c\xa6\x12\x04\xa9\x86\x0a\x28\xb8\xf8\x4e\xf0\x10\xa9\x84\xd7\x5b\x6c\x77\x5e\x7f\x5b\xea\x74\xbf\x80\xae\xbb\xac\xb7\xa2\x5c\xd1\xe7\xb9\xd8\x48\xa9\xb2\xbd\xce\x34\xff\x0b\xd3\x6e\xab\x93\x75\xfb\xa5\x17\xe1\x56\xce\xcd\xe0\xfc\x7a\x5c\xa7\x8a\x62\x22\xaa\xf5\x90\xd1\x79\x6e\x6c\x69\xd7\xc7\xd6\xe9\x91\xb9\x33\x69\xc4\x20\x9e\x09\x53\xc5\x3a\x03\xa7\xe2\x46\x6a\xcb\xf5\xaa\xbf\x1d\xb7\xe9\x18\x64\x63\xf7\xb6\xf0\xd6\x51\xa9\xf6\x67\x90\xc4\x37\x6e\xb7\x17\x27\x3a\x48\xe2\x98\xa1\xa9\x19\xf9\x28\xab\xf2\xf0\x1e\x4e\xa5\xbd\xda\x86\xe2\x1e\x1d\xfb\x77\xf5\xc9\x31\xbc\x94\x72\x97\x80\xf2\xe3\xb6\x2c\x56\xdd\x75\xb8\x3e\xa6\x4b\xee\x9c\xe4\xe4\x6d\x9f\xdc\xcd\x44\x2e\xa4\x84\x03\x6c\x59\x51\xf3\x7e\x7b\x19\x23\x47\x59\xd0\xd6\xa3\xb2\x9b\x9c\x1c\xdc\x65\x84\x60\x79\x17\x93\x91\x90\x8f\xa0\xb0\x59\x10\x85\xc9\xde\xf6\x7c\x45\x14\x8b\x1d\x5e\xf5\x75\xb9\x1d\xfe\x92\xbe\x77\x92\xb8\xf8\xf7\x3d\xfd\x3b\x00\x00\xff\xff\x2a\xc7\xc8\x57\xc4\x03\x00\x00")

func dataServerPemBytes() ([]byte, error) {
	return bindataRead(
		_dataServerPem,
		"data/server.pem",
	)
}

func dataServerPem() (*asset, error) {
	bytes, err := dataServerPemBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/server.pem", size: 964, mode: os.FileMode(420), modTime: time.Unix(1465410894, 0)}
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
	"data/server.crt": dataServerCrt,
	"data/server.key": dataServerKey,
	"data/server.pem": dataServerPem,
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
	"data": &bintree{nil, map[string]*bintree{
		"server.crt": &bintree{dataServerCrt, map[string]*bintree{}},
		"server.key": &bintree{dataServerKey, map[string]*bintree{}},
		"server.pem": &bintree{dataServerPem, map[string]*bintree{}},
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

