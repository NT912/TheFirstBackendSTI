package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"Status": "success", "data": data})
}

func Error(c *gin.Context, msg string, code int) {
	c.JSON(code, gin.H{"Status": "Error", "message": msg})
}
