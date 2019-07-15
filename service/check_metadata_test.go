package service

import (
	"testing"
)


func TestMetaData(t *testing.T) {
	var metaData = []struct{
		metaData map[string]string
	}{
		{map[string]string {
			"x":"a",
		}},
		{map[string]string{
			"x-":"a",
		}},
		{map[string]string{
			"x-qs-":"a",
		}},
		{map[string]string{
			"x-abc":"a,,",
		}},
		{map[string]string{
			"x-qs-meta":"a",
		}},
		{map[string]string{
			"x-qs-metas":"a",
		}},
		{map[string]string{
			"x-qs-meta-a":"a",
		}},
		{map[string]string{
			"x-qs-meta-,,":"a",
		}},
		{map[string]string{
			"x-qs-metas":"a-数据",
		}},
		{map[string]string{
			"x-qs-meta-aa":"a,,",
		}},

	}


	for _,tt := range metaData{
		flag := IsMetaDataValid(&tt.metaData)
		if flag!=nil {
			t.Errorf("%v",tt)
		}
	}
}
