package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(dbURL string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("❌ Faild to connect DB: ", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal("❌ Failed to Ping DB: ", err)
	}

	log.Println("Connected to PostgreSQL!")
	return pool
}
