package handlers

import (
	"net/http"

	// "strconv" // No longer directly used here, potentially needed for param conversion
	"workout_app_backend/internal/models"
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
}

func GetExerciseHandlerInstance(exerciseModel *models.ExerciseModel) *ExerciseHandler {
	return &ExerciseHandler{exerciseModel: exerciseModel}
}

func (h *ExerciseHandler) CreateExercise(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("CreateExercise not implemented"))
}

func (h *ExerciseHandler) GetExercise(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("GetExercise not implemented"))
}

func (h *ExerciseHandler) ListExercisesByWorkout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("ListExercisesByWorkout not implemented"))
}

func (h *ExerciseHandler) UpdateExercise(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("UpdateExercise not implemented"))
}

// DeleteExercise handles DELETE /api/exercises/{exerciseID} (or nested)
func (h *ExerciseHandler) DeleteExercise(w http.ResponseWriter, r *http.Request) {
	// TODO: Get exerciseID from path param
	// TODO: Implement delete logic
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("DeleteExercise not implemented"))
}
