package local

import (
	"go.nandlabs.io/commons/vfs"
	"net/url"
	"os"
)

var fileScheme = "file"
var emptyScheme = ""
var localFsSchemes []string = []string{"file", ""}

type OsFs struct {
	*vfs.BaseVFS
}

func (l *OsFs) CopyAll(src, dst string) error {
	//TODO implement me
	panic("implement me")
}

func (l *OsFs) Create(u string) (vfs.VFile, error) {
	//TODO implement me
	panic("implement me")
}

func (l *OsFs) Home() (vfs.VFile, error) {
	panic("implement me")
}
func (l *OsFs) Open(u *url.URL) (file vfs.VFile, err error) {
	if u.Scheme == "" {
		u.Scheme = fileScheme
	}
	var f *os.File
	f, err = os.Open(u.Path)
	file = &OsFile{
		file: f,
		u:    u,
	}

	return
}

func (l *OsFs) Root() (vfs.VFile, error) {
	panic("implement me")
}

func (l *OsFs) Mkdir(u string) (vfs.VFile, error) {
	//TODO implement me
	panic("implement me")
}

func (l *OsFs) MkdirAll(u string) (vfs.VFile, error) {
	//TODO implement me
	panic("implement me")
}

func (l *OsFs) Schemes() []string {
	return localFsSchemes
}
