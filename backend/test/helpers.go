package test

import (
	"context"
	"os"
	"testing"
	"time"

	"workoutapp/internal/database"

	"github.com/joho/godotenv"
)

// TestDBConfig holds test database configuration
type TestDBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// GetTestDBConfig returns test database configuration
func GetTestDBConfig() *TestDBConfig {
	// Load environment variables from .env.test if it exists
	_ = godotenv.Load(".env.test")

	return &TestDBConfig{
		Host:     getEnvOrDefault("DB_HOST", "localhost"),
		Port:     getEnvOrDefault("DB_PORT", "5432"),
		User:     getEnvOrDefault("DB_USER", "postgres"),
		Password: getEnvOrDefault("DB_PASSWORD", "postgres"),
		DBName:   getEnvOrDefault("DB_NAME", "workout_test"),
	}
}

// SetupTestDB initializes a test database connection
func SetupTestDB(t *testing.T) (*database.Database, context.Context, func()) {
	// Create a new context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// Initialize database
	db, err := database.GetInstance()
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}

	// Initialize tables
	if err := db.InitializeTables(ctx); err != nil {
		t.Fatalf("Failed to initialize test tables: %v", err)
	}

	// Return cleanup function
	cleanup := func() {
		cancel()
		db.Close()
	}

	return db, ctx, cleanup
}

// getEnvOrDefault returns the value of the environment variable or a default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
