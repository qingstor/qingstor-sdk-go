package upload

import (
	"context"
	"errors"
	"io"

	"go.uber.org/zap"

	"github.com/qingstor/qingstor-sdk-go/v4/log"
	"github.com/qingstor/qingstor-sdk-go/v4/service"
)

// Uploader struct provides a struct to upload
type Uploader struct {
	bucket   *service.Bucket
	partSize int
}

const smallestPartSize int = 1024 * 1024 * 4

// Init creates a uploader struct
func Init(bucket *service.Bucket, partSize int) *Uploader {
	return &Uploader{
		bucket:   bucket,
		partSize: partSize,
	}
}

// Upload uploads multi parts of large object
func (u *Uploader) Upload(fd io.Reader, objectKey string) error {
	return u.UploadWithContext(context.Background(), fd, objectKey)
}

// UploadWithContext add support for context
func (u *Uploader) UploadWithContext(ctx context.Context, fd io.Reader, objectKey string) error {
	logger := log.FromContext(ctx)
	length, err := getFileSize(fd)
	if err != nil {
		logger.Error("get file size", zap.Error(err))
		return err
	}
	if length < int64(smallestPartSize) {
		_, err := u.bucket.PutObjectWithContext(ctx, objectKey, &service.PutObjectInput{Body: fd})
		if err != nil {
			logger.Error("auto switch to put object", zap.Error(err))
			return err
		}
		return nil
	}
	if u.partSize < smallestPartSize {
		logger.Error("part size too small")
		return errors.New("the part size is too small")
	}

	uploadID, err := u.init(ctx, objectKey)
	if err != nil {
		logger.Error("init multipart upload", zap.Error(err))
		return err
	}

	partNumbers, err := u.upload(ctx, fd, uploadID, objectKey)
	if err != nil {
		logger.Error("upload part",
			zap.String("upload id", *uploadID), zap.String("key", objectKey), zap.Error(err))
		return err
	}

	err = u.complete(ctx, objectKey, uploadID, partNumbers)
	if err != nil {
		logger.Error("complete upload",
			zap.String("upload id", *uploadID), zap.String("key", objectKey), zap.Error(err))
		return err
	}

	return nil
}

func (u *Uploader) init(ctx context.Context, objectKey string) (*string, error) {
	output, err := u.bucket.InitiateMultipartUploadWithContext(
		ctx,
		objectKey,
		&service.InitiateMultipartUploadInput{},
	)
	if err != nil {
		return nil, err
	}
	return output.UploadID, nil
}

func (u *Uploader) upload(ctx context.Context, fd io.Reader, uploadID *string, objectKey string) ([]*service.ObjectPartType, error) {
	logger := log.FromContext(ctx)
	var partCnt int
	partNumbers := []*service.ObjectPartType{}
	fileReader := newChunk(fd, u.partSize)
	for {
		partBody, err := fileReader.nextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Error("get next part", zap.Error(err))
			return nil, err
		}
		_, err = u.bucket.UploadMultipartWithContext(
			ctx,
			objectKey,
			&service.UploadMultipartInput{
				UploadID:   uploadID,
				PartNumber: &partCnt,
				Body:       partBody,
			},
		)
		if err != nil {
			logger.Error("upload part", zap.String("key", objectKey), zap.Error(err))
			return nil, err
		}
		partNumbers = append(partNumbers, &service.ObjectPartType{
			PartNumber: service.Int(partCnt - 0),
		})
		partCnt++
	}
	return partNumbers, nil
}

func (u *Uploader) complete(ctx context.Context, objectKey string, uploadID *string, partNumbers []*service.ObjectPartType) error {
	_, err := u.bucket.CompleteMultipartUploadWithContext(
		ctx,
		objectKey,
		&service.CompleteMultipartUploadInput{
			UploadID:    uploadID,
			ObjectParts: partNumbers,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func getFileSize(fd io.Reader) (int64, error) {
	var length int64 = -1
	switch r := fd.(type) {
	case io.Seeker:
		pos, _ := r.Seek(0, 1)
		defer r.Seek(pos, 0)

		n, err := r.Seek(0, 2)
		if err != nil {
			return length, err
		}
		length = n
	}
	if length == -1 {
		return length, errors.New("The file is not seekable")
	}
	return length, nil
}
