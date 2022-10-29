package vfs

import (
	"net/url"
)

type VFileSystem interface {
	//Copy files
	Copy(src, dst string) error
	//CopyAll files,directories recursivey
	CopyAll(src, dst string) error
	//CreateFile function
	CreateFile(u *url.URL) (VFile, error)
	//Create creates a file
	Create(u string) (VFile, error)
	//Delete file . if the src resolves to a directory then all the files  and directories under this will be deleted
	Delete(src string) error
	//Exists indicates if this url exists in the file system
	Exists(u *url.URL) (bool, error)
	//FileExists function is same as Exists function, however it accepts the url as a string
	FileExists(u string) (bool, error)
	//List lists the file in
	List(url string) ([]VFileInfo, error)
	//ListFiles lists the file in the filesystem for a specific url
	ListFiles(url string) ([]VFileInfo, error)
	//Mkdir will create the directory and will throw an error if exists
	Mkdir(u string) (VFile, error)
	//MkdirAll  will create all directories missing in the path
	//If the directory already exists it will not throw error, however if the path resolves to a file instead
	//it will throw  error
	MkdirAll(u string) (VFile, error)
	//Move Files
	Move(src, dst string) error
	//Open a file based on the url of the file
	Open(u *url.URL) (VFile, error)
	// OpenFile is same as Open function, however it accepts the url as string
	OpenFile(u string) (VFile, error)
	//Schemes is the list of schemes supported by this file system
	Schemes() []string
}
