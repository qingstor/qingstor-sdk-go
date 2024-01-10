# 下载对象

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

然后设置 GetObject 方法用到的输入参数（使用 GetObjectInput 存储）。

```go
	input := &service.GetObjectInput{}
```

请注意 GetObjectInput 中 field 不是必须设置的，具体可以参考[官方 API 文档](https://docsv4.qingcloud.com/user_guide/storage/object_storage/api/object/basic_opt/get/)。

然后调用 GetObject 方法下载对象。objectKey 设置要获取的对象的 filepath（位于当前 bucket 中）。

```go
	// Please replace this file path with some file exists on your bucket.
	objectKey := "your-picture-uploaded.jpg"
	if output, err := bucketService.GetObject(objectKey, input); err != nil {
		fmt.Printf("Download object(%s) in bucket(name: %s) failed with given error: %s\n", objectKey, bucketName, err)
	} else {
		data, _ := ioutil.ReadAll(output.Body)
		_ = output.Close()
		err := ioutil.WriteFile("/tmp/picture_downloaded.jpg", data, 0644)
		if err != nil {
			panic(err)
		}
	}
```
