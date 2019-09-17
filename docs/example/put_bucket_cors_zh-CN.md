# PUT Bucket CORS

## 请求消息体

|      名称       |  类型   | 描述                                                                                                            | 是否必要 |
| :-------------: | :-----: | :-------------------------------------------------------------------------------------------------------------- | :------: |
|   cors_rules    |  Array  | 跨源的规则配置，每组配置项中的元素解释如下。                                                                    |   Yes    |
| allowed_origin  | String  | 用户所期望的跨源请求来源,可以用 ‘*’ 来进行通配。                                                                |   Yes    |
| allowed_methods |  Array  | 设置源所允许的 HTTP 方法。可指定以下值的组合: “GET”, “PUT”, “POST”, “DELETE”, “HEAD”, 或者使用 ‘*’ 来进行设置。 |   Yes    |
| allowed_headers |  Array  | 设置源所允许的 HTTP header 。 可以用 ‘*’ 来进行通配。                                                           |    No    |
| expose_headers  |  Array  | 设置客户能够从其应用程序（例如，从 JavaScript XMLHttpRequest 对象）进行访问的HTTP 响应头。                      |    No    |
| max_age_seconds | Integer | 设置在预检请求(Options)被资源、HTTP 方法和源识别之后，浏览器将为预检请求缓存响应的时间（以秒为单位）。          |    No    |

访问 [API Docs](https://docs.qingcloud.com/qingstor/api/bucket/cors/put_cors.html) 以查看更多关于请求消息体的信息。

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

然后您可以 PUT Bucket CORS

```go
	toPtr := func(s string) *string { return &s }
	maxAgeSeconds := 200
	body := service.PutBucketCORSInput{CORSRules: []*service.CORSRuleType{
		{AllowedHeaders: []*string{toPtr("x-qs-date"), toPtr("Content-Type"), toPtr("Content-MD5"), toPtr("Authorization")},
			AllowedMethods: []*string{toPtr("PUT"), toPtr("GET"), toPtr("DELETE"), toPtr("POST")},
			AllowedOrigin:  toPtr("http://*.qingcloud.com"),
			ExposeHeaders:  []*string{toPtr("x-qs-date")},
			MaxAgeSeconds:  &maxAgeSeconds,
		},
	}}
	if output, err := bucketService.PutCORS(&body); err != nil {
		fmt.Printf("Set CORS of bucket(name: %s) failed with given error: %s\n", bucketName, err)
	} else {
		fmt.Printf("The status code expected: 200(actually: %d)\n", *output.StatusCode)
	}
```