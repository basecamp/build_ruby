// Code generated by go-bindata.
// sources:
// data/Dockerfile-centos.template
// data/Dockerfile-lucid.template
// data/Dockerfile.template
// data/patches/1.0.0/01_for_tests
// data/patches/1.0.0/02_for_tests
// data/patches/2.1.0/01_readline.patch
// data/patches/2.1.1/01_readline.patch
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

var _dataDockerfileCentosTemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x54\x4d\x6f\xe3\x36\x10\xbd\xfb\x57\x0c\xb2\x87\x3d\x6c\x18\x61\xd1\x62\x11\x07\xf0\x61\x91\x34\x45\xd0\x36\x59\x38\x4d\x7d\xe8\xf6\x40\x89\x23\x9b\x35\x3f\x04\x92\x8a\xe3\x6c\xf3\xdf\xfb\x44\x49\x8e\xbc\x1f\x06\x2c\x72\xde\x0c\x87\xc3\x99\x37\x73\xbd\xbc\xfb\x83\xbe\x7c\x39\xbb\xd2\x31\x05\xff\xf2\x32\x5b\x3e\xdc\xd2\xbe\xb5\xd4\x36\x4a\x26\x26\xb1\x3f\x20\xda\xc5\x24\x8d\x01\x44\xeb\xaa\xea\xfe\xa2\x7a\xf7\x8e\xac\xdc\x32\xc9\x36\xf9\xbc\xd9\x72\x70\x6c\x84\xe2\x47\x36\x14\x1a\x2b\xca\x56\x1b\x45\x9f\x67\x84\xdf\xda\xe8\xb2\x1a\x74\xd8\xd6\xb5\x1e\x84\xb5\x2a\xed\xb0\x75\x55\x1b\x22\xc7\x41\xea\xcf\x05\x96\xca\x68\xc7\x03\xe8\x1b\x76\x31\x9a\x57\x47\x7b\x69\x47\xe9\x19\xe2\xf1\xd9\xb6\xdc\xe7\xcf\x18\x13\xb6\x6b\xb6\x91\x70\x8f\xa1\x24\x43\x7e\xdf\xdf\x27\x45\x1b\x43\x51\x6a\x57\x40\x79\x72\x4a\x27\xc3\x6b\xbb\x6d\xdd\x64\x44\x08\xa8\x95\x0e\x8b\xd1\xb4\x07\x9d\x17\x41\xf9\xea\x55\xd0\x27\xff\x64\x9f\xf9\x82\x2e\xb5\x7e\xe7\x8c\x97\xea\x21\x98\x97\x97\xff\x70\x23\xf9\xa7\xe7\x4b\x2a\x92\x6d\x66\xab\xbb\xe5\x6f\x57\x37\xcb\x2c\x14\x39\x4a\x1c\x58\x62\xfd\x8b\x43\xd4\xde\x0d\x05\xb9\xbc\xfe\xfd\xe3\xaf\xf7\x8b\xb7\xc2\xca\x50\x6d\x16\x4f\xe7\x1f\xc4\x87\x9f\x49\xdc\xfd\x44\xa2\xc6\x95\xb5\x8c\x09\xaa\xb4\x21\xb1\x06\xb4\x46\x3a\x49\xac\x72\xad\x56\xfc\x94\x82\xc4\x0a\xb3\xd6\xb5\x91\x95\x68\x64\x90\x96\x13\x87\x1e\x85\xc8\x2e\x6d\x18\x49\xef\x01\xe3\xdd\x3a\x7f\x7a\xd1\xea\x18\x35\x90\x5a\xb3\x51\x42\x3b\x9d\xb4\x34\xfa\x19\xf1\x41\x3f\xb8\x7c\x94\x41\xcb\xd2\x80\x2e\xab\xc6\x6b\x07\xdf\x02\x48\x17\xcf\x6a\x87\x95\x05\xc8\x05\x1f\xdd\x09\xc5\x95\x41\x00\x09\x8f\x13\xb2\xee\x2c\x91\xe7\xc4\x16\x31\x40\xab\x6d\x63\x74\xa5\x93\xa8\x5b\x57\x65\x9b\x89\x7d\x3e\xdd\x04\xae\x60\xaf\xa6\x8a\x38\xbe\xa4\xda\x42\x51\xea\xd4\xc7\x5a\x79\xdb\x48\x78\x8d\x49\x2d\x74\xf4\xf3\xf3\xf9\xfc\xe2\xfd\x7c\x3e\x27\x64\xed\xd3\xcd\xe5\x5b\x3a\x2b\x2a\xef\x6a\xbd\x6e\x03\x67\xb6\x08\x01\xef\xb5\x7e\x5a\x14\xbe\x49\xb9\x1c\x5f\x57\x63\x30\x63\xd7\xbd\x56\xc4\x0d\x72\xa7\x06\x4c\xe9\x98\xc1\x81\x38\x02\x9c\x38\xb6\xee\x38\x20\x02\x1b\x84\xfc\xc8\xb3\x37\x74\xcf\x1d\x0d\x93\x27\xef\xcc\x9e\x64\x5d\x73\x95\x28\x7a\xcb\xf4\xfe\x6c\x4e\x91\x83\x46\x45\x70\x3b\x96\x53\x2a\xdb\x44\x28\x12\xe1\x4d\x60\x1e\x2e\xcd\xdd\x16\x13\x37\x17\xb3\x37\x70\xd6\xfd\x3a\xa6\x64\x58\xfc\x7b\x7e\xe8\xd6\xab\x5f\xee\xff\x04\xc3\x16\x99\x61\x60\x72\xb6\x4e\x1b\xe4\x25\x48\xc7\xd9\x67\x3e\x83\x8d\xeb\x77\xc3\xc9\x53\xda\xc9\x48\x65\xf0\x5b\x76\x67\x74\x8f\xba\xa4\x84\x1a\x92\x4e\x98\x0d\x30\x4a\xbe\x73\xb4\xf3\x88\xb4\x63\x54\xca\xb1\x59\xe9\x54\xa4\x9d\x0f\xdb\x48\x35\xe2\xbc\x98\x4d\x82\x42\x32\x6f\x5b\x7b\xf9\xe9\x61\x60\xf5\xf4\xb6\x6f\xe3\x7c\xed\x8d\x6c\x0c\x68\xe8\x68\x11\x09\x5d\x38\x0a\xa9\x9b\x32\xa3\xe0\xe8\xbb\x3d\x34\xaa\x65\xd7\x8e\x1f\xd1\x41\xaf\xd0\x23\xfd\xc0\x18\xf0\x0d\xd8\x99\xf9\x85\x78\xb3\xb1\x3a\x9a\x60\x9f\x47\xf0\x68\x94\x1d\xd0\xc9\x4c\x3b\x60\xdf\x1b\x6e\x80\xbf\x9a\x6f\x07\xfc\x78\xd0\x4d\xef\x9b\x4c\xbc\x03\xfc\xcd\xe8\x13\xfd\x90\x29\x26\xa9\x6b\xa8\xc0\xc3\xae\xb5\xe1\x5b\xcc\x80\xc3\x63\x41\xf7\xd9\xff\x01\x00\x00\xff\xff\xe2\x94\xb3\x2e\x0a\x06\x00\x00")

