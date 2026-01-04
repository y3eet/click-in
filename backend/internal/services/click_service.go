package services

import (
	"github.com/y3eet/click-in/internal/models"
	"github.com/y3eet/click-in/internal/repositories"
)

type ClickService struct {
	repo *repositories.ClickRepository
}

func NewClickService(repo *repositories.ClickRepository) *ClickService {
	return &ClickService{repo: repo}
}

func (s *ClickService) CreateNewClick(click *models.Click) error {
	return s.repo.Create(click)
}

func (s *ClickService) CountClicksByClickableID(clickableID uint) (int64, error) {
	return s.repo.CountClicksByClickableID(clickableID)
}

func (s *ClickService) CountClicksByUserID(userID uint) (int64, error) {
	return s.repo.CountClicksByUserID(userID)
}

func (s *ClickService) CountClicksByClickableAndUser(clickableID uint, userID uint) (int64, error) {
	return s.repo.CountClicksByClickableAndUser(clickableID, userID)
}
