package database

import (
	"github.com/y3eet/click-in/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate models
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&models.RefreshToken{}); err != nil {
		return nil, err
	}

	return db, nil
}