func dataDockerfileCentosTemplateBytes() ([]byte, error) {
	return bindataRead(
		_dataDockerfileCentosTemplate,
		"data/Dockerfile-centos.template",
	)
}

func dataDockerfileCentosTemplate() (*asset, error) {
	bytes, err := dataDockerfileCentosTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/Dockerfile-centos.template", size: 1546, mode: os.FileMode(420), modTime: time.Unix(1448756444, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataDockerfileLucidTemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x54\x51\x6f\xdb\x36\x10\x7e\xf7\xaf\x38\xb4\x43\x0b\x0c\xa5\xb4\x74\x5d\x11\x67\x73\x81\x20\x59\x86\x60\x5b\x12\x24\xcb\xf2\xd2\x87\x52\xe2\x49\xe6\x22\x91\x02\x49\x25\x71\x32\xff\xf7\x7d\xa4\x2c\x5b\x49\x3b\x3f\x58\xc7\x8f\x77\xc7\xef\xc8\xef\xee\xe4\xf2\xfc\x4f\x7a\x7a\xca\x8e\xb5\x0f\xce\xae\xd7\xb3\xcb\xeb\x33\xe2\x72\x69\xe9\x95\xe2\x82\x96\x21\x74\x07\x79\xee\xb9\xec\x9d\x0e\xab\xac\x2f\x7a\x13\xfa\xac\xb4\x6d\x3e\x98\xd4\xf4\xa5\x56\x62\x74\xa0\x56\x6a\xf3\x8a\x3e\x7d\xa2\x9c\x43\x99\xcb\x2e\xe4\xde\xf6\xae\x64\x9f\x35\x38\x21\x65\x07\x28\x6a\x0e\xd4\x77\x4a\x06\x7e\x06\x69\xe3\x83\x6c\x1a\x12\x2b\x72\x7d\xb1\xda\xcb\xe6\xd9\x9e\xa8\x7a\x20\x45\xaf\x1b\x25\xd8\x7b\x36\x41\xcb\x86\x3e\xcf\x08\xbf\x46\x17\xe5\x47\xa1\xf8\x2e\x5a\x55\xa5\x47\xb3\x56\x45\x3b\xda\x06\xcc\x3c\xfb\x9f\xd2\x7a\x1b\xe6\x58\xaa\x46\x1b\x1e\xbd\xbc\x6f\x46\x73\x25\xdb\xc1\x7e\xc4\x62\xaf\x7e\x1e\x67\x3b\x36\xd1\x77\x4b\x6f\x42\x34\x3a\xe2\xb0\x26\x95\x14\x8d\xf1\xfa\x3a\x67\x55\x5f\x06\x6d\x4d\x56\x56\x59\x0c\xa8\xb9\xf5\x99\x75\x75\x3e\x2e\xb6\x86\x78\x9f\x7d\xc8\xde\x67\xa1\x7e\xa4\x7f\x83\x74\x64\x1f\x1e\x8f\x28\x0f\x6d\x37\x64\x55\xc9\x7e\xe1\x4d\x6f\xde\xec\x68\x90\xe7\xd0\x77\x99\x2b\x52\x00\x9c\x06\x74\xbc\xda\xaa\x6b\x49\x88\x42\x1b\xa5\xdd\x22\xef\xbd\xcb\x61\x03\x31\x56\x38\x65\xcb\x8d\xa5\x77\x35\x44\x71\xd8\x7b\xd3\x58\xa9\xae\x5d\xb3\x5e\xbf\x60\xf5\xf4\xe4\xa4\xa9\x99\xb2\x0b\x19\xca\x25\xfb\xf5\xfa\xf0\xf8\x38\x06\xad\xd7\x94\x63\x97\x8d\x82\xaa\x6e\xce\x2f\x7f\x3f\x3e\xbd\xdc\x91\x17\xf0\xb8\xc4\xf7\x6f\x76\x1e\xf7\xb2\x11\x5e\x65\x1d\x69\x50\xa5\x2f\x91\x55\xde\x78\xca\xbf\xcf\xba\x98\xf8\xcb\xcf\xa4\x2c\x25\x93\x44\xf7\x03\xfd\x42\xdf\xe9\x08\x99\x41\x40\x47\x27\x7f\x1c\xfe\x76\xb5\x78\x2b\x5a\xe9\xca\xe5\xe2\x61\xff\xa3\xf8\xf8\x81\xc4\xf9\x8f\x24\x2a\xd4\x53\x49\x1f\xb0\x15\x10\x5b\x03\xaa\xa1\x0f\x12\x37\x49\x69\x37\xfc\x10\x9c\xc4\x17\x6e\xbd\xe9\x3d\x2b\xd1\x49\x27\x5b\x0e\xec\x06\x14\x4b\x88\x0e\xa5\xb1\x1f\x80\xc6\x9a\x3a\xfd\x0d\xcb\x56\x7b\xaf\x81\x54\x9a\x21\x51\x6d\x74\x14\xa8\x7e\x44\x61\xd8\xdf\xa4\xbc\x93\x4e\xcb\xa2\x61\x20\x9d\xd5\x06\xb9\x05\x90\xc8\xe7\xe6\x1e\x5f\x16\xe8\x3e\xe4\x88\x11\x8a\xcb\x06\x04\xa2\x5a\x84\xac\xa2\x27\x1e\x2e\x70\x0b\x0e\xd8\xd5\x6d\xd7\xe8\x52\x07\x74\x85\x49\x8a\x12\x13\xff\x14\xdd\x39\x2e\xe1\xaf\xa6\x1b\x7e\xac\xa4\xbc\xc5\x46\xa1\xc3\xc0\x15\x6d\x8c\x1b\x25\x1c\xa0\x16\xda\xdb\xf9\xfe\x7c\x7e\xb0\x37\x9f\xcf\x09\xb7\x76\x71\x7a\xf4\x96\xb2\xbc\xb4\xa6\xd2\x75\xef\x38\xb5\x80\x10\xc8\x5e\xe9\x87\x45\x6e\xd1\xd9\xf1\x1d\x5f\x3e\xe3\xc6\x8d\x4d\xac\x56\xf8\x25\xee\x4e\x6d\x30\xa5\x7d\x02\x37\x4a\x14\x51\x6d\xcf\xbc\xa3\xc4\x84\xe3\x06\x94\xef\x78\xf6\x9a\xae\x18\x0a\xa7\x60\xc9\x9a\x66\x45\xb2\xaa\xb8\x0c\xe4\x6d\xcb\x04\x49\x43\xe6\x4e\xe3\x45\x70\x3a\x3e\xef\x30\x20\x02\xe1\x91\x08\x35\x41\x3a\x38\xb4\x95\xb7\x4c\x3e\x70\x77\x30\x7b\x8d\x64\xf1\x17\x95\x92\x60\xf1\xcf\xfe\xb6\x21\x8e\x7f\xbd\xfa\x0b\xd2\x5c\x24\x69\xa2\x3b\x92\x77\x58\xe2\x5e\x20\x6c\x4e\x39\x53\x0c\x0c\x33\x58\x9b\xc8\x77\x74\x2f\x3d\x15\xce\xde\xb2\xc9\xe8\x0a\xef\x12\x02\xde\x90\x74\x1c\x6d\x70\x0a\x36\x26\xba\xb7\x60\x1a\x15\x15\x12\xb7\x56\x1a\xe5\xe9\xde\xba\x5b\x4f\x15\x78\x1e\xcc\x26\xa4\x70\x99\x67\x7d\x7b\x74\x71\xbd\x69\x87\xe9\x69\x5f\xf3\xdc\x35\xd5\xd0\x3b\x68\xec\x61\x4c\x09\x4f\x68\xee\x71\x11\x28\x0e\xf2\xcd\xc2\xd0\x37\x9b\x6f\xdc\x96\xb1\x71\x0f\xd1\x41\x3b\xe8\x8e\xfe\xc7\x19\xf0\x29\xd4\x99\xf4\x05\xbe\xc9\x59\x4d\x26\xf2\xe7\x09\x34\x8e\xe6\x29\xb6\x9d\xd1\x53\xf0\x1b\xc3\x7a\xd8\x78\x36\xaf\xa7\x1b\xe3\xe0\x9e\x62\xdb\x09\xbe\x05\xbf\x1a\xe5\x62\x18\x5f\xf9\xe4\xd6\x3a\xca\x51\xd3\x89\x6e\xf8\x0c\xed\xbf\xad\x13\x4a\x9f\xfd\x17\x00\x00\xff\xff\x67\x40\x99\xf7\x26\x07\x00\x00")

