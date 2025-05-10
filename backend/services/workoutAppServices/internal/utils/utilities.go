package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
)

func IsTestEnv() bool {
	// Check if we're running under go test
	for _, arg := range os.Args {
		// Check for test.run argument
		if strings.HasPrefix(arg, "-test.run") {
			return true
		}
	}

	// Check binary name
	return strings.HasSuffix(os.Args[0], ".test") || // Check if binary is a test binary
		strings.Contains(os.Args[0], "/_test/") || // Check if running in test directory
		strings.Contains(os.Args[0], "\\_test\\") // Windows path
}

func LoadEnv() error {
	// Get the current file's directory
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("unable to get current file path")
	}

	// Try multiple possible .env file locations
	possiblePaths := []string{
		filepath.Join(filepath.Dir(file), "..", "..", ".env"),             // backend/services/workoutAppServices/.env
		filepath.Join(filepath.Dir(file), "..", "..", "..", ".env"),       // backend/services/.env
		filepath.Join(filepath.Dir(file), "..", "..", "..", "..", ".env"), // backend/.env
		".env", // Current directory
	}

	var envLoaded bool
	for _, envPath := range possiblePaths {
		if err := godotenv.Load(envPath); err == nil {
			log.Printf(".env file loaded from: %s\n", envPath)
			envLoaded = true
			break
		}
	}

	if !envLoaded {
		if IsTestEnv() {
			// In test environment, set default test values
			setTestEnvDefaults()
			log.Println("Using test environment defaults")
		} else {
			log.Println("No .env file found in any of the expected locations")
		}
	}

	// Validate required environment variables
	requiredVars := []string{
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
	}

	missingVars := []string{}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			missingVars = append(missingVars, v)
		}
	}

	if len(missingVars) > 0 {
		return fmt.Errorf("missing required environment variables: %v", missingVars)
	}

	return nil
}

func setTestEnvDefaults() {
	defaults := map[string]string{
		"DB_HOST":     "localhost",
		"DB_PORT":     "5432",
		"DB_USER":     "postgres",
		"DB_PASSWORD": "postgres",
		"DB_NAME":     "workout_app_test",
	}

	for k, v := range defaults {
		if os.Getenv(k) == "" {
			os.Setenv(k, v)
		}
	}
}
