package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateProduct(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return prodCollection.UpdateOne(ctx, filter, update)
}
