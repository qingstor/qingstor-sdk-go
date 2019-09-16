# OPTIONS Object

## api 信息

### 请求头(Request Headers)

|              名称              |  类型  | 描述                                           | 是否必要 |
| :----------------------------: | :----: | :--------------------------------------------- | :------: |
|             Origin             | String | 跨源请求的源。                                 |   Yes    |
| Access-Control-Request-Method  | String | 跨源请求的 HTTP method 。                      |   Yes    |
| Access-Control-Request-Headers | String | 跨源请求中的 HTTP headers (逗号分割的字符串)。 |    No    |

访问 [API Docs](https://docs.qingcloud.com/qingstor/api/object/options.html) 以查看更多关于请求头的信息。

### 响应头(Response Headers)

|             名称              |  类型  | 描述                                                                                           |
| :---------------------------: | :----: | :--------------------------------------------------------------------------------------------- |
|  Access-Control-Allow-Origin  | String | 跨源请求所允许的源。如果跨源请求没有被允许，该头信息将不会存在于响应头中。                     |
|    Access-Control-Max-Age     | String | 预检请求的结果被缓存的时间（单位为秒）。                                                       |
| Access-Control-Allow-Methods  | String | 跨源请求中的 HTTP method 。如果跨源请求没有被允许，该头信息将不会存在于响应头中。              |
| Access-Control-Allow-Headers  | String | 跨源请求中可以被允许发送的 HTTP headers (逗号分割的字符串)。                                   |
| Access-Control-Expose-Headers | String | 跨源请求的响应中,客户端（如 JavaScript Client） 可以获取到的 HTTP headers (逗号分割的字符串)。 |

访问 [API Docs](https://docs.qingcloud.com/qingstor/api/object/options.html) 以查看更多关于响应头的信息。

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

然后设置 OptionsObject 方法用到的输入参数（使用 OptionsObjectInput 存储）。

```go
	toPtr := func(s string) *string { return &s }
	// OptionsObject will filter allowed options
	input := &service.OptionsObjectInput{
		AccessControlRequestHeaders: toPtr("content-length,content-type"),
		AccessControlRequestMethod:  toPtr("DELETE,GET,PUT,PATCH"),
		Origin:                      toPtr("http://*.qingcloud.com"),
	}
```

请注意 OptionsObjectInput 中 field 不是都必须设置的，具体可以参考[官方 API 文档](https://docs.qingcloud.com/qingstor/api/object/options)。

然后调用 OptionsObject 方法下载对象。objectKey 设置要 options 的对象的 filepath（位于当前 bucket 中）。

```go
	// Please replace this file path with some file exists on your bucket.
	objectKey := "your-picture-your_file_test_options.zip"
	if output, err := bucketService.OptionsObject(objectKey, input); err != nil {
		fmt.Printf("The attempt to get allowed options a object(name: %s) failed with given error: %s\n", bucketName, err)
	} else {
		fmt.Printf("The status code expected: 200(actually: %d)\n", * output.StatusCode)
		b, _ := json.MarshalIndent(output, "", "\t")
		fmt.Printf("The allowed options of object(%s):\n %s\n", objectKey, string(b))
	}
```

响应将返回过滤过后的所有被允许的操作（包括 header, method, origin）。