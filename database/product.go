package database

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/hmuir28/go-ecommerce/models"
)

var (
	ErrCantFindProduct = errors.New("can't find the product")
	ErrCantDecodeProducts = errors.New("can't find the product")
	ErrUserIdIsNotValid = errors.New("this user is not valid")
	ErrCantUpdateUser = errors.New("cannot add this product to the cart")
	ErrCantRemoveItemCart = errors.New("cannot remove this item from the cart")
	ErrCantGetItem = errors.New("was unable to get the item from the cart")
	ErrCantBuyCartItem = errors.New("cannot update the purchase")
)

func CreateProduct(ctx context.Context, product models.Product) error {
	_, insertErr := prodCollection.InsertOne(ctx, product)
	return insertErr
}

func FindProductByID(ctx context.Context, product *models.Product, id primitive.ObjectID) error {
	return prodCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
}

