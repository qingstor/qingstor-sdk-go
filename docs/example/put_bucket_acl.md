# PutACL Example

## Request Elements

|    Name    |  Type  | Description                                                                                                                        |
| :--------: | :----: | :--------------------------------------------------------------------------------------------------------------------------------- |
|    acl     |  List  | Supports to set 0 or more grantees                                                                                                 |
|  grantee   |  Dict  | Specifies the Type(user, group). When type is user, need user id; when type is group, only supports QS_ALL_USERS(all of the users) |
| permission | String | Specifies the permission (READ, WRITE, FULL_CONTROL) given to the grantee.                                                         |

## Code Snippet

Initialize the Qingstor object with your AccessKeyID and SecretAccessKey.

```go
import (
	"github.com/yunify/qingstor-sdk-go/v3/config"
	"github.com/yunify/qingstor-sdk-go/v3/service"
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

then you can put bucket ACL

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