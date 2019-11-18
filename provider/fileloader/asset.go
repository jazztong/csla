// Code generated by go-bindata. DO NOT EDIT.
// sources:
// template/aws-api-lambda-golang/.gitignore.tmpl
// template/aws-api-lambda-golang/Makefile.tmpl
// template/aws-api-lambda-golang/main.go.tmpl

package fileloader


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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataTemplateAwsapilambdagolangGitignoretmpl = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x08\x49\x2d\x2e\x51\x28\x4a\x2d\x2e\xcd\x29\xe1\x4a\xce\x2f\x4b" +
	"\x2d\x4a\x4c\x4f\xd5\xcb\x2f\x2d\xe1\x2a\x4a\x2d\xc8\x2f\x2a\xd1\xcb\x2a\xce\xcf\xe3\x2a\x49\x2d\x2e\xd1\xab\xc8" +
	"\xcd\x01\x04\x00\x00\xff\xff\x06\x42\xc9\x12\x2f\x00\x00\x00")

func bindataTemplateAwsapilambdagolangGitignoretmplBytes() ([]byte, error) {
	return bindataRead(
		_bindataTemplateAwsapilambdagolangGitignoretmpl,
		"template/aws-api-lambda-golang/.gitignore.tmpl",
	)
}



func bindataTemplateAwsapilambdagolangGitignoretmpl() (*asset, error) {
	bytes, err := bindataTemplateAwsapilambdagolangGitignoretmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "template/aws-api-lambda-golang/.gitignore.tmpl",
		size: 47,
		md5checksum: "",
		mode: os.FileMode(493),
		modTime: time.Unix(1573661157, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataTemplateAwsapilambdagolangMakefiletmpl = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xce\x41\xae\xc3\x30\x08\x04\xd0\xf5\xf7\x29\xe6\x02\x90\x7d\xa4\x7f" +
	"\x98\x28\xa1\xc8\x55\x6c\x22\x8c\xa3\xf6\xf6\x95\xdd\xee\x78\x33\x2c\x26\xa4\xc5\x9a\xfe\xd4\xa0\x12\x50\x1b\xe6" +
	"\x30\x3b\xdb\xf2\x45\xeb\x65\xd4\xbf\x13\x44\xcf\x5e\x73\x3c\xf2\x29\x98\xbf\xaf\x72\x8e\xb0\x59\x9d\x99\xcb\x65" +
	"\x1e\x3c\x0c\x22\xd0\x6e\xb7\xf8\xe5\x36\xca\xff\x89\x4d\x85\xad\x07\xc8\xb7\x5d\x40\x37\x78\x61\xe6\x39\xa1\xd8" +
	"\x81\xc8\xc7\x3b\x25\xef\x75\x4d\x00\xa0\x06\xef\x15\x65\xcb\x95\xd5\x3e\x01\x00\x00\xff\xff\xd8\x96\x89\x01\xae" +
	"\x00\x00\x00")

func bindataTemplateAwsapilambdagolangMakefiletmplBytes() ([]byte, error) {
	return bindataRead(
		_bindataTemplateAwsapilambdagolangMakefiletmpl,
		"template/aws-api-lambda-golang/Makefile.tmpl",
	)
}



func bindataTemplateAwsapilambdagolangMakefiletmpl() (*asset, error) {
	bytes, err := bindataTemplateAwsapilambdagolangMakefiletmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "template/aws-api-lambda-golang/Makefile.tmpl",
		size: 174,
		md5checksum: "",
		mode: os.FileMode(493),
		modTime: time.Unix(1573661227, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataTemplateAwsapilambdagolangMaingotmpl = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8d\x31\x4f\xc3\x30\x10\x46\x67\xdf\xaf\x38\xdd\x64\x23\xab\xcd\x0a" +
	"\x1b\x4b\xd5\x81\x09\x06\x66\x2b\x39\x83\x85\x7d\x36\xf6\x85\x0c\xa8\xff\xbd\x6a\x94\xf5\x7d\x4f\xdf\x6b\x61\xfe" +
	"\x09\x5f\x8c\x25\x24\x01\x48\xa5\xd5\xae\x68\xc1\x50\x2c\x4a\x60\x48\x58\xcf\xdf\xaa\x8d\xc0\x01\xc4\x55\xe6\xdd" +
	"\xb4\x0e\xff\xc1\x3c\xf8\xe9\x1a\x64\xc9\x7c\x59\x65\xb6\x74\x26\x8f\x0f\xc7\x6e\xb8\x6f\xef\x3c\x5a\x95\xc1\x9f" +
	"\x3d\x29\x77\x8f\x1d\x9f\x0e\xfe\xbb\xf2\xd0\xfd\xc4\xc4\xa2\xa7\x4b\xeb\x49\x34\xda\xcd\x23\x5d\x39\xe7\xea\x71" +
	"\xab\x3d\x2f\xe4\xc0\xdc\xdc\x51\x7a\x4b\x43\x59\x5e\x65\xf9\xe0\xfe\xc7\x96\x5e\x9e\xa7\x69\x22\x8f\x92\xb2\x83" +
	"\x1b\xdc\x03\x00\x00\xff\xff\x1f\xaa\xd5\x95\xca\x00\x00\x00")

func bindataTemplateAwsapilambdagolangMaingotmplBytes() ([]byte, error) {
	return bindataRead(
		_bindataTemplateAwsapilambdagolangMaingotmpl,
		"template/aws-api-lambda-golang/main.go.tmpl",
	)
}



func bindataTemplateAwsapilambdagolangMaingotmpl() (*asset, error) {
	bytes, err := bindataTemplateAwsapilambdagolangMaingotmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "template/aws-api-lambda-golang/main.go.tmpl",
		size: 202,
		md5checksum: "",
		mode: os.FileMode(493),
		modTime: time.Unix(1573661076, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"template/aws-api-lambda-golang/.gitignore.tmpl": bindataTemplateAwsapilambdagolangGitignoretmpl,
	"template/aws-api-lambda-golang/Makefile.tmpl":   bindataTemplateAwsapilambdagolangMakefiletmpl,
	"template/aws-api-lambda-golang/main.go.tmpl":    bindataTemplateAwsapilambdagolangMaingotmpl,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"template": {Func: nil, Children: map[string]*bintree{
		"aws-api-lambda-golang": {Func: nil, Children: map[string]*bintree{
			".gitignore.tmpl": {Func: bindataTemplateAwsapilambdagolangGitignoretmpl, Children: map[string]*bintree{}},
			"Makefile.tmpl": {Func: bindataTemplateAwsapilambdagolangMakefiletmpl, Children: map[string]*bintree{}},
			"main.go.tmpl": {Func: bindataTemplateAwsapilambdagolangMaingotmpl, Children: map[string]*bintree{}},
		}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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