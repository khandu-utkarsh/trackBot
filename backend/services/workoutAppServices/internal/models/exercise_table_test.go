package models

import (
	"context"
	"strings"
	"testing"
	testutils "workout_app_backend/internal/testutils"
)

func setupTestExerciseModel(t *testing.T) (*ExerciseModel, *WorkoutModel, *UserModel, *Workout, *User, context.Context, func()) {
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

	// Initialize workout model and create test workout
	workoutModel := GetWorkoutModelInstance(db, "workouts_test", "users_test")
	if err := workoutModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize workouts table: %v", err)
	}

	testWorkout := &Workout{
		UserID: testUser.ID,
	}
	id, err = workoutModel.Create(ctx, testWorkout)
	if err != nil {
		t.Fatalf("Failed to create test workout: %v", err)
	}
	testWorkout.ID = id

	// Initialize exercise model
	exerciseModel := GetExerciseModelInstance(db, "exercises_test", "workouts_test")
	if err := exerciseModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize exercises table: %v", err)
	}

	return exerciseModel, workoutModel, userModel, testWorkout, testUser, ctx, cleanup
}

func TestExerciseModel_CreateCardio(t *testing.T) {
	exerciseModel, _, _, testWorkout, _, ctx, cleanup := setupTestExerciseModel(t)
	defer cleanup()

	tests := []struct {
		name     string
		exercise *CardioExercise
		wantErr  bool
		errMsg   string
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
					WorkoutID: 99999,
					Name:      "Running",
					Type:      ExerciseTypeCardio,
				},
				Distance: 5000,
				Duration: 1800,
			},
			wantErr: true,
			errMsg:  "violates foreign key constraint",
		},
		{
			name: "Empty exercise name",
			exercise: &CardioExercise{
				BaseExercise: BaseExercise{
					WorkoutID: testWorkout.ID,
					Name:      "",
					Type:      ExerciseTypeCardio,
				},
				Distance: 5000,
				Duration: 1800,
			},
			wantErr: true,
			errMsg:  "exercise name cannot be empty",
		},
		{
			name: "Invalid exercise type",
			exercise: &CardioExercise{
				BaseExercise: BaseExercise{
					WorkoutID: testWorkout.ID,
					Name:      "Running",
					Type:      "invalid",
				},
				Distance: 5000,
				Duration: 1800,
			},
			wantErr: true,
			errMsg:  "invalid exercise type",
		},
		{
			name:     "Nil exercise",
			exercise: nil,
			wantErr:  true,
			errMsg:   "exercise cannot be nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := exerciseModel.CreateCardio(ctx, tt.exercise)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExerciseModel.CreateCardio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if id <= 0 {
					t.Error("ExerciseModel.CreateCardio() returned invalid ID")
				}
				// Verify the exercise was created correctly
				exercises, err := exerciseModel.ListByWorkout(ctx, tt.exercise.WorkoutID)
				if err != nil {
					t.Errorf("Failed to get created exercise: %v", err)
					return
				}
				found := false
				for _, e := range exercises {
					switch exercise := e.(type) {
					case *CardioExercise:
						if exercise.ID == id {
							found = true
							if exercise.Name != tt.exercise.Name {
								t.Errorf("Created exercise name = %v, want %v", exercise.Name, tt.exercise.Name)
							}
							if exercise.Type != tt.exercise.Type {
								t.Errorf("Created exercise type = %v, want %v", exercise.Type, tt.exercise.Type)
							}
							break
						}
					}
				}
				if !found {
					t.Error("Created exercise not found in list")
				}
			} else if tt.errMsg != "" && err != nil {
				// Check if the error message contains the expected message
				if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("ExerciseModel.CreateCardio() error message = %v, want to contain %v", err, tt.errMsg)
				}
			}
		})
	}
}

func TestExerciseModel_CreateWeights(t *testing.T) {
	exerciseModel, _, _, testWorkout, _, ctx, cleanup := setupTestExerciseModel(t)
	defer cleanup()

	tests := []struct {
		name     string
		exercise *WeightExercise
		wantErr  bool
		errMsg   string
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
					WorkoutID: 99999,
					Name:      "Bench Press",
					Type:      ExerciseTypeWeights,
				},
				Sets:   3,
				Reps:   10,
				Weight: 60.0,
			},
			wantErr: true,
			errMsg:  "violates foreign key constraint",
		},
		{
			name: "Empty exercise name",
			exercise: &WeightExercise{
				BaseExercise: BaseExercise{
					WorkoutID: testWorkout.ID,
					Name:      "",
					Type:      ExerciseTypeWeights,
				},
				Sets:   3,
				Reps:   10,
				Weight: 60.0,
			},
			wantErr: true,
			errMsg:  "exercise name cannot be empty",
		},
		{
			name: "Invalid exercise type",
			exercise: &WeightExercise{
				BaseExercise: BaseExercise{
					WorkoutID: testWorkout.ID,
					Name:      "Bench Press",
					Type:      "invalid",
				},
				Sets:   3,
				Reps:   10,
				Weight: 60.0,
			},
			wantErr: true,
			errMsg:  "invalid exercise type",
		},
		{
			name:     "Nil exercise",
			exercise: nil,
			wantErr:  true,
			errMsg:   "exercise cannot be nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := exerciseModel.CreateWeights(ctx, tt.exercise)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExerciseModel.CreateWeights() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if id <= 0 {
					t.Error("ExerciseModel.CreateWeights() returned invalid ID")
				}
				// Verify the exercise was created correctly
				exercises, err := exerciseModel.ListByWorkout(ctx, tt.exercise.WorkoutID)
				if err != nil {
					t.Errorf("Failed to get created exercise: %v", err)
					return
				}
				found := false
				for _, e := range exercises {
					switch exercise := e.(type) {
					case *WeightExercise:
						if exercise.ID == id {
							found = true
							if exercise.Name != tt.exercise.Name {
								t.Errorf("Created exercise name = %v, want %v", exercise.Name, tt.exercise.Name)
							}
							if exercise.Type != tt.exercise.Type {
								t.Errorf("Created exercise type = %v, want %v", exercise.Type, tt.exercise.Type)
							}
							break
						}
					}
				}
				if !found {
					t.Error("Created exercise not found in list")
				}
			} else if tt.errMsg != "" && err != nil {
				// Check if the error message contains the expected message
				if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("ExerciseModel.CreateWeights() error message = %v, want to contain %v", err, tt.errMsg)
				}
			}
		})
	}
}

