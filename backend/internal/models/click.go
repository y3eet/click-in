package models

import "time"

type Click struct {
	ID uint `gorm:"primarykey" json:"id"`

	UserID uint  `gorm:"not null" json:"user_id"`
	User   *User `gorm:"foreignKey:UserID" json:"user"`

	ClickableID uint       `gorm:"not null" json:"clickable_id"`
	Clickable   *Clickable `gorm:"foreignKey:ClickableID" json:"clickable"`

	CreatedAt time.Time `json:"created_at"`
}
