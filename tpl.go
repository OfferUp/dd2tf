// Code generated by go-bindata.
// sources:
// tmpl/monitor.tmpl
// tmpl/screenboard.tmpl
// tmpl/timeboard.tmpl
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

var _tmplMonitorTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x94\xcd\x6e\xdb\x3a\x10\x85\xf7\x7a\x8a\x81\xd6\xf7\xfa\x09\x92\x45\x91\xb8\x88\x17\xb5\xd1\xc0\x40\x16\x45\x41\x10\xd2\xd8\x22\x42\x93\x0e\x45\x59\x30\x58\xbe\x7b\xc1\x3f\x4b\x54\xe8\xd4\x2b\xf1\x9c\x33\x9f\x46\xa3\x91\x15\xf6\x72\x50\x0d\x42\xdd\x52\x4d\x5b\x79\x24\x27\x29\x98\x96\xaa\x86\xda\x18\x58\x6d\x5a\xb0\xb6\x06\x53\x01\x08\x7a\x42\xc8\x7f\x8f\x21\xb4\x75\x8e\xb5\x75\x05\xa0\xaf\xe7\x3b\xa1\xbd\x73\x42\xc8\x98\xff\x81\x1d\x60\xb5\xa7\xc7\x1e\xac\x75\x65\xee\x6a\x59\xf6\xcb\x18\x45\xc5\x11\x43\xd0\xda\xda\x98\x95\xb5\xf5\x7f\xc6\xa0\x68\xad\xfd\x1d\x49\x28\xda\x00\x39\x61\xdf\xd3\x23\xe6\x90\x87\x87\xf5\x6e\x5f\xb9\x06\x7e\x44\xdb\xda\xca\x49\x00\xd8\x37\x94\x53\xcd\xa4\x20\xa9\x74\x9e\xdf\x9d\x9d\xd5\xaf\xd6\xb7\xd8\x82\x50\x01\x7c\x0c\xa8\xae\xf0\x08\x2e\xff\xd3\x5f\xff\xf1\xd8\x33\x3e\x75\x54\xd1\x46\xa3\xf2\x4f\x18\x5b\x1d\x99\xee\x6e\xe0\xd0\x74\x9a\xc5\x56\x6a\x76\xb8\x6e\xe5\x33\xd5\x34\x38\xc2\x2b\x44\x48\xe2\x5e\x4d\x78\x1a\x3f\xed\x4f\xc9\x6c\x0a\x09\xf8\x8a\x01\xb0\x11\x1a\xd5\x85\xf2\xe0\xaa\xa8\x12\x96\xe4\x00\x2d\xa7\x8b\xe0\x70\xff\x6f\x43\xcb\x74\xd6\x28\xf5\xca\xbc\xc7\x59\xa6\x48\xda\xb3\x13\xca\x41\xbf\xc4\x1d\x08\x27\xd2\x45\x46\xee\x16\x01\x1b\xd1\xf0\xa1\xc5\x69\x8f\x58\x10\x88\xdf\xa7\x80\x99\x67\xbe\x1c\xd7\xc7\xc0\x14\x7e\x1f\x38\x7f\x63\xa2\x95\x63\x9a\x97\x97\xc9\x61\xe0\x9c\x8c\xc1\x48\x13\x2b\x16\x94\x47\x86\xe3\x8b\xec\xf5\x33\x72\x7a\x8d\x33\xc3\x91\x74\xb2\xd7\xa4\xf5\x5a\x9c\xda\xa7\x58\x91\xb6\xbe\x50\x3e\xf8\x85\x9c\x25\xf1\x26\x66\xc8\x62\x76\x46\x9d\xbf\x8d\x4e\x61\xdf\x49\xde\xce\x36\x33\x2c\xec\xd2\xd1\xd3\xd9\xfd\x31\x4c\x88\xdd\x7b\x08\x00\xc8\xf7\xd8\xc0\x24\x65\xcf\x32\xd5\xbc\x51\x25\x98\x38\x26\x79\x8c\xc7\x50\xbd\x30\xbf\x46\xbc\x62\x23\x2f\xee\x1b\xcc\x51\x44\x25\x3d\x63\x2e\xd3\x77\xd8\x4f\x8a\x69\xd6\xa4\x0f\x02\xa0\x49\xe7\x00\x5b\xda\xff\xa0\x2c\x6f\x9a\x68\xcb\x1e\xef\xe5\x33\x7c\x71\x41\x4a\x87\xca\x56\x7f\x03\x00\x00\xff\xff\x9e\x63\xc6\x89\xe8\x05\x00\x00")

func tmplMonitorTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplMonitorTmpl,
		"tmpl/monitor.tmpl",
	)
}

func tmplMonitorTmpl() (*asset, error) {
	bytes, err := tmplMonitorTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/monitor.tmpl", size: 1512, mode: os.FileMode(420), modTime: time.Unix(1597095598, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplScreenboardTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x5a\xd9\x52\xe3\x3c\x1a\xbd\xe7\x29\x54\xb9\x9e\xe6\x7f\x02\x2e\x68\x96\xbf\xa9\x82\x6e\x86\xd0\xf4\x2c\x35\xe5\x12\xf6\x17\x47\x85\x6d\xa5\x25\x39\x24\x64\xf2\xee\x53\xda\x25\xc7\x96\x45\xd5\xfc\xdc\xc4\x3a\xe7\xe8\x68\xdf\x61\xc0\x69\xcf\x4a\x40\x8b\x0a\x0b\x5c\xd1\xba\xe0\x25\x03\xe8\x5e\x29\x66\xd5\x02\x2d\x0e\x07\x74\x7e\x57\xa1\xe3\x71\x81\x0e\x67\x08\x09\x22\x1a\x40\x17\x1a\x7f\x56\x81\xe3\x71\x71\x86\xd0\xe1\xf0\x05\x91\x15\x3a\x7f\x02\x5c\xfd\xe8\x9a\x3d\x3a\x1e\xcf\x10\x62\x80\xab\x82\xca\xe0\x05\x92\x31\x62\x56\xc6\x81\xae\xf2\x01\x69\xb0\x5c\x63\x06\x06\xe3\xfa\x5b\xc7\x0d\x89\xd1\x98\xcf\xd0\x6e\x1a\x2c\xe0\x05\x33\x82\x5f\x1b\xe0\x9e\x7e\x27\x62\x9d\x14\x30\xdc\xd5\x80\xce\x35\x20\x8c\xae\xd8\x1a\xa1\x2a\x3a\x42\x1d\x6e\x41\xfe\x9a\xe2\x7f\x97\x41\x5d\x7a\x84\x36\x0c\x56\x64\xe7\xb8\x47\x1d\xb4\x6c\x05\x2b\xdc\x37\xc2\xb2\xd7\x26\xa8\xe9\xd1\x02\x8d\x05\xa2\xc2\xfc\x22\x55\x0d\x62\xaa\x08\xef\x8a\x35\xf9\x16\xfb\x8d\x6f\x33\xf9\x6d\xb3\xb5\xb3\xe8\x3f\x1c\xb4\xb7\xd0\x3f\x1d\xe4\xea\xd7\x34\xb7\xf6\x3c\xe9\x08\xcf\xb0\x13\x51\x1c\x57\x84\x81\xc5\x65\x43\xea\x2e\xf2\x29\xb0\x82\x42\x37\x2b\x9a\xb7\x5b\x92\x8f\x38\x57\x05\x97\x48\x68\x66\x24\x49\xaf\x6f\x40\xea\xb5\xb0\xe8\x5a\x87\x8c\x89\xe3\x92\x0e\xbf\x48\x25\xd6\x16\x7c\x57\x01\x13\xdf\x32\xe9\xc2\xe8\xfa\xd3\xe5\x90\xdf\xb6\x04\x19\xf5\x7a\x45\x1b\xca\x2c\x58\xaa\x80\x89\x6d\x99\x99\x7a\x6c\x21\xc4\xcc\x70\x09\x50\x21\xbf\x2f\x4c\x7f\x42\xa8\x21\x5b\x28\xf8\x06\xbb\x36\xbb\x27\x5b\x58\xca\xb0\x4d\xe9\x38\x95\xde\x68\xf2\x0d\x5c\xc3\x6a\x2c\x07\x11\x21\x48\x03\x45\x05\x2b\x97\x0d\x6b\xf0\x42\x3e\xac\x06\xa1\x2d\xf9\xb0\xb9\xd2\xf8\x22\x50\x07\xa9\x07\xb5\xd7\x73\x41\xdb\x9f\x1d\x11\x9e\x2b\x15\x56\xf4\x12\xb4\x55\x19\xca\x66\x4d\x2f\x7b\x41\x79\x89\xfd\x88\x41\x08\x3b\xc8\x38\x86\x9a\x59\x43\xd9\x0f\xa2\xa1\xa3\xbb\xc9\x60\xec\x04\xa2\x59\xc7\xa5\xd8\x87\xd9\xf3\xf5\x3e\x20\xb8\x0a\xfa\xd6\xf7\x0e\x8f\xb8\x01\x21\x02\x29\x42\x1b\x03\xd9\x89\xd0\x29\x16\x51\xec\x28\x43\x27\x86\xb7\x0d\xd9\x8c\x98\x16\x2b\x89\xeb\xe5\x60\x42\x99\xf0\xbe\x25\x4d\xf3\x40\xba\x90\x5a\x91\xa6\x29\x5a\xe2\xaa\xcf\x4b\xb2\x72\xab\xe4\x78\x77\xea\x88\x77\x91\x23\xde\xa5\x1d\x8f\xe8\x8f\x3f\x54\x25\xa7\x5a\x6c\xb4\x09\x9f\xe0\x77\x0f\xdc\x2e\x03\x71\x2b\x8e\x73\xd1\x32\x21\xff\x98\x96\x8d\xb4\xee\xdf\x7b\x60\xfb\xb0\x70\xbf\x6d\xa9\x2c\x93\x55\x4b\x66\xc9\x71\xf8\xe4\x72\x34\xe3\xa3\x12\x1d\x9a\xfd\x96\x60\x11\x5a\x86\xb2\x2c\xdf\x07\x10\x8c\x94\x21\xd3\x6a\xc4\x18\x3a\x3e\xaf\xb4\xb0\x13\xb7\xa4\x11\xc0\xa2\x32\xcb\x91\xba\xd2\x70\x30\x54\x9d\x30\xcb\xfa\x9e\xb4\xe1\x14\x25\xa7\xe1\xd6\x4f\x4f\x96\xcd\xb2\xba\xac\x6b\x06\x35\x16\x34\xca\x25\xf6\xa8\x9d\xa1\x42\x5d\x96\xf3\x15\x6d\x37\x98\xc1\x33\x0d\xc9\x52\x83\x85\xa0\x7e\x5d\xf2\xb2\x3c\xdf\xb5\xec\xb8\xc3\xe6\x2f\x15\x1a\xb5\x7f\x24\xcc\xb2\xfe\xc1\x2a\x60\x5f\xa3\x9e\x4e\x25\x54\xbc\xba\xdd\x90\x97\xe4\x3b\x5e\x13\x76\x6a\x59\x11\x16\x79\x6a\x51\x96\xe9\xcd\x4e\x30\x7c\x45\x9b\x90\x03\x89\x15\x25\x6d\xac\x69\x20\xca\x32\xbd\xeb\x4a\x06\x98\xc3\x9f\x94\x46\x3c\x31\x78\x51\x4b\xc2\x98\x0f\xc4\x99\xfd\xa1\xab\x88\x20\xb4\xc3\xcd\x2d\x65\x2d\x0e\xe7\xa3\x70\xb6\x1a\xd1\x7d\x19\x08\x4f\xa6\x2e\xd9\xb1\x5c\xac\x62\xa5\xa2\x05\xf3\xd8\xc4\xfe\xc8\xc6\x9c\xdc\x25\x4d\x96\x29\xb9\xf0\x65\x2d\x7d\x33\xb6\x7a\x54\x0c\x87\xa5\x1d\x3f\xe1\xc0\x8c\x94\xd9\xfe\x77\xdd\x16\x98\x88\x39\xa2\x31\xd7\xc4\x46\x91\xed\xf9\x82\x9b\x7e\x50\x11\x5b\x05\xd9\x4d\x98\xe1\xf3\x33\xd9\xe2\x1a\x7e\x3e\xdd\x0f\xb2\x29\xd1\xa2\x67\xae\xa7\x07\xb2\x19\xeb\xd9\xdd\x41\x1e\x34\xb2\x65\x4a\x6c\x9a\xc6\xb6\x4d\x7f\x79\xff\x19\x4e\x8f\x33\x8b\xed\x8c\x5b\x74\xa0\xd1\x7f\x89\x63\xcd\x74\xf5\x0f\x76\x37\x9f\xad\xf3\x67\x5c\xeb\x75\x32\x9a\x39\x04\xae\xcd\x82\xca\xd1\x05\xfa\xf7\xe1\x60\x66\x07\xaf\x3e\x1e\x17\x87\xc3\xf9\xf1\xb8\xf8\xdb\xe1\x00\x5d\x75\x3c\xfe\x27\xbd\x03\x33\x1b\xa1\x4f\xed\xc1\x46\x37\x65\x37\x5b\xe8\xa2\xe9\xcb\x77\x92\x51\xea\x64\x5e\x03\xa9\x0a\xfa\xcd\xe4\xbe\x4b\x66\x5b\x89\x07\x59\xfa\x72\x92\xc9\x71\x64\x90\xf1\x07\xcc\xde\xa2\x6a\xf6\x19\x1f\xa5\x4e\x32\xde\x2a\xd5\xc8\x56\xf2\xff\xb5\x07\xbc\xc7\xaf\x10\x2d\x80\x8d\x02\xec\x16\xc8\xb0\x59\x56\x27\xf3\xd6\xcc\xac\x35\xda\x67\xc6\x2a\x3f\xbb\xc7\x48\x03\x7b\xa6\x3d\x1b\xe1\x13\x17\x05\xd1\xa5\x87\xdc\x5b\x46\x77\x1e\x5e\x90\x3c\xf6\x87\x27\xde\xf0\xa0\x1b\x1d\x71\x27\xe2\x3e\x32\x28\x09\x27\xd4\x1d\xa7\x36\x0e\xf0\x57\x60\x4e\x31\x7b\xf1\x11\xdf\x09\x65\x1c\x6b\x27\xbc\x6e\x69\x17\xd5\xcd\x8a\x76\x71\xdd\x04\x82\xa4\xd1\x65\x03\x4c\xdc\x5d\x5b\x18\xcb\x60\x41\xdc\x4e\xc8\xd3\x69\x97\x5e\xd0\x27\x58\x31\xe0\x6e\x1e\xc5\xbd\xa0\x05\x33\x58\x70\x0f\xe0\x65\x49\xc7\x7b\xa8\x03\xb4\xd1\x21\xdb\xfd\x2d\x97\xe1\x10\x56\x92\x76\x89\xaa\x29\x12\x25\xed\xa2\x43\xa2\x3a\x8d\x8d\x4f\x56\x53\x7d\xd0\xaf\xf3\xc1\xda\x1e\x2e\xeb\x13\x11\xaf\xd6\x50\xbe\xb9\x2b\x2f\x15\x70\x07\x00\xcd\x24\xa3\xff\xc9\x68\xbf\x21\x5d\x6d\xf1\xda\x86\x8d\x49\xc0\xcf\xfb\x44\x26\x91\xc3\xfc\xcd\x5b\xf9\xf6\x48\xb9\xbf\xe1\x2a\xdf\x8a\x0d\xe5\xfe\xf6\xd2\xd2\xb3\x2e\x37\x55\x0d\x91\x0d\x48\x20\xf0\x31\x82\xf4\x25\xe8\xf3\x83\x6b\x8e\xb5\x68\x5d\x7b\x28\xfc\xbf\x08\x78\x89\x37\xf0\x1d\xde\x1b\xd2\x41\x56\xae\xc2\x1c\x85\x99\x99\x8b\xfb\xb5\x2e\xc3\x2d\xfb\x6b\x1d\x6d\xd7\x3d\x9b\x34\x51\x8b\x6d\xd8\xd1\xd5\x6c\x1d\xf5\xf3\x50\x92\xf4\x5a\x92\x8f\xa0\xb7\x70\x1d\x32\x26\x8e\x4b\x3a\x3c\x60\x56\xfb\xdb\xa7\x56\x87\xec\x15\x83\xe5\xd2\xe5\xe9\xb6\xae\x24\xdd\xd6\x15\x41\xa1\xe9\xcc\x03\xdb\x92\x12\xcc\x8f\x2b\x84\x0e\x16\xe6\xd7\x95\x66\x28\xce\xb1\x36\x6f\x21\x91\xaf\x7a\x2e\x89\x4d\xa3\x27\x93\xe9\x9a\x86\x17\x60\xe1\xda\x22\x5b\xac\xd8\x1a\xcc\x57\x7a\x20\x4b\xcf\x78\x78\x4f\x7b\x31\xf0\x6c\x14\x38\x74\x1d\x4a\xd3\x2d\xda\x73\xb1\x5c\xd3\xf7\x6f\xc4\x1f\x6a\xdb\x9e\x8b\x82\xaf\xe9\x7b\xb1\x96\xa8\x6d\xdf\x58\x99\xe5\x7a\xc3\x18\x65\x23\xbe\xa0\xf1\x81\xb3\x53\x67\x79\xdf\x63\x01\x5d\xb9\x3f\x35\x6f\x0c\x31\x70\xf7\xfa\x2c\xfb\xaf\x0c\xf0\x5b\x45\xdf\xbb\xd3\x04\x5e\x1d\x35\x48\x22\x8c\x93\x95\xc8\x35\xe1\x82\x91\xd7\x5e\x04\x8d\xea\xd3\xa9\x42\x76\x90\xd4\x20\x66\x56\x6a\x4f\xe6\x9d\xf4\x9e\x70\x71\x9a\x9a\x7d\x45\x2d\x1a\x49\x0f\x92\x1b\x44\x4d\x26\x77\x4d\xf8\xa6\xc1\x7b\x7d\x03\x62\xc9\x4a\x83\xf6\x82\xc3\x3e\x22\x0e\xa4\xf3\x0f\x44\x8f\x0c\x56\xc0\xa0\xf3\x53\x80\x9a\x46\x8b\x8d\xc7\xc3\xfb\x90\x48\x9e\x5e\x36\x48\x05\xff\x02\x46\xaf\x68\xdf\xf9\xa1\xb0\x26\x15\x14\x1f\xc0\x68\x51\x6a\xdc\x2e\x26\x43\xf5\xcc\xb4\xd9\xe1\x1a\x96\x02\x8b\x9e\xcb\xda\x8c\x5e\x21\x5b\x45\x16\x5c\xb1\xba\x29\xa2\x87\xc9\xa9\xc8\xd9\x29\x86\x2f\x9b\x23\x29\xea\xf7\xc6\xf0\xb5\x6e\x2a\xf2\xe7\x52\x0c\x17\xad\xb1\x14\xc3\x25\x6c\x2a\xf2\xe7\x52\x8c\x36\xe0\x63\x49\x46\xfb\xf1\xc9\xe8\xe9\x33\x03\x66\xb8\xe5\x21\xaa\xcf\x96\x31\xbe\xd1\xa1\xe1\x1b\xdf\x92\x86\x17\x57\x9c\xfa\x2b\x2b\xc3\x64\x3d\xa0\x79\x74\xf2\x89\x35\xe1\xa0\x7a\xac\x87\x55\xb7\xf6\x43\x46\x73\x19\xaf\x6e\x38\x2a\x89\x0a\xda\xa2\x18\x6e\xc2\x64\xf2\x58\x38\x3e\xe2\xfb\xb6\xe3\xc1\x48\x57\x41\xfd\x78\x66\x49\xbb\xa5\xbb\x5a\x63\x86\xcb\xf0\x8a\x65\x6a\x29\xa5\x35\x07\x97\xfb\x46\x87\xec\xda\x69\xb9\x91\x4e\x20\x8f\xba\xfa\x7f\x12\x12\xff\xdf\x70\x3c\xfb\x5f\x00\x00\x00\xff\xff\x73\x3a\x2a\x67\x8e\x22\x00\x00")

func tmplScreenboardTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplScreenboardTmpl,
		"tmpl/screenboard.tmpl",
	)
}

func tmplScreenboardTmpl() (*asset, error) {
	bytes, err := tmplScreenboardTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/screenboard.tmpl", size: 8846, mode: os.FileMode(420), modTime: time.Unix(1597096224, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmplTimeboardTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\xc9\x6e\xdb\x3c\x10\xbe\xfb\x29\x06\x42\x0e\xff\x0f\xc4\x7e\x80\x02\x3e\xa4\x09\x1c\x14\xe8\x92\x26\x41\x7a\x28\x0a\x81\x96\x46\x0a\x51\x6a\x09\x45\x25\x71\x08\xbe\x7b\xc1\xe1\x26\xdb\x8a\x7b\xa8\x4f\x9c\xf9\xbe\x59\x39\x43\x4b\xe2\xd0\x8d\xb2\x40\xc8\x4a\xa6\x58\xd9\xd5\xb9\xe2\x0d\x6e\x3b\x26\xcb\x0c\x32\xad\x61\xf5\xa9\x04\x63\x32\xd0\x0b\x00\xc5\x95\x40\x70\xbf\xb5\x43\xef\x49\x65\x4c\xb6\x00\x28\x71\x28\x24\xef\x15\xef\xda\x00\x5f\x4d\x54\x8e\x24\x91\x95\x79\xd7\x8a\x1d\xf9\xb0\x9c\x5b\x64\xe5\x37\xab\x58\x1a\xb3\x00\xd0\xfa\x85\xab\x47\x58\x5d\x4b\xd6\x3f\x0e\x51\x29\x59\x5b\x23\xac\x80\xc4\xda\x62\x94\x52\x48\x6a\x26\x9d\xe4\xea\x0a\x2b\xde\x72\x4a\xc2\xb9\x03\x78\xe6\x6f\xa9\x88\x07\xfe\x66\x81\x60\xb4\x04\x5e\xc1\xea\x62\x54\xdd\x50\x30\x81\x16\x62\x51\xf0\x26\x09\x35\x26\xb3\x26\xd8\x96\xde\x75\x70\x70\x23\xb1\xe0\x83\x0f\xda\x47\xc1\x3b\x48\xe8\xbb\x0e\xae\x65\x37\xf6\xd4\x81\xda\x9e\x60\x0d\x3f\xb5\x3e\xab\x9d\xf6\xc3\x3a\x10\x8c\x09\xdd\x39\xe3\x6d\x89\xaf\xe7\x70\x86\x02\x9b\x03\x06\xaf\x3c\x6c\xcc\xb9\xd6\x14\x2c\xd3\x9a\x98\x74\x22\xcd\xaf\xf9\x44\xee\x8a\xae\x47\x4a\x64\xb0\x27\x9f\xc8\xe0\xb4\x36\x8c\x23\x9c\x4a\x24\x31\xfe\x29\x11\xb5\x73\xf7\x41\xda\x81\xa4\xb5\x9f\x03\x80\x9e\x09\x54\x0a\xf7\xe6\x93\x2c\x56\x37\x1e\x09\x77\x1c\xb9\x79\x25\x78\x3f\xcb\xdd\x58\x20\xf0\xcd\x7c\x3a\x5f\x98\xfc\x8d\x92\x1a\x63\x55\x6e\xd8\xf6\x94\x7e\x6a\xbd\x5d\x43\x50\xcc\x57\xed\x7a\x4c\x8b\x64\x85\x94\xdf\x33\x13\x63\x1c\xb6\x07\x12\x12\xa8\x35\x85\xff\xcc\xb6\x28\x6c\x1c\x41\x07\x4f\x76\xda\xa3\xa1\x4a\x25\x1c\x1f\x62\x51\xae\x82\x5b\x7c\x1a\x71\x50\xb3\x25\x48\x87\xc5\x1a\x9e\x26\xbd\xfe\x3e\xa2\xdc\xa5\x45\x8a\x79\x52\x69\x4b\x63\xa8\xde\x83\x72\xb5\xb6\x29\x80\xf7\x1e\x4d\x2e\xea\x5a\x62\xcd\x54\x27\x5d\x12\x56\xd9\x22\x64\x19\xfc\x77\x85\xb7\x58\xdd\x29\xc9\xdb\x7a\xca\xfb\x9f\xd6\x34\x99\x85\x3d\x4d\x9a\x18\x8d\x1c\xda\xa8\xcb\xc3\xb0\x71\xbe\xb4\xf6\xbd\x70\x9a\xc8\x3b\x1c\x39\x6b\x49\xbb\xee\x07\xcc\x6e\xba\x3f\x86\x3d\x4f\xa3\xb7\xdf\xef\x64\xfc\x83\x97\xea\xd1\x9a\xbe\xd0\xc1\x1b\x3a\xed\x09\xb3\xd3\x6d\xdd\xb7\x99\xd4\xe9\x0b\x0f\xad\x88\x80\x1b\xe9\xcb\xae\x2d\xe9\xa5\x64\x62\xd3\xc9\x86\xa9\x01\xa6\xc3\xfd\x2e\x1c\x9e\xe7\xd4\xd2\x22\x51\xf3\x8a\xb8\x7b\x5d\x83\x69\xdb\x4e\x77\x6d\x7f\x42\xac\xe7\xa6\x67\x72\x7a\xcb\x97\x49\x93\xa6\x2f\x15\x15\xf6\x67\x7e\xa9\x8e\x03\xc4\x66\x8c\x83\xea\x9a\x8f\xf5\x65\x27\xc8\x73\x41\x72\xbe\xad\xf3\x82\x34\x21\xfa\x01\xed\xaf\x1e\x37\x87\x1e\xab\x59\x8f\x9b\xf7\x3d\x4e\xaf\xcd\x01\x73\xa7\x45\xa2\x1e\xef\xfa\xbc\x64\xfc\x03\x3c\xd5\xa5\xeb\xbf\xc7\xa6\x17\x4c\xe1\x03\x93\x9c\x6d\x05\xc6\xa7\x6f\xf2\xef\x6c\xbf\x13\x3c\x2d\x7f\xf6\x3c\x7f\xf3\x2d\x6b\x70\xf2\x5e\x7c\xb5\x62\xb8\xaf\x5e\x62\xc5\x5f\x61\xf2\xf7\x68\xc5\x80\x96\x58\xb1\x51\xa8\xf4\x59\xe1\x44\x07\x1f\xe7\x6c\x16\x7f\x02\x00\x00\xff\xff\xf4\xa7\x64\x9f\xd8\x08\x00\x00")

func tmplTimeboardTmplBytes() ([]byte, error) {
	return bindataRead(
		_tmplTimeboardTmpl,
		"tmpl/timeboard.tmpl",
	)
}

func tmplTimeboardTmpl() (*asset, error) {
	bytes, err := tmplTimeboardTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmpl/timeboard.tmpl", size: 2264, mode: os.FileMode(420), modTime: time.Unix(1597095598, 0)}
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
	"tmpl/monitor.tmpl": tmplMonitorTmpl,
	"tmpl/screenboard.tmpl": tmplScreenboardTmpl,
	"tmpl/timeboard.tmpl": tmplTimeboardTmpl,
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
	"tmpl": &bintree{nil, map[string]*bintree{
		"monitor.tmpl": &bintree{tmplMonitorTmpl, map[string]*bintree{}},
		"screenboard.tmpl": &bintree{tmplScreenboardTmpl, map[string]*bintree{}},
		"timeboard.tmpl": &bintree{tmplTimeboardTmpl, map[string]*bintree{}},
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

