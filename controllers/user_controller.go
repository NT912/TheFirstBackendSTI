package controllers

import (
	"fmt"
	"net/http"
	"nhatruong/firstGoBackend/services"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := services.GetAllUser()
	if err != nil {
		fmt.Println("DB Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}
