package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config - application configuration.
type Config struct {
	PostgresEndpoint       string
	PostgresMaxConnections uint
}

// New - creates new Config and loads environment variables.
func New() (*Config, error) {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	// Parse environment variables into Config struct
	cfg := &Config{
		PostgresEndpoint:       os.Getenv("POSTGRES_ENDPOINT"),
		PostgresMaxConnections: 30, // default value
	}

	// Override PostgresMaxConnections if the environment variable is set
	if maxConnStr := os.Getenv("POSTGRES_MAX_CONNECTIONS"); maxConnStr != "" {
		maxConn, err := strconv.ParseUint(maxConnStr, 10, 32)
		if err != nil {
			return nil, errors.New("invalid value for POSTGRES_MAX_CONNECTIONS")
		}
		cfg.PostgresMaxConnections = uint(maxConn)
	}

	// Validate required fields
	if cfg.PostgresEndpoint == "" {
		return nil, errors.New("POSTGRES_ENDPOINT is required")
	}

	return cfg, nil
}
