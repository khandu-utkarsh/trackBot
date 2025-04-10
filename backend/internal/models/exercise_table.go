package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"workout_app_backend/internal/database"
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
	db         database.Database
	name       string
	foreignKey string
}

// GetExerciseModelInstance creates a new ExerciseModel instance
func GetExerciseModelInstance(db database.Database, name string, foreignKey string) *ExerciseModel {
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

	return m.db.CreateTable(ctx, m.name, schema)
}

// CreateCardio creates a new cardio exercise
func (m *ExerciseModel) CreateCardio(ctx context.Context, exercise *CardioExercise) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s (workout_id, name, Type, Notes, distance, duration) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", m.name)

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
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error creating cardio exercise: %v", err)
	}

	return id, nil
}

// CreateWeights creates a new weight training exercise
func (m *ExerciseModel) CreateWeights(ctx context.Context, exercise *WeightExercise) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s (workout_id, name, Type, Notes, sets, reps, weight) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", m.name)

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
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("error creating weight exercise: %v", err)
	}

	return id, nil
}

// Get retrieves an exercise by ID and returns the correct struct type (CardioExercise or WeightExercise)
func (m *ExerciseModel) Get(ctx context.Context, id int64) (interface{}, error) {
	query := fmt.Sprintf(`
		SELECT id, workout_id, name, Type, Notes,
		       distance, duration,
		       sets, reps, weight,
		       created_at, updated_at
		FROM %s
		WHERE id = $1
	`, m.name)

	var (
		base         BaseExercise
		exerciseType ExerciseType
		distance     sql.NullFloat64
		duration     sql.NullInt64
		sets         sql.NullInt64
		reps         sql.NullInt64
		weight       sql.NullFloat64
	)

	err := m.db.QueryRowContext(ctx, query, id).Scan(
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
		return nil, fmt.Errorf("exercise not found")
	}
	if err != nil {
		return nil, fmt.Errorf("error getting exercise: %v", err)
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

// ListByWorkout retrieves all exercises for a workout
func (m *ExerciseModel) ListByWorkout(ctx context.Context, workoutID int64) ([]interface{}, error) {
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
		return nil, fmt.Errorf("error listing exercises: %v", err)
	}
	defer rows.Close()

	var exercises []interface{}

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

		err := rows.Scan(
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
		if err != nil {
			return nil, fmt.Errorf("error scanning exercise: %v", err)
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
			// You can choose to skip or error out for unknown types
			return nil, fmt.Errorf("unknown exercise type: %s", exerciseType)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating exercises: %v", err)
	}

	return exercises, nil
}

func (m *ExerciseModel) Update(ctx context.Context, exercise *BaseExercise) error {
	_, err := m.db.ExecContext(ctx, fmt.Sprintf("UPDATE %s SET name = $1, Notes = $2, updated_at = $3 WHERE id = $4", m.name),
		exercise.Name, exercise.Notes, time.Now(), exercise.ID)
	return err
}

func (m *ExerciseModel) Delete(ctx context.Context, id int64) error {
	_, err := m.db.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE id = $1", m.name), id)
	return err
}
