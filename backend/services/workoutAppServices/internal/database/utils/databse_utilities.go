package database_utilities

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
)

var databaseUtilitiesLogger *log.Logger

func init() {
	databaseUtilitiesLogger = log.New(os.Stdout, "Database Utilities: ", log.LstdFlags)
}

// TableExists checks if a table exists in the database
func TableExists(db *sql.DB, ctx context.Context, tableName string) (bool, error) {
	databaseUtilitiesLogger.Println("Checking if table exists: ", tableName)
	// Using to_regclass is more efficient than querying information_schema
	query := "SELECT to_regclass($1) IS NOT NULL"

	var exists bool
	err := db.QueryRowContext(ctx, query, tableName).Scan(&exists)
	if err != nil {
		databaseUtilitiesLogger.Println("Error checking if table exists: ", err) //! Logging the error.
		return false, fmt.Errorf("error checking if table %s exists: %w", tableName, err)
	}

	return exists, nil
}

// CreateTable creates a table with the given schema if it doesn't exist
func CreateTable(db *sql.DB, ctx context.Context, tableName string, schema string) error {
	databaseUtilitiesLogger.Println("Creating table: ", tableName)
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", tableName, schema)
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		databaseUtilitiesLogger.Println("Error creating table: ", err) //! Logging the error.
		return fmt.Errorf("error creating table %s: %w", tableName, err)
	}
	databaseUtilitiesLogger.Println("Table ", tableName, " created or already exists")
	return nil
}
