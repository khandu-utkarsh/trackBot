package database

import (
	"os"
	"testing"
	"time"
)

func TestGetLocalDBConfig(t *testing.T) {
	// Set up test environment variables
	os.Setenv("DB_USER_LOCAL", "testuser")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "testdb")
	os.Setenv("TEST_ENV", "true")

	config, err := GetLocalDBConfig()
	if err != nil {
		t.Fatalf("Failed to get local DB config: %v", err)
	}

	// Verify config values
	if config.Username != "testuser" {
		t.Errorf("Expected username 'testuser', got '%s'", config.Username)
	}
	if config.Host != "localhost" {
		t.Errorf("Expected host 'localhost', got '%s'", config.Host)
	}
	if config.Port != "5432" {
		t.Errorf("Expected port '5432', got '%s'", config.Port)
	}
	if config.DBName != "testdb" {
		t.Errorf("Expected DB name 'testdb', got '%s'", config.DBName)
	}
	if config.MaxOpenConns != 25 {
		t.Errorf("Expected MaxOpenConns 25, got %d", config.MaxOpenConns)
	}
	if config.MaxIdleConns != 5 {
		t.Errorf("Expected MaxIdleConns 5, got %d", config.MaxIdleConns)
	}
	if config.ConnMaxLifetime != 5*time.Minute {
		t.Errorf("Expected ConnMaxLifetime 5m, got %v", config.ConnMaxLifetime)
	}
	if config.DatabaseType != "local-postgres" {
		t.Errorf("Expected DatabaseType 'local-postgres', got '%s'", config.DatabaseType)
	}

	// Clean up
	os.Unsetenv("DB_USER_LOCAL")
	os.Unsetenv("DB_PORT")
	os.Unsetenv("DB_NAME")
	os.Unsetenv("TEST_ENV")
}

func TestGetLocalDBConfigConnectionString(t *testing.T) {
	// Test local connection string
	config := &DBConfig{
		Host:     "localhost",
		Port:     "5432",
		Username: "testuser",
		DBName:   "testdb",
	}
	connStr := config.GetConnectionString()
	expected := "host=localhost port=5432 user=testuser dbname=testdb sslmode=disable"
	if connStr != expected {
		t.Errorf("Expected connection string '%s', got '%s'", expected, connStr)
	}
}

func TestGetAWSDBConfigConnectionString(t *testing.T) {
	// Test local connection string
	config := &DBConfig{
		Host:     "localhost",
		Port:     "5432",
		Username: "testuser",
		DBName:   "testdb",
		Password: "testpass",
	}
	connStr := config.GetAWSDBConfigConnectionString()
	expected := "host=localhost port=5432 user=testuser password=testpass dbname=testdb sslmode=disable"
	if connStr != expected {
		t.Errorf("Expected connection string '%s', got '%s'", expected, connStr)
	}
}

func TestGetDBConfig(t *testing.T) {

	// Test local config
	config, err := GetDBConfig()
	if err != nil {
		t.Fatalf("Failed to get DB config in test environment: %v", err)
	}
	if config.DatabaseType != "local-postgres" {
		t.Errorf("Expected local-postgres config, got %s", config.DatabaseType)
	}

	// Note: AWS config testing would require mocking AWS services
	// This is typically done in integration tests
}
