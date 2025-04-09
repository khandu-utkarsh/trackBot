package testutils

import (
	"testing"
	"workout_app_backend/internal/database"
)

// SetupTestDB creates a new test database connection and returns a cleanup function
func SetupTestDB(t *testing.T) (database.Database, func()) {
	db, err := database.NewTestDB()
	if err != nil {
		t.Fatalf("Failed to create test database: %v", err)
	}

	// Cleanup function to drop all tables after tests
	cleanup := func() {
		database.CleanupTestDB(t, db)
		db.Close()
	}

	return db, cleanup
}
