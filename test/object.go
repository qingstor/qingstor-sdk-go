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

package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/DATA-DOG/godog"

	"github.com/yunify/qingstor-sdk-go/request"
	qs "github.com/yunify/qingstor-sdk-go/service"
)

// ObjectFeatureContext provides feature context for object.
func ObjectFeatureContext(s *godog.Suite) {
	s.Step(`^put object with key "([^"]*)"$`, putObjectWithKey)
	s.Step(`^put object status code is (\d+)$`, putObjectStatusCodeIs)

	s.Step(`^copy object with key "([^"]*)"$`, copyObjectWithKey)
	s.Step(`^copy object status code is (\d+)$`, copyObjectStatusCodeIs)

	s.Step(`^move object with key "([^"]*)"$`, moveObjectWithKey)
	s.Step(`^move object status code is (\d+)$`, moveObjectStatusCodeIs)

	s.Step(`^get object$`, getObject)
	s.Step(`^get object status code is (\d+)$`, getObjectStatusCodeIs)
	s.Step(`^get object content length is (\d+)$`, getObjectContentLengthIs)
	s.Step(`^get object with content type "([^"]*)"$`, getObjectWithContentType)
	s.Step(`^get object content type is "([^"]*)"$`, getObjectContentTypeIs)
	s.Step(`^get object with query signature$`, getObjectWithQuerySignature)
	s.Step(`^get object with query signature content length is (\d+)$`, getObjectWithQuerySignatureContentLengthIs)

	s.Step(`^head object$`, headObject)
	s.Step(`^head object status code is (\d+)$`, headObjectStatusCodeIs)

	s.Step(`^options object with method "([^"]*)" and origin "([^"]*)"$`, optionsObjectWithMethodAndOrigin)
	s.Step(`^options object status code is (\d+)$`, optionsObjectStatusCodeIs)

	s.Step(`^delete object$`, deleteObject)
	s.Step(`^delete object status code is (\d+)$`, deleteObjectStatusCodeIs)
	s.Step(`^delete the move object$`, deleteTheMoveObject)
	s.Step(`^delete the move object status code is (\d+)$`, deleteTheMoveObjectStatusCodeIs)
}

// --------------------------------------------------------------------------

var theObjectKey string
var theCopyObjectKey string
var theMoveObjectKey string

var putObjectOutput *qs.PutObjectOutput
var copyObjectOutput *qs.PutObjectOutput
var moveObjectOutput *qs.PutObjectOutput

func putObjectWithKey(objectKey string) error {
	theObjectKey = objectKey

	_, err = exec.Command("dd", "if=/dev/zero", "of=/tmp/sdk_bin", "bs=1048576", "count=1").Output()
	if err != nil {
		return err
	}
	defer os.Remove("/tmp/sdk_bin")

	file, err := os.Open("/tmp/sdk_bin")
	if err != nil {
		return err
	}
	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)
	hashInBytes := hash.Sum(nil)[:16]
	md5String := hex.EncodeToString(hashInBytes)

	//file.Seek(0, io.SeekStart)
	file.Seek(0, 0)
	putObjectOutput, err = bucket.PutObject(theObjectKey, &qs.PutObjectInput{
		ContentType: "text/plain",
		ContentMD5:  md5String,
		Body:        file,
	})
	return err
}

func putObjectStatusCodeIs(statusCode int) error {
	if putObjectOutput != nil {
		return checkEqual(putObjectOutput.StatusCode, statusCode)
	}
	return err
}

func copyObjectWithKey(objectKey string) error {
	theCopyObjectKey = objectKey
	copyObjectOutput, err = bucket.PutObject(theCopyObjectKey, &qs.PutObjectInput{
		XQSCopySource: "/" + tc.BucketName + "/" + theObjectKey,
	})
	return err
}

func copyObjectStatusCodeIs(statusCode int) error {
	if copyObjectOutput != nil {
		return checkEqual(copyObjectOutput.StatusCode, statusCode)
	}
	return err
}

func moveObjectWithKey(objectKey string) error {
	theMoveObjectKey = objectKey
	moveObjectOutput, err = bucket.PutObject(theMoveObjectKey, &qs.PutObjectInput{
		XQSMoveSource: "/" + tc.BucketName + "/" + theCopyObjectKey,
	})
	return err
}

func moveObjectStatusCodeIs(statusCode int) error {
	if moveObjectOutput != nil {
		return checkEqual(moveObjectOutput.StatusCode, statusCode)
	}
	return err
}