func dataDockerfileLucidTemplateBytes() ([]byte, error) {
	return bindataRead(
		_dataDockerfileLucidTemplate,
		"data/Dockerfile-lucid.template",
	)
}

func dataDockerfileLucidTemplate() (*asset, error) {
	bytes, err := dataDockerfileLucidTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/Dockerfile-lucid.template", size: 1830, mode: os.FileMode(420), modTime: time.Unix(1448756529, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataDockerfileTemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x54\x5b\x4f\x23\x37\x14\x7e\xcf\xaf\x38\x62\x2b\xad\x54\xe1\x4c\xd1\xb6\x88\xd0\xe6\x01\x91\x52\xa1\xb6\x80\x42\x29\x0f\xdd\x4a\xeb\x8c\xcf\x04\x17\x8f\x3d\xb2\x3d\x84\x4b\xf3\xdf\xfb\xd9\x33\x13\x86\xee\x6e\x1e\x32\xf6\x77\x2e\x3e\x97\xef\x9c\xb3\xe5\xe5\xef\xf4\xf2\x32\x5d\xe8\x10\xbd\xdb\x6e\x27\xcb\x9b\x0b\x92\x4d\x14\x6b\x8e\xd4\x36\x4a\x46\x7e\x03\x69\x1b\xa2\x34\x86\xc4\x13\xf9\x76\xf5\x74\x30\x9d\x4d\x3f\xd0\xaa\xd5\x46\x09\x0e\x81\x6d\xd4\xd2\xd0\xc7\x09\xe1\x67\xf4\xaa\x3c\x14\x8a\x1f\xd2\xa9\xaa\xf4\x70\x5c\xab\x55\x3d\x9c\x6d\xd9\xfa\xc0\xe1\x87\x7c\xdf\x99\x79\x96\xca\x68\xcb\x83\x56\x08\x66\x38\x3e\xc9\xba\x3b\x3f\xe3\x72\xb0\x1e\xd9\xc1\x93\xc9\xa1\xfe\xb5\x57\xb4\xc1\x17\x2b\x6d\x8b\x35\xd7\x7b\xfb\xb4\xd7\x07\x9d\x8e\x55\x93\x11\x21\x20\x56\xda\xcf\x07\xd5\x0e\xb4\x4e\x78\xe5\xca\xd7\x8b\xde\xfb\x3b\xfb\x4c\xce\x73\x99\xdc\xc6\x1a\x27\xd5\x8d\x37\xdb\xed\xbf\x51\x7a\x72\x8f\xcf\xa7\x54\xc4\xba\x99\xbc\xbc\x78\x69\xd7\x4c\xd3\x2b\x19\xcb\x3b\x0e\xdb\xed\xc9\x62\x91\x8c\xb6\x5b\x2a\x20\x65\xab\x50\xdf\xdb\xcb\xe5\xaf\x8b\xf3\x65\x36\x29\x52\x09\x05\x34\x96\xf8\xfe\xc9\x3e\x68\x67\xfb\x16\x54\xce\x93\x46\xb5\xe9\x53\x4e\xc4\x04\x2a\xbe\x9d\x36\xc9\xf1\xa7\x1f\x49\x39\xca\x47\x12\xcd\x77\xf4\x13\x7d\xa3\x13\x64\xbb\x46\x9d\x9e\xfd\x76\xf2\xcb\xf5\xfc\xbd\xa8\xa5\x2f\xef\xe6\x8f\x47\x87\xe2\xf0\x7b\x12\x97\x1f\x48\x54\xc8\xa8\x92\x21\x42\x14\x61\xbb\x06\xb4\x46\x33\x48\xdc\xe6\x8e\xde\xf2\x63\xf4\x12\x5f\xa8\xb5\xb6\x0d\xac\x44\x23\xbd\xac\x39\xb2\xef\x50\x5c\xd1\x61\xa4\xc6\xa1\x03\x8c\xb3\xeb\xfc\xd7\x5d\x6b\x1d\x82\x06\x52\x69\x06\x1f\xb4\xd5\x89\x0d\xfa\x19\x89\x41\xde\xbb\x7c\x90\x5e\xcb\x95\x61\x20\x8d\xd3\x16\xbe\x05\x90\x14\xcf\xed\x06\x5f\x16\xe0\x21\x7c\x24\x0b\xc5\xa5\x41\x00\x11\x55\x11\xb2\x4a\x9a\x68\x63\xe4\x1a\x31\x40\xaa\xeb\xc6\xe8\x52\x47\x51\xb5\xb6\xcc\x3a\x23\xfd\x6c\xdd\x78\x2e\xa1\xaf\xc6\x82\x30\x64\x52\xde\x43\xb0\xd2\xb1\x8b\xb5\x74\x35\x2a\x4a\x78\x40\xcd\x75\x70\xb3\xa3\xd9\xec\xf8\x60\x36\x9b\x11\xaa\x76\x75\x7e\xfa\x9e\xa6\x45\xe9\x6c\xa5\xd7\xad\xe7\xcc\x37\x21\xe0\xbd\xd2\x8f\xf3\xc2\x35\x31\xf7\xf1\xff\x6d\xec\xd5\xd8\xa6\x6c\x45\xb8\x43\xed\x54\x8f\x29\x1d\x32\xd8\xf3\x52\x80\x72\x6f\xb5\x13\xc5\x84\x67\x83\x90\x1f\x78\xf2\x8e\xae\x99\xeb\x40\xd1\x91\xb3\xe6\x89\x64\x55\x71\x19\x29\xb8\x9a\x09\x03\x48\x81\xbd\x46\x47\xf0\x3a\x3e\xfb\x98\xc6\x48\x68\x12\x21\x27\x50\x07\x8f\xd6\xf2\x9e\x29\x44\x6e\x8e\x27\xef\xe0\x2c\xfd\x12\x53\x32\x2c\xfe\x39\xda\xcd\xf4\xe2\xe7\xeb\x3f\x40\xcd\x79\xa6\x26\x06\x25\x6b\xc7\x3b\xd4\x05\xc4\xe6\xec\x33\xdb\xe0\x60\xbb\x53\x6f\xb9\x4f\x1b\x19\x68\xe5\xdd\x3d\xdb\x29\x5d\xa3\x2f\x31\xa2\x87\xa4\xd3\x0a\x81\x52\x74\xc9\xd1\xc6\x21\xd2\xc4\xa8\x98\x63\xab\xa5\x55\x81\x36\xce\xdf\x07\xaa\x10\xe7\xf1\x64\x14\x14\x8a\x79\xd1\xd6\xa7\x57\x37\xfd\x38\x8c\x5f\xfb\x3c\xce\xd7\xa1\xea\x66\xa7\xa9\xfb\x9d\x20\x02\x61\xc8\x87\x4b\x24\xc5\xab\xe1\x62\xe9\x8b\xc3\x37\x88\x65\x1a\xdc\x13\x4c\xd0\x2b\xf4\x40\x5f\x51\x06\x7c\x0e\x76\x66\x7e\x21\xde\xac\xac\x46\xeb\xef\xe3\x08\x1a\xf6\xe0\x18\xdb\x2d\xc4\x31\xf8\x85\xcd\xd8\x09\xde\x2c\xc7\xb1\x60\xd8\x92\x63\x6c\xb7\x2e\x77\xe0\x67\x7b\x53\x74\xeb\xab\x18\x55\xad\xa1\x02\x39\x9d\x69\xc3\x17\x18\xff\x5d\x9e\x60\xfa\xe4\xbf\x00\x00\x00\xff\xff\x4f\xd6\xc5\x3b\x30\x06\x00\x00")

