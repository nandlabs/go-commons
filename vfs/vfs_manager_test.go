package vfs

import (
	"net/url"
	"os"
	"reflect"
	"testing"
)

func getParsedUrl(input string) (u *url.URL) {
	u, _ = url.Parse(input)
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
				u: getParsedUrl("file://./tests"),
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
				u: getParsedUrl("file://./tests/dummy"),
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
				u: getParsedUrl("file://./tests/abc.txt"),
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
				u: getParsedUrl("file://./tests/dummy/dummy.txt"),
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

func TestFileSystems_Copy(t *testing.T) {
	type fields struct {
		fileSystems map[string]VFileSystem
	}
	type args struct {
		src  *url.URL
		dest *url.URL
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantFileName string
		wantErr      bool
	}{
		{
			name: "Copy File",
			fields: fields{
				fileSystems: map[string]VFileSystem{"file": &OsFs{}, "": &OsFs{}},
			},
			args: args{
				src:  getParsedUrl("file://./tests/abc.txt"),
				dest: getParsedUrl("file://./tests/dummy/abc-copy.txt"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fileSystems{
				fileSystems: tt.fields.fileSystems,
			}
			err := fs.Copy(tt.args.src, tt.args.dest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}

}
