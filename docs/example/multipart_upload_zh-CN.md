## 大文件分段上传

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

然后设置 UploadMultipart 方法用到的输入参数（使用 UploadMultipartInput 存储）。
其中会涉及到 UploadID, PartNumber, ContentLength, Body 几个参数。

首先需要申请一个 UploadID。用于后续在上传分段时，在请求参数中附加该 Upload ID，则表明分段属于同一个对象。objKey 用于指定分段上传完成后的 filepath。

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

为了得到分段上传的 part 信息。我们自定义一个 getPartsCount 用于得到相关信息。

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

开始分段上传。

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

查看已上传的分段。可以尝试访问以下方法查看是否所有的分段都已上传完成。它将返回所有的已上传分段信息。

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

当所有分段都已经上传完毕后，您可以使用下面的方法来标记上传完成，所有分段将拼接为 objectKey 指定的对象。
ETag 信息不是必须设置，具体可以参考 [api docs](https://docs.qingcloud.com/qingstor/api/object/multipart/complete_multipart_upload.html)。

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

如果您想要取消分段上传，只需要指定 objectKey 和 uploadID 即可。

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
