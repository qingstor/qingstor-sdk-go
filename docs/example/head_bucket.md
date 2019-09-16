# Head a Bucket

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

After the object is created, we need to perform the actual Bucket meta information operation:

```go
	if resp, err := bucketService.Head(); err != nil {
		fmt.Printf("The attempt to access a bucket(name: %s) failed with given error: %s\n", bucketName, err)
	} else {
		fmt.Printf("The status code expected: 200(actually: %d)\n", *resp.StatusCode)
	}
```

The function that appears in the code above:
- `bucketService.Head()` In the `pek3b` area, try to use HEAD to get a Bucket message named `your-bucket-name`.

The object that appears in the above code:
- The `resp` object is the return value of the `bucketService.Head()` method.
- `resp.StatusCode` stores the http status code for the api operation.

