package db

import (
	"os"

	"example.com/blog-app-backend-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDB() (err error) {

	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	err = DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Category{}, &models.Comment{}, &models.Role{})

	return err
}
