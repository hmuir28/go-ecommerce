package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hmuir28/go-ecommerce/controllers"
)

func ProductRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/products", controllers.CreateProduct())
	incomingRoutes.GET("/products/:id", controllers.FindProductById())
	incomingRoutes.GET("/products/search", controllers.FindProductByQuery())
}

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
}
