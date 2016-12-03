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

package unpacker

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yunify/qingstor-sdk-go/request/data"
	"reflect"
)

func TestBaseUnpacker_UnpackHTTPRequest(t *testing.T) {
	type FakeOutput struct {
		StatusCode int

		A  string `location:"data" json:"a" name:"a"`
		B  string `location:"data" json:"b" name:"b"`
		CD int    `location:"data" json:"cd" name:"cd"`
	}

	httpResponse := &http.Response{Header: http.Header{}}
	httpResponse.StatusCode = 200
	httpResponse.Header.Set("Content-Type", "application/json")
	responseString := `{"a": "el_a", "b": "el_b", "cd": 1024}`
	httpResponse.Body = ioutil.NopCloser(bytes.NewReader([]byte(responseString)))

	output := &FakeOutput{}
	outputValue := reflect.ValueOf(output)
	unpacker := BaseUnpacker{}
	err := unpacker.UnpackHTTPRequest(&data.Operation{}, httpResponse, &outputValue)
	assert.Nil(t, err)
	assert.Equal(t, "el_a", output.A)
	assert.Equal(t, "el_b", output.B)
	assert.Equal(t, 1024, output.CD)
}
