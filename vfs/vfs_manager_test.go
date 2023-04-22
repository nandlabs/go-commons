package vfs

import (
	"fmt"
	"net/url"
	"os"
	"reflect"
	"testing"
)

func GetParsedUrl(input string) (output *url.URL) {
	currentPath, _ := os.Getwd()
	u, _ := url.Parse(input)
	path := currentPath + u.Path
	output, _ = url.Parse(u.Scheme + "://" + path)
	return
}

func GetRawPath(input string) (output string) {
	currentPath, _ := os.Getwd()
	u, _ := url.Parse(input)
	path := currentPath + u.Path
	output = u.Scheme + "://" + path
	return
}

func TestGetManager(t *testing.T) {
	tests := []struct {
		name string
		want VFileSystem
	}{
		// TODO: Add test cases.
		{
			name: "localFs",
			want: &fileSystems{
				fileSystems: map[string]VFileSystem{"file": &OsFs{BaseVFS: &BaseVFS{VFileSystem: &OsFs{}}}, "": &OsFs{BaseVFS: &BaseVFS{VFileSystem: &OsFs{}}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetManager(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileSystems_MkdirRaw(t *testing.T) {
	testManager := GetManager()
	u := GetRawPath("file:///tests")
	gotDir, err := testManager.MkdirRaw(u)
	if err != nil {
		t.Errorf("MkdirRaw() error = %v", err)
		return
	}
	info, err := gotDir.Info()
	if info.Name() != "tests" {
		t.Errorf("Info() got DirName = %v, want %v", info.Name(), "tests")
	}
}

func TestFileSystems_Mkdir(t *testing.T) {
	testManager := GetManager()
	u := GetParsedUrl("file:///tests1")
	_, err := testManager.Mkdir(u)
	if err != nil {
		t.Errorf("Mkdir() error = %v", err)
	}
}

func TestFileSystems_MkdirAll(t *testing.T) {
	testManager := GetManager()
	u := GetParsedUrl("file:///all/tests")
	_, err := testManager.MkdirAll(u)
	if err != nil {
		t.Errorf("MkdirAll() error = %v", err)
	}
}

func TestFileSystems_MkdirAllRaw(t *testing.T) {
	testManager := GetManager()
	u := GetRawPath("file:///all-raw/tests")
	_, err := testManager.MkdirAllRaw(u)
	if err != nil {
		t.Errorf("MkdirAllRaw() error = %v", err)
	}
}

func TestFileSystems_RemoveDir(t *testing.T) {
	testManager := GetManager()
	u := GetParsedUrl("file:///dummy")
	_, err := testManager.Mkdir(u)
	if err != nil {
		t.Errorf("Mkdir() error = %v", err)
	}
	err = testManager.RemoveDir(u)
	if err != nil {
		t.Errorf("RemoveDir() error = %v", err)
	}
}

func TestFileSystems_CreateRaw(t *testing.T) {
	testManager := GetManager()
	u := GetRawPath("file:///raw-abc.txt")
	_, err := testManager.CreateRaw(u)
	if err != nil {
		t.Errorf("CreateRaw() error = %v", err)
	}
}

func TestFileSystems_Delete(t *testing.T) {
	testManager := GetManager()
	u := GetParsedUrl("file:///raw-abc.txt")
	err := testManager.Delete(u)
	if err != nil {
		t.Errorf("Delete() error = %v", err)
	}
}

func TestFileSystems_DeleteRaw(t *testing.T) {
	testManager := GetManager()
	u := GetRawPath("file:///raw-abc-1.txt")

	_, err := testManager.CreateRaw(u)
	if err != nil {
		t.Errorf("CreateRaw() error = %v", err)
	}

	err = testManager.DeleteRaw(u)
	if err != nil {
		t.Errorf("DeleteRaw() error = %v", err)
	}
}

func TestFileSystems_List(t *testing.T) {
	testManager := GetManager()

	u := GetParsedUrl("file:///tests")
	_, err := testManager.List(u)
	if err != nil {
		t.Errorf("List() error = %v", err)
	}
}

func TestFileSystems_ListRaw(t *testing.T) {
	testManager := GetManager()

	u := GetRawPath("file:///tests")
	_, err := testManager.ListRaw(u)
	if err != nil {
		t.Errorf("ListRaw() error = %v", err)
	}
}

func TestFileSystems_Move(t *testing.T) {
	testManager := GetManager()

	src := GetParsedUrl("file:///tests/abc.txt")
	dest := GetParsedUrl("file:///tests/move-abc.txt")
	_, err := testManager.Create(src)
	if err != nil {
		t.Errorf("Create() error = %v", err)
	}
	err = testManager.Move(src, dest)
	if err != nil {
		t.Errorf("Move() error = %v", err)
	}
}

func TestFileSystems_MoveRaw(t *testing.T) {
	testManager := GetManager()

	src := GetRawPath("file:///tests/abc-raw.txt")
	dest := GetRawPath("file:///tests/move-abc-raw.txt")
	_, err := testManager.CreateRaw(src)
	if err != nil {
		t.Errorf("CreateRaw() error = %v", err)
	}
	err = testManager.MoveRaw(src, dest)
	if err != nil {
		t.Errorf("MoveRaw() error = %v", err)
	}
}

func TestFileSystems_Open(t *testing.T) {
	testManager := GetManager()

	u := GetParsedUrl("file:///tests/move-abc-raw.txt")
	file, err := testManager.Open(u)
	if err != nil {
		t.Errorf("Open() error = %v", err)
	}
	info, err := file.Info()
	fmt.Println(info.Name())
}

func TestFileSystems_OpenRaw(t *testing.T) {
	testManager := GetManager()

	u := GetRawPath("file:///tests/move-abc-raw.txt")
	_, err := testManager.OpenRaw(u)
	if err != nil {
		t.Errorf("OpenRaw() error = %v", err)
	}
}

// TODO : not implemented yet
//func TestFileSystems_Copy(t *testing.T) {
//	testManager := GetManager()
//
//	src := GetParsedUrl("file:///tests/abc-copy-tst.txt")
//	dest := GetParsedUrl("file:///tests1")
//	_, err := testManager.Create(src)
//	if err != nil {
//		t.Errorf("Create() error = %v", err)
//	}
//	err = testManager.Copy(src, dest)
//	if err != nil {
//		t.Errorf("Copy() error = %v", err)
//	}
//}

func TestFileSystems_Schemes(t *testing.T) {
	testManager := GetManager()

	output := testManager.Schemes()
	fmt.Println(output)
}

func Test_InvalidFS(t *testing.T) {
	testManager := GetManager()
	u := GetRawPath("dummy:///raw-abc.txt")
	_, err := testManager.CreateRaw(u)
	if err == nil {
		t.Errorf("CreateRaw() error = %v", err)
	}
	if err.Error() != "Unsupported scheme dummy for in the url "+u {
		t.Errorf("Test_InvalidFS() error = %v", err)
	}
}

func TestFileSystems_IsSupported(t *testing.T) {
	var scheme string
	var output bool

	testManager := GetManager()

	scheme = "file"
	output = testManager.IsSupported(scheme)
	if output == false {
		t.Error()
	}
	scheme = "test"
	output = testManager.IsSupported(scheme)
	if output == true {
		t.Error()
	}
}

func Test_RemoveAllTestsDirs(t *testing.T) {
	os.RemoveAll("./tests")
	os.RemoveAll("./all")
	os.RemoveAll("./all-raw")
	os.RemoveAll("./tests1")
}
