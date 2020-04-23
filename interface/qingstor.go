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

// Package iface provides QingStor Service API interface (API Version 2016-01-06)
package iface

import (
	"github.com/qingstor/qingstor-sdk-go/v4/service"
)

// Service is the method set for all public service API.
type Service interface {
	// Bucket initializes a new bucket.
	Bucket(bucketName string, zone string) (*service.Bucket, error)

	// ListBuckets does Retrieve the bucket list.
	ListBuckets(input *service.ListBucketsInput) (*service.ListBucketsOutput, error)
}

// Bucket is the method set for all public bucket API.
type Bucket interface {

	// bucket is the method set for bucket sub service.
	bucket

	// object is the method set for object sub service.
	object
}
