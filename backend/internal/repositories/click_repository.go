package repositories

import (
	"github.com/y3eet/click-in/internal/models"
	"gorm.io/gorm"
)

type ClickRepository struct {
	db *gorm.DB
}

func NewClickRepository(db *gorm.DB) *ClickRepository {
	return &ClickRepository{db: db}
}

func (r *ClickRepository) Create(click *models.Click) error {
	return r.db.Create(click).Error
}

func (r *ClickRepository) CountClicksByClickableID(clickableID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Click{}).Where("clickable_id = ?", clickableID).Count(&count).Error
	return count, err
}

func (r *ClickRepository) CountClicksByUserID(userID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Click{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}

func (r *ClickRepository) CountClicksByClickableAndUser(clickableID uint, userID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Click{}).Where("clickable_id = ? AND user_id = ?", clickableID, userID).Count(&count).Error
	return count, err
}
