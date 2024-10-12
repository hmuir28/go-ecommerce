package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddProductToCart(c context.Context, prodCollection *mongo.Collection, userCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {
	return nil
}

func RemoveCartItem(c context.Context, prodCollection *mongo.Collection, userCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {
	return nil
}

func GetItemFromCart() {

}

func BuyItemFromCart(c context.Context, userCollection *mongo.Collection, userID string) error {
	return nil
}

func InstantBuyer(c context.Context, prodCollection *mongo.Collection, userCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {
	return nil
}
