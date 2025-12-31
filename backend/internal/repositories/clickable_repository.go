package repositories

import (
	"github.com/y3eet/click-in/internal/models"
	"gorm.io/gorm"
)

type ClickableRepository struct {
	db *gorm.DB
}

func NewClickableRepository(db *gorm.DB) *ClickableRepository {
	return &ClickableRepository{db: db}
}

func (r ClickableRepository) Create(clickable *models.Clickable) error {
	return r.db.Create(clickable).Error
}

func (r ClickableRepository) FindByID(id uint) (*models.Clickable, error) {
	var clickable models.Clickable
	err := r.db.Where(&models.Clickable{ID: id}).First(&clickable).Error
	return &clickable, err
}

func (r ClickableRepository) FindByName(name string) (*models.Clickable, error) {
	var clickable models.Clickable
	err := r.db.Where(&models.Clickable{Name: name}).First(&clickable).Error
	return &clickable, err
}

func (r ClickableRepository) GetAll() ([]models.Clickable, error) {
	var clickable []models.Clickable
	err := r.db.Find(&clickable).Error
	return clickable, err
}
