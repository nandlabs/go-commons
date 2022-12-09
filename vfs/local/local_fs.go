package local

import (
	"net/url"
	"os"

	"go.nandlabs.io/commons/vfs"
)

const (
	fileScheme  = "file"
	emptyScheme = ""
)

var localFsSchemes []string = []string{fileScheme, emptyScheme}

type OsFs struct {
	*vfs.BaseVFS
}

func (o OsFs) Create(u *url.URL) (file vfs.VFile, err error) {
	var f *os.File
	f, err = os.Create(u.Path)
	if err == nil {
		file = &OsFile{
			file:     f,
			Location: u,
			fs:       o,
		}
	}
	return
}

func (o OsFs) Mkdir(u *url.URL) (file vfs.VFile, err error) {

	err = os.Mkdir(u.Path, os.ModePerm)
	if err == nil {
		file, err = o.Open(u)
	}
	return
}

func (o OsFs) MkdirAll(u *url.URL) (file vfs.VFile, err error) {
	err = os.MkdirAll(u.Path, os.ModePerm)
	if err == nil {
		file, err = o.Open(u)
	}
	return
}

func (o OsFs) Open(u *url.URL) (file vfs.VFile, err error) {
	var f *os.File
	f, err = os.Open(u.Path)
	if err == nil {
		file = &OsFile{
			file:     f,
			Location: u,
			fs:       o,
		}
	}
	return
}

func (o OsFs) Schemes() []string {
	return localFsSchemes
}
