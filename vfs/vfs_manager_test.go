package vfs

import (
	"fmt"
	"net/url"
	"reflect"
	"sync"
	"testing"
)

func TestGetManager(t *testing.T) {
	tests := []struct {
		name string
		want VFileSystem
	}{
		// TODO: Add test cases.
		{
			name: "localFs",
			want: &OsFs{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetManager(); !reflect.DeepEqual(got, tt.want) {
				fmt.Println(got)
				t.Errorf("GetManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	type args struct {
		fs VFileSystem
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Register(tt.args.fs)
		})
	}
}

func Test_fileSystems_Copy(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		fileSystems map[string]VFileSystem
	}
	type args struct {
		src *url.URL
		dst *url.URL
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				mutex:       tt.fields.mutex,
				fileSystems: tt.fields.fileSystems,
			}
			if err := fs.Copy(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fileSystems_CopyRaw(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		fileSystems map[string]VFileSystem
	}
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				mutex:       tt.fields.mutex,
				fileSystems: tt.fields.fileSystems,
			}
			if err := fs.CopyRaw(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("CopyRaw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fileSystems_Create(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		fileSystems map[string]VFileSystem
	}
	type args struct {
		u *url.URL
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantFile VFile
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				mutex:       tt.fields.mutex,
				fileSystems: tt.fields.fileSystems,
			}
			gotFile, err := fs.Create(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFile, tt.wantFile) {
				t.Errorf("Create() gotFile = %v, want %v", gotFile, tt.wantFile)
			}
		})
	}
}

func Test_fileSystems_CreateRaw(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		fileSystems map[string]VFileSystem
	}
	type args struct {
		raw string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantFile VFile
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				mutex:       tt.fields.mutex,
				fileSystems: tt.fields.fileSystems,
			}
			gotFile, err := fs.CreateRaw(tt.args.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateRaw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFile, tt.wantFile) {
				t.Errorf("CreateRaw() gotFile = %v, want %v", gotFile, tt.wantFile)
			}
		})
	}
}

func Test_fileSystems_Delete(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		fileSystems map[string]VFileSystem
	}
	type args struct {
		u *url.URL
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				mutex:       tt.fields.mutex,
				fileSystems: tt.fields.fileSystems,
			}
			if err := fs.Delete(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fileSystems_DeleteRaw(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		fileSystems map[string]VFileSystem
	}
	type args struct {
		raw string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				mutex:       tt.fields.mutex,
				fileSystems: tt.fields.fileSystems,
			}
			if err := fs.DeleteRaw(tt.args.raw); (err != nil) != tt.wantErr {
				t.Errorf("DeleteRaw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fileSystems_IsSupported(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		fileSystems map[string]VFileSystem
	}
	type args struct {
		scheme string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		wantSupported bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				mutex:       tt.fields.mutex,
				fileSystems: tt.fields.fileSystems,
			}
			if gotSupported := fs.IsSupported(tt.args.scheme); gotSupported != tt.wantSupported {
				t.Errorf("IsSupported() = %v, want %v", gotSupported, tt.wantSupported)
			}
		})
	}
}

func Test_fileSystems_List(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		fileSystems map[string]VFileSystem
	}
	type args struct {
		url *url.URL
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantFiles []VFile
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				mutex:       tt.fields.mutex,
				fileSystems: tt.fields.fileSystems,
			}
			gotFiles, err := fs.List(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFiles, tt.wantFiles) {
				t.Errorf("List() gotFiles = %v, want %v", gotFiles, tt.wantFiles)
			}
		})
	}
}

func Test_fileSystems_ListRaw(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		fileSystems map[string]VFileSystem
	}
	type args struct {
		raw string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantFiles []VFile
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				mutex:       tt.fields.mutex,
				fileSystems: tt.fields.fileSystems,
			}
			gotFiles, err := fs.ListRaw(tt.args.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListRaw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFiles, tt.wantFiles) {
				t.Errorf("ListRaw() gotFiles = %v, want %v", gotFiles, tt.wantFiles)
			}
		})
	}
}

func Test_fileSystems_Mkdir(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		fileSystems map[string]VFileSystem
	}
	type args struct {
		url *url.URL
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantFile VFile
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				mutex:       tt.fields.mutex,
				fileSystems: tt.fields.fileSystems,
			}
			gotFile, err := fs.Mkdir(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mkdir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFile, tt.wantFile) {
				t.Errorf("Mkdir() gotFile = %v, want %v", gotFile, tt.wantFile)
			}
		})
	}
}

