// Code generated by go-bindata.
// sources:
// torrc
// DO NOT EDIT!

package torctl

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

var _torrc = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x51\xc1\x8e\xd3\x30\x10\xbd\xfb\x2b\x46\xda\x0b\x48\x34\x12\x17\xee\x65\xbb\x2b\x2a\x0a\x44\x0a\xd2\x1e\x10\x07\x2b\x9e\x24\x56\x1d\x4f\x99\x71\x9b\x86\x36\xff\xce\x38\x24\x88\x3d\xc5\xf3\xe6\xbd\x37\xef\x29\x0f\xb0\x6f\x20\x52\xdc\xfc\x46\xa6\x77\x90\x78\x84\x44\x30\xb0\x4f\x98\x1f\xce\xcb\x11\x02\x8a\x40\xc3\xf8\xeb\x8c\x31\x05\x25\x74\x36\xc2\x80\x30\xd0\x39\x38\xa0\xd4\x21\x0f\x5e\xb0\x30\xdb\x0b\x79\xb7\x53\xc9\x4b\xd6\x0b\xbc\x37\xe6\x01\x5e\x74\x3d\x7b\x09\x46\x07\x81\xda\xd6\xc7\x16\x7a\xf5\xb4\x2d\x4a\x01\xf0\x4c\xdc\xdb\x04\x5e\xa0\xf7\xb1\xc2\x0b\xaa\x78\xfc\xb1\xe9\xed\x75\x1d\x7e\xaa\xcd\x1b\x49\x0e\x99\xef\xfa\xa1\x73\xba\xcb\x28\x6a\x75\x6f\x7c\x40\x78\xde\x1f\x9e\xbe\x6e\xbf\x3c\xbd\x2d\xcc\x81\x5a\x6d\x93\x7c\x8d\xf0\x97\x98\x13\x7c\xf4\x7a\x58\x03\xa4\x4e\x6f\x58\xe7\x38\xf7\xd1\x39\x78\x49\x18\xf3\xab\xa6\x18\xb1\x4e\x9e\x62\x2e\x4a\x3d\x54\xdf\x1e\x3f\x57\x1b\x39\xa1\x3d\x6a\x5a\xf5\xb0\xa7\x53\xf0\xb5\x9d\x29\x85\xa9\xa8\x3e\x4a\x49\x9c\xe0\x76\x83\xa2\x64\xba\x8e\xf3\x34\x4d\xb0\x2f\x2f\x1f\xbe\xb3\x6d\x1a\x5f\x43\xc9\xd8\x20\x67\xc4\x98\x47\x8a\x89\x29\x1c\xe6\x9b\xdb\x25\x44\x56\x2f\x8b\x15\x9a\xa6\x95\xfa\xcf\xff\xff\x59\xd7\xe6\x93\x95\x0e\xdd\x8a\x5a\x91\x81\xd8\xbd\x62\x2e\x58\x26\x66\xc5\xce\x26\xbb\xf3\xac\x0d\x49\xff\x6f\x26\x2e\x48\x5e\xfe\x09\x00\x00\xff\xff\xd2\xd8\x96\xed\x02\x02\x00\x00")

func torrcBytes() ([]byte, error) {
	return bindataRead(
		_torrc,
		"torrc",
	)
}

func torrc() (*asset, error) {
	bytes, err := torrcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "torrc", size: 514, mode: os.FileMode(436), modTime: time.Unix(1445830334, 0)}
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
	"torrc": torrc,
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
	"torrc": &bintree{torrc, map[string]*bintree{}},
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
