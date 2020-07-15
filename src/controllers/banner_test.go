package controllers_test

import (
	"bytes"
	"encoding/json"
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
