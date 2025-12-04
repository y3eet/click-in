package repositories

import (
	"github.com/y3eet/click-in/internal/models"
	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

func (r RefreshTokenRepository) Create(refreshToken *models.RefreshToken) error {
	return r.db.Create(refreshToken).Error
}

func (r RefreshTokenRepository) FindByUA(userAgent string, userId uint) (*models.RefreshToken, error) {
	var token models.RefreshToken
	err := r.db.Where(&models.RefreshToken{UserAgent: userAgent, UserID: userId}).First(&token).Error
	return &token, err
}

func (r RefreshTokenRepository) GetFirst(refreshTokenModel *models.RefreshToken) (*models.RefreshToken, error) {
	var token models.RefreshToken
	err := r.db.Where(refreshTokenModel).First(&token).Error
	return &token, err
}

func (r RefreshTokenRepository) FindByToken(refreshToken string) (*models.RefreshToken, error) {
	var token models.RefreshToken
	err := r.db.Where(&models.RefreshToken{Token: refreshToken}).First(&token).Error
	return &token, err
}

func (r RefreshTokenRepository) Update(refreshToken *models.RefreshToken) error {
	return r.db.Save(refreshToken).Error
}

func (r RefreshTokenRepository) DeleteByToken(refreshToken string) error {
	return r.db.Where(&models.RefreshToken{Token: refreshToken}).Delete(&models.RefreshToken{}).Error
}

func (r RefreshTokenRepository) Delete(id uint) {
	r.db.Delete(&models.User{}, id)
}