func dataDockerfileTemplateBytes() ([]byte, error) {
	return bindataRead(
		_dataDockerfileTemplate,
		"data/Dockerfile.template",
	)
}

func dataDockerfileTemplate() (*asset, error) {
	bytes, err := dataDockerfileTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/Dockerfile.template", size: 1584, mode: os.FileMode(420), modTime: time.Unix(1448756512, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataPatches10001_for_tests = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\x4a\x2c\x02\x04\x00\x00\xff\xff\xaa\x8c\xff\x76\x03\x00\x00\x00")

func dataPatches10001_for_testsBytes() ([]byte, error) {
	return bindataRead(
		_dataPatches10001_for_tests,
		"data/patches/1.0.0/01_for_tests",
	)
}

func dataPatches10001_for_tests() (*asset, error) {
	bytes, err := dataPatches10001_for_testsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/patches/1.0.0/01_for_tests", size: 3, mode: os.FileMode(420), modTime: time.Unix(1448756444, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataPatches10002_for_tests = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\x4a\xac\x02\x04\x00\x00\xff\xff\x98\x04\x24\x78\x03\x00\x00\x00")

func dataPatches10002_for_testsBytes() ([]byte, error) {
	return bindataRead(
		_dataPatches10002_for_tests,
		"data/patches/1.0.0/02_for_tests",
	)
}

