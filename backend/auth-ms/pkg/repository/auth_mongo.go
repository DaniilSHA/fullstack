package repository

import (
	"context"
	"fullstack/backend/auth-ms/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthMongo struct {
	collection *mongo.Collection
}

func newAuthMongo(database *mongo.Database, collection string) *AuthMongo {
	return &AuthMongo{
		collection: database.Collection(collection),
	}
}

func (a *AuthMongo) CreateUser(ctx context.Context, user models.UserDto) (int, error) {
	result, err := a.collection.InsertOne(ctx, user)
	if err != nil {
		logrus.Errorf("failed to create user %s", err)
	}
	return 0, nil
}

func (a *AuthMongo) FindById(ctx context.Context, id string) (models.User, error) {
	return nil, nil
}

func (a *AuthMongo) UpdateUser(ctx context.Context, user models.User) error {
	return nil
}

func (a *AuthMongo) DeleteUser(ctx context.Context, id string) error {
	return nil
}
