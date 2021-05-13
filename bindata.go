// Code generated by go-bindata. DO NOT EDIT.
// sources:
// templates/index.html
// templates/package.html
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

var _templatesIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x4f\x6f\xd3\x30\x14\xbf\xf3\x29\x1e\x56\x77\x23\x7e\x08\x71\x9a\x9c\x72\xa0\x68\x1c\x90\xa8\x06\x1c\x38\xba\xf6\x4b\xec\xd5\xb1\x2b\xdb\xe9\xa8\xa6\x7c\x77\xd4\xa4\x5b\xb7\x96\x64\x93\xd0\x72\x49\x9c\xf7\xfb\x63\x3f\xbf\xf7\xc4\xdb\xc5\xf7\xcf\x3f\x7f\x2f\xbf\x80\xc9\x8d\x9b\xbf\x11\xc3\x0b\x00\x40\x18\x92\x7a\xf8\xec\x97\xce\xfa\x35\x44\x72\x25\x4b\x79\xe7\x28\x19\xa2\xcc\xc0\x44\xaa\x4a\x66\x72\xde\xa4\x4b\x44\xa5\xfd\x4d\xe2\xca\x85\x56\x57\x4e\x46\xe2\x2a\x34\x28\x6f\xe4\x1f\x74\x76\x95\x30\xad\xc9\x51\x0e\x1e\x3f\xf0\xf7\xfc\xe3\xc3\x92\x37\xd6\x73\x95\x12\x03\x3c\x58\xe3\xd1\x5b\xac\x82\xde\x3d\xda\x86\xb6\x5b\x50\x4e\xa6\x54\x32\x15\x7c\x96\xd6\x53\x64\xc7\xf8\x29\x26\x86\xdb\x93\x68\x8f\xc8\x72\xe5\xe8\x1e\xd3\x16\x55\xeb\x5c\x71\x6b\x75\x36\xff\x00\x0f\x84\xa7\xd9\x38\x8f\xc7\xf1\xe0\x41\x60\xbe\x94\x6a\x2d\x6b\x12\x98\xcd\xf3\xe0\x1f\xa1\x8d\xea\x85\xd8\x45\x50\x6d\x43\x3e\xcb\x6c\x83\x9f\xa6\x08\x1c\xdb\xe9\x9e\x37\x7a\x46\x91\x9f\x5e\xc3\xe9\x73\x77\x07\x51\xfa\x9a\x60\xb6\xa6\xdd\x3b\x98\x6d\xa5\x6b\x09\x2e\x4b\xe0\x87\x43\x27\xe8\xba\x29\xf6\xcc\x36\x9b\x10\xf3\x52\x66\xb3\xa7\x6d\xa2\xf5\xb9\x02\x76\xb1\xc5\x8b\x2d\x83\x19\xff\x75\xfd\xad\xd7\x9e\x92\x79\xc1\x25\xe8\xf9\x89\x57\xd7\x09\xcc\x13\x17\x7b\x4f\x9b\x04\xf4\x20\x79\xe8\x05\xc4\xbd\x45\x9f\x00\x7e\x4d\x9b\x00\x5d\xc7\xe6\x67\xbf\x04\xca\x67\x4c\x5f\x65\x5b\xfc\x2a\xe8\xa0\xf8\xd7\x90\x32\x74\x1d\x9e\x25\x63\xa4\xfc\xcf\x54\x6d\x53\x43\x8a\x6a\x2f\x6b\x9b\x9a\x27\x63\xc9\xe9\xc4\x6d\xc0\x95\xd4\x35\x61\xbd\xb7\x29\x22\x55\x14\xc9\x2b\x2a\x56\xae\xa5\x4f\xfd\xdc\x28\xab\x10\x8b\x6c\xa8\xe8\x81\x0c\xa4\xcb\x25\xbb\x0a\x8b\xa0\x1e\xda\x7f\x3a\x2d\xff\x95\xb7\xf1\xf2\x87\xa1\x0e\xc9\xeb\xb1\x12\x13\x38\xd2\x04\x02\xfb\x71\x72\x32\x84\x50\xdb\xed\xa3\xb9\x75\x5c\x0a\x1c\x64\x04\x0e\xc3\xf6\x6f\x00\x00\x00\xff\xff\x27\x3c\xf8\x31\x84\x05\x00\x00")

func templatesIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesIndexHtml,
		"templates/index.html",
	)
}

func templatesIndexHtml() (*asset, error) {
	bytes, err := templatesIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/index.html", size: 1412, mode: os.FileMode(420), modTime: time.Unix(1620900995, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesPackageHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xd1\x31\x4f\xf3\x30\x10\x06\xe0\xfd\xfb\x15\xf7\x85\xb9\x35\x33\x38\x1d\x28\x88\xa5\x82\xaa\x82\x81\xd1\x24\x6f\x62\x4b\x8e\x2f\x38\x97\x4a\x28\xf2\x7f\x47\xa9\x2b\x15\x44\x11\x62\xb2\x7d\xb6\x1e\x9f\x5f\xeb\xff\xb7\x8f\xeb\xa7\x97\xed\x1d\x59\xe9\xfc\xea\x9f\xce\x03\x11\x91\xb6\x30\x75\x9e\x1e\x96\x1d\xc4\x50\x30\x1d\xca\xa2\xe5\x85\xeb\x7a\x8e\x52\x50\xc5\x41\x10\xa4\x2c\xa6\x89\x96\x6b\x13\x38\xb8\xca\xf8\xe7\xdd\x86\x52\xa2\xd6\x09\x59\x91\x7e\xb8\x52\x6a\xde\xdf\xa1\x67\x4a\xa9\xf8\x51\x1d\x78\x8c\x15\x7e\x51\xcf\x88\xe7\x6a\x4a\x22\x70\x28\xdc\x44\x13\x2a\x4b\x29\x4d\xaa\x76\xf1\x6f\x87\xd5\xd4\x38\x8f\x74\xb1\x99\xbc\x0b\xf8\xde\xfa\x6c\x2d\xf0\x36\xba\x7d\x59\x44\x34\x11\x83\xfd\xd4\xfe\xe5\x35\x8d\xd1\x97\xb3\x7b\xcf\x35\x57\xf9\x05\x47\x44\xab\x53\xc2\xfa\x95\xeb\xf7\x93\xfd\xc0\x62\x5d\x68\x49\x98\x06\x80\x2c\x22\x96\xb4\xf5\x30\x03\x48\x1b\xb2\x11\x4d\xce\xe6\x8b\xda\xf1\x1e\x64\x3c\x87\x56\x2b\xb3\x5a\x1e\x2f\xc9\xb2\x56\xf9\x67\x3f\x02\x00\x00\xff\xff\x0b\xa0\xa0\x2d\xf1\x01\x00\x00")

func templatesPackageHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesPackageHtml,
		"templates/package.html",
	)
}

func templatesPackageHtml() (*asset, error) {
	bytes, err := templatesPackageHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/package.html", size: 497, mode: os.FileMode(420), modTime: time.Unix(1620907647, 0)}
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
	"templates/index.html":   templatesIndexHtml,
	"templates/package.html": templatesPackageHtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"index.html":   &bintree{templatesIndexHtml, map[string]*bintree{}},
		"package.html": &bintree{templatesPackageHtml, map[string]*bintree{}},
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
