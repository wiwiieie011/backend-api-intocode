package config

import (
	"os"
	"wiwieie011/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDB() {
	dsn := os.Getenv("DB")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("connection failed")
	}

	DB.AutoMigrate(&models.Group{}, &models.Note{}, &models.Student{})
}
