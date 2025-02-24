package main

import (
	"fmt"
	"ombs/internal/config"
	controller "ombs/internal/controller"
	"ombs/internal/service"
	"time"

	"github.com/labstack/echo/v4"

	"ombs/internal/drivers/postgres/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	// _, err = postgresql(cfg.PostgresEndpoint, int32(cfg.PostgresMaxConnections))
	_, err = postgresql(cfg.PostgresEndpoint)

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

func postgresql(connStr string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to establish connection to PostgreSQL: %w", err)
	}
	psqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get SQL DB instance: %w", err)
	}

	// Set connection pool settings
	psqlDB.SetMaxOpenConns(10)
	psqlDB.SetMaxIdleConns(5)
	psqlDB.SetConnMaxLifetime(time.Hour)

	psqlDB.Ping()
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.OracleNodeStatus{})

	// Insert initial data
	seedDatabase(db)

	// Now migrate the dependent table
	db.AutoMigrate(&models.OracleNode{})
	db.AutoMigrate(&models.JWTIssuance{})
	db.AutoMigrate(&models.Leader{})
	db.AutoMigrate(&models.LeaderTransaction{})
	db.AutoMigrate(&models.Round{})
	db.AutoMigrate(&models.Batch{})
	db.AutoMigrate(&models.Signature{})

	return db, nil
}

func seedDatabase(db *gorm.DB) {
	roles := []models.Role{
		{ID: 1, RoleName: "Admin"},
		{ID: 2, RoleName: "Leader"},
		{ID: 3, RoleName: "Secondary Leader"},
	}

	statuses := []models.OracleNodeStatus{
		{ID: 1, StatusName: "Active"},
		{ID: 2, StatusName: "Inactive"},
		{ID: 3, StatusName: "Suspended"},
	}

	// Insert roles (ignore duplicates)
	for _, role := range roles {
		db.FirstOrCreate(&role, models.Role{ID: role.ID})
	}

	// Insert statuses (ignore duplicates)
	for _, status := range statuses {
		db.FirstOrCreate(&status, models.OracleNodeStatus{ID: status.ID})
	}
}
