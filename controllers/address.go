package controllers

import (
	"time"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/hmuir28/go-ecommerce/models"
	"github.com/hmuir28/go-ecommerce/database"
)

func AddAddress() {

}

func EditHomeAddress() {

}

func EditWorkAddress() {

}

func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		addressId := c.Query("id")

		if addressId == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid Search Index"})
			c.Abort()
			return
		}

		addresses := make([]models.Address, 0)

		id, err := primitive.ObjectIDFromHex(addressId)

		if err != nil {
			c.IndentedJSON(500, "Internal Server Error")
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.D{primitive.E{Key: "_id", Value: id}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}

		_, err = database.UpdateProduct(ctx, filter, update)

		if err != nil {
			c.IndentedJSON(404, "Wrong command")
			return
		}

		defer cancel()

		ctx.Done()
		c.IndentedJSON(200, "Address deleted successfully")
	}
}
