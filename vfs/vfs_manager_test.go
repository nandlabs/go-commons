package vfs

import (
	"fmt"
	"net/url"
	"os"
	"reflect"
	"testing"
)

func getParsedUrl(input string) (output *url.URL) {
	currentPath, _ := os.Getwd()
	u, _ := url.Parse(input)
	path := currentPath + u.Path
	output, _ = url.Parse(u.Scheme + "://" + path)
	return
}

func getRawPath(input string) (output string) {
	currentPath, _ := os.Getwd()
	u, _ := url.Parse(input)
	path := currentPath + u.Path
	output = u.Scheme + "://" + path
	return
}

func init() {
	// keep this for the time being
	os.RemoveAll("./tests")
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
				fileSystems: map[string]VFileSystem{"file": &OsFs{}, "": &OsFs{}},
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

func TestFileSystems_Mkdir(t *testing.T) {
	type fields struct {
		fileSystems map[string]VFileSystem
	}
	type args struct {
		u *url.URL
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantDirName string
		wantErr     bool
	}{
		{
			name: "Create Directory",
			fields: fields{
				fileSystems: map[string]VFileSystem{"file": &OsFs{}, "": &OsFs{}},
			},
			args: args{
				u: getParsedUrl("file:///tests"),
			},
			wantDirName: "tests",
			wantErr:     false,
		},
		{
			name: "Create Sub-Directory",
			fields: fields{
				fileSystems: map[string]VFileSystem{"file": &OsFs{}, "": &OsFs{}},
			},
			args: args{
				u: getParsedUrl("file:///tests/dummy"),
			},
			wantDirName: "dummy",
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				fileSystems: tt.fields.fileSystems,
			}
			err := fs.RemoveDir(tt.args.u)
			if err != nil && !os.IsNotExist(err) {
				t.Errorf("Error removing directory: %v", err)
			}
			gotDir, err := fs.Mkdir(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mkdir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			info, err := gotDir.Info()
			if info.Name() != tt.wantDirName {
				t.Errorf("Mkdir() got DirName = %v, want %v", info.Name(), tt.wantDirName)
			}
		})
	}
}

func TestFileSystems_Create(t *testing.T) {
	type fields struct {
		fileSystems map[string]VFileSystem
	}
	type args struct {
		u *url.URL
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantFileName string
		wantErr      bool
	}{
		{
			name: "Create File",
			fields: fields{
				fileSystems: map[string]VFileSystem{"file": &OsFs{}, "": &OsFs{}},
			},
			args: args{
				u: getParsedUrl("file:///tests/abc.txt"),
			},
			wantFileName: "abc.txt",
			wantErr:      false,
		},
		{
			name: "Create File in Sub Directory",
			fields: fields{
				fileSystems: map[string]VFileSystem{"file": &OsFs{}, "": &OsFs{}},
			},
			args: args{
				u: getParsedUrl("file:///tests/dummy/dummy.txt"),
			},
			wantFileName: "dummy.txt",
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				fileSystems: tt.fields.fileSystems,
			}
			gotFile, err := fs.Create(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			info, err := gotFile.Info()
			if info.Name() != tt.wantFileName {
				t.Errorf("Create() got FileName = %v, want %v", info.Name(), tt.wantFileName)
			}
		})
	}
}

