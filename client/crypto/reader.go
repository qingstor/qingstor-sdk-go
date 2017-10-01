package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"io"
	"strings"
)

const blockSize int = 16

// EncryptReader embed io.Reader and have cipher.Block.
type EncryptReader struct {
	io.Reader
	block    cipher.Block
	isPadded bool
}

// DecryptReader embed io.ReadCloser and have cipher.Block.
type DecryptReader struct {
	io.ReadCloser
	block cipher.Block
}

// NewEncryptReader return *EncryptReader to encrypt data.
func NewEncryptReader(body io.Reader, key string) (*EncryptReader, error) {
	if body == nil || strings.Trim(key, " ") == "" {
		return nil, errors.New("body or key is null")
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	return &EncryptReader{
		Reader: body,
		block:  block,
	}, nil
}

// NewDecryptReader return *DecryptReader to decrypt data.
func NewDecryptReader(body io.ReadCloser, key string) (*DecryptReader, error) {
	if body == nil || strings.Trim(key, " ") == "" {
		return nil, errors.New("body or key is null")
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	return &DecryptReader{
		ReadCloser: body,
		block:      block,
	}, nil
}

// Read reads up to len(b) bytes from the EncryptReader.
// It returns the number of bytes read and any error encountered.
// At end of EncryptReader, Read returns 0, io.EOF.
func (encryptReader *EncryptReader) Read(p []byte) (n int, err error) {
	temp := make([]byte, blockSize)
	n, err = encryptReader.Reader.Read(temp)
	if err != nil {
		if !(err == io.EOF && !encryptReader.isPadded) {
			return n, err
		}
	}

	for err == nil && n < blockSize {
		tmp := make([]byte, blockSize-n)
		tmpN := 0
		tmpN, err = encryptReader.Reader.Read(tmp)
		if err != nil {
			if !(err == io.EOF && n != 0) {
				return n, err
			}
		}

		j := 0
		for i := n; i < n+tmpN; i++ {
			temp[i] = tmp[j]
			j++
		}
		n += tmpN
	}

	if n < blockSize {
		temp = PKCS7Padding(temp[:n], blockSize)
		encryptReader.isPadded = true
		n = len(temp)
	}

	encrypter := NewECBEncrypter(encryptReader.block)
	encrypter.CryptBlocks(p, temp)
	return n, err
}

// Read reads up to len(b) bytes from the DecryptReader.
// It returns the number of bytes read and any error encountered.
// At end of DecryptReader, Read returns 0, io.EOF.
func (decryptReader *DecryptReader) Read(p []byte) (n int, err error) {

	temp := make([]byte, blockSize)
	n, err = decryptReader.ReadCloser.Read(temp)
	if err != nil {
		if !(err == io.EOF && n != 0) {
			return n, err
		}
	}

	for err == nil && n < blockSize {
		tmp := make([]byte, blockSize-n)
		tmpN := 0
		tmpN, err = decryptReader.ReadCloser.Read(tmp)
		if err != nil {
			if !(err == io.EOF && n != 0) {
				return n, err
			}
		}

		j := 0
		for i := n; i < n+tmpN; i++ {
			temp[i] = tmp[j]
			j++
		}
		n += tmpN
	}
	decrypter := NewECBDecrypter(decryptReader.block)
	decryptedData := make([]byte, n)
	decrypter.CryptBlocks(decryptedData, temp[:n])
	if err == io.EOF {
		decryptedData = PKCS7UnPadding(decryptedData)
	}
	copy(p, decryptedData)
	return len(decryptedData), err
}

// Close closes the DecryptReader, rendering it unusable for I/O.
// It returns an error, if any.
func (decryptReader *DecryptReader) Close() error {
	return decryptReader.ReadCloser.Close()
}

// PKCS7Padding is in whole bytes.
// The value of each added byte is the number of bytes that are added,
// i.e. N bytes, each of value N are added.
// The number of bytes added will depend on the block boundary to which
// the message needs to be extended.
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding remove the filled data
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
