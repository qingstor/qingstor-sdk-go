# PUT Bucket External Mirror

## 请求消息体

|    名称     |  类型  | 描述                                                                                                                                                                                                                                           | 是否必要 |
| :---------: | :----: | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :------: |
| source_site | String | 外部镜像回源的源站。源站形式为 `<protocol>://<host>[:port]/[path]` 。 protocol的值可为 “http” 或 “https”，默认为 “http”。port 默认为 protocol 对应的端口。path 可为空。 如果存储空间多次设置不同的源站，该存储空间的源站采用最后一次设置的值。 |   Yes    |

访问 [API Docs](https://docs.qingcloud.com/qingstor/api/bucket/external_mirror/put_external_mirror.html) 以查看更多关于请求消息体的信息。

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

然后您可以 PUT Bucket External Mirror

```go
	sourceSite := "http://example.com:80/image/"
	body := service.PutBucketExternalMirrorInput{SourceSite: &sourceSite}
	if output, err := bucketService.PutExternalMirror(&body); err != nil {
		fmt.Printf("Set external mirror of bucket(name: %s) failed with given error: %s\n", bucketName, err)
	} else {
		fmt.Printf("The status code expected: 200(actually: %d)\n", *output.StatusCode)
	}
```