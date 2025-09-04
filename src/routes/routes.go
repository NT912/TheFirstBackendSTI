package routes

import (
	"nhatruong/firstGoBackend/src/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(authController *controllers.AuthController) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/register", authController.Register)
		api.POST("/login", authController.Login)
	}

	return r
}
