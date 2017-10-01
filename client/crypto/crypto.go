package crypto

import (
	"errors"

	"github.com/yunify/qingstor-sdk-go/service"
)

// Algorithm is the set of all unsigned 64-bit integers.
type Algorithm uint8

var zero int64

const (
	// AES256 represents aes56 algorithm
	AES256 Algorithm = 1 << iota
)

// Client is advanced client of crypto
type Client struct {
	*service.Bucket
	key       string
	algorithm Algorithm
}

// NewClient uses the bucket, key and algorithm to initialize the advanced
// client of crypto.
func NewClient(bucket *service.Bucket, key string,
	algorithm Algorithm) *Client {
	return &Client{
		Bucket:    bucket,
		key:       key,
		algorithm: algorithm,
	}
}

// PutObject does Upload the encrypted object.
func (client *Client) PutObject(objectKey string,
	input *service.PutObjectInput) (*service.PutObjectOutput, error) {
	key, err := getDataCryptoKey(client.key, client.algorithm)
	if err != nil {
		return nil, err
	}
	input.ContentLength = &zero
	input.Body, err = NewEncryptReader(input.Body, key)
	if err != nil {
		return nil, err
	}
	saveDataCryptoKey(input, key, client.algorithm)
	return client.Bucket.PutObject(objectKey, input)
}

// GetObject does Retrieve the decrypted object.
func (client *Client) GetObject(objectKey string,
	input *service.GetObjectInput) (*service.GetObjectOutput, error) {
	output, err := client.Bucket.GetObject(objectKey, input)
	key, err := getDataCryptoKey(client.key, client.algorithm)
	output.Body, err = NewDecryptReader(output.Body, key)
	if err != nil {
		return nil, err
	}
	return output, err
}

func getDataCryptoKey(key string, algorithm Algorithm) (string, error) {
	switch algorithm {
	case AES256:
		return key, nil
	default:
		return "", errors.New("Not support for algorithm")
	}
}

func saveDataCryptoKey(input interface{}, key string, algorithm Algorithm) error {
	switch algorithm {
	case AES256:
		return nil
	default:
		return errors.New("Not support for algorithm")
	}
}
