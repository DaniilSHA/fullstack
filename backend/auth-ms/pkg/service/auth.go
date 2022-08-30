package service

import (
	"fullstack/backend/auth-ms/models"
	"fullstack/backend/auth-ms/pkg/repository"
)

type AuthService struct {
	repo repository.Authentication
}

func NewAuthService(repo repository.Authentication) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.UserDto) (models.User, error) {
	return s.repo.CreateUser(user)
}
