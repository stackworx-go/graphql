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

var _templateClientTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\x41\x4f\xf3\x30\x0c\x86\xcf\xcd\xaf\xb0\x72\xfa\x3e\x89\x35\x77\xae\xdb\x0e\x5c\x36\x90\x06\x17\x84\xb4\xb4\x33\x59\x58\x6b\x67\x89\x0b\xab\xaa\xfe\x77\x94\x95\xa1\x09\x21\x7c\x4a\xec\xe7\x51\xf2\xda\x18\x98\xf3\x0e\xc1\x21\x61\xb4\x82\x3b\xa8\x7a\x70\x5e\xf6\x5d\x55\xd6\xdc\x9a\x24\xb6\x3e\x7c\x70\x3c\xcd\x1c\x1b\x77\x6c\x1c\xd2\x2c\x62\x63\xfb\x1b\x58\xac\x61\xb5\xde\xc0\x72\x71\xb7\x29\x55\xb0\xf5\xc1\x3a\x84\x61\x80\xf2\x7e\x3a\xaf\x6c\x8b\x30\x8e\x4a\xf9\x36\x70\x14\xf8\xa7\x00\x00\x74\xd5\x0b\x26\xad\x0a\x8d\x54\xf3\xce\x93\x33\x6f\x89\x29\x37\x3c\x1b\xcf\x9d\xf8\x26\x5f\x08\xc5\xec\x45\x82\x56\xff\x95\x92\x3e\x20\xcc\x1b\x8f\x24\x90\x24\x76\xb5\xc0\xa0\x8a\x3c\x2e\xa7\xae\x2a\x1e\x63\x93\x47\x9e\x9c\x1a\xbf\x84\x88\xc7\x0e\xd3\xb5\xf1\xd0\x61\xec\xe1\x52\x13\x0e\x3f\x6b\x9b\xff\x73\xab\x8f\x99\xd5\x5b\x55\x18\x03\xeb\x90\x97\xe3\x99\xce\x99\xfe\xf6\xf8\x9a\xcd\xfe\x93\x8d\xde\x56\x0d\xa6\x33\xe5\x49\x30\xbe\xda\x1a\x87\xf1\x62\xbc\x5f\x80\x4c\x2f\x4f\x82\x94\x3c\xd3\x19\x6f\x6d\x78\x9e\x9e\x7b\xf9\x45\xc4\x6f\x56\x6f\x73\xea\xcf\x00\x00\x00\xff\xff\xdd\x30\x8c\xde\xcd\x01\x00\x00")

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

	info := bindataFileInfo{name: "template/client.tmpl", size: 461, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateRequestTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\xc1\x8e\xd4\x30\x0c\x3d\xc7\x5f\x61\xf6\x80\xda\x55\x27\xdc\x47\x9a\x03\xbb\x17\x38\xb0\x5a\x90\x96\xf3\x66\x5a\x57\x04\x65\x92\x92\xa4\xa0\x2a\xf2\xbf\x23\x67\x02\xaa\xe0\x34\xb7\xda\xcf\x7e\xef\xf9\xa5\x3f\x4d\xc4\x52\x50\x3f\x99\x0b\x21\x33\x9e\xf0\x55\xca\xcf\x2b\xc5\x0d\x99\x5f\x01\x4a\x41\x3b\xa3\xfe\xe8\x97\x35\xe3\x81\x19\xa5\xd3\x4a\x66\x28\xe5\x80\xe4\x27\xf9\xac\xc0\xb3\xd9\x5c\x30\xd7\x7a\x5e\xfd\x88\xdd\x88\xf7\x8f\xce\x92\xcf\xfd\x5e\xa9\x03\x44\x44\xd9\x16\xf6\x0f\x26\x55\x46\x66\x5b\x89\x77\x83\xb5\x5f\x4a\x15\x39\x30\x43\x8f\xdd\xfd\x0e\x6e\x7a\x03\x52\x8c\x21\xf6\x58\x10\x85\x59\x45\xfa\xb1\x52\xca\x0f\x61\xda\x2a\x86\xc7\x13\x7e\x4f\xc1\xeb\x4f\x26\xa6\x6f\xc6\x75\x6d\xa0\x80\x52\xf5\xda\xe3\x5e\x74\x00\xa5\xfe\xf5\x26\xf1\x80\x52\x5f\x4d\xb4\xe6\xec\x28\x1d\xb1\x7a\x1d\xea\x21\xaa\x39\x64\x06\xc5\x3d\x80\xb2\x73\x55\x7d\x73\x42\x6f\x1d\x8a\x4a\xa4\xbc\x46\x2f\x65\x35\x04\x8a\x01\x54\xa4\xb4\xfc\xf5\x37\xea\xe7\x90\x72\x37\xea\x97\xe8\x06\xbc\x33\xcb\xe2\xec\x68\xb2\x0d\xfe\x9d\x58\xbf\x1b\xf0\xbc\x65\x4a\xfa\x89\x7e\x3d\xac\xf3\x4c\xb1\xdb\x5d\xd9\xf7\xe2\x03\x01\x40\x4d\x34\x53\x44\xa1\xd6\x82\xe8\x47\x17\x12\x75\xe2\xea\xbc\x8f\xc3\x86\x35\x5b\xa7\xbf\x90\x99\xde\x3b\x09\xa4\xcd\xdf\x60\x5f\x7e\x9f\xa5\xbd\xf8\xff\x8f\x02\x4a\x38\x5a\xee\x2f\xfe\xd2\x92\xbf\x9a\x78\xdb\xf6\x6e\x0a\xab\x36\xff\x6c\x0e\x02\x03\xc3\xef\x00\x00\x00\xff\xff\xe6\x9e\x58\x55\xc4\x02\x00\x00")

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

	info := bindataFileInfo{name: "template/request.tmpl", size: 708, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
