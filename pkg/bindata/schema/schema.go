// Copyright 2019 Publr Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-bindata.
// sources:
// data/schema/posts.sql
// data/schema/sites.sql
// data/schema/users.sql
// DO NOT EDIT!

package schema

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

var _dataSchemaPostsSQL = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x55\x4f\x73\xe2\xb8\x17\xbc\xf3\x29\xba\x72\x82\x2a\x20\x4c\xaa\x72\xf8\xfd\xe6\xe4\x01\xb1\xeb\x1a\x22\x67\x6d\xb3\x33\x39\x61\x61\x3f\x6c\xd5\xda\x92\x57\x92\x87\xf0\xed\xb7\x84\x49\x02\x99\x9a\xcc\xfe\xa9\x9d\xaa\xe5\xc0\x41\xaf\x5f\xab\x5f\xf7\xb3\x3d\x99\x60\xae\xdb\x83\x91\x65\xe5\x70\x33\x7b\xf7\x3f\xdc\x77\xdb\xda\x20\xe8\x5c\xa5\x8d\x9d\x0e\x26\x93\xc1\x64\x82\x95\xcc\x49\x59\x2a\xd0\xa9\x82\x0c\x5c\x45\x08\x5a\x91\x57\xf4\x54\x19\xe3\x57\x32\x56\x6a\x85\x9b\xe9\x0c\x43\x0f\xb8\x3a\x95\xae\x46\xef\x3d\xc5\x41\x77\x68\xc4\x01\x4a\x3b\x74\x96\xe0\x2a\x69\xb1\x93\x35\x81\x1e\x73\x6a\x1d\xa4\x42\xae\x9b\xb6\x96\x42\xe5\x84\xbd\x74\xd5\xf1\x9e\x13\x8b\x57\x82\x87\x13\x87\xde\x3a\x21\x15\x04\x72\xdd\x1e\xa0\x77\xe7\x40\x08\x77\x12\xed\x7f\x95\x73\xed\xff\xaf\xaf\xf7\xfb\xfd\x54\x1c\x05\x4f\xb5\x29\xaf\xeb\x1e\x6a\xaf\x57\xe1\x9c\xf1\x84\x4d\x6e\xa6\xb3\x53\xd3\x5a\xd5\x64\x2d\x0c\xfd\xde\x49\x43\x05\xb6\x07\x88\xb6\xad\x65\x2e\xb6\x35\xa1\x16\x7b\x68\x03\x51\x1a\xa2\x02\x4e\x7b\xd1\x7b\x23\x9d\x54\xe5\x18\x56\xef\xdc\x5e\x18\xf2\x34\x85\xb4\xce\xc8\x6d\xe7\x2e\x3c\x7b\x92\x28\xed\x05\x40\x2b\x08\x85\xab\x20\x41\x98\x5c\xe1\x43\x90\x84\xc9\xd8\x93\x7c\x0a\xd3\x9f\xa3\x75\x8a\x4f\x41\x1c\x07\x3c\x0d\x59\x82\x28\xc6\x3c\xe2\x8b\x30\x0d\x23\x9e\x20\x5a\x22\xe0\x0f\xf8\x18\xf2\xc5\x18\x24\x5d\x45\x06\xf4\xd8\x1a\x3f\x81\x36\x90\xde\x4d\x2a\x8e\xd6\x25\x44\x17\x12\x76\xba\x97\x64\x5b\xca\xe5\x4e\xe6\xa8\x85\x2a\x3b\x51\x12\x4a\xfd\x85\x8c\x92\xaa\x44\x4b\xa6\x91\xd6\xa7\x6a\x21\x54\xe1\x69\x6a\xd9\x48\x27\xdc\xf1\xe8\xab\xb9\xa6\x83\xc1\x3c\x66\x41\xca\x90\x06\x1f\x56\x0c\xe1\x12\x3c\x4a\xc1\x3e\x87\x49\x9a\x20\x6b\xb5\x75\x36\xc3\x70\xe0\x83\xc9\x64\x91\x41\x2a\x47\x25\x19\x04\xeb\x34\xda\x84\x7c\x1e\xb3\x3b\xc6\xd3\x63\x13\x5f\xaf\x56\xe3\x1e\xe9\xa4\xab\x29\xc3\x17\x61\xf2\x4a\x98\xe1\xed\x6c\xf4\x1a\x61\xeb\xae\x7c\x13\x50\xb9\xa6\xce\xe0\xe8\xd1\xbd\xae\xc8\x46\x94\x67\xe4\x37\xb7\xb7\xa3\x53\xa5\xed\xb6\xb5\xb4\x15\x15\x19\xb6\x5a\xd7\x24\xd4\x73\x2f\x16\x6c\x19\xac\x57\x29\x66\x27\x6c\x6e\x48\x38\x72\xb2\xa1\x0c\xfe\xdf\x3a\xd1\xb4\x97\xd0\xf9\x3a\x8e\x19\x4f\x37\x69\x78\xc7\x92\x34\xb8\xbb\xbf\xbc\xe6\xcd\xde\x33\xbd\x5d\x5b\xfc\xe5\x9b\x10\x71\xac\xef\x17\x3e\x98\x6f\xa8\xb8\x8f\xc3\xbb\x20\x7e\xc0\x47\xf6\x80\xa1\x8f\xe6\x64\x42\xc8\x17\xec\xb3\xcf\xea\x71\xd3\x7b\x3c\xec\xbd\x1e\x0d\x46\xef\xbf\x1f\xf6\xc6\x4a\x47\x7f\x27\xf1\xbe\xf9\x7b\xa1\x7a\xf6\x4d\xa1\x1b\x21\xd5\x0b\xee\xdd\xec\x6b\xe0\x3f\x48\xe7\xc7\xda\xbd\xe6\xe1\x2f\x6b\x76\x3c\xce\xba\xdf\x36\x17\x1e\x9e\x99\x32\xbe\x9c\xfd\xd4\x3c\x8f\x78\x92\xc6\x41\xc8\x53\x64\xf9\x45\xef\x32\x8a\x59\xf8\x13\xef\x79\x77\x3d\xef\x2b\xca\x11\x62\xb6\x64\x31\xe3\x73\xf6\xfc\xa4\x3e\x65\x7d\x3e\x4f\x90\xcc\x83\x05\xf3\x27\x0b\xb6\x62\x2f\x27\x7f\x6e\x1f\x44\xff\x55\xf9\xd7\x36\xa2\xe7\xdf\x74\x96\x8c\x12\xcd\xdb\xef\x8c\xff\xf4\x52\xbc\x18\x79\xb9\x16\xaf\x0d\xf8\xf6\x6a\x3c\x33\xfc\x98\xe5\xf8\x23\x00\x00\xff\xff\x24\x79\x17\xf9\x68\x08\x00\x00")

