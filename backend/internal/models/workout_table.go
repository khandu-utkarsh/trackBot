package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"workout_app_backend/internal/database"
)

// Workout represents a workout session
type Workout struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// WorkoutModel handles workout-related database operations
type WorkoutModel struct {
	db         database.Database
	name       string
	foreignKey string
}

// GetWorkoutModelInstance creates a new WorkoutModel instance
func GetWorkoutModelInstance(db database.Database, name string, foreignKey string) *WorkoutModel {
	return &WorkoutModel{db: db, name: name, foreignKey: foreignKey}
}

// Initialize creates the workouts table if it doesn't exist
func (m *WorkoutModel) Initialize(ctx context.Context) error {
	schema := fmt.Sprintf(`
		id SERIAL PRIMARY KEY,
		user_id BIGINT NOT NULL REFERENCES %s(id) ON DELETE CASCADE,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	`, m.foreignKey)

	return m.db.CreateTable(ctx, m.name, schema)
}

// Create creates a new workout
func (m *WorkoutModel) Create(ctx context.Context, workout *Workout) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s (user_id, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id", m.name)

	var id int64
	err := m.db.QueryRowContext(
		ctx,
		query,
		workout.UserID,
		workout.CreatedAt,
		workout.UpdatedAt,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error creating workout: %v", err)
	}

	return id, nil
}

// Get retrieves a workout by ID
func (m *WorkoutModel) Get(ctx context.Context, id int64) (*Workout, error) {
	query := fmt.Sprintf(`
		SELECT id, user_id, created_at, updated_at
		FROM %s
		WHERE id = $1
	`, m.name)

	workout := &Workout{}
	err := m.db.QueryRowContext(ctx, query, id).Scan(
		&workout.ID,
		&workout.UserID,
		&workout.CreatedAt,
		&workout.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("workout not found")
	}
	if err != nil {
		return nil, fmt.Errorf("error getting workout: %v", err)
	}

	return workout, nil
}

// List retrieves all workouts for a user
func (m *WorkoutModel) List(ctx context.Context, userID int64) ([]*Workout, error) {
	query := fmt.Sprintf(`
		SELECT id, user_id, created_at, updated_at
		FROM %s
		WHERE user_id = $1
		ORDER BY created_at DESC
	`, m.name)

	rows, err := m.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("error listing workouts: %v", err)
	}
	defer rows.Close()

	var workouts []*Workout
	for rows.Next() {
		workout := &Workout{}
		err := rows.Scan(
			&workout.ID,
			&workout.UserID,
			&workout.CreatedAt,
			&workout.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning workout: %v", err)
		}
		workouts = append(workouts, workout)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating workouts: %v", err)
	}

	return workouts, nil
}

func (m *WorkoutModel) Update(ctx context.Context, workout *Workout) error {
	query := fmt.Sprintf(
		"UPDATE %s SET user_id = $1, created_at = $2, updated_at = $3 WHERE id = $4 RETURNING id",
		m.name,
	)

	var updatedID int64
	err := m.db.QueryRowContext(ctx, query, workout.UserID, workout.CreatedAt, workout.UpdatedAt, workout.ID).Scan(&updatedID)
	if err == sql.ErrNoRows {
		return fmt.Errorf("workout with ID %d not found", workout.ID)
	}
	if err != nil {
		return fmt.Errorf("error deleting workout: %v", err)
	}

	return nil
}

func (m *WorkoutModel) Delete(ctx context.Context, id int64) (int64, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 RETURNING id", m.name)

	var deletedID int64
	err := m.db.QueryRowContext(ctx, query, id).Scan(&deletedID)
	if err == sql.ErrNoRows {
		return -1, fmt.Errorf("workout with ID %d not found", id)
	}
	if err != nil {
		return 0, fmt.Errorf("error deleting workout: %v", err)
	}

	return deletedID, nil
}
