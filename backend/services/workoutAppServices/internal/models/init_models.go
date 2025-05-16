package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	database_utilities "workout_app_backend/internal/database/utils"
)

var (
	ErrMigrationFailed = errors.New("migration failed")
)

// Migration represents a database migration
type Migration struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	AppliedAt time.Time `json:"applied_at"`
}

// MigrationModel handles migration-related database operations
type MigrationModel struct {
	db *sql.DB
}

// GetMigrationModelInstance creates a new MigrationModel instance
func GetMigrationModelInstance(db *sql.DB) *MigrationModel {
	return &MigrationModel{db: db}
}

// Initialize creates the migrations table if it doesn't exist
func (m *MigrationModel) Initialize(ctx context.Context) error {
	schema := `
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL UNIQUE,
		applied_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	`
	return database_utilities.CreateTable(m.db, ctx, "migrations", schema)
}

// HasMigration checks if a migration has been applied
func (m *MigrationModel) HasMigration(ctx context.Context, name string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM migrations WHERE name = $1)"
	err := m.db.QueryRowContext(ctx, query, name).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking migration: %w", err)
	}
	return exists, nil
}

// ApplyMigration records a migration as applied
func (m *MigrationModel) ApplyMigration(ctx context.Context, name string) error {
	query := "INSERT INTO migrations (name) VALUES ($1)"
	_, err := m.db.ExecContext(ctx, query, name)
	if err != nil {
		return fmt.Errorf("error applying migration: %w", err)
	}
	return nil
}

// InitializeModels initializes all database tables using migrations
func InitializeModels(db *sql.DB) error {
	ctx := context.Background()

	// Start a transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer tx.Rollback()

	// Initialize migrations table first
	migrationModel := GetMigrationModelInstance(db)
	if err := migrationModel.Initialize(ctx); err != nil {
		return fmt.Errorf("error initializing migrations table: %w", err)
	}

	// Define migrations
	migrations := []struct {
		name string
		fn   func(context.Context, *sql.Tx) error
	}{
		{
			name: "create_users_table",
			fn: func(ctx context.Context, tx *sql.Tx) error {
				userModel := GetUserModelInstance(db, "users")
				return userModel.Initialize(ctx)
			},
		},
		{
			name: "create_workouts_table",
			fn: func(ctx context.Context, tx *sql.Tx) error {
				workoutModel := GetWorkoutModelInstance(db, "workouts", "users")
				return workoutModel.Initialize(ctx)
			},
		},
		{
			name: "create_exercises_table",
			fn: func(ctx context.Context, tx *sql.Tx) error {
				exerciseModel := GetExerciseModelInstance(db, "exercises", "workouts")
				return exerciseModel.Initialize(ctx)
			},
		},
	}

	// Apply each migration if not already applied
	for _, migration := range migrations {
		exists, err := migrationModel.HasMigration(ctx, migration.name)
		if err != nil {
			return fmt.Errorf("error checking migration %s: %w", migration.name, err)
		}

		if !exists {
			log.Printf("Applying migration: %s", migration.name)
			if err := migration.fn(ctx, tx); err != nil {
				return fmt.Errorf("error applying migration %s: %w", migration.name, err)
			}

			if err := migrationModel.ApplyMigration(ctx, migration.name); err != nil {
				return fmt.Errorf("error recording migration %s: %w", migration.name, err)
			}
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing migrations: %w", err)
	}

	log.Println("All migrations completed successfully")
	return nil
}
