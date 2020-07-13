package repositories

import (
	"database/sql"
	"errors"

	"github.com/spencerfeng/banner_maker-api/src/models"

	restError "github.com/spencerfeng/banner_maker-api/src/restError"
)

const (
	queryInsertBanner = "INSERT INTO banners(layers) VALUES(?);"
)

// BannerRepository ...
type BannerRepository struct {
	db *sql.DB
}

// Save ...
func (r *BannerRepository) Save(b *models.BannerToDB) (int64, restError.RestError) {
	stmt, err := r.db.Prepare(queryInsertBanner)
	if err != nil {
		return 0, restError.NewInternalServerError("error when trying to save banner", errors.New("database error"))
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(b.Layers)
	if saveErr != nil {
		return 0, restError.NewInternalServerError("error when trying to save banner", errors.New("database error"))
	}

	bannerID, err := insertResult.LastInsertId()
	if err != nil {
		return 0, restError.NewInternalServerError("error when trying to save banner", errors.New("database error"))
	}

	return bannerID, nil
}

// NewBannerRepository ...
func NewBannerRepository(db *sql.DB) *BannerRepository {
	return &BannerRepository{
		db: db,
	}
}
