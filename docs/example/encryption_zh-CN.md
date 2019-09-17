# 数据加密

## 代码片段

上传时可对数据进行加密。

访问该链接 [https://docs.qingcloud.com/qingstor/api/common/encryption.html#object-storage-encryption-headers](https://docs.qingcloud.com/qingstor/api/common/encryption.html#object-storage-encryption-headers) .
以更好的理解数据加密解密的过程。

上传文件时加密。通过设置 PutObjectInput 中的相关项来进行加密操作。

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

下载加密文件需要对文件进行解密，同样是通过设置 Input 对应参数。请参考以下示例：

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