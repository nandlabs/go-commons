package vfs

import (
	"errors"
	"net/url"

	"go.nandlabs.io/commons/ioutils"
)

type BaseVFS struct {
	VFileSystem
}

func (b *BaseVFS) OpenLocation(l string) (file VFile, err error) {
	var u *url.URL
	u, err = url.Parse(l)
	if err == nil {
		file, err = b.Open(u)
	}
	return
}

func (b *BaseVFS) Copy(src, dst string) (err error) {
	var srcFile VFile
	var srfFileInfo VFileInfo
	srcFile, err = b.OpenFile(src)
	if err == nil {
		defer ioutils.CloserFunc(srcFile)
		srfFileInfo, err = srcFile.Info()
		if err == nil {
			if srfFileInfo.IsDir() {
				err = errors.New("unable to execute  Copy directory use CopyAll")
			} else {
				err = srcFile.Copy(dst)
			}
		}
	}
	return
}

func (b *BaseVFS) Move(src, dst string) (err error) {
	var srcFile VFile
	srcFile, err = b.OpenFile(src)
	if err == nil {
		defer ioutils.CloserFunc(srcFile)
		err = srcFile.Move(dst)
	}
	return
}

func (b *BaseVFS) Delete(src string) (err error) {
	var srcFile VFile
	srcFile, err = b.OpenFile(src)
	if err == nil {
		defer ioutils.CloserFunc(srcFile)
		err = srcFile.Delete()
	}
	return
}
