# GET Bucket Notification

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

对象创建完毕后，我们需要执行真正的获取 Bucket Notification 操作：

```go
	if output, err := bucketService.GetNotification(); err != nil {
		fmt.Printf("Get notifications of bucket(name: %s) failed with given error: %s\n", bucketName, err)
	} else {
		b, _ := json.Marshal(output.Notifications)
		fmt.Println("The notifications of this bucket: ", string(b))
	}
```