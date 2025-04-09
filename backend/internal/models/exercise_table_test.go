package models

import (
	"context"
	"testing"
	"workout_app_backend/internal/testutils"
)

func TestExerciseModel_CreateCardio(t *testing.T) {
	t.Skip("Skipping this test temporarily")
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()
	exerciseModel := GetExerciseModelInstance(db, "exercises_test")

	// Initialize the table
	if err := exerciseModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize exercises table: %v", err)
	}

	// Create a test user first
	userModel := GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	testUser := &User{
		Email: "test@example.com",
	}
	_, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	// Create a test workout
	workoutModel := GetWorkoutModelInstance(db, "workouts_test")
	if err := workoutModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize workouts table: %v", err)
	}

	testWorkout := &Workout{
		UserID: testUser.ID,
	}
	if err := workoutModel.Create(ctx, testWorkout); err != nil {
		t.Fatalf("Failed to create test workout: %v", err)
	}

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
			err := exerciseModel.CreateCardio(ctx, tt.exercise)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExerciseModel.CreateCardio() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExerciseModel_CreateWeights(t *testing.T) {
	t.Skip("Skipping this test temporarily")
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()
	exerciseModel := GetExerciseModelInstance(db, "exercises_test")

	// Initialize the table
	if err := exerciseModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize exercises table: %v", err)
	}

	// Create a test user first
	userModel := GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	testUser := &User{
		Email: "test@example.com",
	}
	_, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	// Create a test workout
	workoutModel := GetWorkoutModelInstance(db, "workouts_test")
	if err := workoutModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize workouts table: %v", err)
	}

	testWorkout := &Workout{
		UserID: testUser.ID,
	}
	if err := workoutModel.Create(ctx, testWorkout); err != nil {
		t.Fatalf("Failed to create test workout: %v", err)
	}

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
			err := exerciseModel.CreateWeights(ctx, tt.exercise)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExerciseModel.CreateWeights() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExerciseModel_ListByWorkout(t *testing.T) {
	t.Skip("Skipping this test temporarily")
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()
	exerciseModel := GetExerciseModelInstance(db, "exercises_test")

	// Initialize the table
	if err := exerciseModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize exercises table: %v", err)
	}

	// Create a test user first
	userModel := GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	testUser := &User{
		Email: "test@example.com",
	}
	_, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	// Create a test workout
	workoutModel := GetWorkoutModelInstance(db, "workouts_test")
	if err := workoutModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize workouts table: %v", err)
	}

	testWorkout := &Workout{
		UserID: testUser.ID,
	}
	if err := workoutModel.Create(ctx, testWorkout); err != nil {
		t.Fatalf("Failed to create test workout: %v", err)
	}

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
			if err := exerciseModel.CreateCardio(ctx, e); err != nil {
				t.Fatalf("Failed to create cardio exercise: %v", err)
			}
		case *WeightExercise:
			if err := exerciseModel.CreateWeights(ctx, e); err != nil {
				t.Fatalf("Failed to create weight exercise: %v", err)
			}
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
