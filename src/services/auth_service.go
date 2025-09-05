package services

import (
	"context"
	"errors"
	"log"
	"nhatruong/firstGoBackend/src/models"
	"nhatruong/firstGoBackend/src/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: repo}
}

func (s *AuthService) Register(ctx context.Context, name, email, password string) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &models.User{Name: name, Email: email, Password: string(hashed)}
	return s.userRepo.Create(ctx, user)
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*models.User, error) {
	log.Printf("AuthService: Attempting to log in user with email: '%s'", email)
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		log.Printf("AuthService: Error finding user by email: %v", err)
		return nil, errors.New("User not found")
	}

	log.Printf("AuthService: User found: %+v", user)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Printf("AuthService: Invalid password for user %s", email)
		return nil, errors.New("Invalid Password")
	}

	return user, nil
}
