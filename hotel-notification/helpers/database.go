package helpers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hotel-notification/internal/models"
	"os"
)

var DB *gorm.DB

func SetupDB() {
	var err error

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("NOTIFICATION_DB_HOST"),
		os.Getenv("NOTIFICATION_DB_PORT"),
		os.Getenv("NOTIFICATION_DB_USER"),
		os.Getenv("NOTIFICATION_DB_PASSWORD"),
		os.Getenv("NOTIFICATION_DB_NAME"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		Logger.Error("Failed to connect to database: ", err)
		return
	}

	err = DB.AutoMigrate(&models.NotificationTemplate{}, &models.NotificationHistory{})
	if err != nil {
		Logger.Error("Failed to migrate database: ", err)
		return
	}

	Logger.Info("Connected to database")
}
