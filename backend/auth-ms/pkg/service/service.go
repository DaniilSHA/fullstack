package service

import (
	"fullstack/backend/auth-ms/models"
	"fullstack/backend/auth-ms/pkg/repository"
)

type Authentication interface {
	CreateUser(user models.User) (int, error)
}

type Service struct {
	Authentication
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
