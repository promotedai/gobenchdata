// Code generated by fileb0x at "2019-05-04 21:13:29.864085 -0700 PDT m=+0.002442422" from config file "b0x.yml" DO NOT EDIT.
// modification hash(b941deb49cfd9723171b7707d0b1c2da.844d3c884b87b2b2ca4c2e55c010f13a)

package internal

import (
	"bytes"

	"context"
	"io"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct {
	// Prefix allows to limit the path of all requests. F.e. a prefix "css" would allow only calls to /css/*
	Prefix string
}

// FileWebAppJs is "web/app.js"
var FileWebAppJs = []byte("\x27\x75\x73\x65\x20\x73\x74\x72\x69\x63\x74\x27\x3b\x0a\x0a\x2f\x2f\x20\x47\x65\x6e\x65\x72\x61\x74\x65\x20\x6f\x6e\x65\x20\x63\x68\x61\x72\x74\x20\x70\x65\x72\x20\x73\x75\x69\x74\x65\x0a\x65\x78\x70\x6f\x72\x74\x20\x61\x73\x79\x6e\x63\x20\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x67\x65\x6e\x65\x72\x61\x74\x65\x43\x68\x61\x72\x74\x73\x28\x7b\x0a\x20\x20\x64\x69\x76\x2c\x20\x20\x2f\x2f\x20\x64\x69\x76\x20\x74\x6f\x20\x70\x6f\x70\x75\x6c\x61\x74\x65\x20\x77\x69\x74\x68\x20\x63\x68\x61\x72\x74\x73\x20\x0a\x20\x20\x6a\x73\x6f\x6e\x2c\x20\x2f\x2f\x20\x70\x61\x74\x68\x20\x74\x6f\x20\x4a\x53\x4f\x4e\x20\x64\x61\x74\x61\x62\x61\x73\x65\x0a\x7d\x29\x20\x7b\x0a\x20\x20\x6c\x65\x74\x20\x72\x75\x6e\x73\x20\x3d\x20\x5b\x5d\x3b\x0a\x20\x20\x74\x72\x79\x20\x7b\x0a\x20\x20\x20\x20\x72\x75\x6e\x73\x20\x3d\x20\x61\x77\x61\x69\x74\x20\x72\x65\x61\x64\x4a\x53\x4f\x4e\x28\x6a\x73\x6f\x6e\x29\x3b\x0a\x20\x20\x7d\x20\x63\x61\x74\x63\x68\x20\x28\x65\x29\x20\x7b\x0a\x20\x20\x20\x20\x63\x6f\x6e\x73\x74\x20\x65\x72\x72\x20\x3d\x20\x64\x6f\x63\x75\x6d\x65\x6e\x74\x2e\x63\x72\x65\x61\x74\x65\x45\x6c\x65\x6d\x65\x6e\x74\x28\x27\x64\x69\x76\x27\x29\x3b\x0a\x20\x20\x20\x20\x64\x69\x76\x2e\x61\x70\x70\x65\x6e\x64\x43\x68\x69\x6c\x64\x28\x65\x72\x72\x29\x3b\x0a\x20\x20\x20\x20\x65\x72\x72\x2e\x69\x6e\x6e\x65\x72\x54\x65\x78\x74\x20\x3d\x20\x65\x3b\x0a\x20\x20\x7d\x0a\x0a\x20\x20\x63\x6f\x6e\x73\x74\x20\x6c\x61\x62\x65\x6c\x73\x20\x3d\x20\x72\x75\x6e\x73\x2e\x73\x6f\x72\x74\x28\x28\x61\x2c\x20\x62\x29\x20\x3d\x3e\x20\x61\x2e\x44\x61\x74\x65\x20\x2d\x20\x62\x2e\x44\x61\x74\x65\x29\x2e\x6d\x61\x70\x28\x72\x20\x3d\x3e\x20\x6c\x61\x62\x65\x6c\x28\x72\x29\x29\x3b\x0a\x20\x20\x63\x6f\x6e\x73\x74\x20\x63\x68\x61\x72\x74\x73\x20\x3d\x20\x7b\x7d\x3b\x0a\x20\x20\x6c\x65\x74\x20\x6c\x65\x6e\x20\x3d\x20\x30\x3b\x0a\x20\x20\x2f\x2f\x20\x72\x75\x6e\x73\x20\x73\x68\x6f\x75\x6c\x64\x20\x73\x74\x61\x72\x74\x20\x66\x72\x6f\x6d\x20\x74\x68\x65\x20\x6d\x6f\x73\x74\x20\x72\x65\x63\x65\x6e\x74\x20\x72\x75\x6e\x0a\x20\x20\x72\x75\x6e\x73\x2e\x66\x6f\x72\x45\x61\x63\x68\x28\x72\x75\x6e\x20\x3d\x3e\x20\x7b\x0a\x20\x20\x20\x20\x6c\x65\x6e\x2b\x2b\x3b\x0a\x0a\x20\x20\x20\x20\x2f\x2f\x20\x61\x64\x64\x20\x64\x61\x74\x61\x20\x66\x72\x6f\x6d\x20\x65\x61\x63\x68\x20\x73\x75\x69\x74\x65\x0a\x20\x20\x20\x20\x72\x75\x6e\x2e\x53\x75\x69\x74\x65\x73\x2e\x66\x6f\x72\x45\x61\x63\x68\x28\x73\x75\x69\x74\x65\x20\x3d\x3e\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x69\x66\x20\x28\x63\x68\x61\x72\x74\x73\x5b\x73\x75\x69\x74\x65\x2e\x50\x6b\x67\x5d\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2f\x2f\x20\x69\x66\x20\x74\x68\x65\x20\x63\x68\x61\x72\x74\x20\x64\x69\x76\x20\x77\x61\x73\x20\x61\x6c\x72\x65\x61\x64\x79\x20\x73\x65\x74\x20\x75\x70\x2c\x20\x61\x70\x70\x65\x6e\x64\x20\x64\x61\x74\x61\x20\x74\x6f\x20\x63\x68\x61\x72\x74\x2e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2f\x2f\x20\x69\x66\x20\x74\x68\x65\x20\x64\x61\x74\x61\x73\x65\x74\x20\x69\x73\x20\x69\x73\x6e\x27\x74\x20\x69\x6e\x20\x74\x68\x65\x20\x64\x61\x74\x61\x73\x65\x74\x73\x2c\x20\x74\x68\x65\x6e\x20\x69\x74\x20\x6e\x6f\x20\x6c\x6f\x6e\x67\x65\x72\x20\x65\x78\x69\x73\x74\x73\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2f\x2f\x20\x61\x6e\x64\x20\x77\x65\x27\x6c\x6c\x20\x69\x67\x6e\x6f\x72\x65\x20\x69\x74\x2e\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x73\x75\x69\x74\x65\x2e\x42\x65\x6e\x63\x68\x6d\x61\x72\x6b\x73\x2e\x66\x6f\x72\x45\x61\x63\x68\x28\x62\x65\x6e\x63\x68\x20\x3d\x3e\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x63\x6f\x6e\x73\x74\x20\x7b\x20\x64\x61\x74\x61\x3a\x20\x7b\x20\x64\x61\x74\x61\x73\x65\x74\x73\x20\x7d\x20\x7d\x20\x3d\x20\x63\x68\x61\x72\x74\x73\x5b\x73\x75\x69\x74\x65\x2e\x50\x6b\x67\x5d\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x63\x6f\x6e\x73\x74\x20\x64\x61\x74\x61\x73\x65\x74\x20\x3d\x20\x64\x61\x74\x61\x73\x65\x74\x73\x2e\x66\x69\x6e\x64\x28\x65\x20\x3d\x3e\x20\x28\x65\x2e\x6c\x61\x62\x65\x6c\x20\x3d\x3d\x3d\x20\x62\x65\x6e\x63\x68\x2e\x4e\x61\x6d\x65\x29\x29\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x69\x66\x20\x28\x64\x61\x74\x61\x73\x65\x74\x29\x20\x64\x61\x74\x61\x73\x65\x74\x2e\x64\x61\x74\x61\x2e\x70\x75\x73\x68\x28\x6e\x65\x77\x50\x6f\x69\x6e\x74\x28\x72\x75\x6e\x2c\x20\x62\x65\x6e\x63\x68\x2e\x4e\x73\x50\x65\x72\x4f\x70\x29\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x7d\x20\x65\x6c\x73\x65\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2f\x2f\x20\x63\x72\x65\x61\x74\x65\x20\x65\x6c\x65\x6d\x65\x6e\x74\x73\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x63\x6f\x6e\x73\x74\x20\x63\x61\x6e\x76\x61\x73\x20\x3d\x20\x64\x6f\x63\x75\x6d\x65\x6e\x74\x2e\x63\x72\x65\x61\x74\x65\x45\x6c\x65\x6d\x65\x6e\x74\x28\x27\x63\x61\x6e\x76\x61\x73\x27\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x63\x61\x6e\x76\x61\x73\x2e\x69\x64\x20\x3d\x20\x73\x75\x69\x74\x65\x2e\x50\x6b\x67\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x63\x6f\x6e\x73\x74\x20\x63\x74\x78\x20\x3d\x20\x63\x61\x6e\x76\x61\x73\x2e\x67\x65\x74\x43\x6f\x6e\x74\x65\x78\x74\x28\x27\x32\x64\x27\x29\x3b\x0a\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2f\x2f\x20\x63\x72\x65\x61\x74\x65\x20\x63\x68\x61\x72\x74\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x6c\x65\x74\x20\x69\x20\x3d\x20\x72\x61\x6e\x64\x6f\x6d\x49\x6e\x74\x28\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x63\x6f\x6e\x73\x74\x20\x7b\x20\x42\x65\x6e\x63\x68\x6d\x61\x72\x6b\x73\x3a\x20\x62\x65\x6e\x63\x68\x6d\x61\x72\x6b\x73\x20\x7d\x20\x3d\x20\x73\x75\x69\x74\x65\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x63\x68\x61\x72\x74\x73\x5b\x73\x75\x69\x74\x65\x2e\x50\x6b\x67\x5d\x20\x3d\x20\x6e\x65\x77\x20\x43\x68\x61\x72\x74\x28\x63\x74\x78\x2c\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x74\x79\x70\x65\x3a\x20\x27\x6c\x69\x6e\x65\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x64\x61\x74\x61\x3a\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6c\x61\x62\x65\x6c\x73\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x64\x61\x74\x61\x73\x65\x74\x73\x3a\x20\x62\x65\x6e\x63\x68\x6d\x61\x72\x6b\x73\x2e\x6d\x61\x70\x28\x62\x65\x6e\x63\x68\x20\x3d\x3e\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x69\x20\x2b\x3d\x20\x33\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6c\x61\x62\x65\x6c\x3a\x20\x62\x65\x6e\x63\x68\x2e\x4e\x61\x6d\x65\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x64\x61\x74\x61\x3a\x20\x5b\x6e\x65\x77\x50\x6f\x69\x6e\x74\x28\x72\x75\x6e\x2c\x20\x62\x65\x6e\x63\x68\x2e\x4e\x73\x50\x65\x72\x4f\x70\x29\x5d\x2c\x0a\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x66\x69\x6c\x6c\x3a\x20\x66\x61\x6c\x73\x65\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x62\x61\x63\x6b\x67\x72\x6f\x75\x6e\x64\x43\x6f\x6c\x6f\x72\x3a\x20\x67\x65\x74\x43\x6f\x6c\x6f\x72\x28\x69\x29\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x62\x6f\x72\x64\x65\x72\x43\x6f\x6c\x6f\x72\x3a\x20\x67\x65\x74\x43\x6f\x6c\x6f\x72\x28\x69\x29\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x70\x6f\x69\x6e\x74\x52\x61\x64\x69\x75\x73\x3a\x20\x34\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x29\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6f\x70\x74\x69\x6f\x6e\x73\x3a\x20\x63\x68\x61\x72\x74\x4f\x70\x74\x69\x6f\x6e\x73\x28\x73\x75\x69\x74\x65\x29\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x29\x3b\x0a\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x2f\x2f\x20\x61\x74\x74\x61\x63\x68\x20\x74\x6f\x20\x64\x6f\x6d\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x63\x6f\x6e\x73\x74\x20\x63\x61\x6e\x76\x61\x73\x44\x69\x76\x20\x3d\x20\x64\x6f\x63\x75\x6d\x65\x6e\x74\x2e\x63\x72\x65\x61\x74\x65\x45\x6c\x65\x6d\x65\x6e\x74\x28\x27\x64\x69\x76\x27\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x63\x61\x6e\x76\x61\x73\x44\x69\x76\x2e\x73\x65\x74\x41\x74\x74\x72\x69\x62\x75\x74\x65\x28\x27\x63\x6c\x61\x73\x73\x27\x2c\x20\x27\x63\x61\x6e\x76\x61\x73\x77\x72\x61\x70\x70\x65\x72\x27\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x63\x61\x6e\x76\x61\x73\x44\x69\x76\x2e\x61\x70\x70\x65\x6e\x64\x43\x68\x69\x6c\x64\x28\x63\x61\x6e\x76\x61\x73\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x64\x69\x76\x2e\x61\x70\x70\x65\x6e\x64\x43\x68\x69\x6c\x64\x28\x63\x61\x6e\x76\x61\x73\x44\x69\x76\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x64\x69\x76\x2e\x61\x70\x70\x65\x6e\x64\x43\x68\x69\x6c\x64\x28\x64\x6f\x63\x75\x6d\x65\x6e\x74\x2e\x63\x72\x65\x61\x74\x65\x45\x6c\x65\x6d\x65\x6e\x74\x28\x27\x62\x72\x27\x29\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x7d\x29\x3b\x0a\x0a\x20\x20\x20\x20\x2f\x2f\x20\x66\x69\x6c\x6c\x20\x6d\x69\x73\x73\x69\x6e\x67\x20\x64\x61\x74\x61\x20\x66\x72\x6f\x6d\x20\x64\x61\x74\x61\x73\x65\x74\x73\x0a\x20\x20\x20\x20\x4f\x62\x6a\x65\x63\x74\x2e\x76\x61\x6c\x75\x65\x73\x28\x63\x68\x61\x72\x74\x73\x29\x2e\x66\x6f\x72\x45\x61\x63\x68\x28\x63\x20\x3d\x3e\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x63\x6f\x6e\x73\x74\x20\x7b\x20\x64\x61\x74\x61\x3a\x20\x7b\x20\x64\x61\x74\x61\x73\x65\x74\x73\x20\x7d\x20\x7d\x20\x3d\x20\x63\x3b\x0a\x20\x20\x20\x20\x20\x20\x64\x61\x74\x61\x73\x65\x74\x73\x2e\x66\x6f\x72\x45\x61\x63\x68\x28\x64\x20\x3d\x3e\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x63\x6f\x6e\x73\x74\x20\x7b\x20\x64\x61\x74\x61\x20\x7d\x20\x3d\x20\x64\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x69\x66\x20\x28\x64\x61\x74\x61\x2e\x6c\x65\x6e\x67\x74\x68\x20\x3c\x20\x6c\x65\x6e\x29\x20\x64\x61\x74\x61\x2e\x75\x6e\x73\x68\x69\x66\x74\x28\x6e\x65\x77\x50\x6f\x69\x6e\x74\x28\x72\x75\x6e\x2c\x20\x4e\x61\x4e\x29\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x7d\x29\x3b\x0a\x20\x20\x20\x20\x7d\x29\x0a\x20\x20\x7d\x29\x0a\x7d\x0a\x0a\x63\x6f\x6e\x73\x74\x20\x63\x68\x61\x72\x74\x4f\x70\x74\x69\x6f\x6e\x73\x20\x3d\x20\x28\x73\x75\x69\x74\x65\x29\x20\x3d\x3e\x20\x28\x7b\x0a\x20\x20\x72\x65\x73\x70\x6f\x6e\x73\x69\x76\x65\x3a\x20\x74\x72\x75\x65\x2c\x0a\x20\x20\x61\x73\x70\x65\x63\x74\x52\x61\x74\x69\x6f\x3a\x20\x31\x2e\x35\x2c\x0a\x20\x20\x74\x69\x74\x6c\x65\x3a\x20\x7b\x0a\x20\x20\x20\x20\x64\x69\x73\x70\x6c\x61\x79\x3a\x20\x74\x72\x75\x65\x2c\x0a\x20\x20\x20\x20\x74\x65\x78\x74\x3a\x20\x73\x75\x69\x74\x65\x2e\x50\x6b\x67\x2c\x0a\x20\x20\x7d\x2c\x0a\x20\x20\x74\x6f\x6f\x6c\x74\x69\x70\x73\x3a\x20\x7b\x0a\x20\x20\x20\x20\x6d\x6f\x64\x65\x3a\x20\x27\x69\x6e\x64\x65\x78\x27\x2c\x0a\x20\x20\x20\x20\x69\x6e\x74\x65\x72\x73\x65\x63\x74\x3a\x20\x66\x61\x6c\x73\x65\x2c\x0a\x20\x20\x7d\x2c\x0a\x20\x20\x68\x6f\x76\x65\x72\x3a\x20\x7b\x0a\x20\x20\x20\x20\x6d\x6f\x64\x65\x3a\x20\x27\x6e\x65\x61\x72\x65\x73\x74\x27\x2c\x0a\x20\x20\x20\x20\x69\x6e\x74\x65\x72\x73\x65\x63\x74\x3a\x20\x74\x72\x75\x65\x0a\x20\x20\x7d\x2c\x0a\x20\x20\x73\x63\x61\x6c\x65\x73\x3a\x20\x7b\x0a\x20\x20\x20\x20\x79\x41\x78\x65\x73\x3a\x20\x5b\x7b\x0a\x20\x20\x20\x20\x20\x20\x64\x69\x73\x70\x6c\x61\x79\x3a\x20\x74\x72\x75\x65\x2c\x0a\x20\x20\x20\x20\x20\x20\x73\x63\x61\x6c\x65\x4c\x61\x62\x65\x6c\x3a\x20\x7b\x20\x64\x69\x73\x70\x6c\x61\x79\x3a\x20\x74\x72\x75\x65\x2c\x20\x6c\x61\x62\x65\x6c\x53\x74\x72\x69\x6e\x67\x3a\x20\x27\x6e\x73\x27\x20\x7d\x2c\x0a\x20\x20\x20\x20\x20\x20\x74\x69\x63\x6b\x73\x3a\x20\x7b\x20\x62\x65\x67\x69\x6e\x41\x74\x5a\x65\x72\x6f\x3a\x20\x74\x72\x75\x65\x20\x7d\x2c\x0a\x20\x20\x20\x20\x7d\x5d\x2c\x0a\x20\x20\x7d\x2c\x0a\x7d\x29\x0a\x0a\x63\x6f\x6e\x73\x74\x20\x6e\x65\x77\x50\x6f\x69\x6e\x74\x20\x3d\x20\x28\x72\x75\x6e\x2c\x20\x76\x61\x6c\x29\x20\x3d\x3e\x20\x28\x7b\x0a\x20\x20\x74\x3a\x20\x6e\x65\x77\x20\x44\x61\x74\x65\x28\x72\x75\x6e\x2e\x44\x61\x74\x65\x2a\x31\x30\x30\x30\x29\x2c\x0a\x20\x20\x79\x3a\x20\x76\x61\x6c\x2c\x0a\x7d\x29\x0a\x0a\x0a\x63\x6f\x6e\x73\x74\x20\x63\x68\x61\x72\x74\x43\x6f\x6c\x6f\x72\x73\x20\x3d\x20\x7b\x0a\x09\x72\x65\x64\x3a\x20\x27\x72\x67\x62\x28\x32\x35\x35\x2c\x20\x39\x39\x2c\x20\x31\x33\x32\x29\x27\x2c\x0a\x09\x6f\x72\x61\x6e\x67\x65\x3a\x20\x27\x72\x67\x62\x28\x32\x35\x35\x2c\x20\x31\x35\x39\x2c\x20\x36\x34\x29\x27\x2c\x0a\x09\x79\x65\x6c\x6c\x6f\x77\x3a\x20\x27\x72\x67\x62\x28\x32\x35\x35\x2c\x20\x32\x30\x35\x2c\x20\x38\x36\x29\x27\x2c\x0a\x09\x67\x72\x65\x65\x6e\x3a\x20\x27\x72\x67\x62\x28\x37\x35\x2c\x20\x31\x39\x32\x2c\x20\x31\x39\x32\x29\x27\x2c\x0a\x09\x62\x6c\x75\x65\x3a\x20\x27\x72\x67\x62\x28\x35\x34\x2c\x20\x31\x36\x32\x2c\x20\x32\x33\x35\x29\x27\x2c\x0a\x09\x70\x75\x72\x70\x6c\x65\x3a\x20\x27\x72\x67\x62\x28\x31\x35\x33\x2c\x20\x31\x30\x32\x2c\x20\x32\x35\x35\x29\x27\x2c\x0a\x09\x67\x72\x65\x79\x3a\x20\x27\x72\x67\x62\x28\x32\x30\x31\x2c\x20\x32\x30\x33\x2c\x20\x32\x30\x37\x29\x27\x0a\x7d\x3b\x0a\x0a\x63\x6f\x6e\x73\x74\x20\x63\x68\x61\x72\x74\x43\x6f\x6c\x6f\x72\x73\x4c\x69\x73\x74\x20\x3d\x20\x4f\x62\x6a\x65\x63\x74\x2e\x76\x61\x6c\x75\x65\x73\x28\x63\x68\x61\x72\x74\x43\x6f\x6c\x6f\x72\x73\x29\x3b\x0a\x0a\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x72\x61\x6e\x64\x6f\x6d\x49\x6e\x74\x28\x29\x20\x7b\x0a\x20\x20\x72\x65\x74\x75\x72\x6e\x20\x4d\x61\x74\x68\x2e\x66\x6c\x6f\x6f\x72\x28\x4d\x61\x74\x68\x2e\x72\x61\x6e\x64\x6f\x6d\x28\x29\x20\x2a\x20\x28\x63\x68\x61\x72\x74\x43\x6f\x6c\x6f\x72\x73\x4c\x69\x73\x74\x2e\x6c\x65\x6e\x67\x74\x68\x20\x2b\x20\x31\x29\x29\x3b\x0a\x7d\x0a\x0a\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x67\x65\x74\x43\x6f\x6c\x6f\x72\x28\x69\x29\x20\x7b\x0a\x20\x20\x72\x65\x74\x75\x72\x6e\x20\x63\x68\x61\x72\x74\x43\x6f\x6c\x6f\x72\x73\x4c\x69\x73\x74\x5b\x69\x20\x25\x20\x63\x68\x61\x72\x74\x43\x6f\x6c\x6f\x72\x73\x4c\x69\x73\x74\x2e\x6c\x65\x6e\x67\x74\x68\x5d\x3b\x0a\x7d\x0a\x0a\x61\x73\x79\x6e\x63\x20\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x72\x65\x61\x64\x4a\x53\x4f\x4e\x28\x70\x61\x74\x68\x29\x20\x7b\x0a\x20\x20\x72\x65\x74\x75\x72\x6e\x20\x28\x61\x77\x61\x69\x74\x20\x66\x65\x74\x63\x68\x28\x70\x61\x74\x68\x29\x29\x2e\x6a\x73\x6f\x6e\x28\x29\x3b\x0a\x7d\x0a\x0a\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x6c\x61\x62\x65\x6c\x28\x72\x75\x6e\x29\x20\x7b\x0a\x20\x20\x63\x6f\x6e\x73\x74\x20\x64\x20\x3d\x20\x6e\x65\x77\x20\x44\x61\x74\x65\x28\x72\x75\x6e\x2e\x44\x61\x74\x65\x2a\x31\x30\x30\x30\x29\x3b\x0a\x20\x20\x6c\x65\x74\x20\x6d\x6f\x6e\x74\x68\x20\x3d\x20\x64\x2e\x67\x65\x74\x4d\x6f\x6e\x74\x68\x28\x29\x3b\x0a\x20\x20\x72\x65\x74\x75\x72\x6e\x20\x60\x24\x7b\x72\x75\x6e\x2e\x56\x65\x72\x73\x69\x6f\x6e\x2e\x73\x75\x62\x73\x74\x72\x69\x6e\x67\x28\x30\x2c\x20\x37\x29\x7d\x20\x28\x24\x7b\x2b\x2b\x6d\x6f\x6e\x74\x68\x7d\x2f\x24\x7b\x64\x2e\x67\x65\x74\x44\x61\x74\x65\x28\x29\x7d\x29\x60\x3b\x0a\x7d\x0a")

// FileWebIndexHTML is "web/index.html"
var FileWebIndexHTML = []byte("\x3c\x21\x44\x4f\x43\x54\x59\x50\x45\x20\x68\x74\x6d\x6c\x3e\x0a\x3c\x68\x74\x6d\x6c\x20\x6c\x61\x6e\x67\x3d\x22\x65\x6e\x22\x3e\x0a\x20\x20\x3c\x68\x65\x61\x64\x3e\x0a\x20\x20\x20\x20\x3c\x6d\x65\x74\x61\x20\x63\x68\x61\x72\x73\x65\x74\x3d\x22\x75\x74\x66\x2d\x38\x22\x20\x2f\x3e\x0a\x20\x20\x20\x20\x3c\x6c\x69\x6e\x6b\x20\x72\x65\x6c\x3d\x22\x73\x68\x6f\x72\x74\x63\x75\x74\x20\x69\x63\x6f\x6e\x22\x20\x68\x72\x65\x66\x3d\x22\x25\x50\x55\x42\x4c\x49\x43\x5f\x55\x52\x4c\x25\x2f\x66\x61\x76\x69\x63\x6f\x6e\x2e\x69\x63\x6f\x22\x20\x2f\x3e\x0a\x20\x20\x20\x20\x3c\x6d\x65\x74\x61\x0a\x20\x20\x20\x20\x20\x20\x6e\x61\x6d\x65\x3d\x22\x76\x69\x65\x77\x70\x6f\x72\x74\x22\x0a\x20\x20\x20\x20\x20\x20\x63\x6f\x6e\x74\x65\x6e\x74\x3d\x22\x77\x69\x64\x74\x68\x3d\x64\x65\x76\x69\x63\x65\x2d\x77\x69\x64\x74\x68\x2c\x20\x69\x6e\x69\x74\x69\x61\x6c\x2d\x73\x63\x61\x6c\x65\x3d\x31\x2c\x20\x73\x68\x72\x69\x6e\x6b\x2d\x74\x6f\x2d\x66\x69\x74\x3d\x6e\x6f\x22\x0a\x20\x20\x20\x20\x2f\x3e\x0a\x0a\x20\x20\x20\x20\x3c\x6d\x65\x74\x61\x20\x6e\x61\x6d\x65\x3d\x22\x74\x68\x65\x6d\x65\x2d\x63\x6f\x6c\x6f\x72\x22\x20\x63\x6f\x6e\x74\x65\x6e\x74\x3d\x22\x23\x30\x30\x30\x30\x30\x30\x22\x20\x2f\x3e\x0a\x20\x20\x20\x20\x3c\x73\x74\x79\x6c\x65\x3e\x0a\x20\x20\x20\x20\x20\x20\x2a\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x74\x65\x78\x74\x2d\x61\x6c\x69\x67\x6e\x3a\x20\x63\x65\x6e\x74\x65\x72\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x63\x6f\x6c\x6f\x72\x3a\x20\x23\x36\x36\x36\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x66\x6f\x6e\x74\x2d\x66\x61\x6d\x69\x6c\x79\x3a\x20\x27\x48\x65\x6c\x76\x65\x74\x69\x63\x61\x20\x4e\x65\x75\x65\x27\x3b\x0a\x20\x20\x20\x20\x20\x20\x7d\x0a\x0a\x20\x20\x20\x20\x20\x20\x2e\x68\x69\x64\x64\x65\x6e\x20\x7b\x20\x64\x69\x73\x70\x6c\x61\x79\x3a\x20\x6e\x6f\x6e\x65\x3b\x20\x7d\x0a\x0a\x20\x20\x20\x20\x20\x20\x62\x6f\x64\x79\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x70\x61\x64\x64\x69\x6e\x67\x2d\x6c\x65\x66\x74\x3a\x20\x20\x34\x30\x30\x70\x78\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x70\x61\x64\x64\x69\x6e\x67\x2d\x72\x69\x67\x68\x74\x3a\x20\x34\x30\x30\x70\x78\x3b\x0a\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x40\x6d\x65\x64\x69\x61\x20\x6f\x6e\x6c\x79\x20\x73\x63\x72\x65\x65\x6e\x20\x61\x6e\x64\x20\x28\x6d\x61\x78\x2d\x77\x69\x64\x74\x68\x3a\x20\x31\x36\x30\x30\x70\x78\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x62\x6f\x64\x79\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x70\x61\x64\x64\x69\x6e\x67\x2d\x6c\x65\x66\x74\x3a\x20\x20\x33\x30\x30\x70\x78\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x70\x61\x64\x64\x69\x6e\x67\x2d\x72\x69\x67\x68\x74\x3a\x20\x33\x30\x30\x70\x78\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x40\x6d\x65\x64\x69\x61\x20\x6f\x6e\x6c\x79\x20\x73\x63\x72\x65\x65\x6e\x20\x61\x6e\x64\x20\x28\x6d\x61\x78\x2d\x77\x69\x64\x74\x68\x3a\x20\x31\x32\x30\x30\x70\x78\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x62\x6f\x64\x79\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x70\x61\x64\x64\x69\x6e\x67\x2d\x6c\x65\x66\x74\x3a\x20\x20\x31\x35\x30\x70\x78\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x70\x61\x64\x64\x69\x6e\x67\x2d\x72\x69\x67\x68\x74\x3a\x20\x31\x35\x30\x70\x78\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x40\x6d\x65\x64\x69\x61\x20\x6f\x6e\x6c\x79\x20\x73\x63\x72\x65\x65\x6e\x20\x61\x6e\x64\x20\x28\x6d\x61\x78\x2d\x77\x69\x64\x74\x68\x3a\x20\x39\x30\x30\x70\x78\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x62\x6f\x64\x79\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x70\x61\x64\x64\x69\x6e\x67\x2d\x6c\x65\x66\x74\x3a\x20\x20\x35\x30\x70\x78\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x70\x61\x64\x64\x69\x6e\x67\x2d\x72\x69\x67\x68\x74\x3a\x20\x35\x30\x70\x78\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x40\x6d\x65\x64\x69\x61\x20\x6f\x6e\x6c\x79\x20\x73\x63\x72\x65\x65\x6e\x20\x61\x6e\x64\x20\x28\x6d\x61\x78\x2d\x77\x69\x64\x74\x68\x3a\x20\x36\x30\x30\x70\x78\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x62\x6f\x64\x79\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x70\x61\x64\x64\x69\x6e\x67\x2d\x6c\x65\x66\x74\x3a\x20\x20\x30\x70\x78\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x70\x61\x64\x64\x69\x6e\x67\x2d\x72\x69\x67\x68\x74\x3a\x20\x30\x70\x78\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x3c\x2f\x73\x74\x79\x6c\x65\x3e\x0a\x0a\x20\x20\x20\x20\x3c\x73\x63\x72\x69\x70\x74\x20\x73\x72\x63\x3d\x22\x68\x74\x74\x70\x73\x3a\x2f\x2f\x63\x64\x6e\x2e\x6a\x73\x64\x65\x6c\x69\x76\x72\x2e\x6e\x65\x74\x2f\x6e\x70\x6d\x2f\x63\x68\x61\x72\x74\x2e\x6a\x73\x40\x32\x2e\x38\x2e\x30\x22\x3e\x3c\x2f\x73\x63\x72\x69\x70\x74\x3e\x0a\x0a\x20\x20\x20\x20\x3c\x74\x69\x74\x6c\x65\x3e\x7b\x7b\x2e\x54\x69\x74\x6c\x65\x7d\x7d\x3c\x2f\x74\x69\x74\x6c\x65\x3e\x0a\x20\x20\x3c\x2f\x68\x65\x61\x64\x3e\x0a\x0a\x20\x20\x3c\x62\x6f\x64\x79\x3e\x0a\x20\x20\x20\x20\x3c\x21\x2d\x2d\x20\x64\x65\x66\x69\x6e\x65\x20\x74\x65\x6d\x70\x6c\x61\x74\x65\x20\x76\x61\x72\x69\x61\x62\x6c\x65\x73\x20\x69\x6e\x20\x64\x69\x76\x73\x20\x74\x6f\x20\x61\x76\x6f\x69\x64\x20\x62\x72\x65\x61\x6b\x69\x6e\x67\x20\x6a\x73\x20\x63\x6f\x64\x65\x20\x69\x6e\x20\x64\x65\x76\x67\x20\x2d\x2d\x3e\x0a\x20\x20\x20\x20\x3c\x64\x69\x76\x20\x69\x64\x3d\x22\x62\x65\x6e\x63\x68\x6d\x61\x72\x6b\x73\x2d\x70\x61\x74\x68\x22\x20\x68\x69\x64\x64\x65\x6e\x3e\x7b\x7b\x2e\x42\x65\x6e\x63\x68\x6d\x61\x72\x6b\x73\x50\x61\x74\x68\x7d\x7d\x3c\x2f\x64\x69\x76\x3e\x0a\x20\x20\x0a\x20\x20\x20\x20\x3c\x21\x2d\x2d\x20\x68\x65\x61\x64\x65\x72\x73\x20\x2d\x2d\x3e\x0a\x20\x20\x20\x20\x3c\x64\x69\x76\x3e\x0a\x20\x20\x20\x20\x20\x20\x3c\x68\x31\x3e\x7b\x7b\x2e\x54\x69\x74\x6c\x65\x7d\x7d\x3c\x2f\x68\x31\x3e\x0a\x20\x20\x20\x20\x20\x20\x3c\x68\x34\x3e\x67\x65\x6e\x65\x72\x61\x74\x65\x64\x20\x77\x69\x74\x68\x20\x3c\x61\x20\x68\x72\x65\x66\x3d\x22\x68\x74\x74\x70\x73\x3a\x2f\x2f\x67\x69\x74\x68\x75\x62\x2e\x63\x6f\x6d\x2f\x6d\x61\x72\x6b\x65\x74\x70\x6c\x61\x63\x65\x2f\x61\x63\x74\x69\x6f\x6e\x73\x2f\x67\x6f\x62\x65\x6e\x63\x68\x64\x61\x74\x61\x2d\x74\x6f\x2d\x67\x68\x2d\x70\x61\x67\x65\x73\x22\x3e\x67\x6f\x62\x65\x6e\x63\x68\x64\x61\x74\x61\x3c\x2f\x61\x3e\x3c\x2f\x68\x34\x3e\x0a\x20\x20\x20\x20\x3c\x2f\x64\x69\x76\x3e\x0a\x0a\x20\x20\x20\x20\x3c\x62\x72\x20\x2f\x3e\x0a\x0a\x20\x20\x20\x20\x3c\x21\x2d\x2d\x20\x63\x68\x61\x72\x74\x20\x63\x6f\x6e\x74\x65\x6e\x74\x2c\x20\x67\x65\x6e\x65\x72\x61\x74\x65\x64\x20\x75\x73\x69\x6e\x67\x20\x61\x70\x70\x2e\x6a\x73\x20\x2d\x2d\x3e\x0a\x20\x20\x20\x20\x3c\x64\x69\x76\x3e\x0a\x20\x20\x20\x20\x20\x20\x3c\x64\x69\x76\x20\x69\x64\x3d\x22\x63\x68\x61\x72\x74\x73\x22\x3e\x3c\x2f\x64\x69\x76\x3e\x0a\x20\x20\x20\x20\x3c\x2f\x64\x69\x76\x3e\x0a\x0a\x20\x20\x20\x20\x3c\x21\x2d\x2d\x20\x6a\x61\x76\x61\x73\x63\x72\x69\x70\x74\x20\x2d\x2d\x3e\x0a\x20\x20\x20\x20\x3c\x73\x63\x72\x69\x70\x74\x20\x74\x79\x70\x65\x3d\x22\x6d\x6f\x64\x75\x6c\x65\x22\x3e\x0a\x20\x20\x20\x20\x20\x20\x69\x6d\x70\x6f\x72\x74\x20\x7b\x20\x67\x65\x6e\x65\x72\x61\x74\x65\x43\x68\x61\x72\x74\x73\x20\x7d\x20\x66\x72\x6f\x6d\x20\x27\x2e\x2f\x61\x70\x70\x2e\x6a\x73\x27\x3b\x0a\x20\x20\x20\x20\x20\x20\x63\x6f\x6e\x73\x74\x20\x63\x68\x61\x72\x74\x73\x20\x3d\x20\x64\x6f\x63\x75\x6d\x65\x6e\x74\x2e\x67\x65\x74\x45\x6c\x65\x6d\x65\x6e\x74\x42\x79\x49\x64\x28\x27\x63\x68\x61\x72\x74\x73\x27\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x63\x6f\x6e\x73\x74\x20\x62\x65\x6e\x63\x68\x6d\x61\x72\x6b\x73\x20\x3d\x20\x64\x6f\x63\x75\x6d\x65\x6e\x74\x2e\x67\x65\x74\x45\x6c\x65\x6d\x65\x6e\x74\x42\x79\x49\x64\x28\x27\x62\x65\x6e\x63\x68\x6d\x61\x72\x6b\x73\x2d\x70\x61\x74\x68\x27\x29\x3b\x0a\x20\x20\x20\x20\x20\x20\x77\x69\x6e\x64\x6f\x77\x2e\x6f\x6e\x6c\x6f\x61\x64\x20\x3d\x20\x67\x65\x6e\x65\x72\x61\x74\x65\x43\x68\x61\x72\x74\x73\x28\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x64\x69\x76\x3a\x20\x63\x68\x61\x72\x74\x73\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x6a\x73\x6f\x6e\x3a\x20\x28\x62\x65\x6e\x63\x68\x6d\x61\x72\x6b\x73\x2e\x69\x6e\x6e\x65\x72\x54\x65\x78\x74\x20\x3d\x3d\x3d\x20\x27\x7b\x7b\x2e\x42\x65\x6e\x63\x68\x6d\x61\x72\x6b\x73\x50\x61\x74\x68\x7d\x7d\x27\x29\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3f\x20\x27\x62\x65\x6e\x63\x68\x6d\x61\x72\x6b\x73\x2e\x6a\x73\x6f\x6e\x27\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x62\x65\x6e\x63\x68\x6d\x61\x72\x6b\x73\x2e\x69\x6e\x6e\x65\x72\x54\x65\x78\x74\x2c\x0a\x20\x20\x20\x20\x20\x20\x7d\x29\x2e\x66\x69\x6e\x61\x6c\x6c\x79\x28\x63\x6f\x6e\x73\x6f\x6c\x65\x2e\x6c\x6f\x67\x28\x27\x63\x68\x61\x72\x74\x73\x20\x6c\x6f\x61\x64\x65\x64\x21\x27\x29\x29\x3b\x0a\x20\x20\x20\x20\x3c\x2f\x73\x63\x72\x69\x70\x74\x3e\x0a\x20\x20\x3c\x2f\x62\x6f\x64\x79\x3e\x0a\x3c\x2f\x68\x74\x6d\x6c\x3e\x0a")

// FileWebStyleCSS is "web/style.css"
var FileWebStyleCSS = []byte("\x2a\x20\x7b\x0a\x20\x20\x74\x65\x78\x74\x2d\x61\x6c\x69\x67\x6e\x3a\x20\x63\x65\x6e\x74\x65\x72\x3b\x0a\x20\x20\x63\x6f\x6c\x6f\x72\x3a\x20\x23\x36\x36\x36\x3b\x0a\x20\x20\x66\x6f\x6e\x74\x2d\x66\x61\x6d\x69\x6c\x79\x3a\x20\x27\x48\x65\x6c\x76\x65\x74\x69\x63\x61\x20\x4e\x65\x75\x65\x27\x3b\x0a\x7d\x0a\x0a\x2e\x68\x69\x64\x64\x65\x6e\x20\x7b\x20\x64\x69\x73\x70\x6c\x61\x79\x3a\x20\x6e\x6f\x6e\x65\x3b\x20\x7d\x0a\x0a\x62\x6f\x64\x79\x20\x7b\x0a\x20\x20\x70\x61\x64\x64\x69\x6e\x67\x2d\x6c\x65\x66\x74\x3a\x20\x20\x34\x30\x30\x70\x78\x3b\x0a\x20\x20\x70\x61\x64\x64\x69\x6e\x67\x2d\x72\x69\x67\x68\x74\x3a\x20\x34\x30\x30\x70\x78\x3b\x0a\x7d\x0a\x40\x6d\x65\x64\x69\x61\x20\x6f\x6e\x6c\x79\x20\x73\x63\x72\x65\x65\x6e\x20\x61\x6e\x64\x20\x28\x6d\x61\x78\x2d\x77\x69\x64\x74\x68\x3a\x20\x31\x36\x30\x30\x70\x78\x29\x20\x7b\x0a\x20\x20\x62\x6f\x64\x79\x20\x7b\x0a\x20\x20\x20\x20\x70\x61\x64\x64\x69\x6e\x67\x2d\x6c\x65\x66\x74\x3a\x20\x20\x33\x30\x30\x70\x78\x3b\x0a\x20\x20\x20\x20\x70\x61\x64\x64\x69\x6e\x67\x2d\x72\x69\x67\x68\x74\x3a\x20\x33\x30\x30\x70\x78\x3b\x0a\x20\x20\x7d\x0a\x7d\x0a\x40\x6d\x65\x64\x69\x61\x20\x6f\x6e\x6c\x79\x20\x73\x63\x72\x65\x65\x6e\x20\x61\x6e\x64\x20\x28\x6d\x61\x78\x2d\x77\x69\x64\x74\x68\x3a\x20\x31\x32\x30\x30\x70\x78\x29\x20\x7b\x0a\x20\x20\x62\x6f\x64\x79\x20\x7b\x0a\x20\x20\x20\x20\x70\x61\x64\x64\x69\x6e\x67\x2d\x6c\x65\x66\x74\x3a\x20\x20\x31\x35\x30\x70\x78\x3b\x0a\x20\x20\x20\x20\x70\x61\x64\x64\x69\x6e\x67\x2d\x72\x69\x67\x68\x74\x3a\x20\x31\x35\x30\x70\x78\x3b\x0a\x20\x20\x7d\x0a\x7d\x0a\x40\x6d\x65\x64\x69\x61\x20\x6f\x6e\x6c\x79\x20\x73\x63\x72\x65\x65\x6e\x20\x61\x6e\x64\x20\x28\x6d\x61\x78\x2d\x77\x69\x64\x74\x68\x3a\x20\x39\x30\x30\x70\x78\x29\x20\x7b\x0a\x20\x20\x62\x6f\x64\x79\x20\x7b\x0a\x20\x20\x20\x20\x70\x61\x64\x64\x69\x6e\x67\x2d\x6c\x65\x66\x74\x3a\x20\x20\x35\x30\x70\x78\x3b\x0a\x20\x20\x20\x20\x70\x61\x64\x64\x69\x6e\x67\x2d\x72\x69\x67\x68\x74\x3a\x20\x35\x30\x70\x78\x3b\x0a\x20\x20\x7d\x0a\x7d\x0a\x40\x6d\x65\x64\x69\x61\x20\x6f\x6e\x6c\x79\x20\x73\x63\x72\x65\x65\x6e\x20\x61\x6e\x64\x20\x28\x6d\x61\x78\x2d\x77\x69\x64\x74\x68\x3a\x20\x36\x30\x30\x70\x78\x29\x20\x7b\x0a\x20\x20\x62\x6f\x64\x79\x20\x7b\x0a\x20\x20\x20\x20\x70\x61\x64\x64\x69\x6e\x67\x2d\x6c\x65\x66\x74\x3a\x20\x20\x30\x70\x78\x3b\x0a\x20\x20\x20\x20\x70\x61\x64\x64\x69\x6e\x67\x2d\x72\x69\x67\x68\x74\x3a\x20\x30\x70\x78\x3b\x0a\x20\x20\x7d\x0a\x7d\x0a")

func init() {
	err := CTX.Err()
	if err != nil {
		panic(err)
	}

	err = FS.Mkdir(CTX, "web/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	var f webdav.File

	f, err = FS.OpenFile(CTX, "web/app.js", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(FileWebAppJs)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "web/index.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(FileWebIndexHTML)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "web/style.css", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(FileWebStyleCSS)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}

}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {
	path = hfs.Prefix + path

	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}
