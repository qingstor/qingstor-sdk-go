# GetDownObjectMulti Example

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

Please note that the field in GetObjectInput is not necessarily required to be set. The parameter that must be set manually here is the Range parameter. For details, please refer to [Official API Documentation](https://docs.qingcloud.com/qingstor/api/object/get).

Then call the GetObject method to download the object and set 5M as the segment size. objectKey sets the filepath of the object to be fetched (in the current bucket).

```go
	objectKey := "your_zip_fetch_with_seg.zip"
	partSize := 1024 * 1024 * 5 // 5M every part.
	for i := 0; ; i++ {
		lo := partSize * i
		hi := partSize*(i+1) - 1
		byteRange := fmt.Sprintf("bytes=%d-%d", lo, hi)
		input := &service.GetObjectInput{
			Range: &byteRange,
		}
		output, err := bucketService.GetObject(objectKey, input)
		if err != nil {
			panic(err)
		}
		f, _ := os.OpenFile("/tmp/"+objectKey, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		data, _ := ioutil.ReadAll(output.Body)
		_, err = f.Write(data)
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
		if len(data) < partSize {
			break
		}
	}
```

The file will be saved to /tmp/{objectKey} (replace with your objectKey).