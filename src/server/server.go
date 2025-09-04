package server

import (
	"log"
	"nhatruong/firstGoBackend/src/config"
	"nhatruong/firstGoBackend/src/db"
	"nhatruong/firstGoBackend/src/routes"

	"github.com/gin-gonic/gin"
)

func Run() error {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	dbPool, err := db.ConnectDB(cfg.DBURL)
	if err != nil {
		return err
	}
	defer dbPool.Close()

	// Init router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router, dbPool, cfg)

	// Start server
	log.Println("🚀 Server running on port: ", cfg.Port)
	return router.Run(":" + cfg.Port)
}
