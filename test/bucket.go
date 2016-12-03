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
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"

	"github.com/yunify/qingstor-sdk-go/request/errs"
	qs "github.com/yunify/qingstor-sdk-go/service"
)

// BucketFeatureContext provides feature context for bucket.
func BucketFeatureContext(s *godog.Suite) {
	s.Step(`^initialize the bucket$`, initializeTheBucket)
	s.Step(`^the bucket is initialized$`, theBucketIsInitialized)

	s.Step(`^put bucket$`, putBucketFake)
	s.Step(`^put bucket status code is (\d+)$`, putBucketStatusCodeIsFake)
	s.Step(`^put same bucket again$`, putSameBucketAgain)
	s.Step(`^put same bucket again status code is (\d+)$$`, putSameBucketAgainStatusCodeIs)

	s.Step(`^list objects$`, listObjects)
	s.Step(`^list objects status code is (\d+)$`, listObjectsStatusCodeIs)
	s.Step(`^list objects keys count is (\d+)$`, listObjectsKeysCountIs)

	s.Step(`^head bucket$`, headBucket)
	s.Step(`^head bucket status code is (\d+)$`, headBucketStatusCodeIs)

	s.Step(`^delete bucket$`, deleteBucketFake)
	s.Step(`^delete bucket status code is (\d+)$`, deleteBucketStatusCodeIsFake)

	s.Step(`^delete multiple objects:$`, deleteMultipleObjects)
	s.Step(`^delete multiple objects code is (\d+)$`, deleteMultipleObjectsCodeIs)

	s.Step(`^get bucket statistics$`, getBucketStatistics)
	s.Step(`^get bucket statistics status code is (\d+)$`, getBucketStatisticsStatusCodeIs)
	s.Step(`^get bucket statistics status is "([^"]*)"$`, getBucketStatisticsStatusIs)
}

// --------------------------------------------------------------------------

var bucket *qs.Bucket

func initializeTheBucket() error {
	bucket, err = qsService.Bucket(tc.BucketName, tc.Zone)
	return err
}

func theBucketIsInitialized() error {
	if bucket == nil {
		return errors.New("Bucket is not initialized")
	}
	return nil
}

// --------------------------------------------------------------------------

var putBucketOutput *qs.PutBucketOutput

func putBucket() error {
	putBucketOutput, err = bucket.Put()
	return err
}

func putBucketFake() error {
	return nil
}

func putBucketStatusCodeIs(statusCode int) error {
	if putBucketOutput != nil {
		return checkEqual(putBucketOutput.StatusCode, statusCode)
	}
	return nil
}

func putBucketStatusCodeIsFake(_ int) error {
	return nil
}

// --------------------------------------------------------------------------

func putSameBucketAgain() error {
	_, err = bucket.Put()
	return nil
}

func putSameBucketAgainStatusCodeIs(statusCode int) error {
	switch e := err.(type) {
	case *errs.QingStorError:
		return checkEqual(e.StatusCode, statusCode)
	}

	return fmt.Errorf("put same bucket again should get \"%d\"", statusCode)
}

// --------------------------------------------------------------------------

var listObjectsOutput *qs.ListObjectsOutput

func listObjects() error {
	listObjectsOutput, err = bucket.ListObjects(&qs.ListObjectsInput{
		Delimiter: "/",
		Limit:     1000,
		Prefix:    "Test/",
		Marker:    "Next",
	})
	return err
}

func listObjectsStatusCodeIs(statusCode int) error {
	if listObjectsOutput != nil {
		return checkEqual(listObjectsOutput.StatusCode, statusCode)
	}
	return err
}

func listObjectsKeysCountIs(count int) error {
	return checkEqual(len(listObjectsOutput.Keys), count)
}

// --------------------------------------------------------------------------

var headBucketOutput *qs.HeadBucketOutput

func headBucket() error {
	headBucketOutput, err = bucket.Head()
	return err
}

func headBucketStatusCodeIs(statusCode int) error {
	if headBucketOutput != nil {
		return checkEqual(headBucketOutput.StatusCode, statusCode)
	}
	return err
}

// --------------------------------------------------------------------------

var deleteBucketOutput *qs.DeleteBucketOutput

func deleteBucket() error {
	deleteBucketOutput, err = bucket.Delete()
	return err
}

func deleteBucketFake() error {
	return nil
}

func deleteBucketStatusCodeIs(statusCode int) error {
	if deleteBucketOutput != nil {
		return checkEqual(deleteBucketOutput.StatusCode, statusCode)
	}
	return err
}

func deleteBucketStatusCodeIsFake(_ int) error {
	return nil
}

// --------------------------------------------------------------------------

var deleteMultipleObjectsOutput *qs.DeleteMultipleObjectsOutput

func deleteMultipleObjects(requestJSON *gherkin.DocString) error {
	_, err := bucket.PutObject("object_0", nil)
	if err != nil {
		return err
	}
	_, err = bucket.PutObject("object_1", nil)
	if err != nil {
		return err
	}
	_, err = bucket.PutObject("object_2", nil)
	if err != nil {
		return err
	}

	deleteMultipleObjectsInput := &qs.DeleteMultipleObjectsInput{}
	err = json.Unmarshal([]byte(requestJSON.Content), deleteMultipleObjectsInput)
	if err != nil {
		return err
	}

	requestData := map[string]interface{}{
		"objects": deleteMultipleObjectsInput.Objects,
		"quiet":   deleteMultipleObjectsInput.Quiet,
	}
	jsonBytes, err := json.Marshal(requestData)
	if err != nil {
		return err
	}
	md5Value := md5.Sum(jsonBytes)

	deleteMultipleObjectsOutput, err = bucket.DeleteMultipleObjects(
		&qs.DeleteMultipleObjectsInput{
			Objects:    deleteMultipleObjectsInput.Objects,
			Quiet:      deleteMultipleObjectsInput.Quiet,
			ContentMD5: base64.StdEncoding.EncodeToString(md5Value[:]),
		},
	)
	return err
}

func deleteMultipleObjectsCodeIs(statusCode int) error {
	if deleteMultipleObjectsOutput != nil {
		return checkEqual(deleteMultipleObjectsOutput.StatusCode, statusCode)
	}
	return err
}

// --------------------------------------------------------------------------

var getBucketStatisticsOutput *qs.GetBucketStatisticsOutput

func getBucketStatistics() error {
	getBucketStatisticsOutput, err = bucket.GetStatistics()
	return err
}

func getBucketStatisticsStatusCodeIs(statusCode int) error {
	if getBucketStatisticsOutput != nil {
		return checkEqual(getBucketStatisticsOutput.StatusCode, statusCode)
	}
	return err
}

func getBucketStatisticsStatusIs(status string) error {
	if getBucketStatisticsOutput != nil {
		return checkEqual(getBucketStatisticsOutput.Status, status)
	}
	return err
}
