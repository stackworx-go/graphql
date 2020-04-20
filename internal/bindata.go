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

var _templateClientTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x92\x4d\x6f\xdb\x30\x0c\x86\xcf\xd6\xaf\x20\x74\x4a\x86\xd5\x02\x76\xdc\xb5\xcd\x86\x5e\x9a\x7d\x64\xbb\x14\x45\xa3\x38\x8c\xa3\xc5\x96\x14\x8a\xce\x1a\x18\xfe\xef\x03\xfd\x11\x18\xc3\x30\x1d\x6c\x93\x7e\x1f\xf2\x15\x25\x63\xe0\x3e\xec\x11\x4a\xf4\x48\x96\x71\x0f\xbb\x2b\x94\x8e\x8f\xcd\x2e\x2f\x42\x6d\x12\xdb\xe2\xf4\x3b\xd0\xdb\x5d\x19\x4c\x49\x36\x1e\xcf\xd5\x5d\x51\x39\xf4\xfc\x1e\x1e\xd6\xf0\xb4\xde\xc0\xea\xe1\x71\x93\xab\x68\x8b\x93\x2d\x11\xda\x16\xf2\x2f\xc3\xf7\x93\xad\x11\xba\x4e\x29\x57\xc7\x40\x0c\x0b\x95\xe9\x43\xcd\x5a\x01\x00\xe8\xdd\x95\x31\x69\x95\x69\xf4\x45\xd8\x3b\x5f\x9a\x5f\x29\x78\x49\xb8\x60\x5c\x68\xd8\x55\x12\x78\x64\x73\x64\x8e\x5a\xa9\x4c\xcf\x9c\x5d\xf0\xc4\xf6\x68\xca\x73\x15\x2d\x25\x24\x73\xf9\x20\x01\x12\x05\xd2\x6a\xa9\x14\x5f\x23\xc2\x7d\x6f\x15\x12\x53\x53\x30\xb4\x2a\x93\x52\xf9\x90\x55\xd9\x0f\xaa\xe4\x97\xf3\xa5\xea\x46\x80\xf0\xdc\x60\x9a\x13\x5f\x1b\xa4\x2b\x4c\x6b\x90\xc3\xdf\x6b\x2b\xde\x3f\xea\xb3\x68\xf5\x56\x65\xc6\xc0\x3a\xca\x44\x5d\xf0\xfd\x18\xfe\xcf\x85\xb9\x56\xf8\x9f\x96\x9c\xdd\x55\x98\x7a\x95\xf3\x8c\x74\xb0\x05\xb6\xdd\x44\x5c\x26\x81\xa8\x57\x6f\x8c\x3e\xb9\xe0\x7b\x79\x6d\xe3\xf3\xd0\xee\xe5\x1f\x20\xde\xb4\x7a\x7b\xdb\xb5\x3c\xfc\x68\x73\xdc\xf6\x66\x96\x12\xe7\x23\xfe\xfa\x3a\x69\x67\xf8\xe7\xe1\x62\xac\x64\xf6\xb3\x12\x7d\x9c\xe0\xf9\x65\x3a\x97\xbc\xcf\x08\x76\x68\x7c\x01\x0b\x84\x77\x73\x74\x09\xfd\x6b\xb1\x9c\x7a\xb6\x2a\x23\xe4\x86\x3c\xe8\x51\x07\xdf\xc6\xf3\xf9\x64\x5d\x85\x7b\x2d\xb5\xe4\xc6\x3d\xfa\xd8\xf0\xf7\xbe\x73\x82\xae\xfb\x13\x00\x00\xff\xff\x39\x69\x96\x4f\xd7\x02\x00\x00")

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

	info := bindataFileInfo{name: "template/client.tmpl", size: 727, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateRequestTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x4d\x4f\xdc\x30\x10\x3d\xdb\xbf\x62\x1a\x01\x4a\x50\x48\x51\x8f\x91\xa8\x04\xf4\xf3\x50\x44\x69\xe9\x05\x21\xad\x89\x27\x8b\x2b\xaf\x13\x6c\x07\xba\xb2\xfc\xdf\xab\x71\xb2\x6d\x60\x39\xc0\x2d\x9e\x79\xcf\x7e\xef\x79\x9c\x7b\x61\x21\x04\xa8\xce\xc4\x0a\x21\x46\x38\x82\x05\x2d\xbf\x0f\x68\xd7\x10\xe3\x82\x73\x5a\x9e\x8b\xb5\xee\x84\x84\x18\x39\xf7\xeb\x1e\xc1\xa2\xeb\x3b\xe3\x70\x4e\x75\xde\x0e\x8d\x87\xc0\xd9\x07\xe1\x05\xec\xcf\x7a\x1b\xfe\xe2\xb7\xeb\x4c\x9d\x49\xe1\x45\xb6\xe0\xec\xa3\xb5\x9d\x75\x70\x75\xbd\xbc\xd3\x48\xdf\x55\xaa\x4c\xb0\x54\x71\x0b\x1e\x39\x6f\x07\xd3\x40\xde\xc0\xfe\xa9\x56\x68\x7c\x31\x97\x9c\x73\x00\x80\x10\x0e\x40\xb5\x50\x7d\x11\xee\xab\xe9\x07\x1f\x63\x08\x56\x98\x25\xc2\x8e\x32\x12\xff\x94\xb0\x83\x1a\x57\x68\x3c\xd4\x47\x50\x1d\xdb\xe5\x40\x0b\x47\x38\xd5\x4e\xa0\x18\xcb\x10\xd0\x48\x2a\xfe\xf7\x45\xdf\x3f\xc9\x34\x95\x37\x5d\x34\x12\x0e\x62\xe4\x05\xe4\xcf\x18\x2d\x21\x89\x2f\x20\x00\x90\x3c\x66\xf1\x6e\x40\xe7\x4f\x3a\xb9\x4e\x3d\x12\x41\x26\xab\x6f\xc2\xba\x5b\xa1\xf3\x09\x10\x38\x63\x29\xfb\x7a\x6e\xb1\xe4\x8c\x3d\x35\x48\xca\x38\x63\xbf\x84\x55\xe2\x46\xa3\xab\x61\x25\xfa\x2b\xe7\xad\x32\xcb\x6b\x65\x3c\xda\x56\x34\x18\x22\xed\x98\xc8\x2f\x4b\x83\xd0\xd9\xec\xe8\x6c\x4b\x48\xda\x8c\xec\x27\x70\x2c\x53\xfc\x8f\x8a\xb1\xe0\x9c\xa9\x36\xf9\x7c\x73\x04\x46\x69\x1a\x0a\x66\xd1\x0f\xd6\xd0\x32\x45\xc0\x59\xe4\x9c\xd1\x20\xfd\x4b\xa4\xa9\xce\x3b\xe7\xf3\xa6\xba\xb4\xba\x84\x4c\xf4\xbd\x56\x8d\xf0\xaa\x33\x6f\x29\xac\xac\x84\x9b\xb5\x47\x57\x9d\xe1\xc3\xc9\xd0\xb6\x68\xf3\x59\xae\x45\x01\xf0\xc2\x73\x21\x21\x25\xb6\x68\xd3\x24\x57\xc4\xaf\x4e\x75\xe7\x30\x27\xed\x37\xf3\x6b\x52\xdd\xe0\x95\xae\x2e\x50\xc8\x63\x4d\x17\x35\xe1\x5f\x61\x52\xb5\xe3\x31\x3f\xbc\xf0\x83\x3b\xed\x24\x12\xe5\xdd\xe1\xe1\x16\xa5\x5d\xf9\xf1\x0d\xb4\x79\x76\x31\x7a\x83\x4f\x42\x69\x94\xf0\xa0\xfc\x2d\xb8\xb4\x03\x34\x9d\xc4\x1a\x76\x65\x09\x24\xb5\x86\xdd\xfb\xac\x7c\x7a\xc4\xd8\x2b\x46\x05\xf4\xcc\xfb\xe9\x11\x3e\xf3\x76\x39\x23\x1b\xd3\x48\x5e\x9a\xd5\x34\x94\x63\x0e\x7b\x13\xf1\x75\x86\x35\x9a\x7c\x22\x8e\x8e\x5c\x01\xef\x61\xdb\xf1\xde\x67\x2b\xfa\xdb\x3b\x9d\x30\x61\x44\xd6\xf0\x98\x19\x37\xb3\x92\x68\x9b\x1e\xfd\x65\x4a\xda\x84\x47\xfe\x37\x00\x00\xff\xff\x34\xcd\x6c\x30\xc7\x04\x00\x00")

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

	info := bindataFileInfo{name: "template/request.tmpl", size: 1223, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
