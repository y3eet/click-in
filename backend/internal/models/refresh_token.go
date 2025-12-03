package models

import (
	"time"
)

type RefreshToken struct {
	ID     uint   `gorm:"primarykey" json:"id"`
	Token  string `gorm:"not null;unique" json:"token"`
	UserID uint   `gorm:"not null" json:"user_id"`
	User   *User  `gorm:"foreignKey:UserID" json:"user"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
