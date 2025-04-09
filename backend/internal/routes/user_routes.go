package routes

import (
	"workout_app_backend/internal/handlers"

	"github.com/go-chi/chi/v5"
)

// RegisterUserRoutes sets up the REST-compliant routes for user-related actions.
func RegisterUserRoutes(r chi.Router, userHandler *handlers.UserHandler) {
	r.Route("/users", func(r chi.Router) {
		// List all users
		r.Get("/", userHandler.ListUsers) // GET /api/users

		// Create a new user
		r.Post("/", userHandler.CreateUser) // POST /api/users

		// Routes for a specific user
		r.Route("/{userID}", func(r chi.Router) {
			r.Get("/", userHandler.GetUser)       // GET /api/users/{userID}
			r.Put("/", userHandler.UpdateUser)    // PUT /api/users/{userID}
			r.Delete("/", userHandler.DeleteUser) // DELETE /api/users/{userID}
		})
	})
}
