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

package request

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/pengsrc/go-shared/convert"
	"github.com/qingstor/log"

	"github.com/qingstor/qingstor-sdk-go/v4/request/builder"
	"github.com/qingstor/qingstor-sdk-go/v4/request/data"
	"github.com/qingstor/qingstor-sdk-go/v4/request/errors"
	"github.com/qingstor/qingstor-sdk-go/v4/request/response"
	"github.com/qingstor/qingstor-sdk-go/v4/request/signer"
)

// A Request can build, sign, send and unpack API request.
type Request struct {
	Operation *data.Operation
	Input     *reflect.Value
	Output    *reflect.Value

	HTTPRequest  *http.Request
	HTTPResponse *http.Response
}

// New create a Request from given Operation, Input and Output.
// It returns a Request.
func New(o *data.Operation, i data.Input, x interface{}) (*Request, error) {
	input := reflect.ValueOf(i)
	if input.IsValid() && input.Elem().IsValid() {
		err := i.Validate()
		if err != nil {
			return nil, err
		}
	}
	output := reflect.ValueOf(x)

	return &Request{
		Operation: o,
		Input:     &input,
		Output:    &output,
	}, nil
}

// SendWithContext sends API request with given ctx.
// It returns error if error occurred.
func (r *Request) SendWithContext(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}

	err := r.BuildWithContext(ctx)
	if err != nil {
		return err
	}

	err = r.SignWithContext(ctx)
	if err != nil {
		return err
	}

	err = r.DoWithContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Send sends API request.
// It returns error if error occurred.
// Deprecated: Use SendWithContext instead
func (r *Request) Send() error {
	return r.SendWithContext(context.Background())
}

// BuildWithContext checks and builds the API request with given ctx.
// It returns error if error occurred.
func (r *Request) BuildWithContext(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}

	err := r.check(ctx)
	if err != nil {
		return err
	}

	err = r.build(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Build checks and builds the API request.
// It returns error if error occurred.
// Deprecated: Use BuildWithContext instead
func (r *Request) Build() error {
	return r.BuildWithContext(context.Background())
}

// DoWithContext sends and unpacks the API request with given ctx.
// It returns error if error occurred.
func (r *Request) DoWithContext(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}

	err := r.send(ctx)
	if err != nil {
		return err
	}

	err = r.unpack(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Do sends and unpacks the API request.
// It returns error if error occurred.
// Deprecated: Use DoWithContext instead
func (r *Request) Do() error {
	return r.DoWithContext(context.Background())
}

// SignWithContext sign the API request by setting the authorization header with given ctx.
// It returns error if error occurred.
func (r *Request) SignWithContext(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}

	if r.Operation.Config.AccessKeyID != "" && r.Operation.Config.SecretAccessKey != "" {
		err := r.sign(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// Sign sign the API request by setting the authorization header.
// It returns error if error occurred.
// Deprecated: Use SignWithContext instead
func (r *Request) Sign() error {
	return r.SignWithContext(context.Background())
}

// SignQuery sign the API request by appending query string.
// It returns error if error occurred.
func (r *Request) SignQuery(timeoutSeconds int) error {
	err := r.signQuery(int(time.Now().Unix()) + timeoutSeconds)
	if err != nil {
		return err
	}

	return nil
}

// ApplySignature applies the Authorization header.
// It returns error if error occurred.
func (r *Request) ApplySignature(authorization string) error {
	r.HTTPRequest.Header.Set("Authorization", authorization)
	return nil
}

// ApplyQuerySignature applies the query signature.
// It returns error if error occurred.
func (r *Request) ApplyQuerySignature(accessKeyID string, expires int, signature string) error {
	queryValue := r.HTTPRequest.URL.Query()
	queryValue.Set("access_key_id", accessKeyID)
	queryValue.Set("expires", strconv.Itoa(expires))
	queryValue.Set("signature", signature)

	r.HTTPRequest.URL.RawQuery = queryValue.Encode()
	return nil
}

func (r *Request) check(ctx context.Context) error {
	if r.Operation.Config.AccessKeyID == "" && r.Operation.Config.SecretAccessKey != "" {
		return errors.NewSDKError(
			errors.WithAction("check request"),
			errors.WithError(fmt.Errorf("access key not provided")),
		)
	}

	if r.Operation.Config.SecretAccessKey == "" && r.Operation.Config.AccessKeyID != "" {
		return errors.NewSDKError(
			errors.WithAction("check request"),
			errors.WithError(fmt.Errorf("secret access key not provided")),
		)
	}

	return nil
}

func (r *Request) build(ctx context.Context) error {
	b := &builder.Builder{}
	httpRequest, err := b.BuildHTTPRequest(ctx, r.Operation, r.Input)
	if err != nil {
		return err
	}

	r.HTTPRequest = httpRequest
	return nil
}

func (r *Request) sign(ctx context.Context) error {
	s := &signer.QingStorSigner{
		AccessKeyID:            r.Operation.Config.AccessKeyID,
		SecretAccessKey:        r.Operation.Config.SecretAccessKey,
		EnableVirtualHostStyle: r.Operation.Config.EnableVirtualHostStyle,
	}
	err := s.WriteSignature(r.HTTPRequest)
	if err != nil {
		return err
	}

	return nil
}

func (r *Request) signQuery(expires int) error {
	s := &signer.QingStorSigner{
		AccessKeyID:     r.Operation.Config.AccessKeyID,
		SecretAccessKey: r.Operation.Config.SecretAccessKey,
	}
	err := s.WriteQuerySignature(r.HTTPRequest, expires)
	if err != nil {
		return err
	}

	return nil
}

func (r *Request) send(ctx context.Context) error {
	logger := log.FromContext(ctx)
	var resp *http.Response
	var err error

	if r.Operation.Config.Connection == nil {
		r.Operation.Config.InitHTTPClient()
	}

	logger.Info(
		log.String("title", "Sending request"),
		log.Int("date", convert.StringToTimestamp(r.HTTPRequest.Header.Get("Date"), convert.RFC822)),
		log.String("method", r.Operation.RequestMethod),
		log.String("host", r.HTTPRequest.Host),
	)

	resp, err = r.Operation.Config.Connection.Do(r.HTTPRequest)
	if err != nil {
		return errors.NewSDKError(
			errors.WithAction("do request in send"),
			errors.WithError(err),
		)
	}

	r.HTTPResponse = resp

	return nil
}

func (r *Request) unpack(ctx context.Context) error {
	err := response.UnpackToOutputWithContext(ctx, r.Operation, r.HTTPResponse, r.Output)
	if err != nil {
		return err
	}

	return nil
}