func dataSchemaPostsSQLBytes() ([]byte, error) {
	return bindataRead(
		_dataSchemaPostsSQL,
		"data/schema/posts.sql",
	)
}

func dataSchemaPostsSQL() (*asset, error) {
	bytes, err := dataSchemaPostsSQLBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/schema/posts.sql", size: 2152, mode: os.FileMode(420), modTime: time.Unix(1557587009, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSchemaSitesSQL = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\x4f\x6f\xdb\x38\x14\xc4\xef\xfe\x14\x03\x9f\x6c\xc0\xff\x12\x60\x0f\xbb\x39\x29\xb6\x82\x15\x62\x4b\x5e\x89\xda\xd4\xa7\x98\x96\x9e\xa5\x87\xca\xa4\x4a\x52\x71\xfc\xed\x0b\xfa\x0f\xda\x34\xe8\xa1\x3a\xe8\xf0\xf8\x7b\xc3\x99\xe1\x78\x8c\xb9\x6e\x4f\x86\xab\xda\xe1\x7e\x76\xf7\x37\xd6\xdd\xae\x31\x08\x3a\x57\x6b\x63\x27\xbd\xf1\xb8\x37\x1e\x63\xc9\x05\x29\x4b\x25\x3a\x55\x92\x81\xab\x09\x41\x2b\x8b\x9a\x6e\x27\x23\xfc\x4f\xc6\xb2\x56\xb8\x9f\xcc\x30\xf0\x40\xff\x7a\xd4\x1f\x3e\x78\x89\x93\xee\x70\x90\x27\x28\xed\xd0\x59\x82\xab\xd9\x62\xcf\x0d\x81\xde\x0b\x6a\x1d\x58\xa1\xd0\x87\xb6\x61\xa9\x0a\xc2\x91\x5d\x7d\xbe\xe7\xaa\xe2\x9d\x60\x73\xd5\xd0\x3b\x27\x59\x41\xa2\xd0\xed\x09\x7a\xff\x33\x08\xe9\xae\xa6\xfd\x57\x3b\xd7\xfe\x33\x9d\x1e\x8f\xc7\x89\x3c\x1b\x9e\x68\x53\x4d\x9b\x0b\x6a\xa7\xcb\x68\x1e\xc6\x59\x38\xbe\x9f\xcc\xae\x4b\xb9\x6a\xc8\x5a\x18\xfa\xd6\xb1\xa1\x12\xbb\x13\x64\xdb\x36\x5c\xc8\x5d\x43\x68\xe4\x11\xda\x40\x56\x86\xa8\x84\xd3\xde\xf4\xd1\xb0\x63\x55\x8d\x60\xf5\xde\x1d\xa5\x21\x2f\x53\xb2\x75\x86\x77\x9d\xfb\xd0\xd9\xcd\x22\xdb\x0f\x80\x56\x90\x0a\xfd\x20\x43\x94\xf5\xf1\x18\x64\x51\x36\xf2\x22\x2f\x91\xf8\x37\xc9\x05\x5e\x82\x34\x0d\x62\x11\x85\x19\x92\x14\xf3\x24\x5e\x44\x22\x4a\xe2\x0c\xc9\x13\x82\x78\x83\xe7\x28\x5e\x8c\x40\xec\x6a\x32\xa0\xf7\xd6\xf8\x04\xda\x80\x7d\x9b\x54\x9e\xab\xcb\x88\x3e\x58\xd8\xeb\x8b\x25\xdb\x52\xc1\x7b\x2e\xd0\x48\x55\x75\xb2\x22\x54\xfa\x8d\x8c\x62\x55\xa1\x25\x73\x60\xeb\x5f\xd5\x42\xaa\xd2\xcb\x34\x7c\x60\x27\xdd\x79\xf4\x29\xd7\xa4\xd7\x9b\xa7\x61\x20\x42\x88\xe0\x71\x19\x22\x7a\x42\x9c\x08\x84\x5f\xa2\x4c\x64\xd8\x5a\x76\x64\xb7\x18\xf4\xfc\xc3\x6c\xb9\xdc\x82\x95\xa3\x8a\x0c\x82\x5c\x24\xaf\x51\x3c\x4f\xc3\x55\x18\x8b\xf3\x52\x9c\x2f\x97\xa3\x0b\xe9\xd8\x35\xb4\xc5\x9b\x34\x45\x2d\xcd\xe0\xaf\xd9\xf0\x57\xa2\xd4\x07\xc9\xea\x07\x72\x37\xfb\xcc\x14\x86\xa4\x23\xc7\x07\xda\xc2\xff\xad\x93\x87\xf6\x0c\x60\x11\x3e\x05\xf9\x52\x60\x9e\xa7\x69\x18\x8b\x57\x11\xad\xc2\x4c\x04\xab\xf5\x75\xb5\x6b\xcb\x3f\x5e\x45\x12\x23\x5f\x2f\x7c\x19\xbf\x91\x5d\xa7\xd1\x2a\x48\x37\x78\x0e\x37\x18\xf8\x3a\x86\x97\x79\x1e\x47\xff\xe5\xe1\x79\xbc\xed\xbe\xbe\xde\xb2\x0d\x6e\x29\x87\xbd\xe1\xc3\xf7\x00\x00\x00\xff\xff\x41\x8c\x7e\x58\xb7\x03\x00\x00")

