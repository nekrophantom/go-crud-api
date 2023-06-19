package db

import (
	"fmt"
	// "log"
	"os"

	// "github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {

	dbHost		:= os.Getenv("DB_HOST")
	dbPort 		:= os.Getenv("DB_PORT")
	dbName 		:= os.Getenv("DB_NAME")
	dbUsername 	:= os.Getenv("DB_USERNAME")
	dbPassword 	:= os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db

	return nil
}

func GetDB() *gorm.DB {
	return DB
}