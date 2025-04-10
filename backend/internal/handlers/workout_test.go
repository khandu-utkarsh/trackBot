package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"workout_app_backend/internal/models"
	"workout_app_backend/internal/testutils"

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

	handler := GetWorkoutHandlerInstance(workoutModel)

	req := httptest.NewRequest("GET", "/api/users/"+strconv.FormatInt(userID, 10)+"/workouts", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("userID", strconv.FormatInt(userID, 10))
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()
	handler.ListWorkouts(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response []models.Workout
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if len(response) != len(testWorkouts) {
		t.Errorf("Expected %d workouts, got %d", len(testWorkouts), len(response))
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

	handler := GetWorkoutHandlerInstance(workoutModel)

	tests := []struct {
		name       string
		workout    *models.Workout
		wantStatus int
	}{
		{
			name: "Valid workout",
			workout: &models.Workout{
				UserID: userID,
			},
			wantStatus: http.StatusCreated,
		},
		{
			name: "Invalid user ID",
			workout: &models.Workout{
				UserID: 99999,
			},
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.workout)
			if err != nil {
				t.Fatalf("Failed to marshal workout: %v", err)
			}

			req := httptest.NewRequest("POST", "/api/users/"+strconv.FormatInt(tt.workout.UserID, 10)+"/workouts", bytes.NewBuffer(body))
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("userID", strconv.FormatInt(tt.workout.UserID, 10))
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
				if response.UserID != userID {
					t.Errorf("Expected user ID %d, got %d", userID, response.UserID)
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

	handler := GetWorkoutHandlerInstance(workoutModel)

	tests := []struct {
		name       string
		workoutID  int64
		wantStatus int
	}{
		{
			name:       "Existing workout",
			workoutID:  workoutID,
			wantStatus: http.StatusOK,
		},
		{
			name:       "Non-existent workout",
			workoutID:  99999,
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/api/users/"+strconv.FormatInt(userID, 10)+"/workouts/"+strconv.FormatInt(tt.workoutID, 10), nil)
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

	handler := GetWorkoutHandlerInstance(workoutModel)

	tests := []struct {
		name          string
		workoutID     int64
		updateWorkout *models.Workout
		wantStatus    int
	}{
		{
			name:      "Valid update",
			workoutID: workoutID,
			updateWorkout: &models.Workout{
				UserID: userID,
			},
			wantStatus: http.StatusOK,
		},
		{
			name:      "Non-existent workout",
			workoutID: 99999,
			updateWorkout: &models.Workout{
				UserID: userID,
			},
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.updateWorkout)
			if err != nil {
				t.Fatalf("Failed to marshal workout: %v", err)
			}

			req := httptest.NewRequest("PUT", "/api/users/"+strconv.FormatInt(tt.updateWorkout.UserID, 10)+"/workouts/"+strconv.FormatInt(tt.workoutID, 10), bytes.NewBuffer(body))
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("userID", strconv.FormatInt(tt.updateWorkout.UserID, 10))
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

	handler := GetWorkoutHandlerInstance(workoutModel)

	tests := []struct {
		name       string
		workoutID  int64
		wantStatus int
	}{
		{
			name:       "Existing workout",
			workoutID:  workoutID,
			wantStatus: http.StatusOK,
		},
		{
			name:       "Non-existent workout",
			workoutID:  99999,
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("DELETE", "/api/users/"+strconv.FormatInt(userID, 10)+"/workouts/"+strconv.FormatInt(tt.workoutID, 10), nil)
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
					t.Errorf("Expected deleted ID %d, got %d", tt.workoutID, response["deleted_id"])
				}
			}
		})
	}
}
