package app

import (
	"fmt"
	"log"
	"os"

	"github.com/spencerfeng/banner_maker-api/src/services"

	"github.com/gin-gonic/gin"
	"github.com/spencerfeng/banner_maker-api/src/controllers"
	"github.com/spencerfeng/banner_maker-api/src/database/sqldb"
	"github.com/spencerfeng/banner_maker-api/src/models"
	"github.com/spencerfeng/banner_maker-api/src/repositories"
)

var (
	mysqlUser     = os.Getenv("MYSQL_USER")
	mysqlPassword = os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase = os.Getenv("MYSQL_DATABASE")
	mysqlHost     = os.Getenv("MYSQL_HOST")
	mysqlPort     = os.Getenv("MYSQL_PORT")
)

// SetupRouter ...
func SetupRouter(bannerRepository models.BannerRepositoryInterface, imageService services.ImageService) *gin.Engine {
	router := gin.Default()

	bannerBaseHandler := controllers.NewBannerBaseHandler(bannerRepository)
	baseImageHandler := controllers.NewBaseImageHandler(imageService)

	router.POST("/banners", bannerBaseHandler.Create)
	router.POST("/images", baseImageHandler.Upload)

	return router
}

// StartApplication ...
func StartApplication() {
	db, err := sqldb.NewDB(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase))
	if err != nil {
		log.Panic(err)
	}

	// initialise repositories
	bannerRepository := repositories.NewBannerRepository(db)

	// initialise services
	imageService, imageServiceErr := services.NewImageService()
	if imageServiceErr != nil {
		log.Panic(imageServiceErr)
	}

	router := SetupRouter(bannerRepository, imageService)

	router.Run(":8082")
}
