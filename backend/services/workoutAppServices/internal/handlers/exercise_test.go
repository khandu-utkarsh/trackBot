package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"workout_app_backend/services/workoutAppServices/internal/models"
	"workout_app_backend/services/workoutAppServices/internal/testutils"

	"github.com/go-chi/chi/v5"
)

func TestExerciseHandler_CreateExercise(t *testing.T) {
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()

	// Create test tables
	userModel := models.GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	workoutModel := models.GetWorkoutModelInstance(db, "workouts_test", "users_test")
	if err := workoutModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize workouts table: %v", err)
	}

	exerciseModel := models.GetExerciseModelInstance(db, "exercises_test", "workouts_test")
	if err := exerciseModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize exercises table: %v", err)
	}

	// Create a test user and workout
	testUser := &models.User{
		Email: "test@example.com",
	}
	userID, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	testWorkout := &models.Workout{
		UserID: userID,
	}
	workoutID, err := workoutModel.Create(ctx, testWorkout)
	if err != nil {
		t.Fatalf("Failed to create test workout: %v", err)
	}

	handler := GetExerciseHandlerInstance(exerciseModel, workoutModel)

	tests := []struct {
		name       string
		method     string
		exercise   interface{}
		wantStatus int
		wantError  bool
	}{
		{
			name:   "Valid cardio exercise",
			method: http.MethodPost,
			exercise: &models.CardioExercise{
				BaseExercise: models.BaseExercise{
					WorkoutID: workoutID,
					Name:      "Running",
					Type:      models.ExerciseTypeCardio,
					Notes:     "Morning run",
				},
				Distance: 5000,
				Duration: 1800,
			},
			wantStatus: http.StatusCreated,
			wantError:  false,
		},
		{
			name:   "Valid weight exercise",
			method: http.MethodPost,
			exercise: &models.WeightExercise{
				BaseExercise: models.BaseExercise{
					WorkoutID: workoutID,
					Name:      "Bench Press",
					Type:      models.ExerciseTypeWeights,
					Notes:     "3 sets",
				},
				Sets:   3,
				Reps:   10,
				Weight: 60.0,
			},
			wantStatus: http.StatusCreated,
			wantError:  false,
		},
		{
			name:   "Invalid exercise type",
			method: http.MethodPost,
			exercise: map[string]interface{}{
				"type": "invalid",
			},
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
		{
			name:   "Invalid method",
			method: http.MethodGet,
			exercise: &models.CardioExercise{
				BaseExercise: models.BaseExercise{
					WorkoutID: workoutID,
					Name:      "Running",
					Type:      models.ExerciseTypeCardio,
				},
				Distance: 5000,
				Duration: 1800,
			},
			wantStatus: http.StatusMethodNotAllowed,
			wantError:  true,
		},
		{
			name:       "Invalid JSON",
			method:     http.MethodPost,
			exercise:   nil,
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body []byte
			var err error
			if tt.exercise != nil {
				body, err = json.Marshal(tt.exercise)
				if err != nil {
					t.Fatalf("Failed to marshal exercise: %v", err)
				}
			} else {
				body = []byte("invalid json")
			}

			var workoutIDStr string
			if tt.exercise != nil {
				switch e := tt.exercise.(type) {
				case *models.CardioExercise:
					workoutIDStr = strconv.FormatInt(e.WorkoutID, 10)
				case *models.WeightExercise:
					workoutIDStr = strconv.FormatInt(e.WorkoutID, 10)
				default:
					workoutIDStr = strconv.FormatInt(workoutID, 10)
				}
			} else {
				workoutIDStr = strconv.FormatInt(workoutID, 10)
			}

			req := httptest.NewRequest(tt.method, "/api/users/"+strconv.FormatInt(userID, 10)+"/workouts/"+workoutIDStr+"/exercises", bytes.NewBuffer(body))
			w := httptest.NewRecorder()
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("userID", strconv.FormatInt(userID, 10))
			rctx.URLParams.Add("workoutID", workoutIDStr)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			handler.CreateExercise(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusCreated {
				// Decode response based on exercise type
				switch tt.exercise.(type) {
				case *models.CardioExercise:
					var response models.CardioExercise
					if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
						t.Fatalf("Failed to decode response: %v", err)
					}
					if response.ID <= 0 {
						t.Error("Expected non-zero exercise ID")
					}
					if response.WorkoutID != workoutID {
						t.Errorf("Expected workout ID %d, got %d", workoutID, response.WorkoutID)
					}
				case *models.WeightExercise:
					var response models.WeightExercise
					if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
						t.Fatalf("Failed to decode response: %v", err)
					}
					if response.ID <= 0 {
						t.Error("Expected non-zero exercise ID")
					}
					if response.WorkoutID != workoutID {
						t.Errorf("Expected workout ID %d, got %d", workoutID, response.WorkoutID)
					}
				default:
					t.Fatal("Unexpected exercise type in test case")
				}
			}

			if tt.wantError {
				var errorResponse map[string]string
				if err := json.NewDecoder(w.Body).Decode(&errorResponse); err != nil {
					t.Fatalf("Failed to decode error response: %v", err)
				}
				if _, ok := errorResponse["error"]; !ok {
					t.Error("Expected error response to contain 'error' field")
				}
			}
		})
	}
}

