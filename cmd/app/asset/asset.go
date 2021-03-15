// Code generated for package asset by go-bindata DO NOT EDIT. (@generated)
// sources:
// test/Makefile
// test/config.go
// test/main.go
package asset

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

var _testMakefile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8f\x41\x4b\xf3\x40\x10\x86\xcf\x99\x5f\xf1\x92\x2f\x87\x84\x8f\x6c\xef\x2b\x01\x83\x4d\x55\x28\x2a\x6d\x3c\x08\x42\x59\x37\x93\x4d\x60\xb2\x5b\xd2\x08\x0a\xfe\x78\x69\x94\xa8\xa7\xf7\x30\xef\x3c\xf3\x4c\x5d\xee\xae\xab\x5a\x17\x13\x9f\x26\xa2\xfd\xee\xea\xb0\xb9\xdd\x56\x7b\xe8\x02\x49\x7a\xea\x58\x04\x6d\xef\x1b\x28\xb5\x52\x0a\xf9\xf4\x7e\x64\xa0\xc5\x07\xdc\xc8\x47\xc4\xcf\xca\x85\x18\x40\x46\xff\x7e\x76\xff\x17\x18\x4c\xef\x95\x0b\x68\xc5\xb8\x73\xda\xe0\xdb\xde\x1d\xde\x06\x51\x2e\x10\xa9\x75\xb5\x29\x1f\xb7\xb5\x36\x22\x44\x46\x44\x27\xe9\x97\x49\x06\x8a\x2e\xd9\x76\x01\xf1\x3a\x78\x46\x29\xa2\x62\x10\x91\x63\x6f\x43\xc3\x56\x83\x22\xfb\xed\xb3\xba\x80\x0b\x70\xec\x79\x34\x13\x13\x2d\x10\x9d\xa4\x8b\x4e\x46\x91\x0b\x78\x79\xed\xa5\x41\x1e\xb0\x74\xfe\xde\x99\xe7\xbd\x77\x2a\x26\x22\xf5\x70\x73\x7f\xf7\xa4\x61\x85\x8d\x07\xcd\xa1\x29\x1a\x07\xe4\x63\xfb\x9b\x70\xfe\x6c\xee\xe4\xd6\xd8\x8e\xe9\x33\x00\x00\xff\xff\x24\x60\x9f\x30\x4e\x01\x00\x00")

func testMakefileBytes() ([]byte, error) {
	return bindataRead(
		_testMakefile,
		"test/Makefile",
	)
}

func testMakefile() (*asset, error) {
	bytes, err := testMakefileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/Makefile", size: 334, mode: os.FileMode(420), modTime: time.Unix(1615781269, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testConfigGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x5f\x4b\xc3\x30\x14\xc5\x9f\x73\x3f\xc5\x35\x30\x69\x60\x74\xef\x93\x3d\xe8\x54\x10\xf6\x30\xdc\xab\xe0\x62\x9a\x74\xc1\x36\x09\xb7\xc9\x40\xa4\xdf\x5d\xd2\x66\xe0\xbf\xc7\xfe\x38\xe7\x77\x4f\x13\xa4\x7a\x97\xad\xc6\x5e\x5a\x07\x60\xfb\xe0\x29\x62\x05\x8c\x9b\x3e\x72\x00\xc6\x5b\x1b\x4f\xe9\xad\x56\xbe\x5f\xdd\x25\x72\xf1\x90\x86\x93\x5d\x45\xdf\x77\x1c\x04\x40\xfc\x08\x1a\x95\x77\xc6\xb6\x4f\xce\x78\x1c\x22\x25\x15\xf1\x13\xd8\xf3\x7e\x7b\xdb\x34\x84\x19\x59\xd7\xe2\x31\x77\xd6\xbc\x60\x7e\x04\xb6\xf3\xed\xa3\xed\xf4\xef\x44\xc1\x25\xb1\x97\xf1\xf4\x4f\x22\xe3\x92\xd8\xe9\xb3\xee\xfe\x26\x26\xcc\x8f\x30\x02\x9c\x25\x95\x8d\xdf\xa6\x02\x98\xe4\x14\x76\x5e\x36\xdb\x09\x56\x26\x14\x8b\xc0\x4a\x13\xa1\x26\xf2\x24\xf2\xbf\x58\x83\xaf\xcb\xfc\x8d\x1b\xcc\x07\xea\x7b\xad\x7c\xa3\xf3\xcc\xca\x84\x25\x5e\xcf\x5a\x71\x33\x65\xae\x36\xe8\x6c\x97\x7b\xcc\xf4\xb1\xde\x93\x75\xd1\x54\x7c\xee\xe0\x7c\x0c\x1f\xb2\x7c\xbd\x18\xf0\xc5\xf1\x49\x5d\x4f\xa4\x12\x02\x18\x23\x1d\x13\x39\x60\x23\xc0\x0f\xc5\xce\xcb\xe6\x22\x30\xf9\xe9\x16\x03\x1e\x92\x52\x7a\x28\x1e\x13\x04\xb0\x79\xa7\xb3\x1d\x5c\x44\x23\x7c\x05\x00\x00\xff\xff\x65\x84\xe3\x8b\xe8\x01\x00\x00")

func testConfigGoBytes() ([]byte, error) {
	return bindataRead(
		_testConfigGo,
		"test/config.go",
	)
}

