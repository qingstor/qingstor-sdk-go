# OPTIONS Object

## Request Headers

|              Name              |  Type  | Description                                                                     | Required |
| :----------------------------: | :----: | :------------------------------------------------------------------------------ | :------: |
|             Origin             | String | Identifies the origin of the cross-origin request.                              |   Yes    |
| Access-Control-Request-Method  | String | Identifies what HTTP method will be used in the actual request.                 |   Yes    |
| Access-Control-Request-Headers | String | A comma-delimited list of HTTP headers that will be sent in the actual request. |    No    |

See [API Docs](https://docs.qingcloud.com/qingstor/api/object/options.html) for more information about request headers.

## Response Headers

|             Name              |  Type  | Description                                                                                                                                                                                                                                                                     |
| :---------------------------: | :----: | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
|  Access-Control-Allow-Origin  | String | The origin you sent in your request. If the origin in your request is not allowed, QingStor will not include this header in the response.                                                                                                                                       |
|    Access-Control-Max-Age     | String | How long, in seconds, the results of the preflight request can be cached.                                                                                                                                                                                                       |
| Access-Control-Allow-Methods  | String | The HTTP method that was sent in the original request. If the method in the request is not allowed, QingStor will not include this header in the response.                                                                                                                      |
| Access-Control-Allow-Headers  | String | A comma-delimited list of HTTP headers that the browser can send in the actual request. If any of the requested headers is not allowed, QingStor will not include that header in the response, nor will the response contain any of the headers with the Access-Control prefix. |
| Access-Control-Expose-Headers | String | A comma-delimited list of HTTP headers. This header provides the JavaScript client with access to these headers in the response to the actual request.                                                                                                                          |

See [API Docs](https://docs.qingcloud.com/qingstor/api/object/options.html) for more information about response headers.

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

Then set the input parameters used by the OptionsObject method (stored in OptionsObjectInput).

```go
	toPtr := func(s string) *string { return &s }
	// OptionsObject will filter allowed options
	input := &service.OptionsObjectInput{
		AccessControlRequestHeaders: toPtr("content-length,content-type"),
		AccessControlRequestMethod:  toPtr("DELETE,GET,PUT,PATCH"),
		Origin:                      toPtr("http://*.qingcloud.com"),
	}
```

Please note that not all fields in OptionsObjectInput required to be set. For details, please refer to [Official API Documentation](https://docs.qingcloud.com/qingstor/api/object/options).

Then call the OptionsObject method to download the object. objectKey Sets the filepath of the object to be options (in the current bucket).

```go
	// Please replace this file path with some file exists on your bucket.
	objectKey := "your-picture-your_file_test_options.zip"
	if output, err := bucketService.OptionsObject(objectKey, input); err != nil {
		fmt.Printf("The attempt to get allowed options a object(name: %s) failed with given error: %s\n", bucketName, err)
	} else {
		fmt.Printf("The status code expected: 200(actually: %d)\n", * output.StatusCode)
		b, _ := json.MarshalIndent(output, "", "\t")
		fmt.Printf("The allowed options of object(%s):\n %s\n", objectKey, string(b))
	}
```

The response will return all allowed operations (including header, method, origin) after filtering.