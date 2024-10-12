package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddProductToCart(c context.Context, prodCollection *mongo.Collection, userCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {
	return nil
}

func RemoveCartItem(c context.Context, prodCollection *mongo.Collection, userCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {
	return nil
}

func FindProductFromCart(ctx context.Context, filteredMatch bson.D, unwind bson.D, grouping bson.D) (cur *mongo.Cursor, err error) {
	return userCollection.Aggregate(ctx, mongo.Pipeline{
		filteredMatch,
		unwind,
		grouping,
	})
}

func BuyItemFromCart(c context.Context, userCollection *mongo.Collection, userID string) error {
	return nil
}

func InstantBuyer(c context.Context, prodCollection *mongo.Collection, userCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {
	return nil
}
