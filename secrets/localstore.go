package secrets

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"go.nandlabs.io/commons/errutils"
	"io"
	"os"
	"strings"
	"sync"
)

const (
	LocalStoreProvider = "LocalStore"
)

//LocalStore will store the credential in a local file.
type LocalStore struct {
	credentials map[string]*Credential
	storeFile   string
	masterKey   string
	mutex       sync.RWMutex
}

func NewLocalStore(storeFile, masterKey string) (ls *LocalStore, err error) {
	var fileContent []byte
	var decryptedContent string
	var credentials = make(map[string]*Credential)
	fileContent, err = os.ReadFile(storeFile)
	if err == nil {
		decryptedContent, err = aesDecrypt([]byte(masterKey), string(fileContent))
		decoder := gob.NewDecoder(strings.NewReader(decryptedContent))
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
	var encodedData string
	encoder := gob.NewEncoder(b)
	err = encoder.Encode(ls.credentials)
	if err == nil {

		base64Encoded := base64.StdEncoding.EncodeToString(b.Bytes())
		encodedData, err = aesEncrypt([]byte(ls.masterKey), base64Encoded)
		if err == nil {
			err = os.WriteFile(ls.storeFile, []byte(encodedData), 0600)
		}
	}

	return
}

func aesEncrypt(key []byte, message string) (encoded string, err error) {
	var msgBytes, cipherText []byte
	var block cipher.Block

	msgBytes = []byte(message)
	block, err = aes.NewCipher(key)

	if err == nil {
		cipherText = make([]byte, aes.BlockSize+len(msgBytes))
		//iv is the ciphertext up to the blocksize (16)
		iv := cipherText[:aes.BlockSize]
		if _, err = io.ReadFull(rand.Reader, iv); err == nil {
			stream := cipher.NewCFBEncrypter(block, iv)
			stream.XORKeyStream(cipherText[aes.BlockSize:], msgBytes)
			encoded = base64.RawStdEncoding.EncodeToString(cipherText)
		}
	}

	//Return string encoded in base64
	return
}

func aesDecrypt(key []byte, content string) (decryptedText string, err error) {
	var block cipher.Block
	var cipherText []byte
	cipherText, err = base64.RawStdEncoding.DecodeString(content)
	if err == nil {
		block, err = aes.NewCipher(key)
		if err == nil {
			if len(cipherText) < aes.BlockSize {
				err = errutils.FmtError("The Cipher Length is too Short. Min required %d", aes.BlockSize)
				return
			}
			iv := cipherText[:aes.BlockSize]
			cipherText = cipherText[aes.BlockSize:]
			//Decrypt the message
			stream := cipher.NewCFBDecrypter(block, iv)
			stream.XORKeyStream(cipherText, cipherText)
			decryptedText = string(cipherText)
		}
	}

	return
}
