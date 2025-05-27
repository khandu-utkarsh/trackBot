package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	models "workout_app_backend/internal/models"

	"github.com/go-chi/chi/v5"
	// Needed for URL params
)

//!Functions needed
//CreateExercise
//GetExercise
//ListExercisesByWorkout
//UpdateExercise
//DeleteExercise

type ExerciseHandler struct {
	exerciseModel *models.ExerciseModel
	workoutModel  *models.WorkoutModel
}

func GetExerciseHandlerInstance(exerciseModel *models.ExerciseModel, workoutModel *models.WorkoutModel) *ExerciseHandler {
	return &ExerciseHandler{exerciseModel: exerciseModel, workoutModel: workoutModel}
}

// CreateExercise handles POST /api/users/{userID}/workouts/{workoutID}/exercises
func (h *ExerciseHandler) CreateExercise(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Creating Exercise Handler 1")

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	fmt.Println("Creating Exercise Handler 2")

	workoutIDStr := chi.URLParam(r, "workoutID")
	workoutID, err := strconv.ParseInt(workoutIDStr, 10, 64)
	if err != nil || workoutID <= 0 {
		respondWithError(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	fmt.Println("Creating Exercise Handler 3")

	// Read the request body once
	var exerciseData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&exerciseData); err != nil {
		respondWithError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println("Creating Exercise Handler 4")

	// Get the exercise type
	exerciseType, ok := exerciseData["type"].(string)
	if !ok {
		respondWithError(w, "Invalid exercise type", http.StatusBadRequest)
		return
	}

	fmt.Println("Creating Exercise Handler 5")

	// Verify workout exists and belongs to user before creating exercise
	ctx := r.Context()
	workout, err := h.workoutModel.Get(ctx, workoutID)
	if err != nil {
		if errors.Is(err, models.ErrWorkoutNotFound) {
			respondWithError(w, "Workout not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to verify workout", http.StatusInternalServerError)
		return
	}

	fmt.Println("Creating Exercise Handler 6")

	if workout.UserID != userID {
		respondWithError(w, "Workout does not belong to user", http.StatusForbidden)
		return
	}

	var id int64
	switch exerciseType {
	case "cardio":

		fmt.Println("Creating Exercise Handler 7")

		var cardioExercise models.CardioExercise
		// Convert the map back to JSON and then decode into CardioExercise
		jsonData, err := json.Marshal(exerciseData)
		if err != nil {
			respondWithError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(jsonData, &cardioExercise); err != nil {
			respondWithError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		cardioExercise.WorkoutID = workoutID
		fmt.Println(" Creating the Cardio Exercise: ", cardioExercise)
		id, _ = h.exerciseModel.CreateCardio(ctx, &cardioExercise)
	case "weights":
		fmt.Println("Creating Exercise Handler 8")
		var weightExercise models.WeightExercise
		// Convert the map back to JSON and then decode into WeightExercise
		jsonData, err := json.Marshal(exerciseData)
		if err != nil {
			respondWithError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(jsonData, &weightExercise); err != nil {
			respondWithError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		weightExercise.WorkoutID = workoutID
		fmt.Println(" Creating the Weight Exercise: ", weightExercise)
		id, _ = h.exerciseModel.CreateWeights(ctx, &weightExercise)

		fmt.Println("Creating Exercise Handler 9.1, created exercise id: ", id)
	default:
		fmt.Println("Creating Exercise Handler 9, Invalid exercise type: ", exerciseType)

		respondWithError(w, "Invalid exercise type", http.StatusBadRequest)
		return
	}

	fmt.Println("Creating Exercise Handler 10")

	// Get the created exercise to return complete data
	createdExercise, err := h.exerciseModel.Get(ctx, id)
	if err != nil {
		respondWithError(w, "Failed to get created exercise", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, createdExercise)
}

// GetExercise handles GET /api/users/{userID}/workouts/{workoutID}/exercises/{exerciseID}
func (h *ExerciseHandler) GetExercise(w http.ResponseWriter, r *http.Request) {
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

	workoutIDStr := chi.URLParam(r, "workoutID")
	workoutID, err := strconv.ParseInt(workoutIDStr, 10, 64)
	if err != nil || workoutID <= 0 {
		respondWithError(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	exerciseIDStr := chi.URLParam(r, "exerciseID")
	exerciseID, err := strconv.ParseInt(exerciseIDStr, 10, 64)
	if err != nil || exerciseID <= 0 {
		respondWithError(w, "Invalid exercise ID", http.StatusBadRequest)
		return
	}

	// Verify workout exists and belongs to user
	ctx := r.Context()
	workout, err := h.workoutModel.Get(ctx, workoutID)
	if err != nil {
		if errors.Is(err, models.ErrWorkoutNotFound) {
			respondWithError(w, "Workout not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to verify workout", http.StatusInternalServerError)
		return
	}

	if workout.UserID != userID {
		respondWithError(w, "Workout does not belong to user", http.StatusForbidden)
		return
	}

	exercise, err := h.exerciseModel.Get(ctx, exerciseID)
	if err != nil {
		if errors.Is(err, models.ErrExerciseNotFound) {
			respondWithError(w, "Exercise not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to get exercise", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, exercise)
}

// ListExercisesByWorkout handles GET /api/users/{userID}/workouts/{workoutID}/exercises
func (h *ExerciseHandler) ListExercisesByWorkout(w http.ResponseWriter, r *http.Request) {
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

	workoutIDStr := chi.URLParam(r, "workoutID")
	workoutID, err := strconv.ParseInt(workoutIDStr, 10, 64)
	if err != nil || workoutID <= 0 {
		respondWithError(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	// Verify workout exists and belongs to user
	ctx := r.Context()
	workout, err := h.workoutModel.Get(ctx, workoutID)
	if err != nil {
		if errors.Is(err, models.ErrWorkoutNotFound) {
			respondWithError(w, "Workout not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to verify workout", http.StatusInternalServerError)
		return
	}

	if workout.UserID != userID {
		respondWithError(w, "Workout does not belong to user", http.StatusForbidden)
		return
	}

	exercises, err := h.exerciseModel.ListByWorkout(ctx, workoutID)
	if err != nil {
		respondWithError(w, "Failed to list exercises", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, exercises)
}

// UpdateExercise handles PUT /api/users/{userID}/workouts/{workoutID}/exercises/{exerciseID}
func (h *ExerciseHandler) UpdateExercise(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	workoutIDStr := chi.URLParam(r, "workoutID")
	workoutID, err := strconv.ParseInt(workoutIDStr, 10, 64)
	if err != nil || workoutID <= 0 {
		respondWithError(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	exerciseIDStr := chi.URLParam(r, "exerciseID")
	exerciseID, err := strconv.ParseInt(exerciseIDStr, 10, 64)
	if err != nil || exerciseID <= 0 {
		respondWithError(w, "Invalid exercise ID", http.StatusBadRequest)
		return
	}

	// Verify workout exists and belongs to user
	ctx := r.Context()
	workout, err := h.workoutModel.Get(ctx, workoutID)
	if err != nil {
		if errors.Is(err, models.ErrWorkoutNotFound) {
			respondWithError(w, "Workout not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to verify workout", http.StatusInternalServerError)
		return
	}

	if workout.UserID != userID {
		respondWithError(w, "Workout does not belong to user", http.StatusForbidden)
		return
	}

	// Get existing exercise to determine its type
	existingExercise, err := h.exerciseModel.Get(ctx, exerciseID)
	if err != nil {
		if errors.Is(err, models.ErrExerciseNotFound) {
			respondWithError(w, "Exercise not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to get exercise", http.StatusInternalServerError)
		return
	}

	// Update based on exercise type
	switch existingExercise.(type) {
	case *models.CardioExercise:
		var cardioExercise models.CardioExercise
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&cardioExercise); err != nil {
			respondWithError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		cardioExercise.ID = exerciseID
		cardioExercise.WorkoutID = workoutID
		if err := h.exerciseModel.UpdateCardio(ctx, &cardioExercise); err != nil {
			respondWithError(w, "Failed to update exercise", http.StatusInternalServerError)
			return
		}
	case *models.WeightExercise:
		var weightExercise models.WeightExercise
		if err := json.NewDecoder(r.Body).Decode(&weightExercise); err != nil {
			respondWithError(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		weightExercise.ID = exerciseID
		weightExercise.WorkoutID = workoutID
		if err := h.exerciseModel.UpdateWeights(ctx, &weightExercise); err != nil {
			respondWithError(w, "Failed to update exercise", http.StatusInternalServerError)
			return
		}
	default:
		respondWithError(w, "Invalid exercise type", http.StatusBadRequest)
		return
	}

	// Get the updated exercise to return complete data
	updatedExercise, err := h.exerciseModel.Get(ctx, exerciseID)
	if err != nil {
		respondWithError(w, "Failed to get updated exercise", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedExercise)
}

// DeleteExercise handles DELETE /api/users/{userID}/workouts/{workoutID}/exercises/{exerciseID}
func (h *ExerciseHandler) DeleteExercise(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	workoutIDStr := chi.URLParam(r, "workoutID")
	workoutID, err := strconv.ParseInt(workoutIDStr, 10, 64)
	if err != nil || workoutID <= 0 {
		respondWithError(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	exerciseIDStr := chi.URLParam(r, "exerciseID")
	exerciseID, err := strconv.ParseInt(exerciseIDStr, 10, 64)
	if err != nil {
		respondWithError(w, "Invalid exercise ID", http.StatusBadRequest)
		return
	}

	// Verify workout exists and belongs to user
	ctx := r.Context()
	workout, err := h.workoutModel.Get(ctx, workoutID)
	if err != nil {
		if errors.Is(err, models.ErrWorkoutNotFound) {
			respondWithError(w, "Workout not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to verify workout", http.StatusInternalServerError)
		return
	}

	if workout.UserID != userID {
		respondWithError(w, "Workout does not belong to user", http.StatusForbidden)
		return
	}

	if err := h.exerciseModel.Delete(ctx, exerciseID); err != nil {
		if errors.Is(err, models.ErrExerciseNotFound) {
			respondWithError(w, "Exercise not found", http.StatusNotFound)
			return
		} else if errors.Is(err, models.ErrInvalidInput) {
			respondWithError(w, "Invalid exercise ID", http.StatusBadRequest)
			return
		}
		respondWithError(w, "Failed to delete exercise", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]int64{"deleted_id": exerciseID})
}
