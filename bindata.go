// Code generated by go-bindata.
// sources:
// data/Dockerfile-centos.template
// data/Dockerfile-lucid.template
// data/Dockerfile.template
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

var _dataDockerfileCentosTemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x52\x4d\x73\xd3\x30\x10\xbd\xe7\x57\xec\xf8\x5a\x54\x1f\x60\x3a\x5c\x72\xe8\x24\x84\xe9\x00\x09\x93\x12\x38\x50\x0e\xb2\xb5\x76\x45\xf5\xe1\x91\xe4\x90\xb4\xe4\xbf\xb3\x56\x64\x27\x2e\x6d\x66\x92\xec\x3e\xad\x76\xf5\xde\xbe\xc5\x7a\xf5\x05\x9e\x9e\x2e\xe7\xd2\x07\x67\x0f\x87\xc9\x7a\xb3\x84\x7d\xab\xa1\x6d\x04\x0f\x08\x6c\x3f\x20\xd2\xf8\xc0\x95\x22\x08\xea\xb2\xec\xbe\xac\xbc\xb8\x00\xcd\x1f\x10\x78\x1b\x6c\x0c\x1e\xd0\x19\x54\x4c\xe0\x16\x15\xb8\x46\xb3\xa2\x95\x4a\xc0\xdd\x04\xe8\x53\x2b\x59\x94\xe9\x8c\xc2\xaa\x92\x29\xa9\x45\xa1\x53\x68\xca\xd6\x79\xf4\x29\x3b\xde\x73\xc8\x85\x92\x06\x13\x68\x1b\x34\xde\xab\x53\xa3\x3d\xd7\x7d\xf6\x48\xe9\xf8\x6e\x5b\xec\xe3\x4f\xff\x26\x0a\x6b\xd4\x1e\x68\x8e\x82\xc0\x5d\xe4\xf7\x33\xcb\x5b\xef\xf2\x42\x9a\x9c\x0e\xb3\x37\x90\x25\xb6\x5d\x58\x35\x11\x61\x8c\x8e\x85\x74\xd3\xbe\xf4\x08\x1a\xcb\x9c\xb0\xe5\x29\x91\xd9\xaf\xd8\x33\x0e\xe8\xa4\xb5\x7f\x8c\xb2\x5c\x6c\x9c\x3a\x1c\xfe\xd2\x44\xb0\xbb\xc7\x19\xe4\x41\x37\x93\x1f\xab\xf5\xa7\xf9\xcd\x3a\x26\x79\x7c\x25\x5d\x58\xd3\xff\x77\x74\x5e\x5a\x93\x16\x32\x5b\x7c\xbe\xfe\x78\x3b\xcd\x98\xe6\xae\xbc\x9f\xee\xde\x5f\xb1\xab\x77\xc0\x56\x6f\x33\xb8\xcc\x4b\x6b\x2a\x59\xb7\x0e\x23\x63\xc6\x1a\x87\x95\xdc\x4d\x73\xdb\x84\xd8\xf2\x79\xc7\x54\x86\x86\x17\x0a\x99\xbf\xe7\x0e\x45\xc2\x84\xf4\x11\x4c\xe4\x19\xf1\x1a\x57\x77\x3c\x98\x43\xc5\x83\xdc\x62\x7c\x5a\xdc\x3a\xfb\x4d\x43\x96\xad\x9e\x7d\xdd\x50\xff\xde\x28\xf3\x0f\xb7\xdf\x88\xdc\x34\x92\x23\x11\x27\x27\xba\xf1\x2a\x41\x69\x49\xcc\x03\x09\xdb\x27\xa1\x33\x4e\x9f\x18\x78\x51\x96\xfe\x98\x77\x0a\x5f\x93\x28\x27\x68\x0b\xaf\x14\x13\x7c\x13\xd0\xd1\xdb\xa3\xb0\xb1\x58\x8c\x4c\x79\xd7\x83\x23\x77\x0e\xe8\x99\x4d\x07\xec\x25\xbf\x12\xfc\xcc\xb2\x03\x3e\xf6\xee\xf9\xbc\x33\x13\x0f\xf0\x7f\x6e\x66\x47\xdf\xe4\x67\xd2\x35\x90\x13\xb1\x85\x54\xb8\xe4\x1a\x07\xb2\xb4\xfd\xc9\xbf\x00\x00\x00\xff\xff\x1b\x1d\xd6\x78\xdd\x03\x00\x00")

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

	info := bindataFileInfo{name: "data/Dockerfile-centos.template", size: 989, mode: os.FileMode(420), modTime: time.Unix(1447200546, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataDockerfileLucidTemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x53\xcd\x6e\xdb\x3c\x10\xbc\xfb\x29\x16\x3e\xe4\x46\x12\xc9\x97\x2f\x68\x0b\x38\x40\x90\x34\x45\xd0\x36\x29\x9c\xa6\xbd\xe4\x42\x89\x2b\x99\x2d\x45\x0a\xfc\x49\x63\xa7\x7e\xf7\x92\x92\x28\xcb\x6e\xea\x8b\x87\xa3\x59\x72\x67\xb0\x7b\xbd\xbc\xfb\x0c\x2f\x2f\xf4\x4a\x3a\x6f\xcd\x76\x3b\x5b\x3e\xdc\x02\x96\x2b\x03\x73\x81\x05\xac\xbc\x6f\xdf\x31\xe6\xb0\x0c\x56\xfa\x35\x0d\x45\xd0\x3e\xd0\xd2\x34\xac\x87\xa0\x42\x29\x05\xc9\x02\x68\xb8\xd4\x73\x38\x3f\x07\x86\xbe\x64\xbc\xf5\xcc\x99\x60\x4b\x74\x54\xc5\x17\xba\xdb\x23\x49\x6a\xf4\x10\x5a\xc1\x3d\xee\x51\x52\x3b\xcf\x95\x02\xb2\x06\x1b\x8a\xf5\x31\x7d\x4b\x8f\x49\x15\x22\x53\x04\xa9\x04\x41\xe7\x50\x7b\xc9\x15\x3c\xce\x20\xfe\x94\x2c\xca\x33\x22\xf0\x29\xa1\xaa\x92\x19\xd6\xa2\x68\x32\xd6\xb1\x33\x87\xee\xff\xee\x3c\x96\x59\xe4\x42\x49\x8d\x59\xe5\x9c\xca\x70\xcd\x9b\x1e\x6f\xe2\xe1\xb8\xde\xaf\x33\x2d\xea\xa4\x1d\xdb\x9b\x34\x9a\x84\xf1\x31\xd5\x59\x4a\x20\xc7\xd7\x5a\x23\x42\xe9\xa5\xd1\xb4\xac\x68\x2a\xa8\xb1\x71\xd4\xd8\x9a\xe5\xc3\x08\xc8\x09\x3d\xa5\x27\xd4\xd7\x1b\xf8\xed\xb9\x05\xf3\xbc\xb9\x04\xe6\x9b\xb6\xbf\x55\x74\xf8\x40\x0d\x47\x47\xbb\x36\xc0\xa1\x0f\x2d\xb5\x45\x57\x10\x45\x3d\x9b\xa3\xad\xda\x06\x08\x29\xa4\x16\xd2\x2e\x58\x70\x96\x45\x1c\x19\x6d\x88\x15\xa6\x1c\x90\xdc\x79\x48\xc3\x61\x7e\x69\x65\xb8\x78\xb0\x6a\xbb\x3d\xe8\xea\xfb\xdd\xf2\xe3\xd5\xcd\x72\xd7\x16\x89\x05\xcb\xf8\xff\x0d\xad\x8b\x8e\x87\x91\xba\xbc\xfe\x74\xf1\xe1\x7e\x31\x27\x0d\xb7\xe5\x6a\xf1\xfc\xe6\x8c\x9c\x9d\x02\xb9\xfb\x6f\x0e\x94\x95\x46\x57\xb2\x0e\x16\xbb\x9c\x09\x69\x2d\x56\xf2\x79\xc1\x4c\x1c\x9f\x74\xe5\xe1\x8d\x83\x0c\x35\x2f\x14\x12\xb7\xe2\x16\xc5\xc0\x09\xe9\x3a\x72\xb0\x4b\x92\xa5\x3d\x75\xf2\x41\x2c\x2a\xee\xe5\x53\x3f\x7c\x0d\xff\x89\x40\x7e\xc4\x47\x6e\x43\x73\xf9\xe5\x61\xe8\xb8\xa3\x73\x6a\x57\xef\xef\xbf\x46\x97\x8b\xce\x65\x8c\x70\xb6\xf3\xdd\x89\x53\xaa\xfd\x8c\x10\x07\x31\xd9\x7c\xf0\x90\xb6\x68\x38\x68\x78\x35\x9f\xfc\x99\xa7\xa8\x2f\x62\x3a\x3b\xea\x09\xfe\x21\x8e\xf4\x8d\x47\xcb\x7d\x9f\x70\x27\x16\x93\x75\x78\x9c\x50\x79\x2f\xa6\xdc\xb8\x20\x53\xf2\x95\x4d\xe9\x3f\xec\x2d\xcb\xf4\x43\xde\x9a\x29\x37\xae\xcf\x48\xfe\xb5\x47\xa4\x9f\x1d\x36\x49\xad\x05\x16\x3d\x5d\x4b\x85\xb7\xbc\xc1\xd1\x67\x9c\x80\xd9\x9f\x00\x00\x00\xff\xff\xae\xce\x58\xf1\xa3\x04\x00\x00")

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

	info := bindataFileInfo{name: "data/Dockerfile-lucid.template", size: 1187, mode: os.FileMode(420), modTime: time.Unix(1447206021, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataDockerfileTemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x52\xc1\x72\xd3\x30\x10\xbd\xe7\x2b\x76\x7c\x66\xe3\xe9\x14\x32\x70\xf0\xa1\x93\x10\xa6\x03\x24\x8c\x4b\xe0\x40\x39\xc8\xd6\xda\x15\xc8\xb2\x47\x92\x43\xd2\x92\x7f\x47\x96\x2d\xc7\x81\x92\x4b\x76\xdf\x3e\xad\xf5\x9e\xde\x3a\xdd\x7e\x84\xa7\xa7\xf9\x4a\x18\xab\xeb\xd3\x69\x96\xee\x36\xc0\x1a\x8b\x25\x59\x68\x1b\xce\x2c\x5d\x40\x42\x19\xcb\xa4\x04\x3c\x82\x6e\xb3\xe3\xd5\xfc\xcd\xfc\x1a\xb2\x56\x48\x8e\x64\x0c\x29\x2b\x98\x84\xfb\x19\xb8\x9f\x14\x59\xbe\x40\x4e\xfb\xae\x2a\x0a\x11\xca\x92\x67\x55\xa8\x55\xde\x6a\x43\xe6\x95\xef\xc7\x63\x9a\x18\x97\x42\x51\x60\x19\x23\x43\x79\x64\x55\x5f\x3f\xba\xe6\xaa\x9c\x9c\x73\x9b\xa4\xbf\xea\xb7\x28\x6e\x8d\x8e\x33\xa1\xe2\x92\xaa\xe8\x05\x44\xc3\xa5\xbb\xb2\x68\x3c\x82\xe8\xc6\x5c\xe8\x24\x50\x7b\x50\xd5\xa8\x79\x9d\x9f\x1b\x11\x7d\xf7\x3b\xbb\xe5\xde\xa6\xfa\x97\x92\x35\xe3\x3b\x2d\x4f\xa7\xdf\x96\x69\xa8\x0f\x8f\x4b\x88\x6d\xd5\xcc\xbe\x6e\xd3\xf7\xab\xdb\xd4\x37\x71\x67\x0e\xba\x03\xa9\xfb\xff\x42\xda\x88\x5a\x0d\xe6\x2e\xd7\x1f\x6e\xde\xdd\x25\x11\x56\x4c\xe7\x0f\xc9\xe1\xf5\x02\x17\x2f\x01\xb7\xd7\x11\xcc\xe3\xbc\x56\x85\x28\x5b\x4d\x5e\x14\x62\xa3\xa9\x10\x87\x24\xae\x1b\xeb\x57\xfe\xbd\x71\xa0\x91\x62\x99\x24\x34\x0f\x4c\x13\x1f\x30\x2e\x8c\x07\x07\xf1\xe8\x74\x5d\xb2\x3b\x1d\xa8\x49\x32\x2b\xf6\xfd\x23\x57\xec\x27\x01\xfe\x70\x1f\xd9\xb4\xd5\xf2\xd3\x6e\xb8\xb1\x87\xc3\xc3\xaf\xde\xde\x7d\x76\x2a\x13\xaf\xd2\xb9\x39\x3b\xeb\xf6\x64\x07\x0d\x0f\x82\x06\x9c\xc3\xa1\xb1\xc0\x29\x0b\x8d\x82\x67\xfd\x09\x63\xd6\x59\x7d\xe3\xdc\x39\x43\x7b\xf8\x0f\xd9\xc1\xb7\x96\xb4\x13\xe1\x1d\xf6\x64\x3e\xc9\xde\xfd\x04\x0a\x21\x9c\x62\x63\x1a\xa7\xe0\x33\xb1\xec\x07\x17\xc9\x9c\x0e\x42\x44\xa7\xd8\x98\xd5\x11\xfc\x27\xb4\xd8\x67\x27\x9e\xb8\xd6\x40\xec\x34\xad\x85\xa4\x0d\xab\x68\xd4\xe9\x12\x30\xfb\x13\x00\x00\xff\xff\x4d\xab\xf4\x3b\xad\x03\x00\x00")

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

	info := bindataFileInfo{name: "data/Dockerfile.template", size: 941, mode: os.FileMode(420), modTime: time.Unix(1447206027, 0)}
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

