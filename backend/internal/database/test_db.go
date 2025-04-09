package database

import (
	"context"
	"fmt"
	"testing"
	"workout_app_backend/internal/utils"

	_ "github.com/lib/pq"
)

// NewTestDB creates a new test database connection
func NewTestDB() (Database, error) {

	utils.LoadEnv()

	//!Testing on AWS Cluster
	db, err := GetInstance()
	if err != nil {
		return nil, fmt.Errorf("error getting db config: %v", err)
	}
	// Verify the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to test database: %v", err)
	}

	return db, nil
}

// CleanupTestDB drops all tables in the test database
func CleanupTestDB(t *testing.T, db Database) {
	t.Helper()
	ctx := context.Background()

	// Drop all tables
	_, err := db.ExecContext(ctx, `
		DO $$ DECLARE
			r RECORD;
		BEGIN
			FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP
				EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE';
			END LOOP;
		END $$;
	`)
	if err != nil {
		t.Fatalf("Failed to cleanup test database: %v", err)
	}
}
