package database

import (
	"fmt"
	"workoutapp/internal/models"
)

// Exercise operations
func (db *Database) CreateExercise(exercise *models.Exercise) error {
	query := `
		INSERT INTO exercises (
			workout_id, name, type, order_index,
			sets, reps, weight,
			duration, distance,
			notes
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id, created_at`

	err := db.QueryRow(query,
		exercise.WorkoutID,
		exercise.Name,
		exercise.Type,
		exercise.Order,
		exercise.Sets,
		exercise.Reps,
		exercise.Weight,
		exercise.Duration,
		exercise.Distance,
		exercise.Notes,
	).Scan(&exercise.ID, &exercise.CreatedAt)

	if err != nil {
		return fmt.Errorf("error creating exercise: %v", err)
	}
	return nil
}

func (db *Database) GetExercisesForWorkout(workoutID int64) ([]models.Exercise, error) {
	query := `
		SELECT id, workout_id, name, type, order_index,
			sets, reps, weight,
			duration, distance,
			notes, created_at
		FROM exercises
		WHERE workout_id = $1
		ORDER BY order_index`

	rows, err := db.Query(query, workoutID)
	if err != nil {
		return nil, fmt.Errorf("error getting exercises: %v", err)
	}
	defer rows.Close()

	var exercises []models.Exercise
	for rows.Next() {
		var exercise models.Exercise
		err := rows.Scan(
			&exercise.ID,
			&exercise.WorkoutID,
			&exercise.Name,
			&exercise.Type,
			&exercise.Order,
			&exercise.Sets,
			&exercise.Reps,
			&exercise.Weight,
			&exercise.Duration,
			&exercise.Distance,
			&exercise.Notes,
			&exercise.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning exercise: %v", err)
		}
		exercises = append(exercises, exercise)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating exercises: %v", err)
	}

	return exercises, nil
}

func (db *Database) UpdateExercise(exercise *models.Exercise) error {
	query := `
		UPDATE exercises 
		SET name = $1, type = $2, order_index = $3,
			sets = $4, reps = $5, weight = $6,
			duration = $7, distance = $8,
			notes = $9
		WHERE id = $10 AND workout_id = $11`

	result, err := db.Exec(query,
		exercise.Name,
		exercise.Type,
		exercise.Order,
		exercise.Sets,
		exercise.Reps,
		exercise.Weight,
		exercise.Duration,
		exercise.Distance,
		exercise.Notes,
		exercise.ID,
		exercise.WorkoutID,
	)

	if err != nil {
		return fmt.Errorf("error updating exercise: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("exercise not found")
	}

	return nil
}

func (db *Database) DeleteExercise(id int64, workoutID int64) error {
	query := `DELETE FROM exercises WHERE id = $1 AND workout_id = $2`

	result, err := db.Exec(query, id, workoutID)
	if err != nil {
		return fmt.Errorf("error deleting exercise: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("exercise not found")
	}

	return nil
}
