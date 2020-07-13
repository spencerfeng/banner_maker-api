package services

import (
	"encoding/json"

	"github.com/spencerfeng/banner_maker-api/src/models"
	restError "github.com/spencerfeng/banner_maker-api/src/restError"
)

// BannerService ...
type BannerService interface {
	CreateBanner(r models.BannerRepositoryInterface, b models.Banner) (*models.Banner, restError.RestError)
}

type bannerService struct{}

func (s *bannerService) CreateBanner(r models.BannerRepositoryInterface, b models.Banner) (*models.Banner, restError.RestError) {
	bl, err := json.Marshal(b.Layers)

	if err != nil {
		restErr := restError.NewBadRequestError("invalid json body")
		return nil, restErr
	}

	layersJSONStr := string(bl)

	bannerToDB := &models.BannerToDB{
		Layers: layersJSONStr,
	}

	bannerID, saveErr := r.Save(bannerToDB)

	if saveErr != nil {
		return nil, saveErr
	}

	b.ID = bannerID

	return &b, nil
}

// NewBannerService ...
func NewBannerService() BannerService {
	return &bannerService{}
}
