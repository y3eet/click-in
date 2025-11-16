package database

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/y3eet/click-in/internal/models"
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

    return db, nil
}