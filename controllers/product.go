package controllers

import (
	"time"
	"context"
	"errors"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/hmuir28/go-ecommerce/models"
	"github.com/hmuir28/go-ecommerce/database"
)

func CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		var product models.Product

		if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		insertErr := database.CreateProduct(ctx, product)

		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "the user did not get created"})
			return
		}

		defer cancel()

		c.JSON(http.StatusCreated, "Product created successfully!")
	}
}

func ProductViewerAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func FindProductById() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		productId := c.Param("id")

		defer cancel()

		if productId == "" {
			log.Println("product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return
		}

		var foundProduct models.Product
		
		id, err := primitive.ObjectIDFromHex(productId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}

		err = database.FindProductByID(ctx, &foundProduct, id)

		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Product not found"})
			return
		}

		c.JSON(http.StatusFound, foundProduct)
	}
}

func FindProductByQuery() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
