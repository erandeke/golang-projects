package routes

import (
	"go-ecommerce/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signUp", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addProduct", controllers.AddProduct())
	incomingRoutes.GET("/users/productView", controllers.ViewProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
}
