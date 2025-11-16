package services

import (
	"github.com/y3eet/click-in/internal/models"
	"github.com/y3eet/click-in/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(email, password, username string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    email,
		Password: string(hashedPassword),
		Username: username,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}
