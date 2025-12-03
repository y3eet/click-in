package services

import (
	"github.com/y3eet/click-in/internal/models"
	"github.com/y3eet/click-in/internal/repositories"
)

type RefreshTokenService struct {
	repo *repositories.RefreshTokenRepository
}

func NewRefreshTokenService(repo *repositories.RefreshTokenRepository) *RefreshTokenService {
	return &RefreshTokenService{repo: repo}
}

func (r RefreshTokenService) CreateRefreshToken(refreshToken *models.RefreshToken) error {
	token, err := r.repo.FindByUA(refreshToken.UserAgent)
	if err == nil && token != nil {
		token.Token = refreshToken.Token
		token.ExpiresAt = refreshToken.ExpiresAt
		return r.repo.Update(token)
	}
	return r.repo.Create(refreshToken)
}

func (r RefreshTokenService) GetRefreshTokenByToken(token string) (*models.RefreshToken, error) {
	return r.repo.FindByToken(token)
}

func (r RefreshTokenService) DeleteRefreshTokenByToken(token string) error {
	return r.repo.DeleteByToken(token)
}
