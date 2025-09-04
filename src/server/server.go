package server

import (
	"nhatruong/firstGoBackend/src/config"
	"nhatruong/firstGoBackend/src/controllers"
	"nhatruong/firstGoBackend/src/db"
	"nhatruong/firstGoBackend/src/repository"
	"nhatruong/firstGoBackend/src/routes"
	"nhatruong/firstGoBackend/src/services"
)

func Run() {
	cfg := config.LoadConfig()
	dbPool := db.ConnectDB(cfg.DBUrl)

	userRepo := repository.NewUserRepository(dbPool)
	authSerive := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authSerive)

	r := routes.SetupRoutes(authController)
	r.Run(":" + cfg.Port)
}