func dataSchemaSitesSQLBytes() ([]byte, error) {
	return bindataRead(
		_dataSchemaSitesSQL,
		"data/schema/sites.sql",
	)
}

func dataSchemaSitesSQL() (*asset, error) {
	bytes, err := dataSchemaSitesSQLBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/schema/sites.sql", size: 951, mode: os.FileMode(420), modTime: time.Unix(1557587009, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataSchemaUsersSQL = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x53\x4d\x93\xe2\x36\x14\xbc\xf3\x2b\xba\xe6\x34\x54\x01\xc3\x4e\x55\x0e\xc9\x9c\xbc\x20\x12\xd7\x32\xf2\xc4\x96\xb3\x33\x27\x10\xf6\xc3\x56\xad\x91\x1c\x49\x5e\x86\x7f\x9f\x92\x81\x30\x6c\xb2\x9b\x8f\xdb\xfa\xe0\xc3\x53\x77\xab\x5f\xb7\x3d\x1e\x63\x66\xda\x83\x55\x55\xed\x71\x3f\x7d\xf7\x23\x9e\xba\x4d\x63\x11\x75\xbe\x36\xd6\x4d\x06\xe3\xf1\x60\x3c\xc6\x52\x15\xa4\x1d\x95\xe8\x74\x49\x16\xbe\x26\x44\xad\x2c\x6a\x3a\x9f\x8c\xf0\x1b\x59\xa7\x8c\xc6\xfd\x64\x8a\xdb\x00\xb8\x39\x1d\xdd\x0c\x1f\x82\xc4\xc1\x74\xd8\xc9\x03\xb4\xf1\xe8\x1c\xc1\xd7\xca\x61\xab\x1a\x02\xbd\x16\xd4\x7a\x28\x8d\xc2\xec\xda\x46\x49\x5d\x10\xf6\xca\xd7\xfd\x3d\x27\x95\xe0\x04\x2f\x27\x0d\xb3\xf1\x52\x69\x48\x14\xa6\x3d\xc0\x6c\xdf\x02\x21\xfd\xc9\x74\x78\x6a\xef\xdb\x9f\xee\xee\xf6\xfb\xfd\x44\xf6\x86\x27\xc6\x56\x77\xcd\x11\xea\xee\x96\xf1\x8c\xf1\x8c\x8d\xef\x27\xd3\x13\x29\xd7\x0d\x39\x07\x4b\xbf\x77\xca\x52\x89\xcd\x01\xb2\x6d\x1b\x55\xc8\x4d\x43\x68\xe4\x1e\xc6\x42\x56\x96\xa8\x84\x37\xc1\xf4\xde\x2a\xaf\x74\x35\x82\x33\x5b\xbf\x97\x96\x82\x4c\xa9\x9c\xb7\x6a\xd3\xf9\xab\xcc\xce\x16\x95\xbb\x02\x18\x0d\xa9\x71\x13\x65\x88\xb3\x1b\xbc\x8f\xb2\x38\x1b\x05\x91\x8f\xb1\xf8\x25\xc9\x05\x3e\x46\x69\x1a\x71\x11\xb3\x0c\x49\x8a\x59\xc2\xe7\xb1\x88\x13\x9e\x21\x59\x20\xe2\x2f\xf8\x10\xf3\xf9\x08\xa4\x7c\x4d\x16\xf4\xda\xda\xb0\x81\xb1\x50\x21\x4d\x2a\xfb\xe8\x32\xa2\x2b\x0b\x5b\x73\xb4\xe4\x5a\x2a\xd4\x56\x15\x68\xa4\xae\x3a\x59\x11\x2a\xf3\x99\xac\x56\xba\x42\x4b\x76\xa7\x5c\x68\xd5\x41\xea\x32\xc8\x34\x6a\xa7\xbc\xf4\xfd\xe8\x2f\x7b\x4d\x06\x83\x59\xca\x22\xc1\x20\xa2\xf7\x4b\x86\x78\x01\x9e\x08\xb0\xe7\x38\x13\x19\xd6\x9d\x23\xeb\xd6\xb8\x1d\x84\x62\xd6\xaa\x5c\x43\x69\x4f\x15\x59\x44\xb9\x48\x56\x31\x9f\xa5\xec\x91\x71\xd1\x93\x78\xbe\x5c\x8e\x8e\x48\xda\x49\xd5\xac\xf1\x59\xda\xa2\x96\xf6\xf6\x87\xe9\xf0\x4b\x44\x50\xd6\x72\x47\xdf\x04\x6d\xbb\xa6\xf9\x47\x50\x61\x49\x7a\xf2\x2a\xc0\xc2\xdb\x79\xb9\x6b\x7b\x00\xe6\x6c\x11\xe5\x4b\x81\x59\x9e\xa6\x8c\x8b\x95\x88\x1f\x59\x26\xa2\xc7\xa7\xb3\x89\xb6\xfc\xcf\x54\x24\x1c\xf9\xd3\x3c\x44\xf6\x15\xd9\xa7\x34\x7e\x8c\xd2\x17\x7c\x60\x2f\xb8\x0d\xa1\x0d\x8f\xf3\x98\xcf\xd9\x73\x48\xf1\x75\x75\xd9\xfe\xf6\x92\xc4\x09\xb6\xc8\x97\x4b\xc1\x9e\x05\xd6\x5b\xbf\x3a\x17\x70\x41\x8d\xde\xc4\x32\x1c\x0c\x1f\xbe\x5d\xa0\x53\x9e\x56\xff\xb7\xc5\x9e\x5c\x9a\x9d\x54\xfa\xd2\xc0\xbb\xe9\xdf\x97\xb9\xfa\x57\x8d\x5a\xd3\x50\x7f\xfd\x77\xdb\x62\xce\xe3\x5f\x73\xd6\x8f\xd7\xdd\xa7\xd5\x55\xc0\x5f\x04\x31\xba\x8e\xf0\x24\x30\x4b\x78\x26\xd2\x28\xe6\x02\xeb\xe2\x4c\x5d\x24\x29\x8b\x7f\xe6\x47\xd9\xed\xa7\xaf\x29\x0e\x91\xb2\x05\x4b\x19\x9f\xb1\x3f\xff\xce\xb7\x5f\xd0\xdb\xb5\xa2\x6c\x16\xcd\x59\x98\xcc\xd9\x92\x5d\x26\x83\xe1\xc3\x1f\x01\x00\x00\xff\xff\xd9\xa8\x2a\x01\x43\x06\x00\x00")

func dataSchemaUsersSQLBytes() ([]byte, error) {
	return bindataRead(
		_dataSchemaUsersSQL,
		"data/schema/users.sql",
	)
}

func dataSchemaUsersSQL() (*asset, error) {
	bytes, err := dataSchemaUsersSQLBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/schema/users.sql", size: 1603, mode: os.FileMode(420), modTime: time.Unix(1557587009, 0)}
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
	"data/schema/posts.sql": dataSchemaPostsSQL,
	"data/schema/sites.sql": dataSchemaSitesSQL,
	"data/schema/users.sql": dataSchemaUsersSQL,
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
	"data": {nil, map[string]*bintree{
		"schema": {nil, map[string]*bintree{
			"posts.sql": {dataSchemaPostsSQL, map[string]*bintree{}},
			"sites.sql": {dataSchemaSitesSQL, map[string]*bintree{}},
			"users.sql": {dataSchemaUsersSQL, map[string]*bintree{}},
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
