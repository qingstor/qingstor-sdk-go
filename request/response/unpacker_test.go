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

package response

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/yunify/qingstor-sdk-go/v3/request/data"
	"github.com/yunify/qingstor-sdk-go/v3/request/errors"
)

func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

func IntValue(v *int) int {
	if v != nil {
		return *v
	}
	return 0
}

func Int64Value(v *int64) int64 {
	if v != nil {
		return *v
	}
	return 0
}

func TimeValue(v *time.Time) time.Time {
	if v != nil {
		return *v
	}
	return time.Time{}
}

func TestSimpleUnpackHTTPRequest(t *testing.T) {
	type FakeOutput struct {
		StatusCode *int

		A  *string `location:"elements" json:"a" name:"a"`
		B  *string `location:"elements" json:"b" name:"b"`
		CD *int    `location:"elements" json:"cd" name:"cd"`
		EF *int64  `location:"elements" json:"ef" name:"ef"`
	}

	httpResponse := &http.Response{Header: http.Header{}}
	httpResponse.StatusCode = 200
	httpResponse.Header.Set("Content-Type", "application/json")
	responseString := `{"a": "el_a", "b": "el_b", "cd": 1024, "ef": 2048}`
	httpResponse.Body = ioutil.NopCloser(bytes.NewReader([]byte(responseString)))

	output := &FakeOutput{}
	outputValue := reflect.ValueOf(output)
	u := unpacker{operation: &data.Operation{}, resp: httpResponse, output: &outputValue}
	err := u.unpackResponse()
	assert.Nil(t, err)
	assert.Equal(t, 200, IntValue(output.StatusCode))
	assert.Equal(t, "el_a", StringValue(output.A))
	assert.Equal(t, "el_b", StringValue(output.B))
	assert.Equal(t, 1024, IntValue(output.CD))
	assert.Equal(t, int64(2048), Int64Value(output.EF))
}

func TestUnpackHTTPRequest(t *testing.T) {
	type Bucket struct {
		// Created time of the Bucket
		Created *time.Time `json:"created" name:"created" format:"RFC 822"`
		// QingCloud Zone ID
		Location *string `json:"location" name:"location"`
		// Bucket name
		Name *string `json:"name" name:"name"`
		// URL to access the Bucket
		URL *string `json:"url" name:"url"`
	}

	type ListBucketsOutput struct {
		StatusCode *int `location:"statusCode"`
		Error      *errors.QingStorError
		RequestID  *string `location:"requestID"`

		XTestHeader *string    `json:"X-Test-Header" name:"X-Test-Header" location:"headers"`
		XTestTime   *time.Time `json:"X-Test-Time" name:"X-Test-Time" format:"RFC 822" location:"headers"`

		// Buckets information
		Buckets []*Bucket `json:"buckets" name:"buckets"`
		// Bucket count
		Count *int `json:"count" name:"count"`
	}

	httpResponse := &http.Response{Header: http.Header{
		"X-Test-Header": []string{"test-header"},
		"X-Test-Time":   []string{"Thu, 01 Sep 2016 07:30:00 GMT"},
	}}
	httpResponse.StatusCode = 200
	httpResponse.Header.Set("Content-Type", "application/json")
	responseString := `{
	  "count": 2,
	  "buckets": [
	    {
	      "name": "test-bucket",
	      "location": "pek3a",
	      "url": "https://test-bucket.pek3a.qingstor.com",
	      "created": "2015-07-11T04:45:57Z"
	    },
	    {
	      "name": "test-photos",
	      "location": "pek3a",
	      "url": "https://test-photos.pek3a.qingstor.com",
	      "created": "2015-07-12T09:40:32Z"
	    }
	  ]
	}`
	httpResponse.Body = ioutil.NopCloser(bytes.NewReader([]byte(responseString)))

	output := &ListBucketsOutput{}
	outputValue := reflect.ValueOf(output)
	u := unpacker{operation: &data.Operation{}, resp: httpResponse, output: &outputValue}
	err := u.unpackResponse()
	assert.Nil(t, err)
	assert.Equal(t, "test-header", StringValue(output.XTestHeader))
	assert.Equal(t, time.Date(2016, 9, 1, 7, 30, 0, 0, time.UTC), TimeValue(output.XTestTime))
	assert.Equal(t, 2, IntValue(output.Count))
	assert.Equal(t, "test-bucket", StringValue(output.Buckets[0].Name))
	assert.Equal(t, "pek3a", StringValue(output.Buckets[0].Location))
	assert.Equal(t, time.Date(2015, 7, 12, 9, 40, 32, 0, time.UTC), TimeValue(output.Buckets[1].Created))
}

