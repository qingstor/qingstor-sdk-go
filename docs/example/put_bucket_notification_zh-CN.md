# PUT Bucket Notification

## 请求消息体

|      名称      |  类型  | 描述                                                                                                                                                                                                                           | 是否必要 |
| :------------: | :----: | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :------: |
| notifications  | Array  | bucket notification 的配置规则，配置项中的元素解释见下                                                                                                                                                                         |   Yes    |
|       id       | String | 通知配置的标识                                                                                                                                                                                                                 |   Yes    |
|  event_types   | Array  | 事件的类型，每当该类型的事件被触发时，发出通知。<br>目前支持的类型为: <br> - “create_object”: 创建对象完成 <br> - “delete_object”: 删除对象完成<br> - “abort_multipart”: 终止分段上传<br> - “complete_multipart”: 完成分段上传 |   Yes    |
| object_filters | Array  | 对象名匹配规则(glob patterns)                                                                                                                                                                                                  |    no    |
|   cloudfunc    | String | 事件处理云服务，接收通知中触发的事件并进行处理。目前支持:<br> - tupu-porn: 图谱鉴黄服务<br> - notifier: 通知服务, 将 QingStor 事件推送到 notify_url<br> - image: 图片基本处理服务                                              |   Yes    |
| cloudfunc_args | Object | 提供给 cloudfunc 的自定义参数                                                                                                                                                                                                  |    No    |
|   notify_url   | String | 通知事件处理结果的 url ，当事件处理完成后，会将处理结果以 POST 方式向 notify_url 请求。<br>如果 POST 超时，将会重试，超时时间是 5s， 重试间隔为 1s。                                                                           |    No    |

### 图片基本处理服务参数

当设置 cloudfunc 为 image 时, 需要设置 cloudfunc_args 为以下参数，对象存储将按照指定的图片处理规则对图片进行处理，并将结果另存回对象存储。

|    名称     |  类型  | 描述                                                                                                                 | 是否必要 |
| :---------: | :----: | :------------------------------------------------------------------------------------------------------------------- | :------: |
|   action    | String | 图片的具体操作参数, 见 [图片基本处理服务](https://docs.qingcloud.com/qingstor/data_process/image_process/index.html) |   Yes    |
| key_prefix  | String | 处理后 object 名称的前缀, 默认为 “gen”                                                                               |    No    |
| key_seprate | String | key_prefix 和 object 之间的分隔符，默认为 “_“                                                                        |    No    |
| save_bucket | String | 另存为的目标 bucket 名称，默认为当前 object 所在 bucket                                                              |    No    |

访问 [API Docs](https://docs.qingcloud.com/qingstor/api/bucket/notification/put_notification.html) 以查看更多关于请求消息体的信息。

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

然后您可以 PUT Bucket Notification.
下面的代码对 bucket 设置了每次有新对象创建都进行是否包含色情信息的检验，如果有的话，通知 notify_url 中设置的 url。

```go
	toPtr := func(s string) *string { return &s }
	body := service.PutBucketNotificationInput{Notifications: []*service.NotificationType{{
		Cloudfunc: toPtr("tupu-porn"), // 图普鉴黄
		// CloudfuncArgs: nil,
		EventTypes:    []*string{toPtr("create_object")},
		ID:            toPtr("notification-1"),
		NotifyURL:     toPtr("http://user_notify_url"),
		ObjectFilters: []*string{toPtr("*")},
	}}}
	if output, err := bucketService.PutNotification(&body); err != nil {
		fmt.Printf("Set notifications of bucket(name: %s) failed with given error: %s\n", bucketName, err)
	} else {
		fmt.Printf("The status code expected: 200(actually: %d)\n", *output.StatusCode)
	}
```