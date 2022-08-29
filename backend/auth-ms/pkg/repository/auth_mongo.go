package repository

import "fullstack/backend/auth-ms/models"

type AuthMongo struct {
	db
}

func newAuthMongo(db) *AuthMongo {
	return &AuthMongo{db}
}

func (a *AuthMongo) CreateUser(user models.User) (int, error) {
	return 0, int
}
