# 对象拷贝(PUT Object - Copy)

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

然后设置 PutObject 方法用到的输入参数（核心参数：XQSCopySource）。

```go
	// Please replace this path with some file exists on your bucket.
	sourcePath := "/your-bucket-name/your-picture-uploaded.jpg"
	input := &service.PutObjectInput{
		XQSCopySource: &sourcePath,
	}
```

请注意 PutObjectInput 中 field 不是都必须设置的，具体可以参考[官方 API 文档](https://docs.qingcloud.com/qingstor/api/object/copy)。

然后调用 PutObject 方法拷贝对象。objectKey 设置拷贝后的 filepath（位于当前 bucket 中）。

```go
	if output, err := bucketService.PutObject(objectKey, input); err != nil {
		fmt.Printf("Copy object from source storage space(%s) to target path(%s) failed with given error: %s\n", sourcePath, objectKey, err)
	} else {
		fmt.Printf("The status code expected: 201(actually: %d)\n", *output.StatusCode)
	}
```