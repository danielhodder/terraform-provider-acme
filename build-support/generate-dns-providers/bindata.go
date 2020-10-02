// Code generated by go-bindata.
// sources:
// templates/acme-provider-sidebar-template.tmpl
// templates/dns-provider-doc-template.tmpl
// templates/dns-provider-go-template.tmpl
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

var _templatesAcmeProviderSidebarTemplateTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x54\xc1\x6e\xdb\x30\x0c\xbd\xe7\x2b\x38\x0f\x01\xb6\x83\xe2\x7b\x21\x1b\x18\xba\x1d\x37\x0c\xdb\x07\x18\xaa\x45\x27\x04\x64\xca\xa0\x64\xaf\x41\x90\x7f\x1f\x92\xda\x8e\xdd\xac\xab\xd3\x1e\x25\x52\x8f\xef\x3d\x92\xd2\x6b\xf8\x23\xa6\x29\x9c\xd9\xfb\x36\xc2\x1d\x31\xa3\x80\xf5\xb0\xce\x57\x00\x7a\x0d\xa5\xe7\x88\x1c\x8b\xca\x0b\xdc\x05\xb2\xf8\x60\x2e\x71\x00\xfd\x41\x29\x30\x6d\xf4\x6a\x8b\x8c\x62\x22\x5a\xa8\xc4\xd7\x60\xca\x1a\x55\x23\xbe\x23\x8b\xa2\xfa\x87\x2a\x62\xdd\x38\x13\x71\x13\xeb\xc6\x81\x52\x3d\x88\xa5\x0e\x4a\x67\x42\xc8\x12\xeb\xcb\x30\xa4\xc3\x8e\xac\x45\x56\x8d\x10\x47\x30\x55\x45\x8f\x2a\xfa\x26\x01\xf1\x0e\xb3\xa4\xf4\x75\xe3\xb0\x46\x8e\x46\xf6\xc9\x13\x14\x80\x6e\xdd\x80\xc5\xa6\x83\x11\x8f\x4d\x37\xe6\x00\x68\x47\x7a\x9d\x41\x5f\xa8\x28\x5b\x11\xe4\xf8\xe9\xa9\xfc\xce\xd7\x98\x7c\x86\x75\x7e\xc9\x07\xd0\x06\x76\x82\x55\x96\xa4\xa7\x9c\x74\x90\x16\x52\x62\x8b\x8f\x9b\x5d\xac\x5d\x92\x7f\x71\x0e\x7e\x0e\x11\x9d\x9a\x49\xc1\xd4\x51\xbe\x5a\x56\xff\xec\xdd\x19\x76\x39\x8b\xd3\x9b\x39\x95\xfb\xef\xdf\x46\x2e\xef\xa3\x22\x18\x7c\x2b\xe5\x7f\x3c\xf9\x98\xe4\xbf\xfa\xa4\xb9\xec\xab\x7e\xb0\xe9\x54\x47\x81\x1e\x1c\x26\xd3\xb4\x5b\x78\x28\xc1\x2d\x85\x28\x26\x92\xe7\x2b\x52\xaf\xdb\x24\xe9\x14\xa0\xf7\xeb\x14\x29\xa6\xf7\xcf\x74\x0c\xbe\xbd\x91\x72\x89\x12\xa9\xa2\xd2\xc4\x6b\x1b\x97\x30\x9e\xbc\x9f\x12\x9e\x5c\xbf\xca\x57\xa7\xad\x7b\x3e\x05\xcb\x95\x58\x0e\xe3\x3e\x87\x1b\xe7\xd2\x72\x28\xfe\xbd\x30\x5f\x7f\xfc\x7e\x61\x61\x6e\x98\x9c\xc3\x41\x81\x18\xde\x22\x6c\x8e\xc7\x1b\x1b\x34\x93\xa5\x0e\x87\xcd\xbd\xb7\x78\x3c\xbe\xa1\x47\x73\x91\x23\x52\x2f\x74\x3c\x2f\x18\xab\x93\x1e\x64\x3b\xd3\xf2\x72\xf3\x2e\x11\x9d\x5a\xea\xfa\x7f\x1b\xd9\x0e\x7f\x78\x06\x7b\x42\x77\x3e\x5e\x02\xab\xbf\x01\x00\x00\xff\xff\xa7\xdf\x84\x96\xfa\x05\x00\x00")

func templatesAcmeProviderSidebarTemplateTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesAcmeProviderSidebarTemplateTmpl,
		"templates/acme-provider-sidebar-template.tmpl",
	)
}

