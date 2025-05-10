package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	database_utilities "workout_app_backend/services/workoutAppServices/internal/database/utils"
)

// Common errors
var (
	ErrExerciseNotFound = errors.New("exercise not found")
)

// ExerciseType represents the type of exercise
type ExerciseType string

const (
	ExerciseTypeCardio  ExerciseType = "cardio"
	ExerciseTypeWeights ExerciseType = "weights"
)

// BaseExercise represents the common fields for all exercise types
type BaseExercise struct {
	ID        int64        `json:"id"`
	WorkoutID int64        `json:"workout_id"`
	Name      string       `json:"name"`
	Type      ExerciseType `json:"type"`
	Notes     string       `json:"notes"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

// CardioExercise represents a cardio exercise with specific attributes
type CardioExercise struct {
	BaseExercise
	Distance float64 `json:"distance"` // in meters
	Duration int     `json:"duration"` // in seconds
}

// WeightExercise represents a weight training exercise with specific attributes
type WeightExercise struct {
	BaseExercise
	Sets   int     `json:"sets"`
	Reps   int     `json:"reps"`
	Weight float64 `json:"weight"` // in kilograms
}

// ExerciseModel handles exercise-related database operations
type ExerciseModel struct {
	db         *sql.DB
	name       string
	foreignKey string
}

// GetExerciseModelInstance creates a new ExerciseModel instance
func GetExerciseModelInstance(db *sql.DB, name string, foreignKey string) *ExerciseModel {
	return &ExerciseModel{db: db, name: name, foreignKey: foreignKey}
}

// Initialize creates the exercises table if it doesn't exist
func (m *ExerciseModel) Initialize(ctx context.Context) error {
	schema := fmt.Sprintf(`
		id SERIAL PRIMARY KEY,
		workout_id INTEGER NOT NULL REFERENCES %s(id) ON DELETE CASCADE,
		name VARCHAR(255) NOT NULL,
		Type VARCHAR(20) NOT NULL CHECK (Type IN ('cardio', 'weights')),
		Notes TEXT,
		
		-- Cardio specific fields
		distance FLOAT,
		duration INTEGER,
		
		-- Weight training specific fields
		sets INTEGER,
		reps INTEGER,
		weight FLOAT,
		
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	`, m.foreignKey)

	return database_utilities.CreateTable(m.db, ctx, m.name, schema)
}

// validateBaseExercise checks if the base exercise data is valid
func (m *ExerciseModel) validateBaseExercise(exercise *BaseExercise) error {
	if exercise == nil {
		return fmt.Errorf("%w: exercise cannot be nil", ErrInvalidInput)
	}
	if exercise.WorkoutID <= 0 {
		return fmt.Errorf("%w: invalid workout ID", ErrInvalidInput)
	}
	if exercise.Name == "" {
		return fmt.Errorf("%w: exercise name cannot be empty", ErrInvalidInput)
	}
	return nil
}

// validateCardioExercise checks if the cardio exercise data is valid
func (m *ExerciseModel) validateCardioExercise(exercise *CardioExercise) error {

	if exercise == nil {
		return fmt.Errorf("%w: exercise cannot be nil", ErrInvalidInput)
	}

	if err := m.validateBaseExercise(&exercise.BaseExercise); err != nil {
		return err
	}
	if exercise.Type != ExerciseTypeCardio {
		return fmt.Errorf("%w: invalid exercise type for cardio exercise", ErrInvalidInput)
	}
	if exercise.Distance < 0 {
		return fmt.Errorf("%w: distance cannot be negative", ErrInvalidInput)
	}
	if exercise.Duration <= 0 {
		return fmt.Errorf("%w: duration must be positive", ErrInvalidInput)
	}
	return nil
}

// validateWeightExercise checks if the weight exercise data is valid
func (m *ExerciseModel) validateWeightExercise(exercise *WeightExercise) error {

	if exercise == nil {
		return fmt.Errorf("%w: exercise cannot be nil", ErrInvalidInput)
	}

	if err := m.validateBaseExercise(&exercise.BaseExercise); err != nil {
		return err
	}
	if exercise.Type != ExerciseTypeWeights {
		return fmt.Errorf("%w: invalid exercise type for weight exercise", ErrInvalidInput)
	}
	if exercise.Sets <= 0 {
		return fmt.Errorf("%w: sets must be positive", ErrInvalidInput)
	}
	if exercise.Reps <= 0 {
		return fmt.Errorf("%w: reps must be positive", ErrInvalidInput)
	}
	if exercise.Weight < 0 {
		return fmt.Errorf("%w: weight cannot be negative", ErrInvalidInput)
	}
	return nil
}

// scanExercise scans a database row into the appropriate exercise type
func (m *ExerciseModel) scanExercise(row *sql.Row) (interface{}, error) {
	var (
		base         BaseExercise
		exerciseType ExerciseType
		distance     sql.NullFloat64
		duration     sql.NullInt64
		sets         sql.NullInt64
		reps         sql.NullInt64
		weight       sql.NullFloat64
	)

	err := row.Scan(
		&base.ID,
		&base.WorkoutID,
		&base.Name,
		&exerciseType,
		&base.Notes,
		&distance,
		&duration,
		&sets,
		&reps,
		&weight,
		&base.CreatedAt,
		&base.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, ErrExerciseNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("error scanning exercise: %w", err)
	}

	base.Type = exerciseType

	switch exerciseType {
	case ExerciseTypeCardio:
		return &CardioExercise{
			BaseExercise: base,
			Distance:     distance.Float64,
			Duration:     int(duration.Int64),
		}, nil

	case ExerciseTypeWeights:
		return &WeightExercise{
			BaseExercise: base,
			Sets:         int(sets.Int64),
			Reps:         int(reps.Int64),
			Weight:       weight.Float64,
		}, nil

	default:
		return nil, fmt.Errorf("unknown exercise type: %s", exerciseType)
	}
}

// CreateCardio creates a new cardio exercise
func (m *ExerciseModel) CreateCardio(ctx context.Context, exercise *CardioExercise) (int64, error) {
	if err := m.validateCardioExercise(exercise); err != nil {
		return 0, err
	}

	now := time.Now()
	query := fmt.Sprintf("INSERT INTO %s (workout_id, name, Type, Notes, distance, duration, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", m.name)

	var id int64
	err := m.db.QueryRowContext(
		ctx,
		query,
		exercise.WorkoutID,
		exercise.Name,
		ExerciseTypeCardio,
		exercise.Notes,
		exercise.Distance,
		exercise.Duration,
		now,
		now,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error creating cardio exercise: %w", err)
	}

	return id, nil
}

// CreateWeights creates a new weight training exercise
func (m *ExerciseModel) CreateWeights(ctx context.Context, exercise *WeightExercise) (int64, error) {
	if err := m.validateWeightExercise(exercise); err != nil {
		return 0, err
	}

	now := time.Now()
	query := fmt.Sprintf("INSERT INTO %s (workout_id, name, Type, Notes, sets, reps, weight, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id", m.name)

	var id int64
	err := m.db.QueryRowContext(
		ctx,
		query,
		exercise.WorkoutID,
		exercise.Name,
		ExerciseTypeWeights,
		exercise.Notes,
		exercise.Sets,
		exercise.Reps,
		exercise.Weight,
		now,
		now,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error creating weight exercise: %w", err)
	}

	return id, nil
}

// Get retrieves an exercise by ID and returns the correct struct type
func (m *ExerciseModel) Get(ctx context.Context, id int64) (interface{}, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: invalid exercise ID", ErrInvalidInput)
	}

	query := fmt.Sprintf(`
		SELECT id, workout_id, name, Type, Notes,
		       distance, duration,
		       sets, reps, weight,
		       created_at, updated_at
		FROM %s
		WHERE id = $1
	`, m.name)

	return m.scanExercise(m.db.QueryRowContext(ctx, query, id))
}

// ListByWorkout retrieves all exercises for a workout
func (m *ExerciseModel) ListByWorkout(ctx context.Context, workoutID int64) ([]interface{}, error) {
	if workoutID <= 0 {
		return nil, fmt.Errorf("%w: invalid workout ID", ErrInvalidInput)
	}

	query := fmt.Sprintf(`
		SELECT id, workout_id, name, Type, Notes,
		       distance, duration,
		       sets, reps, weight,
		       created_at, updated_at
		FROM %s
		WHERE workout_id = $1
		ORDER BY id
	`, m.name)

	rows, err := m.db.QueryContext(ctx, query, workoutID)
	if err != nil {
		return nil, fmt.Errorf("error querying exercises: %w", err)
	}
	defer rows.Close()

	exercises := make([]interface{}, 0)
	for rows.Next() {
		var (
			base         BaseExercise
			exerciseType ExerciseType
			distance     sql.NullFloat64
			duration     sql.NullInt64
			sets         sql.NullInt64
			reps         sql.NullInt64
			weight       sql.NullFloat64
		)

		if err := rows.Scan(
			&base.ID,
			&base.WorkoutID,
			&base.Name,
			&exerciseType,
			&base.Notes,
			&distance,
			&duration,
			&sets,
			&reps,
			&weight,
			&base.CreatedAt,
			&base.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("error scanning exercise row: %w", err)
		}

		base.Type = exerciseType

		switch exerciseType {
		case ExerciseTypeCardio:
			exercises = append(exercises, &CardioExercise{
				BaseExercise: base,
				Distance:     distance.Float64,
				Duration:     int(duration.Int64),
			})
		case ExerciseTypeWeights:
			exercises = append(exercises, &WeightExercise{
				BaseExercise: base,
				Sets:         int(sets.Int64),
				Reps:         int(reps.Int64),
				Weight:       weight.Float64,
			})
		default:
			return nil, fmt.Errorf("unknown exercise type: %s", exerciseType)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating exercise rows: %w", err)
	}

	return exercises, nil
}

// Update updates an existing exercise
func (m *ExerciseModel) Update(ctx context.Context, exercise *BaseExercise) error {
	if err := m.validateBaseExercise(exercise); err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET name = $1, Notes = $2, updated_at = $3 WHERE id = $4", m.name)
	result, err := m.db.ExecContext(ctx, query, exercise.Name, exercise.Notes, time.Now(), exercise.ID)
	if err != nil {
		return fmt.Errorf("error updating exercise: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return ErrExerciseNotFound
	}

	return nil
}

// Delete removes an exercise from the database
func (m *ExerciseModel) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("%w: invalid exercise ID", ErrInvalidInput)
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", m.name)
	result, err := m.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting exercise: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return ErrExerciseNotFound
	}

	return nil
}

// UpdateCardio updates an existing cardio exercise
func (m *ExerciseModel) UpdateCardio(ctx context.Context, exercise *CardioExercise) error {
	if err := m.validateCardioExercise(exercise); err != nil {
		return err
	}

	query := fmt.Sprintf(`
		UPDATE %s 
		SET name = $1, Notes = $2, distance = $3, duration = $4, updated_at = $5 
		WHERE id = $6
	`, m.name)

	result, err := m.db.ExecContext(
		ctx,
		query,
		exercise.Name,
		exercise.Notes,
		exercise.Distance,
		exercise.Duration,
		time.Now(),
		exercise.ID,
	)
	if err != nil {
		return fmt.Errorf("error updating cardio exercise: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return ErrExerciseNotFound
	}

	return nil
}

// UpdateWeights updates an existing weight training exercise
func (m *ExerciseModel) UpdateWeights(ctx context.Context, exercise *WeightExercise) error {
	if err := m.validateWeightExercise(exercise); err != nil {
		return err
	}

	query := fmt.Sprintf(`
		UPDATE %s 
		SET name = $1, Notes = $2, sets = $3, reps = $4, weight = $5, updated_at = $6 
		WHERE id = $7
	`, m.name)

	result, err := m.db.ExecContext(
		ctx,
		query,
		exercise.Name,
		exercise.Notes,
		exercise.Sets,
		exercise.Reps,
		exercise.Weight,
		time.Now(),
		exercise.ID,
	)
	if err != nil {
		return fmt.Errorf("error updating weight exercise: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return ErrExerciseNotFound
	}

	return nil
}