func TestUnpackHTTPRequestWithError(t *testing.T) {
	type ListBucketsOutput struct {
		StatusCode *int `location:"statusCode"`
		Error      *errors.QingStorError
		RequestID  *string `location:"requestID"`
	}

	httpResponse := &http.Response{Header: http.Header{}}
	httpResponse.StatusCode = 400
	httpResponse.Header.Set("Content-Type", "application/json")
	responseString := `{
	  "code": "bad_request",
	  "message": "Invalid argument(s) or invalid argument value(s)",
	  "request_id": "aa08cf7a43f611e5886952542e6ce14b",
	  "url": "http://docs.qingcloud.com/object_storage/api/bucket/get.html"
	}`
	httpResponse.Body = ioutil.NopCloser(bytes.NewReader([]byte(responseString)))

	output := &ListBucketsOutput{}
	outputValue := reflect.ValueOf(output)
	u := unpacker{operation: &data.Operation{}, resp: httpResponse, output: &outputValue}
	err := u.unpackResponse()
	assert.NotNil(t, err)
	switch e := err.(type) {
	case *errors.QingStorError:
		assert.Equal(t, "bad_request", e.Code)
		assert.Equal(t, "aa08cf7a43f611e5886952542e6ce14b", e.RequestID)
	}
}

func TestUnpackHeadHTTPRequestWithError(t *testing.T) {
	type HeadBucketsOutput struct {
		StatusCode *int `location:"statusCode"`
		Error      *errors.QingStorError
		RequestID  *string `location:"requestID"`
	}

	httpResponse := &http.Response{Header: http.Header{}}
	httpResponse.StatusCode = 404
	httpResponse.Header.Set("Content-Type", "application/json")
	httpResponse.Header.Set("X-QS-Request-ID", "aa08cf7a43f611e5886952542e6ce14b")
	httpResponse.Body = ioutil.NopCloser(strings.NewReader(""))

	output := &HeadBucketsOutput{}
	outputValue := reflect.ValueOf(output)
	u := unpacker{operation: &data.Operation{}, resp: httpResponse, output: &outputValue}
	err := u.unpackResponse()
	assert.NotNil(t, err)
	switch e := err.(type) {
	case *errors.QingStorError:
		assert.Equal(t, "aa08cf7a43f611e5886952542e6ce14b", e.RequestID)
	}
}

func TestUnpackHTTPRequestWithEmptyError(t *testing.T) {
	type ListBucketsOutput struct {
		StatusCode *int `location:"statusCode"`
		Error      *errors.QingStorError
		RequestID  *string `location:"requestID"`
	}

	httpResponse := &http.Response{Header: http.Header{}}
	httpResponse.StatusCode = 400
	httpResponse.Body = ioutil.NopCloser(strings.NewReader(""))
	httpResponse.Header.Set("X-QS-Request-ID", "aa08cf7a43f611e5886952542e6ce14b")

	output := &ListBucketsOutput{}
	outputValue := reflect.ValueOf(output)
	u := unpacker{operation: &data.Operation{}, resp: httpResponse, output: &outputValue}
	err := u.unpackResponse()
	assert.NotNil(t, err)
	switch e := err.(type) {
	case *errors.QingStorError:
		assert.Equal(t, 400, e.StatusCode)
		assert.Equal(t, "aa08cf7a43f611e5886952542e6ce14b", e.RequestID)
	}
}
