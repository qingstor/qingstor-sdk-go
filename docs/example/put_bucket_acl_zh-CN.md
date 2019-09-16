# Put ACL 示例

## 请求消息体

|    名称    |  类型  | 描述                                                                                                                        |
| :--------: | :----: | :-------------------------------------------------------------------------------------------------------------------------- |
|    acl     |  List  | 支持设置 0 到多个被授权者                                                                                                   |
|  grantee   |  Dict  | 支持 user, group 两种类型，当设置 user 类型时，需要给出 user id；当设置 group 类型时，目前只支持 QS_ALL_USERS，代表所有用户 |
| permission | String | 支持三种权限为 READ, WRITE, FULL_CONTROL                                                                                    |

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

然后您可以 put bucket ACL

```go
	toPtr := func(s string) *string { return &s }
	acls := []*service.ACLType{
		{
			Grantee: &service.GranteeType{
				ID:   toPtr("usr-XxXxXxX"), // your should assign a real user id to this variable.
				Type: toPtr("user"),
			},
			Permission: toPtr("WRITE"),
		},
		{
			Grantee: &service.GranteeType{
				Name: toPtr("QS_ALL_USERS"),
				Type: toPtr("group"),
			},
			Permission: toPtr("READ"),
		},
	}
	body := service.PutBucketACLInput{ACL: acls}
	if output, err := bucketService.PutACL(&body); err != nil {
		fmt.Printf("Set acl of bucket(name: %s) failed with given error: %s\n", bucketName, err)
	} else {
		fmt.Printf("The status code expected: 200(actually: %d)\n", *output.StatusCode)
	}
```