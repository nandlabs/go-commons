package vfs

import (
	"io/fs"
	"io/ioutil"
	"net/url"
	"os"

	"go.nandlabs.io/commons/errutils"
	"go.nandlabs.io/commons/fsutils"
	"go.nandlabs.io/commons/textutils"
)

type OsFile struct {
	*BaseFile
	file     *os.File
	Location *url.URL
	fs       VFileSystem
}

func (o *OsFile) Close() error {
	return o.file.Close()
}

func (o *OsFile) Read(b []byte) (int, error) {
	return o.file.Read(b)
}

func (o *OsFile) Write(b []byte) (int, error) {
	return o.file.Write(b)
}

func (o *OsFile) Seek(offset int64, whence int) (int64, error) {
	return o.file.Seek(offset, whence)
}

func (o *OsFile) ContentType() string {
	return fsutils.LookupContentType(o.Location.Path)
}

func (o *OsFile) ListAll() (files []VFile, err error) {
	var fis []fs.FileInfo
	manager := GetManager()
	fis, err = ioutil.ReadDir(o.Location.Path)
	if err == nil {
		var children []VFile
		var child VFile
		var childUrl *url.URL
		for _, fi := range fis {
			childUrl, err = o.Location.Parse(textutils.ForwardSlashStr + fi.Name())
			if err == nil {
				child, err = manager.Open(childUrl)
				if err == nil {
					children = append(children, child)
				} else {
					break
				}
			} else {
				break
			}
		}
		if err == nil {
			files = children
		}
	}

	return
}

func (o *OsFile) Delete() error {
	return os.Remove(o.Location.Path)
}

func (o *OsFile) Find(filter FileFilter) (files []VFile, err error) {
	err = o.fs.Walk(o.Location, func(file VFile) (err error) {
		var filterPass bool
		filterPass, err = filter(file)
		if err == nil && filterPass {
			files = append(files, file)
		}
		return
	})
	return
}

func (o *OsFile) Info() (VFileInfo, error) {
	return o.file.Stat()
}

func (o *OsFile) Parent() (file VFile, err error) {
	var fileInfos []fs.FileInfo
	fileInfos, err = ioutil.ReadDir(o.Location.Path)
	if err == nil {
		for _, info := range fileInfos {
			var f *os.File
			var u *url.URL
			u, _ = o.Location.Parse("../" + info.Name())
			f, err = os.Open(u.Path)
			if err == nil {
				file = &OsFile{
					file:     f,
					Location: u,
				}
			}
		}
	}
	return
}

func (o *OsFile) Url() *url.URL {
	return o.Location
}

func (o *OsFile) AddProperty(name string, value string) error {
	return errutils.FmtError("Unsupported operation AddProperty for scheme")
}

func (o *OsFile) GetProperty(name string) (v string, err error) {
	err = errutils.FmtError("Unsupported operation GetProperty for scheme")
	return

}
