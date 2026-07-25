package config

import (
	"errors"
)

func (c *Config) Validate() error {

	if c.Server.Port == "" {
		return errors.New("APP_PORT is required")
	}

	if c.Database.Host == "" {
		return errors.New("DB_HOST is required")
	}

	if c.Database.User == "" {
		return errors.New("DB_USER is required")
	}

	if c.Database.Password == "" {
		return errors.New("DB_PASSWORD is required")
	}

	if c.Database.Name == "" {
		return errors.New("DB_NAME is required")
	}

	if c.JWT.Secret == "" {
		return errors.New("JWT_SECRET is required")
	}

	return nil
}