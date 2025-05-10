package routes

import (
	"workout_app_backend/services/workoutAppServices/internal/handlers"

	"github.com/go-chi/chi/v5"
)

// RegisterWorkoutRoutes sets up the RESTful routes for workout-related actions.
func RegisterWorkoutRoutes(r chi.Router, workoutHandler *handlers.WorkoutHandler) {
	r.Route("/users/{userID}", func(r chi.Router) {
		r.Route("/workouts", func(r chi.Router) {
			// List all workouts (optionally filter by ?user_id=)
			r.Get("/", workoutHandler.ListWorkouts) // GET /api/users/{userID}/workouts

			// Create a new workout
			r.Post("/", workoutHandler.CreateWorkout) // POST /api/users/{userID}/workouts

			// Routes for a specific workout
			r.Route("/{workoutID}", func(r chi.Router) {
				r.Get("/", workoutHandler.GetWorkout)       // GET /api/users/{userID}/workouts/{workoutID}
				r.Put("/", workoutHandler.UpdateWorkout)    // PUT /api/users/{userID}/workouts/{workoutID}
				r.Delete("/", workoutHandler.DeleteWorkout) // DELETE /api/users/{userID}/workouts/{workoutID}
			})
		})
	})
}
