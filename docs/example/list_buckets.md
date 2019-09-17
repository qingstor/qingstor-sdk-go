# List Buckets

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

Then you can get all your Buckets

```go
resp, _ := qingStor.ListBuckets(nil)
if *resp.StatusCode == 200 {
    buckets := resp.Buckets // all buckets info
    // print bucket info in human readable format.
    for _, bucket := range buckets {
        b, _ := json.Marshal(bucket)
        fmt.Println(string(b))
    }
}
```