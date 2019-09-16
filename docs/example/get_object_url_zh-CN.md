# 获取文件的下载地址

## 代码片段

使用您的 AccessKeyID 和 SecretAccessKey 初始化 Qingstor 对象。

```go
import (
	"github.com/yunify/qingstor-sdk-go/v3/config"
	"github.com/yunify/qingstor-sdk-go/v3/service"
	"github.com/yunify/qingstor-sdk-go/v3/utils"
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

然后设置 GetObjectRequest 方法用到的输入参数（使用 GetObjectInput 存储）。

```go
    input := &service.GetObjectInput{}
```

请注意 GetObjectInput 中 field 不是必须设置的，具体可以参考[官方 API 文档](https://docs.qingcloud.com/qingstor/api/object/get)。

然后你可以获得该对象的签名地址。objectKey 设置要获取的对象的 filepath（位于当前 bucket 中）。

```go
	// Please replace this file path with some file exists on your bucket.
	objectKey := "your-picture-uploaded.jpg"
	req, _, _ := bucketService.GetObjectRequest(objectKey, input)
	_ = req.Build()
	// the url expired after 600 sec.
	_ = req.SignQuery(600)
	fmt.Println(req.HTTPRequest.URL)
```

打印出的 url 是可以直接在浏览器中打开的，如果是浏览器支持预览的格式，浏览器会其进行预览，否则已默认文件名下载保存。
如果您想要设置保存的文件名，直接执行下载动作，可以进行如下设置：

```go
	encodedName := utils.URLQueryEscape("特殊?$&ab c=符号.jpg")
	fmt.Println(encodedName)
	disposition := fmt.Sprintf("attachment; filename=\"%s\"; filename*=utf-8''%s", encodedName, encodedName)
    input := &service.GetObjectInput{ResponseContentDisposition: &disposition}
```
