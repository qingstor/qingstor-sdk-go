# PUT Bucket Lifecycle

## Request Elements

|               Name                |  Type   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | Required |
| :-------------------------------: | :-----: | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :------: |
|       The elements of rule        |  List   | rule are Lifecycle rules. The rules are of type Dict and the valid keys are "id", "status", "filter", "expiration", "abort_incomplete_multipart_upload" and "transition". The total number of rules cannot exceed 100, and only one type of operation is allowed in each rule. The same bucket, prefix and support operations ( expiration, abort_incomplete_multipart_upload, transition) cannot be duplicated, otherwise return 400 invalid_request contains duplicate rule information see [Error Message] (https://docsv4.qingcloud.com/user_guide/storage/object_storage/api/error_code/). |   Yes    |
|                id                 | String  | Identifier of the rule. It can be any UTF-8 encoded character and cannot exceed 255 bytes in length. In a Bucket Lifecycle, the rule identifier must be unique. This string can be used to describe the purpose of the policy. If the id is repeated, it returns 400 invalid_request .                                                                                                                                                                                                                                                                                                                     |   Yes    |
|              status               | String  | The status of this rule. Its value can be either "enabled" (for effective) or "disabled" (for disabled).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |   Yes    |
|              filter               |  Dict   | is used to match the filter condition of Object. The valid key is “prefix”.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |   Yes    |
|              prefix               | String  | The prefix match policy is used to match the Object name, and the empty string means to match the Object in the entire Bucket. The default is an empty string.                                                                                                                                                                                                                                                                                                                                                                                                                                             |    No    |
|            expiration             |  Dict   | The rule for deleting an Object with a valid key of "days". "days" must be a positive integer, otherwise return 400 invalid_request. The object that matches the prefix is ​​deleted after the specified number of days (days) at the last modification time.                                                                                                                                                                                                                                                                                                                                              |    No    |
| abort_incomplete_multipart_upload |  Dict   | Rules for canceling unfinished multipart uploads. The valid key is "days_after_initiation". "days_after_initiation" must be a positive integer, otherwise return 400 invalid_request.                                                                                                                                                                                                                                                                                                                                                                                                                      |    No    |
|            transition             |  Dict   | The rule for changing the storage level. The valid keys are "days", "storage_class". Days must be >= 30, otherwise return 400 invalid_request. For objects that match the prefix (prefix), change to low frequency storage after the specified number of days (days) at the last modification time.                                                                                                                                                                                                                                                                                                        |    No    |
|               days                | Integer | Executes the operation after the specified number of days of the last modification time of the object.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |    No    |
|       days_after_initiation       | Integer | Performs after the specified number of days to initialize the segment upload.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |   Yes    |
|           storage_class           | Integer | The storage_class to be changed to, the supported values ​​are "STANDARD", "STANDARD_IA".                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |   Yes    |

See [API Docs](https://docsv4.qingcloud.com/user_guide/storage/object_storage/api/bucket/lifecycle/put_lifecycle/) for more information about request elements.

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

Then you can PUT Bucket Lifecycle.
The following code sets the log information under the bucket (stored in the logs/ directory) to automatically perform the delete operation after 180 days.

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
