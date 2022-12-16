package vfs

import (
	"go.nandlabs.io/commons/ioutils"
	"io"
	"net/url"
)

type BaseVFS struct {
	VFileSystem
}

func (b *BaseVFS) Copy(src, dst *url.URL) (err error) {
	var srcFile VFile
	var srfFileInfo VFileInfo
	srcFile, err = b.Open(src)

	if err == nil {
		defer ioutils.CloserFunc(srcFile)
		srfFileInfo, err = srcFile.Info()
		if err == nil {
			if srfFileInfo.IsDir() {

				err = b.Walk(src, func(file VFile) (err error) {
					var fileInfo VFileInfo
					fileInfo, err = srcFile.Info()
					if err == nil {
						if fileInfo.IsDir() {

						}
					}

					return
				})

				// Create directories and copy files
			} else {
				var destFile VFile
				destFile, err = manager.Create(dst)
				defer ioutils.CloserFunc(destFile)
				if err == nil {
					_, err = io.Copy(srcFile, destFile)
				}
			}
		}
	}
	return
}

func (b *BaseVFS) CopyRaw(src, dst string) (err error) {
	var srcUrl, dstUrl *url.URL
	srcUrl, err = url.Parse(src)
	if err == nil {
		dstUrl, err = url.Parse(dst)
		if err == nil {
			err = b.Copy(srcUrl, dstUrl)
		}
	}
	return
}

func (b *BaseVFS) CreateRaw(u string) (file VFile, err error) {
	var fileUrl *url.URL
	fileUrl, err = url.Parse(u)
	if err == nil {
		if err == nil {
			file, err = b.Create(fileUrl)
		}
	}
	return
}

func (b *BaseVFS) Delete(src *url.URL) (err error) {
	var srcFile VFile
	srcFile, err = b.Open(src)
	if err == nil {
		defer ioutils.CloserFunc(srcFile)
		err = srcFile.Delete()
	}
	return
}

func (b *BaseVFS) DeleteRaw(u string) (err error) {
	var fileUrl *url.URL
	fileUrl, err = url.Parse(u)
	if err == nil {
		err = b.Delete(fileUrl)
	}
	return
}

func (b *BaseVFS) List(src *url.URL) (files []VFile, err error) {
	var srcFile VFile
	srcFile, err = b.Open(src)
	if err == nil {
		defer ioutils.CloserFunc(srcFile)
		files, err = srcFile.ListAll()
	}
	return
}

func (b *BaseVFS) ListRaw(src string) (files []VFile, err error) {
	var fileUrl *url.URL
	fileUrl, err = url.Parse(src)
	if err == nil {
		files, err = b.List(fileUrl)
	}
	return
}

func (b *BaseVFS) MkdirRaw(u string) (vFile VFile, err error) {
	var fileUrl *url.URL
	fileUrl, err = url.Parse(u)
	if err == nil {
		vFile, err = b.Mkdir(fileUrl)
	}
	return
}
func (b *BaseVFS) MkdirAllRaw(u string) (vFile VFile, err error) {
	var fileUrl *url.URL
	fileUrl, err = url.Parse(u)
	if err == nil {
		vFile, err = b.MkdirAll(fileUrl)
	}
	return
}

func (b *BaseVFS) Move(src, dst *url.URL) (err error) {
	err = b.Copy(src, dst)
	if err == nil {
		err = b.Delete(src)
	}
	return
}

func (b *BaseVFS) MoveRaw(src, dst string) (err error) {
	var srcUrl, dstUrl *url.URL
	srcUrl, err = url.Parse(src)
	if err == nil {
		dstUrl, err = url.Parse(dst)
		if err == nil {
			err = b.Move(srcUrl, dstUrl)
		}
	}
	return
}

func (b *BaseVFS) OpenRaw(l string) (file VFile, err error) {
	var u *url.URL
	u, err = url.Parse(l)
	if err == nil {
		file, err = b.Open(u)
	}
	return
}

func (b *BaseVFS) Walk(u *url.URL, fn WalkFn) (err error) {
	var src VFile
	var srcFi VFileInfo
	var childInfo VFileInfo
	var children []VFile
	src, err = manager.Open(u)
	if err == nil {
		srcFi, err = src.Info()
		if err == nil {
			if srcFi.IsDir() {
				children, err = src.ListAll()
				if err == nil {
					for _, child := range children {
						childInfo, err = child.Info()
						if err == nil {

							if childInfo.IsDir() {
								err = b.Walk(child.Url(), fn)
							} else {
								err = fn(child)
							}
							if err != nil {
								break
							}
						}
					}
				}
			}

		}
	}
	return
}

func (b *BaseVFS) WalkRaw(raw string, fn WalkFn) (err error) {
	var u *url.URL
	u, err = url.Parse(raw)
	if err == nil {
		err = b.Walk(u, fn)
	}
	return
}