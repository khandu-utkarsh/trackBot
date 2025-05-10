package models

import (
	"context"
	"strings"
	"testing"
	"workout_app_backend/services/workoutAppServices/internal/testutils"
)

func setupTestWorkoutModel(t *testing.T) (*WorkoutModel, *UserModel, *User, context.Context, func()) {
	db, cleanup := testutils.SetupTestDB(t)
	ctx := context.Background()

	// Initialize user model and create test user
	userModel := GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	testUser := &User{
		Email: "test@example.com",
	}
	id, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}
	testUser.ID = id

	// Initialize workout model
	workoutModel := GetWorkoutModelInstance(db, "workouts_test", "users_test")
	if err := workoutModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize workouts table: %v", err)
	}

	return workoutModel, userModel, testUser, ctx, cleanup
}

func TestWorkoutModel_Create(t *testing.T) {
	workoutModel, _, testUser, ctx, cleanup := setupTestWorkoutModel(t)
	defer cleanup()

	tests := []struct {
		name    string
		workout *Workout
		wantErr bool
		errMsg  string
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
			errMsg:  "violates foreign key constraint",
		},
		{
			name:    "Nil workout",
			workout: nil,
			wantErr: true,
			errMsg:  "workout cannot be nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := workoutModel.Create(ctx, tt.workout)
			if (err != nil) != tt.wantErr {
				t.Errorf("WorkoutModel.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if id <= 0 {
					t.Error("WorkoutModel.Create() returned invalid ID")
				}
				// Verify the workout was created correctly
				workout, err := workoutModel.Get(ctx, id)
				if err != nil {
					t.Errorf("Failed to get created workout: %v", err)
				}
				if workout.UserID != tt.workout.UserID {
					t.Errorf("Created workout UserID = %v, want %v", workout.UserID, tt.workout.UserID)
				}
			} else if tt.errMsg != "" && err != nil {
				// Check if the error message contains the expected message
				if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("WorkoutModel.Create() error message = %v, want to contain %v", err, tt.errMsg)
				}
			}
		})
	}
}

func TestWorkoutModel_Get(t *testing.T) {
	workoutModel, _, testUser, ctx, cleanup := setupTestWorkoutModel(t)
	defer cleanup()

	// Create a test workout
	testWorkout := &Workout{
		UserID: testUser.ID,
	}
	id, err := workoutModel.Create(ctx, testWorkout)
	if err != nil {
		t.Fatalf("Failed to create test workout: %v", err)
	}
	testWorkout.ID = id

	tests := []struct {
		name    string
		id      int64
		want    *Workout
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Existing workout",
			id:      testWorkout.ID,
			want:    testWorkout,
			wantErr: false,
		},
		{
			name:    "Non-existent workout",
			id:      99999,
			wantErr: true,
			errMsg:  "workout not found",
		},
		{
			name:    "Invalid ID",
			id:      0,
			wantErr: true,
			errMsg:  "invalid workout ID",
		},
		{
			name:    "Negative ID",
			id:      -1,
			wantErr: true,
			errMsg:  "invalid workout ID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			workout, err := workoutModel.Get(ctx, tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("WorkoutModel.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if workout == nil {
					t.Error("WorkoutModel.Get() returned nil workout")
					return
				}
				if workout.ID != tt.want.ID {
					t.Errorf("WorkoutModel.Get() ID = %v, want %v", workout.ID, tt.want.ID)
				}
				if workout.UserID != tt.want.UserID {
					t.Errorf("WorkoutModel.Get() UserID = %v, want %v", workout.UserID, tt.want.UserID)
				}
			} else if tt.errMsg != "" && err != nil {
				// Check if the error message contains the expected message
				if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("WorkoutModel.Get() error message = %v, want to contain %v", err, tt.errMsg)
				}
			}
		})
	}
}

func TestWorkoutModel_List(t *testing.T) {
	workoutModel, _, testUser, ctx, cleanup := setupTestWorkoutModel(t)
	defer cleanup()

	// Create multiple test workouts
	workouts := []*Workout{
		{UserID: testUser.ID},
		{UserID: testUser.ID},
		{UserID: testUser.ID},
	}

	for _, workout := range workouts {
		id, err := workoutModel.Create(ctx, workout)
		if err != nil {
			t.Fatalf("Failed to create test workout: %v", err)
		}
		workout.ID = id
	}

	// Test listing workouts
	list, err := workoutModel.List(ctx, testUser.ID)
	if err != nil {
		t.Fatalf("WorkoutModel.List() error = %v", err)
	}

	if len(list) != len(workouts) {
		t.Errorf("WorkoutModel.List() returned %d workouts, want %d", len(list), len(workouts))
	}

	// Verify each workout's data
	workoutMap := make(map[int64]*Workout)
	for _, workout := range workouts {
		workoutMap[workout.ID] = workout
	}

	for _, workout := range list {
		original, exists := workoutMap[workout.ID]
		if !exists {
			t.Errorf("WorkoutModel.List() returned unexpected workout with ID %d", workout.ID)
			continue
		}
		if workout.UserID != original.UserID {
			t.Errorf("WorkoutModel.List() workout UserID = %v, want %v", workout.UserID, original.UserID)
		}
	}
}

func TestWorkoutModel_Delete(t *testing.T) {
	workoutModel, _, testUser, ctx, cleanup := setupTestWorkoutModel(t)
	defer cleanup()

	// Create a test workout
	testWorkout := &Workout{
		UserID: testUser.ID,
	}
	id, err := workoutModel.Create(ctx, testWorkout)
	if err != nil {
		t.Fatalf("Failed to create test workout: %v", err)
	}

	tests := []struct {
		name    string
		id      int64
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Existing workout",
			id:      id,
			wantErr: false,
		},
		{
			name:    "Non-existent workout",
			id:      99999,
			wantErr: true,
			errMsg:  "workout not found",
		},
		{
			name:    "Invalid ID",
			id:      0,
			wantErr: true,
			errMsg:  "invalid workout ID",
		},
		{
			name:    "Negative ID",
			id:      -1,
			wantErr: true,
			errMsg:  "invalid workout ID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := workoutModel.Delete(ctx, tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("WorkoutModel.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// Verify the deletion
				_, err := workoutModel.Get(ctx, tt.id)
				if err == nil {
					t.Error("WorkoutModel.Delete() did not delete the workout")
				} else if err != ErrWorkoutNotFound {
					t.Errorf("Unexpected error after deletion: %v", err)
				}
			} else if tt.errMsg != "" && err != nil {
				// Check if the error message contains the expected message
				if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("WorkoutModel.Delete() error message = %v, want to contain %v", err, tt.errMsg)
				}
			}
		})
	}
}
