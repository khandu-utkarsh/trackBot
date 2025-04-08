package models

import (
	"context"
	"log"
	"workout_app_backend/internal/database"
)

func InitializeModels(db database.Database) error {
	workoutTable := NewWorkoutModel(db)
	exerciseTable := NewExerciseModel(db)
	if err := workoutTable.Initialize(context.Background()); err != nil {
		log.Fatal(err)
	}
	if err := exerciseTable.Initialize(context.Background()); err != nil {
		log.Fatal(err)
	}
	return nil
}
