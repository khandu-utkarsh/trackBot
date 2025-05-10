package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq" // PostgreSQL driver -- needed to execute anything on postgres RDMBS
)

var (
	instance *sql.DB
	mu       sync.Mutex
)

// GetInstance returns the database instance, reconnecting if necessary
func GetInstance() (*sql.DB, error) {
	mu.Lock()
	defer mu.Unlock()

	// If instance is nil or connection is closed, initialize a new connection
	if instance == nil {
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

func initDB() (*sql.DB, error) {

	//!Thee

	// Get database configuration from config package
	dbConfig, err := GetDBConfig()
	if err != nil {
		return nil, fmt.Errorf("error getting database config: %v", err)
	}

	// Open database connection using config

	connectionString := dbConfig.GetConnectionString()
	db, err := sql.Open("postgres", connectionString)
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
	return db, nil
}
