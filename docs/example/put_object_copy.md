# CopyObject Example

## Code Snippet

Initialize the Qingstor object with your AccessKeyID and SecretAccessKey.

```go
import (
	"github.com/yunify/qingstor-sdk-go/v3/config"
	"github.com/yunify/qingstor-sdk-go/v3/service"
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

Then set the input parameters used by the PutObject method (core parameter: XQSCopySource).

```go
	// Please replace this path with some file exists on your bucket.
	sourcePath := "/your-bucket-name/your-picture-uploaded.jpg"
	input := &service.PutObjectInput{
		XQSCopySource: &sourcePath,
	}
```

Please note that not all fields in PutObjectInput required to be set. For details, please refer to [Official API Documentation](https://docs.qingcloud.com/qingstor/api/object/copy).

Then call the PutObject method to copy the object. objectKey Sets the copied filepath (located in the current bucket).

```go
	if output, err := bucketService.PutObject(objectKey, input); err != nil {
		fmt.Printf("Copy object from source storage space(%s) to target path(%s) failed with given error: %s\n", sourcePath, objectKey, err)
	} else {
		fmt.Printf("The status code expected: 201(actually: %d)\n", *output.StatusCode)
	}
```