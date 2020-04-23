# GET Bucket Lifecycle

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

After created the object, we need perform the action to GET Bucket Lifecycle：

```go
	if output, err := bucketService.GetLifecycle(); err != nil {
		fmt.Printf("Get life cycles of bucket(name: %s) failed with given error: %s\n", bucketName, err)
	} else {
		b, _ := json.Marshal(output.Rule)
		fmt.Println("The life cycles of this bucket: ", string(b))
	}
```