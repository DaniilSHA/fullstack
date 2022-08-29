package service

import "fullstack/backend/auth-ms/pkg/repository"

type Authentication interface {
}

type Service struct {
	Authentication
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
