package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string

	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type JWTConfig struct {
	Secret string
	Expiry time.Duration
}

func Load() *Config {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtExpiry, err := time.ParseDuration(os.Getenv("JWT_EXPIRY"))
	if err != nil {
		log.Fatal("Invalid JWT_EXPIRY value in .env")
	}

	return &Config{
		AppName: os.Getenv("APP_NAME"),

		Server: ServerConfig{
			Port: os.Getenv("APP_PORT"),
		},

		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},

		JWT: JWTConfig{
			Secret: os.Getenv("JWT_SECRET"),
			Expiry: jwtExpiry,
		},
	}
}