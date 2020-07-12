package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/spencerfeng/banner_maker-api/src/app"
)

func main() {
	app.StartApplication()
}
