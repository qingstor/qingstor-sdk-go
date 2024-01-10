package utils

import (
	"strconv"
	"strings"

	"github.com/qingstor/qingstor-sdk-go/v4/request/errors"
)

var metadataPrefix = "x-qs-meta-"

// validKeyChars contains all valid key char: (a-z, A-Z, -, _, .)
var validKeyChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_."

// IsMetaDataValid check whether the metadata-KV follows rule in API document
// https://docsv4.qingcloud.com/user_guide/storage/object_storage/api/metadata/#%E5%AF%B9%E8%87%AA%E5%AE%9A%E4%B9%89%E5%85%83%E6%95%B0%E6%8D%AE%E7%9A%84%E9%99%90%E5%88%B6
func IsMetaDataValid(XQSMetaData *map[string]string) error {
	metadataSize := 0

	for k, v := range *XQSMetaData {
		if !strings.HasPrefix(strings.ToLower(k), metadataPrefix) {
			return newMetaDataInvalidError(k, v)
		}

		// check key (a-z, A-Z, -, _, .)
		for _, c := range k {
			if !strings.ContainsRune(validKeyChars, c) {
				return newMetaDataInvalidError(k, v)
			}
		}

		// check value (must printable)
		for _, c := range v {
			if c <= 32 || c > 126 {
				return newMetaDataInvalidError(k, v)
			}
		}

		// check key length, without prefix "x-qs-meta-"
		keyLen := len(strings.TrimPrefix(k, metadataPrefix))
		if keyLen > 512 {
			return newMetaDataInvalidError(k, v)
		}

		metadataSize = metadataSize + keyLen + len(v)
	}

	// check metadata size, must le 2048
	if metadataSize > 2048 {
		return newMetaDataTooLargeError(metadataSize)
	}
	return nil
}

func newMetaDataInvalidError(key, value string) error {
	return errors.ParameterValueNotAllowedError{
		ParameterName:  "XQSMetaData",
		ParameterValue: "map[" + key + "]=" + value,
		AllowedValues:  []string{"https://docsv4.qingcloud.com/user_guide/storage/object_storage/api/metadata/"},
	}
}

func newMetaDataTooLargeError(size int) error {
	return errors.ParameterValueNotAllowedError{
		ParameterName:  "XQSMetaData size too large",
		ParameterValue: strconv.Itoa(size),
		AllowedValues:  []string{"https://docsv4.qingcloud.com/user_guide/storage/object_storage/api/metadata/"},
	}
}
