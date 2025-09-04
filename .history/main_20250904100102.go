package main

import "github.com/jackc/pgx"

func main() {
	// Connection string
	dsn := "postgres://admin:123456@localhost:5432/mydb"

	// Connect DB
	conn, err := pgx.Connect()
}
