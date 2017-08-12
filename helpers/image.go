package helpers

import (
	"fmt"
	"github.com/yunify/qingstor-sdk-go/service"
	"reflect"
	"strings"
)

const (
	// ActionSep is seperator of action.
	ActionSep = ":"

	// OPSep is seperator of operation.
	OPSep = "|"

	// KVSep is seperator of Key-Value.
	KVSep = "_"

	//KVPairSep is seperator of args.
	KVPairSep = ","
)

const (
	// InfoOperation is string of info operation.
	InfoOperation string = "info"

	// CropOperation is string of crop operation.
	CropOperation string = "crop"

	// FormatOperation is string of format operation.
	FormatOperation string = "format"

	// ResizeOperation is string of resize operation.
	ResizeOperation string = "resize"

	// RotateOperation is string of rotate operation.
	RotateOperation string = "rotate"

	// WaterMarkOperation is string of watermark operation.
	WaterMarkOperation string = "watermark"

	// WaterMarkImageOperation is string of watermark image operation.
	WaterMarkImageOperation string = "watermark_image"
)

// ResizeMode is the type of resize mode.
type ResizeMode int

const (
	// ResizeFixed resizes image to fix  width and height.
	ResizeFixed ResizeMode = iota

	// ResizeForce resizes image to force witdth and height.
	ResizeForce

	// ResizeThumbnail resizes image to thumbnail width and height.
	ResizeThumbnail
)

// CropGravity is the type of crop gravity.
type CropGravity int

const (

	// CropCenter crops image to center width and height.
	CropCenter CropGravity = iota

	// CropNorth crops image to north width and height.
	CropNorth

	// CropEast crops image to east width and height.
	CropEast

	// CropSouth crops image to south width and height.
	CropSouth

	// CropWest crops image to west width and height.
	CropWest

	// CropNorthWest crops image to north west width and height.
	CropNorthWest

	// CropNorthEast crops image to north east width and height.
	CropNorthEast

	// CropSouthWest crops image to south west width and height.
	CropSouthWest

	// CropSouthEast crops image to south east width and height.
	CropSouthEast

	// CropAuto crops image to auto width and height.
	CropAuto
)

// Image is the combination type usually image process
type Image struct {
	ImageOutput *service.ImageProcessOutput
	bucket      *service.Bucket
	uri         *string
	name        *string
	Err         error
}

// InitImage initialize image.
func InitImage(b *service.Bucket, objectKey string) *Image {
	return &Image{
		bucket: b,
		name:   &objectKey,
	}
}

// Info gets the information of the image.
func (image *Image) Info() (*service.ImageProcessOutput, error) {
	image.setActionParam(InfoOperation, nil)
	return imageProcess(image).ImageOutput, image.Err
}

// RotateParam is param of the rotate operation.
type RotateParam struct {
	Angle int `schema:"a"`
}

// Rotate image.
func (image *Image) Rotate(param *RotateParam) *Image {
	image.setActionParam(RotateOperation, param)
	return imageProcess(image)
}

// ResizeParam is param of the resize operation.
type ResizeParam struct {
	Width  int        `schema:"w,omitempty"`
	Height int        `schema:"h,omitempty"`
	Mode   ResizeMode `schema:"m"`
}

// Resize image.
func (image *Image) Resize(param *ResizeParam) *Image {
	image.setActionParam(ResizeOperation, param)
	return imageProcess(image)
}

// CropParam is param of the crop operation.
type CropParam struct {
	Width   int         `schema:"w,omitempty"`
	Height  int         `schema:"h,omitempty"`
	Gravity CropGravity `schema:"g"`
}

// Crop image.
func (image *Image) Crop(param *CropParam) *Image {
	image.setActionParam(CropOperation, param)
	return imageProcess(image)
}

// FormatParam is param of the format operation.
type FormatParam struct {
	Type string `schema:"t"`
}

// Format image.
func (image *Image) Format(param *FormatParam) *Image {
	image.setActionParam(FormatOperation, param)
	return imageProcess(image)
}

// WaterMarkParam is param of the wartermark operation.
type WaterMarkParam struct {
	Dpi     int     `schema:"d,omitempty"`
	Opacity float64 `schema:"p,omitempty"`
	Text    string  `schema:"t"`
	Color   string  `schema:"c"`
}

// WaterMark  text content.
func (image *Image) WaterMark(param *WaterMarkParam) *Image {
	image.setActionParam(WaterMarkOperation, param)
	return imageProcess(image)
}

// WaterMarkImageParam is param of the  waterMark image operation
type WaterMarkImageParam struct {
	Left    int     `schema:"l"`
	Top     int     `schema:"t"`
	Opacity float64 `schema:"p,omitempty"`
	URL     string  `schema:"u"`
}

// WaterMarkImage is watermark image operation.
func (image *Image) WaterMarkImage(param *WaterMarkImageParam) *Image {
	image.setActionParam(WaterMarkImageOperation, param)
	return imageProcess(image)
}

func (image *Image) setActionParam(operation string, in interface{}) {
	var uri string
	if image.uri != nil {
		if in != nil {
			uri = fmt.Sprintf("%s%s%s%s%s", *image.uri, OPSep, operation,
				ActionSep, buildOptParamStr(in))
		} else {
			uri = fmt.Sprintf("%s%s%s", *image.uri, OPSep, operation)
		}
	} else {
		if in != nil {
			uri = fmt.Sprintf("%s%s%s", operation, ActionSep,
				buildOptParamStr(in))
		} else {
			uri = fmt.Sprintf("%s", operation)
		}
	}
	image.uri = &uri
}

func imageProcess(image *Image) *Image {
	output, err := image.bucket.
		ImageProcess(*image.name, &service.ImageProcessInput{Action: image.uri})
	image.Err = err
	image.ImageOutput = output
	return image
}

func buildOptParamStr(param interface{}) string {
	v := reflect.ValueOf(param).Elem()
	var kvPairs []string

	for i := 0; i < v.NumField(); i++ {
		vf := v.Field(i)
		tf := v.Type().Field(i)
		key := tf.Tag.Get("schema")
		value := vf.Interface()
		tagValues := strings.Split(key, ",")
		if isEmptyValue(vf) &&
			len(tagValues) == 2 &&
			tagValues[1] == "omitempty" {
			continue
		}
		key = tagValues[0]
		kvPairs = append(kvPairs, fmt.Sprintf("%v%s%v", key, KVSep, value))
	}
	return strings.Join(kvPairs, KVPairSep)
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array,
		reflect.Map,
		reflect.Slice,
		reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		return v.Int() == 0
	case reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32,
		reflect.Float64:
		return v.Float() == 0
	case reflect.Ptr:
		return v.IsNil()
	}
	return false
}
