package models

import (
	restError "github.com/spencerfeng/banner_maker-api/src/restError"
)

// Banner ...
type Banner struct {
	ID     int64         `json:"id"`
	Layers []BannerLayer `json:"layers"`
}

// BannerToDB ...
type BannerToDB struct {
	ID     int64
	Layers string
}

// BannerLayer ...
type BannerLayer struct {
	Type       string                `json:"type"`
	Properties BannerLayerProperties `json:"properties"`
}

// BannerLayerProperties ...
type BannerLayerProperties struct {
	X      int64  `json:"x"`
	Y      int64  `json:"y"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
	URL    string `json:"url"`
}

// BannerRepositoryInterface ...
type BannerRepositoryInterface interface {
	Save(bannerToDB *BannerToDB) restError.RestError
}
