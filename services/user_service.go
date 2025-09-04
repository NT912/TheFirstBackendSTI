package services

import (
	"context"
	"nhatruong/firstGoBackend/config"
	"nhatruong/firstGoBackend/models"
)

func GetAllUser() ([]models.User, error) {
	rows, err := config.DB.Query(context.Background(), "SELECT id, name, email from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return users, nil
}
