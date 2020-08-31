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
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/pengsrc/go-shared/convert"
	"github.com/qingstor/log"

	"github.com/qingstor/qingstor-sdk-go/v4/request/data"
	"github.com/qingstor/qingstor-sdk-go/v4/request/errors"
)

// unpacker is the response unpacker for QingStor service.
type unpacker struct {
	operation *data.Operation
	resp      *http.Response
	output    *reflect.Value
}

// UnpackToOutputWithContext unpack the http response with an operation, http response and an output with given ctx.
func UnpackToOutputWithContext(ctx context.Context, o *data.Operation, r *http.Response, x *reflect.Value) error {
	if ctx == nil {
		ctx = context.Background()
	}

	u := &unpacker{
		operation: o,
		resp:      r,
		output:    x,
	}
	return u.unpackResponse(ctx)
}

// UnpackToOutput unpack the http response with an operation, http response and an output with given ctx.
// Deprecated: Use UnpackToOutputWithContext instead
func UnpackToOutput(o *data.Operation, r *http.Response, x *reflect.Value) error {
	return UnpackToOutputWithContext(context.Background(), o, r, x)
}

func (b *unpacker) unpackResponse(ctx context.Context) error {
	err := b.exposeStatusCode(ctx)
	if err != nil {
		return err
	}
	err = b.parseResponseHeaders(ctx)
	if err != nil {
		return err
	}
	err = b.parseResponseBody(ctx)
	if err != nil {
		return err
	}
	err = b.parseResponseElements(ctx)
	if err != nil {
		return err
	}
	err = b.parseError(ctx)
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

func (b *unpacker) exposeStatusCode(ctx context.Context) error {
	logger := log.FromContext(ctx)
	value := b.output.Elem().FieldByName("StatusCode")
	if value.IsValid() {
		switch value.Interface().(type) {
		case *int:
			logger.Info(
				log.String("title", "QingStor response code"),
				log.Int("status_code", int64(b.resp.StatusCode)),
				log.Int("date", convert.StringToTimestamp(b.resp.Header.Get("Date"), convert.RFC822)),
			)
			value.Set(reflect.ValueOf(&b.resp.StatusCode))
		}
	}

	return nil
}

func (b *unpacker) parseResponseHeaders(ctx context.Context) error {
	logger := log.FromContext(ctx)
	logger.Info(
		log.String("title", "QingStor response header"),
		log.Int("date", convert.StringToTimestamp(b.resp.Header.Get("Date"), convert.RFC822)),
		log.String("header", fmt.Sprint(b.resp.Header)),
	)

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

func (b *unpacker) parseResponseBody(ctx context.Context) error {
	logger := log.FromContext(ctx)
	if b.isResponseRight() {
		value := b.output.Elem().FieldByName("Body")
		if value.IsValid() {
			switch value.Type().String() {
			case "string":
				buffer := &bytes.Buffer{}
				buffer.ReadFrom(b.resp.Body)
				b.resp.Body.Close()

				logger.Info(
					log.String("title", "QingStor response body"),
					log.Int("date", convert.StringToTimestamp(b.resp.Header.Get("Date"), convert.RFC822)),
					log.Bytes("body", buffer.Bytes()),
				)

				value.SetString(string(buffer.Bytes()))
			case "io.ReadCloser":
				value.Set(reflect.ValueOf(b.resp.Body))
			}
		}
	}

	return nil
}

func (b *unpacker) parseResponseElements(ctx context.Context) error {
	logger := log.FromContext(ctx)
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

	logger.Info(
		log.String("title", "QingStor response body"),
		log.Int("date", convert.StringToTimestamp(b.resp.Header.Get("Date"), convert.RFC822)),
		log.Bytes("body", buffer.Bytes()),
	)

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

func (b *unpacker) parseError(ctx context.Context) error {
	if b.isResponseRight() {
		return nil
	}

	respCode, requestID := b.resp.StatusCode, b.resp.Header.Get(http.CanonicalHeaderKey("X-QS-Request-ID"))
	logger := log.FromContext(ctx).WithFields(log.String("requestID", requestID))

	qsError := &errors.QingStorError{
		StatusCode: respCode,
		RequestID:  requestID,
	}

	// QingStor nginx could refuse user's request directly and only return status code.
	// We should handle this and return qsError directly.
	if b.resp.ContentLength <= 0 {
		return qsError
	}

	buffer := &bytes.Buffer{}
	_, err := io.Copy(buffer, b.resp.Body)
	if err != nil {
		logger.Error(
			log.String("action", "copy_from_error_response_body"),
			log.String("err", err.Error()),
		)
		return errors.NewUnhandledError(
			errors.WithRequestID(requestID),
			errors.WithStatusCode(respCode),
			errors.WithDetail(err.Error()),
			errors.WithMessage("copy from error response body failed"),
		)

	}

	// close response body after copy
	if err = b.resp.Body.Close(); err != nil {
		logger.Error(
			log.String("action", "close_error_response_body"),
			log.String("err", err.Error()),
		)
		return errors.NewUnhandledError(
			errors.WithRequestID(requestID),
			errors.WithStatusCode(respCode),
			errors.WithDetail(err.Error()),
			errors.WithMessage("close error response body failed"),
		)
	}

	// don't handle non-json error (qs error is surely json format), return body as it is
	if !strings.HasPrefix(b.resp.Header.Get("Content-Type"), "application/json") {
		return errors.NewUnhandledError(
			errors.WithRequestID(requestID),
			errors.WithStatusCode(respCode),
			errors.WithDetail(buffer.String()),
			errors.WithMessage("content not json"),
		)
	}

	// if body is blank, return qsError directly
	if buffer.Len() <= 0 {
		return qsError
	}

	// try to unmarshal body as qsError, if failed return unhandled error
	if err = json.Unmarshal(buffer.Bytes(), qsError); err != nil {
		logger.Error(
			log.String("action", "close_error_response_body_failed"),
			log.Bytes("body", buffer.Bytes()),
			log.String("err", err.Error()),
		)
		return errors.NewUnhandledError(
			errors.WithRequestID(requestID),
			errors.WithStatusCode(respCode),
			errors.WithDetail(buffer.String()),
			errors.WithMessage(err.Error()),
		)
	}

	return qsError
}
