# 列取分段上传(List Multipart Uploads)

获取正在进行的分段上传对象的列表。当一个对象通过 Initiate Multipart 接口开启了分段上传模式，在调用 Complete Multipart 或 Abort Multipart 接口之前，该对象处于“正在进行分段上传”的状态，此对象将会出现在该接口返回的列表里。

与 GET Bucket (List Objects) 接口类似，用户可以通过传递 prefix, delimiter 请求参数，指定获取某个目录下面正在进行的分段上传。列表按照对象名称的 alphanumeric 顺序从小到大排序。如果同名对象有多个分段上传，翻页被截断后只显示了一部分，下次翻页可通过 upload_id_marker 参数，获取该 upload_id 往后按创建时间排序后剩下的分段上传。

如果用户只想获取某个对象已经上传的分段，请查阅 [API Docs](https://docs.qingcloud.com/qingstor/api/object/multipart/list_multipart.html#object-storage-api-list-multipart).

## 请求参数

在 List Bucket Objects 时添加筛选条件

参考[对应的 API 文档](https://docs.qingcloud.com/qingstor/api/bucket/list_multipart_uploads.html)，您可以在对应的 Input 设置并添加如下筛选条件：

|      参数名      |  类型   | 描述                                                                                                | 是否必要 |
| :--------------: | :-----: | :-------------------------------------------------------------------------------------------------- | :------: |
|      prefix      | String  | 限定返回的分段上传对象名必须以 prefix 作为前缀                                                      |    No    |
|    delimiter     |  Char   | 用于给对象名分组的字符。返回的对象名是从指定的 prefix 开始，到第一次出现 delimiter 字符之间的对象名 |    No    |
|    key_marker    | String  | 设定结果从 key 之后按字母排序的第一个分段上传开始返回                                               |    No    |
| upload_id_marker | String  | 设定结果从 upload_id 之后按时间排序的第一个分段上传开始返回                                         |    No    |
|      limit       | Integer | 限定此次返回的分段对象的最大数量，默认值为 200，最大允许设置 1000                                   |    No    |

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

对象创建完毕后，我们需要执行真正的列取分段上传(List Multipart Uploads)操作：

下面的操作罗列出 Movies 目录下(不包含子目录及其文件)所有尚未调用 Complete MultiUpload 接口的对象，限制 6 个。

```go
	toPtr := func(s string) *string { return &s }
	limit := 6
	resp, err := bucketService.ListMultipartUploads(&service.ListMultipartUploadsInput{
		Delimiter:      toPtr("/"),
		KeyMarker:      nil,
		Limit:          &limit,
		Prefix:         toPtr("Movies/"),
		UploadIDMarker: nil,
	})
	if err != nil {
		fmt.Printf("List Objects(multiUploaded and the complete api not been called) on bucket: %s failed with given error: %s\n", bucketName, err)
	} else {
		for _, objInfo := range resp.Uploads {
			b, _ := json.Marshal(objInfo)
			fmt.Println(string(b))
		}
	}
```
