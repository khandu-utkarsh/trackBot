package database_test

import (
	"testing"

	"workoutapp/test"
)

func TestDatabaseConnection(t *testing.T) {
	// Test case 1: Valid database connection
	db, ctx, cleanup := test.SetupTestDB(t)
	defer cleanup()

	// Verify connection is alive
	err := db.PingContext(ctx)
	if err != nil {
		t.Errorf("Database connection failed: %v", err)
	}

	// Test case 2: Database operations work
	var result int
	err = db.QueryRowContext(ctx, "SELECT 1").Scan(&result)
	if err != nil {
		t.Errorf("Failed to execute simple query: %v", err)
	}
	if result != 1 {
		t.Errorf("Expected query result 1, got %d", result)
	}
}
