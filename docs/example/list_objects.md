# ListObjects Example

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

then you can list objects

```go
    resp, err := bucketService.ListObjects(&service.ListObjectsInput{})
    if err != nil {
        fmt.Printf("List Objects on bucket: %s failed with given error: %s\n", bucketName, err)
    } else {
        for _, objInfo := range resp.Keys {
            b, _ := json.Marshal(objInfo)
            fmt.Println(string(b))
        }
    }
```

Add some options which act as filter when list bucket objects

You can set options below in ListObjectsInput. See controlled [API Docs](https://docs.qingcloud.com/qingstor/api/bucket/get).

| Parameter name |  Type   | Description                                                                                                                                                                                                                                                         | Required |
| :------------: | :-----: | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | :------: |
|     prefix     | String  | Limits the response to keys that begin with the specified prefix.                                                                                                                                                                                                   |    No    |
|   delimiter    |  Char   | A delimiter is a character you use to group keys.<br/>If you specify a prefix, all keys that contain the same string between the prefix and the first occurrence of the delimiter after the prefix are grouped under a single result element called CommonPrefixes. |    No    |
|     marker     | String  | Specifies the key to start with when listing objects in a bucket.                                                                                                                                                                                                   |    No    |
|     limit      | Integer | Sets the maximum number of objects returned in the response body. Default is 200, maximum is 1000.                                                                                                                                                                  |    No    |

The following code shows all the objects in the *test* folder in Bucket (without subfolders), sorted by file name by default.

```go
// List objects return objects which start with `prefix` and behind an object named `marker`.
// objects are located in the bucket bound with `bucketService`.
// The records returned is limit to 10.
func listObjects(bucketService *service.Bucket, prefix string, marker string) *string {
	delimiter := "/"
	limit := 10
	params := &service.ListObjectsInput{
		Delimiter: &delimiter,
		Limit:     &limit,
		Marker:    &marker,
		Prefix:    &prefix,
	}
	resp, err := bucketService.ListObjects(params)
	if err != nil {
		fmt.Printf("List Objects on bucket: %s failed with given error: %s\n", *bucketService.Properties.BucketName, err)
		return nil
	} else {
		fmt.Println("=============List Objects=============")
		for _, objInfo := range resp.Keys {
			b, _ := json.Marshal(objInfo)
			fmt.Println(string(b))
		}
		fmt.Println("=====End======")
		return resp.NextMarker
	}
}
```

If the return value is not empty, there is still data on the next page, you can continue to access. The following is an example of a call:

```go
	bucketName := "your-bucket-name"
	zoneName := "pek3b"
	bucketService, _ := qingStor.Bucket(bucketName, zoneName)
	nextMarker := listObjects(bucketService, "test/", "")
	for nextMarker != nil && *nextMarker != "" { // result have next page
		nextMarker = listObjects(bucketService, "test/", *nextMarker)
	}
```