package secrets

import (
	"bytes"
	"context"
	"encoding/gob"
	"os"
	"sync"

	"go.nandlabs.io/commons/errutils"
)

const (
	LocalStoreProvider = "LocalStore"
)

//LocalStore will want the credential in a local file.
type LocalStore struct {
	credentials map[string]*Credential
	storeFile   string
	masterKey   string
	mutex       sync.RWMutex
}

func NewLocalStore(storeFile, masterKey string) (ls *LocalStore, err error) {
	var fileContent []byte
	var decryptedContent []byte
	var credentials = make(map[string]*Credential)
	fileContent, err = os.ReadFile(storeFile)
	if err == nil {
		decryptedContent, err = AesDecrypt(fileContent, []byte(masterKey))
		decoder := gob.NewDecoder(bytes.NewReader(decryptedContent))
		err = decoder.Decode(credentials)
		if err == nil {
			ls = &LocalStore{
				credentials: credentials,
				storeFile:   storeFile,
				masterKey:   masterKey,
				mutex:       sync.RWMutex{},
			}
		}
	}
	return
}

func (ls *LocalStore) Get(key string, ctx context.Context) (cred *Credential, err error) {
	ls.mutex.RLock()
	defer ls.mutex.RUnlock()
	if v, ok := ls.credentials[key]; ok {
		cred = v
	} else {
		err = errutils.FmtError("Unable to find a credential with key %s", key)
	}

	return
}

func (ls *LocalStore) Write(key string, credential *Credential, ctx context.Context) (err error) {
	ls.mutex.Lock()
	defer ls.mutex.Unlock()
	ls.credentials[key] = credential
	var b = &bytes.Buffer{}
	var encodedData []byte
	encoder := gob.NewEncoder(b)
	err = encoder.Encode(ls.credentials)
	if err == nil {
		encodedData, err = AesEncrypt([]byte(ls.masterKey), b.Bytes())
		if err == nil {
			err = os.WriteFile(ls.storeFile, encodedData, 0600)
		}
	}
	return
}

func (ls *LocalStore) Provider() string {
	return LocalStoreProvider
}
