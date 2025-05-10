package routes

import (
	"workout_app_backend/services/workoutAppServices/internal/handlers"

	"github.com/go-chi/chi/v5"
)

// RegisterExerciseRoutes registers exercise routes under a given workout ID.
func RegisterExerciseRoutes(r chi.Router, exerciseHandler *handlers.ExerciseHandler) {
	r.Route("/users/{userID}", func(r chi.Router) {
		r.Route("/workouts/{workoutID}/exercises", func(r chi.Router) {
			// List all exercises under a workout
			r.Get("/", exerciseHandler.ListExercisesByWorkout)

			// Create a new exercise under a workout
			r.Post("/", exerciseHandler.CreateExercise)

			// Actions on a specific exercise under a workout
			r.Route("/{exerciseID}", func(r chi.Router) {
				r.Get("/", exerciseHandler.GetExercise)
				r.Put("/", exerciseHandler.UpdateExercise)
				r.Delete("/", exerciseHandler.DeleteExercise)
			})
		})
	})
}
