package config

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	// Connection string
	dsn := "postgres://admin:123456@localhost:5432/mydb"

	// Connect DB
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Ping(context.Background())

	fmt.Println("✅ Connected to PostgreSql!")

	DB = pool
}
