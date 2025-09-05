package controllers

import (
	"context"
	"net/http"
	"nhatruong/firstGoBackend/internal/models"
	"nhatruong/firstGoBackend/internal/services"
	"nhatruong/firstGoBackend/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (ac *AuthController) Register(c *gin.Context) {
	var req models.RegisterRequest

	// Bind JSON vao struct
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Goi Service
	user, err := ac.AuthService.Register(context.Background(), &req)
	if err != nil {
		utils.Error(c, err.Error(), http.StatusBadRequest)
		return
	}

	utils.Success(c, user)
}

func (ac *AuthController) Login(c *gin.Context) {
	var req models.LoginRequest

	// Bind JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Error(c, "Invailid request", http.StatusBadRequest)
		return
	}

	// Goi service
	token, err := ac.AuthService.Login(context.Background(), &models.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		utils.Error(c, err.Error(), http.StatusBadRequest)
		return
	}

	// Set cookie JWT (Song 24h)
	c.SetCookie("jwt_token", token, 3600*24, "/", "", false, true)

	// Response
	utils.Success(c, gin.H{"token": token})

}
