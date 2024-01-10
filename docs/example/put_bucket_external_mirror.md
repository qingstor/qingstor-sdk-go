# PUT Bucket External Mirror

## Request Elements

|    Name     |  Type  | Description                                                                                                                                                                                                                                                                                                                                                                          | Required |
| :---------: | :----: | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :------: |
| source_site | String | Source site of external mirror source. Source site is like this: `<protocol>://<host>[:port]/[path]` . Valid values of protocol: “http” or “https”, default “http”. Port defaults to the port corresponding to the protocol. Path can be empty. If the storage space has multiple source sites for many times, the source site of the storage space will use the last setting value. |   Yes    |

See [API Docs](https://docsv4.qingcloud.com/user_guide/storage/object_storage/api/bucket/external_mirror/put_external_mirror/) for more information about request elements.

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

then you can PUT Bucket External Mirror


```go
	sourceSite := "http://example.com:80/image/"
	body := service.PutBucketExternalMirrorInput{SourceSite: &sourceSite}
	if output, err := bucketService.PutExternalMirror(&body); err != nil {
		fmt.Printf("Set external mirror of bucket(name: %s) failed with given error: %s\n", bucketName, err)
	} else {
		fmt.Printf("The status code expected: 200(actually: %d)\n", *output.StatusCode)
	}
```