package controllers

import (
	"net/http"

	"github.com/spencerfeng/banner_maker-api/src/models"

	"github.com/gin-gonic/gin"
	restError "github.com/spencerfeng/banner_maker-api/src/restError"
	"github.com/spencerfeng/banner_maker-api/src/services"
)

// BaseBannerHandler ...
type BaseBannerHandler struct {
	bannerRepository models.BannerRepositoryInterface
	bannerService    services.BannerService
}

// NewBannerBaseHandler ...
func NewBannerBaseHandler(bannerRepository models.BannerRepositoryInterface) *BaseBannerHandler {
	return &BaseBannerHandler{
		bannerRepository: bannerRepository,
		bannerService:    services.NewBannerService(),
	}
}

// Create ...
func (h *BaseBannerHandler) Create(c *gin.Context) {
	var b models.Banner
	if err := c.ShouldBindJSON(&b); err != nil {
		restErr := restError.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, saveErr := h.bannerService.CreateBanner(h.bannerRepository, b)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}
