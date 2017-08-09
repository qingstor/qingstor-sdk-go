package main

import (
	"errors"
	"github.com/DATA-DOG/godog"
	"github.com/yunify/qingstor-sdk-go/helpers"
	qs "github.com/yunify/qingstor-sdk-go/service"
	"os"
	"path"
	"strings"
)

func ImageFeatureContext(s *godog.Suite) {
	s.Step(`^initialize image by the image\'s address "([^"]*)"$`, initializeImageByTheImagesAddress)
	s.Step(`^The image is initialized$`, theImageIsInitialized)

	s.Step(`^crop the image with keys (\d+), (\d+), (\d+)$`, cropTheImageWithKeys)
	s.Step(`^crop the image status code is (\d+)$`, cropTheImageStatusCodeIs)

	s.Step(`^format the image with Type "([^"]*)"$`, formatTheImageWithType)
	s.Step(`^format the image status code is (\d+)$`, formatTheImageStatusCodeIs)

	s.Step(`^resize the image with keys (\d+), (\d+), (\d+)$`, resizeTheImageWithKeys)
	s.Step(`^resize the image status code is (\d+)$`, resizeTheImageStatusCodeIs)

	s.Step(`^rotate the image with Angle (\d+)$`, rotateTheImageWithAngle)
	s.Step(`^rotate the image status code is (\d+)$`, rotateTheImageStatusCodeIs)

	s.Step(`^watermark the image with keys (\d+), "([^"]*)", "([^"]*)"$`, watermarkTheImageWithKeys)
	s.Step(`^watermark the image status code is (\d+)$`, watermarkTheImageStatusCodeIs)

	s.Step(`^watermarkImage the image with keys (\d+), (\d+), "([^"]*)"$`, watermarkImageTheImageWithKeys)
	s.Step(`^watermarkImage the image status code is (\d+)$`, watermarkImageTheImageStatusCodeIs)

	s.Step(`^delete the image$`, deleteTheImage)
	s.Step(`^delete the image status code is (\d+)$`, deleteTheImageStatusCodeIs)

}

var image *helpers.Image
var imageName string

func initializeImageByTheImagesAddress(address string) error {
	if bucket == nil {
		return errors.New("The bucket is not exist")
	}
	if address == "" {
		return errors.New("The image's address is not exist")
	}

	file, err := os.Open(path.Join("features", "data", address))
	if err != nil {
		return err
	}
	defer file.Close()

	imageName = getImageName(address)

	_, err = bucket.PutObject(imageName, &qs.PutObjectInput{Body: file})
	if err != nil {
		return err
	}

	image = helpers.InitImage(bucket, imageName)
	if image.Err != nil {
		return image.Err
	}

	return nil
}

func theImageIsInitialized() error {
	if image == nil {
		return errors.New("The Image is not initialized")
	}
	return nil
}

func getImageName(address string) (imageName string) {

	winResult := strings.Split(address, "\\")
	otherResult := strings.Split(address, "/")

	if len(winResult) == 1 && len(otherResult) == 1 {
		if winResult[0] == otherResult[0] {
			imageName = winResult[0]
		}
	} else if len(winResult) == 1 && len(otherResult) > 1 {
		imageName = otherResult[len(otherResult)-1]
	} else if len(otherResult) == 1 && len(winResult) > 1 {
		imageName = winResult[len(winResult)-1]
	}

	return
}

var imageOutput *qs.ImageProcessOutput

func cropTheImageWithKeys(width, height, gravity int) error {
	image = image.Crop(&helpers.CropParam{
		Width:   width,
		Height:  height,
		Gravity: helpers.CropGravity(gravity),
	})

	if image.Err != nil {
		return image.Err
	}
	imageOutput = image.ImageOutput
	return nil
}

func cropTheImageStatusCodeIs(statusCode int) error {
	return checkEqual(qs.IntValue(imageOutput.StatusCode), statusCode)
}

func formatTheImageWithType(types string) error {
	image = image.Format(&helpers.FormatParam{Type: types})
	if image.Err != nil {
		return image.Err
	}
	imageOutput = image.ImageOutput
	return nil
}

func formatTheImageStatusCodeIs(statusCode int) error {
	return checkEqual(qs.IntValue(imageOutput.StatusCode), statusCode)
}

func resizeTheImageWithKeys(width, height, mode int) error {
	image = image.Resize(&helpers.ResizeParam{
		Width:  width,
		Height: height,
		Mode:   helpers.ResizeMode(mode),
	})
	if image.Err != nil {
		return image.Err
	}
	imageOutput = image.ImageOutput
	return nil
}

func resizeTheImageStatusCodeIs(statusCode int) error {
	return checkEqual(qs.IntValue(imageOutput.StatusCode), statusCode)
}

func rotateTheImageWithAngle(angle int) error {
	image = image.Rotate(&helpers.RotateParam{Angle: angle})
	if image.Err != nil {
		return image.Err
	}
	imageOutput = image.ImageOutput
	return nil
}

func rotateTheImageStatusCodeIs(statusCode int) error {
	return checkEqual(qs.IntValue(imageOutput.StatusCode), statusCode)
}

func watermarkTheImageWithKeys(dpi int, text, color string) error {
	image = image.WaterMark(&helpers.WaterMarkParam{
		Text:  text,
		Dpi:   dpi,
		Color: color,
	})
	if image.Err != nil {
		return image.Err
	}
	imageOutput = image.ImageOutput
	return nil
}

func watermarkTheImageStatusCodeIs(statusCode int) error {
	return checkEqual(qs.IntValue(imageOutput.StatusCode), statusCode)
}

func watermarkImageTheImageWithKeys(left, top int, url string) error {
	image = image.WaterMarkImage(&helpers.WaterMarkImageParam{
		Url:  url,
		Left: left,
		Top:  top,
	})
	if image.Err != nil {
		return image.Err
	}
	imageOutput = image.ImageOutput
	return nil
}

func watermarkImageTheImageStatusCodeIs(statusCode int) error {
	return checkEqual(qs.IntValue(imageOutput.StatusCode), statusCode)
}

var oOutput *qs.DeleteObjectOutput

func deleteTheImage() error {
	if bucket == nil {
		return errors.New("The bucket has not been initialized yet")
	}

	// Init image by the bucket and the object name(eg: default.jpg)
	oOutput, err = bucket.DeleteObject(imageName)

	if err != nil {
		return err
	}
	return nil
}

func deleteTheImageStatusCodeIs(statusCode int) error {
	return checkEqual(qs.IntValue(oOutput.StatusCode), statusCode)
}
