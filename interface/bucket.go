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
	"context"

	"github.com/qingstor/qingstor-sdk-go/v4/service"
)

// bucket is the method set for bucket sub service.
type bucket interface {

	// Delete does Delete a bucket.
	Delete() (*service.DeleteBucketOutput, error)
	DeleteWithContext(ctx context.Context) (*service.DeleteBucketOutput, error)

	// DeleteCNAME does Delete bucket CNAME setting of the bucket.
	DeleteCNAME(input *service.DeleteBucketCNAMEInput) (*service.DeleteBucketCNAMEOutput, error)
	DeleteCNAMEWithContext(ctx context.Context, input *service.DeleteBucketCNAMEInput) (*service.DeleteBucketCNAMEOutput, error)

	// DeleteCORS does Delete CORS information of the bucket.
	DeleteCORS() (*service.DeleteBucketCORSOutput, error)
	DeleteCORSWithContext(ctx context.Context) (*service.DeleteBucketCORSOutput, error)

	// DeleteExternalMirror does Delete external mirror of the bucket.
	DeleteExternalMirror() (*service.DeleteBucketExternalMirrorOutput, error)
	DeleteExternalMirrorWithContext(ctx context.Context) (*service.DeleteBucketExternalMirrorOutput, error)

	// DeleteLifecycle does Delete Lifecycle information of the bucket.
	DeleteLifecycle() (*service.DeleteBucketLifecycleOutput, error)
	DeleteLifecycleWithContext(ctx context.Context) (*service.DeleteBucketLifecycleOutput, error)

	// DeleteLogging does Delete bucket logging setting of the bucket.
	DeleteLogging() (*service.DeleteBucketLoggingOutput, error)
	DeleteLoggingWithContext(ctx context.Context) (*service.DeleteBucketLoggingOutput, error)

	// DeleteNotification does Delete Notification information of the bucket.
	DeleteNotification() (*service.DeleteBucketNotificationOutput, error)
	DeleteNotificationWithContext(ctx context.Context) (*service.DeleteBucketNotificationOutput, error)

	// DeletePolicy does Delete policy information of the bucket.
	DeletePolicy() (*service.DeleteBucketPolicyOutput, error)
	DeletePolicyWithContext(ctx context.Context) (*service.DeleteBucketPolicyOutput, error)

	// DeleteReplication does Delete Replication information of the bucket.
	DeleteReplication() (*service.DeleteBucketReplicationOutput, error)
	DeleteReplicationWithContext(ctx context.Context) (*service.DeleteBucketReplicationOutput, error)

	// DeleteMultipleObjects does Delete multiple objects from the bucket.
	DeleteMultipleObjects(input *service.DeleteMultipleObjectsInput) (*service.DeleteMultipleObjectsOutput, error)
	DeleteMultipleObjectsWithContext(ctx context.Context, input *service.DeleteMultipleObjectsInput) (*service.DeleteMultipleObjectsOutput, error)

	// GetACL does Get ACL information of the bucket.
	GetACL() (*service.GetBucketACLOutput, error)
	GetACLWithContext(ctx context.Context) (*service.GetBucketACLOutput, error)

	// GetCNAME does Get bucket CNAME setting of the bucket.
	GetCNAME(input *service.GetBucketCNAMEInput) (*service.GetBucketCNAMEOutput, error)
	GetCNAMEWithContext(ctx context.Context, input *service.GetBucketCNAMEInput) (*service.GetBucketCNAMEOutput, error)

	// GetCORS does Get CORS information of the bucket.
	GetCORS() (*service.GetBucketCORSOutput, error)
	GetCORSWithContext(ctx context.Context) (*service.GetBucketCORSOutput, error)

	// GetExternalMirror does Get external mirror of the bucket.
	GetExternalMirror() (*service.GetBucketExternalMirrorOutput, error)
	GetExternalMirrorWithContext(ctx context.Context) (*service.GetBucketExternalMirrorOutput, error)

	// GetLifecycle does Get Lifecycle information of the bucket.
	GetLifecycle() (*service.GetBucketLifecycleOutput, error)
	GetLifecycleWithContext(ctx context.Context) (*service.GetBucketLifecycleOutput, error)

	// GetLogging does Get bucket logging setting of the bucket.
	GetLogging() (*service.GetBucketLoggingOutput, error)
	GetLoggingWithContext(ctx context.Context) (*service.GetBucketLoggingOutput, error)

	// GetNotification does Get Notification information of the bucket.
	GetNotification() (*service.GetBucketNotificationOutput, error)
	GetNotificationWithContext(ctx context.Context) (*service.GetBucketNotificationOutput, error)

	// GetPolicy does Get policy information of the bucket.
	GetPolicy() (*service.GetBucketPolicyOutput, error)
	GetPolicyWithContext(ctx context.Context) (*service.GetBucketPolicyOutput, error)

	// GetReplication does Get Replication information of the bucket.
	GetReplication() (*service.GetBucketReplicationOutput, error)
	GetReplicationWithContext(ctx context.Context) (*service.GetBucketReplicationOutput, error)

	// GetStatistics does Get statistics information of the bucket.
	GetStatistics() (*service.GetBucketStatisticsOutput, error)
	GetStatisticsWithContext(ctx context.Context) (*service.GetBucketStatisticsOutput, error)

	// Head does Check whether the bucket exists and available.
	Head() (*service.HeadBucketOutput, error)
	HeadWithContext(ctx context.Context) (*service.HeadBucketOutput, error)

	// ListMultipartUploads does List multipart uploads in the bucket.
	ListMultipartUploads(input *service.ListMultipartUploadsInput) (*service.ListMultipartUploadsOutput, error)
	ListMultipartUploadsWithContext(ctx context.Context, input *service.ListMultipartUploadsInput) (*service.ListMultipartUploadsOutput, error)

	// ListObjects does Retrieve the object list in a bucket.
	ListObjects(input *service.ListObjectsInput) (*service.ListObjectsOutput, error)
	ListObjectsWithContext(ctx context.Context, input *service.ListObjectsInput) (*service.ListObjectsOutput, error)

	// Put does Create a new bucket.
	Put() (*service.PutBucketOutput, error)
	PutWithContext(ctx context.Context) (*service.PutBucketOutput, error)

	// PutACL does Set ACL information of the bucket.
	PutACL(input *service.PutBucketACLInput) (*service.PutBucketACLOutput, error)
	PutACLWithContext(ctx context.Context, input *service.PutBucketACLInput) (*service.PutBucketACLOutput, error)

	// PutCNAME does Set bucket CNAME of the bucket.
	PutCNAME(input *service.PutBucketCNAMEInput) (*service.PutBucketCNAMEOutput, error)
	PutCNAMEWithContext(ctx context.Context, input *service.PutBucketCNAMEInput) (*service.PutBucketCNAMEOutput, error)

	// PutCORS does Set CORS information of the bucket.
	PutCORS(input *service.PutBucketCORSInput) (*service.PutBucketCORSOutput, error)
	PutCORSWithContext(ctx context.Context, input *service.PutBucketCORSInput) (*service.PutBucketCORSOutput, error)

	// PutExternalMirror does Set external mirror of the bucket.
	PutExternalMirror(input *service.PutBucketExternalMirrorInput) (*service.PutBucketExternalMirrorOutput, error)
	PutExternalMirrorWithContext(ctx context.Context, input *service.PutBucketExternalMirrorInput) (*service.PutBucketExternalMirrorOutput, error)

	// PutLifecycle does Set Lifecycle information of the bucket.
	PutLifecycle(input *service.PutBucketLifecycleInput) (*service.PutBucketLifecycleOutput, error)
	PutLifecycleWithContext(ctx context.Context, input *service.PutBucketLifecycleInput) (*service.PutBucketLifecycleOutput, error)

	// PutLogging does Set bucket logging of the bucket.
	PutLogging(input *service.PutBucketLoggingInput) (*service.PutBucketLoggingOutput, error)
	PutLoggingWithContext(ctx context.Context, input *service.PutBucketLoggingInput) (*service.PutBucketLoggingOutput, error)

	// PutNotification does Set Notification information of the bucket.
	PutNotification(input *service.PutBucketNotificationInput) (*service.PutBucketNotificationOutput, error)
	PutNotificationWithContext(ctx context.Context, input *service.PutBucketNotificationInput) (*service.PutBucketNotificationOutput, error)

	// PutPolicy does Set policy information of the bucket.
	PutPolicy(input *service.PutBucketPolicyInput) (*service.PutBucketPolicyOutput, error)
	PutPolicyWithContext(ctx context.Context, input *service.PutBucketPolicyInput) (*service.PutBucketPolicyOutput, error)

	// PutReplication does Set Replication information of the bucket.
	PutReplication(input *service.PutBucketReplicationInput) (*service.PutBucketReplicationOutput, error)
	PutReplicationWithContext(ctx context.Context, input *service.PutBucketReplicationInput) (*service.PutBucketReplicationOutput, error)
}
