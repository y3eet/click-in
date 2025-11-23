package repositories

import (
	"github.com/y3eet/click-in/internal/models"
	"gorm.io/gorm"
)

type EntityRepository struct {
	db *gorm.DB
}

func NewEntityRepository(db *gorm.DB) *EntityRepository {
	return &EntityRepository{db: db}
}

func (r EntityRepository) Create(entity *models.Entity) error {
	return r.db.Create(entity).Error
}

func (r EntityRepository) FindByID(id uint) (*models.Entity, error) {
	var entity models.Entity
	err := r.db.Where(&models.Entity{ID: id}).First(&entity).Error
	return &entity, err
}

func (r EntityRepository) GetAll() ([]models.Entity, error) {
	var entities []models.Entity
	err := r.db.Find(&entities).Error
	return entities, err
}