// --------------------------------------------------------------------------

var getObjectOutput *qs.GetObjectOutput

func getObject() error {
	getObjectOutput, err = bucket.GetObject(theObjectKey, nil)
	return err
}

func getObjectStatusCodeIs(statusCode int) error {
	if getObjectOutput != nil {
		return checkEqual(getObjectOutput.StatusCode, statusCode)
	}
	return err
}

func getObjectContentLengthIs(length int) error {
	buffer := &bytes.Buffer{}
	buffer.ReadFrom(getObjectOutput.Body)
	getObjectOutput.Body.Close()

	ioutil.WriteFile("/tmp/sdk_bin", buffer.Bytes(), 0644)
	defer os.Remove("/tmp/sdk_bin")

	return checkEqual(len(buffer.Bytes()), length)
}

// --------------------------------------------------------------------------

var getObjectWithContentTypeRequest *request.Request

func getObjectWithContentType(contentType string) error {
	getObjectWithContentTypeRequest, _, err = bucket.GetObjectRequest(
		theObjectKey,
		&qs.GetObjectInput{
			ResponseContentType: contentType,
		},
	)
	if err != nil {
		return err
	}
	err = getObjectWithContentTypeRequest.Send()
	return err
}

func getObjectContentTypeIs(contentType string) error {
	return checkEqual(
		getObjectWithContentTypeRequest.HTTPResponse.Header.Get("Content-Type"),
		contentType,
	)
}

// --------------------------------------------------------------------------

var getObjectWithQuerySignatureURL string

func getObjectWithQuerySignature() error {
	getObjectRequest, _, err := bucket.GetObjectRequest(theObjectKey, nil)
	if err != nil {
		return err
	}

	err = getObjectRequest.SignQuery(10)
	if err != nil {
		return err
	}

	getObjectWithQuerySignatureURL = getObjectRequest.HTTPRequest.URL.String()
	return nil
}

func getObjectWithQuerySignatureContentLengthIs(length int) error {
	getObjectResponse, err := http.Get(getObjectWithQuerySignatureURL)
	if err != nil {
		return err
	}

	buffer := &bytes.Buffer{}
	buffer.ReadFrom(getObjectResponse.Body)
	getObjectResponse.Body.Close()

	ioutil.WriteFile("/tmp/sdk_bin", buffer.Bytes(), 0644)
	defer os.Remove("/tmp/sdk_bin")

	return checkEqual(len(buffer.Bytes()), length)
}

// --------------------------------------------------------------------------

var headObjectOutput *qs.HeadObjectOutput

func headObject() error {
	headObjectOutput, err = bucket.HeadObject(theObjectKey, nil)
	return err
}

func headObjectStatusCodeIs(statusCode int) error {
	if headObjectOutput != nil {
		return checkEqual(headObjectOutput.StatusCode, statusCode)
	}
	return err
}

// --------------------------------------------------------------------------

var optionsObjectOutput *qs.OptionsObjectOutput

func optionsObjectWithMethodAndOrigin(method, origin string) error {
	optionsObjectOutput, err = bucket.OptionsObject(
		theObjectKey,
		&qs.OptionsObjectInput{
			AccessControlRequestMethod: method,
			Origin: origin,
		},
	)
	return err
}

func optionsObjectStatusCodeIs(statusCode int) error {
	if optionsObjectOutput != nil {
		return checkEqual(optionsObjectOutput.StatusCode, statusCode)
	}
	return err
}

// --------------------------------------------------------------------------

var deleteObjectOutput *qs.DeleteObjectOutput
var deleteTheMoveObjectOutput *qs.DeleteObjectOutput

func deleteObject() error {
	deleteObjectOutput, err = bucket.DeleteObject(theObjectKey)
	return err
}

func deleteObjectStatusCodeIs(statusCode int) error {
	if deleteObjectOutput != nil {
		return checkEqual(deleteObjectOutput.StatusCode, statusCode)
	}
	return err
}

func deleteTheMoveObject() error {
	deleteTheMoveObjectOutput, err = bucket.DeleteObject(theMoveObjectKey)
	return err
}

func deleteTheMoveObjectStatusCodeIs(statusCode int) error {
	if deleteTheMoveObjectOutput != nil {
		return checkEqual(deleteTheMoveObjectOutput.StatusCode, statusCode)
	}
	return err
}
