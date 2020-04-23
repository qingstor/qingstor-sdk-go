# Get Bucket Statistics

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

After the object is created, we need to perform the actual Bucket usage statistics:

```go
	if resp, err := bucketService.GetStatistics(); err != nil {
		fmt.Printf("Get bucket(name: %s) statistics failed with given error: %s\n", bucketName, err)
	} else {
		fmt.Printf("Bucket info: {size: %d, count: %d, location: %s, url: %s, created: %s}\n",
			*resp.Size, *resp.Count, *resp.Location, *resp.URL, *resp.Created)
	}
```

