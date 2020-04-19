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

var _templateClientTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x92\x4d\x6f\xdb\x30\x0c\x86\xcf\xe6\xaf\x20\x74\x4a\x86\xd5\x02\x76\xdc\xb5\xcd\x86\x5e\x9a\x7d\x74\xbb\x14\x05\xa2\x38\x8c\xa3\xc5\x96\x14\x8a\xca\x6a\x18\xfe\xef\x83\x9c\x38\x30\x86\x61\xbc\xd8\xa4\x9f\x97\x1f\x2f\xac\x35\xde\xfb\x1d\x61\x4d\x8e\xd8\x08\xed\x70\xdb\x61\x6d\xe5\x90\xb6\x65\xe5\x5b\x1d\xc5\x54\xc7\xdf\x9e\xdf\xee\x6a\xaf\xeb\x53\x53\x93\xbb\x63\x6a\x4c\xf7\x1e\x1f\xd6\xf8\xb4\x7e\xc6\xd5\xc3\xe3\x73\x09\xc1\x54\x47\x53\x13\xf6\x3d\x96\x5f\x2e\xef\x4f\xa6\x25\x1c\x06\x00\xdb\x06\xcf\x82\x0b\x28\xd4\xbe\x15\x05\x88\x88\x6a\xdb\x09\x45\x05\x85\x22\x57\xf9\x9d\x75\xb5\xfe\x15\xbd\xcb\x05\xeb\xb5\xf5\x49\x6c\x93\x13\x47\xa2\x0f\x22\x41\x01\x14\x6a\xb6\xd7\x99\x8e\x62\x0e\x79\xa3\x60\x38\x12\xeb\xf3\x87\x9c\x10\xb3\x67\x05\x4b\x00\xe9\x02\xe1\x7d\x63\xc9\x09\x46\xe1\x54\x09\xf6\x50\xe4\x56\xe5\xa5\x0a\xc5\x0f\x6e\xf2\x27\xeb\x6a\x18\xae\x02\xa6\x53\xa2\x38\x57\x7c\x4d\xc4\x1d\x4e\x71\xc1\xf1\xef\xd8\xe4\xdd\x3f\xaa\x53\x66\xd5\x06\x0a\xad\x71\x1d\xb2\x9f\xd6\xbb\xd1\x86\xff\xeb\xfc\x9c\xcd\xfa\x9f\x86\xad\xd9\x36\x14\x47\xca\x3a\x21\xde\x9b\x8a\xfa\x61\x52\x9c\x27\x20\xd3\xab\x37\x21\x17\xad\x77\x23\xde\x9a\xf0\x72\x19\xf7\xfa\x0f\x21\xdd\x58\xb5\xb9\x5d\xfd\x99\x4d\x38\x9c\x9a\x55\x36\x6f\x76\xfa\x98\x47\x7c\x79\x9d\x8c\x2d\xc7\x4a\x96\xed\x93\xab\x70\x41\xf8\x6e\x2e\x5d\xe2\xf8\x58\x2c\xa7\x73\x7b\x28\x98\x24\xb1\x43\x75\xe5\xf0\xdb\xd5\xe0\x4f\xc6\x36\xb4\x53\xb9\x57\xfe\x65\x1e\x5d\x48\xf2\x7d\x9c\x1c\x71\x18\xfe\x04\x00\x00\xff\xff\x6c\x26\xd0\xf2\x96\x02\x00\x00")

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

	info := bindataFileInfo{name: "template/client.tmpl", size: 662, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
