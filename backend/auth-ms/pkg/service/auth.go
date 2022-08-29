package service

import (
	"fullstack/backend/auth-ms/models"
	"fullstack/backend/auth-ms/pkg/repository"
)

type AuthService struct {
	repo repository.Repository
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) {
	return s.repo.CreateUser(user)
}
