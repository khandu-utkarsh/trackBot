package models

import (
	"context"
	"log"
	"workout_app_backend/internal/database"
)

func InitializeModels(db database.Database) error {

	userTable := GetUserModelInstance(db, "users").Initialize(context.Background())
	workoutTable := GetWorkoutModelInstance(db, "workouts", "users").Initialize(context.Background())
	exerciseTable := GetExerciseModelInstance(db, "exercises", "workouts").Initialize(context.Background())

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
