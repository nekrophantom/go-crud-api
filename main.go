package main

import (
	"crud_api/config"
	"crud_api/db"
	"crud_api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	
	e := echo.New()

	config.LoadConfig(e)

	err := db.Init()
	if err != nil {
		panic(err)
	}
	// defer db.GetDB()

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":9000"))
}