package database

import (
	"context"

	"github.com/hmuir28/go-ecommerce/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CountUserDocumentsByKeyValue(ctx context.Context, key, value string) (int64, error) {
	return userCollection.CountDocuments(ctx, bson.M{key: value})
}

func CreateUser(ctx context.Context, user models.User) error {
	_, insertErr := userCollection.InsertOne(ctx, user)
	return insertErr
}

func FindUserByEmail(ctx context.Context, user *models.User, email string) error {
	return userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&user)
}
