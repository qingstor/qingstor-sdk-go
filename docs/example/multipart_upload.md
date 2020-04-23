# UploadMultipart Example

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


Then set the input parameters used by the UploadMultipart method (stored in UploadMultipartInput).
Required parameters include UploadID, PartNumber, ContentLength, Body.

First you need to apply for an UploadID. For subsequent uploading of the segment, appending the Upload ID to the request parameter indicates that the segment belongs to the same object. objKey is used to specify the filepath after the segment upload is completed.

```go
func InitUploadID(bucketService *service.Bucket, objKey string) (*string, error) {
	output, err := bucketService.InitiateMultipartUpload(
		objKey,
		&service.InitiateMultipartUploadInput{},
	)
	if err != nil {
		fmt.Printf("The attempt to generate upload id failed with given error: %s\n", err)
		return nil, err
	} else {
		fmt.Printf("The status code expected: 200(actually: %d)\n", * output.StatusCode)
		return output.UploadID, err
	}
}
```

In order to get the part information of the segment upload. We defined a function named getPartsCount to get the relevant information.

```go
// getPartsCount return a file(from path) can be split to how many parts(depends on chunkSize) and file size.
func getPartsCount(path string, chunkSize int) (parts int, size int) {
	file, _ := os.Open(path)
	defer file.Close()
	fileInfo, _ := file.Stat()
	sz := fileInfo.Size()
	return int(math.Ceil(float64(sz) / float64(chunkSize))), int(sz)
}
```

Start a multipart upload.

```go
	objectKey := "your_file_multi_uploaded.zip"
	uploadID, _ := testInitUploadID(bucketService, objectKey)
	fmt.Println(uploadID)
	filePath := "/home/max/Pictures/test/your_file_multi_uploaded.zip"
	chunkSize := 5 * (1 << 20)
	partsCount, size := getPartsCount(filePath, chunkSize)
	f, _ := os.Open(filePath)
	for i := 0; i < partsCount; i++ {
		partSize := int64(math.Min(float64(chunkSize), float64(size-i*chunkSize)))
		partBuffer := make([]byte, partSize)
		_, _ = f.Read(partBuffer)
		partNumber := i
		uploadOutput, _ := bucketService.UploadMultipart(
			objectKey,
			&service.UploadMultipartInput{
				UploadID:      uploadID,
				PartNumber:    &partNumber,
				ContentLength: &partSize,
				Body:          bytes.NewReader(partBuffer),
			},
		)
		fmt.Println("201 expected, actually:", *uploadOutput.StatusCode)
	}
	_ = f.Close()
```

View the uploaded segments. Try the following methods to see if all the segments have been uploaded. It will return all uploaded segmentation information.

```go
// https://docs.qingcloud.com/qingstor/api/object/multipart/list_multipart.html
func ListMultiParts(bucketService *service.Bucket, objectKey string, uploadID *string) ([]*service.ObjectPartType, error) {
	output, err := bucketService.ListMultipart(
		objectKey,
		&service.ListMultipartInput{
			UploadID: uploadID,
		},
	)
	if err != nil {
		fmt.Printf("List uploaded parts failed with given error: %s\n", err)
		return nil, err
	} else {
		fmt.Printf("The status code expected: 200(actually: %d)\n", *output.StatusCode)
		// b, _ := json.MarshalIndent(output, "", "\t")
		// fmt.Printf("The multipart info of object(%s) with uploadID(%s):\n%s\n", objectKey, *uploadID, string(b))
		return output.ObjectParts, err
	}
}
```

Once all the segments have been uploaded, you can use the following method to mark the upload completion, all segments will be stitched to the object specified by objectKey.
ETag header is not required to be set. For details, please refer to [api docs] (https://docs.qingcloud.com/qingstor/api/object/multipart/complete_multipart_upload.html).

```go
func CompleteMultiParts(bucketService *service.Bucket, filepath string, objectKey string, uploadID *string, parts []*service.ObjectPartType) {
	f, _ := os.Open(filepath)
	hash := md5.New()
	if _, err := io.Copy(hash, f); err != nil {
		log.Fatal(err)
	}
	_ = f.Close()
	checksum := hex.EncodeToString(hash.Sum(nil))
	output, err := bucketService.CompleteMultipartUpload(
		objectKey,
		&service.CompleteMultipartUploadInput{
			ETag:        &checksum,
			UploadID:    uploadID,
			ObjectParts: parts,
		},
	)
	if err != nil {
		fmt.Printf("Complete uploaded parts failed with given error: %s\n", err)
	} else {
		fmt.Printf("The status code expected: 201(actually: %d)\n", *output.StatusCode)
	}
}
```

If you want to cancel a multipart upload, just specify objectKey and uploadID.

```go
func AbortMultiUpload(bucketService *service.Bucket, objectKey string, uploadID *string) {
	output, err := bucketService.AbortMultipartUpload(objectKey, &service.AbortMultipartUploadInput{UploadID: uploadID})
	if err != nil {
		fmt.Printf("Abort multiparts upload failed with given error: %s\n", err)
	} else {
		fmt.Printf("The status code expected: 204(actually: %d)\n", *output.StatusCode)
	}
}
```
