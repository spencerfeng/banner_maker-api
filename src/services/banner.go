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

	if err := r.Save(bannerToDB); err != nil {
		return nil, err
	}

	b.ID = (*bannerToDB).ID

	return &b, nil
}

// NewBannerService ...
func NewBannerService() BannerService {
	return &bannerService{}
}
