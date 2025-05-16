package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	models "workout_app_backend/internal/models"

	"github.com/go-chi/chi/v5"
)

//!Functions needed
//ListWorkouts
//CreateWorkout
//GetWorkout
//UpdateWorkout
//DeleteWorkout

type WorkoutHandler struct {
	workoutModel *models.WorkoutModel
	userModel    *models.UserModel
}

func GetWorkoutHandlerInstance(workoutModel *models.WorkoutModel, userModel *models.UserModel) *WorkoutHandler {
	return &WorkoutHandler{workoutModel: workoutModel, userModel: userModel}
}

// ListWorkouts handles GET /api/users/{userID}/workouts
func (h *WorkoutHandler) ListWorkouts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if userID <= 0 {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	//!These are optional params
	yearStr := r.URL.Query().Get("year")
	monthStr := r.URL.Query().Get("month")
	dayStr := r.URL.Query().Get("day")

	ctx := r.Context()
	workouts, err := h.workoutModel.List(ctx, models.WorkoutListParams{
		UserID: userID,
		Year:   yearStr,
		Month:  monthStr,
		Day:    dayStr,
	})
	if err != nil {
		respondWithError(w, "Failed to list workouts", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, workouts)
}

// CreateWorkout handles POST /api/users/{userID}/workouts
func (h *WorkoutHandler) CreateWorkout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var workout models.Workout
	if err := json.NewDecoder(r.Body).Decode(&workout); err != nil {
		respondWithError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set the user ID from the URL path
	workout.UserID = userID

	// Verify user exists before creating workout
	ctx := r.Context()
	uout, err := h.userModel.Get(ctx, userID)
	fmt.Println(uout)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			respondWithError(w, "User not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to verify user", http.StatusInternalServerError)
		return
	}

	id, err := h.workoutModel.Create(ctx, &workout)
	if err != nil {
		respondWithError(w, "Failed to create workout", http.StatusInternalServerError)
		return
	}

	// Get the created workout to return complete data
	createdWorkout, err := h.workoutModel.Get(ctx, id)
	if err != nil {
		respondWithError(w, "Failed to get created workout", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, createdWorkout)
}

// GetWorkout handles GET /api/users/{userID}/workouts/{workoutID}
func (h *WorkoutHandler) GetWorkout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	workoutIDStr := chi.URLParam(r, "workoutID")
	workoutID, err := strconv.ParseInt(workoutIDStr, 10, 64)
	if err != nil || workoutID <= 0 {
		respondWithError(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	if workoutID <= 0 {
		respondWithError(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	workout, err := h.workoutModel.Get(ctx, workoutID)
	if err != nil {
		if errors.Is(err, models.ErrWorkoutNotFound) {
			respondWithError(w, "Workout not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to get workout", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, workout)
}

// UpdateWorkout handles PUT /api/users/{userID}/workouts/{workoutID}
func (h *WorkoutHandler) UpdateWorkout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	workoutIDStr := chi.URLParam(r, "workoutID")
	workoutID, err := strconv.ParseInt(workoutIDStr, 10, 64)
	if err != nil {
		respondWithError(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	if workoutID <= 0 {
		respondWithError(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if userID <= 0 {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var workout models.Workout
	if err := json.NewDecoder(r.Body).Decode(&workout); err != nil {
		respondWithError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set the IDs from the URL path
	workout.ID = workoutID
	workout.UserID = userID

	// Verify user exists before updating workout
	ctx := r.Context()
	_, err = h.userModel.Get(ctx, userID)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			respondWithError(w, "User not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to verify user", http.StatusInternalServerError)
		return
	}

	if err := h.workoutModel.Update(ctx, &workout); err != nil {
		if errors.Is(err, models.ErrWorkoutNotFound) {
			respondWithError(w, "Workout not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to update workout", http.StatusInternalServerError)
		return
	}

	// Get the updated workout to return complete data
	updatedWorkout, err := h.workoutModel.Get(ctx, workoutID)
	if err != nil {
		respondWithError(w, "Failed to get updated workout", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedWorkout)
}

// DeleteWorkout handles DELETE /api/users/{userID}/workouts/{workoutID}
func (h *WorkoutHandler) DeleteWorkout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	workoutIDStr := chi.URLParam(r, "workoutID")
	workoutID, err := strconv.ParseInt(workoutIDStr, 10, 64)
	if err != nil {
		respondWithError(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	if workoutID <= 0 {
		respondWithError(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if userID <= 0 {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Verify user exists before deleting workout
	ctx := r.Context()
	_, err = h.userModel.Get(ctx, userID)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			respondWithError(w, "User not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to verify user", http.StatusInternalServerError)
		return
	}

	if err := h.workoutModel.Delete(ctx, workoutID); err != nil {
		if errors.Is(err, models.ErrWorkoutNotFound) {
			respondWithError(w, "Workout not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to delete workout", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]int64{"deleted_id": workoutID})
}
