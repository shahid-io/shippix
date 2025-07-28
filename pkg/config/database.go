package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host string
	Port int
	User string
	Pass string
	Name string
	// You can add fields like SSLMode, Timeout, etc. if required
}

func LoadDatabaseConfig() *DatabaseConfig {
	if err := godotenv.Load(".env"); err != nil {
		log.Println(".env file not found, relying on environment variables")
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Printf("Invalid DB_PORT value, defaulting to 5432: %v", err)
		port = 5432
	}

	return &DatabaseConfig{
		Host: os.Getenv("DB_HOST"),
		Port: port,
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Name: os.Getenv("DB_NAME"),
		// Fill in for sslmode or timeout if required
	}
}