func TestExerciseModel_ListByWorkout(t *testing.T) {
	exerciseModel, _, _, testWorkout, _, ctx, cleanup := setupTestExerciseModel(t)
	defer cleanup()

	// Create multiple test exercises
	exercises := []interface{}{
		&CardioExercise{
			BaseExercise: BaseExercise{
				WorkoutID: testWorkout.ID,
				Name:      "Running",
				Type:      ExerciseTypeCardio,
				Notes:     "Morning run",
			},
			Distance: 5000,
			Duration: 1800,
		},
		&WeightExercise{
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
	}

	// Create exercises and store their IDs
	exerciseIDs := make(map[string]int64)
	for _, exercise := range exercises {
		switch e := exercise.(type) {
		case *CardioExercise:
			id, err := exerciseModel.CreateCardio(ctx, e)
			if err != nil {
				t.Fatalf("Failed to create cardio exercise: %v", err)
			}
			e.ID = id
			exerciseIDs[e.Name] = id
		case *WeightExercise:
			id, err := exerciseModel.CreateWeights(ctx, e)
			if err != nil {
				t.Fatalf("Failed to create weight exercise: %v", err)
			}
			e.ID = id
			exerciseIDs[e.Name] = id
		}
	}

	// Test listing exercises
	list, err := exerciseModel.ListByWorkout(ctx, testWorkout.ID)
	if err != nil {
		t.Fatalf("ExerciseModel.ListByWorkout() error = %v", err)
	}

	if len(list) != len(exercises) {
		t.Errorf("ExerciseModel.ListByWorkout() returned %d exercises, want %d", len(list), len(exercises))
	}

	// Verify each exercise's data
	for _, exercise := range list {
		switch e := exercise.(type) {
		case *CardioExercise:
			expectedID, exists := exerciseIDs[e.Name]
			if !exists {
				t.Errorf("ExerciseModel.ListByWorkout() returned unexpected exercise with name %s", e.Name)
				continue
			}
			if e.ID != expectedID {
				t.Errorf("ExerciseModel.ListByWorkout() exercise ID = %v, want %v", e.ID, expectedID)
			}
			if e.WorkoutID != testWorkout.ID {
				t.Errorf("ExerciseModel.ListByWorkout() exercise WorkoutID = %v, want %v", e.WorkoutID, testWorkout.ID)
			}
		case *WeightExercise:
			expectedID, exists := exerciseIDs[e.Name]
			if !exists {
				t.Errorf("ExerciseModel.ListByWorkout() returned unexpected exercise with name %s", e.Name)
				continue
			}
			if e.ID != expectedID {
				t.Errorf("ExerciseModel.ListByWorkout() exercise ID = %v, want %v", e.ID, expectedID)
			}
			if e.WorkoutID != testWorkout.ID {
				t.Errorf("ExerciseModel.ListByWorkout() exercise WorkoutID = %v, want %v", e.WorkoutID, testWorkout.ID)
			}
		}
	}
}

func TestExerciseModel_Delete(t *testing.T) {
	exerciseModel, _, _, testWorkout, _, ctx, cleanup := setupTestExerciseModel(t)
	defer cleanup()

	// Create a test exercise
	testExercise := &CardioExercise{
		BaseExercise: BaseExercise{
			WorkoutID: testWorkout.ID,
			Name:      "Running",
			Type:      ExerciseTypeCardio,
		},
		Distance: 5000,
		Duration: 1800,
	}
	id, err := exerciseModel.CreateCardio(ctx, testExercise)
	if err != nil {
		t.Fatalf("Failed to create test exercise: %v", err)
	}

	tests := []struct {
		name    string
		id      int64
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Existing exercise",
			id:      id,
			wantErr: false,
		},
		{
			name:    "Non-existent exercise",
			id:      99999,
			wantErr: true,
			errMsg:  "exercise not found",
		},
		{
			name:    "Invalid ID",
			id:      0,
			wantErr: true,
			errMsg:  "invalid exercise ID",
		},
		{
			name:    "Negative ID",
			id:      -1,
			wantErr: true,
			errMsg:  "invalid exercise ID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := exerciseModel.Delete(ctx, tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExerciseModel.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// Verify the deletion
				exercises, err := exerciseModel.ListByWorkout(ctx, testWorkout.ID)
				if err != nil {
					t.Errorf("Failed to list exercises: %v", err)
					return
				}
				for _, e := range exercises {
					switch exercise := e.(type) {
					case *CardioExercise, *WeightExercise:
						if exercise.(interface{ GetID() int64 }).GetID() == tt.id {
							t.Error("ExerciseModel.Delete() did not delete the exercise")
							return
						}
					}
				}
			} else if tt.errMsg != "" && err != nil {
				// Check if the error message contains the expected message
				if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("ExerciseModel.Delete() error message = %v, want to contain %v", err, tt.errMsg)
				}
			}
		})
	}
}
