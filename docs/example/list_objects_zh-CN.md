# 列取对象

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

然后您可以得到 Bucket 中的所有对象的记录

```go
    resp, err := bucketService.ListObjects(&service.ListObjectsInput{})
    if err != nil {
        fmt.Printf("List Objects on bucket: %s failed with given error: %s\n", bucketName, err)
    } else {
        for _, objInfo := range resp.Keys {
            b, _ := json.Marshal(objInfo)
            fmt.Println(string(b))
        }
    }
```

还可以在 List Bucket Objects 时添加筛选条件

参考[对应的 API 文档](https://docsv4.qingcloud.com/user_guide/storage/object_storage/api/bucket/basic_opt/get/)，您可以在对应的 Input 设置并添加如下筛选条件：

|   参数名称    |   类型    |                                   描述                                    | 是否必须 |
|:---------:|:-------:|:-----------------------------------------------------------------------:|:----:|
|  prefix   | String  |                    限定返回的 object key 必须以 prefix 作为前缀                     |  否   |
| delimiter |  Char   | 是一个用于对 Object 名字进行分组的字符。所有名字包含指定的前缀且第一次出现 delimiter 字符之间的 object 作为一组元素 |  否   |
|  marker   | String  |                      设定结果从 marker 之后按字母排序的第一个开始返回                       |  否   |
|   limit   | Integer |                限定此次返回 object 的最大数量，默认值为 200，最大允许设置 1000                 |  否   |

以下代码是展示 Bucket 内 *test* 文件夹的所有对象（不包含子文件夹），默认以文件名排序。

```go
// List objects return objects which start with `prefix` and behind an object named `marker`.
// objects are located in the bucket bound with `bucketService`.
// The records returned is limit to 10.
func listObjects(bucketService *service.Bucket, prefix string, marker string) *string {
	delimiter := "/"
	limit := 10
	params := &service.ListObjectsInput{
		Delimiter: &delimiter,
		Limit:     &limit,
		Marker:    &marker,
		Prefix:    &prefix,
	}
	resp, err := bucketService.ListObjects(params)
	if err != nil {
		fmt.Printf("List Objects on bucket: %s failed with given error: %s\n", *bucketService.Properties.BucketName, err)
		return nil
	} else {
		fmt.Println("=============List Objects=============")
		for _, objInfo := range resp.Keys {
			b, _ := json.Marshal(objInfo)
			fmt.Println(string(b))
		}
		fmt.Println("=====End======")
		return resp.NextMarker
	}
}
```

如返回值不为空，说明还有下一页数据，可以继续访问。下面是一个调用示例：

```go
	bucketName := "your-bucket-name"
	zoneName := "pek3b"
	bucketService, _ := qingStor.Bucket(bucketName, zoneName)
	nextMarker := listObjects(bucketService, "test/", "")
	for nextMarker != nil && *nextMarker != "" { // result have next page
		nextMarker = listObjects(bucketService, "test/", *nextMarker)
	}
```