package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/spencerfeng/banner_maker-api/src/app"
	"github.com/stretchr/testify/mock"

	"github.com/gin-gonic/gin"
	"github.com/spencerfeng/banner_maker-api/src/models"
	restError "github.com/spencerfeng/banner_maker-api/src/restError"
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

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

type BannerServiceMock struct {
	mock.Mock
}

func (s *BannerServiceMock) CreateBanner(r models.BannerRepositoryInterface, b models.Banner) (*models.Banner, restError.RestError) {
	args := s.Called(r, b)

	if args.Get(0) == nil {
		return nil, args.Get(1).(restError.RestError)
	}

	return args.Get(0).(*models.Banner), nil
}

func TestCreateBanner(t *testing.T) {
	bannerRepoMock := new(BannerRepositoryMock)
	bannerRepoMock.On("Save", mock.Anything).Return(int64(1), nil)

	reqBodyJSONStr := []byte(`{
		"layers": [
			{
				"type": "image",
				"properties": {
					"x": 0,
					"y": 0,
					"width": 200,
					"height": 200,
					"url": "https://test-image.com"
				}
			}
		]
	}`)

	expectedResBody := gin.H{
		"id": float64(1),
		"layers": []interface{}{
			map[string]interface{}{
				"type": "image",
				"properties": map[string]interface{}{
					"x":      float64(0),
					"y":      float64(0),
					"width":  float64(200),
					"height": float64(200),
					"url":    "https://test-image.com",
				},
			},
		},
	}

	router := app.SetupRouter(bannerRepoMock)

	w := performRequest(router, "POST", "/banners", bytes.NewBuffer(reqBodyJSONStr))

	var resBody map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &resBody)

	if err == nil {
		assert.Equal(t, w.Code, http.StatusCreated)

		idVal, idExists := resBody["id"]
		assert.True(t, idExists)
		assert.Equal(t, expectedResBody["id"], idVal)

		layersVal, layersExists := resBody["layers"]
		assert.True(t, layersExists)
		assert.Equal(t, expectedResBody["layers"], layersVal)
	}
}

func TestCreateBannerFailDueToDatabaseSaveError(t *testing.T) {
	bannerServiceMock := new(BannerServiceMock)
	bannerServiceMock.On("CreateBanner", mock.Anything).Return(nil, restError.NewInternalServerError("error when trying to save banner", errors.New("database error")))

	bannerRepoMock := new(BannerRepositoryMock)
	bannerRepoMock.On("Save", mock.Anything).Return(int64(0), restError.NewInternalServerError("error when trying to save banner", errors.New("database error")))

	reqBodyJSONStr := []byte(`{
		"layers": [
			{
				"type": "image",
				"properties": {
					"x": 0,
					"y": 0,
					"width": 200,
					"height": 200,
					"url": "https://test-image.com"
				}
			}
		]
	}`)

	expectedResBody := gin.H{
		"message": "error when trying to save banner",
		"status":  float64(500),
		"error":   "internal_server_error",
		"causes":  []interface{}{"database error"},
	}

	router := app.SetupRouter(bannerRepoMock)

	w := performRequest(router, "POST", "/banners", bytes.NewBuffer(reqBodyJSONStr))

	var resBody map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &resBody)

	if err == nil {
		assert.Equal(t, w.Code, http.StatusInternalServerError)

		messageVal, messageExists := resBody["message"]
		assert.True(t, messageExists)
		assert.Equal(t, expectedResBody["message"], messageVal)

		statusVal, statusExists := resBody["status"]
		assert.True(t, statusExists)
		assert.Equal(t, expectedResBody["status"], statusVal)

		errorVal, errorExists := resBody["error"]
		assert.True(t, errorExists)
		assert.Equal(t, expectedResBody["error"], errorVal)

		causesVal, causesExists := resBody["causes"]
		assert.True(t, causesExists)
		assert.Equal(t, expectedResBody["causes"], causesVal)
	}
}
