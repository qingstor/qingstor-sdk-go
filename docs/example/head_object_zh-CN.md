# 获取文件的元数据 

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

然后设置 HeadObject 方法用到的输入参数（使用 HeadObjectInput 存储）。

```go
	input := &service.HeadObjectInput{}
```

请注意 HeadObjectInput 中 field 不是必须设置的，具体可以参考[官方 API 文档](https://docsv4.qingcloud.com/user_guide/storage/object_storage/api/object/basic_opt/head/)。

然后调用 HeadObject 方法获取对象元信息，测试是否可以被访问。objectKey 设置要获取的对象的 filepath（位于当前 bucket 中）。

```go
	// Please replace this file path with some file exists on your bucket.
	objectKey := "your_file.zip"
	if output, err := bucketService.HeadObject(objectKey, input); err != nil {
		fmt.Printf("The attempt to access a object(name: %s) metadata failed with given error: %s\n", bucketName, err)
	} else {
		fmt.Printf("The status code expected: 200(actually: %d)\n", * output.StatusCode)
		b, _ := json.MarshalIndent(output, "", "\t")
		fmt.Printf("The metadata of object(%s):\n %s\n", objectKey, string(b))
	}
```

操作正确返回的话，响应状态码将会是 200。
