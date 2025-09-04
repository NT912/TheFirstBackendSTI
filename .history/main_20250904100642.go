package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
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

	fmt.Println("✅ Connected to PostgreSql!")

	// Test Query
	var greeting string
	err = conn.QueryRow(context.Background(), "SELECT 'Hello from PostgreSql!").Scan(&greeting)
}
