package app

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spencerfeng/banner_maker-api/src/controllers"
	"github.com/spencerfeng/banner_maker-api/src/database/sqldb"
	"github.com/spencerfeng/banner_maker-api/src/repositories"
)

var (
	router        = gin.Default()
	mysqlUser     = os.Getenv("MYSQL_USER")
	mysqlPassword = os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase = os.Getenv("MYSQL_DATABASE")
	mysqlHost     = os.Getenv("MYSQL_HOST")
	mysqlPort     = os.Getenv("MYSQL_PORT")
)

// StartApplication ...
func StartApplication() {
	db, err := sqldb.NewDB(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase))
	if err != nil {
		log.Panic(err)
	}

	// initialise repositories
	bannerRepository := repositories.NewBannerRepository(db)

	bannerBaseHandler := controllers.NewBannerBaseHandler(bannerRepository)

	router.POST("/banners", bannerBaseHandler.Create)

	router.Run(":8082")
}