func dataPatches10002_for_tests() (*asset, error) {
	bytes, err := dataPatches10002_for_testsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/patches/1.0.0/02_for_tests", size: 3, mode: os.FileMode(420), modTime: time.Unix(1448756444, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataPatches21001_readlinePatch = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x54\xeb\x4f\xdb\x3e\x14\xfd\xfc\xf3\x5f\x71\x15\x7e\x52\x93\xe5\xb1\x3e\xc2\xa3\x03\xa4\x32\x1e\x2b\x1a\x02\x44\xd9\xbe\x5a\x26\xb9\xa5\x16\xa9\x13\xd9\x4e\x07\xff\xfd\xec\xd0\x84\x05\xb5\x15\x6c\x9a\xb4\x7c\x48\x5d\xe7\xdc\x73\x4f\xce\x3d\xce\xb9\x48\xf1\xf1\x13\xe0\xa3\xfe\x28\x91\xa5\x19\x17\xd8\x2c\xa2\x84\x1c\xfe\xf9\x45\xc2\x30\x5c\xc7\xff\x9f\x2b\x71\xc1\x15\xcf\x05\xc4\xdb\xfd\x7e\xec\x11\xdf\xf7\xdf\x08\xde\xf6\xc8\x68\x04\x61\x6f\xb8\x1b\x07\xbb\xe0\x2f\x7f\x47\x23\x02\x04\xec\x25\x33\xca\xb4\xc6\x79\xa1\x31\xa5\x49\x3e\x2f\x32\xd4\xa6\x96\x4e\x4b\x91\xd8\x05\x1c\x42\xcd\xbd\x19\xb8\x4f\x60\x8b\x4f\x21\xc5\xa9\x81\xa6\xee\xf8\xe8\xfb\x29\xbd\xb9\xa0\xd7\x37\xa7\xf4\xfc\xf2\xfa\xdb\x2d\x1d\x5f\x5d\x7d\xf5\x48\xb8\x6c\x5a\x48\xa4\x5c\x14\xa5\xa6\xb3\x3c\x7f\x30\x5d\xdc\xb3\xba\xe3\x07\xaf\xe9\xd8\x46\xed\x13\x7f\x6d\xb5\xd9\xb4\xcb\x4a\x0e\xd5\x9b\x38\x60\x0b\x45\xca\xa7\x95\x5c\xa3\x16\x6a\xa5\xc7\x47\xb7\xc7\x63\x3a\x39\xff\x72\x79\x74\x31\x69\xcc\x49\x98\x4e\x66\x54\xf1\x7b\xc1\x32\x65\x1a\x75\xf7\xc9\xaa\x30\x98\x3f\x49\x2e\xa6\x91\xbc\xfb\x3b\x61\x78\xe1\x6f\xcf\x77\x30\x5c\x11\x86\x75\xe0\xb8\x5b\x87\x21\xd8\xb1\x51\x08\x7a\xdd\x2a\x09\xe6\x45\x51\x97\x52\x80\x2a\x0b\x94\xae\xb5\x30\x80\x99\x61\x43\xa9\x3c\x02\xc6\x2e\x93\x16\xdf\x7a\xd5\xc4\x6c\xc6\x16\x48\xf5\x53\x81\xae\xbd\x79\x76\x30\x2d\x0e\xbb\xfb\x0b\x87\x6f\x39\x7c\x02\x29\x97\xd4\x4a\xe3\xf7\x6e\x27\x29\xa5\x42\xd5\xf1\xda\xbb\x62\xf5\xb6\x46\x39\x4f\x58\xd1\x79\x7e\x81\x61\x1c\xc4\xe0\x0f\xf7\x82\xbd\x4a\x7f\x5b\x95\x95\xef\x3a\x66\x74\x12\x53\xae\x8a\x8c\x3d\x39\xde\x5a\x0c\x17\x0a\xa5\xa6\xda\x58\xb6\x01\x95\xa2\x09\x3b\xd6\x28\xbf\x14\x19\x2a\xb5\xca\x0c\xa7\x1d\x43\xa7\x32\xe6\xff\x93\xd3\xb3\x09\x1c\x1c\x80\x13\x9e\xb4\x9f\x1f\xd6\x99\x77\x1a\x87\x12\xc3\x6a\x5a\xcd\xd9\x83\x39\x48\x99\x65\x5c\x76\x31\x5c\xff\x52\xf0\xe2\xfe\x7b\x82\x37\x58\xce\xad\x0a\x9e\xb9\x0f\x37\xcd\xed\x77\x66\x02\xef\x9c\xc9\x16\xbc\xfa\x60\x70\x05\x6c\xc1\x78\xc6\xee\x32\x04\xc5\x45\x82\x0d\x57\x18\x47\x7d\x70\xfb\xdd\x6e\xcf\x8b\x9e\x6b\x9b\x4f\x15\xb7\x1d\xe7\xf9\x02\x53\x60\xfa\xa5\x60\x27\x1a\xd8\x82\x5e\x5c\x17\x8c\xf3\x1f\xb8\x40\x19\x80\x89\xa4\xb6\x10\x70\x2f\x51\x7f\x9e\x9c\xc0\x4e\xd4\x8b\x06\x01\x54\x60\x48\x73\x54\xa2\xa3\xc1\x8a\x7f\x25\x30\xb2\xc7\xf4\x4d\x49\xb2\xe7\xf5\x67\x00\x00\x00\xff\xff\xdd\x85\xa5\xd3\xb0\x06\x00\x00")

