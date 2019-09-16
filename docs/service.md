# Service Initialization

First, we need to initialize a QingStor service to call the services provided by QingStor.

```go
import (
	"github.com/yunify/qingstor-sdk-go/v3/config"
	"github.com/yunify/qingstor-sdk-go/v3/service"
)

const accessKeyId = "YOUR-ACCESS-KEY-ID"
const secretAccessKey = "YOUR--SECRET-ACCESS-KEY"

var conf, _ = config.New(accessKeyId, secretAccessKey)
var qingStor, _ = service.Init(conf)
var bucketService, _ = qingStor.Bucket("your-bucket-name", "zone-name")
```

The object that appears in the above code:
- The `conf` object carries the user's authentication information and configuration.
- The `qingStor` object is used to operate the QingStor object storage service, which is used to call all Service level APIs or to create a specified Bucket object to call Bucket and Object level APIs.
- The `bucketService` object is bound to the specified bucket and provides a series of object storage operations for the bucket.