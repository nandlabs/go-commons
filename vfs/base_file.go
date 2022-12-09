package vfs

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
)

type BaseFile struct {
	VFile
}

func (b *BaseFile) AsString() (s string, err error) {
	var bytes []byte
	bytes, err = ioutil.ReadAll(b)
	if err == nil {
		s = string(bytes)
	}
	return
}

func (b *BaseFile) AsBytes() ([]byte, error) {
	return ioutil.ReadAll(b)
}

func (b *BaseFile) DeleteMatching(filter FileFilter) (err error) {
	var info VFileInfo
	info, err = b.Info()

	if err == nil {
		if info.IsDir() {
			err = errors.New(fmt.Sprintf("Invalid operation DeleteMatching %s is a file", b.Url().String()))
		} else {
			var files []VFile
			files, err = b.Find(filter)
			if err == nil {
				for _, file := range files {
					err = file.Delete()
					if err != nil {
						break
					}
				}
			}
		}
	}

	return
}

func (b *BaseFile) WriteString(s string) (int, error) {
	return io.WriteString(b, s)
}
