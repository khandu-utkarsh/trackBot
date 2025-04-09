package models

import (
	"context"
	"log"
	"workout_app_backend/internal/database"
)

func InitializeModels(db database.Database) error {

	userTable := GetUserModelInstance(db).Initialize(context.Background())
	workoutTable := GetWorkoutModelInstance(db).Initialize(context.Background())
	exerciseTable := GetExerciseModelInstance(db).Initialize(context.Background())

	if userTable != nil {
		log.Fatal(userTable)
	}
	if workoutTable != nil {
		log.Fatal(workoutTable)
	}
	if exerciseTable != nil {
		log.Fatal(exerciseTable)
	}
	return nil
}
