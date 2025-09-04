package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	DBUrl     string
	JwtSecret string
}

func LoadConfig() *Config {
	err := godotenv.Load() // Load tu file .env neu co
	if err != nil {
		log.Println("⚠️ No .env file found, using system env var")
	}

	return &Config{
		Port:      getEnv("PORT", "8080"),
		DBUrl:     getEnv("DB_URL", ""),
		JwtSecret: getEnv("JWT_SECRET", "changeme"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
