package mongodb

import (
	"context"
	"errors"
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

func NewAuthMongo(database *mongo.Database, collection string) *AuthMongo {
	return &AuthMongo{
		collection: database.Collection(collection),
	}
}

func (a *AuthMongo) CreateUser(ctx context.Context, user *models.User) (string, error) {
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
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			//TODO ErrEntityNotFound
			return u, fmt.Errorf("not found")
		}
		return u, fmt.Errorf("failed to find user by id: %s due to error: %s", id, err)
	}
	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user from db (id: %s) due to error: %s", id, err)
	}

	return u, nil
}

func (a *AuthMongo) UpdateUser(ctx context.Context, user models.User) error {
	oid, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return fmt.Errorf("failed to convert hex to objectId, hex: %s", user.Id)
	}

	filter := bson.M{"_id": oid}

	userBytes, err := bson.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal document, error: %v", err)
	}

	var updateUserObj bson.M
	err = bson.Unmarshal(userBytes, &updateUserObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal user bytes, error: %v", err)
	}

	delete(updateUserObj, "_id")

	update := bson.M{
		"$set": updateUserObj,
	}

	result, err := a.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to execute update user query, error: %v", err)
	}

	if result.MatchedCount == 0 {
		//TODO ErrEntityNotFound
		return fmt.Errorf("not found")
	}
	logrus.Tracef("matched %d documents and modified %d documents", result.MatchedCount, result.ModifiedCount)

	return nil
}

func (a *AuthMongo) DeleteUser(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert hex to objectId, hex: %s", id)
	}

	filter := bson.M{"_id": oid}

	result, err := a.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete, err: %v", err)
	}

	if result.DeletedCount == 0 {
		//TODO ErrEntityNotFound
		return fmt.Errorf("not found")
	}

	logrus.Tracef("deleted %d documents", result.DeletedCount)

	return nil
}
