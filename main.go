package main

import (
	"nhatruong/firstGoBackend/config"
	"nhatruong/firstGoBackend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// ConnectDB
	config.ConnectDB()

	// Setup gin
	router := gin.Default()

	// Register routes
	routes.UserRoutes(router)

	// Run server
	router.Run(":8080")
}
