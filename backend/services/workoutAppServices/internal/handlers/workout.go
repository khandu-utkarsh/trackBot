package handlers

import (
	"errors"
	"net/http"
	api_models "workout_app_backend/internal/generated"
	models "workout_app_backend/internal/models"
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
	logRequest("ListWorkouts")

	if err := validateHTTPMethod(r, http.MethodGet); err != nil {
		handleHTTPError(w, err)
		return
	}

	userID, err := parseIDFromURL(r, "userID")
	if err != nil {
		handleHTTPError(w, err)
		return
	}

	// Extract optional query parameters
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

	// Convert to API response model
	apiWorkouts := make([]api_models.Workout, len(workouts))
	for i, workout := range workouts {
		apiWorkouts[i] = api_models.Workout{
			Id:     workout.ID,
			UserId: workout.UserID,
		}
		apiWorkouts[i].SetCreatedAt(workout.CreatedAt)
		apiWorkouts[i].SetUpdatedAt(workout.UpdatedAt)
	}

	response := &api_models.ListWorkoutsResponse{
		Workouts: apiWorkouts,
	}

	respondWithJSON(w, http.StatusOK, response)
}

// CreateWorkout handles POST /api/users/{userID}/workouts
func (h *WorkoutHandler) CreateWorkout(w http.ResponseWriter, r *http.Request) {
	logRequest("CreateWorkout")

	if err := validateHTTPMethod(r, http.MethodPost); err != nil {
		handleHTTPError(w, err)
		return
	}

	userID, err := parseIDFromURL(r, "userID")
	if err != nil {
		handleHTTPError(w, err)
		return
	}

	// 1. Parse request using generated API model
	var createWorkoutReq api_models.CreateWorkoutRequest
	if err := decodeJSONBody(r, &createWorkoutReq); err != nil {
		handleHTTPError(w, err)
		return
	}

	// 2. Convert API model to internal domain model
	workout := &models.Workout{
		UserID: userID, // Set from URL path
	}

	// 3. Verify user exists before creating workout
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

	// 4. Business logic using internal models
	id, err := h.workoutModel.Create(ctx, workout)
	if err != nil {
		respondWithError(w, "Failed to create workout", http.StatusInternalServerError)
		return
	}

	// 5. Convert to API response model
	response := &api_models.CreateWorkoutResponse{
		Id: id,
	}

	// 6. Return API response
	respondWithJSON(w, http.StatusCreated, response)
}

// GetWorkout handles GET /api/users/{userID}/workouts/{workoutID}
func (h *WorkoutHandler) GetWorkout(w http.ResponseWriter, r *http.Request) {
	logRequest("GetWorkout")

	if err := validateHTTPMethod(r, http.MethodGet); err != nil {
		handleHTTPError(w, err)
		return
	}

	userID, workoutID, err := parseUserAndWorkoutIDs(r)
	if err != nil {
		handleHTTPError(w, err)
		return
	}

	ctx := r.Context()

	// Verify user exists
	_, err = h.userModel.Get(ctx, userID)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			respondWithError(w, "User not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to verify user", http.StatusInternalServerError)
		return
	}

	workout, err := h.workoutModel.Get(ctx, workoutID)
	if err != nil {
		if errors.Is(err, models.ErrWorkoutNotFound) {
			respondWithError(w, "Workout not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to get workout", http.StatusInternalServerError)
		return
	}

	// Verify workout belongs to user
	if workout.UserID != userID {
		respondWithError(w, "Workout not found", http.StatusNotFound)
		return
	}

	// Convert to API response model
	workoutResponse := &api_models.Workout{
		Id:     workout.ID,
		UserId: workout.UserID,
	}
	workoutResponse.SetCreatedAt(workout.CreatedAt)
	workoutResponse.SetUpdatedAt(workout.UpdatedAt)

	respondWithJSON(w, http.StatusOK, workoutResponse)
}

// UpdateWorkout handles PUT /api/users/{userID}/workouts/{workoutID}
func (h *WorkoutHandler) UpdateWorkout(w http.ResponseWriter, r *http.Request) {
	logRequest("UpdateWorkout")

	if err := validateHTTPMethod(r, http.MethodPut); err != nil {
		handleHTTPError(w, err)
		return
	}

	userID, workoutID, err := parseUserAndWorkoutIDs(r)
	if err != nil {
		handleHTTPError(w, err)
		return
	}

	// 1. Parse request using generated API model
	var updateWorkoutReq api_models.UpdateWorkoutRequest
	if err := decodeJSONBody(r, &updateWorkoutReq); err != nil {
		handleHTTPError(w, err)
		return
	}

	// 2. Verify user exists
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

	// 3. Convert API model to internal domain model
	workout := &models.Workout{
		ID:     workoutID, // Set from URL path
		UserID: userID,    // Set from URL path
	}

	// Override UserID if provided in request
	if updateWorkoutReq.UserId != nil {
		workout.UserID = *updateWorkoutReq.UserId

		// Verify new user exists if different
		if *updateWorkoutReq.UserId != userID {
			_, err = h.userModel.Get(ctx, *updateWorkoutReq.UserId)
			if err != nil {
				if errors.Is(err, models.ErrUserNotFound) {
					respondWithError(w, "Target user not found", http.StatusNotFound)
					return
				}
				respondWithError(w, "Failed to verify target user", http.StatusInternalServerError)
				return
			}
		}
	}

	// 4. Business logic using internal models
	if err := h.workoutModel.Update(ctx, workout); err != nil {
		if errors.Is(err, models.ErrWorkoutNotFound) {
			respondWithError(w, "Workout not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to update workout", http.StatusInternalServerError)
		return
	}

	// 5. Get updated workout and convert to API response model
	updatedWorkout, err := h.workoutModel.Get(ctx, workoutID)
	if err != nil {
		respondWithError(w, "Failed to get updated workout", http.StatusInternalServerError)
		return
	}

	workoutResponse := &api_models.Workout{
		Id:     updatedWorkout.ID,
		UserId: updatedWorkout.UserID,
	}
	workoutResponse.SetCreatedAt(updatedWorkout.CreatedAt)
	workoutResponse.SetUpdatedAt(updatedWorkout.UpdatedAt)

	respondWithJSON(w, http.StatusOK, workoutResponse)
}

// DeleteWorkout handles DELETE /api/users/{userID}/workouts/{workoutID}
func (h *WorkoutHandler) DeleteWorkout(w http.ResponseWriter, r *http.Request) {
	logRequest("DeleteWorkout")

	if err := validateHTTPMethod(r, http.MethodDelete); err != nil {
		handleHTTPError(w, err)
		return
	}

	userID, workoutID, err := parseUserAndWorkoutIDs(r)
	if err != nil {
		handleHTTPError(w, err)
		return
	}

	ctx := r.Context()

	// Verify user exists
	_, err = h.userModel.Get(ctx, userID)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			respondWithError(w, "User not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to verify user", http.StatusInternalServerError)
		return
	}

	// Verify workout exists and belongs to user
	workout, err := h.workoutModel.Get(ctx, workoutID)
	if err != nil {
		if errors.Is(err, models.ErrWorkoutNotFound) {
			respondWithError(w, "Workout not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to get workout", http.StatusInternalServerError)
		return
	}

	if workout.UserID != userID {
		respondWithError(w, "Workout not found", http.StatusNotFound)
		return
	}

	if err := h.workoutModel.Delete(ctx, workoutID); err != nil {
		respondWithError(w, "Failed to delete workout", http.StatusInternalServerError)
		return
	}

	// Return 204 No Content as per OpenAPI spec
	w.WriteHeader(http.StatusNoContent)
}