func TestExerciseHandler_GetExercise(t *testing.T) {
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()

	// Create test tables
	userModel := models.GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	workoutModel := models.GetWorkoutModelInstance(db, "workouts_test", "users_test")
	if err := workoutModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize workouts table: %v", err)
	}

	exerciseModel := models.GetExerciseModelInstance(db, "exercises_test", "workouts_test")
	if err := exerciseModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize exercises table: %v", err)
	}

	// Create a test user and workout
	testUser := &models.User{
		Email: "test@example.com",
	}
	userID, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	testWorkout := &models.Workout{
		UserID: userID,
	}
	workoutID, err := workoutModel.Create(ctx, testWorkout)
	if err != nil {
		t.Fatalf("Failed to create test workout: %v", err)
	}

	// Create a test exercise
	testExercise := &models.CardioExercise{
		BaseExercise: models.BaseExercise{
			WorkoutID: workoutID,
			Name:      "Running",
			Type:      models.ExerciseTypeCardio,
			Notes:     "Morning run",
		},
		Distance: 5000,
		Duration: 1800,
	}
	exerciseID, err := exerciseModel.CreateCardio(ctx, testExercise)
	if err != nil {
		t.Fatalf("Failed to create test exercise: %v", err)
	}

	handler := GetExerciseHandlerInstance(exerciseModel, workoutModel)

	tests := []struct {
		name       string
		method     string
		exerciseID int64
		wantStatus int
		wantError  bool
	}{
		{
			name:       "Existing exercise",
			method:     http.MethodGet,
			exerciseID: exerciseID,
			wantStatus: http.StatusOK,
			wantError:  false,
		},
		{
			name:       "Non-existent exercise",
			method:     http.MethodGet,
			exerciseID: 99999,
			wantStatus: http.StatusNotFound,
			wantError:  true,
		},
		{
			name:       "Invalid method",
			method:     http.MethodPost,
			exerciseID: exerciseID,
			wantStatus: http.StatusMethodNotAllowed,
			wantError:  true,
		},
		{
			name:       "Invalid exercise ID format",
			method:     http.MethodGet,
			exerciseID: -1,
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			workoutIDStr := strconv.FormatInt(workoutID, 10)
			req := httptest.NewRequest(tt.method, "/api/users/"+strconv.FormatInt(userID, 10)+"/workouts/"+workoutIDStr+"/exercises/"+strconv.FormatInt(tt.exerciseID, 10), nil)
			w := httptest.NewRecorder()
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("userID", strconv.FormatInt(userID, 10))
			rctx.URLParams.Add("workoutID", workoutIDStr)
			rctx.URLParams.Add("exerciseID", strconv.FormatInt(tt.exerciseID, 10))
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			handler.GetExercise(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusOK {
				var response models.CardioExercise
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if response.ID != tt.exerciseID {
					t.Errorf("Expected exercise ID %d, got %d", tt.exerciseID, response.ID)
				}
			}

			if tt.wantError {
				var errorResponse map[string]string
				if err := json.NewDecoder(w.Body).Decode(&errorResponse); err != nil {
					t.Fatalf("Failed to decode error response: %v", err)
				}
				if _, ok := errorResponse["error"]; !ok {
					t.Error("Expected error response to contain 'error' field")
				}
			}
		})
	}
}

