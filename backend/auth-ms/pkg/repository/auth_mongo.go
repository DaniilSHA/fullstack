package repository

import (
	"context"
	"fmt"
	"fullstack/backend/auth-ms/models"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (a *AuthMongo) CreateUser(ctx context.Context, user models.UserDto) (string, error) {
	logrus.Debug("creating user")
	result, err := a.collection.InsertOne(ctx, user)
	if err != nil {
		logrus.Errorf("failed to create user %s", err)
	}
	logrus.Debug("converted insertedId to objectId")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	logrus.Trace(user)
	return "", fmt.Errorf("failed to convert object id to hex, objectId:%s", oid)
}

func (a *AuthMongo) FindById(ctx context.Context, id string) (u models.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("failed to convert hex to objectId, hex:%s", id)
	}
	filter := bson.M{"_id": oid}

	result := a.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return u, fmt.Errorf("failed to find user by id: %s due to error: %s", id, err)
	}
	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user from db (id: %s) due to error: %s", id, err)
	}

	return u, nil
}

func (a *AuthMongo) UpdateUser(ctx context.Context, user models.User) error {
	return nil
}

func (a *AuthMongo) DeleteUser(ctx context.Context, id string) error {
	return nil
}
