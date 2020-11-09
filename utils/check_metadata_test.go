package utils

import (
	"errors"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	sdkErr "github.com/qingstor/qingstor-sdk-go/v4/request/errors"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))

func TestMetaData(t *testing.T) {
	t.Run("normal metadata", func(t *testing.T) {
		m := map[string]string{
			"x-qs-meta-a": "a",
		}
		err := IsMetaDataValid(&m)
		assert.Nil(t, err)
	})

	t.Run("non prefix", func(t *testing.T) {
		m := map[string]string{
			"a": "a",
		}
		err := IsMetaDataValid(&m)
		assert.NotNil(t, err)

		ae := sdkErr.ParameterValueNotAllowedError{}
		assert.True(t, errors.As(err, &ae))
	})

	t.Run("invalid key char", func(t *testing.T) {
		var s strings.Builder
		s.WriteString(metadataPrefix)
		s.WriteByte(',')
		m := map[string]string{
			s.String(): "a",
		}
		err := IsMetaDataValid(&m)
		assert.NotNil(t, err)

		ae := sdkErr.ParameterValueNotAllowedError{}
		assert.True(t, errors.As(err, &ae))
	})

	t.Run("invalid value char", func(t *testing.T) {
		var s strings.Builder
		s.WriteByte(byte(32))
		m := map[string]string{
			metadataPrefix + "a": s.String(),
		}
		err := IsMetaDataValid(&m)
		assert.NotNil(t, err)

		ae := sdkErr.ParameterValueNotAllowedError{}
		assert.True(t, errors.As(err, &ae))
	})

	t.Run("key too long", func(t *testing.T) {
		key := randomKeyString(513)
		m := map[string]string{
			metadataPrefix + key: "a",
		}
		err := IsMetaDataValid(&m)
		assert.NotNil(t, err)

		ae := sdkErr.ParameterValueNotAllowedError{}
		assert.True(t, errors.As(err, &ae))
	})

	t.Run("metadata oversize", func(t *testing.T) {
		m := map[string]string{
			metadataPrefix + randomKeyString(512): randomPrintableString(512),
			metadataPrefix + randomKeyString(512): randomPrintableString(513),
		}

		err := IsMetaDataValid(&m)
		assert.NotNil(t, err)

		ae := sdkErr.ParameterValueNotAllowedError{}
		assert.True(t, errors.As(err, &ae))
		assert.Equal(t, ae.ParameterValue, "2049")
	})
}

func randomPrintableString(size int) string {
	var s strings.Builder
	for i := 0; i < size; i++ {
		printableChar := r.Int31n(94) + 33 // ascii in [33, 127)
		s.WriteByte(byte(printableChar))
	}
	return s.String()
}

func randomKeyString(size int) string {
	var s strings.Builder
	for i := 0; i < size; i++ {
		pos := r.Intn(len(validKeyChars))
		s.WriteByte(validKeyChars[pos])
	}
	return s.String()
}
