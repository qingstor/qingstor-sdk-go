// +-------------------------------------------------------------------------
// | Copyright (C) 2016 Yunify, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------

package signer

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yunify/qingstor-sdk-go/utils"
)

func TestQingStorSignerWriteSignature(t *testing.T) {
	url := "https://qingstor.com/?acl&upload_id=fde133b5f6d932cd9c79bac3c7318da1&part_number=0&other=abc"
	httpRequest, err := http.NewRequest("GET", url, nil)
	httpRequest.Header.Set("Date", utils.TimeToString(time.Time{}, "RFC 822"))
	httpRequest.Header.Set("X-QS-Test-2", "Test 2")
	httpRequest.Header.Set("X-QS-Test-1", "Test 1")
	assert.Nil(t, err)

	s := QingStorSigner{
		AccessKeyID:     "ENV_ACCESS_KEY_ID",
		SecretAccessKey: "ENV_SECRET_ACCESS_KEY",
	}

	err = s.WriteSignature(httpRequest)
	assert.Nil(t, err)

	signature := "QS ENV_ACCESS_KEY_ID:bvglZF9iMOv1RaCTxPYWxexmt1UN2m5WKngYnhDEp2c="
	assert.Equal(t, signature, httpRequest.Header.Get("Authorization"))
}

func TestQingStorSignerWriteSignatureChinese(t *testing.T) {
	url := "https://zone.qingstor.com/bucket-name/中文"
	httpRequest, err := http.NewRequest("GET", url, nil)
	httpRequest.Header.Set("Date", utils.TimeToString(time.Time{}, "RFC 822"))
	assert.Nil(t, err)

	s := QingStorSigner{
		AccessKeyID:     "ENV_ACCESS_KEY_ID",
		SecretAccessKey: "ENV_SECRET_ACCESS_KEY",
	}

	err = s.WriteSignature(httpRequest)
	assert.Nil(t, err)

	signature := "QS ENV_ACCESS_KEY_ID:XsTXX50kzqBf92zLG1aIUIJmZ0hqIHoaHgkumwnV3fs="
	assert.Equal(t, signature, httpRequest.Header.Get("Authorization"))
}

func TestQingStorSignerWriteQuerySignature(t *testing.T) {
	url := "https://qingstor.com/?acl&upload_id=fde133b5f6d932cd9c79bac3c7318da1&part_number=0"
	httpRequest, err := http.NewRequest("GET", url, nil)
	httpRequest.Header.Set("Date", utils.TimeToString(time.Time{}, "RFC 822"))
	httpRequest.Header.Set("X-QS-Test-2", "Test 2")
	httpRequest.Header.Set("X-QS-Test-1", "Test 1")
	assert.Nil(t, err)

	s := QingStorSigner{
		AccessKeyID:     "ENV_ACCESS_KEY_ID",
		SecretAccessKey: "ENV_SECRET_ACCESS_KEY",
	}

	err = s.WriteQuerySignature(httpRequest, utils.StringToUnixInt(httpRequest.Header.Get("Date"), "RFC 822"))
	assert.Nil(t, err)

	targetURL := "https://qingstor.com/?acl&upload_id=fde133b5f6d932cd9c79bac3c7318da1&part_number=0&access_key_id=ENV_ACCESS_KEY_ID&expires=-62135596800&signature=gTdB/cmD6rjv8CbFRDfFbHc64q442rYNAp99Hm7fBl4="
	assert.Equal(t, httpRequest.URL.String(), targetURL)
}
