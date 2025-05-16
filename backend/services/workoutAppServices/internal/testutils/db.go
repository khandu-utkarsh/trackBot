package testutils

import (
	"database/sql"
	"testing"
	database "workout_app_backend/internal/database/init"
	utils "workout_app_backend/internal/utils"
)

// SetupTestDB creates a new test database connection and returns a cleanup function
func SetupTestDB(t *testing.T) (*sql.DB, func()) {

	utils.LoadEnv()

	//!This should return a new database instance for the test
	db, err := database.GetInstance()
	if err != nil {
		t.Fatalf("Failed to create test database: %v", err)
	}

	// Cleanup function to drop all tables after tests
	cleanup := func() {
		// First, drop all test tables explicitly
		testTables := []string{"users_test", "workouts_test", "exercises_test"}
		for _, table := range testTables {
			_, err := db.Exec("DROP TABLE IF EXISTS " + table + " CASCADE")
			if err != nil {
				t.Errorf("Failed to drop test table %s: %v", table, err)
			}
		}

		// Then drop any remaining tables in public schema
		_, err := db.Exec(`
			DO $$ DECLARE
				r RECORD;
			BEGIN
				FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP
					EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE';
				END LOOP;
			END $$;
		`)
		if err != nil {
			t.Errorf("Failed to drop remaining tables: %v", err)
		}

		// Close the database connection
		if err := db.Close(); err != nil {
			t.Errorf("Failed to close database connection: %v", err)
		}
	}

	return db, cleanup
}
