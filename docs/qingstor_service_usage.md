# QingStor Service Usage Guide

Import the QingStor and initialize service with a config, and you are ready to use the initialized service. Service only contains one API, and it is "ListBuckets".
To use bucket related APIs, you need to initialize a bucket from service using "Bucket" function.

Each API function take a Input struct and return an Output struct. The Input struct consists of request params, request headers, request elements and request body, and the Output holds the HTTP status code, QingStor request ID, response headers, response elements, response body and error (if error occurred).

You can use a specified version of a service by import a service package with a date suffix.

``` go
import (
	// Import the latest version API
	"github.com/yunify/qingstor-sdk-go/qingstor"
)
```

### Code Snippet

Initialize the QingStor service with a configuration

``` go
qsService, _ := qingstor.Init(configuration)
```

List buckets

``` go
qsOutput, _ := qsService.ListBuckets(nil)

// Print the HTTP status code.
// Example: 200
fmt.Println(qsOutput.StatusCode)

// Print the bucket count.
// Example: 5
fmt.Println(qsOutput.Count)

// Print the name of first bucket.
// Example: "test-bucket"
fmt.Println(qsOutput.Buckets[0].Name)
```

Initialize a QingStor bucket

``` go
bucket, _ := qsService.Bucket("test-bucket", "pek3a")
```

List objects in the bucket

``` go
bOutput, _ := bucket.ListObjects(nil)

// Print the HTTP status code.
// Example: 200
fmt.Println(bOutput.StatusCode)

// Print the key count.
// Example: 7
fmt.Println(len(bOutput.Keys))
```

Set ACL of the bucket

``` go
bACLOutput, _ := bucket.PutACL(&qingstor.PutBucketACLInput{
	ACL: []*qingstor.ACLType{{
		Grantee: &qingstor.GranteeType{
			Type: "user",
			ID:   "usr-xxxxxxxx",
		},
		Permission: "FULL_CONTROL",
	}},
})

// Print the HTTP status code.
// Example: 200
fmt.Println(bACLOutput.StatusCode)
```

Put object

``` go
// Open file
file, _ := os.Open("~/Desktop/Screenshot.jpg")
defer file.Close()

// Calculate MD5
hash := md5.New()
io.Copy(hash, file)
hashInBytes := hash.Sum(nil)[:16]
md5String := hex.EncodeToString(hashInBytes)

// Put object
oOutput, _ := bucket.PutObject(
	"Screenshot.jpg",
	&qingstor.PutObjectInput{
		ContentLength: 102475,       // Obtain automatically if empty
		ContentType:   "image/jpeg", // Detect automatically if empty
		ContentMD5:    md5String,
		Body:          file,
	},
)

// Print the HTTP status code.
// Example: 201
fmt.Println(oOutput.StatusCode)
```

Delete object

``` go
oOutput, _ := bucket.DeleteObject("Screenshot.jpg")

// Print the HTTP status code.
// Example: 204
fmt.Println(oOutput.StatusCode)
```

Initialize Multipart Upload

``` go
aOutput, _ := bucket.InitiateMultipartUpload(
	"QingCloudInsight.mov",
	&qingstor.InitiateMultipartUploadInput{
		ContentType: "video/quicktime",
	},
)

// Print the HTTP status code.
// Example: 200
fmt.Println(aOutput.StatusCode)

// Print the upload ID.
// Example: "9d37dd6ccee643075ca4e597ad65655c"
fmt.Println(aOutput.UploadID)
```

Upload Multipart

``` go
aOutput, _ := bucket.UploadMultipart(
	"QingCloudInsight.mov",
	&qingstor.UploadMultipartInput{
		UploadID:   "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		PartNumber: 0,
		ContentMD5: md5String0,
		Body:       file0,
	},
)

// Print the HTTP status code.
// Example: 201
fmt.Println(aOutput.StatusCode)

aOutput, _ = bucket.UploadMultipart(
	"QingCloudInsight.mov",
	&qingstor.UploadMultipartInput{
		UploadID:   "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		PartNumber: 1,
		ContentMD5: md5String1,
		Body:       file1,
	},
)

// Print the HTTP status code.
// Example: 201
fmt.Println(aOutput.StatusCode)

aOutput, _ = bucket.UploadMultipart(
	"QingCloudInsight.mov"
	&qingstor.UploadMultipartInput{
		UploadID:   "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		PartNumber: 2,
		ContentMD5: md5String2,
		Body:       file2,
	},
)

// Print the HTTP status code.
// Example: 201
fmt.Println(aOutput.StatusCode)
```

Complete Multipart Upload

``` go
aOutput, _ := bucket.CompleteMultipartUpload(
	"QingCloudInsight.mov",
	&qingstor.CompleteMultipartUploadInput{
		UploadID:    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		ObjectParts: []*qingstor.ObjectPart{{
			PartNumber: 0,
		}, {
			PartNumber: 1,
		}, {
			PartNumber: 2,
		}},
	},
)

// Print the HTTP status code.
// Example: 200
fmt.Println(aOutput.StatusCode)
```

Abort Multipart Upload

``` go
aOutput, err := bucket.AbortMultipartUpload(
	"QingCloudInsight.mov"
	&qingstor.AbortMultipartUploadInput{
		UploadID:  "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	},
)

// Print the error message.
// Example: QingStor Error: StatusCode 400, Code...
fmt.Println(err)
```
