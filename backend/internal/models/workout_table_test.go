package models

import (
	"context"
	"testing"
	"workout_app_backend/internal/testutils"
)

func TestWorkoutModel_Create(t *testing.T) {
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()

	// Create a test user first
	//!Then create a workout
	userModel := GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	workoutModel := GetWorkoutModelInstance(db, "workouts_test", "users_test")
	if err := workoutModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize workouts table: %v", err)
	}

	testUser := &User{
		Email: "test@example.com",
	}

	id, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}
	testUser.ID = id

	// Test cases
	tests := []struct {
		name    string
		workout *Workout
		wantErr bool
	}{
		{
			name: "Valid workout",
			workout: &Workout{
				UserID: testUser.ID,
			},
			wantErr: false,
		},
		{
			name: "Invalid user ID",
			workout: &Workout{
				UserID: 99999, // Non-existent user
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := workoutModel.Create(ctx, tt.workout)
			if (err != nil) != tt.wantErr {
				t.Errorf("WorkoutModel.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.workout.ID = id
		})
	}
}

func TestWorkoutModel_Get(t *testing.T) {
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()

	// Create a test user
	userModel := GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	workoutModel := GetWorkoutModelInstance(db, "workouts_test", "users_test")
	if err := workoutModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize workouts table: %v", err)
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
		name    string
		id      int64
		wantErr bool
	}{
		{
			name:    "Existing workout",
			id:      testWorkout.ID,
			wantErr: false,
		},
		{
			name:    "Non-existent workout",
			id:      99999,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			workout, err := workoutModel.Get(ctx, tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("WorkoutModel.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && workout == nil {
				t.Error("WorkoutModel.Get() returned nil workout when error was expected")
			}
		})
	}
}

func TestWorkoutModel_List(t *testing.T) {
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()

	// Create a test user
	userModel := GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	workoutModel := GetWorkoutModelInstance(db, "workouts_test", "users_test")
	if err := workoutModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize workouts table: %v", err)
	}

	testUser := &User{
		Email: "test@example.com",
	}
	id, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}
	testUser.ID = id

	// Create multiple test workouts
	workouts := []*Workout{
		{
			UserID: testUser.ID,
		},
		{
			UserID: testUser.ID,
		},
	}

	for _, workout := range workouts {
		id, err := workoutModel.Create(ctx, workout)
		if err != nil {
			t.Fatalf("Failed to create test workout: %v", err)
		}
		workout.ID = id
	}

	// Test listing workouts
	listedWorkouts, err := workoutModel.List(ctx, testUser.ID)
	if err != nil {
		t.Fatalf("WorkoutModel.List() error = %v", err)
	}

	if len(listedWorkouts) < len(workouts) {
		t.Errorf("WorkoutModel.List() returned %d workouts, want at least %d",
			len(listedWorkouts), len(workouts))
	}
}
