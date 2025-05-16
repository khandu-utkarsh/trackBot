package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	models "workout_app_backend/internal/models"
	testutils "workout_app_backend/internal/testutils"

	"github.com/go-chi/chi/v5"
)

func TestWorkoutHandler_ListWorkouts(t *testing.T) {
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()

	// Create a test user
	userModel := models.GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	workoutModel := models.GetWorkoutModelInstance(db, "workouts_test", "users_test")
	if err := workoutModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize workouts table: %v", err)
	}

	testUser := &models.User{
		Email: "test@example.com",
	}
	userID, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	// Create test workouts
	testWorkouts := []*models.Workout{
		{UserID: userID},
		{UserID: userID},
	}
	for _, workout := range testWorkouts {
		_, err := workoutModel.Create(ctx, workout)
		if err != nil {
			t.Fatalf("Failed to create test workout: %v", err)
		}
	}

	handler := GetWorkoutHandlerInstance(workoutModel, userModel)

	tests := []struct {
		name       string
		method     string
		wantStatus int
		wantError  bool
	}{
		{
			name:       "Valid GET request",
			method:     http.MethodGet,
			wantStatus: http.StatusOK,
			wantError:  false,
		},
		{
			name:       "Invalid method",
			method:     http.MethodPost,
			wantStatus: http.StatusMethodNotAllowed,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/api/users/"+strconv.FormatInt(userID, 10)+"/workouts", nil)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("userID", strconv.FormatInt(userID, 10))
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			w := httptest.NewRecorder()
			handler.ListWorkouts(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusOK {
				var response []models.Workout
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if len(response) != len(testWorkouts) {
					t.Errorf("Expected %d workouts, got %d", len(testWorkouts), len(response))
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

func TestWorkoutHandler_CreateWorkout(t *testing.T) {
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()

	// Create a test user
	userModel := models.GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	workoutModel := models.GetWorkoutModelInstance(db, "workouts_test", "users_test")
	if err := workoutModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize workouts table: %v", err)
	}

	testUser := &models.User{
		Email: "test@example.com",
	}
	userID, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	handler := GetWorkoutHandlerInstance(workoutModel, userModel)

	tests := []struct {
		name       string
		method     string
		workout    *models.Workout
		wantStatus int
		wantError  bool
	}{
		{
			name:   "Valid workout",
			method: http.MethodPost,
			workout: &models.Workout{
				UserID: userID,
			},
			wantStatus: http.StatusCreated,
			wantError:  false,
		},
		{
			name:   "Invalid user ID",
			method: http.MethodPost,
			workout: &models.Workout{
				UserID: 99999,
			},
			wantStatus: http.StatusNotFound,
			wantError:  true,
		},
		{
			name:   "Invalid method",
			method: http.MethodGet,
			workout: &models.Workout{
				UserID: userID,
			},
			wantStatus: http.StatusMethodNotAllowed,
			wantError:  true,
		},
		{
			name:       "Invalid JSON",
			method:     http.MethodPost,
			workout:    nil,
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body []byte
			var err error
			if tt.workout != nil {
				body, err = json.Marshal(tt.workout)
				if err != nil {
					t.Fatalf("Failed to marshal workout: %v", err)
				}
			} else {
				body = []byte("invalid json")
			}

			var userIDStr string
			if tt.workout != nil {
				userIDStr = strconv.FormatInt(tt.workout.UserID, 10)
			} else {
				userIDStr = "" // or "0", "-1", or "invalid" based on what you're testing
			}

			req := httptest.NewRequest(tt.method, "/api/users/"+userIDStr+"/workouts", bytes.NewBuffer(body))
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("userID", userIDStr)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			w := httptest.NewRecorder()
			handler.CreateWorkout(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusCreated {
				var response models.Workout
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if response.UserID != tt.workout.UserID {
					t.Errorf("Expected user ID %d, got %d", tt.workout.UserID, response.UserID)
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

func TestWorkoutHandler_GetWorkout(t *testing.T) {
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()

	// Create a test user
	userModel := models.GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	workoutModel := models.GetWorkoutModelInstance(db, "workouts_test", "users_test")
	if err := workoutModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize workouts table: %v", err)
	}

	testUser := &models.User{
		Email: "test@example.com",
	}
	userID, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	// Create a test workout
	testWorkout := &models.Workout{
		UserID: userID,
	}
	workoutID, err := workoutModel.Create(ctx, testWorkout)
	if err != nil {
		t.Fatalf("Failed to create test workout: %v", err)
	}

	handler := GetWorkoutHandlerInstance(workoutModel, userModel)

	tests := []struct {
		name       string
		method     string
		workoutID  int64
		wantStatus int
		wantError  bool
	}{
		{
			name:       "Existing workout",
			method:     http.MethodGet,
			workoutID:  workoutID,
			wantStatus: http.StatusOK,
			wantError:  false,
		},
		{
			name:       "Non-existent workout",
			method:     http.MethodGet,
			workoutID:  99999,
			wantStatus: http.StatusNotFound,
			wantError:  true,
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
			req := httptest.NewRequest(tt.method, "/api/users/"+strconv.FormatInt(userID, 10)+"/workouts/"+strconv.FormatInt(tt.workoutID, 10), nil)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("userID", strconv.FormatInt(userID, 10))
			rctx.URLParams.Add("workoutID", strconv.FormatInt(tt.workoutID, 10))
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			w := httptest.NewRecorder()
			handler.GetWorkout(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusOK {
				var response models.Workout
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if response.ID != tt.workoutID {
					t.Errorf("Expected workout ID %d, got %d", tt.workoutID, response.ID)
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

func TestWorkoutHandler_UpdateWorkout(t *testing.T) {
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()

	// Create a test user
	userModel := models.GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	workoutModel := models.GetWorkoutModelInstance(db, "workouts_test", "users_test")
	if err := workoutModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize workouts table: %v", err)
	}

	testUser := &models.User{
		Email: "test@example.com",
	}
	userID, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	// Create a test workout
	testWorkout := &models.Workout{
		UserID: userID,
	}
	workoutID, err := workoutModel.Create(ctx, testWorkout)
	if err != nil {
		t.Fatalf("Failed to create test workout: %v", err)
	}

	handler := GetWorkoutHandlerInstance(workoutModel, userModel)

	tests := []struct {
		name          string
		method        string
		workoutID     int64
		updateWorkout *models.Workout
		wantStatus    int
		wantError     bool
	}{
		{
			name:      "Valid update",
			method:    http.MethodPut,
			workoutID: workoutID,
			updateWorkout: &models.Workout{
				UserID: userID,
			},
			wantStatus: http.StatusOK,
			wantError:  false,
		},
		{
			name:      "Non-existent workout",
			method:    http.MethodPut,
			workoutID: 99999,
			updateWorkout: &models.Workout{
				UserID: userID,
			},
			wantStatus: http.StatusNotFound,
			wantError:  true,
		},
		{
			name:      "Invalid method",
			method:    http.MethodPost,
			workoutID: workoutID,
			updateWorkout: &models.Workout{
				UserID: userID,
			},
			wantStatus: http.StatusMethodNotAllowed,
			wantError:  true,
		},
		{
			name:      "Invalid user ID",
			method:    http.MethodPut,
			workoutID: workoutID,
			updateWorkout: &models.Workout{
				UserID: 99999,
			},
			wantStatus: http.StatusNotFound,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.updateWorkout)
			if err != nil {
				t.Fatalf("Failed to marshal workout: %v", err)
			}

			var userIDStr string
			if tt.updateWorkout != nil {
				userIDStr = strconv.FormatInt(tt.updateWorkout.UserID, 10)
			} else {
				userIDStr = "" // or "0", "-1", or "invalid" based on what you're testing
			}
			req := httptest.NewRequest(tt.method, "/api/users/"+userIDStr+"/workouts/"+strconv.FormatInt(tt.workoutID, 10), bytes.NewBuffer(body))
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("userID", userIDStr)
			rctx.URLParams.Add("workoutID", strconv.FormatInt(tt.workoutID, 10))
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			w := httptest.NewRecorder()
			handler.UpdateWorkout(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusOK {
				var response models.Workout
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if response.ID != tt.workoutID {
					t.Errorf("Expected workout ID %d, got %d", tt.workoutID, response.ID)
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

func TestWorkoutHandler_DeleteWorkout(t *testing.T) {
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()

	// Create a test user
	userModel := models.GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	workoutModel := models.GetWorkoutModelInstance(db, "workouts_test", "users_test")
	if err := workoutModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize workouts table: %v", err)
	}

	testUser := &models.User{
		Email: "test@example.com",
	}
	userID, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	// Create a test workout
	testWorkout := &models.Workout{
		UserID: userID,
	}
	workoutID, err := workoutModel.Create(ctx, testWorkout)
	if err != nil {
		t.Fatalf("Failed to create test workout: %v", err)
	}

	handler := GetWorkoutHandlerInstance(workoutModel, userModel)

	tests := []struct {
		name       string
		method     string
		workoutID  int64
		wantStatus int
		wantError  bool
	}{
		{
			name:       "Existing workout",
			method:     http.MethodDelete,
			workoutID:  workoutID,
			wantStatus: http.StatusOK,
			wantError:  false,
		},
		{
			name:       "Non-existent workout",
			method:     http.MethodDelete,
			workoutID:  99999,
			wantStatus: http.StatusNotFound,
			wantError:  true,
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
			method:     http.MethodDelete,
			workoutID:  -1,
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/api/users/"+strconv.FormatInt(userID, 10)+"/workouts/"+strconv.FormatInt(tt.workoutID, 10), nil)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("userID", strconv.FormatInt(userID, 10))
			rctx.URLParams.Add("workoutID", strconv.FormatInt(tt.workoutID, 10))
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			w := httptest.NewRecorder()
			handler.DeleteWorkout(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusOK {
				var response map[string]int64
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if response["deleted_id"] != tt.workoutID {
					t.Errorf("Expected deleted_id %d, got %d", tt.workoutID, response["deleted_id"])
				}

				// Verify the workout was actually deleted
				_, err := workoutModel.Get(ctx, tt.workoutID)
				if err == nil {
					t.Error("Expected workout to be deleted")
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