func templatesAcmeProviderSidebarTemplateTmpl() (*asset, error) {
	bytes, err := templatesAcmeProviderSidebarTemplateTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/acme-provider-sidebar-template.tmpl", size: 1530, mode: os.FileMode(420), modTime: time.Unix(1601656654, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDnsProviderDocTemplateTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x51\x6b\xe3\x46\x10\x7e\xdf\x5f\x31\x38\x81\xde\x99\x48\x7e\x37\xbd\x03\xe3\xe6\xe0\xe0\x9a\x86\x4b\xda\x17\x63\xac\xb1\x34\x92\x96\xac\x76\xc5\xec\xca\x39\xa3\xea\xbf\x97\x91\x63\xc9\x76\x4b\x92\x16\xfa\x94\x58\x3b\x3b\x33\xdf\xf7\xcd\x37\x1b\x45\x91\x32\xb8\x77\x4d\x98\xc3\x04\xd3\x8a\x26\xaa\xc6\x82\x36\x41\x07\x43\x73\x98\x2c\x96\xbf\xde\xce\xa1\x6d\xe3\x3b\xac\xa8\xeb\xe0\x97\xbb\x07\x58\x96\x68\x0c\xd9\x82\xe0\x9e\xdd\x4e\x67\xc4\x13\xe5\x75\x46\x5b\xe4\x4d\xda\x30\x93\x95\x64\x99\x4b\x7d\x24\x19\xa3\xcc\xfa\xa8\x7e\x89\xf4\x51\xdb\xc6\x4b\x97\x51\xd7\x4d\x54\x46\x3e\x65\x5d\x07\xed\xec\x1c\xfe\x8c\x14\x1c\x13\x7a\x40\x60\xf2\xae\xe1\x94\x20\x38\xa8\xd0\x62\x41\x90\x12\x07\x9d\xeb\x14\x03\x79\x70\x16\xd0\x82\xf4\x07\xcb\x45\xac\x04\xc8\xcf\x5b\xfe\xac\x54\xf4\x19\xa6\xd3\xbb\xdf\x1e\x6f\xe7\xd3\x29\x3c\x96\x04\xb9\x33\xc6\x3d\x6b\x5b\x40\xe6\xd2\xa6\x22\x1b\x50\x4a\x82\xf6\x80\x4d\x70\x51\x41\x96\x18\x03\x65\x90\xb3\xab\x20\x94\xa4\xfa\xb4\xc7\x9e\x7f\xf2\xb0\xb8\xff\x0a\x46\x6f\x19\x79\x0f\x2b\x43\x85\x5b\x7f\x28\x43\xa8\xfd\x7c\x36\x2b\x5c\x8f\x32\x2e\x74\x28\x9b\x6d\xac\xdd\x4c\xce\x67\x1f\x63\xf5\xe0\x2a\x02\x4f\xa9\x14\xf3\x50\xe1\x1e\x98\x72\x62\x01\x24\x21\x90\x69\xa6\x34\x98\x3d\x44\xa0\x2d\x54\xce\x07\x48\xd1\x93\xbf\x91\x16\x3c\xa9\xe1\x2a\xd6\xb5\xd9\xcb\xb5\x50\x12\x3c\x12\x33\xe6\x8e\xab\xa1\x3f\x40\x0f\xcf\x64\x4c\xac\xd4\xd5\x9b\x52\x29\x25\x94\x24\x83\x0a\x49\x1f\x97\x0e\x71\x43\xd2\x14\x2d\x6c\x09\x1a\x4f\x99\x94\xae\x89\xfb\xa2\x67\xd1\x1e\x72\xc7\x4a\xba\x5a\x25\x42\xc2\xe6\x44\xa1\x64\xbd\x3a\x4a\x78\x18\x83\x93\xb3\xf5\xa8\xee\xb3\x0e\xa5\x6a\x5b\x9d\x43\xfc\xfb\xf7\x6f\x5d\xb7\x1a\x00\xac\x3f\xb4\xed\xe1\xdb\xc7\xb6\x25\xe3\xa9\xeb\x86\xb3\xb6\x8d\x80\x6c\xd6\x75\xb1\x52\xaf\x94\x99\xc3\x4c\xc6\x70\x36\x4c\xdf\x4c\x42\x66\x3c\x3b\x09\x8a\xcb\x50\x19\xa5\xbe\x38\x86\xd4\x55\xb5\xa1\x40\xa0\xad\x60\x3d\x4c\x89\xb3\x50\xba\x67\xa1\xa0\xf1\x04\xa1\xd4\x7e\xe4\x48\x9a\xef\x45\x39\x43\x2f\xe0\xd5\xb1\xa9\x1b\xf0\x44\xb0\x2a\x89\xe9\x15\x42\x7a\x8f\x8c\xb4\xae\x5f\x85\x75\x19\xfc\x6e\x94\x57\x8d\xd7\xb6\xb8\xb8\xae\xd4\xd5\x15\xdc\xfe\x40\x81\xae\x54\x92\x24\x65\x6a\x86\xee\x0f\x3b\xe1\x54\xd6\x09\x4c\xce\x7e\xb5\x0a\x20\x8e\x63\xa5\x00\x32\xeb\x37\xe3\x20\xc9\x01\x8c\x54\x7d\x82\xc9\x89\xf3\x01\x3a\xd5\x49\x31\xa5\x44\x4a\x51\x7f\xe9\x6c\xae\x8b\x86\x7b\xd6\xe3\x7b\x26\x4f\x36\x74\x9d\x74\xb7\xe0\xa2\xb7\x2d\x7c\x17\xff\x90\x4d\xe9\x30\xc5\xa3\xb1\xf1\x25\xc2\x1f\xc7\x96\x74\x28\x89\xa1\x46\x2f\xf3\x8b\x1e\xc8\xee\x34\x3b\xdb\xa7\xd9\x21\x6b\xdc\x1a\x71\x9a\x63\x35\xd8\x30\x94\xec\x9a\xe2\x45\xd0\xb4\x6f\x27\x81\xad\x71\xe9\x93\xf8\x53\xd6\xc2\x2a\x39\xc3\xf8\xda\x88\x9f\xb3\x1c\x21\x17\xeb\xa1\xcb\x31\xdd\x7f\x33\x4d\x0c\x32\xad\x95\x63\x82\x8c\x02\x6a\xe3\xfb\x29\x53\xff\xe3\x94\xf5\x00\xde\x3f\x68\x67\x34\x29\xf5\xd5\x02\x66\x99\x16\x65\x6f\x2e\xb4\x42\xe3\x9d\x08\xe6\x83\x63\xca\x84\x19\x04\xe3\x52\x34\x90\x6b\x43\x37\xa3\xc5\x6a\x0c\xa5\xf2\x4d\x5d\x1b\x4d\x19\x6c\xf7\xd0\xff\xbf\x17\xf5\xe5\x78\xe0\x76\xf4\xe4\xe6\xcb\xd7\x6f\xb7\x09\xf8\x26\xcf\xf5\x8f\x18\x1e\x46\x86\xfe\x86\x55\x6a\x09\xc4\x88\x0e\x46\x58\xcb\x62\x3b\x30\x7c\xb2\x0b\x84\xaf\xb7\xaf\xfe\x5b\x3b\x1e\xa7\xb1\xcf\xe4\xa3\xdc\xf1\xf0\x5a\x46\x03\x57\xaa\x6d\x19\xc5\x56\xd7\x4f\x37\x70\xbd\x83\xf9\xa7\x4b\xc7\x2c\x99\x32\xb2\x41\xa3\xf1\x5d\xa7\xa6\xb2\xe2\xaf\x9f\x64\xbf\x47\xd0\xb6\xd7\x3b\x59\x95\xc3\xd6\x7c\x3b\xdb\xe2\x45\x2e\x34\xef\x49\x26\xfe\xbd\xb5\xbb\x3f\x90\x17\x46\xcb\x13\xd6\x75\x17\x06\x1d\x2c\x07\xc8\x04\xd3\xe9\xf0\x8a\x45\xbe\xa6\x54\x58\x99\x4e\x01\x0f\x97\x7b\xee\x7b\x4d\xb7\x6e\x47\x2a\x3d\x6d\x0c\x76\x68\x1a\xf2\x73\xf5\x4f\x10\x2e\x5b\x38\xeb\xbb\x4f\xde\xa7\x4e\x0e\x10\x92\x73\x0c\xc7\xbf\xe3\x97\x73\x12\xfe\x0a\x00\x00\xff\xff\x19\x3c\xe2\xf7\x2a\x09\x00\x00")

func templatesDnsProviderDocTemplateTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesDnsProviderDocTemplateTmpl,
		"templates/dns-provider-doc-template.tmpl",
	)
}

func templatesDnsProviderDocTemplateTmpl() (*asset, error) {
	bytes, err := templatesDnsProviderDocTemplateTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/dns-provider-doc-template.tmpl", size: 2346, mode: os.FileMode(420), modTime: time.Unix(1601656654, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDnsProviderGoTemplateTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x4f\x8b\xdb\x30\x10\xc5\xef\xfa\x14\xd3\x10\xa8\x03\x59\xf9\xbe\x90\x43\xd8\x64\x7b\x6a\x08\x2c\xe4\x52\x7a\x98\xb5\xc6\x8e\xb0\x3c\x32\x92\xec\x12\x8c\xbe\x7b\x91\x13\xbb\xe9\x9f\x94\x3d\xd9\x33\xd2\x7b\xf3\x7b\xf6\xe4\x39\x6c\xbb\x60\x9f\x2a\x62\x72\x18\x48\x41\xa9\x0d\x49\xd8\x59\x60\x1b\x80\x94\x0e\x52\xb4\x58\xd4\x58\x11\x60\xd1\x90\x10\xba\x69\xad\x0b\x90\x89\xc5\x30\x2c\xe5\xb1\xae\x8e\x18\xce\x31\xe6\xc5\x19\x8d\x21\xae\x68\x21\x86\xe1\x09\x1c\x72\x45\x20\x8f\xce\xf6\x5a\x91\xf3\x31\xfe\x29\x68\xa7\xa3\x5c\xb1\xcf\x87\x41\x7e\xb1\xc7\xba\x8a\xf1\xaa\x27\x56\x31\x8a\x95\x10\x79\x0e\x8a\xfd\xe4\xf3\x8a\x45\xb0\xee\xf2\xda\x71\x01\xda\x03\x42\xd9\x71\x11\xb4\x65\x08\x67\x0c\x50\xa0\x31\xa9\x3b\x79\x7f\xf6\x49\x5f\x58\xf6\xc1\x75\x49\x09\xc8\x0a\x1c\x85\xce\xb1\x87\x70\xa6\xf9\x26\x68\x0e\xe4\x4a\x2c\x48\x8a\x70\x69\xe9\xd1\xd0\x34\x2f\x5b\x41\x36\xc7\x9d\x23\xae\x81\x9c\xb3\xee\x11\xf2\x0d\xf7\x56\x94\x09\xc5\x18\xb0\xe5\x48\xd1\xa3\xd1\x0a\x76\x87\xb7\x99\x67\x04\xf7\x5d\x9b\x3e\x36\x29\x78\xbf\xc0\xf6\xe5\xeb\x7e\x3e\x96\xa2\x47\xf7\xaf\x29\x1b\x68\xb0\xfd\xe6\x83\xd3\x5c\x7d\x7f\x90\x61\xf8\xcf\x1f\x92\x2f\x56\x51\x8c\x8b\xe7\x0f\x24\xbd\x19\xe9\x12\xe4\x9e\xfb\x13\xba\xad\xd1\xe8\x29\x39\x35\xd8\xee\xb9\xd7\xce\x72\x43\x1c\x4e\xe8\x34\xbe\x1b\x3a\xa1\xe9\xc8\x67\x77\x84\xd7\xc7\x3d\xd0\xb2\x5e\xc3\xb2\x87\xe7\xcd\xdf\xa6\x69\x81\xea\x91\x2d\xbd\xf5\x31\x2e\xd6\x77\xbb\x12\x57\x77\x45\x3b\x32\x26\x97\x5f\x8b\x25\x0f\xf4\x63\x77\x78\x9b\x52\x64\x2b\xa1\xcb\xf1\xd6\xa7\x0d\xb0\x36\x30\x88\xeb\x62\xa4\x62\x94\x8b\x28\xa6\x56\xbb\x4e\x5d\x11\x7f\x1b\x28\x7e\x06\x00\x00\xff\xff\x8a\x29\xdb\x58\x3f\x03\x00\x00")

func templatesDnsProviderGoTemplateTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesDnsProviderGoTemplateTmpl,
		"templates/dns-provider-go-template.tmpl",
	)
}

func templatesDnsProviderGoTemplateTmpl() (*asset, error) {
	bytes, err := templatesDnsProviderGoTemplateTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/dns-provider-go-template.tmpl", size: 831, mode: os.FileMode(420), modTime: time.Unix(1601656654, 0)}
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
	"templates/acme-provider-sidebar-template.tmpl": templatesAcmeProviderSidebarTemplateTmpl,
	"templates/dns-provider-doc-template.tmpl": templatesDnsProviderDocTemplateTmpl,
	"templates/dns-provider-go-template.tmpl": templatesDnsProviderGoTemplateTmpl,
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
		"acme-provider-sidebar-template.tmpl": &bintree{templatesAcmeProviderSidebarTemplateTmpl, map[string]*bintree{}},
		"dns-provider-doc-template.tmpl": &bintree{templatesDnsProviderDocTemplateTmpl, map[string]*bintree{}},
		"dns-provider-go-template.tmpl": &bintree{templatesDnsProviderGoTemplateTmpl, map[string]*bintree{}},
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

