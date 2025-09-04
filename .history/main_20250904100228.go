package main

import (
	"context"
	"log"

	"github.com/jackc/pgx"
)

func main() {
	// Connection string
	dsn := "postgres://admin:123456@localhost:5432/mydb"

	// Connect DB
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())
}
