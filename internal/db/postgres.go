package db

import (
	"database/sql"
	"fmt"
	"time"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase(dbURL string) (*Database, error) {
	// Mo ket noi
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("❌ Failed to connect DB: %v", err)
	}

	// cau hinh connection pool
	db.SetMaxOpenConns(25)                 // so ket noi toi da
	db.SetMaxIdleConns(5)                  // so ket noi nhan roi giu san
	db.SetConnMaxLifetime(5 * time.Minute) // tu Refresh ket noi sau 5 phut

	// xac nhan da ket noi
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("❌ Failed to Ping DB: %v", err)
	}

	return &Database{DB: db}, nil
}
