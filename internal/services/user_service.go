package services

import (
	"context"

	"github.com/om1cael/scooter-api/internal/models"
	"github.com/om1cael/scooter-api/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, user *models.User) error
}

type userService struct {
	repository repositories.UserRepository
}

func (s *userService) Register(ctx context.Context, name, email, password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(passwordHash),
	}

	return s.repository.Create(ctx, user)
}
