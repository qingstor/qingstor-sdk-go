# PUT Bucket CORS

## Request Elements

|      Name       |  Type   | Description                                                                                                                                                                        | Required |
| :-------------: | :-----: | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :------: |
|   cors_rules    |  Array  | A set of origins and methods (cross-origin access that you want to allow). The elements in each set of configuration items are explained as follows.                               |   Yes    |
| allowed_origin  | String  | An origin that you want to allow cross-domain requests from. This can contain at most one * wild character.                                                                        |   Yes    |
| allowed_methods |  Array  | An HTTP method that you want to allow the origin to execute. A combination of the following values can be specified: “GET”, “PUT”, “POST”, “DELETE”, “HEAD”, or use ‘*’ to set up. |   Yes    |
| allowed_headers |  Array  | An HTTP header that you want to allow the origin to execute. This can contain at most one * wild character.                                                                        |    No    |
| expose_headers  |  Array  | One or more headers in the response that you want customers to be able to access from their applications (for example, from a JavaScript XMLHttpRequest object).                   |    No    |
| max_age_seconds | Integer | The time in seconds that your browser is to cache the preflight response for the specified resource.(seconds)                                                                      |    No    |

See [API Docs](https://docsv4.qingcloud.com/user_guide/storage/object_storage/api/bucket/cors/put_cors/) for more information about request elements.

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

then you can PUT Bucket CORS

```go
	toPtr := func(s string) *string { return &s }
	maxAgeSeconds := 200
	body := service.PutBucketCORSInput{CORSRules: []*service.CORSRuleType{
		{AllowedHeaders: []*string{toPtr("x-qs-date"), toPtr("Content-Type"), toPtr("Content-MD5"), toPtr("Authorization")},
			AllowedMethods: []*string{toPtr("PUT"), toPtr("GET"), toPtr("DELETE"), toPtr("POST")},
			AllowedOrigin:  toPtr("http://*.qingcloud.com"),
			ExposeHeaders:  []*string{toPtr("x-qs-date")},
			MaxAgeSeconds:  &maxAgeSeconds,
		},
	}}
	if output, err := bucketService.PutCORS(&body); err != nil {
		fmt.Printf("Set CORS of bucket(name: %s) failed with given error: %s\n", bucketName, err)
	} else {
		fmt.Printf("The status code expected: 200(actually: %d)\n", *output.StatusCode)
	}
```