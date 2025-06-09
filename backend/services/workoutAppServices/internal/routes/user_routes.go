package routes

import (
	handlers "workout_app_backend/internal/handlers"

	"github.com/go-chi/chi/v5"
)

// RegisterUserRoutes sets up the REST-compliant routes for user-related actions.
func RegisterUserRoutes(r chi.Router, userHandler *handlers.UserHandler, workoutHandler *handlers.WorkoutHandler, exerciseHandler *handlers.ExerciseHandler, conversationHandler *handlers.ConversationHandler, messageHandler *handlers.MessageHandler) {
	r.Route("/users", func(r chi.Router) {
		// Routes for a specific user
		r.Route("/{userID}", func(r chi.Router) {
			r.Get("/email", userHandler.GetUserByEmail) // GET /api/users/{userID}/email

			// Register workout routes under user
			RegisterWorkoutRoutes(r, workoutHandler)
			// Register exercise routes under user
			RegisterExerciseRoutes(r, exerciseHandler)
			// Register conversation routes under user
			RegisterConversationRoutes(r, conversationHandler, messageHandler)
		})
	})
}
