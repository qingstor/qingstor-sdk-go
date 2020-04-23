# List Multipart Uploads

Get a list of ongoing multipart upload objects. When an object has the segment upload mode enabled via the Initiate Multipart interface, the object is in the "multiply uploading" state before the Complete Multipart or Abort Multipart interface is called. This object will appear in the list returned by the interface.

Similar to the GET Bucket (List Objects) interface, the user can specify the segment upload that is being performed under a directory by passing the prefix, delimiter request parameter. The list is sorted from small to large in alphanumeric order of object names. If there are multiple segments uploaded by the same-named object, only one part of the page is truncated, and the next page-turning can be obtained by upload_id_marker parameter to obtain the segment upload after the upload_id is sorted by creation time.

If the user only wants to get the segment that an object has uploaded, please refer to [API Docs] (https://docs.qingcloud.com/qingstor/api/object/multipart/list_multipart.html#object-storage-api-list -multipart).

## Request Parameters

You can add some filter options when list multipart uploads.

You can set options below in ListMultipartUploadsInput. See controlled [API Docs](https://docs.qingcloud.com/qingstor/api/bucket/list_multipart_uploads.html).

|  Parameter name  |  Type   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | Required |
| :--------------: | :-----: | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | :------: |
|      prefix      | String  | Limits the response to keys that begin with the specified prefix.                                                                                                                                                                                                                                                                                                                                                                                                                                 |    No    |
|    delimiter     |  Char   | A delimiter is a character you use to group keys.<br/>If you specify a prefix, all keys that contain the same string between the prefix and the first occurrence of the delimiter after the prefix are grouped under a single result element called CommonPrefixes.                                                                                                                                                                                                                               |    No    |
|    key_marker    | String  | Together with upload-id-marker, this parameter specifies the multipart upload after which listing should begin.<br>If upload-id-marker is not specified, only the keys lexicographically greater than the specified key-marker will be included in the list.<br>If upload-id-marker is specified, any multipart uploads for a key equal to the key-marker might also be included, provided those multipart uploads have upload IDs lexicographically greater than the specified upload-id-marker. |    No    |
| upload_id_marker | String  | Together with key-marker, specifies the multipart upload after which listing should begin. If key-marker is not specified, the upload-id-marker parameter is ignored. Otherwise, any multipart uploads for a key equal to the key-marker might be included in the list only if they have an upload ID lexicographically greater than the specified upload-id-marker.                                                                                                                              |    No    |
|      limit       | Integer | Sets the maximum number of objects returned in the response body. Default is 200, maximum is 1000.                                                                                                                                                                                                                                                                                                                                                                                                |    No    |

## Code Snippet

Initialize the Qingstor object with your AccessKeyID and SecretAccessKey.

```go
import (
	"github.com/qingstor/qingstor-sdk-go/v4/config"
	"github.com/qingstor/qingstor-sdk-go/v4/service"
)

var conf, _ = config.New("YOUR-ACCESS-KEY-ID", "YOUR--SECRET-ACCESS-KEY")
var qingStor, _ = service.Init(conf)
```

Initialize a Bucket object according to the bucket name you set for subsequent creation:

```go
bucketName := "your-bucket-name"
zoneName := "pek3b"
bucketService, _ := qingStor.Bucket(bucketName, zoneName)
```

After the object is created, we need to perform a real List Multipart Uploads operation:

The following operations list all objects that have not yet called the Complete MultiUpload interface under the Movies directory (without subdirectories and their files), limited to six.

```go
	toPtr := func(s string) *string { return &s }
	limit := 6
	resp, err := bucketService.ListMultipartUploads(&service.ListMultipartUploadsInput{
		Delimiter:      toPtr("/"),
		KeyMarker:      nil,
		Limit:          &limit,
		Prefix:         toPtr("Movies/"),
		UploadIDMarker: nil,
	})
	if err != nil {
		fmt.Printf("List Objects(multiUploaded and the complete api not been called) on bucket: %s failed with given error: %s\n", bucketName, err)
	} else {
		for _, objInfo := range resp.Uploads {
			b, _ := json.Marshal(objInfo)
			fmt.Println(string(b))
		}
	}
```
