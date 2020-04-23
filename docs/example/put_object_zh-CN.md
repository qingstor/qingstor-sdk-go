# 上传对象

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

然后设置 PutObject 方法可能用到的输入参数。

```go
	filepath := "/tmp/your-picture.jpg"
	file, _ := os.Open(filepath)
	defer func() {
		_ = file.Close()
	}()
	// Calculate MD5
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	hashInBytes := hash.Sum(nil)[:16]
	md5String := hex.EncodeToString(hashInBytes)
	toPtr := func(s string) *string { return &s }
	input := &service.PutObjectInput{
		ContentMD5:      toPtr(md5String),    // optional. You can manually calculate this to check uploaded file is intact or not.
		ContentType:     toPtr("image/jpeg"), // ContentType and ContentLength will be detected automatically if empty
		Body:            file,
		XQSStorageClass: toPtr("STANDARD"), // optional. default to be “STANDARD”. value can be "STANDARD" or “STANDARD_IA”.
	}
```

请注意 PutObjectInput 中 field 不是都必须设置的，具体可以参考[官方 API 文档](https://docs.qingcloud.com/qingstor/api/object/put)。

然后调用 PutObject 方法上传对象。objectKey 设置上传后的 filepath。

```go
	objectKey := "your-picture-uploaded.jpg"
	if output, err := bucketService.PutObject(objectKey, input); err != nil {
		fmt.Printf("Put object to bucket(name: %s) failed with given error: %s\n", bucketName, err)
	} else {
		fmt.Printf("The status code expected: 201(actually: %d)\n", *output.StatusCode)
	}
```