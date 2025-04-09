package routes

import (
	"log"
	"net/http"
	"workout_app_backend/internal/handlers" // Import handlers package
	"workout_app_backend/internal/middleware"

	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(userHandler *handlers.UserHandler, workoutHandler *handlers.WorkoutHandler, exerciseHandler *handlers.ExerciseHandler) *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.CorsConfigured())
	r.Use(middleware.Logger())
	r.Use(chi_middleware.Recoverer)

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Group routes under /api
	r.Route("/api", func(r chi.Router) {
		RegisterUserRoutes(r, userHandler)
		RegisterWorkoutRoutes(r, workoutHandler)
		RegisterExerciseRoutes(r, exerciseHandler)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Route not found", http.StatusNotFound)
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("[%s] %s", method, route)
		return nil
	})

	return r
}
