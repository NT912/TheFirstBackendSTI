package controllers

import (
	"context"
	"net/http"
	"nhatruong/firstGoBackend/src/services"
	"nhatruong/firstGoBackend/src/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (ac *AuthController) Register(c *gin.Context) {
	var body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		utils.Error(c, "Invalid input", http.StatusBadRequest)
		return
	}

	err := ac.AuthService.Register(context.Background(), strings.TrimSpace(body.Name), strings.TrimSpace(body.Email), body.Password)
	if err != nil {
		utils.Error(c, err.Error(), http.StatusBadRequest)
		return
	}
	utils.Success(c, "User registered successfully!")
}

func (ac *AuthController) Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		utils.Error(c, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := ac.AuthService.Login(context.Background(), strings.TrimSpace(body.Email), body.Password)
	if err != nil {
		utils.Error(c, err.Error(), http.StatusUnauthorized)
		return
	}

	utils.Success(c, user)
}