//func TestFileSystems_Copy(t *testing.T) {
//	type fields struct {
//		fileSystems map[string]VFileSystem
//	}
//	type args struct {
//		src  *url.URL
//		dest *url.URL
//	}
//	tests := []struct {
//		name         string
//		fields       fields
//		args         args
//		wantFileName string
//		wantErr      bool
//	}{
//		{
//			name: "Copy File",
//			fields: fields{
//				// TODO : this is the issue, we are creating fileSystem with blank structs
//				fileSystems: map[string]VFileSystem{"file": &OsFs{}, "": &OsFs{}},
//			},
//			args: args{
//				src:  getParsedUrl("file:///tests/abc.txt"),
//				dest: getParsedUrl("file:///tests/dummy/abc-copy.txt"),
//			},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			fs := &fileSystems{
//				fileSystems: tt.fields.fileSystems,
//			}
//			err := fs.Copy(tt.args.src, tt.args.dest)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//		})
//	}
//}
//
//func TestFileSystems_CopyRaw(t *testing.T) {
//	type fields struct {
//		fileSystems map[string]VFileSystem
//	}
//	type args struct {
//		src  string
//		dest string
//	}
//	tests := []struct {
//		name         string
//		fields       fields
//		args         args
//		wantFileName string
//		wantErr      bool
//	}{
//		{
//			name: "Copy Raw File",
//			fields: fields{
//				fileSystems: map[string]VFileSystem{"file": &OsFs{}, "": &OsFs{}},
//			},
//			args: args{
//				src:  getRawPath("file:///tests/abc.txt"),
//				dest: getRawPath("file:///tests/dummy/abc-copy.txt"),
//			},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			fs := &fileSystems{
//				fileSystems: tt.fields.fileSystems,
//			}
//			err := fs.CopyRaw(tt.args.src, tt.args.dest)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//		})
//	}
//}

// var inputFileSystems = map[string]VFileSystem{"file": &OsFs{}, "": &OsFs{}}
var testManager = GetManager()

func TestFileSystems_CreateRaw(t *testing.T) {
	u := getRawPath("file:///tests/raw-abc.txt")
	fmt.Println(testManager)
	_, err := testManager.CreateRaw(u)
	if err != nil {
		t.Errorf("CreateRaw() error = %v", err)
	}
}

func TestFileSystems_Delete(t *testing.T) {
	u := getParsedUrl("file:///tests/raw-abc.txt")
	err := testManager.Delete(u)
	if err != nil {
		t.Errorf("Delete() error = %v", err)
	}
}

func TestFileSystems_DeleteRaw(t *testing.T) {
	u := getRawPath("file:///tests/raw-abc.txt")
	err := testManager.DeleteRaw(u)
	if err != nil {
		t.Errorf("DeleteRaw() error = %v", err)
	}
}

func TestFileSystems_List(t *testing.T) {
	u := getParsedUrl("file:///tests")
	_, err := testManager.List(u)
	if err != nil {
		t.Errorf("List() error = %v", err)
	}
}

func TestFileSystems_ListRaw(t *testing.T) {
	u := getRawPath("file:///tests")
	_, err := testManager.ListRaw(u)
	if err != nil {
		t.Errorf("ListRaw() error = %v", err)
	}
}

func TestFileSystems_Move(t *testing.T) {
	src := getParsedUrl("file:///tests/abc.txt")
	dest := getParsedUrl("file:///tests/move-abc.txt")
	err := testManager.Move(src, dest)
	if err != nil {
		t.Errorf("Move() error = %v", err)
	}
}

func TestFileSystems_MoveRaw(t *testing.T) {
	src := getRawPath("file:///tests/move-abc.txt")
	dest := getRawPath("file:///tests/move-abc-raw.txt")
	err := testManager.MoveRaw(src, dest)
	if err != nil {
		t.Errorf("MoveRaw() error = %v", err)
	}
}

func TestFileSystems_Open(t *testing.T) {
	u := getParsedUrl("file:///tests/abc.txt")
	_, err := testManager.Open(u)
	if err != nil {
		t.Errorf("Open() error = %v", err)
	}
}

func TestFileSystems_OpenRaw(t *testing.T) {
	u := getRawPath("file:///tests/abc.txt")
	_, err := testManager.OpenRaw(u)
	if err != nil {
		t.Errorf("OpenRaw() error = %v", err)
	}
}

func TestFileSystems_Schemes(t *testing.T) {
	output := testManager.Schemes()
	fmt.Println(output)
}

func TestFileSystems_IsSupported(t *testing.T) {
	var scheme string
	var output bool

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
