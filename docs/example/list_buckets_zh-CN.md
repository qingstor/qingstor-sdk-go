# 列取 Buckets

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

然后您可以获取您所有的 Buckets

```go
resp, _ := qingStor.ListBuckets(nil)
if *resp.StatusCode == 200 {
    buckets := resp.Buckets // all buckets info
    // print bucket info in human readable format.
    for _, bucket := range buckets {
        b, _ := json.Marshal(bucket)
        fmt.Println(string(b))
    }
}
```