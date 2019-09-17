# Delete a Bucket

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

After created the object, we need perform the action to delete a Bucket：

```go
	if resp, err := bucketService.Delete(); err != nil {
		fmt.Printf("Delete bucket(name: %s) failed with given error: %s\n", bucketName, err)
	} else {
		fmt.Printf("The status code expected: 204(actually: %d)\n", *resp.StatusCode)
	}
```

The function that appears in the code above:
- `bucketService.Delete()` Deletes a Bucket named `your-bucket-name` in the `pek3b` field.

The object that appears in the above code:
- The `resp` object is the return value of the `bucketService.Delete()` method.
- `resp.StatusCode` stores the http status code for the api operation.

