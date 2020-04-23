# GetObject Example

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

Then set the input parameters used by the GetObject method (using GetObjectInput storage).

```go
	input := &service.GetObjectInput{}
```

Please note that the field in GetObjectInput is not necessarily required to be set. For details, please refer to [Official API Documentation](https://docs.qingcloud.com/qingstor/api/object/get).

Then call the GetObject method to download the object. objectKey Sets the filepath of the object to be fetched (in the current bucket).

```go
	// Please replace this file path with some file exists on your bucket.
	objectKey := "your-picture-uploaded.jpg"
	if output, err := bucketService.GetObject(objectKey, input); err != nil {
		fmt.Printf("Download object(%s) in bucket(name: %s) failed with given error: %s\n", objectKey, bucketName, err)
	} else {
		data, _ := ioutil.ReadAll(output.Body)
		_ = output.Close()
		err := ioutil.WriteFile("/tmp/picture_downloaded.jpg", data, 0644)
		if err != nil {
			panic(err)
		}
	}
```