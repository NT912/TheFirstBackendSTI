package services

import (
	"context"
	"errors"
	"nhatruong/firstGoBackend/internal/models"
	"nhatruong/firstGoBackend/internal/repository"
	"nhatruong/firstGoBackend/internal/utils"
	"nhatruong/firstGoBackend/internal/validation"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrEmailUsed         = errors.New("❌ Email already used")
	ErrInvalidCredential = errors.New("❌ Invailid email or password")
)

type AuthService struct {
	userRepo    *repository.UserRepository
	jwtSecret   string
	tokenExpiry time.Duration
	validate    *validator.Validate
}

func NewAuthService(repo *repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:    repo,
		jwtSecret:   jwtSecret,
		tokenExpiry: 24 * time.Hour,
		validate:    validator.New(),
	}
}

func (s *AuthService) Register(ctx context.Context, req *models.RegisterRequest) (*models.User, error) {
	// Validate input struct (email format, min password) bang validator
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	// Check password policy them bang utils.ValidatePassword
	if err := validation.ValidatePassword(req.Password); err != nil {
		return nil, err
	}

	// Nomalize email
	email := strings.TrimSpace(strings.ToLower(req.Email))

	// Kiem tra da co mail chua
	existingUser, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, ErrEmailUsed
	}

	// hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Tao user struct
	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	// Luu vao DB
	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	// Tra user ve, khong bao gom HashPassword
	user.Password = ""
	return user, nil
}

func (s *AuthService) Login(ctx context.Context, req *models.LoginRequest) (string, error) {
	// Validate input
	if err := s.validate.Struct(req); err != nil {
		return "", err
	}

	// Normalize email
	email := strings.TrimSpace(strings.ToLower(req.Email))

	// lay user tu DB
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil || user == nil {
		return "", ErrInvalidCredential
	}

	// Check Password: dam bao goi dung thu tu (plain, hash)
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return "", ErrInvalidCredential
	}

	// Tao JWT
	token, err := s.generateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) generateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.Id,
		"email":   user.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // token 24h
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return t.SignedString([]byte(s.jwtSecret))
}
