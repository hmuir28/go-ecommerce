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

func FindProducts() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		cursor, err := database.FindProducts(ctx)

		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong!"})
			return
		}

		var productList []models.Product

		err = cursor.All(ctx, &productList)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		defer cursor.Close(ctx)

		if err := cursor.Err(); err != nil {
			log.Println(err)
			c.IndentedJSON(400, "invalid")
			return
		}

		defer cancel()

		c.JSON(http.StatusFound, productList)
	}
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
	return func(c *gin.Context) {
		var productList []models.Product
		queryParam := c.Query("name")

		if queryParam == "" {
			log.Println("query is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid search index"})
			c.Abort()
			return
		}
		
		var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)

		defer cancel()

		cursor, err := database.FindProductByName(ctx, queryParam)

		if err != nil {
			c.IndentedJSON(404, "something went wrong while fetching data")
			return
		}

		err = cursor.All(ctx, &productList)

		if err != nil {
			log.Println(err)
			c.IndentedJSON(400, "invalid")
			return
		}

		defer cursor.Close(ctx)

		if err := cursor.Err(); err != nil {
			log.Println(err)
			c.IndentedJSON(400, "invalid")
			return
		}

		defer cancel()

		c.JSON(http.StatusFound, productList)
	}
}
