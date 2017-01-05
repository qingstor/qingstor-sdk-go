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
	"encoding/json"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"

	qs "github.com/yunify/qingstor-sdk-go/service"
)

// BucketExternalMirrorFeatureContext provides feature context for bucket external mirror.
func BucketExternalMirrorFeatureContext(s *godog.Suite) {
	s.Step(`^put bucket external mirror:$`, putBucketExternalMirror)
	s.Step(`^put bucket external mirror status code is (\d+)$`, putBucketExternalMirrorStatusCodeIs)

	s.Step(`^get bucket external mirror$`, getBucketExternalMirror)
	s.Step(`^get bucket external mirror status code is (\d+)$`, getBucketExternalMirrorStatusCodeIs)
	s.Step(`^get bucket external mirror should have source_site "([^"]*)"$`, getBucketExternalMirrorShouldHaveSourceSite)

	s.Step(`^delete bucket external mirror$`, deleteBucketExternalMirror)
	s.Step(`^delete bucket external mirror status code is (\d+)$`, deleteBucketExternalMirrorStatusCodeIs)
}

// --------------------------------------------------------------------------

var putBucketExternalMirrorOutput *qs.PutBucketExternalMirrorOutput

func putBucketExternalMirror(ExternalMirrorJSONText *gherkin.DocString) error {
	putBucketExternalMirrorInput := &qs.PutBucketExternalMirrorInput{}
	err = json.Unmarshal([]byte(ExternalMirrorJSONText.Content), putBucketExternalMirrorInput)
	if err != nil {
		return err
	}

	putBucketExternalMirrorOutput, err = bucket.PutExternalMirror(putBucketExternalMirrorInput)
	return err
}

func putBucketExternalMirrorStatusCodeIs(statusCode int) error {
	if putBucketExternalMirrorOutput != nil {
		return checkEqual(qs.IntValue(putBucketExternalMirrorOutput.StatusCode), statusCode)
	}
	return err
}

// --------------------------------------------------------------------------

var getBucketExternalMirrorOutput *qs.GetBucketExternalMirrorOutput

func getBucketExternalMirror() error {
	getBucketExternalMirrorOutput, err = bucket.GetExternalMirror()
	return err
}

func getBucketExternalMirrorStatusCodeIs(statusCode int) error {
	if getBucketExternalMirrorOutput != nil {
		return checkEqual(qs.IntValue(getBucketExternalMirrorOutput.StatusCode), statusCode)
	}
	return err
}

func getBucketExternalMirrorShouldHaveSourceSite(sourceSite string) error {
	if getBucketExternalMirrorOutput != nil {
		return checkEqual(qs.StringValue(getBucketExternalMirrorOutput.SourceSite), sourceSite)
	}
	return err
}

// --------------------------------------------------------------------------

var deleteBucketExternalMirrorOutput *qs.DeleteBucketExternalMirrorOutput

func deleteBucketExternalMirror() error {
	deleteBucketExternalMirrorOutput, err = bucket.DeleteExternalMirror()
	return err
}

func deleteBucketExternalMirrorStatusCodeIs(statusCode int) error {
	if deleteBucketExternalMirrorOutput != nil {
		return checkEqual(qs.IntValue(deleteBucketExternalMirrorOutput.StatusCode), statusCode)
	}
	return err
}
