package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hmuir28/go-ecommerce/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview/:id", controllers.FindProductById())
	incomingRoutes.GET("/users/search", controllers.FindProductByQuery())
}
