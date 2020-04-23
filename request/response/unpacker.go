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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/pengsrc/go-shared/convert"

	"github.com/qingstor/qingstor-sdk-go/v4/logger"
	"github.com/qingstor/qingstor-sdk-go/v4/request/data"
	"github.com/qingstor/qingstor-sdk-go/v4/request/errors"
)

// unpacker is the response unpacker for QingStor service.
type unpacker struct {
	operation *data.Operation
	resp      *http.Response
	output    *reflect.Value
}

// UnpackToOutput unpack the http response with an operation, http response and an output.
func UnpackToOutput(o *data.Operation, r *http.Response, x *reflect.Value) error {
	u := &unpacker{
		operation: o,
		resp:      r,
		output:    x,
	}
	return u.unpackResponse()
}

func (b *unpacker) unpackResponse() error {
	err := b.exposeStatusCode()
	if err != nil {
		return err
	}
	err = b.parseResponseHeaders()
	if err != nil {
		return err
	}
	err = b.parseResponseBody()
	if err != nil {
		return err
	}
	err = b.parseResponseElements()
	if err != nil {
		return err
	}
	err = b.parseError()
	if err != nil {
		return err
	}

	// Close body for every API except GetObject and ImageProcess.
	if b.operation.APIName != "GET Object" && b.operation.APIName != "Image Process" && b.resp.Body != nil {
		err = b.resp.Body.Close()
		if err != nil {
			return err
		}

		b.resp.Body = nil
	}

	return nil
}

func (b *unpacker) exposeStatusCode() error {
	value := b.output.Elem().FieldByName("StatusCode")
	if value.IsValid() {
		switch value.Interface().(type) {
		case *int:
			logger.Infof(nil, fmt.Sprintf(
				"QingStor response status code: [%d] %d",
				convert.StringToTimestamp(b.resp.Header.Get("Date"), convert.RFC822),
				b.resp.StatusCode,
			))
			value.Set(reflect.ValueOf(&b.resp.StatusCode))
		}
	}

	return nil
}

func (b *unpacker) parseResponseHeaders() error {
	logger.Infof(nil, fmt.Sprintf(
		"QingStor response headers: [%d] %s",
		convert.StringToTimestamp(b.resp.Header.Get("Date"), convert.RFC822),
		fmt.Sprint(b.resp.Header),
	))

	if !b.isResponseRight() {
		return nil
	}
	fields := b.output.Elem()
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		fieldTagName := fields.Type().Field(i).Tag.Get("name")
		fieldTagLocation := fields.Type().Field(i).Tag.Get("location")
		if fieldTagName == "X-QS-MetaData" { // handle situation that custom metadata exists
			m := make(map[string]string)
			for k, v := range b.resp.Header {
				kLower := strings.ToLower(k)
				if strings.HasPrefix(kLower, "x-qs-meta-") {
					if len(v) > 0 {
						m[k] = v[0]
					}
				}
			}
			if len(m) > 0 {
				field.Set(reflect.ValueOf(&m))
			}
			continue
		}

		fieldStringValue := b.resp.Header.Get(fieldTagName)

		// Empty value should be ignored.
		if fieldStringValue == "" {
			continue
		}

		if fieldTagName != "" && fieldTagLocation == "headers" {
			switch field.Interface().(type) {
			case *string:
				field.Set(reflect.ValueOf(&fieldStringValue))
			case *int:
				intValue, err := strconv.Atoi(fieldStringValue)
				if err != nil {
					return err
				}
				field.Set(reflect.ValueOf(&intValue))
			case *int64:
				int64Value, err := strconv.ParseInt(fieldStringValue, 10, 64)
				if err != nil {
					return err
				}
				field.Set(reflect.ValueOf(&int64Value))
			case *bool:
			case *time.Time:
				formatString := fields.Type().Field(i).Tag.Get("format")
				format := ""
				switch formatString {
				case "RFC 822":
					format = convert.RFC822
				case "ISO 8601":
					format = convert.ISO8601
				}
				timeValue, err := convert.StringToTime(fieldStringValue, format)
				if err != nil {
					return err
				}
				field.Set(reflect.ValueOf(&timeValue))
			}
		}
	}
	return nil
}

func (b *unpacker) parseResponseBody() error {
	if b.isResponseRight() {
		value := b.output.Elem().FieldByName("Body")
		if value.IsValid() {
			switch value.Type().String() {
			case "string":
				buffer := &bytes.Buffer{}
				buffer.ReadFrom(b.resp.Body)
				b.resp.Body.Close()

				logger.Infof(nil, fmt.Sprintf(
					"QingStor response body string: [%d] %s",
					convert.StringToTimestamp(b.resp.Header.Get("Date"), convert.RFC822),
					string(buffer.Bytes()),
				))

				value.SetString(string(buffer.Bytes()))
			case "io.ReadCloser":
				value.Set(reflect.ValueOf(b.resp.Body))
			}
		}
	}

	return nil
}

func (b *unpacker) parseResponseElements() error {
	if !b.isResponseRight() {
		return nil
	}

	// Do not parse GetObject and ImageProcess's body.
	if b.operation.APIName == "GET Object" ||
		b.operation.APIName == "Image Process" {
		return nil
	}

	if !strings.HasPrefix(b.resp.Header.Get("Content-Type"), "application/json") {
		return nil
	}

	buffer := &bytes.Buffer{}
	buffer.ReadFrom(b.resp.Body)
	b.resp.Body.Close()

	if buffer.Len() == 0 {
		return nil
	}

	logger.Infof(nil, fmt.Sprintf(
		"QingStor response body string: [%d] %s",
		convert.StringToTimestamp(b.resp.Header.Get("Date"), convert.RFC822),
		string(buffer.Bytes()),
	))

	err := json.Unmarshal(buffer.Bytes(), b.output.Interface())
	if err != nil {
		return err
	}

	return nil
}

func (b *unpacker) isResponseRight() bool {
	rightStatusCodes := b.operation.StatusCodes
	if len(rightStatusCodes) == 0 {
		rightStatusCodes = append(rightStatusCodes, 200)
	}

	flag := false
	for _, statusCode := range rightStatusCodes {
		if statusCode == b.resp.StatusCode {
			flag = true
		}
	}

	return flag
}

func (b *unpacker) parseError() error {
	if b.isResponseRight() {
		return nil
	}

	qsError := &errors.QingStorError{
		StatusCode: b.resp.StatusCode,
		RequestID:  b.resp.Header.Get(http.CanonicalHeaderKey("X-QS-Request-ID")),
	}

	// QingStor nginx could refuse user's request directly and only return status code.
	// We should handle this and return qsError directly.
	if b.resp.ContentLength <= 0 {
		return qsError
	}
	if !strings.HasPrefix(b.resp.Header.Get("Content-Type"), "application/json") {
		return qsError
	}

	buffer := &bytes.Buffer{}
	_, err := io.Copy(buffer, b.resp.Body)
	if err != nil {
		logger.Errorf(nil, "Copy from error response body failed for %v", err)
		return err
	}
	err = b.resp.Body.Close()
	if err != nil {
		logger.Errorf(nil, "Close error response body failed for %v", err)
		return err
	}

	if buffer.Len() > 0 && json.Valid(buffer.Bytes()) {
		err := json.Unmarshal(buffer.Bytes(), qsError)
		if err != nil {
			return err
		}
	}
	return qsError
}
