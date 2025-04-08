package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"workout_app_backend/internal/config"

	_ "github.com/lib/pq"
)

// Database represents a generic database interface
type Database interface {
	// Basic operations
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	Ping() error
	Close() error

	// Transaction support
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)

	// Table operations
	TableExists(ctx context.Context, tableName string) (bool, error)
	CreateTable(ctx context.Context, tableName string, schema string) error
}

// SQLDatabase implements the Database interface using SQL
type SQLDatabase struct {
	*sql.DB
}

var (
	instance *SQLDatabase
	mu       sync.Mutex
)

// GetInstance returns the database instance, reconnecting if necessary
func GetInstance() (Database, error) {
	mu.Lock()
	defer mu.Unlock()

	// If instance is nil or connection is closed, initialize a new connection
	if instance == nil || instance.DB == nil {
		var err error
		instance, err = initDB()
		if err != nil {
			return nil, fmt.Errorf("failed to initialize database: %v", err)
		}
		return instance, nil
	}

	// Test if the connection is still alive
	if err := instance.Ping(); err != nil {
		// Connection is dead, close it and create a new one
		instance.Close()
		var err error
		instance, err = initDB()
		if err != nil {
			return nil, fmt.Errorf("failed to reinitialize database: %v", err)
		}
	}

	return instance, nil
}

func initDB() (*SQLDatabase, error) {
	// Get database configuration from config package
	dbConfig, err := config.GetDBConfig()
	if err != nil {
		return nil, fmt.Errorf("error getting database config: %v", err)
	}

	// Open database connection using config
	db, err := sql.Open("postgres", dbConfig.GetConnectionString())
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(dbConfig.MaxOpenConns)
	db.SetMaxIdleConns(dbConfig.MaxIdleConns)
	db.SetConnMaxLifetime(dbConfig.ConnMaxLifetime)

	// Test the connection
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	log.Println("Successfully connected to database")
	return &SQLDatabase{db}, nil
}

// TableExists checks if a table exists in the database
func (db *SQLDatabase) TableExists(ctx context.Context, tableName string) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT FROM information_schema.tables 
			WHERE table_schema = 'public' 
			AND table_name = $1
		)
	`

	var exists bool
	err := db.QueryRowContext(ctx, query, tableName).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking if table %s exists: %v", tableName, err)
	}

	return exists, nil
}

// CreateTable creates a table with the given schema if it doesn't exist
func (db *SQLDatabase) CreateTable(ctx context.Context, tableName string, schema string) error {
	exists, err := db.TableExists(ctx, tableName)
	if err != nil {
		return err
	}

	if exists {
		log.Printf("Table %s already exists", tableName)
		return nil
	}

	query := fmt.Sprintf("CREATE TABLE %s (%s)", tableName, schema)
	_, err = db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("error creating table %s: %v", tableName, err)
	}

	log.Printf("Table %s created successfully", tableName)
	return nil
}

// Close closes the database connection
func (db *SQLDatabase) Close() error {
	if db.DB != nil {
		return db.DB.Close()
	}
	return nil
}
