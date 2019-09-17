# PutObjects Example

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

Then set the input parameters that the PutObject method might use.

```go
	filepath := "/tmp/your-picture.jpg"
	file, _ := os.Open(filepath)
	defer func() {
		_ = file.Close()
	}()
	// Calculate MD5
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	hashInBytes := hash.Sum(nil)[:16]
	md5String := hex.EncodeToString(hashInBytes)
	toPtr := func(s string) *string { return &s }
	input := &service.PutObjectInput{
		ContentMD5:      toPtr(md5String),    // optional. You can manually calculate this to check uploaded file is intact or not.
		ContentType:     toPtr("image/jpeg"), // ContentType and ContentLength will be detected automatically if empty
		Body:            file,
		XQSStorageClass: toPtr("STANDARD"), // optional. default to be “STANDARD”. value can be "STANDARD" or “STANDARD_IA”.
	}
```

Please note that not all fields in PutObjectInput required to be set. For details, please refer to [Official API Documentation](https://docs.qingcloud.com/qingstor/api/object/put).

Then call the PutObject method to upload the object. objectKey Sets the filepath after uploading.

```go
	objectKey := "your-picture-uploaded.jpg"
	if output, err := bucketService.PutObject(objectKey, input); err != nil {
		fmt.Printf("Put object to bucket(name: %s) failed with given error: %s\n", bucketName, err)
	} else {
		fmt.Printf("The status code expected: 201(actually: %d)\n", *output.StatusCode)
	}
```