# 删除多个对象

## 代码片段

使用您的 AccessKeyID 和 SecretAccessKey 初始化 Qingstor 对象。

```go
import (
	"github.com/qingstor/qingstor-sdk-go/v4/config"
	"github.com/qingstor/qingstor-sdk-go/v4/service"
)

var conf, _ = config.New("YOUR-ACCESS-KEY-ID", "YOUR--SECRET-ACCESS-KEY")
var qingStor, _ = service.Init(conf)
```

然后根据要操作的 bucket 信息（zone, bucket name）来初始化 Bucket。

```go
	bucketName := "your-bucket-name"
	zoneName := "pek3b"
	bucketService, _ := qingStor.Bucket(bucketName, zoneName)
```

然后设置 DeleteObject 方法用到的输入参数（使用 DeleteMultipleObjectsInput 存储）。`Quiet` 指定是否返回被删除的对象列表。

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

请注意 DeleteMultipleObjectsInput 中的 field 不是都必须设置的，具体可以参考[官方 API 文档](https://docs.qingcloud.com/qingstor/api/bucket/delete_multiple)。

然后调用 DeleteMultipleObjects 方法删除对象。objectKey 设置要删除的对象的 filepath（位于当前 bucket 中）。

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

操作正确返回的话，响应状态码将会是 200。