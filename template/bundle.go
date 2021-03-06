// Code generated by go-bindata. DO NOT EDIT.
//  memcopy: true
//  compress: true
//  decompress: once
//  metadata: true
//  asset-dir: true
//  restore: true
// sources:
//  template/minimal/index.html.tmpl
//  template/minimal/main.css.tmpl
//  template/minimal/main.js.tmpl

package template

import (
	"bytes"
	"compress/flate"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/tmthrgd/go-bindata/restore"
)

type asset struct {
	name string
	data string
	size int64
	mode os.FileMode
	time time.Time

	once  sync.Once
	bytes []byte
	err   error
}

func (a *asset) Name() string {
	return a.name
}

func (a *asset) Size() int64 {
	return a.size
}

func (a *asset) Mode() os.FileMode {
	return a.mode
}

func (a *asset) ModTime() time.Time {
	return a.time
}

func (*asset) IsDir() bool {
	return false
}

func (*asset) Sys() interface{} {
	return nil
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]*asset{
	"template/minimal/index.html.tmpl": &asset{
		name: "index.html.tmpl",
		data: "" +
			"\xec\x56\xcd\x8e\xdb\x36\x10\xbe\x17\xe8\x3b\xb0\x6c\x8e\x91\xb8\xce\xa1\x68\x03\xca\xc0\x76\x93" +
			"\x00\x41\x5b\xb4\x87\x4d\xd1\x1e\xc7\xe4\x58\x1a\x2f\x45\xaa\x24\xe5\x5d\x43\xf0\xbb\x17\xb4\x69" +
			"\x5b\xb2\x1d\x34\x45\xf6\xd0\x02\x39\x89\xe4\xfc\x7c\xf3\xf7\x91\x92\xdf\xbc\xf9\xf5\xee\xfe\xcf" +
			"\xdf\xde\xb2\x26\xb6\x66\xfe\xf5\x57\x32\x7f\x19\x93\x0d\x82\xde\xad\x18\x93\x2d\x46\x60\xaa\x01" +
			"\x1f\x30\x56\xfc\xc3\xfd\xbb\xe2\x7b\x3e\x91\x35\x31\x76\x05\xfe\xd5\xd3\xba\xe2\x7f\x14\x1f\x6e" +
			"\x8b\x3b\xd7\x76\x10\x69\x61\x90\x33\xe5\x6c\x44\x1b\x2b\x4e\x58\xa1\xae\x71\x6a\x6a\xa1\xc5\x8a" +
			"\xaf\x09\x1f\x3b\xe7\xe3\x48\xfb\x91\x74\x6c\x2a\x8d\x6b\x52\x58\xec\x36\x2f\x19\x59\x8a\x04\xa6" +
			"\x08\x0a\x0c\x56\xb3\xf2\xe6\x9a\x2f\xe8\x63\xe3\xfc\xc8\xd3\x30\x94\xbf\x60\x84\xf2\x76\x27\xd8" +
			"\x6e\xaf\x19\x69\x0c\xca\x53\x17\xc9\xd9\x2b\x96\x6f\x4e\xd2\x91\x79\xa4\x68\x70\x7e\xd0\xb9\x4f" +
			"\xbb\xed\x56\x8a\xc3\x31\x2d\x99\x45\xb6\x17\xfe\xe8\xc1\x6a\xb2\x75\xf9\x5e\x39\xcb\x38\xdf\x6e" +
			"\xb3\x0f\x43\xf6\x81\x79\x34\x15\x27\x95\xa0\xe3\xa6\xc3\x8a\x53\x0b\x35\x8a\xce\xd6\x9c\x35\x1e" +
			"\x97\x15\x0f\x11\x22\x29\x41\x6d\x2d\x92\x5e\x99\x44\xf3\x61\x40\xab\xaf\x78\x0a\x71\x63\x30\x34" +
			"\x88\xf1\x60\x9e\xfa\x13\x5e\x0b\xa1\xb4\x5d\x85\x52\x19\xd7\xeb\xa5\x01\x8f\xa5\x72\xad\x80\x15" +
			"\x3c\x09\x43\x8b\x20\x1a\xaa\x1b\x43\x75\x13\xcb\x55\x10\x3f\x94\xb3\x57\xe5\x8d\xd8\x3b\x13\x10" +
			"\x5d\x5b\x38\x8b\x85\x06\xff\x50\xb6\x64\x4b\x15\xc2\xb1\x12\xcf\x80\xbd\xe8\x4d\x0b\xe2\xa6\xfc" +
			"\xae\x7c\x25\x54\xc8\xfb\x7f\x05\x94\x6b\x94\x8c\x5b\x38\xb3\xdb\xb7\x8f\x05\xaf\x3e\xab\x18\xa7" +
			"\xb3\x14\xd8\x2a\xf0\xb9\x14\x7b\xd7\xcf\x0c\x64\xc0\xd6\x3d\xd4\x18\xc4\x06\x5a\xf3\x49\x68\x39" +
			"\xfd\x55\xce\xfe\x42\x5d\x8a\x23\xa5\xe5\xc2\xe9\x4d\xf6\xc1\x98\xb4\xb0\x66\xa4\x2b\xde\xba\x05" +
			"\x19\x2c\x2c\xac\x39\x53\x06\x42\xa8\x38\x85\xa2\x21\xad\xd1\x16\x11\x16\x06\x23\x3f\x5a\x31\x26" +
			"\x35\xed\xed\x2c\xac\x8b\x45\xef\x6b\xf4\x63\x71\x0a\xaf\x03\x9b\x82\x48\x9f\xcf\x11\x48\xa1\x69" +
			"\x7d\x8e\x9c\x23\x4c\xe0\x99\xaf\x67\xe8\x23\xa5\xa4\x00\x64\xd1\x33\x0a\xc5\xd2\xf4\xa4\xa7\xba" +
			"\x8c\x49\x6a\xeb\x83\xb6\x71\xb5\xe3\x0c\x4c\xac\xf8\xcf\xbb\xe5\xb8\xbe\x89\x82\x49\x61\x4f\xc1" +
			"\x09\xe0\x59\x94\x8c\x0d\x83\x07\x5b\x23\x7b\x41\x56\xe3\xd3\x4b\xf6\x02\x0d\xb6\x68\x23\x7b\x5d" +
			"\xb1\xf2\xd6\x47\x52\x06\xc3\x81\xc0\x97\x51\x6b\xef\x3a\xed\x1e\x6d\x71\x0c\xff\x22\xea\x6e\x5c" +
			"\x86\xd4\x5f\xf4\x9c\x2d\x9d\xaf\xb8\x6a\x50\x3d\x2c\xdc\x53\x31\x0c\xa4\xdf\x79\xd7\x1e\xd1\xcb" +
			"\x9f\x70\x93\x6e\xb1\x61\x98\x9e\x48\xd1\x9d\xbb\xef\xcd\xb4\xcc\xc6\x40\x17\x30\x99\xe6\xc4\x02" +
			"\xaa\x74\x29\xbe\xcf\xf9\xe5\xed\xdb\x53\x9a\x47\x88\xdf\xc1\xf4\x38\xcd\x35\x33\x7a\x2e\x21\x33" +
			"\xf8\xdb\x53\xac\x53\x4f\xa3\x90\xaf\x09\xa4\x80\xb9\x14\x86\xa6\x37\xe2\xa9\x2f\xbd\xb9\xd2\xa8" +
			"\x0b\xdd\x69\xff\xa4\xb0\x30\xda\x41\x20\x8d\x87\x62\xa4\xf5\x02\x76\xc3\x94\xd9\xb1\xa7\x0e\xff" +
			"\xc8\x8c\x66\xfd\xff\xe8\x9c\x7e\x99\xd2\xff\xf3\x94\xee\x06\xf3\xb4\xcf\xc0\x87\x82\xa4\x8b\x9e" +
			"\xe5\xb3\x51\x41\x3e\xad\xd3\xb2\x99\x1d\xfc\xc0\x5e\x78\x68\xde\x95\xae\x34\xb3\x67\xa9\xb7\xcc" +
			"\x48\xbb\x87\xe5\x9f\x0a\xfd\xb1\x91\xcc\x3e\x2e\xe6\xf0\x94\x4f\xaa\xcb\x38\x99\xeb\xfd\x6a\x66" +
			"\xe7\x1e\xce\xb8\x7a\x41\xe6\xcc\x27\x04\x7d\xb7\x17\x5f\x84\x7e\x75\xc2\x2e\x9f\x8e\x8b\x27\x4f" +
			"\xe4\xa4\x8e\x67\x79\x32\xf2\x47\x8a\x8c\x93\xdf\xfa\xfc\xc2\x4b\xb1\xfb\xa3\xff\x3b\x00\x00\xff" +
			"\xff",
		size: 3049,
		mode: 0666,
		time: time.Unix(1529704804, 441809400),
	},
	"template/minimal/main.css.tmpl": &asset{
		name: "main.css.tmpl",
		data: "" +
			"\xac\x57\xdf\x6f\xdb\x36\x10\x7e\x0f\x90\xff\x81\x6b\x11\x34\x06\x4a\x59\x92\xed\x34\x55\x30\xac" +
			"\xd8\xf3\x30\x0c\xeb\x5e\xf6\x48\x49\x67\x8b\x0b\x45\x12\x24\xe5\x1f\x0d\xf2\xbf\x0f\x22\xf5\x83" +
			"\xa2\xe5\x24\x5d\x07\x18\x81\xc4\xf0\x8e\xf7\x7d\xbc\xfb\xee\xf4\x85\xd6\x52\x28\x83\x1a\xc5\x6e" +
			"\x3f\x54\xc6\x48\x9d\x2d\x97\x5b\xc1\x8d\x8e\x76\x42\xec\x18\x10\x49\x75\x54\x88\x7a\x59\x68\xfd" +
			"\xcb\x96\xd4\x94\x9d\x7e\xfe\xb3\xc9\xe9\xe3\x87\xc5\xc3\xf5\xd5\xac\x79\x51\xf2\xe8\x1f\x5d\x02" +
			"\xa3\x7b\x15\x71\x30\x4b\x2e\xeb\xa5\x39\x49\xd8\x92\x02\x30\x03\xb2\x6b\x00\x6b\x49\x94\x21\xfc" +
			"\x4b\x1c\xc5\xd1\x7a\xbd\xa4\xbc\x84\x63\x54\x53\x1e\x15\x5a\x5b\xd7\xd7\x57\x95\xa9\xd9\x47\x94" +
			"\x8b\xf2\x84\x9e\xae\xaf\x10\xaa\x89\xda\x51\x9e\xa1\xf8\xa1\x7d\x93\xa4\x2c\x29\xdf\xf5\xaf\x15" +
			"\xd0\x5d\x65\x32\x94\xc4\xf1\x8d\x5d\x68\x41\x60\x4d\xbf\x81\xb7\x96\x93\xe2\x71\xa7\x44\xc3\x4b" +
			"\x5c\x08\x26\x54\x86\xde\x3f\x3d\x45\xbf\x0e\xab\xcf\xcf\xa3\xa9\xc3\x9a\x21\x0b\xf6\x23\xd2\x84" +
			"\x6b\xac\x41\xd1\xed\xc3\xf5\xd5\x73\x1b\x5e\xc4\xc4\x4e\x04\x91\x21\xd2\x18\x81\xd6\x1b\x79\xb4" +
			"\x4f\xd6\x5b\x49\xb5\x64\xe4\x94\xa1\x9c\x89\xe2\xf1\xc1\xed\x3f\xe2\x03\x2d\x4d\x65\x63\x93\xc7" +
			"\x61\xb1\x47\xb1\x71\x8b\xee\x1c\x4d\x4b\xc8\x89\x72\x47\x9d\xe1\xec\xfc\xa4\xf7\xbd\x9f\x8e\x18" +
			"\x6c\x84\xcc\x50\x1a\x6d\x14\xd4\x6e\x5d\x68\x6a\xa8\xe0\x19\xda\xd2\x23\x94\x97\x62\xb3\x66\x8e" +
			"\x53\x06\x5b\xd3\x3f\xcf\x53\xf7\x3b\xd9\x77\x9c\x89\x3d\xa8\x2d\x13\x07\x7c\xca\x46\xe8\xb9\x38" +
			"\x62\x5d\x91\x52\x1c\x5a\x6e\x62\x94\xb6\xc4\xe0\xb5\x3c\x22\xb5\xcb\xc9\x6d\xfc\x11\x75\xbf\x28" +
			"\xfe\xbc\x08\xf1\xe2\x42\x70\x03\xdc\xf8\x14\x3b\x50\x6b\x07\xc9\xed\xe6\x64\x8f\x2b\x20\x25\x74" +
			"\x04\x19\x38\x1a\x6c\x14\xe1\x7a\x2b\x54\x9d\xa1\x46\x4a\x50\x05\xd1\x30\x5c\x6d\x86\xa2\x3b\x05" +
			"\x35\xfa\xcd\x66\x22\xfa\xea\x32\x31\xb8\x62\x84\x3c\x98\x7f\x28\x5a\x13\x75\xea\xa0\x32\x30\x06" +
			"\x54\x9b\xc1\x85\x4d\xc0\x28\xd9\x74\x14\xdb\x63\x3b\x92\xa7\xbb\x50\x94\xea\x99\x90\xb3\xaa\xe5" +
			"\xcd\x05\x7e\xe6\x36\x0d\x50\x16\x82\x31\x22\x35\x4c\x13\x61\x20\x7b\x28\x88\x24\x96\x47\x14\xa3" +
			"\x95\x97\x42\x13\x73\x46\x9d\x87\xc1\x60\x6d\xf7\xb7\x7f\x93\xcd\x25\x23\xe2\x6c\x3c\x52\xbe\x42" +
			"\x21\x78\x39\xd2\xd2\x53\xbb\x79\x95\xdb\x39\xf7\x3e\x13\xa2\xa5\xc0\x9c\x32\x14\x47\x9f\x86\xed" +
			"\xa5\x12\xb2\x14\x07\x6e\xd3\x82\x50\xde\x6f\x9e\x49\xe1\x01\xd8\x67\x0b\x2c\x46\x6b\xaf\x9c\xde" +
			"\xd7\x22\xa7\x0c\x30\x27\x7b\xe7\xe0\x1b\xb6\xe2\x93\xa1\xd5\xc5\x32\xb9\x50\x13\x63\x09\xdf\x04" +
			"\xa8\xbc\xc4\x9d\x94\x63\xb2\xb9\x99\xd6\x1d\x17\x1c\x26\xb6\x42\x02\x47\xe7\x5e\xce\x50\x06\x16" +
			"\x4f\x6f\xa9\xd1\x40\x3b\x1c\x1d\xad\x8b\xbc\x51\xbb\xa1\x7e\x5c\x9c\x9d\x98\x38\xb8\xc3\xeb\x1c" +
			"\x39\xbd\x00\xf5\xfa\x33\x9e\xd2\x2d\x04\x04\x17\x8d\xd2\x6d\x6c\x52\x50\x6e\x40\xd9\x35\x7c\x80" +
			"\xfc\x91\x76\x75\xdb\x9d\x10\x47\x1b\x8d\x80\x68\xc0\x94\x63\xd1\x98\xb3\x0a\x9b\xd9\x70\x06\x49" +
			"\x4b\xc2\x2f\xe6\x49\xa7\x29\xb9\x30\x46\xd4\x59\x57\x30\x3e\x48\x92\x6b\xc1\x1a\x03\x13\x5c\x69" +
			"\xb7\x6d\x7a\xfb\x3e\xfb\x73\xa2\x91\x0b\x55\x82\xc2\x8a\x94\xb4\xd1\x2d\x5b\xce\x68\x48\xf5\xe4" +
			"\x2c\xbb\xe6\x39\x49\x5f\x25\x25\x7d\x1b\x2b\x19\x37\x15\x2e\x2a\xca\xca\xdb\x64\xe1\x5d\x7d\xfc" +
			"\x06\x83\xd4\x37\xb8\xf3\x6a\xeb\xb2\xc9\xca\x37\x49\xd2\x40\x66\x6c\x0e\xbf\x21\xc2\x09\x27\x4e" +
			"\xdf\x95\x30\xc4\xc0\x6d\xb2\xda\x94\xb0\x5b\x8c\x7c\x5c\xfe\x6f\x10\xf4\x9b\x02\xe8\x11\x8f\xd2" +
			"\xe4\xdd\x17\x5e\xc5\xdf\xe9\x6e\xf5\x2a\x1e\xfc\x32\x20\xfc\x12\x22\xa2\x0c\x2d\x18\xcc\xe8\x4f" +
			"\xea\x77\xce\x5e\x60\xa4\xea\xb6\x06\x39\x1a\xdb\x36\x19\x76\x7f\x37\x52\x15\xad\x5a\x07\x6d\xa7" +
			"\xdb\xdb\x0f\x43\x89\xeb\xb2\x71\xff\xf0\x93\x9b\x15\x09\x37\x93\xf4\x0e\x67\x84\xd5\xbd\x15\xed" +
			"\x70\x40\x48\x17\xb3\x13\xc5\xcb\xbb\xfd\x19\xd0\x47\x3e\xce\x94\x67\xe3\xd4\x38\x56\xa2\xbb\x1b" +
			"\xf7\x9b\xa0\x4a\x1d\xa6\xd8\x1b\xb8\x2e\x68\x73\x7f\x0b\xd1\x44\xc8\xbd\x90\xa2\xfb\x9e\x5e\x4f" +
			"\xad\xff\x82\xa3\xe9\x24\xe3\x50\x51\x63\x07\xe5\x02\xda\x46\xa1\x6a\xc2\x26\x08\x26\x63\x4f\x38" +
			"\x3d\xc4\x51\xd2\x7b\x77\x8d\x39\x89\xd2\x1f\x19\x7a\x02\xad\x1c\x9c\xdb\x0b\xa1\xdf\xec\x99\x1d" +
			"\xd0\xf6\x4a\x43\x16\xbe\x37\xd6\x1f\x1a\xd0\x26\x29\xbf\xf9\x6f\x91\x66\x64\x6b\xfa\x78\x7d\x65" +
			"\x6f\x1b\x36\x52\x20\x81\x18\xa4\x0b\x25\x18\xb3\xd9\x30\x37\x10\x39\x39\x8d\xd2\xf1\x96\xed\xa1" +
			"\x19\x7a\xf7\xee\xd2\xf0\x1d\xf6\x98\xb1\x15\x29\x60\xc4\xd0\x3d\xf8\xad\x67\x4c\xa0\xa0\x72\x57" +
			"\x63\x7f\xff\x52\x43\x49\x49\x1b\x2a\x00\x47\x84\x97\xe8\xd6\xfb\xfa\xf8\x74\x77\x2f\x8f\x9d\x16" +
			"\x79\x55\x71\xb9\x0c\x86\x44\xf0\x3b\xd5\x73\x7f\xd6\x23\x9c\xb6\x8a\xd4\xa0\x51\x2e\x1a\x5e\x74" +
			"\xb2\xb2\x89\x6f\x7a\xb7\x9e\x90\xd9\x47\x46\x0c\xfc\x7d\x8b\x57\xf1\xcd\x62\xea\xa9\x57\x87\xff" +
			"\xcd\x63\x96\xf5\x2e\xdd\xad\x0d\x1f\x54\x1d\x13\xf7\x1d\xe3\xe7\x33\x54\xa0\x2d\x8b\x40\xbd\x26" +
			"\xc4\xf7\x5f\x73\x97\x8e\xc4\xa6\x6a\xea\x3c\xdb\x43\x9b\x6b\x84\x9d\xa7\xd7\x5c\x1e\xbd\x72\xd6" +
			"\x34\x66\x46\x65\x36\x14\x80\x4b\xef\x3e\x3d\x6c\x5a\x21\x2d\x18\x2d\xe7\x41\xd5\x94\xe3\x60\x7a" +
			"\x7b\xfe\x37\x00\x00\xff\xff",
		size: 4129,
		mode: 0666,
		time: time.Unix(1529710666, 430493700),
	},
	"template/minimal/main.js.tmpl": &asset{
		name: "main.js.tmpl",
		data: "" +
			"\x94\x90\xcd\x4e\x42\x31\x10\x85\xf7\x24\xbc\xc3\xa4\xab\x7b\x13\xb9\x3c\x00\x2b\x51\x16\x26\x18" +
			"\x17\x3e\xc1\xd0\x0e\xbd\x95\x61\x06\xdb\xde\x46\x62\x78\x77\x53\x21\x44\x22\xd1\xb0\xea\x4f\xe6" +
			"\x9c\x39\xe7\x9b\x4e\x61\xa9\xe8\x12\xa4\xbd\x64\xfc\x80\x3e\xf8\x9e\x83\xef\x73\x10\x0f\x28\x0e" +
			"\x72\x4f\xb0\xd5\x55\x60\x02\xc1\x12\x3c\xe6\xa0\xd2\x8d\x47\x4e\xed\xb0\x25\xc9\x1d\x3a\xb7\x28" +
			"\x24\x79\x19\x52\x26\xa1\xd8\x98\xc7\x97\xe7\x07\x95\x5c\xff\x14\x1d\x39\x73\x07\xeb\x41\x6c\x15" +
			"\x36\x2d\x7c\x8e\x47\x00\x67\xf5\xfb\x40\x71\xff\x4a\x4c\x36\x6b\xbc\x67\x6e\x8c\x0b\x05\x76\x91" +
			"\x4c\xdb\xad\x35\x2e\xd0\xf6\xcd\x59\x5c\x90\x07\x3a\x39\x00\xf4\xfc\x96\xba\x73\xde\x39\xab\xdd" +
			"\x9c\x26\x66\x75\xe0\x50\x8f\x8b\x5d\x9e\xf2\x82\xa9\x5e\xe7\xfb\x27\xd7\x18\xc1\x32\x59\x0d\xd1" +
			"\x53\x34\xed\x95\x1a\x96\x83\xdd\x5c\xc9\xfe\x87\xe3\x11\xd4\x44\xb0\x98\xb6\xb3\x8c\x29\x55\xbb" +
			"\x2e\xab\xf7\x4c\xc7\x85\xba\x23\x31\xed\xf5\x80\xbf\x61\x74\x55\x62\x95\x19\x77\x89\x80\x03\xe0" +
			"\xff\x5c\xbe\x5f\x37\xf5\xb9\xbd\x51\xa4\xad\x96\x8b\x46\xb3\xa3\xd5\xe1\x07\xfc\x43\x3b\xfb\x0a" +
			"\x00\x00\xff\xff",
		size: 607,
		mode: 0666,
		time: time.Unix(1529800820, 581305700),
	},
}

