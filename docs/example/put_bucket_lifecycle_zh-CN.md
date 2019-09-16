# PUT Bucket Lifecycle

## 请求消息体

|               名称                |  类型   | 描述                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | 是否必要 |
| :-------------------------------: | :-----: | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :------: |
|               rule                |  List   | rule 的元素为 Lifecycle 规则。规则为 Dict 类型，有效的键为 “id”、”status”、”filter”、”expiration”、”abort_incomplete_multipart_upload” 和 “transition”。规则总数不能超过 100 条，且每条规则中只允许存在一种类型的操作。同一 bucket, prefix 和 支持操作（ expiration, abort_incomplete_multipart_upload, transition) 不能有重复，否则返回 400 invalid_request 包含重复的规则信息 [参见错误信息](https://docs.qingcloud.com/qingstor/api/common/error_code.html#object-storage-error-code)。 |   Yes    |
|                id                 | String  | 规则的标识符。可为任意 UTF-8 编码字符，长度不能超过 255 个字节，在一个 Bucket Lifecycle 中，规则的标识符必须唯一。该字符串可用来描述策略的用途。如果 id 有重复，会返回 400 invalid_request 。                                                                                                                                                                                                                                                                                              |   Yes    |
|              status               | String  | 该条规则的状态。其值可为 “enabled” (表示生效) 或 “disabled” (表示禁用)。                                                                                                                                                                                                                                                                                                                                                                                                                   |   Yes    |
|              filter               |  Dict   | 用于匹配 Object 的过滤条件，有效的键为 “prefix”。                                                                                                                                                                                                                                                                                                                                                                                                                                          |   Yes    |
|              prefix               | String  | 前缀匹配策略，用于匹配 Object 名称，空字符串表示匹配整个 Bucket 中的 Object。默认值为空字符串。                                                                                                                                                                                                                                                                                                                                                                                            |    No    |
|            expiration             |  Dict   | 用于删除 Object 的规则，有效的键为 “days”。”days” 必须是正整数，否则返回 400 invalid_request。对于匹配前缀（prefix) 的对象在最后修改时间的指定天数（days）后删除该对象。                                                                                                                                                                                                                                                                                                                   |    No    |
| abort_incomplete_multipart_upload |  Dict   | 用于取消未完成的分段上传的规则，有效的键为 “days_after_initiation”。”days_after_initiation” 必须是正整数，否则返回 400 invalid_request。                                                                                                                                                                                                                                                                                                                                                   |    No    |
|            transition             |  Dict   | 用于变更存储级别的规则，有效的键为 “days”, “storage_class”。days 必须 >= 30, 否则返回 400 invalid_request。对于匹配前缀（prefix) 的对象在最后修改时间的指定天数（days）后变更到低频存储。                                                                                                                                                                                                                                                                                                  |    No    |
|               days                | Integer | 在对象最后修改时间的指定天数后执行操作。                                                                                                                                                                                                                                                                                                                                                                                                                                                   |    No    |
|       days_after_initiation       | Integer | 在初始化分段上传的指定天数后执行操作。                                                                                                                                                                                                                                                                                                                                                                                                                                                     |   Yes    |
|           storage_class           | Integer | 要变更至的 storage_class，支持的值为 "STANDARD"、"STANDARD_IA"。                                                                                                                                                                                                                                                                                                                                                                                                                           |   Yes    |

访问 [API Docs](https://docs.qingcloud.com/qingstor/api/bucket/lifecycle/put_lifecycle.html) 以查看更多关于请求消息体的信息。

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

然后您可以 PUT Bucket Lifecycle.
下面的代码设置 bucket 下的日志信息（存放于 logs/ 目录）在 180 天后自动执行删除操作。

```go
	toPtr := func(s string) *string { return &s }
	expireDays := 180
	// choose (expiration, transition, abort_incomplete_multipart_upload) to execute different task.
	body := service.PutBucketLifecycleInput{Rule: []*service.RuleType{{
		Expiration: &service.ExpirationType{Days: &expireDays},
		Filter:     &service.FilterType{Prefix: toPtr("logs/")},
		ID:         toPtr("delete-logs"),
		Status:     toPtr("enabled"),
	}}}
	if output, err := bucketService.PutLifecycle(&body); err != nil {
		fmt.Printf("Set life cycles of bucket(name: %s) failed with given error: %s\n", bucketName, err)
	} else {
		fmt.Printf("The status code expected: 200(actually: %d)\n", *output.StatusCode)
	}
```