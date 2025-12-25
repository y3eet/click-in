package services

import (
	"github.com/y3eet/click-in/internal/models"
	"github.com/y3eet/click-in/internal/repositories"
)

type ClickableService struct {
	repo *repositories.ClickableRepository
}

func NewClickableService(repo *repositories.ClickableRepository) *ClickableService {
	return &ClickableService{repo: repo}
}

func (s ClickableService) CreateNewClickable(clickable *models.Clickable) error {
	return s.repo.Create(clickable)
}

func (s ClickableService) GetAllClickable() ([]models.Clickable, error) {
	return s.repo.GetAll()
}
