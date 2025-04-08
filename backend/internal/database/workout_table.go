package database

import (
	"database/sql"
	"fmt"
	"workoutapp/internal/models"
)

// Workout operations
func (db *Database) CreateWorkout(workout *models.Workout) error {
	query := `
		INSERT INTO workouts (user_id, start_time, end_time)
		VALUES ($1, $2, $3)
		RETURNING id`

	err := db.QueryRow(query,
		workout.UserID,
		workout.StartTime,
		workout.EndTime,
	).Scan(&workout.ID)

	if err != nil {
		return fmt.Errorf("error creating workout: %v", err)
	}
	return nil
}

func (db *Database) GetWorkout(id int64) (*models.Workout, error) {
	workout := &models.Workout{}
	query := `SELECT id, user_id, start_time, end_time FROM workouts WHERE id = $1`

	err := db.QueryRow(query, id).Scan(
		&workout.ID,
		&workout.UserID,
		&workout.StartTime,
		&workout.EndTime,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("workout not found")
		}
		return nil, fmt.Errorf("error getting workout: %v", err)
	}
	return workout, nil
}

func (db *Database) UpdateWorkout(workout *models.Workout) error {
	query := `
		UPDATE workouts 
		SET start_time = $1, end_time = $2
		WHERE id = $3 AND user_id = $4`

	result, err := db.Exec(query,
		workout.StartTime,
		workout.EndTime,
		workout.ID,
		workout.UserID,
	)

	if err != nil {
		return fmt.Errorf("error updating workout: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("workout not found or not owned by user")
	}

	return nil
}

func (db *Database) DeleteWorkout(id int64, userID int64) error {
	query := `DELETE FROM workouts WHERE id = $1 AND user_id = $2`

	result, err := db.Exec(query, id, userID)
	if err != nil {
		return fmt.Errorf("error deleting workout: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("workout not found or not owned by user")
	}

	return nil
}
