# Bucket 使用统计

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

对象创建完毕后，我们需要执行真正的 Bucket 使用统计的操作：

```go
	if resp, err := bucketService.GetStatistics(); err != nil {
		fmt.Printf("Get bucket(name: %s) statistics failed with given error: %s\n", bucketName, err)
	} else {
		fmt.Printf("Bucket info: {size: %d, count: %d, location: %s, url: %s, created: %s}\n",
			*resp.Size, *resp.Count, *resp.Location, *resp.URL, *resp.Created)
	}
```
