package models

import (
	"time"
)

type Entity struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Name     string `gorm:"not null" json:"name"`
	ImageURL string `json:"image_url"`
	Mp3URL   string `json:"mp3_url"`

	UserID uint  `gorm:"not null" json:"user_id"`
	User   *User `gorm:"foreignKey:UserID" json:"user"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
