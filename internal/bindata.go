// Code generated for package internal by go-bindata DO NOT EDIT. (@generated)
// sources:
// template/client.tmpl
// template/request.tmpl
package internal

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateClientTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xc1\x8e\xd3\x30\x10\x86\xcf\x99\xa7\x18\xf9\xd4\x22\x36\x96\x38\x72\xdd\x2d\x88\xcb\x16\xd0\xc2\x65\xb5\x52\xdd\x74\xea\x9a\x3a\xb6\x3b\x9e\x94\x8d\xa2\xbe\x3b\x72\xda\xac\x22\x84\xf0\x25\x99\xc9\xf7\x39\x33\xbf\xd6\x78\x1f\x77\x84\x96\x02\xb1\x11\xda\xe1\xb6\x47\xeb\xe4\xd0\x6d\xeb\x26\xb6\x3a\x8b\x69\x8e\xbf\x23\xbf\xde\xd9\xa8\xed\xc9\x5b\x0a\x77\x4c\xde\xf4\xef\xf1\x61\x8d\x8f\xeb\x27\x5c\x3d\x7c\x79\xaa\x21\x99\xe6\x68\x2c\xe1\x30\x60\xfd\xf5\xfa\xfe\x68\x5a\xc2\xcb\x05\xc0\xb5\x29\xb2\xe0\x02\x2a\xb5\x6f\x45\x01\x22\xa2\xda\xf6\x42\x59\x41\xa5\x28\x34\x71\xe7\x82\xd5\xbf\x72\x0c\xa5\xe1\xa2\x76\xb1\x13\xe7\x4b\x11\x48\xf4\x41\x24\x29\x80\x4a\xcd\xe6\x3a\xd3\x51\xcc\xa1\x4c\x94\x0c\x67\x62\x7d\xfe\x50\x0a\x62\x8e\xac\x60\x09\x20\x7d\x22\xbc\xf7\x8e\x82\x60\x16\xee\x1a\xc1\x01\xaa\x72\x55\x7d\xed\x42\xf5\x83\x7d\xf9\xe4\x82\x85\xcb\x4d\x60\x3a\x75\x94\xe7\xc6\xb7\x8e\xb8\xc7\xe9\x5c\x71\xfc\xfb\x6c\xca\xec\x1f\xd5\xa9\xb0\x6a\x03\x95\xd6\xb8\x4e\x25\x4f\x17\xc3\x18\xc3\xff\xbd\x38\x67\x8b\xff\xd3\xb0\x33\x5b\x4f\x79\xa4\x5c\x10\xe2\xbd\x69\x68\xb8\x4c\xc6\x79\x02\x0a\xbd\x7a\x15\x0a\xd9\xc5\x30\xe2\xad\x49\xcf\xd7\xdf\xbd\xfc\x43\xa4\x37\x56\x6d\xde\xb6\xfe\xcc\x26\x1d\x4e\x7e\x55\xc2\x9b\xad\x3e\xd6\x19\x9f\x5f\xa6\x60\xeb\xb1\x53\xb4\x7d\x17\x1a\x5c\x10\xbe\x9b\xab\x4b\x1c\x1f\x8b\xe5\xb4\xee\x00\x15\x93\x74\x1c\x50\xdd\x38\xfc\x7e\x0b\xf8\x93\x71\x9e\x76\xaa\xdc\xf5\x27\x00\x00\xff\xff\x80\x6e\x6e\x27\x83\x02\x00\x00")

func templateClientTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateClientTmpl,
		"template/client.tmpl",
	)
}

