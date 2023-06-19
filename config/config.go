package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func LoadConfig(e *echo.Echo) {
	err := godotenv.Load()

	if err != nil {
		log.Println("No env file found")
	}
}