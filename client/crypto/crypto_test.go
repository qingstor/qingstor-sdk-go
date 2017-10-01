package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var otherAlgorithm Algorithm
var encryptKey string

func TestGetDataCryptoKey(t *testing.T) {
	encryptKey = "AES256Key-32Characters1234567890"
	key, err := getDataCryptoKey(encryptKey, AES256)
	assert.Equal(t, key, encryptKey)
	assert.Nil(t, err)

	key, err = getDataCryptoKey(encryptKey, otherAlgorithm)
	assert.NotNil(t, err)
}

func TestSaveDataCryptoKey(t *testing.T) {
	err := saveDataCryptoKey(nil, encryptKey, AES256)
	assert.Nil(t, err)

	err = saveDataCryptoKey(nil, encryptKey, otherAlgorithm)
	assert.NotNil(t, err)
}
