package routes

import (
	"log"
	"net/http"
	"time"
	handlers "workout_app_backend/internal/handlers"
	middleware "workout_app_backend/internal/middleware"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(userHandler *handlers.UserHandler, workoutHandler *handlers.WorkoutHandler, exerciseHandler *handlers.ExerciseHandler, conversationHandler *handlers.ConversationHandler, messageHandler *handlers.MessageHandler) *chi.Mux {

	log.Println("Setting up router...")

	r := chi.NewRouter()

	// Public middleware
	r.Use(middleware.CorsConfigured())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	//!Additional Middleware
	r.Use(chimiddleware.Timeout(60 * time.Second))
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.Compress(5))

	// Health check (public)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Setup auth middleware
	authMiddleware := middleware.NewAuthMiddleware()

	// Protected routes under /api
	r.Route("/api", func(r chi.Router) {
		r.Use(authMiddleware.ValidateJWT()) // Protect all /api routes
		RegisterUserRoutes(r, userHandler, workoutHandler, exerciseHandler, conversationHandler, messageHandler)
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
