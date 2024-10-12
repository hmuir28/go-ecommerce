package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection = UserData(Client, "users")
var prodCollection = ProductData(Client, "products")
var Client *mongo.Client = DBSet()

func DBSet() *mongo.Client {
	// username := os.Getenv("MONGO_USERNAME")
	// password := os.Getenv("MONGO_PASSWORD")
	// database := os.Getenv("MONGO_DB")
	username := "admin"
	password := "mongo"
	database := "Ecommerce"

	uri := fmt.Sprintf("mongodb://%s:%s@localhost:27017/%s?authSource=admin", username, password, database)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Println("failed to connect to mongodb")
		return nil
	}

	fmt.Println("Successfully connected to mongodb")
	return client
}

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	var userCollection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)

	return userCollection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	var productCollection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)

	return productCollection
}
