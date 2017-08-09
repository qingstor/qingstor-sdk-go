package helpers

import (
	"fmt"
	"github.com/yunify/qingstor-sdk-go/service"

	"reflect"
	"strings"
)

const (
	ActionSep = ":"
	OPSep     = "|"
	KVSep     = "_"
	KVPairSep = ","
)

const (
	CropOperation           string = "crop"
	FormatOperation         string = "format"
	ResizeOperation         string = "resize"
	RotateOperation         string = "rotate"
	WaterMarkOperation      string = "watermark"
	WaterMarkImageOperation string = "watermark_image"
)

type ResizeMode int

const (
	ResizeFixed ResizeMode = iota
	ResizeForce
	ResizeThumbnail
)

type CropGravity int

const (
	CropCenter CropGravity = iota
	CropNorth
	CropEast
	CropSouth
	CropWest
	CropAuto
)

type Image struct {
	ImageOutput *service.ImageProcessOutput
	bucket      *service.Bucket
	uri         *string
	name        *string
	Err         error
}

func InitImage(b *service.Bucket, objectKey string) *Image {
	return &Image{
		bucket: b,
		name:   &objectKey,
	}
}

type RotateParam struct {
	Angle int `schema:"a"`
}

func (image *Image) Rotate(param *RotateParam) *Image {
	image.setActionParam(fmt.Sprintf("%s%s", RotateOperation, ActionSep), param)
	return imageProcess(image)
}

type ResizeParam struct {
	Width  int        `schema:"w,omitempty"`
	Height int        `schema:"h,omitempty"`
	Mode   ResizeMode `schema:"m"`
}

func (image *Image) Resize(param *ResizeParam) *Image {
	image.setActionParam(fmt.Sprintf("%s%s", ResizeOperation, ActionSep), param)
	return imageProcess(image)
}

type CropParam struct {
	Width   int         `schema:"w,omitempty"`
	Height  int         `schema:"h,omitempty"`
	Gravity CropGravity `schema:"g"`
}

func (image *Image) Crop(param *CropParam) *Image {
	image.setActionParam(fmt.Sprintf("%s%s", CropOperation, ActionSep), param)
	return imageProcess(image)
}

type FormatParam struct {
	Type string `schema:"t"`
}

func (image *Image) Format(param *FormatParam) *Image {
	image.setActionParam(fmt.Sprintf("%s%s", FormatOperation, ActionSep), param)
	return imageProcess(image)
}

type WaterMarkParam struct {
	Dpi     int     `schema:"d,omitempty"`
	Opacity float64 `schema:"p,omitempty"`
	Text    string  `schema:"t"`
	Color   string  `schema:"c"`
}

func (image *Image) WaterMark(param *WaterMarkParam) *Image {
	image.setActionParam(fmt.Sprintf("%s%s", WaterMarkOperation, ActionSep), param)
	return imageProcess(image)
}

type WaterMarkImageParam struct {
	Left    int     `schema:"l"`
	Top     int     `schema:"t"`
	Opacity float64 `schema:"p,omitempty"`
	Url     string  `schema:"u"`
}

func (image *Image) WaterMarkImage(param *WaterMarkImageParam) *Image {
	image.setActionParam(fmt.Sprintf("%s%s", WaterMarkImageOperation, ActionSep), param)
	return imageProcess(image)
}

func (image *Image) setActionParam(operation string, in interface{}) {
	var uri string
	if image.uri != nil {
		uri = fmt.Sprintf("%s%s%s%s", *image.uri, OPSep, operation, buildOptParamStr(in))
	} else {
		uri = fmt.Sprintf("%s%s", operation, buildOptParamStr(in))
	}
	image.uri = &uri
}

func imageProcess(image *Image) *Image {
	output, err := image.bucket.ImageProcess(*image.name, &service.ImageProcessInput{Action: image.uri})
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
