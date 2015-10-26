// Code generated by go-bindata.
// sources:
// main.go
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

var _mainGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x52\xcd\x8e\xd3\x30\x10\x3e\xdb\x4f\x31\xe4\xe4\x88\x2a\xb9\x57\xea\x61\x55\xa8\xc4\x61\x17\xb4\xe2\x86\x10\xf2\x3a\x93\xd4\x6a\x62\x47\x8e\xa1\x94\x2a\xef\xce\xcc\x74\x7f\x82\xca\xa5\x97\xc6\xe3\xf9\xe6\xfb\x19\x77\xb4\xee\x60\x3b\x84\xc1\xfa\xa0\xb5\x1f\xc6\x98\x32\x18\xad\x0a\x9b\xdc\xde\xff\xc2\x3a\xdb\x54\x50\xf9\x74\xca\x38\xf1\xc1\xc5\x61\x4c\x38\x4d\x75\xf7\xc7\x8f\x7c\xd1\x0e\x99\x3f\x3e\xf2\x6f\x1f\x3b\xfe\x44\x82\x96\x5a\xb7\x3f\x83\x13\x66\x53\xc2\x59\xab\xc6\x66\x0b\xeb\x0d\x7c\xfb\xce\x6c\xa6\x38\x9f\xcf\xd5\x3c\xcf\x05\x21\xb5\x4a\x2b\xc0\x94\xb8\xcd\xc4\xd5\x03\x1e\x1f\xd1\x36\x98\x8c\x28\x2f\x6a\x66\x29\x69\x44\xf9\x56\x26\xde\x6d\x20\xf8\x9e\xf9\x15\xa9\x57\x3b\x6a\xf7\x86\x1a\xa5\x56\x33\xa1\xb2\x70\x52\x8a\x05\x45\xe2\xf1\x36\x26\x19\xda\x37\x6f\xd2\x99\x51\xbf\xb3\xa1\xd9\x17\xfa\xcd\x06\x7c\xac\x3e\x7e\xde\x09\x58\xd5\x35\x60\x68\x20\xb6\x4c\x09\xcf\x4b\xe2\xc6\x53\x42\x7b\xa0\xc3\xfc\x36\xba\x70\x76\x65\x4d\x70\xd3\xd1\x67\xb7\x07\x72\x50\x7d\x3d\x8d\xd8\xf6\xb6\x13\xb4\xb3\x13\x8a\x65\xbe\xfd\xe0\xd3\x9a\x19\x68\xcf\xd5\x97\xe4\x43\xee\x83\x29\x1c\xa9\x65\x1f\xba\x35\x00\x14\xf0\x5e\x28\x1e\xec\x80\x1c\x4c\x29\xf1\x0d\x71\xaa\xee\x0f\x8d\x4f\x77\x7d\x6f\x5e\xfa\x2b\x41\xee\x7c\x8f\x9f\x42\x1b\x4d\x59\xdd\xc7\x06\x4d\x59\x42\x5d\x6f\x99\x92\x64\xf7\x08\x34\x84\x2e\xc7\x74\x12\xb6\xff\xa4\xb9\x8e\x23\xbb\xfe\xd7\xf8\x23\x76\xab\x65\x71\x77\x1d\x83\x76\x9d\xac\xbb\x04\xb9\x8e\x71\x7c\x7d\x17\x8a\x72\xb1\xf7\x1a\x64\xe1\xb8\xa5\x34\xe0\x03\x34\x38\x11\x13\xad\x25\x86\x5b\x7d\xab\x1f\x17\x29\x79\xeb\x6d\x1c\x4f\x86\xb4\x73\x12\x11\xaa\xc0\xc5\x90\x31\xe4\x09\x72\x04\x59\x3d\x36\x22\x7b\x93\x0e\x27\xaa\xb6\x7d\x9c\xd0\x3c\xff\x01\x66\x3d\xeb\xbf\x01\x00\x00\xff\xff\x41\x6e\x9f\x83\x81\x03\x00\x00")

func mainGoBytes() ([]byte, error) {
	return bindataRead(
		_mainGo,
		"main.go",
	)
}

func mainGo() (*asset, error) {
	bytes, err := mainGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "main.go", size: 897, mode: os.FileMode(438), modTime: time.Unix(1445739559, 0)}
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
	"main.go": mainGo,
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
	"main.go": &bintree{mainGo, map[string]*bintree{}},
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

