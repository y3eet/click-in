package services

import (
	"time"

	"github.com/y3eet/click-in/internal/constants"
	"github.com/y3eet/click-in/internal/models"
	"github.com/y3eet/click-in/internal/repositories"
)

type RefreshTokenService struct {
	repo *repositories.RefreshTokenRepository
}

func NewRefreshTokenService(repo *repositories.RefreshTokenRepository) *RefreshTokenService {
	return &RefreshTokenService{repo: repo}
}

func (r RefreshTokenService) RefreshToken(oldToken *models.RefreshToken, newToken string) error {
	oldToken.Token = newToken
	oldToken.ExpiresAt = time.Now().Add(constants.RefreshTokenTTL)
	return r.repo.Update(oldToken)
}

func (r RefreshTokenService) CreateRefreshToken(refreshToken *models.RefreshToken) error {
	existingToken, err := r.repo.FindByUA(refreshToken.UserAgent, refreshToken.UserID)
	if err == nil && existingToken != nil {
		if err := r.repo.DeleteByToken(existingToken.Token); err != nil {
			return err
		}
	}
	return r.repo.Create(refreshToken)
}

func (r RefreshTokenService) GetRefreshTokenByToken(token string) (*models.RefreshToken, error) {
	return r.repo.FindByToken(token)
}

func (r RefreshTokenService) DeleteRefreshTokenByToken(token string) error {
	return r.repo.DeleteByToken(token)
}
