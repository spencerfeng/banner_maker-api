package services

import (
	"errors"
	"testing"

	"github.com/spencerfeng/banner_maker-api/src/models"
	restError "github.com/spencerfeng/banner_maker-api/src/restError"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type BannerRepositoryMock struct {
	mock.Mock
}

func (m *BannerRepositoryMock) Save(bannerToDB *models.BannerToDB) (int64, restError.RestError) {
	args := m.Called(bannerToDB)

	if args.Get(1) == nil {
		return args.Get(0).(int64), nil
	}

	return args.Get(0).(int64), args.Get(1).(restError.RestError)
}

func TestCreateBannerSucceed(t *testing.T) {
	bannerRepoMock := new(BannerRepositoryMock)
	bannerRepoMock.On("Save", mock.Anything).Return(int64(1), nil)

	service := bannerService{}
	banner := models.Banner{
		Layers: []models.BannerLayer{
			{
				Type: "image",
				Properties: models.BannerLayerProperties{
					X:      0,
					Y:      0,
					Width:  200,
					Height: 200,
					URL:    "https://test-image.com",
				},
			},
		},
	}

	b, err := service.CreateBanner(bannerRepoMock, banner)

	assert.Nil(t, err)
	assert.Equal(t, b.ID, int64(1))
}

func TestCreateBannerFailDueToRepositorySaveError(t *testing.T) {
	bannerRepoMock := new(BannerRepositoryMock)
	bannerRepoMock.On("Save", mock.Anything).Return(int64(0), restError.NewInternalServerError("error when trying to save banner", errors.New("database error")))

	service := bannerService{}
	banner := models.Banner{
		Layers: []models.BannerLayer{
			{
				Type: "image",
				Properties: models.BannerLayerProperties{
					X:      0,
					Y:      0,
					Width:  200,
					Height: 200,
					URL:    "https://test-image.com",
				},
			},
		},
	}

	b, err := service.CreateBanner(bannerRepoMock, banner)

	assert.Nil(t, b)
	assert.Equal(t, err.Status(), 500)
	assert.Equal(t, err.Message(), "error when trying to save banner")
	assert.Equal(t, err.Error(), "message: error when trying to save banner - status: 500 - error: internal_server_error")
}
