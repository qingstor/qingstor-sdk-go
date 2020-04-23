# GET Bucket CORS

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

then you can GET Bucket CORS

```go
	if output, err := bucketService.GetCORS(); err != nil {
		fmt.Printf("Get CORS rules of bucket(name: %s) failed with given error: %s\n", bucketName, err)
	} else {
		b, _ := json.Marshal(output.CORSRules)
		fmt.Println("The CORS rules of this bucket: ", string(b))
	}
```