func TestExerciseHandler_ListExercisesByWorkout(t *testing.T) {
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()

	// Create test tables
	userModel := models.GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	workoutModel := models.GetWorkoutModelInstance(db, "workouts_test", "users_test")
	if err := workoutModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize workouts table: %v", err)
	}

	exerciseModel := models.GetExerciseModelInstance(db, "exercises_test", "workouts_test")
	if err := exerciseModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize exercises table: %v", err)
	}

	// Create a test user and workout
	testUser := &models.User{
		Email: "test@example.com",
	}
	userID, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	testWorkout := &models.Workout{
		UserID: userID,
	}
	workoutID, err := workoutModel.Create(ctx, testWorkout)
	if err != nil {
		t.Fatalf("Failed to create test workout: %v", err)
	}

	// Create test exercises
	testExercises := []*models.CardioExercise{
		{
			BaseExercise: models.BaseExercise{
				WorkoutID: workoutID,
				Name:      "Running",
				Type:      models.ExerciseTypeCardio,
				Notes:     "Morning run",
			},
			Distance: 5000,
			Duration: 1800,
		},
		{
			BaseExercise: models.BaseExercise{
				WorkoutID: workoutID,
				Name:      "Cycling",
				Type:      models.ExerciseTypeCardio,
				Notes:     "Evening ride",
			},
			Distance: 15000,
			Duration: 2700,
		},
	}
	for _, exercise := range testExercises {
		_, err := exerciseModel.CreateCardio(ctx, exercise)
		if err != nil {
			t.Fatalf("Failed to create test exercise: %v", err)
		}
	}

	handler := GetExerciseHandlerInstance(exerciseModel, workoutModel)

	tests := []struct {
		name       string
		method     string
		workoutID  int64
		wantStatus int
		wantError  bool
	}{
		{
			name:       "Valid workout ID",
			method:     http.MethodGet,
			workoutID:  workoutID,
			wantStatus: http.StatusOK,
			wantError:  false,
		},
		{
			name:       "Non-existent workout",
			method:     http.MethodGet,
			workoutID:  99999,
			wantStatus: http.StatusNotFound, // No, empty list, workout not found
			wantError:  false,
		},
		{
			name:       "Invalid method",
			method:     http.MethodPost,
			workoutID:  workoutID,
			wantStatus: http.StatusMethodNotAllowed,
			wantError:  true,
		},
		{
			name:       "Invalid workout ID format",
			method:     http.MethodGet,
			workoutID:  -1,
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/api/users/"+strconv.FormatInt(userID, 10)+"/workouts/"+strconv.FormatInt(tt.workoutID, 10)+"/exercises", nil)
			w := httptest.NewRecorder()
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("userID", strconv.FormatInt(userID, 10))
			rctx.URLParams.Add("workoutID", strconv.FormatInt(tt.workoutID, 10))
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			handler.ListExercisesByWorkout(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusOK {
				var response []models.CardioExercise
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if tt.workoutID == workoutID && len(response) != len(testExercises) {
					t.Errorf("Expected %d exercises, got %d", len(testExercises), len(response))
				}
			}

			if tt.wantError {
				var errorResponse map[string]string
				if err := json.NewDecoder(w.Body).Decode(&errorResponse); err != nil {
					t.Fatalf("Failed to decode error response: %v", err)
				}
				if _, ok := errorResponse["error"]; !ok {
					t.Error("Expected error response to contain 'error' field")
				}
			}
		})
	}
}

