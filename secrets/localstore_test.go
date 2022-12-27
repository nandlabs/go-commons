package secrets

import (
	"context"
	"os"
	"path"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestLocalStore_Get(t *testing.T) {

	type StoreStub struct {
		credentials map[string]*Credential
		storeFile   string
		masterKey   string
		mutex       sync.RWMutex
	}
	type args struct {
		key string
		ctx context.Context
	}
	tests := []struct {
		name     string
		fields   StoreStub
		args     args
		wantCred *Credential
		wantErr  bool
	}{
		{
			name: "testGet",
			fields: StoreStub{
				credentials: map[string]*Credential{
					"test": {
						Value:       []byte("testValue"),
						LastUpdated: time.Now(),
						Version:     "1.0",
						MetaData:    nil,
					},
				},
				storeFile: "test-File",
				masterKey: "master-key-0001",
				mutex:     sync.RWMutex{},
			},
			args: args{
				key: "test",
				ctx: nil,
			},
			wantCred: &Credential{
				Value:       []byte("testValue"),
				LastUpdated: time.Now(),
				Version:     "1.0",
				MetaData:    nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := &localStore{
				credentials: tt.fields.credentials,
				storeFile:   tt.fields.storeFile,
				masterKey:   tt.fields.masterKey,
				mutex:       tt.fields.mutex,
			}
			gotCred, err := ls.Get(tt.args.key, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCred.Str() != tt.wantCred.Str() {
				t.Errorf("Get() gotCred = %v, want %v", gotCred.Str(), tt.wantCred.Str())
			}
		})
	}
}

func TestLocalStore_Provider(t *testing.T) {
	type fields struct {
		credentials map[string]*Credential
		storeFile   string
		masterKey   string
		mutex       sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{

		{
			name: "local-provider-test",
			fields: fields{
				credentials: nil,
				storeFile:   "",
				masterKey:   "",
				mutex:       sync.RWMutex{},
			},
			want: LocalStoreProvider,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := &localStore{
				credentials: tt.fields.credentials,
				storeFile:   tt.fields.storeFile,
				masterKey:   tt.fields.masterKey,
				mutex:       tt.fields.mutex,
			}
			if got := ls.Provider(); got != tt.want {
				t.Errorf("Provider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalStore_Write(t *testing.T) {
	dir, err := os.Getwd()
	type fields struct {
		credentials map[string]*Credential
		storeFile   string
		masterKey   string
		mutex       sync.RWMutex
	}
	type args struct {
		key        string
		credential *Credential
		ctx        context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{

		{
			name: "write-test",
			fields: fields{
				credentials: make(map[string]*Credential),
				storeFile:   path.Join(dir, "testdata", "test-store.dat"),
				masterKey:   "thisisamasterkey",
				mutex:       sync.RWMutex{},
			},
			args: args{
				key: "test-key-01",
				credential: &Credential{
					Value:       []byte("test-value"),
					LastUpdated: time.Now(),
					Version:     "1.0",
					MetaData:    nil,
				},
				ctx: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls := &localStore{
				credentials: tt.fields.credentials,
				storeFile:   tt.fields.storeFile,
				masterKey:   tt.fields.masterKey,
				mutex:       tt.fields.mutex,
			}
			if err = ls.Write(tt.args.key, tt.args.credential, tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewLocalStore(t *testing.T) {
	type args struct {
		storeFile string
		masterKey string
	}
	tests := []struct {
		name    string
		args    args
		wantLs  *localStore
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLs, err := NewLocalStore(tt.args.storeFile, tt.args.masterKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLocalStore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLs, tt.wantLs) {
				t.Errorf("NewLocalStore() gotLs = %v, want %v", gotLs, tt.wantLs)
			}
		})
	}
}
