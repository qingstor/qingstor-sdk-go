# fetch 对象

## 代码片段

使用您的 AccessKeyID 和 SecretAccessKey 初始化 Qingstor 对象。

```go
import (
	"github.com/yunify/qingstor-sdk-go/v3/config"
	"github.com/yunify/qingstor-sdk-go/v3/service"
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

然后设置 PutObject 方法用到的输入参数（核心参数：XQSFetchSource）。

```go
	// Fetch source looks like this: "protocol://host[:port]/[path]"
	sourceLink := "https://www.qingcloud.com/static/assets/images/icons/common/footer_logo.svg"
	input := &service.PutObjectInput{
		XQSFetchSource: &sourceLink,
	}
```

请注意 PutObjectInput 中的 field 不是都必须设置的，具体可以参考[官方 API 文档](https://docs.qingcloud.com/qingstor/api/object/fetch)。

然后调用 PutObject 方法 fetch 对象。objectKey 设置 put 后的 filepath（位于当前 bucket 中）。

```go
	objectKey := "file-fetched/the_file_fetched.svg"
	if output, err := bucketService.PutObject(objectKey, input); err != nil {
		fmt.Printf("Fetch object from source link(%s) to target path(%s) failed with given error: %s\n", sourceLink, objectKey, err)
	} else {
		fmt.Printf("The status code expected: 201(actually: %d)\n", *output.StatusCode)
	}
```