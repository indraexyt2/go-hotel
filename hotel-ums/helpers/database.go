package helpers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func SetupDB() {
	var err error

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("UMS_DB_HOST"),
		os.Getenv("UMS_DB_PORT"),
		os.Getenv("UMS_DB_USER"),
		os.Getenv("UMS_DB_PASSWORD"),
		os.Getenv("UMS_DB_NAME"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		Logger.Error("Failed to connect to database: ", err)
		return
	}

	Logger.Info("Connected to database")
}
