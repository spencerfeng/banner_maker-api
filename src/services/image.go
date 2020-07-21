package services

import (
	"mime/multipart"

	"github.com/spencerfeng/banner_maker-api/src/providers/uploader"
	restError "github.com/spencerfeng/banner_maker-api/src/restError"
)

// ImageService ...
type ImageService interface {
	Upload(i multipart.File, h *multipart.FileHeader) (string, restError.RestError)
}

type imageService struct {
	uploader uploader.Interface
}

func (s *imageService) Upload(i multipart.File, h *multipart.FileHeader) (string, restError.RestError) {
	imageName, err := s.uploader.UploadFile("image", i, h)
	if err != nil {
		return "", err
	}

	return imageName, nil
}

// NewImageService ...
func NewImageService() (ImageService, error) {
	u, err := uploader.NewS3Provider()
	if err != nil {
		return nil, err
	}

	return &imageService{uploader: u}, nil
}