// AssetAndInfo loads and returns the asset and asset info for the
// given name. It returns an error if the asset could not be found
// or could not be loaded.
func AssetAndInfo(name string) ([]byte, os.FileInfo, error) {
	a, ok := _bindata[filepath.ToSlash(name)]
	if !ok {
		return nil, nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
	}

	a.once.Do(func() {
		fr := flate.NewReader(strings.NewReader(a.data))

		var buf bytes.Buffer
		if _, a.err = io.Copy(&buf, fr); a.err != nil {
			return
		}

		if a.err = fr.Close(); a.err == nil {
			a.bytes = buf.Bytes()
		}
	})
	if a.err != nil {
		return nil, nil, &os.PathError{Op: "read", Path: name, Err: a.err}
	}

	return a.bytes, a, nil
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	a, ok := _bindata[filepath.ToSlash(name)]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
	}

	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	data, _, err := AssetAndInfo(name)
	return data, err
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

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}

	return names
}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	return restore.Asset(dir, name, AssetAndInfo)
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	return restore.Assets(dir, name, AssetDir, AssetAndInfo)
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

	if name != "" {
		var ok bool
		for _, p := range strings.Split(filepath.ToSlash(name), "/") {
			if node, ok = node[p]; !ok {
				return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
			}
		}
	}

	if len(node) == 0 {
		return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
	}

	rv := make([]string, 0, len(node))
	for name := range node {
		rv = append(rv, name)
	}

	return rv, nil
}

type bintree map[string]bintree

var _bintree = bintree{
	"template": bintree{
		"minimal": bintree{
			"index.html.tmpl": bintree{},
			"main.css.tmpl":   bintree{},
			"main.js.tmpl":    bintree{},
		},
	},
}
