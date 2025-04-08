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
	Distance  float64 `json:"distance"`   // in kilometers
	Duration  int     `json:"duration"`   // in seconds
	Pace      float64 `json:"pace"`       // minutes per kilometer
	HeartRate int     `json:"heart_rate"` // beats per minute
}

// WeightExercise represents a weight training exercise with specific attributes
type WeightExercise struct {
	BaseExercise
	Sets     int     `json:"sets"`
	Reps     int     `json:"reps"`
	Weight   float64 `json:"weight"`    // in kilograms
	RestTime int     `json:"rest_time"` // in seconds
}

// ExerciseModel handles exercise-related database operations
type ExerciseModel struct {
	db database.Database
}

// NewExerciseModel creates a new ExerciseModel instance
func NewExerciseModel(db database.Database) *ExerciseModel {
	return &ExerciseModel{db: db}
}

// Initialize creates the exercises table if it doesn't exist
func (m *ExerciseModel) Initialize(ctx context.Context) error {
	schema := `
		id SERIAL PRIMARY KEY,
		workout_id INTEGER NOT NULL REFERENCES workouts(id) ON DELETE CASCADE,
		name VARCHAR(255) NOT NULL,
		type VARCHAR(20) NOT NULL CHECK (type IN ('cardio', 'weights')),
		notes TEXT,
		
		-- Cardio specific fields
		distance FLOAT,
		duration INTEGER,
		pace FLOAT,
		heart_rate INTEGER,
		
		-- Weight training specific fields
		sets INTEGER,
		reps INTEGER,
		weight FLOAT,
		rest_time INTEGER,
		
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	`

	return m.db.CreateTable(ctx, "exercises", schema)
}

// CreateCardio creates a new cardio exercise
func (m *ExerciseModel) CreateCardio(ctx context.Context, exercise *CardioExercise) error {
	query := `
		INSERT INTO exercises (
			workout_id, name, type, notes,
			distance, duration, pace, heart_rate
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, updated_at
	`

	err := m.db.QueryRowContext(
		ctx,
		query,
		exercise.WorkoutID,
		exercise.Name,
		ExerciseTypeCardio,
		exercise.Notes,
		exercise.Distance,
		exercise.Duration,
		exercise.Pace,
		exercise.HeartRate,
	).Scan(&exercise.ID, &exercise.CreatedAt, &exercise.UpdatedAt)

	if err != nil {
		return fmt.Errorf("error creating cardio exercise: %v", err)
	}

	return nil
}

// CreateWeights creates a new weight training exercise
func (m *ExerciseModel) CreateWeights(ctx context.Context, exercise *WeightExercise) error {
	query := `
		INSERT INTO exercises (
			workout_id, name, type, notes,
			sets, reps, weight, rest_time
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, updated_at
	`

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
		exercise.RestTime,
	).Scan(&exercise.ID, &exercise.CreatedAt, &exercise.UpdatedAt)

	if err != nil {
		return fmt.Errorf("error creating weight exercise: %v", err)
	}

	return nil
}

// Get retrieves an exercise by ID
func (m *ExerciseModel) Get(ctx context.Context, id int64) (interface{}, error) {
	query := `
		SELECT id, workout_id, name, type, notes,
			distance, duration, pace, heart_rate,
			sets, reps, weight, rest_time,
			created_at, updated_at
		FROM exercises
		WHERE id = $1
	`

	var (
		base         BaseExercise
		exerciseType ExerciseType
	)

	err := m.db.QueryRowContext(ctx, query, id).Scan(
		&base.ID,
		&base.WorkoutID,
		&base.Name,
		&exerciseType,
		&base.Notes,
		&base.CreatedAt,
		&base.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("exercise not found")
	}
	if err != nil {
		return nil, fmt.Errorf("error getting exercise: %v", err)
	}

	// Create the appropriate exercise type based on the type field
	switch exerciseType {
	case ExerciseTypeCardio:
		cardio := &CardioExercise{BaseExercise: base}
		err = m.db.QueryRowContext(ctx, query, id).Scan(
			&cardio.ID,
			&cardio.WorkoutID,
			&cardio.Name,
			&cardio.Type,
			&cardio.Notes,
			&cardio.Distance,
			&cardio.Duration,
			&cardio.Pace,
			&cardio.HeartRate,
			nil, nil, nil, nil, // weight training fields
			&cardio.CreatedAt,
			&cardio.UpdatedAt,
		)
		return cardio, err
	case ExerciseTypeWeights:
		weights := &WeightExercise{BaseExercise: base}
		err = m.db.QueryRowContext(ctx, query, id).Scan(
			&weights.ID,
			&weights.WorkoutID,
			&weights.Name,
			&weights.Type,
			&weights.Notes,
			nil, nil, nil, nil, // cardio fields
			&weights.Sets,
			&weights.Reps,
			&weights.Weight,
			&weights.RestTime,
			&weights.CreatedAt,
			&weights.UpdatedAt,
		)
		return weights, err
	default:
		return nil, fmt.Errorf("unknown exercise type: %s", exerciseType)
	}
}

// ListByWorkout retrieves all exercises for a workout
func (m *ExerciseModel) ListByWorkout(ctx context.Context, workoutID int64) ([]interface{}, error) {
	query := `
		SELECT id, workout_id, name, type, notes,
			distance, duration, pace, heart_rate,
			sets, reps, weight, rest_time,
			created_at, updated_at
		FROM exercises
		WHERE workout_id = $1
		ORDER BY id
	`

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
		)

		err := rows.Scan(
			&base.ID,
			&base.WorkoutID,
			&base.Name,
			&exerciseType,
			&base.Notes,
			&base.CreatedAt,
			&base.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning exercise: %v", err)
		}

		// Create the appropriate exercise type based on the type field
		switch exerciseType {
		case ExerciseTypeCardio:
			cardio := &CardioExercise{BaseExercise: base}
			err = rows.Scan(
				&cardio.ID,
				&cardio.WorkoutID,
				&cardio.Name,
				&cardio.Type,
				&cardio.Notes,
				&cardio.Distance,
				&cardio.Duration,
				&cardio.Pace,
				&cardio.HeartRate,
				nil, nil, nil, nil, // weight training fields
				&cardio.CreatedAt,
				&cardio.UpdatedAt,
			)
			exercises = append(exercises, cardio)
		case ExerciseTypeWeights:
			weights := &WeightExercise{BaseExercise: base}
			err = rows.Scan(
				&weights.ID,
				&weights.WorkoutID,
				&weights.Name,
				&weights.Type,
				&weights.Notes,
				nil, nil, nil, nil, // cardio fields
				&weights.Sets,
				&weights.Reps,
				&weights.Weight,
				&weights.RestTime,
				&weights.CreatedAt,
				&weights.UpdatedAt,
			)
			exercises = append(exercises, weights)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating exercises: %v", err)
	}

	return exercises, nil
}
