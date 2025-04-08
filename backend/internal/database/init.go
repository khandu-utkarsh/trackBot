package database

import (
	"context"
	"fmt"
	"log"
)

// TableExists checks if a table exists in the database
func (db *Database) TableExists(ctx context.Context, tableName string) (bool, error) {
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

// CreateWorkoutTable creates the workouts table if it doesn't exist
func (db *Database) CreateWorkoutTable(ctx context.Context) error {
	exists, err := db.TableExists(ctx, "workouts")
	if err != nil {
		return err
	}

	if exists {
		log.Println("Workouts table already exists")
		return nil
	}

	query := `
		CREATE TABLE workouts (
			id SERIAL PRIMARY KEY,
			user_id BIGINT NOT NULL,
			name VARCHAR(255) NOT NULL,
			description TEXT,
			date TIMESTAMP WITH TIME ZONE NOT NULL,
			duration INTERVAL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)
	`

	_, err = db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("error creating workouts table: %v", err)
	}

	log.Println("Workouts table created successfully")
	return nil
}

// CreateExerciseTable creates the exercises table if it doesn't exist
func (db *Database) CreateExerciseTable(ctx context.Context) error {
	exists, err := db.TableExists(ctx, "exercises")
	if err != nil {
		return err
	}

	if exists {
		log.Println("Exercises table already exists")
		return nil
	}

	query := `
		CREATE TABLE exercises (
			id SERIAL PRIMARY KEY,
			workout_id INTEGER NOT NULL REFERENCES workouts(id) ON DELETE CASCADE,
			name VARCHAR(255) NOT NULL,
			description TEXT,
			sets INTEGER,
			reps INTEGER,
			weight FLOAT,
			duration INTERVAL,
			distance FLOAT,
			order_index INTEGER NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)
	`

	_, err = db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("error creating exercises table: %v", err)
	}

	log.Println("Exercises table created successfully")
	return nil
}

// InitializeTables creates all necessary tables if they don't exist
func (db *Database) InitializeTables(ctx context.Context) error {
	if err := db.CreateWorkoutTable(ctx); err != nil {
		return fmt.Errorf("error initializing workout table: %v", err)
	}

	if err := db.CreateExerciseTable(ctx); err != nil {
		return fmt.Errorf("error initializing exercise table: %v", err)
	}

	return nil
}
