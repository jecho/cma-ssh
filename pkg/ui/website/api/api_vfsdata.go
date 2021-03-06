// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package api

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Api statically implements the virtual filesystem provided to vfsgen.
var Api = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 12, 14, 22, 22, 56, 20265104, time.UTC),
		},
		"/api.proto": &vfsgen۰CompressedFileInfo{
			name:             "api.proto",
			modTime:          time.Date(2018, 12, 14, 22, 22, 56, 19463401, time.UTC),
			uncompressedSize: 8834,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x39\x5b\x73\xdb\x36\xd6\xef\xfa\x15\x67\xf4\xf2\x39\xdf\x24\x52\xe2\xb4\xdd\x8e\xbd\xde\x5d\xad\xec\x3a\x9a\x26\x92\xc7\x54\xda\xe9\x93\x06\x22\x8f\x48\xd4\x24\x80\x05\x40\x29\x9a\x8e\xff\xfb\xce\x01\x40\x8a\xa4\x28\x3b\x4d\xd2\x87\xd5\x43\x62\xe1\x5c\x70\xee\x17\x68\x3c\x86\xa9\x54\x7b\xcd\xd3\xcc\xc2\xf9\xeb\x37\x3f\x42\xc4\x0a\x53\x8a\x14\xa2\xeb\x08\xa6\xb9\x2c\x13\x98\x33\xcb\xb7\x08\x53\x59\xa8\xd2\x72\x91\xc2\x12\x59\x01\xac\xb4\x99\xd4\x66\x34\x18\x8f\x07\xe3\x31\xbc\xe7\x31\x0a\x83\x09\x94\x22\x41\x0d\x36\x43\x98\x28\x16\x67\x58\x41\x5e\xc2\x2f\xa8\x0d\x97\x02\xce\x47\xaf\xe1\x8c\x10\x86\x01\x34\x7c\x71\x49\x2c\xf6\xb2\x84\x82\xed\x41\x48\x0b\xa5\x41\xb0\x19\x37\xb0\xe1\x39\x02\x7e\x8a\x51\x59\xe0\x02\x62\x59\xa8\x9c\x33\x11\x23\xec\xb8\xcd\xdc\x3d\x81\x0b\x49\x02\xbf\x05\x1e\x72\x6d\x19\x17\xc0\x20\x96\x6a\x0f\x72\xd3\x44\x04\x66\x83\xd0\xf4\xc9\xac\x55\x17\xe3\xf1\x6e\xb7\x1b\x31\x27\xf0\x48\xea\x74\x9c\x7b\x54\x33\x7e\x3f\x9b\xde\xcc\xa3\x9b\x57\xe7\xa3\xd7\x81\xe8\xa3\xc8\xd1\x18\xd0\xf8\x9f\x92\x6b\x4c\x60\xbd\x07\xa6\x54\xce\x63\xb6\xce\x11\x72\xb6\x03\xa9\x81\xa5\x1a\x31\x01\x2b\x49\xe8\x9d\xe6\x64\xb7\x97\x60\xe4\xc6\xee\x98\x46\x62\x93\x70\x63\x35\x5f\x97\xb6\x65\xb3\x4a\x44\x6e\x5a\x08\x52\x00\x13\x30\x9c\x44\x30\x8b\x86\xf0\xef\x49\x34\x8b\x5e\x12\x93\x5f\x67\xcb\x77\x8b\x8f\x4b\xf8\x75\x72\x7f\x3f\x99\x2f\x67\x37\x11\x2c\xee\x61\xba\x98\x5f\xcf\x96\xb3\xc5\x3c\x82\xc5\x4f\x30\x99\xff\x06\x3f\xcf\xe6\xd7\x2f\x01\xb9\xcd\x50\x03\x7e\x52\x9a\x34\x90\x1a\x38\x59\x13\x13\x67\xba\x08\xb1\x25\xc2\x46\x7a\x91\x8c\xc2\x98\x6f\x78\x0c\x39\x13\x69\xc9\x52\x84\x54\x6e\x51\x0b\x8a\x04\x85\xba\xe0\x86\xbc\x6a\x80\x89\x84\xd8\xe4\xbc\xe0\x96\x59\x77\x74\xa4\xd7\x68\xe0\x6e\x0a\x21\x36\x9d\x4f\x97\xf0\xf7\x82\x69\x66\x25\x4f\xfe\x95\x16\x8c\xe7\xa3\x58\x16\xff\x18\x0c\xcc\x5e\x58\xf6\x09\xae\x60\xa8\xb4\xb4\xf2\xed\xf0\x72\x30\x50\x2c\x7e\xa0\xeb\xe3\x82\x19\x93\x5d\x0e\x06\xbc\x50\x52\x5b\x18\xa6\x52\xa6\x39\x8e\x99\xe2\x63\x26\x84\x0c\xb7\x8f\x1c\xe5\xf0\xb2\x46\x73\xdf\xe3\x57\x29\x8a\x57\x66\xc7\xd2\x14\xf5\x58\x2a\x87\xda\x4b\x36\x18\x78\x28\x9c\xa5\x5a\xc5\xa3\x94\x59\xdc\xb1\xbd\x07\xc7\xab\x14\xc5\x2a\x70\x19\x05\x2e\x23\xa9\x50\x30\xc5\xb7\xe7\x15\xe4\x05\x5c\xc1\x1f\x03\x00\x2e\x36\xf2\xc2\xfd\x05\x60\xb9\xcd\xf1\x02\x86\xd3\xbc\x34\x16\x35\x7c\x60\x82\xa5\xa8\x61\x72\x37\x83\x28\x7a\x07\x4a\xcb\x2d\x4f\x50\x0f\x2f\x1d\xfa\xd6\x27\xcd\x05\x0c\xb7\xaf\x47\x6f\x46\xaf\xc3\x71\x2c\x85\x65\xb1\xad\x98\xd2\x47\xb0\x82\xf8\x7e\x20\x73\xc2\x2d\xd3\x6c\x53\xda\x84\x0b\xb9\x0d\x24\xf4\x29\x75\x7e\x01\x43\x0a\x79\x73\x31\x1e\xa7\xdc\x66\xe5\x9a\x2c\x3e\x36\xde\x25\xaf\x62\x11\xdb\x71\x5c\xb0\x57\xc6\x64\x0d\x3a\x24\xd7\x5c\xc0\xf0\xd8\x57\x01\xe9\x91\xfe\x73\xff\xe0\x27\x8b\x5a\xb0\x7c\x95\xc8\xd8\x54\xf2\x7d\xc9\xbd\x09\x9a\x58\x73\x67\x5a\x52\x4b\x6a\x04\xb6\x96\xa5\x85\xcf\xb2\xdc\xe3\x00\xc0\xc4\x19\x16\x68\x2e\xe0\xdd\x72\x79\x17\x5d\x76\x4f\xe8\x20\x96\xc2\x94\xee\x64\x18\x92\x98\xee\x1b\xff\x6e\xa4\x70\x6c\x94\x96\x49\x19\x9f\x82\x3f\x5e\x0e\x06\x06\xf5\x96\xc7\x58\x4b\xe5\x15\xa6\xdc\xe4\x79\xee\x65\x72\x55\x8f\x41\xec\x31\x1c\x5c\xab\x18\xa6\x1a\x99\xc5\x8a\xee\xac\xf5\xf5\x83\x49\x5f\x80\x46\x5b\x6a\x61\x3a\xa0\x7b\x54\xf9\xfe\x45\xc3\xf1\x75\x98\xba\x34\x18\x31\xc5\x47\x64\xe9\x2a\xf8\x0e\x1f\x25\x8d\x85\x0b\x18\xba\x4c\xd9\xbe\x19\x07\x81\x86\x2d\xa4\xb5\x4c\xf6\x84\xf4\xff\x87\xe3\xc7\xe0\xe3\x96\x66\x1a\xad\xe6\xb8\xf5\x35\xc3\x58\x66\x4b\x43\x75\xb6\x56\x93\xea\x01\x70\x6b\xe0\xa1\x5c\x63\x2c\xc5\x86\xa7\xae\xa4\xc4\x52\x08\x8c\x2d\xdf\x72\xbb\xaf\x4d\x71\x8b\xb6\xb6\xc3\xe1\xef\xb6\x11\x0e\xe7\x5f\x6e\x81\x14\x9f\x36\x40\xaf\xa6\x09\xe6\x68\xb1\xc7\x81\xd7\x0e\x50\x0b\xde\xfa\xda\x96\xbd\x05\xfa\x72\xf1\x83\x24\x7f\x5a\x83\xda\x57\x0c\x72\x6e\x2c\xf9\x29\x10\x9a\x1e\x17\xbc\x27\x94\xb3\xf6\xf7\x53\xae\x20\xd8\xb7\x76\xc7\x98\x64\x7c\x5e\xa3\x52\x8b\xaa\x3a\xba\x02\xab\x0b\x97\x9b\xa1\x48\x30\xc5\x81\x52\xb3\xe1\xae\x5b\xb4\x61\x04\x99\x35\xd0\xcf\x0e\xc7\x47\x4a\x86\xf3\x6f\xa6\x60\x10\xf7\x19\xdd\x58\xf2\x7b\x69\x2c\xb0\x27\x8b\xc7\xc4\x21\x05\x2f\xcc\x65\x82\x06\xce\x5a\x67\x6d\x65\x5a\xa0\xaf\xa8\x20\xe5\x37\x2d\x20\xe4\xc2\x52\xa5\x9a\x25\x18\x64\x30\xae\x46\x30\x48\xf9\x16\xc5\x91\xd2\xb7\x68\x3f\x7a\xf4\xa0\x49\xd7\x91\x27\xa1\x47\xae\x3d\x89\xf9\xcd\xa3\x39\x28\xf8\x9c\xd3\xad\xc5\x42\x59\x9a\x18\x2b\x8b\x1c\x3b\xbd\x2d\x34\x9c\xb5\xbf\xb7\x75\x6c\xc3\xbe\xb5\xcb\x8f\xb5\x7a\xce\xf5\x8f\x83\x01\x8a\xb2\xa8\xfa\x64\xe4\x3b\x46\xdd\x2d\xe7\xd2\x82\x41\xeb\xbe\x46\xcb\xc9\xf2\x63\xb4\xfa\x38\x8f\xee\x6e\xa6\xb3\x9f\x66\x37\xd7\x70\x05\xaf\x2f\x2b\xd4\x65\x86\x70\x77\xbf\xf8\x65\x16\xcd\x16\xf3\xd9\xfc\xd6\x75\x1f\x04\x2e\x12\x6a\xcf\x68\x5c\x47\xaa\xba\x10\x37\xb0\x46\x1a\x55\x63\xd7\x43\x93\x91\xe3\xd2\x22\xbf\x82\x37\x2d\xde\xf7\x1f\xe7\xcf\xb2\xcd\x18\xf1\xa5\x10\xf5\x6c\x7d\xb7\x33\xb0\x29\xf3\x7c\x0f\xa5\xa1\x5d\xc0\x5f\x55\x71\xbb\x82\xf3\xf6\x2d\x37\xd3\xc5\x7c\x3a\x7b\xdf\x7f\x13\xb3\x60\x64\x81\xb0\x93\xfa\x81\xf8\x32\xea\x98\x98\xef\x83\x32\x89\x14\x48\x4b\x41\x43\xa4\x97\x60\xca\x38\x03\x66\x42\xfc\x10\x1a\x81\x0b\xe6\x04\x96\x1a\x84\x4c\xb0\x5e\x41\x82\x70\x0d\x21\xae\xe0\x6d\x4b\xc0\x68\xb9\xb8\xbb\xfb\x6c\xf3\xfa\xd6\x94\x04\xff\x05\xca\x2b\xf8\xae\xc5\xf2\xe6\xfe\x7e\x71\xff\x24\x3f\xda\xdd\xd6\x08\xa5\xf0\x26\x74\xc4\x9e\xea\x0a\xbe\x6f\xf1\xba\xbe\xb9\xbd\x9f\x5c\xdf\x5c\x3f\xc9\x2e\x2c\x69\x86\xf6\x49\xed\x8c\x48\x46\x93\xa0\xd1\x58\x1a\x28\xc9\x5d\xb0\x29\x85\x03\xb0\xbc\x1a\x49\x6a\xde\x57\xf0\xc3\x25\x45\x6e\x81\xc6\xd0\xea\xd1\x9d\xd1\x1a\xf1\xcb\x0a\xac\xf6\xcc\xea\x76\x2b\x49\x97\xba\x8a\x07\xeb\xd0\x56\x27\x52\x37\xae\x1f\x85\x5e\xd5\xcf\xe4\x06\x7e\x2e\xd7\xa8\x05\x92\x46\x54\x12\x29\x10\xd0\xfb\xd0\x8c\x60\x2a\x85\xd5\x32\x07\x95\x33\x51\x53\x19\x60\x1a\x21\x41\x4b\x4b\x99\xf0\x9b\x29\x89\xf3\x81\xc5\x19\x17\x18\x29\x8c\x47\x4d\x09\x1e\x7e\x34\xab\xea\xc2\x66\x74\xfe\x9a\xa1\xdb\x13\x5d\xc8\xd8\xae\xbb\xdf\x4d\x06\x3e\xd5\x65\x0e\x19\x4f\xb3\x15\xdb\x32\x9e\xb3\x35\x27\xeb\x1d\x05\xd1\x86\xad\x35\x8f\x83\x25\x4a\xd3\x31\x01\x5a\x52\x6b\x15\x90\x9a\xd1\x12\x64\x36\xb0\xcb\x78\x9c\xb9\xb5\x5f\x73\x83\x4d\x61\x7c\x55\x44\xe5\xf3\x2f\x32\x59\x43\x4f\xb7\x1f\x69\x99\xaf\x9c\x81\x56\xce\x6a\xad\x08\xfa\x5a\xfe\xde\x1d\x35\xe3\x1f\x1a\x4a\x73\x03\x26\x93\x65\x9e\x90\xca\x0c\xb6\x2c\x2f\x11\x72\xfe\x80\xc0\xd5\x85\xdb\x44\x5d\x7a\xef\xa8\xea\x7b\x0c\xae\x6d\xc9\x72\x98\xdd\x8d\x09\x5c\x71\xba\x63\xc6\x90\x13\x59\xfc\x40\xf6\xab\xf6\x2a\x88\x4b\x63\x65\x81\xda\x04\xab\xba\x67\x07\x2b\x49\x87\xa2\x14\x2e\x09\xe8\x6b\x57\x93\x60\x73\xa6\xf8\x0a\x45\xa2\x24\x17\x16\xae\xe0\x6f\xb5\xe0\x77\x9a\x6f\x89\xf4\x01\xf7\xce\x51\xc4\xc3\x98\x0c\xb8\xb0\x12\x8a\x60\xae\x26\x27\xe5\x09\x56\x44\x70\x05\x3f\x9e\xce\x13\xd7\x7b\x1a\x7b\xd1\xe9\xf0\xda\x31\xd3\x4c\x17\x1f\xc0\xdc\xbf\xb5\xa0\xb1\x87\xc0\x93\x0f\x47\xa9\x93\xa0\x65\x3c\x37\xdd\x1c\x0c\xa4\x94\xf1\x4a\x0a\xe3\x2b\x4a\xd5\xf5\x2d\x16\x35\xa2\xcb\x80\x86\x0a\xad\x35\xe4\x73\xf2\x3c\x97\xf2\x01\x13\x28\x55\x7f\x96\xf7\xb2\xee\x98\x66\xd6\x29\xae\xbe\xc0\x9b\xbd\xb1\x58\x1c\x2b\xdf\x54\xe5\xda\x69\xff\xa4\x42\xdd\xf5\xa4\xe9\x11\x66\x29\xb5\x1b\x77\xff\x9f\xf1\xa2\x5b\x49\x7b\xb8\xd5\x72\xff\xac\x56\xc7\x3b\xce\xe1\x86\xa9\xcb\x87\xa6\x6e\x6b\xac\x18\x87\x9a\xd0\xe7\xd7\xa8\x5e\x2b\x89\xb4\x19\x05\x41\x90\xb0\x77\x9e\xf6\x5d\xd8\x5d\xe0\x8f\xd3\xe0\xaf\xf2\x41\x20\x7a\xdf\xbb\x55\x55\xb5\xa3\x27\xdc\x8e\x65\x6e\x22\x1d\x84\xb9\xee\xc4\x5a\x53\x79\x9e\xb4\x64\xe8\x89\xcc\x1e\x9f\x1d\xca\xfc\x24\x49\xb8\x6f\x7b\x3d\xeb\x53\x7b\xa9\x3f\xc1\xd2\x23\xac\x2a\x0d\xba\xb5\xff\x34\x7d\x7b\x06\xac\x9d\xf8\x5d\x9f\x41\x1a\x91\xfd\xbf\x6f\x96\x66\xa6\x35\xde\x45\x5c\xf5\x76\xcf\x22\x4f\x54\xee\x06\x7e\x77\xae\xfa\xd3\x96\xfe\xbe\x65\xe9\xc3\xa8\xf1\x9e\xad\x31\x3f\xd8\x99\x78\x8b\x60\x3f\x06\x39\x01\x9f\x1f\x61\x5c\xbf\xeb\x27\xf0\xb0\x2a\xf2\x2b\xe1\xc3\xfb\xb2\xb7\xb3\x5f\xff\xea\x37\x67\x6a\xb0\xb5\x9c\x9d\x1e\xdc\x12\x93\x06\x3c\x27\x0f\x31\x88\xa2\x77\xc0\xe2\x18\x4d\xab\x61\xd5\x28\x5d\x91\x33\x69\xec\x13\x74\x0e\xdc\x1d\xdf\x5d\x23\xef\xa1\xe1\xc2\xbe\x3d\xf7\xd0\x6e\x3e\x28\x66\xcc\x4e\xea\xa4\x43\x36\xf2\x33\x03\x37\xae\x1d\xf2\x42\xe5\x58\xa0\xa0\xba\xb1\xe3\x36\xe3\xad\x21\x9f\x29\x5e\x71\x5c\x63\xcc\x4a\xe3\x7f\x06\xa1\xd0\x7c\x10\x72\x27\x56\x4e\x56\x53\x2a\x27\x00\x83\x0f\xb3\xe5\x07\x88\x99\x70\xbb\xa9\x6d\xc8\x30\x82\x89\x07\x72\x53\x31\x34\xd6\xed\xa1\xd4\x80\xd7\x39\x16\x4e\x4a\xea\xed\x6b\x46\xd3\x00\x2b\x6d\x86\xc2\x06\x37\x5d\x02\xd2\x7e\xce\x5d\xc0\xed\x21\x91\x4e\xf6\x70\x49\xc5\x90\x88\x1d\x98\x04\xf0\xdc\x79\xa1\x50\x1b\x29\xdc\x8c\xe2\x96\x13\xe7\xce\x11\x2c\x17\xd7\x8b\x8b\x83\xf2\x0d\x6d\x4c\x6b\x66\xad\x6d\xd8\x4d\x01\x17\x6b\xa6\xfe\xd1\xa2\x35\xb0\xd4\x85\xb8\x1b\xe8\x81\xa8\x39\x18\xfa\xbd\x98\xe5\xa0\xca\x75\xce\x63\xef\x7c\xde\xea\xe9\x1e\x12\xa2\xc2\xef\x07\x14\xcb\xb7\x68\x9b\xf3\xbb\x7b\x95\xf6\x0f\x50\x8d\xbe\x73\x78\x69\xf2\x2d\x69\x3c\x06\xdf\x7f\x48\xf0\x8a\xba\x6a\x74\xc7\x74\xdd\x5e\xb5\x01\xa9\x50\xfb\xcc\xa1\xe1\x69\xf1\xf3\x89\x31\xa1\x62\xd5\xf3\x00\x76\x58\xf7\x83\x29\x2d\x4b\xab\xdd\x32\xe5\x34\x39\x29\x69\xb8\x95\x7a\x5f\x23\x06\x43\xa4\xdc\x36\x16\x88\x37\x97\x5d\x46\x19\x33\x59\x55\x94\x88\x13\x4d\xa8\xdc\xf6\x71\xf1\x90\x43\x92\x9d\x1e\x15\xad\x46\x74\xaa\xc6\x39\x32\x01\xbb\x0c\x05\xac\x4b\x9e\xf7\xb2\x25\xe4\x95\xdf\x0f\xeb\x64\x0c\xac\xaf\xe9\x50\x6e\x1c\x6d\xd2\xa5\x75\x87\xab\xc4\xd3\x7d\xd7\xa2\xfb\xe5\xe0\xe1\x54\xd6\x83\x32\xed\x0f\x3c\xac\xab\x4d\x19\x64\xc3\x3e\xdf\xb7\xf8\x4c\x3d\x85\x3e\x2c\x45\x0d\xba\xb8\x02\xd6\x9b\x45\x35\xa4\xe7\xcc\x92\xe7\x80\x5b\x6f\x04\x8f\xe8\x4b\xca\x18\x74\x29\xdc\xcf\x72\x52\x74\x39\xaa\x8a\xb0\x1e\xf9\x1f\x07\x83\x8e\x4a\x8d\xa0\x70\xa0\x9e\x58\x09\xda\xac\x9a\x9d\xb1\x67\xf8\x7a\xea\x19\xee\xc9\xb1\x33\xac\x46\xe8\x36\xd9\x58\x0a\xc3\x13\x74\xf2\x93\x7e\xe1\xc9\xe9\x73\xc6\xeb\xa7\x5f\xf7\x1a\x73\x29\x13\xdd\xa9\x34\xdc\x72\x7a\x28\x75\x62\xb7\x56\x6e\x25\x8d\xe1\xb4\x85\xf9\x9f\xd1\x85\xdc\xb5\xcb\x4e\xd5\xfd\x2a\x9a\xae\xc5\x8e\x9e\xf1\xfe\x22\x1b\xf5\x28\xe0\x98\xec\xb0\xf9\x66\x24\xff\xd9\x6a\xd9\xcd\x07\x82\x93\x32\x77\x57\x3c\x66\xfc\xe2\xc6\xc0\x94\xae\xc9\x6d\xca\xfc\xf4\x16\xd7\x60\xdb\x7d\xc2\xfe\x6b\x2d\xd1\x79\x05\xd8\x51\x65\x11\x6e\x06\x63\x49\xd2\x37\x8a\x9d\x7a\x0e\x60\x49\x52\xbf\x05\x9c\x7f\x06\x7b\x8d\x85\xdc\x22\x6c\xb4\x2c\x9e\xbc\xe3\xde\xe1\x35\x6f\xf2\x94\xf5\x65\x6f\x3b\xf5\xbd\x97\xe6\xa8\xc2\x9f\x1a\x78\x8e\x87\x9e\x37\x75\xb1\x38\xe5\xa4\xaf\x75\xfd\x7f\x03\x00\x00\xff\xff\xcd\xcc\xed\x25\x82\x22\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/api.proto"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr: gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