func TestExerciseHandler_UpdateExercise(t *testing.T) {
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()

	// Create test tables
	userModel := models.GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	workoutModel := models.GetWorkoutModelInstance(db, "workouts_test", "users_test")
	if err := workoutModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize workouts table: %v", err)
	}

	exerciseModel := models.GetExerciseModelInstance(db, "exercises_test", "workouts_test")
	if err := exerciseModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize exercises table: %v", err)
	}

	// Create a test user and workout
	testUser := &models.User{
		Email: "test@example.com",
	}
	userID, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	testWorkout := &models.Workout{
		UserID: userID,
	}
	workoutID, err := workoutModel.Create(ctx, testWorkout)
	if err != nil {
		t.Fatalf("Failed to create test workout: %v", err)
	}

	// Create a test exercise
	testExercise := &models.CardioExercise{
		BaseExercise: models.BaseExercise{
			WorkoutID: workoutID,
			Name:      "Running",
			Type:      models.ExerciseTypeCardio,
			Notes:     "Morning run",
		},
		Distance: 5000,
		Duration: 1800,
	}
	exerciseID, err := exerciseModel.CreateCardio(ctx, testExercise)
	if err != nil {
		t.Fatalf("Failed to create test exercise: %v", err)
	}

	handler := GetExerciseHandlerInstance(exerciseModel, workoutModel)

	tests := []struct {
		name           string
		method         string
		exerciseID     int64
		updateExercise interface{}
		wantStatus     int
		wantError      bool
	}{
		{
			name:       "Valid update",
			method:     http.MethodPut,
			exerciseID: exerciseID,
			updateExercise: &models.CardioExercise{
				BaseExercise: models.BaseExercise{
					WorkoutID: workoutID,
					Name:      "Updated Running",
					Type:      models.ExerciseTypeCardio,
					Notes:     "Updated morning run",
				},
				Distance: 7500,
				Duration: 2700,
			},
			wantStatus: http.StatusOK,
			wantError:  false,
		},
		{
			name:       "Non-existent exercise",
			method:     http.MethodPut,
			exerciseID: 99999,
			updateExercise: &models.CardioExercise{
				BaseExercise: models.BaseExercise{
					WorkoutID: workoutID,
					Name:      "Updated Running",
					Type:      models.ExerciseTypeCardio,
				},
				Distance: 7500,
				Duration: 2700,
			},
			wantStatus: http.StatusNotFound,
			wantError:  true,
		},
		{
			name:       "Invalid method",
			method:     http.MethodPost,
			exerciseID: exerciseID,
			updateExercise: &models.CardioExercise{
				BaseExercise: models.BaseExercise{
					WorkoutID: workoutID,
					Name:      "Updated Running",
					Type:      models.ExerciseTypeCardio,
				},
				Distance: 7500,
				Duration: 2700,
			},
			wantStatus: http.StatusMethodNotAllowed,
			wantError:  true,
		},
		{
			name:       "Invalid exercise data",
			method:     http.MethodPut,
			exerciseID: exerciseID,
			updateExercise: map[string]interface{}{
				"invalid": "data",
			},
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.updateExercise)
			if err != nil {
				t.Fatalf("Failed to marshal exercise: %v", err)
			}

			req := httptest.NewRequest(tt.method, "/api/users/"+strconv.FormatInt(userID, 10)+"/workouts/"+strconv.FormatInt(workoutID, 10)+"/exercises/"+strconv.FormatInt(tt.exerciseID, 10), bytes.NewBuffer(body))
			w := httptest.NewRecorder()
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("userID", strconv.FormatInt(userID, 10))
			rctx.URLParams.Add("workoutID", strconv.FormatInt(workoutID, 10))
			rctx.URLParams.Add("exerciseID", strconv.FormatInt(tt.exerciseID, 10))
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			handler.UpdateExercise(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusOK {
				var response models.CardioExercise
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if response.ID != tt.exerciseID {
					t.Errorf("Expected exercise ID %d, got %d", tt.exerciseID, response.ID)
				}
				if response.WorkoutID != workoutID {
					t.Errorf("Expected workout ID %d, got %d", workoutID, response.WorkoutID)
				}
				// Verify the update was applied
				if response.Name != tt.updateExercise.(*models.CardioExercise).Name {
					t.Errorf("Expected name %s, got %s", tt.updateExercise.(*models.CardioExercise).Name, response.Name)
				}
				if response.Distance != tt.updateExercise.(*models.CardioExercise).Distance {
					t.Errorf("Expected distance %f, got %f", tt.updateExercise.(*models.CardioExercise).Distance, response.Distance)
				}
				if response.Duration != tt.updateExercise.(*models.CardioExercise).Duration {
					t.Errorf("Expected duration %d, got %d", tt.updateExercise.(*models.CardioExercise).Duration, response.Duration)
				}
			}

			if tt.wantError {
				var errorResponse map[string]string
				if err := json.NewDecoder(w.Body).Decode(&errorResponse); err != nil {
					t.Fatalf("Failed to decode error response: %v", err)
				}
				if _, ok := errorResponse["error"]; !ok {
					t.Error("Expected error response to contain 'error' field")
				}
			}
		})
	}
}

