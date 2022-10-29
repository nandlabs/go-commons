package vfs

import (
	"io"
	"io/fs"
	"net/url"
)

type FileFilter func(u string) (bool, error)

//VFile interface provides the basic functions required to interact
type VFile interface {
	//Closer interface included from io package
	io.Closer
	//VFileContent provider interface included
	VFileContent
	//Child of this instance identified by name
	Child(name string) (VFile, error)
	//Children of this file instance. can be nil in case of file object instead of directory
	Children() ([]VFile, error)
	//Copy files to the destination url. This can be in a different VFileSystem if the scheme is supported
	Copy(destUrl string) error
	//Move Files to the destination url.This can be in a different VFileSystem if the scheme is supported
	Move(destUrl string) error
	//Delete the file object. If the file type is directory all  files and subdirectories will be deleted
	Delete() error
	//DeleteMatching will delete only the files that match the filter.
	//Throws error if the files is not a dir type
	//If one of the file deletion fails with an error then it stops processing and returns error
	DeleteMatching(filter FileFilter) error
	//Find files based on filter only works if the file.IsDir() is true
	Find(filter FileFilter) ([]VFile, error)
	//Info  Get the file ifo
	Info() (VFileInfo, error)
	// IsRoot specifies if this is the root of the filesystem
	IsRoot() (bool, error)
	//Parent of the file system
	Parent() (VFile, error)
	//Url of the file
	Url() *url.URL
	// AddProperty will add a property to the file
	AddProperty(name string, value string) error
}

//VFileContent interface providers access to the content
type VFileContent interface {
	io.ReadWriteSeeker
	//AsString content of the file. This should be used very carefully as it is not wise to load a large file in to string
	AsString() (string, error)
	//AsBytes content of the file.This should be used very carefully as it is not wise to load a large file into an array
	AsBytes() ([]byte, error)
	//WriteString method with write the string
	WriteString(s string) (int, error)
	//Encoding of the underlying content. If not set defaults to UTF-8 for text files
	Encoding() string
	//Type of the content stored. This should be the
	Type() string
}

type VFileInfo interface {
	fs.FileInfo
}
