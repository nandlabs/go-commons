package local

import (
	"go.nandlabs.io/commons/vfs"
	"net/url"
	"os"
)

type OsFile struct {
	*vfs.BaseFile
	file *os.File
	u    *url.URL
}

func (l *OsFile) Child(name string) (file vfs.VFile, err error) {
	//var fi vfs.VFileInfo
	//fi, err = l.file.
	//	ioutil.ReadDir()
	//
	return
}

func (l *OsFile) Children() ([]vfs.VFile, error) {
	panic("implement me")
}

func (l *OsFile) Close() error {
	return l.file.Close()
}

func (l *OsFile) Copy(destUrl string) error {

	panic("implement me")
}

func (l *OsFile) Delete() error {
	//TODO implement me
	panic("implement me")
}

func (l *OsFile) DeleteMatching(filter vfs.FileFilter) error {
	//TODO implement me
	panic("implement me")
}

func (l *OsFile) Encoding() string {
	//TODO implement me
	panic("implement me")
}

func (l *OsFile) Exists() (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (l *OsFile) Find(filter vfs.FileFilter) ([]vfs.VFile, error) {
	//TODO implement me
	panic("implement me")
}

func (l *OsFile) Info() (vfs.VFileInfo, error) {
	return l.file.Stat()
}

func (l *OsFile) IsRoot() (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (l *OsFile) Move(destUrl string) error {
	//TODO implement me
	panic("implement me")
}

func (l *OsFile) Parent() (vfs.VFile, error) {
	//TODO implement me
	panic("implement me")
}
func (l *OsFile) Read(b []byte) (n int, err error) {
	return l.file.Read(b)
}

func (l *OsFile) Seek(offset int64, whence int) (int64, error) {
	return l.file.Seek(offset, whence)

}

func (l *OsFile) Type() string {
	//TODO implement me
	panic("implement me")
}

func (l *OsFile) Url() *url.URL {
	return l.Url()
}

func (l *OsFile) Write(b []byte) (n int, err error) {
	return l.file.Write(b)
}
