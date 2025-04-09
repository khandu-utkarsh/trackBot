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
	db database.Database
}

// GetWorkoutModelInstance creates a new WorkoutModel instance
func GetWorkoutModelInstance(db database.Database) *WorkoutModel {
	return &WorkoutModel{db: db}
}

// Initialize creates the workouts table if it doesn't exist
func (m *WorkoutModel) Initialize(ctx context.Context) error {
	schema := `
		id SERIAL PRIMARY KEY,
		user_id BIGINT NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	`

	return m.db.CreateTable(ctx, "workouts", schema)
}

// Create creates a new workout
func (m *WorkoutModel) Create(ctx context.Context, workout *Workout) error {
	query := `
		INSERT INTO workouts (user_id, created_at, updated_at)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`

	err := m.db.QueryRowContext(
		ctx,
		query,
		workout.UserID,
		workout.CreatedAt,
		workout.UpdatedAt,
	).Scan(&workout.ID, &workout.CreatedAt, &workout.UpdatedAt)

	if err != nil {
		return fmt.Errorf("error creating workout: %v", err)
	}

	return nil
}

// Get retrieves a workout by ID
func (m *WorkoutModel) Get(ctx context.Context, id int64) (*Workout, error) {
	query := `
		SELECT id, user_id, name, description, date, duration, created_at, updated_at
		FROM workouts
		WHERE id = $1
	`

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
	query := `
		SELECT id, user_id, created_at, updated_at
		FROM workouts
		WHERE user_id = $1
		ORDER BY date DESC
	`

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
	_, err := m.db.ExecContext(ctx, "UPDATE workouts SET user_id = $1, updated_at = $2 WHERE id = $3", workout.UserID, time.Now(), workout.ID)
	return err
}

func (m *WorkoutModel) Delete(ctx context.Context, id int64) error {
	_, err := m.db.ExecContext(ctx, "DELETE FROM workouts WHERE id = $1", id)
	return err
}
