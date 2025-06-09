package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	api_models "workout_app_backend/internal/generated"
	models "workout_app_backend/internal/models"

	"github.com/go-chi/chi/v5"
)

type ExerciseHandler struct {
	exerciseModel *models.ExerciseModel
	workoutModel  *models.WorkoutModel
}

func GetExerciseHandlerInstance(exerciseModel *models.ExerciseModel, workoutModel *models.WorkoutModel) *ExerciseHandler {
	return &ExerciseHandler{exerciseModel: exerciseModel, workoutModel: workoutModel}
}

// CreateExercise handles POST /api/users/{userID}/workouts/{workoutID}/cardioExercises
func (h *ExerciseHandler) CreateCardioExercises(w http.ResponseWriter, r *http.Request) {
	handlerLogger.Println("CreateExercise request received") //! Logging the request.
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

	workoutIDStr := chi.URLParam(r, "workoutID")
	workoutID, err := strconv.ParseInt(workoutIDStr, 10, 64)
	if err != nil || workoutID <= 0 {
		respondWithError(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	// Read the request body once
	var exerciseData api_models.CreateCardioExerciseRequest
	if err := json.NewDecoder(r.Body).Decode(&exerciseData); err != nil {
		respondWithError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

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

	if workout.UserID != userID {
		respondWithError(w, "Workout does not belong to user", http.StatusForbidden)
		return
	}

	// Convert to internal domain model
	cardioExercise := &models.CardioExercise{
		BaseExercise: models.BaseExercise{
			WorkoutID: workoutID,
			Name:      exerciseData.GetName(),
			Type:      models.ExerciseTypeCardio,
			Notes:     exerciseData.GetNotes(),
		},
		Distance: float64(exerciseData.GetDistance()),
		Duration: int(exerciseData.GetDuration()),
	}

	id, err := h.exerciseModel.CreateCardio(ctx, cardioExercise)
	if err != nil {
		respondWithError(w, "Failed to create cardio exercise", http.StatusInternalServerError)
		return
	}

	// Get the created exercise to return complete data
	createdExercise, err := h.exerciseModel.Get(ctx, id)
	if err != nil {
		respondWithError(w, "Failed to get created exercise", http.StatusInternalServerError)
		return
	}

	createdCardio := createdExercise.(*models.CardioExercise)

	response := api_models.CardioExerciseResponse{
		Id:        id,
		UserId:    userID,
		WorkoutId: workoutID,
		Name:      createdCardio.Name,
		Notes:     &createdCardio.BaseExercise.Notes,
		Distance:  float32(createdCardio.Distance),
		Duration:  int32(createdCardio.Duration),
		CreatedAt: createdCardio.BaseExercise.CreatedAt,
	}

	respondWithJSON(w, http.StatusCreated, response)
}

// CreateExercise handles POST /api/users/{userID}/workouts/{workoutID}/strengthExercises
func (h *ExerciseHandler) CreateStrengthExercises(w http.ResponseWriter, r *http.Request) {
	handlerLogger.Println("CreateExercise request received") //! Logging the request.
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

	workoutIDStr := chi.URLParam(r, "workoutID")
	workoutID, err := strconv.ParseInt(workoutIDStr, 10, 64)
	if err != nil || workoutID <= 0 {
		respondWithError(w, "Invalid workout ID", http.StatusBadRequest)
		return
	}

	// Read the request body once
	var exerciseData api_models.CreateStrengthExerciseRequest
	if err := json.NewDecoder(r.Body).Decode(&exerciseData); err != nil {
		respondWithError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

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

	if workout.UserID != userID {
		respondWithError(w, "Workout does not belong to user", http.StatusForbidden)
		return
	}

	handlerLogger.Println("Creating Weight Exercise")
	// Convert to internal domain model
	weightExercise := &models.WeightExercise{
		BaseExercise: models.BaseExercise{
			WorkoutID: workoutID,
			Name:      exerciseData.GetName(),
			Type:      models.ExerciseTypeWeights,
			Notes:     exerciseData.GetNotes(),
		},
		Sets:   1, // Default to 1 set
		Reps:   int(exerciseData.GetReps()),
		Weight: float64(exerciseData.GetWeight()),
	}

	id, err := h.exerciseModel.CreateWeights(ctx, weightExercise)
	if err != nil {
		respondWithError(w, "Failed to create weight exercise", http.StatusInternalServerError)
		return
	}

	// Get the created exercise to return complete data
	createdExercise, err := h.exerciseModel.Get(ctx, id)
	if err != nil {
		respondWithError(w, "Failed to get created exercise", http.StatusInternalServerError)
		return
	}

	createdWeight := createdExercise.(*models.WeightExercise)

	response := api_models.StrengthExerciseResponse{
		Id:        id,
		UserId:    userID,
		WorkoutId: workoutID,
		Name:      createdWeight.Name,
		Notes:     &createdWeight.BaseExercise.Notes,
		Reps:      int32(createdWeight.Reps),
		Weight:    float32(createdWeight.Weight),
		CreatedAt: createdWeight.BaseExercise.CreatedAt,
	}

	respondWithJSON(w, http.StatusCreated, response)
}

//! Have to implmeneted below methods later on:

/*
// GetExercise handles GET /api/users/{userID}/workouts/{workoutID}/exercises/{exerciseID}
func (h *ExerciseHandler) GetExercise(w http.ResponseWriter, r *http.Request) {
	handlerLogger.Println("GetExercise request received") //! Logging the request.
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

	// Convert to API model
	response := convertExerciseToAPI(exercise)
	respondWithJSON(w, http.StatusOK, response)
}


// ListExercisesByWorkout handles GET /api/users/{userID}/workouts/{workoutID}/exercises
func (h *ExerciseHandler) ListExercisesByWorkout(w http.ResponseWriter, r *http.Request) {
	handlerLogger.Println("ListExercisesByWorkout request received") //! Logging the request.
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

	// Convert to API models and wrap in response
	response := api_models.ListExercisesResponse{
		Exercises: convertExercisesToAPI(exercises),
	}

	respondWithJSON(w, http.StatusOK, response)
}

// UpdateExercise handles PUT /api/users/{userID}/workouts/{workoutID}/exercises/{exerciseID}
func (h *ExerciseHandler) UpdateExercise(w http.ResponseWriter, r *http.Request) {
	handlerLogger.Println("UpdateExercise request received") //! Logging the request.
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
		var request api_models.CardioExercise
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			respondWithError(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Convert to internal domain model
		cardioExercise := &models.CardioExercise{
			BaseExercise: models.BaseExercise{
				ID:        exerciseID,
				WorkoutID: workoutID,
				Name:      request.Name,
				Type:      models.ExerciseTypeCardio,
			},
			Distance: float64(request.Distance),
			Duration: int(request.Duration),
		}
		if request.Notes != nil {
			cardioExercise.Notes = *request.Notes
		}

		if err := h.exerciseModel.UpdateCardio(ctx, cardioExercise); err != nil {
			respondWithError(w, "Failed to update exercise", http.StatusInternalServerError)
			return
		}

	case *models.WeightExercise:
		var request api_models.WeightExercise
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			respondWithError(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Convert to internal domain model
		weightExercise := &models.WeightExercise{
			BaseExercise: models.BaseExercise{
				ID:        exerciseID,
				WorkoutID: workoutID,
				Name:      request.Name,
				Type:      models.ExerciseTypeWeights,
			},
			Sets:   int(request.Sets),
			Reps:   int(request.Reps),
			Weight: float64(request.Weight),
		}
		if request.Notes != nil {
			weightExercise.Notes = *request.Notes
		}

		if err := h.exerciseModel.UpdateWeights(ctx, weightExercise); err != nil {
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

	// Convert to API model
	response := convertExerciseToAPI(updatedExercise)
	respondWithJSON(w, http.StatusOK, response)
}

// DeleteExercise handles DELETE /api/users/{userID}/workouts/{workoutID}/exercises/{exerciseID}
func (h *ExerciseHandler) DeleteExercise(w http.ResponseWriter, r *http.Request) {
	handlerLogger.Println("DeleteExercise request received") //! Logging the request.
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
*/
