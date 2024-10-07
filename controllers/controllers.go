package controllers

import (
	"log"
	"time"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/go-playground/validator/v10"
	
	"github.com/hmuir28/go-ecommerce/models"
)

func HashPassword(password string) string {
	return ""
}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {
	return true, ""
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100 * time.Second)

		defer cancel()

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validate := validator.New()

		validationErr := validate.Struct(user)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}

		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		}

		count, err = UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})

		defer cancel()

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "this phone no. is already in use"})
			return
		}

		password := HashPassword(*user.Password)

		user.Password = &password
		user.Created_At = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()
		
		token, refreshtoken, _ := generate.TokenGenerator(*user.Email, *user.First_Name, *user.Last_Name, user.User_ID)
		user.Token = &token
		user.Refresh_Token = &refreshtoken
		user.UserCart = make([]models.ProductUser, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)

		_, insertErr := UserCollection.InsertOne(ctx, user)

		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "the user did not get created"})
			return
		}

		defer cancel()

		c.JSON(http.StatusCreated, "Successfully signed up!")
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
	
		var ctx, cancel = context.WithTimeout(context.Background(), 100 *time.Second)

		defer cancel()

		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)

		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password incorrect"})
			return
		}

		PasswordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)

		defer cancel()

		if !PasswordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			fmt.Println(msg)
			return
		}

		token, refreshtoken, _ := generate.TokenGenerator(*foundUser.Email, *foundUser.First_Name, *foundUser.Last_Name, foundUser.User_ID)

		defer cancel()

		generate.UpdateAllTokens(token, refreshtoken, foundUser.User_ID)

		c.JSON(http.StatusFound, foundUser)
	}
}

func ProductViewerAdmin() gin.HandlerFunc {

}

func SearchProduct() gin.HandlerFunc {
	
}

func SearchProductByQuery() gin.HandlerFunc {
	
}
