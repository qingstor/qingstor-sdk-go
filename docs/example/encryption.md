# Encryption Example

### Code Snippet

You can encrypt data when uploading.

To understand the process of encryption better, visit the link [https://docsv4.qingcloud.com/user_guide/storage/object_storage/api/object/encryption/](https://docsv4.qingcloud.com/user_guide/storage/object_storage/api/object/encryption/) .

Encrypt when uploading files. The encryption operation is performed by setting related items in PutObjectInput struct.

```go
	toPtr := func(s string) *string { return &s }
    // replace this with some file exists in your file system
	f, _ := os.Open("/tmp/file")
	putInput := &service.PutObjectInput{
		XQSEncryptionCustomerAlgorithm: toPtr("AES256"),
		XQSEncryptionCustomerKey:       toPtr("key"),
		XQSEncryptionCustomerKeyMD5:    toPtr("MD5 of the key"),
		Body:                           f,
	}
    objectKey := "your_file_encrypted"
    output, err := bucketService.PutObject(objectKey, putInput)
```

Downloading an encrypted file requires decrypting the file. It also need set the Input parameter. Please refer to the following example:
```go
	toPtr := func(s string) *string { return &s }
	getInput := service.GetObjectInput{
		XQSEncryptionCustomerAlgorithm: toPtr("AES256"),
		XQSEncryptionCustomerKey:       toPtr("key"),
		XQSEncryptionCustomerKeyMD5:    toPtr("MD5 of the key"),
	}
    // replace this with some object exists in your bucket
    objectKey := "your_file_in_bucket"
    output, err := bucketService.GetObject(objectKey, getInput)
```