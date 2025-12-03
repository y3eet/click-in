package models

import (
	"time"
)

type RefreshToken struct {
	ID     uint   `gorm:"primarykey" json:"id"`
	Token  string `gorm:"not null;unique" json:"token"`
	UserID uint   `gorm:"not null" json:"user_id"`
	User   *User  `gorm:"foreignKey:UserID" json:"user"`

	ExpiresAt time.Time `json:"expires_at"`

	// Fingerprint string `gorm:"size:255" json:"fingerprint"`
	UserAgent      string `gorm:"type:text" json:"user_agent"`
	Browser        string `gorm:"size:100" json:"browser"`
	BrowserVersion string `gorm:"size:100" json:"browser_version"`
	OS             string `gorm:"size:100" json:"os"`
	Platform       string `gorm:"size:100" json:"platform"`
	IsMobile       bool   `json:"is_mobile"`
	IPAddress      string `gorm:"size:45" json:"ip_address"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
