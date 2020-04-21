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

var _templateClientTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x92\xcf\x8e\xd3\x30\x10\xc6\xcf\xf1\x53\x8c\x7c\x6a\x11\xc4\x12\x47\xae\xbb\x05\x56\x42\x5b\x58\x0a\x97\xd5\x6a\xeb\x26\xd3\xd4\x34\xb1\x5d\x7b\x52\x5a\x45\x79\x77\x34\xf9\x53\x59\x08\xe1\x43\xed\x99\xfe\xbe\x99\x7c\x63\x2b\x05\x77\xae\x44\xa8\xd0\x62\xd0\x84\x25\xec\xae\x50\x19\x3a\xb4\xbb\xbc\x70\x8d\x8a\xa4\x8b\xe3\x6f\x17\x2e\xef\x2a\xa7\xaa\xa0\xfd\xe1\x54\xbf\x85\xfb\x35\x3c\xae\x37\xb0\xba\x7f\xd8\xe4\xc2\xeb\xe2\xa8\x2b\x84\xae\x83\xfc\xeb\x78\x7e\xd4\x0d\x42\xdf\x0b\x61\x1a\xef\x02\xc1\x42\x64\x72\x77\x25\x8c\x52\x64\x12\x6d\xe1\x4a\x63\x2b\xf5\x2b\x3a\xcb\x89\x7d\x43\xbc\x19\xa7\x8c\x6b\xc9\xd4\x1c\x58\x24\x75\x20\xf2\x52\x88\x4c\x26\xdf\x73\xc6\x23\xe9\x83\xaa\x4e\xb5\xd7\x21\x62\x50\xe7\xf7\x1c\x60\x08\x2e\x48\xb1\x14\x82\xae\x1e\xe1\xae\x36\x68\x09\x22\x85\xb6\x20\xe8\x44\xc6\xa5\xf2\x31\x2b\xb2\x1f\x4f\x5f\xf8\x2f\x63\x2b\xd1\x4f\x82\x80\xa7\x16\x63\xaa\xf8\xd6\x62\xb8\xc2\xbc\x46\x1c\xfe\x5e\x5b\xb6\xf0\x41\x9e\x98\x95\x5b\x91\x29\x05\x6b\xcf\x73\x34\xce\x0e\x33\xf8\xbf\xce\xa5\x2c\xeb\x7f\xea\x60\xf4\xae\xc6\x38\x50\xc6\x12\x86\xbd\x2e\xb0\xeb\x67\xc5\x79\x06\x98\x5e\x5d\x08\x6d\x34\xce\x0e\x78\xa3\xfd\xf3\xd8\xee\xe5\x1f\x42\xbc\xb1\x72\x7b\x73\xfd\x34\xba\x5e\x7b\x8a\x89\xf3\xcf\xa8\x4b\x0c\x30\x8c\x6c\x3c\xdf\x04\xfc\x63\x27\x5f\x13\xbd\x49\x52\x6c\x75\xea\xf7\xfa\x3a\xb3\x49\xbf\x4f\xe3\xfb\x59\xf1\x65\x25\x25\x86\x38\xc2\xf3\xcb\x7c\x91\xf9\x90\x61\xd9\xbe\xb5\x05\x2c\x10\xde\xa4\xd2\x25\x0c\xdb\x62\x39\xf7\xec\x44\x16\x90\xda\x60\x41\x4e\xdc\x6c\x0d\x3e\x6a\x53\x63\x29\xb9\x16\xbf\xcf\x07\xeb\x5b\xfa\x3e\x74\x8e\xd0\xf7\x7f\x02\x00\x00\xff\xff\xc6\x4f\x64\x2d\xfe\x02\x00\x00")

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

	info := bindataFileInfo{name: "template/client.tmpl", size: 766, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateRequestTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x5b\x4f\xdc\x3c\x10\x7d\xb6\x7f\xc5\x7c\x11\xa0\x04\x85\x7c\xa8\x8f\x2b\x6d\x25\xba\xb4\xa5\x52\x0b\x94\x4b\x5f\x10\xd2\x9a\x64\xb2\xeb\xca\xeb\x64\x6d\x07\xba\xb2\xfc\xdf\xab\x71\xb2\x34\x5c\x1e\xe0\x2d\x93\x39\x67\x7c\xce\xf1\xe5\x5e\x18\xf0\x1e\x8a\x53\xb1\x42\x08\x01\xa6\x30\xa7\xf2\x67\x87\x66\x03\x21\xcc\x39\xa7\xf2\x5c\x6c\x54\x23\x2a\x08\x81\x73\xb7\x69\x11\x0c\xda\xb6\xd1\x16\xc7\x54\xeb\x4c\x57\x3a\xf0\x9c\x1d\x0b\x27\x60\x7f\xd4\xdb\xf2\xe7\xbf\x6d\xa3\x27\x49\x25\x9c\x48\xe6\x9c\x7d\x36\xa6\x31\x16\x6e\x6e\x17\x6b\x85\xf4\x5d\xc4\x3f\x5b\x58\xfc\x65\x93\x39\x0f\x9c\xd7\x9d\x2e\x21\x2d\x61\x7f\xa6\x24\x6a\x97\x8d\x45\xa7\x1c\x00\xc0\xfb\x03\x90\x35\x14\x27\xc2\x7e\xd3\x6d\xe7\x42\xf0\xde\x08\xbd\x40\xd8\x91\xba\xc2\x3f\x39\xec\xa0\xc2\x15\x6a\x07\x93\x29\x14\x47\x66\xd1\x51\x61\x09\xf7\xcf\x04\x7d\x5f\x91\xc3\x10\x72\xef\x51\x57\xb1\x8d\xba\x82\x83\x10\x38\x6b\x5a\x67\x61\xff\x02\xd7\x1d\x5a\x77\xd6\x3a\x9b\x41\xfa\x8a\xd1\x1c\xa2\xf6\x0c\x3c\x00\x89\x63\xa6\x67\x7c\x6a\xaa\x4d\xec\x91\x04\x32\x59\xfc\x10\xc6\x2e\x85\x4a\x07\x80\xe7\x8c\xc5\xec\x27\x63\x83\x39\x67\xec\xb9\x3d\x12\xcb\x19\xfb\x25\x8c\x14\x77\x0a\xed\x04\x56\xa2\xbd\xb1\xce\x48\xbd\xb8\x95\xda\xa1\xa9\x45\x89\x3e\xd0\xc4\x48\x7e\x5b\x16\x84\x4e\x46\x4b\x27\x2f\x84\xc4\x61\x14\x48\x04\x87\x3c\x86\xff\xe4\x67\xc8\x38\x67\xb2\x8e\x3e\xff\x9b\x82\x96\x8a\x0e\x05\x33\xe8\x3a\xa3\xa9\x8c\x11\x70\x16\x38\xa7\x60\x1e\x03\x59\x3a\xd7\x16\xa7\xf8\x30\xc4\x9b\x26\xe7\x67\x97\x57\x49\x0e\x65\x71\x7d\xf1\x3d\x87\xbb\x8d\x43\xdb\xf7\x45\x85\x26\x1d\x65\x9a\xbd\x75\x45\x00\xe8\x91\x71\x23\x9f\x40\xd7\xc5\x49\x9c\x0b\xd3\xd8\x1c\xaa\x47\x95\x43\x5d\x5c\xa2\x4b\x93\x59\xa3\x1d\x6a\x77\x40\x27\x25\xc9\x21\x11\x6d\xab\x64\x29\x9c\x6c\xf4\xff\xb4\xad\x49\x16\x49\xb6\x7d\xf4\x56\x16\xc7\x0d\x29\x7e\xb3\x50\xce\x59\x85\x35\x9a\x78\xd5\x0a\x32\x59\xcc\x54\x63\x31\xa5\x09\x77\xe3\x73\x24\x9b\xce\x49\x55\x50\x2a\x47\x8a\x4e\xd2\x80\x7f\xc7\x2e\xc8\xba\x5f\xe6\xd2\x09\xd7\xd9\x59\x53\x21\x51\x3e\x1c\x1e\xbe\xa0\xd4\x2b\xd7\x5f\xd2\x3a\x4d\x86\x7d\x82\x2f\x42\x2a\xac\xe0\x41\xba\x25\xd8\x38\x01\xca\xa6\xc2\x09\xec\x56\x39\x90\xd4\x09\xec\xde\x27\xf9\xf3\x25\xfa\x5e\xd6\x2b\xa0\x77\xa8\x1d\x5e\x89\x57\x1e\x17\xce\xc8\xc6\x70\x67\xae\xf5\x6a\xb8\x35\x7d\x0e\x7b\x03\xf1\x7d\x86\x15\xea\x74\x20\xf6\x8e\x6c\x06\x1f\xe1\xa5\xe3\xbd\xaf\x46\xb4\xcb\xb5\x8a\x18\xdf\x23\x27\xf0\x94\x19\xb6\xc7\x24\xd2\xb6\x3d\x7a\x06\x73\x1a\xc2\x03\xff\x1b\x00\x00\xff\xff\x8f\x89\xca\xfa\x68\x05\x00\x00")

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

	info := bindataFileInfo{name: "template/request.tmpl", size: 1384, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
