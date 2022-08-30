package service

import (
	"context"
	"fmt"
	"fullstack/backend/auth-ms/models"
	"fullstack/backend/auth-ms/pkg/repository"
)

type AuthService struct {
	repo repository.Authentication
}

func NewAuthService(repo repository.Authentication) *AuthService {
	return &AuthService{repo: repo}
}

func (auth *AuthService) CreateUser(userDto models.UserDto) (string, error) {
	_, err := auth.repo.FindByUsername(context.Background(), userDto.Username)
	if err != nil {
		if err.Error() == "not found" {
			return auth.repo.CreateUser(context.Background(), models.NewUser(userDto))
		}
	}

	return "", fmt.Errorf("username is busy")
}

func (auth *AuthService) CheckUserCredentials(userDto models.UserDto) (string, error) {
	user, err := auth.repo.FindByUsername(context.Background(), userDto.Username)
	if err != nil {
		if err.Error() == "not found" {
			return "", fmt.Errorf("user is not registred")
		}
	}

	checkUser := models.NewUser(userDto)

	if user.Username == checkUser.Username && user.PasswordHash == checkUser.PasswordHash {
		return "", nil
	}

	return "", fmt.Errorf("user don't confirmed")
}
