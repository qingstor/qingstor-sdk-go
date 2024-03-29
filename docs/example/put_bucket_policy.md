# PUT Bucket Policy

## Request Elements

See [API Docs](https://docsv4.qingcloud.com/user_guide/storage/object_storage/api/bucket/policy/put_policy/) for more information about request elements.

## Code Snippet

Initialize the Qingstor object with your AccessKeyID and SecretAccessKey.

```go
import (
	"github.com/qingstor/qingstor-sdk-go/v4/config"
	"github.com/qingstor/qingstor-sdk-go/v4/service"
)

var conf, _ = config.New("YOUR-ACCESS-KEY-ID", "YOUR--SECRET-ACCESS-KEY")
var qingStor, _ = service.Init(conf)
```

Initialize a Bucket object according to the bucket name you set for subsequent creation:

```go
bucketName := "your-bucket-name"
zoneName := "pek3b"
bucketService, _ := qingStor.Bucket(bucketName, zoneName)
```

then you can put Bucket Policy

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