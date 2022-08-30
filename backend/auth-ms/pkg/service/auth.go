package service

import (
	"context"
	"fullstack/backend/auth-ms/models"
	"fullstack/backend/auth-ms/pkg/repository"
)

type AuthService struct {
	repo repository.Authentication
}

func NewAuthService(repo repository.Authentication) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(userDto models.UserDto) (string, error) {
	return s.repo.CreateUser(context.Background(), models.NewUser(userDto))
}
