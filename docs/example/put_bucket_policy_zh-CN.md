# PUT Bucket Policy

## 请求消息体

访问 [API Docs](https://docs.qingcloud.com/qingstor/api/bucket/policy/put_policy.html) 以查看更多关于请求消息体的信息。

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

然后您可以 put bucket Policy

```go
	toPtr := func(s string) *string { return &s }
	strFilter := service.StringLikeType{Referer: []*string{toPtr("*.example1.com"), toPtr("*.example2.com")}}
	conds := service.ConditionType{
		StringLike: &strFilter,
	}
	// here we set only two policy statement. If you want to set two or more, notice that request will be matched by sort order.
	// if matching success, other statement will not be checked and request will be executed.
	stmts := []*service.StatementType{
		{
			Action:    []*string{toPtr("get_object")},
			Condition: &conds,
			Effect:    toPtr("allow"),
			ID:        toPtr("allow certain site to get objects"),
			Resource:  []*string{toPtr(bucketName + "/*")},
			User:      []*string{toPtr("*")}, // match all users.
		},
		{
			Action: []*string{toPtr("list_objects"), toPtr("create_object")},
			Effect:   toPtr("allow"),
			ID:       toPtr("allow user(id: usr-Xxxxx) to list objects and create objects"),
			Resource: []*string{toPtr(bucketName + "/*")},
			User:     []*string{toPtr("usr-Xxxxx")}, // replace with a real user.
		},
	}
	body := service.PutBucketPolicyInput{Statement: stmts}
	if output, err := bucketService.PutPolicy(&body); err != nil {
		fmt.Printf("Set policy of bucket(name: %s) failed with given error: %s\n", bucketName, err)
	} else {
		fmt.Printf("The status code expected: 200(actually: %d)\n", *output.StatusCode)
	}
```