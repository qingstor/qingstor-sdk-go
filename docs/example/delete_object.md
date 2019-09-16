# DeleteObject Example

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

Then call the DeleteObject method to delete the object. objectKey Sets the filepath of the object to be deleted (in the current bucket).

```go
	objectKey := "file_your_want_delete"
	if output, err := bucketService.DeleteObject(objectKey); err != nil {
		fmt.Printf("Delete object(name: %s) in bucket(%s) failed with given error: %s\n", objectKey, bucketName, err)
	} else {
		fmt.Printf("The status code expected: 204(actually: %d)\n", *output.StatusCode)
	}
```

If the operation returns correctly, the response status code will be 204.