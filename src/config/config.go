package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	DBURL     string
	JwtSecret string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load() // Load tu file .env neu co
	if err != nil {
		log.Println("⚠️ No .env file found, using system env var")
	}

	return &Config{
		DBURL:     os.Getenv("DB_URL"),
		JwtSecret: os.Getenv("JWT_SECRET"),
		Port:      os.Getenv("PORT"),
	}, nil
}
