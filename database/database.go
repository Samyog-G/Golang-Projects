package database

import (
	"log"

	"github.com/Samyog-G/Golang-Projects/blog-system/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("./blog_system.db"), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.BlogPost{})
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}
}
