package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	ProviderID string         `gorm:"uniqueIndex;not null" json:"provider_id"`
	Email      string         `gorm:"uniqueIndex;not null" json:"email"`
	Username   string         `json:"username"`
	AvatarURL  string         `json:"avatar_url"`
	Provider   string         `json:"provider"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
