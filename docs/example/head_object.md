# HEAD Object

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

Then set the input parameters used by the HeadObject method (stored in HeadObjectInput).

```go
	input := &service.HeadObjectInput{}
```

Please note that the fields in HeadObjectInput is not required to be set. For details, please refer to [Official API Documentation](https://docsv4.qingcloud.com/user_guide/storage/object_storage/api/object/basic_opt/head/).

Then call the HeadObject method to get the object meta information and test if it can be accessed. objectKey Sets the filepath of the object to be fetched (in the current bucket).

```go
	// Please replace this file path with some file exists on your bucket.
	objectKey := "your_file.zip"
	if output, err := bucketService.HeadObject(objectKey, input); err != nil {
		fmt.Printf("The attempt to access a object(name: %s) metadata failed with given error: %s\n", bucketName, err)
	} else {
		fmt.Printf("The status code expected: 200(actually: %d)\n", * output.StatusCode)
		b, _ := json.MarshalIndent(output, "", "\t")
		fmt.Printf("The metadata of object(%s):\n %s\n", objectKey, string(b))
	}
```

If the operation returns correctly, the response status code will be 200.
