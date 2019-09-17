# GET Bucket ACL

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

After created the object, we need perform the action to get Bucket ACL：

```go
	if output, err := bucketService.GetACL(); err != nil {
		fmt.Printf("Get acl of bucket(name: %s) failed with given error: %s\n", bucketName, err)
	} else {
		fmt.Printf("The owner of this bucket is %s\n", *output.Owner.ID)
		b, _ := json.Marshal(output.ACL)
		fmt.Println("The acl info of this bucket: ", string(b))
	}
```