func Test_fileSystems_MkdirAll(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		fileSystems map[string]VFileSystem
	}
	type args struct {
		url *url.URL
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantFile VFile
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				mutex:       tt.fields.mutex,
				fileSystems: tt.fields.fileSystems,
			}
			gotFile, err := fs.MkdirAll(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("MkdirAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFile, tt.wantFile) {
				t.Errorf("MkdirAll() gotFile = %v, want %v", gotFile, tt.wantFile)
			}
		})
	}
}

func Test_fileSystems_MkdirAllRaw(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		fileSystems map[string]VFileSystem
	}
	type args struct {
		raw string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantFile VFile
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				mutex:       tt.fields.mutex,
				fileSystems: tt.fields.fileSystems,
			}
			gotFile, err := fs.MkdirAllRaw(tt.args.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("MkdirAllRaw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFile, tt.wantFile) {
				t.Errorf("MkdirAllRaw() gotFile = %v, want %v", gotFile, tt.wantFile)
			}
		})
	}
}

func Test_fileSystems_MkdirRaw(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		fileSystems map[string]VFileSystem
	}
	type args struct {
		raw string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantFile VFile
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				mutex:       tt.fields.mutex,
				fileSystems: tt.fields.fileSystems,
			}
			gotFile, err := fs.MkdirRaw(tt.args.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("MkdirRaw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFile, tt.wantFile) {
				t.Errorf("MkdirRaw() gotFile = %v, want %v", gotFile, tt.wantFile)
			}
		})
	}
}

func Test_fileSystems_Move(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		fileSystems map[string]VFileSystem
	}
	type args struct {
		src *url.URL
		dst *url.URL
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				mutex:       tt.fields.mutex,
				fileSystems: tt.fields.fileSystems,
			}
			if err := fs.Move(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("Move() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fileSystems_MoveRaw(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		fileSystems map[string]VFileSystem
	}
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				mutex:       tt.fields.mutex,
				fileSystems: tt.fields.fileSystems,
			}
			if err := fs.MoveRaw(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("MoveRaw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fileSystems_Open(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		fileSystems map[string]VFileSystem
	}
	type args struct {
		url *url.URL
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantFile VFile
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				mutex:       tt.fields.mutex,
				fileSystems: tt.fields.fileSystems,
			}
			gotFile, err := fs.Open(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFile, tt.wantFile) {
				t.Errorf("Open() gotFile = %v, want %v", gotFile, tt.wantFile)
			}
		})
	}
}

func Test_fileSystems_OpenRaw(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		fileSystems map[string]VFileSystem
	}
	type args struct {
		raw string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantFile VFile
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				mutex:       tt.fields.mutex,
				fileSystems: tt.fields.fileSystems,
			}
			gotFile, err := fs.OpenRaw(tt.args.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenRaw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFile, tt.wantFile) {
				t.Errorf("OpenRaw() gotFile = %v, want %v", gotFile, tt.wantFile)
			}
		})
	}
}

func Test_fileSystems_Register(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		fileSystems map[string]VFileSystem
	}
	type args struct {
		vfs VFileSystem
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				mutex:       tt.fields.mutex,
				fileSystems: tt.fields.fileSystems,
			}
			fs.Register(tt.args.vfs)
		})
	}
}

func Test_fileSystems_Schemes(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		fileSystems map[string]VFileSystem
	}
	tests := []struct {
		name        string
		fields      fields
		wantSchemes []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				mutex:       tt.fields.mutex,
				fileSystems: tt.fields.fileSystems,
			}
			if gotSchemes := fs.Schemes(); !reflect.DeepEqual(gotSchemes, tt.wantSchemes) {
				t.Errorf("Schemes() = %v, want %v", gotSchemes, tt.wantSchemes)
			}
		})
	}
}

func Test_fileSystems_Walk(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		fileSystems map[string]VFileSystem
	}
	type args struct {
		url *url.URL
		fn  WalkFn
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				mutex:       tt.fields.mutex,
				fileSystems: tt.fields.fileSystems,
			}
			if err := fs.Walk(tt.args.url, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("Walk() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fileSystems_WalkRaw(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		fileSystems map[string]VFileSystem
	}
	type args struct {
		raw string
		fn  WalkFn
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				mutex:       tt.fields.mutex,
				fileSystems: tt.fields.fileSystems,
			}
			if err := fs.WalkRaw(tt.args.raw, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("WalkRaw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_fileSystems_getFsFor(t *testing.T) {
	type fields struct {
		mutex       sync.Mutex
		fileSystems map[string]VFileSystem
	}
	type args struct {
		src *url.URL
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantVfs VFileSystem
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				mutex:       tt.fields.mutex,
				fileSystems: tt.fields.fileSystems,
			}
			gotVfs, err := fs.getFsFor(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFsFor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVfs, tt.wantVfs) {
				t.Errorf("getFsFor() gotVfs = %v, want %v", gotVfs, tt.wantVfs)
			}
		})
	}
}
