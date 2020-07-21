package uploader

import (
	"errors"
	"mime/multipart"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/globalsign/mgo/bson"
	restError "github.com/spencerfeng/banner_maker-api/src/restError"
)

const (
	region = "ap-southeast-2"
	bucket = "banner-maker-dev"
)

// S3 ...
type S3 struct {
	sess *session.Session
}

// UploadFile ...
func (p *S3) UploadFile(fileBasePath string, file multipart.File, fileHeader *multipart.FileHeader) (string, restError.RestError) {
	key := fileBasePath + "/" + bson.NewObjectId().Hex() + filepath.Ext(fileHeader.Filename)

	uploader := s3manager.NewUploader(p.sess)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
		// The canned ACL to apply to the object. For more information, see Canned ACL
		// (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL).
		ACL: aws.String("public-read"),
	})

	if err != nil {
		return "", restError.NewInternalServerError("error when uploading the file to AWS S3", errors.New("S3 provicer error"))
	}

	return key, nil
}

// NewS3Provider ...
func NewS3Provider() (Interface, error) {
	s, err := session.NewSession(&aws.Config{Region: aws.String(region)})

	if err != nil {
		return nil, err
	}

	return &S3{
		sess: s,
	}, nil
}
