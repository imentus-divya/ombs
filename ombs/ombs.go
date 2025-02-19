package main

import (
	"context"
	"fmt"
	"ombs/internal/config"
	controller "ombs/internal/controller"
	"ombs/internal/service"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

const (
	postgresConnectionTimeout = 10 * time.Second
)

func main() {

	e := echo.New()
	// Load configuration
	cfg, err := config.New() // Loads configuration from environment variables
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	_, err = postgresql(cfg.PostgresEndpoint, int32(cfg.PostgresMaxConnections))
	if err != nil {
		fmt.Println("Error connecting to PostgreSQL:", err)
		return
	}

	// // Initialize Redis client using the loaded config
	// _, err = newRedisClient(cfg.Redis)
	// if err != nil {
	// 	fmt.Println("Error connecting to Redis:", err)
	// 	return
	// }

	// web3Connection()

	// Initialize Services
	authService := service.NewAuthService()

	// Register Controllers
	controller.NewAuthController(e, authService)

	e.Logger.Fatal(e.Start(":3000"))
}

func postgresql(connStr string, maxConns int32) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), postgresConnectionTimeout)
	defer cancel()

	// init postgres database pool
	cfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse postgesql connection config: %w", err)
	}

	cfg.MaxConns = maxConns
	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to establish connection to postgresql: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to get response from postgresql: %w", err)
	}

	return pool, nil
}
