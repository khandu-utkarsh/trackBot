package models

import (
	"context"
	"testing"
	"workout_app_backend/internal/testutils"
)

func TestExerciseModel_CreateCardio(t *testing.T) {
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()

	// Create a test user first
	userModel := GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}
	// Create a test workout
	workoutModel := GetWorkoutModelInstance(db, "workouts_test", "users_test")
	if err := workoutModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize workouts table: %v", err)
	}

	exerciseModel := GetExerciseModelInstance(db, "exercises_test", "workouts_test")
	if err := exerciseModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize exercises table: %v", err)
	}

	testUser := &User{
		Email: "test@example.com",
	}
	id, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}
	testUser.ID = id

	testWorkout := &Workout{
		UserID: testUser.ID,
	}
	id, err = workoutModel.Create(ctx, testWorkout)
	if err != nil {
		t.Fatalf("Failed to create test workout: %v", err)
	}
	testWorkout.ID = id

	// Test cases
	tests := []struct {
		name     string
		exercise *CardioExercise
		wantErr  bool
	}{
		{
			name: "Valid cardio exercise",
			exercise: &CardioExercise{
				BaseExercise: BaseExercise{
					WorkoutID: testWorkout.ID,
					Name:      "Running",
					Type:      ExerciseTypeCardio,
					Notes:     "Morning run",
				},
				Distance: 5000,
				Duration: 1800,
			},
			wantErr: false,
		},
		{
			name: "Invalid workout ID",
			exercise: &CardioExercise{
				BaseExercise: BaseExercise{
					WorkoutID: 99999, // Non-existent workout
					Name:      "Running",
					Type:      ExerciseTypeCardio,
				},
				Distance: 5000,
				Duration: 1800,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := exerciseModel.CreateCardio(ctx, tt.exercise)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExerciseModel.CreateCardio() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.exercise.ID = id
		})
	}
}

func TestExerciseModel_CreateWeights(t *testing.T) {
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()
	// Create a test user first
	userModel := GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}
	// Create a test workout
	workoutModel := GetWorkoutModelInstance(db, "workouts_test", "users_test")
	if err := workoutModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize workouts table: %v", err)
	}

	exerciseModel := GetExerciseModelInstance(db, "exercises_test", "workouts_test")
	if err := exerciseModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize exercises table: %v", err)
	}

	testUser := &User{
		Email: "test@example.com",
	}
	id, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}
	testUser.ID = id
	// Create a test workout

	testWorkout := &Workout{
		UserID: testUser.ID,
	}
	id, err = workoutModel.Create(ctx, testWorkout)
	if err != nil {
		t.Fatalf("Failed to create test workout: %v", err)
	}
	testWorkout.ID = id

	// Test cases
	tests := []struct {
		name     string
		exercise *WeightExercise
		wantErr  bool
	}{
		{
			name: "Valid weight exercise",
			exercise: &WeightExercise{
				BaseExercise: BaseExercise{
					WorkoutID: testWorkout.ID,
					Name:      "Bench Press",
					Type:      ExerciseTypeWeights,
					Notes:     "3 sets",
				},
				Sets:   3,
				Reps:   10,
				Weight: 60.0,
			},
			wantErr: false,
		},
		{
			name: "Invalid workout ID",
			exercise: &WeightExercise{
				BaseExercise: BaseExercise{
					WorkoutID: 99999, // Non-existent workout
					Name:      "Bench Press",
					Type:      ExerciseTypeWeights,
				},
				Sets:   3,
				Reps:   10,
				Weight: 60.0,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := exerciseModel.CreateWeights(ctx, tt.exercise)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExerciseModel.CreateWeights() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.exercise.ID = id
		})
	}
}

func TestExerciseModel_ListByWorkout(t *testing.T) {
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()

	// Create a test user first
	userModel := GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}
	// Create a test workout
	workoutModel := GetWorkoutModelInstance(db, "workouts_test", "users_test")
	if err := workoutModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize workouts table: %v", err)
	}

	exerciseModel := GetExerciseModelInstance(db, "exercises_test", "workouts_test")
	if err := exerciseModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize exercises table: %v", err)
	}

	testUser := &User{
		Email: "test@example.com",
	}
	id, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}
	testUser.ID = id

	testWorkout := &Workout{
		UserID: testUser.ID,
	}
	id, err = workoutModel.Create(ctx, testWorkout)
	if err != nil {
		t.Fatalf("Failed to create test workout: %v", err)
	}
	testWorkout.ID = id

	// Create multiple test exercises
	exercises := []interface{}{
		&CardioExercise{
			BaseExercise: BaseExercise{
				WorkoutID: testWorkout.ID,
				Name:      "Running",
				Type:      ExerciseTypeCardio,
			},
			Distance: 5000,
			Duration: 1800,
		},
		&WeightExercise{
			BaseExercise: BaseExercise{
				WorkoutID: testWorkout.ID,
				Name:      "Bench Press",
				Type:      ExerciseTypeWeights,
			},
			Sets:   3,
			Reps:   10,
			Weight: 60.0,
		},
	}

	for _, exercise := range exercises {
		switch e := exercise.(type) {
		case *CardioExercise:
			id, err := exerciseModel.CreateCardio(ctx, e)
			if err != nil {
				t.Fatalf("Failed to create cardio exercise: %v", err)
			}
			e.ID = id
		case *WeightExercise:
			id, err := exerciseModel.CreateWeights(ctx, e)
			if err != nil {
				t.Fatalf("Failed to create weight exercise: %v", err)
			}
			e.ID = id
		}
	}

	// Test listing exercises
	listedExercises, err := exerciseModel.ListByWorkout(ctx, testWorkout.ID)
	if err != nil {
		t.Fatalf("ExerciseModel.ListByWorkout() error = %v", err)
	}

	if len(listedExercises) < len(exercises) {
		t.Errorf("ExerciseModel.ListByWorkout() returned %d exercises, want at least %d",
			len(listedExercises), len(exercises))
	}
}
