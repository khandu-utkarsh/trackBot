package handlers

import (
	"encoding/json"
	"net/http"
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
	// TODO: Implement logic to list workouts
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("ListWorkouts not implemented"))
}

// CreateWorkout handles POST /api/workouts (or /api/users/{userID}/workouts)
func (h *WorkoutHandler) CreateWorkout(w http.ResponseWriter, r *http.Request) {
	var workout models.Workout
	if err := json.NewDecoder(r.Body).Decode(&workout); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := h.workoutModel.Create(r.Context(), &workout)
	if err != nil {
		http.Error(w, "Failed to create workout", http.StatusInternalServerError)
		return
	}
	workout.ID = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workout)
}

// GetWorkout handles GET /api/workouts/{workoutID} (or /api/users/{userID}/workouts/{workoutID})
func (h *WorkoutHandler) GetWorkout(w http.ResponseWriter, r *http.Request) {
	workoutIDStr := chi.URLParam(r, "workoutID")
	// TODO: Convert workoutIDStr to int64, handle error
	// TODO: Optionally get userID from path if needed for authorization

	// workout, err := h.workoutModel.Get(r.Context(), workoutID) ...

	w.WriteHeader(http.StatusNotImplemented)
	// Use the variable in the response to fix linter error
	w.Write([]byte("GetWorkout not fully implemented for ID: " + workoutIDStr))

	// --- Old code (needs adapting) ---
	/*
		idStr := r.PathValue("id") // Needs change for chi
		...
	*/
}

// UpdateWorkout handles PUT /api/workouts/{workoutID}
func (h *WorkoutHandler) UpdateWorkout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("UpdateWorkout not implemented"))
}

// DeleteWorkout handles DELETE /api/workouts/{workoutID}
func (h *WorkoutHandler) DeleteWorkout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("DeleteWorkout not implemented"))
}
