package repository

import (
	"context"
	"log"
	"nhatruong/firstGoBackend/src/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	_, err := r.db.Exec(ctx, "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
	return err
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	log.Printf("Querying for user with email: %s", email)
	row := r.db.QueryRow(ctx, "SELECT user_id, name, email, password, created_at FROM users WHERE LOWER(email)=LOWER($1)", email)
	u := &models.User{}
	err := row.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.CreatedAt)
	if err != nil {
		log.Printf("Error scanning user data: %v", err)
		return nil, err
	}

	return u, nil
}
