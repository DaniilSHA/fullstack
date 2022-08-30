package service

import (
	"context"
	"fmt"
	"fullstack/backend/auth-ms/internal/config"
	"fullstack/backend/auth-ms/models"
	"fullstack/backend/auth-ms/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type AuthService struct {
	repo repository.Authentication
	cfg  *config.Config
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

func NewAuthService(repo repository.Authentication, cfg *config.Config) *AuthService {
	return &AuthService{repo: repo, cfg: cfg}
}

func (auth *AuthService) CreateUser(userDto models.UserDto) error {
	_, err := auth.repo.FindByUsername(context.Background(), userDto.Username)
	if err != nil {
		if err.Error() == "not found" {
			_, err := auth.repo.CreateUser(context.Background(), models.NewUser(userDto))
			if err != nil {
				return err
			}
			return nil
		}
	}

	return fmt.Errorf("username is busy")
}

func (auth *AuthService) CheckUserCredentials(userDto models.UserDto) (*models.Tokens, error) {
	user, err := auth.repo.FindByUsername(context.Background(), userDto.Username)
	if err != nil {
		if err.Error() == "not found" {
			return nil, fmt.Errorf("user is not registred")
		}
	}

	checkUser := models.NewUser(userDto)

	if user.Username == checkUser.Username && user.PasswordHash == checkUser.PasswordHash {
		return makeTokens(user.Id, auth.cfg.Secret.Jwtkey)
	}

	return nil, fmt.Errorf("user don't confirmed")
}

func makeTokens(userId string, key string) (*models.Tokens, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})
	signedRefreshToken, err := refreshToken.SignedString([]byte(key))
	if err != nil {
		return nil, fmt.Errorf("error while sigh token: %v", err)
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})
	signedAccessToken, err := accessToken.SignedString([]byte(key))
	if err != nil {
		return nil, fmt.Errorf("error while sigh token: %v", err)
	}

	return &models.Tokens{
		AccessToken:  signedAccessToken,
		RefreshToken: signedRefreshToken,
	}, nil
}
