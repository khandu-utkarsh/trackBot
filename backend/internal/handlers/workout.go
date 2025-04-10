package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"workout_app_backend/internal/models"

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
}

func GetWorkoutHandlerInstance(workoutModel *models.WorkoutModel) *WorkoutHandler {
	return &WorkoutHandler{workoutModel: workoutModel}
}

// ListWorkouts handles GET /api/workouts
func (h *WorkoutHandler) ListWorkouts(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	workouts, err := h.workoutModel.List(r.Context(), userID)
	if err != nil {
		http.Error(w, "Failed to list workouts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(workouts); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// CreateWorkout handles POST /api/users/{userID}/workouts
func (h *WorkoutHandler) CreateWorkout(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var workout models.Workout
	if err := json.NewDecoder(r.Body).Decode(&workout); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	workout.UserID = userID
	id, err := h.workoutModel.Create(r.Context(), &workout)
	if err != nil {
		http.Error(w, "Failed to create workout", http.StatusNotFound)
		return
	}
	workout.ID = id

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(workout); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// GetWorkout handles GET /api/users/{userID}/workouts/{workoutID}
func (h *WorkoutHandler) GetWorkout(w http.ResponseWriter, r *http.Request) {
	workoutIDStr := chi.URLParam(r, "workoutID")
	workoutID, err := strconv.ParseInt(workoutIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	workout, err := h.workoutModel.Get(r.Context(), workoutID)
	if err != nil {
		http.Error(w, "Workout not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(workout); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// UpdateWorkout handles PUT /api/users/{userID}/workouts/{workoutID}
func (h *WorkoutHandler) UpdateWorkout(w http.ResponseWriter, r *http.Request) {
	workoutIDStr := chi.URLParam(r, "workoutID")
	workoutID, err := strconv.ParseInt(workoutIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	var workout models.Workout
	if err := json.NewDecoder(r.Body).Decode(&workout); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	workout.ID = workoutID
	if err := h.workoutModel.Update(r.Context(), &workout); err != nil {
		http.Error(w, "Failed to update workout", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(workout); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// DeleteWorkout handles DELETE /api/users/{userID}/workouts/{workoutID}
func (h *WorkoutHandler) DeleteWorkout(w http.ResponseWriter, r *http.Request) {
	workoutIDStr := chi.URLParam(r, "workoutID")
	workoutID, err := strconv.ParseInt(workoutIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	deletedID, err := h.workoutModel.Delete(r.Context(), workoutID)
	if err != nil {
		http.Error(w, "Workout not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]int64{"deleted_id": deletedID}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
