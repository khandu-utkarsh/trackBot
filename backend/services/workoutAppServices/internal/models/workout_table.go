package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	database_utilities "workout_app_backend/internal/database/utils"
)

// Common errors
var (
	ErrWorkoutNotFound = errors.New("workout not found")
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
	db         *sql.DB
	name       string
	foreignKey string
}

// GetWorkoutModelInstance creates a new WorkoutModel instance
func GetWorkoutModelInstance(db *sql.DB, name string, foreignKey string) *WorkoutModel {
	return &WorkoutModel{db: db, name: name, foreignKey: foreignKey}
}

// Initialize creates the workouts table if it doesn't exist
func (m *WorkoutModel) Initialize(ctx context.Context) error {
	modelsLogger.Println("WorkoutModel: Initialize called") //! Logging the request.
	schema := fmt.Sprintf(`
		id SERIAL PRIMARY KEY,
		user_id BIGINT NOT NULL REFERENCES %s(id) ON DELETE CASCADE,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	`, m.foreignKey)

	return database_utilities.CreateTable(m.db, ctx, m.name, schema)
}

// validateWorkout checks if the workout data is valid
func (m *WorkoutModel) validateWorkout(workout *Workout) error {
	modelsLogger.Println("WorkoutModel: validateWorkout called") //! Logging the request.
	if workout == nil {
		return fmt.Errorf("%w: workout cannot be nil", ErrInvalidInput)
	}
	if workout.UserID <= 0 {
		return fmt.Errorf("%w: invalid user ID", ErrInvalidInput)
	}
	return nil
}

// scanWorkout scans a database row into a Workout struct
func (m *WorkoutModel) scanWorkout(row *sql.Row) (*Workout, error) {
	modelsLogger.Println("WorkoutModel: scanWorkout called") //! Logging the request.
	var workout Workout
	err := row.Scan(&workout.ID, &workout.UserID, &workout.CreatedAt, &workout.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, ErrWorkoutNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("error scanning workout: %w", err)
	}
	return &workout, nil
}

// Create creates a new workout
func (m *WorkoutModel) Create(ctx context.Context, workout *Workout) (int64, error) {
	modelsLogger.Println("WorkoutModel: Create called") //! Logging the request.
	if err := m.validateWorkout(workout); err != nil {
		return 0, err
	}

	now := time.Now()
	query := fmt.Sprintf("INSERT INTO %s (user_id, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id", m.name)

	var id int64
	err := m.db.QueryRowContext(ctx, query, workout.UserID, now, now).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error creating workout: %w", err)
	}

	return id, nil
}

// Get retrieves a workout by ID
func (m *WorkoutModel) Get(ctx context.Context, id int64) (*Workout, error) {
	modelsLogger.Println("WorkoutModel: Get called") //! Logging the request.
	if id <= 0 {
		return nil, fmt.Errorf("%w: invalid workout ID", ErrInvalidInput)
	}

	query := fmt.Sprintf(`
		SELECT id, user_id, created_at, updated_at
		FROM %s
		WHERE id = $1
	`, m.name)

	return m.scanWorkout(m.db.QueryRowContext(ctx, query, id))
}

type WorkoutListParams struct {
	UserID int64
	Year   string
	Month  string
	Day    string
}

// List retrieves all workouts for a user
// List retrieves all workouts for a user, with optional date filtering
func (m *WorkoutModel) List(ctx context.Context, params WorkoutListParams) ([]*Workout, error) {
	modelsLogger.Println("WorkoutModel: List called") //! Logging the request.

	if params.UserID <= 0 {
		return nil, fmt.Errorf("%w: invalid user ID", ErrInvalidInput)
	}

	var (
		query string
		args  []any
	)

	if params.Year != "" {
		year, err := strconv.Atoi(params.Year)
		if err != nil {
			return nil, fmt.Errorf("%w: invalid year", ErrInvalidInput)
		}

		if params.Month != "" {
			month, err := strconv.Atoi(params.Month)
			if err != nil {
				return nil, fmt.Errorf("%w: invalid month", ErrInvalidInput)
			}

			if params.Day != "" {
				day, err := strconv.Atoi(params.Day)
				if err != nil {
					return nil, fmt.Errorf("%w: invalid day", ErrInvalidInput)
				}
				// Specific day filter
				dateFilter := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
				query = fmt.Sprintf(`
					SELECT id, user_id, created_at, updated_at
					FROM %s
					WHERE user_id = $1 AND DATE(created_at) = $2
					ORDER BY created_at DESC
				`, m.name)
				args = []any{params.UserID, dateFilter.Format("2006-01-02")}

			} else {
				// Month range filter
				start := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
				end := start.AddDate(0, 1, 0)
				query = fmt.Sprintf(`
					SELECT id, user_id, created_at, updated_at
					FROM %s
					WHERE user_id = $1 AND created_at >= $2 AND created_at < $3
					ORDER BY created_at DESC
				`, m.name)
				args = []any{params.UserID, start, end}
			}
		} else {
			// Year range filter
			start := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
			end := start.AddDate(1, 0, 0)
			query = fmt.Sprintf(`
				SELECT id, user_id, created_at, updated_at
				FROM %s
				WHERE user_id = $1 AND created_at >= $2 AND created_at < $3
				ORDER BY created_at DESC
			`, m.name)
			args = []any{params.UserID, start, end}
		}
	} else {
		// No filters
		query = fmt.Sprintf(`
			SELECT id, user_id, created_at, updated_at
			FROM %s
			WHERE user_id = $1
			ORDER BY created_at DESC
		`, m.name)
		args = []any{params.UserID}
	}

	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error querying workouts: %w", err)
	}
	defer rows.Close()

	var workouts []*Workout
	for rows.Next() {
		var workout Workout
		if err := rows.Scan(&workout.ID, &workout.UserID, &workout.CreatedAt, &workout.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning workout row: %w", err)
		}
		workouts = append(workouts, &workout)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating workout rows: %w", err)
	}

	return workouts, nil
}

// Helper function with defaults
func (m *WorkoutModel) ListWithDefaults(ctx context.Context, userID int64) ([]*Workout, error) {
	modelsLogger.Println("WorkoutModel: ListWithDefaults called") //! Logging the request.
	return m.List(ctx, WorkoutListParams{
		UserID: userID,
		Year:   "",
		Month:  "",
		Day:    "",
	})
}

// Update updates an existing workout
func (m *WorkoutModel) Update(ctx context.Context, workout *Workout) error {
	modelsLogger.Println("WorkoutModel: Update called") //! Logging the request.
	if err := m.validateWorkout(workout); err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET user_id = $1, updated_at = $2 WHERE id = $3", m.name)
	result, err := m.db.ExecContext(ctx, query, workout.UserID, time.Now(), workout.ID)
	if err != nil {
		return fmt.Errorf("error updating workout: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return ErrWorkoutNotFound
	}

	return nil
}

// Delete removes a workout from the database
func (m *WorkoutModel) Delete(ctx context.Context, id int64) error {
	modelsLogger.Println("WorkoutModel: Delete called") //! Logging the request.
	if id <= 0 {
		return fmt.Errorf("%w: invalid workout ID", ErrInvalidInput)
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", m.name)
	result, err := m.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting workout: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return ErrWorkoutNotFound
	}

	return nil
}
