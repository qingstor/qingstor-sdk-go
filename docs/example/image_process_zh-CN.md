# 基本图片处理

用于对用户存储于 QingStor 对象存储上的图片进行各种基本处理，例如格式转换，裁剪，翻转，水印等。

目前支持的图片格式有:

- png
- tiff
- webp
- jpeg
- pdf
- gif
- svg

> 目前不支持对加密过后的图片进行处理，单张图片最大为 10M 。

具体文档说明参考 [API Docs](https://docsv4.qingcloud.com/user_guide/storage/object_storage/api/object/image_process/) 。

## 代码片段

使用您的 AccessKeyID 和 SecretAccessKey 初始化 Qingstor 对象。

```go
import (
	"github.com/qingstor/qingstor-sdk-go/v4/config"
	"github.com/qingstor/qingstor-sdk-go/v4/service"
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

初始化一张图片
```go
img := image.Init(bucket, "imageName")
```

现在，您可以使用高级API或基本图像处理API来执行图像操作。

获取图像的信息
```go
imageProcessOutput, _ := img.Info().Process()
```

裁剪图像。
```go
imageProcessOutput, _  := img.Crop(&image.CropParam{
	...operation_param...
}).Process()
```

旋转图像。
```go
imageProcessOutput, _ := img.Rotate(&image.RotateParam{
	...operation_param...
}).Process()
```
调整图像大小。
```go
imageProcessOutput, _ := img.Resize(&image.ResizeParam{
	...operation_param...
}).Process()
```
为图像添加水印。
```go
imageProcessOutput, _ := img.WaterMark(&image.WaterMarkParam{
	...operation_param...
}).Process()
```
WaterMarkImage图像。
```go
imageProcessOutput, _ : = img.WaterMarkImage(&image.WaterMarkImageParam{
	...operation_param...
}).Process()
```
格式化图像。
```go
imageProcessOutput, _ := img.Format(&image.Format{
	...operation_param...
}).Process()
```
操作管道，图像将按顺序处理。 管道中的最大操作数为10。
```go
// Rotate and then resize the image
imageProcessOutput, _ := img.Rotate(&image.RotateParam{
	... operation_param...
}).Resize(&image.ResizeParam{
	... operation_param...
}).Process()
```
使用原始基本API将图像旋转90度角。
```go
operation := "rotate:a_90"
imageProcessOutput, err := bucket.ImageProcess("imageName", &qs.ImageProcessInput{
	Action: &operation})
```

`operation_param`是图像操作参数，它在`qingstor-sdk-go / client / image / image.go`中定义。
```go
import "github.com/qingstor/qingstor-sdk-go/v4/service"
// client/image/image.go
type Image struct {
	key    *string
	bucket *service.Bucket
	input  *service.ImageProcessInput
}

// About cropping image definition
type CropGravity int
const (
	CropCenter CropGravity = iota
	CropNorth
	CropEast
	CropSouth
	CropWest
	CropNorthWest
	CropNorthEast
	CropSouthWest
	CropSouthEast
	CropAuto
)
type CropParam struct {
	Width   int         `schema:"w,omitempty"`
	Height  int         `schema:"h,omitempty"`
	Gravity CropGravity `schema:"g"`
}

// About rotating image definitions
type RotateParam struct {
	Angle int `schema:"a"`
}

// About resizing image definitions
type ResizeMode int
type ResizeParam struct {
	Width  int        `schema:"w,omitempty"`
	Height int        `schema:"h,omitempty"`
	Mode   ResizeMode `schema:"m"`
}

// On the definition of text watermarking
type WaterMarkParam struct {
	Dpi     int     `schema:"d,omitempty"`
	Opacity float64 `schema:"p,omitempty"`
	Text    string  `schema:"t"`
	Color   string  `schema:"c"`
}

// On the definition of image watermarking
 type WaterMarkImageParam struct {
	Left    int     `schema:"l"`
	Top     int     `schema:"t"`
	Opacity float64 `schema:"p,omitempty"`
	URL     string  `schema:"u"`
}

// About image format conversion definitions
type FormatParam struct {
	Type string `schema:"t"`
}

```

__快速入门代码示例:__

包含一个完整的示例，但代码需要填写您自己的信息。

```go
package main

import (
	"log"

	"github.com/qingstor/qingstor-sdk-go/v4/client/image"
	"github.com/qingstor/qingstor-sdk-go/v4/config"
	qs "github.com/qingstor/qingstor-sdk-go/v4/service"
)

func main() {
	// Load your configuration
	// Replace here with your key pair
	config, err := config.New("ACCESS_KEY_ID", "SECRET_ACCESS_KEY")
	checkErr(err)

	// Initialize QingStror Service
	service, err := qs.Init(config)
	checkErr(err)

	// Initialize Bucket
	// Replace here with your bucketName and zoneID
	bucket, err := service.Bucket("bucketName", "zoneID")
	checkErr(err)

	// Initialize Image
	// Replace here with your your ImageName
	img := image.Init(bucket, "imageName")
	checkErr(err)

	// Because 0 is an invalid parameter, default not modify
	imageProcessOutput, err := img.Crop(&image.CropParam{Width: 0}).Process()
	checkErr(err)
	testOutput(imageProcessOutput)

	// Rotate the image 90 angles
	imageProcessOutput, err = img.Rotate(&image.RotateParam{Angle: 90}).Process()
	checkErr(err)
	testOutput(imageProcessOutput)

	// Text watermark, Watermark text content, encoded by base64.
	imageProcessOutput, err = img.WaterMark(&image.WaterMarkParam{
		Text: "5rC05Y2w5paH5a2X",
	}).Process()
	checkErr(err)
	testOutput(imageProcessOutput)

	// Image watermark, Watermark image url encoded by base64.
	imageProcessOutput, err = img.WaterMarkImage(&image.WaterMarkImageParam{
		URL: "aHR0cHM6Ly9wZWszYS5xaW5nc3Rvci5jb20vaW1nLWRvYy1lZy9xaW5jbG91ZC5wbmc",
	}).Process()
	checkErr(err)
	testOutput(imageProcessOutput)

	// Reszie the image with width 300px and height 400 px
	imageProcessOutput, err = img.Resize(&image.ResizeParam{
		Width:  300,
		Height: 400,
	}).Process()
	checkErr(err)
	testOutput(imageProcessOutput)

	// Swap format to jpeg
	imageProcessOutput, err = img.Format(&image.FormatParam{
		Type: "jpeg",
	}).Process()
	checkErr(err)
	testOutput(imageProcessOutput)

	// Pipline model
	// The maximum number of operations in the pipeline is 10
	imageProcessOutput, err = img.Rotate(&image.RotateParam{
		Angle: 270,
	}).Resize(&image.ResizeParam{
		Width:  300,
		Height: 300,
	}).Process()
	checkErr(err)
	testOutput(imageProcessOutput)

	// Get the information of the image
	imageProcessOutput, err = img.Info().Process()
	checkErr(err)
	testOutput(imageProcessOutput)

	// Use the original api to rotate the image 90 angles
	operation := "rotate:a_90"
	imageProcessOutput, err = bucket.ImageProcess("imageName", &qs.ImageProcessInput{
		Action: &operation})
	checkErr(err)
	defer imageProcessOutput.Close() // Don't forget to close the output otherwise will be leaking http connections
	testOutput(imageProcessOutput)
}

// *qs.ImageProcessOutput: github.com/qingstor/qingstor-sdk-go/v4/service/object.go
func testOutput(out *qs.ImageProcessOutput) {
	log.Println(*out.StatusCode)
	log.Println(*out.RequestID)
	log.Println(out.Body)
	log.Println(*out.ContentLength)
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
```