package routes

import (
	"workout_app_backend/internal/handlers"

	"github.com/go-chi/chi/v5"
)

// RegisterWorkoutRoutes sets up the RESTful routes for workout-related actions.
func RegisterWorkoutRoutes(r chi.Router, workoutHandler *handlers.WorkoutHandler) {
	r.Route("/workouts", func(r chi.Router) {
		// List all workouts (optionally filter by ?user_id=)
		r.Get("/", workoutHandler.ListWorkouts) // GET /api/workouts

		// Create a new workout
		r.Post("/", workoutHandler.CreateWorkout) // POST /api/workouts

		// Routes for a specific workout
		r.Route("/{workoutID}", func(r chi.Router) {
			r.Get("/", workoutHandler.GetWorkout)       // GET /api/workouts/{workoutID}
			r.Put("/", workoutHandler.UpdateWorkout)    // PUT /api/workouts/{workoutID}
			r.Delete("/", workoutHandler.DeleteWorkout) // DELETE /api/workouts/{workoutID}
		})
	})
}
