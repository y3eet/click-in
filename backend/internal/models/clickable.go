package models

import (
	"time"
)

type Clickable struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Name     string `gorm:"not null;unique" json:"name"`
	ImageURL string `json:"image_url"`
	ImageKey string `json:"image_key"`

	Mp3URL string `json:"mp3_url"`
	Mp3Key string `json:"mp3_key"`

	UserID uint  `gorm:"not null" json:"user_id"`
	User   *User `gorm:"foreignKey:UserID" json:"user"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
