package routes

import (
	"nhatruong/firstGoBackend/src/config"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(router *gin.Engine, db *pgxpool.Pool, cfg *config.Config) {
	api := router.Group("/api")

	// Auth routes
	authController := controllers
}