func TestExerciseHandler_DeleteExercise(t *testing.T) {
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()

	// Create test tables
	userModel := models.GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	workoutModel := models.GetWorkoutModelInstance(db, "workouts_test", "users_test")
	if err := workoutModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize workouts table: %v", err)
	}

	exerciseModel := models.GetExerciseModelInstance(db, "exercises_test", "workouts_test")
	if err := exerciseModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize exercises table: %v", err)
	}

	// Create a test user and workout
	testUser := &models.User{
		Email: "test@example.com",
	}
	userID, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	testWorkout := &models.Workout{
		UserID: userID,
	}
	workoutID, err := workoutModel.Create(ctx, testWorkout)
	if err != nil {
		t.Fatalf("Failed to create test workout: %v", err)
	}

	// Create a test exercise
	testExercise := &models.CardioExercise{
		BaseExercise: models.BaseExercise{
			WorkoutID: workoutID,
			Name:      "Running",
			Type:      models.ExerciseTypeCardio,
			Notes:     "Morning run",
		},
		Distance: 5000,
		Duration: 1800,
	}
	exerciseID, err := exerciseModel.CreateCardio(ctx, testExercise)
	if err != nil {
		t.Fatalf("Failed to create test exercise: %v", err)
	}

	handler := GetExerciseHandlerInstance(exerciseModel, workoutModel)

	tests := []struct {
		name       string
		method     string
		exerciseID int64
		wantStatus int
		wantError  bool
	}{
		{
			name:       "Existing exercise",
			method:     http.MethodDelete,
			exerciseID: exerciseID,
			wantStatus: http.StatusOK,
			wantError:  false,
		},
		{
			name:       "Non-existent exercise",
			method:     http.MethodDelete,
			exerciseID: 99999,
			wantStatus: http.StatusNotFound,
			wantError:  true,
		},
		{
			name:       "Invalid method",
			method:     http.MethodPost,
			exerciseID: exerciseID,
			wantStatus: http.StatusMethodNotAllowed,
			wantError:  true,
		},
		{
			name:       "Invalid exercise ID format",
			method:     http.MethodDelete,
			exerciseID: -1,
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/api/users/"+strconv.FormatInt(userID, 10)+"/workouts/"+strconv.FormatInt(workoutID, 10)+"/exercises/"+strconv.FormatInt(tt.exerciseID, 10), nil)
			w := httptest.NewRecorder()
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("userID", strconv.FormatInt(userID, 10))
			rctx.URLParams.Add("workoutID", strconv.FormatInt(workoutID, 10))
			rctx.URLParams.Add("exerciseID", strconv.FormatInt(tt.exerciseID, 10))
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			handler.DeleteExercise(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusOK {
				var response map[string]int64
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if response["deleted_id"] != tt.exerciseID {
					t.Errorf("Expected deleted_id %d, got %d", tt.exerciseID, response["deleted_id"])
				}

				// Verify the exercise was actually deleted
				_, err := exerciseModel.Get(ctx, tt.exerciseID)
				if err == nil {
					t.Error("Expected exercise to be deleted")
				}
			}

			if tt.wantError {
				var errorResponse map[string]string
				if err := json.NewDecoder(w.Body).Decode(&errorResponse); err != nil {
					t.Fatalf("Failed to decode error response: %v", err)
				}
				if _, ok := errorResponse["error"]; !ok {
					t.Error("Expected error response to contain 'error' field")
				}
			}
		})
	}
}
