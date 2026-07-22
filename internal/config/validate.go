package config

import (
	"errors"
)

func (c *Config) Validate() error {

	if c.AppPort == "" {
		return errors.New("APP_PORT is required")
	}

	if c.DBHost == "" {
		return errors.New("DB_HOST is required")
	}

	if c.DBUser == "" {
		return errors.New("DB_USER is required")
	}

	if c.DBPassword == "" {
		return errors.New("DB_PASSWORD is required")
	}

	if c.DBName == "" {
		return errors.New("DB_NAME is required")
	}

	if c.JWTSecret == "" {
		return errors.New("JWT_SECRET is required")
	}

	return nil
}