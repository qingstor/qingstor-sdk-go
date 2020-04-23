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

// object is the method set for object sub service.
type object interface {

	// AbortMultipartUpload does Abort multipart upload.
	AbortMultipartUpload(objectKey string, input *service.AbortMultipartUploadInput) (*service.AbortMultipartUploadOutput, error)

	// CompleteMultipartUpload does Complete multipart upload.
	CompleteMultipartUpload(objectKey string, input *service.CompleteMultipartUploadInput) (*service.CompleteMultipartUploadOutput, error)

	// DeleteObject does Delete the object.
	DeleteObject(objectKey string) (*service.DeleteObjectOutput, error)

	// GetObject does Retrieve the object.
	GetObject(objectKey string, input *service.GetObjectInput) (*service.GetObjectOutput, error)

	// HeadObject does Check whether the object exists and available.
	HeadObject(objectKey string, input *service.HeadObjectInput) (*service.HeadObjectOutput, error)

	// ImageProcess does Image process with the action on the object
	ImageProcess(objectKey string, input *service.ImageProcessInput) (*service.ImageProcessOutput, error)

	// InitiateMultipartUpload does Initial multipart upload on the object.
	InitiateMultipartUpload(objectKey string, input *service.InitiateMultipartUploadInput) (*service.InitiateMultipartUploadOutput, error)

	// ListMultipart does List object parts.
	ListMultipart(objectKey string, input *service.ListMultipartInput) (*service.ListMultipartOutput, error)

	// OptionsObject does Check whether the object accepts a origin with method and header.
	OptionsObject(objectKey string, input *service.OptionsObjectInput) (*service.OptionsObjectOutput, error)

	// PutObject does Upload the object.
	PutObject(objectKey string, input *service.PutObjectInput) (*service.PutObjectOutput, error)

	// UploadMultipart does Upload object multipart.
	UploadMultipart(objectKey string, input *service.UploadMultipartInput) (*service.UploadMultipartOutput, error)
}