func dataPatches21001_readlinePatchBytes() ([]byte, error) {
	return bindataRead(
		_dataPatches21001_readlinePatch,
		"data/patches/2.1.0/01_readline.patch",
	)
}

func dataPatches21001_readlinePatch() (*asset, error) {
	bytes, err := dataPatches21001_readlinePatchBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/patches/2.1.0/01_readline.patch", size: 1712, mode: os.FileMode(420), modTime: time.Unix(1448756444, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataPatches21101_readlinePatch = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x54\xeb\x4f\xdb\x3e\x14\xfd\xfc\xf3\x5f\x71\x15\x7e\x52\x93\xe5\xb1\x3e\xc2\xa3\x03\xa4\x32\x1e\x2b\x1a\x02\x44\xd9\xbe\x5a\x26\xb9\xa5\x16\xa9\x13\xd9\x4e\x07\xff\xfd\xec\xd0\x84\x05\xb5\x15\x6c\x9a\xb4\x7c\x48\x5d\xe7\xdc\x73\x4f\xce\x3d\xce\xb9\x48\xf1\xf1\x13\xe0\xa3\xfe\x28\x91\xa5\x19\x17\xd8\x2c\xa2\x84\x1c\xfe\xf9\x45\xc2\x30\x5c\xc7\xff\x9f\x2b\x71\xc1\x15\xcf\x05\xc4\xdb\xfd\x7e\xec\x11\xdf\xf7\xdf\x08\xde\xf6\xc8\x68\x04\x61\x6f\xb8\x1b\x07\xbb\xe0\x2f\x7f\x47\x23\x02\x04\xec\x25\x33\xca\xb4\xc6\x79\xa1\x31\xa5\x49\x3e\x2f\x32\xd4\xa6\x96\x4e\x4b\x91\xd8\x05\x1c\x42\xcd\xbd\x19\xb8\x4f\x60\x8b\x4f\x21\xc5\xa9\x81\xa6\xee\xf8\xe8\xfb\x29\xbd\xb9\xa0\xd7\x37\xa7\xf4\xfc\xf2\xfa\xdb\x2d\x1d\x5f\x5d\x7d\xf5\x48\xb8\x6c\x5a\x48\xa4\x5c\x14\xa5\xa6\xb3\x3c\x7f\x30\x5d\xdc\xb3\xba\xe3\x07\xaf\xe9\xd8\x46\xed\x13\x7f\x6d\xb5\xd9\xb4\xcb\x4a\x0e\xd5\x9b\x38\x60\x0b\x45\xca\xa7\x95\x5c\xa3\x16\x6a\xa5\xc7\x47\xb7\xc7\x63\x3a\x39\xff\x72\x79\x74\x31\x69\xcc\x49\x98\x4e\x66\x54\xf1\x7b\xc1\x32\x65\x1a\x75\xf7\xc9\xaa\x30\x98\x3f\x49\x2e\xa6\x91\xbc\xfb\x3b\x61\x78\xe1\x6f\xcf\x77\x30\x5c\x11\x86\x75\xe0\xb8\x5b\x87\x21\xd8\xb1\x51\x08\x7a\xdd\x2a\x09\xe6\x45\x51\x97\x52\x80\x2a\x0b\x94\xae\xb5\x30\x80\x99\x61\x43\xa9\x3c\x02\xc6\x2e\x93\x16\xdf\x7a\xd5\xc4\x6c\xc6\x16\x48\xf5\x53\x81\xae\xbd\x79\x76\x30\x2d\x0e\xbb\xfb\x0b\x87\x6f\x39\x7c\x02\x29\x97\xd4\x4a\xe3\xf7\x6e\x27\x29\xa5\x42\xd5\xf1\xda\xbb\x62\xf5\xb6\x46\x39\x4f\x58\xd1\x79\x7e\x81\x61\x1c\xc4\xe0\x0f\xf7\x82\xbd\x4a\x7f\x5b\x95\x95\xef\x3a\x66\x74\x12\x53\xae\x8a\x8c\x3d\x39\xde\x5a\x0c\x17\x0a\xa5\xa6\xda\x58\xb6\x01\x95\xa2\x09\x3b\xd6\x28\xbf\x14\x19\x2a\xb5\xca\x0c\xa7\x1d\x43\xa7\x32\xe6\xff\x93\xd3\xb3\x09\x1c\x1c\x80\x13\x9e\xb4\x9f\x1f\xd6\x99\x77\x1a\x87\x12\xc3\x6a\x5a\xcd\xd9\x83\x39\x48\x99\x65\x5c\x76\x31\x5c\xff\x52\xf0\xe2\xfe\x7b\x82\x37\x58\xce\xad\x0a\x9e\xb9\x0f\x37\xcd\xed\x77\x66\x02\xef\x9c\xc9\x16\xbc\xfa\x60\x70\x05\x6c\xc1\x78\xc6\xee\x32\x04\xc5\x45\x82\x0d\x57\x18\x47\x7d\x70\xfb\xdd\x6e\xcf\x8b\x9e\x6b\x9b\x4f\x15\xb7\x1d\xe7\xf9\x02\x53\x60\xfa\xa5\x60\x27\x1a\xd8\x82\x5e\x5c\x17\x8c\xf3\x1f\xb8\x40\x19\x80\x89\xa4\xb6\x10\x70\x2f\x51\x7f\x9e\x9c\xc0\x4e\xd4\x8b\x06\x01\x54\x60\x48\x73\x54\xa2\xa3\xc1\x8a\x7f\x25\x30\xb2\xc7\xf4\x4d\x49\xb2\xe7\xf5\x67\x00\x00\x00\xff\xff\xdd\x85\xa5\xd3\xb0\x06\x00\x00")

