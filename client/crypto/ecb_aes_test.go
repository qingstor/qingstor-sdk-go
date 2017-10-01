package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"testing"

	"github.com/stretchr/testify/assert"
)

var block cipher.Block

func init() {
	key := "AES256Key-32Characters1234567890"
	var err error
	block, err = aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
}

func TestCryptData(t *testing.T) {

	testdata := []byte("www.qingcloud.com")
	// encrypt data
	src := PKCS7Padding(testdata, 16)
	encryptedData := make([]byte, len(src))
	encrypter := NewECBEncrypter(block)
	encrypter.CryptBlocks(encryptedData, src)

	//decrypt data
	decryptedData := make([]byte, len(encryptedData))
	decrypter := NewECBDecrypter(block)
	decrypter.CryptBlocks(decryptedData, encryptedData)
	decryptedData = PKCS7UnPadding(decryptedData)

	// check result
	assert.Equal(t, testdata, decryptedData)
}