func templateClientTmpl() (*asset, error) {
	bytes, err := templateClientTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/client.tmpl", size: 643, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateRequestTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\x4f\x6b\xdc\x3e\x14\x3c\x4b\x9f\x62\x7e\x86\x04\x3b\x38\xfe\x85\x1e\x0d\x29\x34\xdb\xbf\x87\x86\x34\x25\xbd\x94\xc2\x6a\xed\xe7\xae\x8a\xd6\xf2\x4a\x72\xc2\x22\xf4\xdd\x8b\x64\x6f\x71\x93\x1c\x9a\x9b\xf5\xde\x8c\x66\xe6\xf9\xe9\x5e\x18\x78\x8f\xea\x5a\xec\x08\x21\xe0\x12\xeb\x78\xfc\x32\x92\x39\x20\x84\x35\xe7\xde\x43\x76\xa8\x3e\xf5\xc3\xe8\x70\x1e\x02\x62\x65\x3e\x86\xc0\xbd\x3f\x07\xf5\x6d\xfc\x4c\x8d\x1b\x71\x50\x5a\x4c\x67\x77\x18\x08\x86\xec\xa0\x7b\x4b\x4b\x15\xeb\xcc\xd8\x38\x78\xce\xde\x0a\x27\x70\xb6\xe8\x1d\xf9\xeb\x5f\x56\xf7\x75\xd6\x0a\x27\xb2\x35\x67\xef\x8c\xd1\xc6\xe2\xfb\x8f\x9f\x7b\x45\xf1\xbb\x4a\x95\x19\x96\x2a\x76\xcd\x03\xe7\xdd\xd8\x37\xc8\x1b\x9c\xad\x94\xa4\xde\x15\xcb\x74\x39\x07\x80\xe8\x38\x26\xfa\x28\x6c\x4a\x11\x82\x4c\x61\x16\xc0\x54\xf7\x3e\x05\x3b\x0f\x81\x17\xc8\x9f\xf1\x58\x22\xe9\x16\xf0\x40\xbc\x99\x19\xda\x8f\x64\xdd\x95\x6e\x0f\xa9\x87\xfa\x12\xd1\x5f\xf5\x59\x18\xbb\x15\x2a\x9f\x01\x9e\x33\x96\x26\x5c\x2f\x45\x4b\xce\xd8\x63\x6f\x71\x58\x9c\xb1\x6f\xc2\x48\xb1\x51\x64\x6b\x24\xaf\x65\x0a\xc2\x66\x87\x21\x70\x16\x0a\xce\x99\xec\x92\xea\x7f\x97\xe8\xa5\x8a\xd3\x65\x86\xdc\x68\xfa\x78\x4c\x86\x38\x0b\x9c\xb3\xf8\x47\xfe\xf8\x6b\xaa\x1b\x6d\x5d\xde\x54\x77\x46\x95\xc8\xc4\x30\x28\xd9\x08\x27\x75\xff\x7f\xb4\x9e\x95\xd8\x1c\x1c\xd9\xea\x9a\x1e\xae\xc6\xae\x23\x93\x2f\x52\x16\x05\xf0\x8f\xba\x48\xc8\x96\x3a\x32\x69\x25\xaa\xc8\xaf\x56\x4a\x5b\xca\xa3\xf7\xcd\x72\x68\x52\x8f\x4e\xaa\xea\x96\x44\xfb\x46\xc5\xb1\xcd\xf8\x17\x84\x94\xdd\x24\xf3\xd5\x09\x37\xda\x95\x6e\x29\x52\x5e\x5d\x5c\x3c\xa1\x74\x3b\x37\x2d\x53\x97\x67\xb7\x53\x36\xbc\x17\x52\x51\x8b\x07\xe9\xb6\xb0\xe9\x06\x34\xba\xa5\x1a\x27\x6d\x89\x68\xb5\xc6\xc9\x7d\x56\x3e\x96\x98\x7a\xc5\xe4\x20\x3e\xad\x61\xde\xe6\x67\x1e\x01\x67\x31\xc6\xbc\x20\x77\xfd\x6e\x5e\x91\x69\x0e\xa7\x33\xf1\x65\x81\x15\xf5\xf9\x4c\x9c\x12\xd9\x02\xaf\xf1\x34\xf1\xe9\x07\x23\x86\xed\x5e\x25\x8c\x9f\x90\x35\xfe\x66\x86\xe3\xae\x24\xda\xb1\x17\x9f\x6b\x19\x2f\xe1\x81\xff\x0e\x00\x00\xff\xff\x2d\x2c\x19\x68\x3b\x04\x00\x00")

func templateRequestTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateRequestTmpl,
		"template/request.tmpl",
	)
}

func templateRequestTmpl() (*asset, error) {
	bytes, err := templateRequestTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/request.tmpl", size: 1083, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"template/client.tmpl":  templateClientTmpl,
	"template/request.tmpl": templateRequestTmpl,
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
	"template": &bintree{nil, map[string]*bintree{
		"client.tmpl":  &bintree{templateClientTmpl, map[string]*bintree{}},
		"request.tmpl": &bintree{templateRequestTmpl, map[string]*bintree{}},
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
