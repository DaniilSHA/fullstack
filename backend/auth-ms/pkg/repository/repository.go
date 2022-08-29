package repository

import (
	"context"
	"fullstack/backend/auth-ms/models"
)

type Authentication interface {
	CreateUser(ctx context.Context, user models.User) (int, error)
	FindById(ctx context.Context, id string) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, id string) error
}

type Repository struct {
	Authentication
}

func NewRepository() *Repository {
	return &Repository{}
}
