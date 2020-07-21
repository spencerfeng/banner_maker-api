package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spencerfeng/banner_maker-api/src/services"

	"github.com/gin-gonic/gin"
)

const (
	awsS3Region = "ap-southeast-2"
	awsS3Bucket = "banner-maker-dev"
)

// BaseImageHandler ...
type BaseImageHandler struct {
	imageService services.ImageService
}

// NewBaseImageHandler ...
func NewBaseImageHandler(imageService services.ImageService) *BaseImageHandler {
	return &BaseImageHandler{
		imageService: imageService,
	}
}

// Upload ...
func (h *BaseImageHandler) Upload(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}

	imageName, uploadErr := h.imageService.Upload(file, fileHeader)

	if uploadErr != nil {
		c.JSON(uploadErr.Status(), uploadErr)
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", imageName))
}
