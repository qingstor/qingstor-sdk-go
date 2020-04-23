# 获取 Bucket 元信息(Head Bucket)

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

对象创建完毕后，我们需要执行真正的获取 Bucket 元信息操作：

```go
	if resp, err := bucketService.Head(); err != nil {
		fmt.Printf("The attempt to access a bucket(name: %s) failed with given error: %s\n", bucketName, err)
	} else {
		fmt.Printf("The status code expected: 200(actually: %d)\n", *resp.StatusCode)
	}
```

上面代码中出现的函数：
- `bucketService.Head()` 在 `pek3b` 区域尝试使用 HEAD 获取一个名为 `your-bucket-name` 的 Bucket 信息。 

上面代码中出现的对象：
- `resp` 对象是 `bucketService.Head()` 方法的返回值。
- `resp.StatusCode` 存储了 api 操作的 http 状态码。

