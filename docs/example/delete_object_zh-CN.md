# 删除文件

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

然后调用 DeleteObject 方法删除对象。objectKey 设置要删除的对象的 filepath（位于当前 bucket 中）。

```go
	objectKey := "file_your_want_delete"
	if output, err := bucketService.DeleteObject(objectKey); err != nil {
		fmt.Printf("Delete object(name: %s) in bucket(%s) failed with given error: %s\n", objectKey, bucketName, err)
	} else {
		fmt.Printf("The status code expected: 204(actually: %d)\n", *output.StatusCode)
	}
```

操作正确返回的话，响应状态码将会是 204。