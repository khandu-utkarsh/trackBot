package database_utilities

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

// TableExists checks if a table exists in the database
func TableExists(db *sql.DB, ctx context.Context, tableName string) (bool, error) {
	// Using to_regclass is more efficient than querying information_schema
	query := "SELECT to_regclass($1) IS NOT NULL"

	var exists bool
	err := db.QueryRowContext(ctx, query, tableName).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking if table %s exists: %w", tableName, err)
	}

	return exists, nil
}

// CreateTable creates a table with the given schema if it doesn't exist
func CreateTable(db *sql.DB, ctx context.Context, tableName string, schema string) error {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", tableName, schema)
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("error creating table %s: %w", tableName, err)
	}

	log.Printf("Table %s created or already exists", tableName)
	return nil
}
