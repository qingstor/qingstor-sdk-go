# 创建一个 Bucket

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

按照您设定的 bucket 名称初始化一个 Bucket 对象，以进行后续创建操作：

```go
bucketName := "your-bucket-name"
zoneName := "pek3b"
bucketService, _ := qingStor.Bucket(bucketName, zoneName)
```

`bucketService` 对象用于操作 Bucket，可以使用所有 Bucket 和 Object 级别的 API。现在执行真正的创建 Bucket 操作：

```go
if _, err := bucketService.Put(); err == nil {
    fmt.Printf("Your bucket named \"%s\" in zone \"%s\" has been created successfully\n", bucketName, zoneName)
} else {
    fmt.Printf("Bucket creation failed with given message: %s\n", err)
}
```

`bucketService.Put()` 会在指定 zone 创建 Bucket。 

