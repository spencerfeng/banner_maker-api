package uploader

import (
	"mime/multipart"

	restError "github.com/spencerfeng/banner_maker-api/src/restError"
)

// Interface ...
type Interface interface {
	UploadFile(fileBasePath string, file multipart.File, fileHeader *multipart.FileHeader) (string, restError.RestError)
}
