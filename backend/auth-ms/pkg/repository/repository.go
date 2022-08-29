package repository

import "fullstack/backend/auth-ms/models"

type Authentication interface {
	CreateUser(user models.User) (int, error)
}

type Repository struct {
	Authentication
}

func NewRepository() *Repository {
	return &Repository{}
}
