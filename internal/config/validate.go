package config

import (
	"errors"
	"fmt"
)

// Validate checks whether the loaded configuration is valid.
func (c *Config) Validate() error {
	if err := validateApplication(c); err != nil {
		return err
	}

	if err := validateDatabase(c); err != nil {
		return err
	}

	if err := validateRedis(c); err != nil {
		return err
	}

	if err := validateJWT(c); err != nil {
		return err
	}

	if err := validateSession(c); err != nil {
		return err
	}

	return nil
}

func validateApplication(c *Config) error {
	if c.Application.Name == "" {
		return errors.New("application name cannot be empty")
	}

	if c.Application.Port <= 0 || c.Application.Port > 65535 {
		return fmt.Errorf("invalid application port: %d", c.Application.Port)
	}

	return nil
}

func validateDatabase(c *Config) error {
	if c.Database.Host == "" {
		return errors.New("database host cannot be empty")
	}

	if c.Database.Name == "" {
		return errors.New("database name cannot be empty")
	}

	if c.Database.User == "" {
		return errors.New("database user cannot be empty")
	}

	if c.Database.Port <= 0 || c.Database.Port > 65535 {
		return fmt.Errorf("invalid database port: %d", c.Database.Port)
	}

	return nil
}

func validateRedis(c *Config) error {
	if c.Redis.Host == "" {
		return errors.New("redis host cannot be empty")
	}

	if c.Redis.Port <= 0 || c.Redis.Port > 65535 {
		return fmt.Errorf("invalid redis port: %d", c.Redis.Port)
	}

	return nil
}

func validateJWT(c *Config) error {
	if c.JWT.Secret == "" {
		return errors.New("JWT secret cannot be empty")
	}

	return nil
}

func validateSession(c *Config) error {
	if c.Session.Secret == "" {
		return errors.New("session secret cannot be empty")
	}

	return nil
}