func testConfigGo() (*asset, error) {
	bytes, err := testConfigGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/config.go", size: 488, mode: os.FileMode(420), modTime: time.Unix(1615781043, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testMainGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x52\xdf\x6b\xdb\x3e\x10\x7f\xd6\xfd\x15\xf7\x15\x7c\x8b\x3d\x5c\xa5\x65\x6f\x19\xd9\xc3\xba\x15\x06\xdd\x16\x56\xda\x3d\x6c\xa3\x55\x1c\x49\x11\x95\x25\x23\xc9\x65\x34\xe4\x7f\x1f\xb2\x62\x3b\xde\x60\x2f\xc1\xb9\xfb\xdc\xe7\xc7\xe9\x5a\x5e\x3f\x71\x25\xb0\xe1\xda\x02\xe8\xa6\x75\x3e\x62\x01\x84\xca\x26\x52\x00\x42\x95\x8e\xbb\x6e\xc3\x6a\xd7\x2c\xea\x97\xf3\xb8\x13\x56\x2d\xea\x97\x27\x1d\xcf\x95\x5b\x18\xa7\xe8\x1c\x13\x5a\x79\xf9\x7a\x51\xbb\x8d\xe7\x14\x4a\x80\x67\xee\x13\xdd\xc3\xb3\xf0\x41\x3b\x8b\x88\xb8\x71\xce\x00\x79\xa8\x9d\x95\x5a\x5d\x6b\x23\x30\x44\xaf\xad\x1a\xe0\xde\xb9\x78\xf5\xe9\x3d\xae\xf0\xac\xe7\x61\x57\xae\x69\xb8\xdd\xee\x81\xdc\x05\xb1\x44\x44\x1a\x45\x88\xb4\x02\x72\xbb\x73\x3e\x2e\x91\x86\xc8\x7d\x44\x6e\x71\x68\xdc\x38\xab\x96\x88\x8f\xb3\xc6\x63\x05\xe4\x6b\x67\x13\xc3\x43\xca\x5b\xc1\x01\x40\x76\xb6\x46\x6d\x75\x2c\x4a\xdc\x03\x39\x8a\xb3\x75\xf2\x1b\xa2\xb0\xf1\xda\x70\x15\x8a\x92\xbd\x73\xce\xdc\x73\xbf\x2e\xce\x86\x30\x15\xd2\xe3\x17\x4d\x9f\xb4\x42\xc9\x4d\x10\x15\xd2\xd6\x6b\x1b\x71\xc8\xec\x64\xf6\x55\xfe\x8b\xfe\xb6\xdf\xc1\x51\x60\xda\x4d\x85\x34\xff\x49\x12\x75\xfa\xa1\x63\x09\x65\x5a\x5e\xcb\xe3\x0e\xa5\xf3\x83\xc6\x90\x29\x25\x9c\x67\xfa\xf0\x4b\xd4\x5d\x14\xc5\x84\xd9\x76\x4d\x7b\x9f\x5d\x66\xa8\x6c\x22\x5b\x27\xf3\xb2\xa0\xff\x87\x1f\x7d\xb0\x0b\x76\xc9\x2e\x4e\x88\xfb\xdd\x15\x75\xb3\xc5\x57\xb3\xf7\xa9\x90\x7b\x15\xf0\xfb\xcf\xfc\x9c\x3d\x9f\x96\x38\x3e\xfd\x1e\x08\x99\x09\x02\x21\x5e\xc4\xce\x5b\x20\x87\x1e\x6a\x84\x2d\x4e\xb2\x97\xf8\x16\x2f\xfa\x31\x2d\x51\x78\x8f\xcb\x15\x1a\xc7\xb7\x57\x3d\x62\x86\x7c\xd3\xf7\xff\x5b\xa1\xd5\xa6\x9f\x18\x99\x13\xf5\x01\x85\x09\xa2\xaf\xd7\xcd\x96\xdd\x05\xae\xc4\x1f\xf2\x40\x8c\xcb\xc7\xb8\x5c\xe5\xfb\x62\xf9\xb8\x8d\x53\xeb\xb4\xe1\x54\x66\x0b\x3a\xfa\xcc\xe2\xec\x26\xb7\x27\xab\x03\x7e\x85\x73\xc4\x69\xc6\xa9\x33\x8f\x39\x58\x38\x9d\x4d\x85\xd1\xe0\x67\xde\x88\x2f\x6d\xcc\x8b\x50\xec\x9b\x8e\xbb\x9b\x5c\x2d\x8e\xb3\xe5\xe8\xf8\x6f\x5c\xaa\x16\xc7\x6e\xc6\xb1\x8f\xe9\xf0\x27\xe2\x0a\xa7\xe1\x12\xe0\x00\xbf\x03\x00\x00\xff\xff\xad\x2b\x98\xf9\x1f\x04\x00\x00")

func testMainGoBytes() ([]byte, error) {
	return bindataRead(
		_testMainGo,
		"test/main.go",
	)
}

func testMainGo() (*asset, error) {
	bytes, err := testMainGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/main.go", size: 1055, mode: os.FileMode(420), modTime: time.Unix(1615781010, 0)}
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
	"test/Makefile":  testMakefile,
	"test/config.go": testConfigGo,
	"test/main.go":   testMainGo,
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
	"test": &bintree{nil, map[string]*bintree{
		"Makefile":  &bintree{testMakefile, map[string]*bintree{}},
		"config.go": &bintree{testConfigGo, map[string]*bintree{}},
		"main.go":   &bintree{testMainGo, map[string]*bintree{}},
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
