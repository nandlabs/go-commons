package vfs

import "testing"

var testBaseFs = &BaseVFS{
	VFileSystem: &OsFs{},
}

func init() {
	u := GetRawPath("file:///test-data-base-fs")
	testBaseFs.MkdirRaw(u)
}

func TestBaseVFS_CreateRaw(t *testing.T) {
	u := GetRawPath("file:///test-data-base-fs/abc-copy-tst.txt")
	_, err := testBaseFs.CreateRaw(u)
	if err != nil {
		t.Errorf("CreateRaw() error = %v", err)
	}
}

func TestBaseVFS_CopyRaw(t *testing.T) {
	src := GetRawPath("file:///test-data-base-fs/abc-copy-tst.txt")
	dest := GetRawPath("file:///test-data-base-fs/copy-file-2.txt")

	err := testBaseFs.CopyRaw(src, dest)
	if err != nil {
		t.Errorf("CopyRaw() error = %v", err)
	}
}

//func TestBaseVFS_ListRaw(t *testing.T) {
//	u := GetRawPath("file:///test-data-base-fs")
//	_, err := testBaseFs.ListRaw(u)
//	if err != nil {
//		t.Errorf("ListRaw() error = %v", err)
//	}
//}

func TestBaseVFS_MkdirAllRaw(t *testing.T) {
	u := GetRawPath("file:///test-data-base-fs/dummy/raw")
	_, err := testBaseFs.MkdirAllRaw(u)
	if err != nil {
		t.Errorf("MkdirAllRaw() error = %v", err)
	}
}

func TestBaseVFS_MoveRaw(t *testing.T) {
	src := GetRawPath("file:///test-data-base-fs/abc-copy-tst-raw.txt")
	dest := GetRawPath("file:///test-data-base-fs/copy-file-2-raw.txt")

	_, err := testBaseFs.CreateRaw(src)
	if err != nil {
		t.Errorf("CreateRaw() error = %v", err)
	}
	err = testBaseFs.MoveRaw(src, dest)
	if err != nil {
		t.Errorf("MoveRaw() error = %v", err)
	}
}

func TestBaseVFS_OpenRaw(t *testing.T) {
	u := GetRawPath("file:///test-data-base-fs/abc-copy-tst.txt")
	_, err := testBaseFs.OpenRaw(u)
	if err != nil {
		t.Errorf("OpenRaw() error = %v", err)
	}
}

func TestBaseVFS_DeleteRaw(t *testing.T) {
	u := GetRawPath("file:///test-data-base-fs")
	testBaseFs.DeleteRaw(u)
}
