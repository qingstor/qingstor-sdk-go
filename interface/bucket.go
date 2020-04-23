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

package iface

import (
	"github.com/qingstor/qingstor-sdk-go/v4/service"
)

// bucket is the method set for bucket sub service.
type bucket interface {

	// Delete does Delete a bucket.
	Delete() (*service.DeleteBucketOutput, error)

	// DeleteCORS does Delete CORS information of the bucket.
	DeleteCORS() (*service.DeleteBucketCORSOutput, error)

	// DeleteExternalMirror does Delete external mirror of the bucket.
	DeleteExternalMirror() (*service.DeleteBucketExternalMirrorOutput, error)

	// DeleteLifecycle does Delete Lifecycle information of the bucket.
	DeleteLifecycle() (*service.DeleteBucketLifecycleOutput, error)

	// DeleteNotification does Delete Notification information of the bucket.
	DeleteNotification() (*service.DeleteBucketNotificationOutput, error)

	// DeletePolicy does Delete policy information of the bucket.
	DeletePolicy() (*service.DeleteBucketPolicyOutput, error)

	// DeleteMultipleObjects does Delete multiple objects from the bucket.
	DeleteMultipleObjects(input *service.DeleteMultipleObjectsInput) (*service.DeleteMultipleObjectsOutput, error)

	// GetACL does Get ACL information of the bucket.
	GetACL() (*service.GetBucketACLOutput, error)

	// GetCORS does Get CORS information of the bucket.
	GetCORS() (*service.GetBucketCORSOutput, error)

	// GetExternalMirror does Get external mirror of the bucket.
	GetExternalMirror() (*service.GetBucketExternalMirrorOutput, error)

	// GetLifecycle does Get Lifecycle information of the bucket.
	GetLifecycle() (*service.GetBucketLifecycleOutput, error)

	// GetNotification does Get Notification information of the bucket.
	GetNotification() (*service.GetBucketNotificationOutput, error)

	// GetPolicy does Get policy information of the bucket.
	GetPolicy() (*service.GetBucketPolicyOutput, error)

	// GetStatistics does Get statistics information of the bucket.
	GetStatistics() (*service.GetBucketStatisticsOutput, error)

	// Head does Check whether the bucket exists and available.
	Head() (*service.HeadBucketOutput, error)

	// ListMultipartUploads does List multipart uploads in the bucket.
	ListMultipartUploads(input *service.ListMultipartUploadsInput) (*service.ListMultipartUploadsOutput, error)

	// ListObjects does Retrieve the object list in a bucket.
	ListObjects(input *service.ListObjectsInput) (*service.ListObjectsOutput, error)

	// Put does Create a new bucket.
	Put() (*service.PutBucketOutput, error)

	// PutACL does Set ACL information of the bucket.
	PutACL(input *service.PutBucketACLInput) (*service.PutBucketACLOutput, error)

	// PutCORS does Set CORS information of the bucket.
	PutCORS(input *service.PutBucketCORSInput) (*service.PutBucketCORSOutput, error)

	// PutExternalMirror does Set external mirror of the bucket.
	PutExternalMirror(input *service.PutBucketExternalMirrorInput) (*service.PutBucketExternalMirrorOutput, error)

	// PutLifecycle does Set Lifecycle information of the bucket.
	PutLifecycle(input *service.PutBucketLifecycleInput) (*service.PutBucketLifecycleOutput, error)

	// PutNotification does Set Notification information of the bucket.
	PutNotification(input *service.PutBucketNotificationInput) (*service.PutBucketNotificationOutput, error)

	// PutPolicy does Set policy information of the bucket.
	PutPolicy(input *service.PutBucketPolicyInput) (*service.PutBucketPolicyOutput, error)
}
