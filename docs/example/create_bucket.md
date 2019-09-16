# Create a Bucket

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

The `bucketService` object is used to manipulate the Bucket and can use all Bucket and Object level APIs. Now perform the real creation of the Bucket operation:

```go
if _, err := bucketService.Put(); err == nil {
    fmt.Printf("Your bucket named \"%s\" in zone \"%s\" has been created successfully\n", bucketName, zoneName)
} else {
    fmt.Printf("Bucket creation failed with given message: %s\n", err)
}
```

`bucketService.Put()` will create a Bucket in the specified zone.

