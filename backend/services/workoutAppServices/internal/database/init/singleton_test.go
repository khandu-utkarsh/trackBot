package database

import (
	"database/sql"
	"sync"
	"testing"
	utils "workout_app_backend/internal/utils"
)

func TestGetInstance(t *testing.T) {

	utils.LoadEnv()
	// Test first instance creation
	db1, err := GetInstance()
	if err != nil {
		t.Fatalf("Failed to get first instance: %v", err)
	}
	if db1 == nil {
		t.Fatal("First instance is nil")
	}

	// Test second instance (should be the same as first)
	db2, err := GetInstance()
	if err != nil {
		t.Fatalf("Failed to get second instance: %v", err)
	}
	if db2 != db1 {
		t.Error("Second instance is different from first instance")
	}

	// Test connection is alive
	if err := db1.Ping(); err != nil {
		t.Fatalf("Database connection is not alive: %v", err)
	}
}

func TestGetInstanceReconnection(t *testing.T) {

	utils.LoadEnv()
	// Get initial instance
	db1, err := GetInstance()
	if err != nil {
		t.Fatalf("Failed to get initial instance: %v", err)
	}

	// Close the connection to simulate a dead connection
	db1.Close()

	// Get new instance (should reconnect)
	db2, err := GetInstance()
	if err != nil {
		t.Fatalf("Failed to get new instance after connection death: %v", err)
	}
	if db2 == nil {
		t.Fatal("New instance is nil")
	}

	// Verify new connection is alive
	if err := db2.Ping(); err != nil {
		t.Fatalf("New database connection is not alive: %v", err)
	}
}

func TestGetInstanceConcurrent(t *testing.T) {

	utils.LoadEnv()
	var wg sync.WaitGroup

	// Create a set to store unique DB instances with mutex protection
	var mu sync.Mutex
	dbSet := make(map[*sql.DB]bool)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			db, err := GetInstance()
			if err != nil {
				t.Errorf("Failed to get instance in goroutine: %v", err)
			}
			if db == nil {
				t.Error("Instance is nil in goroutine")
			}

			mu.Lock()
			dbSet[db] = true
			mu.Unlock()
		}()
	}

	// Wait for all goroutines to finish before checking the set
	wg.Wait()

	// Check if we got exactly one unique DB instance
	mu.Lock()
	if len(dbSet) != 1 {
		t.Errorf("Expected 1 unique DB instance, got %d", len(dbSet))
	}
	mu.Unlock()
}