func dataPatches21101_readlinePatchBytes() ([]byte, error) {
	return bindataRead(
		_dataPatches21101_readlinePatch,
		"data/patches/2.1.1/01_readline.patch",
	)
}

func dataPatches21101_readlinePatch() (*asset, error) {
	bytes, err := dataPatches21101_readlinePatchBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/patches/2.1.1/01_readline.patch", size: 1712, mode: os.FileMode(420), modTime: time.Unix(1448756444, 0)}
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
	"data/Dockerfile-centos.template": dataDockerfileCentosTemplate,
	"data/Dockerfile-lucid.template": dataDockerfileLucidTemplate,
	"data/Dockerfile.template": dataDockerfileTemplate,
	"data/patches/1.0.0/01_for_tests": dataPatches10001_for_tests,
	"data/patches/1.0.0/02_for_tests": dataPatches10002_for_tests,
	"data/patches/2.1.0/01_readline.patch": dataPatches21001_readlinePatch,
	"data/patches/2.1.1/01_readline.patch": dataPatches21101_readlinePatch,
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
		"Dockerfile-centos.template": &bintree{dataDockerfileCentosTemplate, map[string]*bintree{}},
		"Dockerfile-lucid.template": &bintree{dataDockerfileLucidTemplate, map[string]*bintree{}},
		"Dockerfile.template": &bintree{dataDockerfileTemplate, map[string]*bintree{}},
		"patches": &bintree{nil, map[string]*bintree{
			"1.0.0": &bintree{nil, map[string]*bintree{
				"01_for_tests": &bintree{dataPatches10001_for_tests, map[string]*bintree{}},
				"02_for_tests": &bintree{dataPatches10002_for_tests, map[string]*bintree{}},
			}},
			"2.1.0": &bintree{nil, map[string]*bintree{
				"01_readline.patch": &bintree{dataPatches21001_readlinePatch, map[string]*bintree{}},
			}},
			"2.1.1": &bintree{nil, map[string]*bintree{
				"01_readline.patch": &bintree{dataPatches21101_readlinePatch, map[string]*bintree{}},
			}},
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

