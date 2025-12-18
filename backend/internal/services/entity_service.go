package services

import (
	"github.com/y3eet/click-in/internal/models"
	"github.com/y3eet/click-in/internal/repositories"
)

type EntityService struct {
	repo *repositories.EntityRepository
}

func NewEntityService(repo *repositories.EntityRepository) *EntityService {
	return &EntityService{repo: repo}
}

func (s EntityService) CreateNewEntity(entity *models.Entity) error {
	return s.repo.Create(entity)
}

func (s EntityService) GetAllEntity() ([]models.Entity, error) {
	return s.repo.GetAll()
}
