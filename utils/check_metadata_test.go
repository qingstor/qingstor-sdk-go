package utils

import (
	"testing"
)

func TestMetaData(t *testing.T) {
	var metaData = []struct {
		metaData map[string]string
	}{
		{map[string]string{
			"x-qs-meta-a": "a",
		}},
		{map[string]string{
			"x-qs-meta-aa": "a,,",
		}},
	}

	for _, tt := range metaData {
		flag := IsMetaDataValid(&tt.metaData)
		if flag != nil {
			t.Errorf("%v", tt)
		}
	}
}
