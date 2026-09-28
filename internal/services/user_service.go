package services

import (
	"context"

	"github.com/om1cael/scooter-api/internal/models"
	"github.com/om1cael/scooter-api/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, name, email, password string) (*models.User, error)
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{repository: repository}
}

func (s *userService) Register(ctx context.Context, name, email, password string) (*models.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(passwordHash),
	}

	err = s.repository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
