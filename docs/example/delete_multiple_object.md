# DeleteMultipleObjects Example

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

Then set the input parameters used by the DeleteObject method (stored in DeleteMultipleObjectsInput). `Quiet` Specifies whether to return a list of deleted objects.

```go
	objects := []string{"file_will_be_delete.jpg", "file_will_be_delete.zip"}
	var keys []*service.KeyType
	for _, objKey := range objects {
		key := objKey
		keys = append(keys, &service.KeyType{
			Key: &key,
		})
	}
	returnDeleteRes := false
	input := &service.DeleteMultipleObjectsInput{
		Objects: keys,
		Quiet:   &returnDeleteRes,
	}
```

Please note that not all fields in DeleteMultipleObjectsInput required to be set. For details, please refer to [Official API Documentation](https://docs.qingcloud.com/qingstor/api/bucket/delete_multiple).

Then call the DeleteMultipleObjects method to delete the object. objectKey Sets the filepath of the object to be deleted (in the current bucket).

```go
	if output, err := bucketService.DeleteMultipleObjects(input); err != nil {
		fmt.Printf("Delete objects(name: %v) failed with given error: %s\n", objects, err)
	} else {
		fmt.Printf("The status code expected: 200(actually: %d)\n", *output.StatusCode)
		fmt.Println("=========== objects been deleted ===========")
		for _, keyType := range output.Deleted {
			fmt.Println(*keyType.Key)
		}
	}
```

If the operation returns correctly, the response status code will be 200.