package routes

import (
	"nhatruong/firstGoBackend/internal/controllers"
	"nhatruong/firstGoBackend/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(authController *controllers.AuthController) *gin.Engine {
	r := gin.Default()

	// Punlic Routes
	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)

	// Protected routes
	auth := r.Group("/auth")
	auth.Use(middlewares.AuthMiddlewares())
	{
		// auth.GET("/profile", authController.Profile)
		// auth.GET("/logout", authController.Logout)
	}

	return r
}
