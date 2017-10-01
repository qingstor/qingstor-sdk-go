package crypto

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var encryptedData []byte
var decryptedData []byte

func init() {
	encryptedData = []byte{238, 109, 53, 209, 187, 241, 63, 39, 188,
		13, 42, 180, 227, 15, 218, 148, 182, 124, 21, 73, 92,
		216, 214, 190, 118, 173, 125, 210, 240, 58, 1, 6}

	decryptedData = []byte{119, 119, 119, 46, 113, 105, 110, 103, 99,
		108, 111, 117, 100, 46, 99, 111}
	encryptedFile, err := os.Create("encryptedTestFile.txt")
	if err != nil {
		panic(err)
	}
	defer encryptedFile.Close()
	encryptedFile.Write(encryptedData)

	decryptedFile, err := os.Create("decryptedTestFile.txt")
	if err != nil {
		panic(err)
	}
	defer decryptedFile.Close()
	decryptedFile.Write(decryptedData)
}

func TestReadData(t *testing.T) {

	key := "AES256Key-32Characters1234567890"
	decryptedFile, err := os.Open("decryptedTestFile.txt")
	if err != nil {
		panic(err)
	}
	defer decryptedFile.Close()
	encrypter, err := NewEncryptReader(decryptedFile, key)
	if err != nil {
		panic(err)
	}
	enData, err := ioutil.ReadAll(encrypter)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, enData, encryptedData)

	encryptedFile, err := os.Open("encryptedTestFile.txt")
	if err != nil {
		panic(err)
	}
	defer encryptedFile.Close()
	decrypter, err := NewDecryptReader(encryptedFile, key)
	deData, err := ioutil.ReadAll(decrypter)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, deData[:len(deData)-16], decryptedData)
	removeTestFiles()
}

func TestPKCS7(t *testing.T) {
	testdata := []byte("www.qingcloud.com")
	blockSize := 16
	src := PKCS7Padding(testdata, blockSize)
	src = PKCS7UnPadding(src)

	assert.Equal(t, testdata, src)
}

func removeTestFiles() {
	err := os.Remove("encryptedTestFile.txt")
	if err != nil {
		panic(err)
	}

	err = os.Remove("decryptedTestFile.txt")
	if err != nil {
		panic(err)
	}
}
