# 大文件分段下载

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

请注意 GetObjectInput 中 field 不是必须设置的，这里必须手动设置的参数是 Range 参数。具体可以参考[官方 API 文档](https://docs.qingcloud.com/qingstor/api/object/get)。

然后调用 GetObject 方法下载对象，将 5M 设置为分段大小。objectKey 设置要获取的对象的 filepath（位于当前 bucket 中）。

```go
	objectKey := "your_zip_fetch_with_seg.zip"
	partSize := 1024 * 1024 * 5 // 5M every part.
	for i := 0; ; i++ {
		lo := partSize * i
		hi := partSize*(i+1) - 1
		byteRange := fmt.Sprintf("bytes=%d-%d", lo, hi)
		input := &service.GetObjectInput{
			Range: &byteRange,
		}
		output, err := bucketService.GetObject(objectKey, input)
		if err != nil {
			panic(err)
		}
		f, _ := os.OpenFile("/tmp/"+objectKey, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		data, _ := ioutil.ReadAll(output.Body)
		_, err = f.Write(data)
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
		if len(data) < partSize {
			break
		}
	}
```

文件将被保存至 /tmp/{objectKey}（请替换为您的 objectKey）。