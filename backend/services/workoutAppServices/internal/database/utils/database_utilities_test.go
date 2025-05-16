package database_utilities

import (
	"context"
	"testing"
	database "workout_app_backend/internal/database/init"
	utils "workout_app_backend/internal/utils"
)

func TestTableExists(t *testing.T) {

	utils.LoadEnv()

	// Get database instance
	db, err := database.GetInstance()
	if err != nil {
		t.Fatalf("Failed to get database instance: %v", err)
	}
	defer db.Close()

	ctx := context.Background()

	// Test non-existent table
	exists, err := TableExists(db, ctx, "non_existent_table")
	if err != nil {
		t.Fatalf("Failed to check non-existent table: %v", err)
	}
	if exists {
		t.Error("Expected non-existent table to return false")
	}

	// Create a test table
	testTable := "test_table_exists"
	schema := "id SERIAL PRIMARY KEY, name TEXT"
	err = CreateTable(db, ctx, testTable, schema)
	if err != nil {
		t.Fatalf("Failed to create test table: %v", err)
	}

	// Test existing table
	exists, err = TableExists(db, ctx, testTable)
	if err != nil {
		t.Fatalf("Failed to check existing table: %v", err)
	}
	if !exists {
		t.Error("Expected existing table to return true")
	}

	// Clean up
	_, err = db.ExecContext(ctx, "DROP TABLE IF EXISTS "+testTable)
	if err != nil {
		t.Fatalf("Failed to clean up test table: %v", err)
	}
}

func TestCreateTable(t *testing.T) {

	utils.LoadEnv()

	// Get database instance
	db, err := database.GetInstance()
	if err != nil {
		t.Fatalf("Failed to get database instance: %v", err)
	}
	defer db.Close()

	ctx := context.Background()

	// Test table creation
	testTable := "test_create_table"
	schema := "id SERIAL PRIMARY KEY, name TEXT, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP"

	// Create table
	err = CreateTable(db, ctx, testTable, schema)
	if err != nil {
		t.Fatalf("Failed to create table: %v", err)
	}

	// Verify table exists
	exists, err := TableExists(db, ctx, testTable)
	if err != nil {
		t.Fatalf("Failed to verify table existence: %v", err)
	}
	if !exists {
		t.Error("Created table does not exist")
	}

	// Test creating table again (should not error)
	err = CreateTable(db, ctx, testTable, schema)
	if err != nil {
		t.Fatalf("Failed to create existing table: %v", err)
	}

	// Clean up
	_, err = db.ExecContext(ctx, "DROP TABLE IF EXISTS "+testTable)
	if err != nil {
		t.Fatalf("Failed to clean up test table: %v", err)
	}
}
