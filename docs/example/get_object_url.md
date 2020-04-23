# GET Object Download Url Example

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

Then set the input parameters used by the GetObjectRequest method (stored in GetObjectInput).

```go
    input := &service.GetObjectInput{}
```

Please note that the field in GetObjectInput is not required to be set. For details, please refer to [Official API Documentation](https://docs.qingcloud.com/qingstor/api/object/get).

Then you can get the signature address of the object. objectKey Sets the filepath of the object to be fetched (in the current bucket).

```go
	// Please replace this file path with some file exists on your bucket.
	objectKey := "your-picture-uploaded.jpg"
	req, _, _ := bucketService.GetObjectRequest(objectKey, input)
	_ = req.Build()
	// the url expired after 600 sec.
	_ = req.SignQuery(600)
	fmt.Println(req.HTTPRequest.URL)
```

The printed url can be opened directly in the browser. If the browser supports the preview format, the browser will preview it, otherwise it will be downloaded and saved with the default file name.
If you want to set the saved file name and execute the download directly, you can set as the following code:

```go
	encodedName := utils.URLQueryEscape("特殊?$&ab c=符号.jpg")
	fmt.Println(encodedName)
	disposition := fmt.Sprintf("attachment; filename=\"%s\"; filename*=utf-8''%s", encodedName, encodedName)
    input := &service.GetObjectInput{ResponseContentDisposition: &disposition}
```
