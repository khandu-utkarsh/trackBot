package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"
	handlers "workout_app_backend/internal/handlers"
	middleware "workout_app_backend/internal/middleware"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(userHandler *handlers.UserHandler,
	workoutHandler *handlers.WorkoutHandler,
	exerciseHandler *handlers.ExerciseHandler,
	conversationHandler *handlers.ConversationHandler,
	messageHandler *handlers.MessageHandler,
	authHandler *handlers.AuthHandler) *chi.Mux {

	log.Println("Setting up router...")

	r := chi.NewRouter()

	// Add debug middleware to log all incoming requests
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("🔵 INCOMING REQUEST: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
			next.ServeHTTP(w, r)
			log.Printf("🟢 COMPLETED REQUEST: %s %s", r.Method, r.URL.Path)
		})
	})

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
		log.Println("🟡 Health check endpoint hit")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Setup auth middleware
	authMiddleware := middleware.NewAuthMiddleware()

	r.Route("/api", func(r chi.Router) {
		log.Println("🔶 Setting up /api routes")

		fmt.Println("Setting up auth routes")
		// Unprotected routes first
		r.Post("/auth/google", func(w http.ResponseWriter, r *http.Request) {
			log.Println("🔑 Google auth endpoint hit")
			authHandler.GoogleLogin(w, r)
		})
		r.Post("/auth/logout", func(w http.ResponseWriter, r *http.Request) {
			log.Println("🔑 Logout endpoint hit")
			authHandler.Logout(w, r)
		})
		r.Get("/auth/me", func(w http.ResponseWriter, r *http.Request) {
			log.Println("🔑 Auth me endpoint hit")
			authHandler.Me(w, r)
		})

		// Protected routes
		r.Group(func(r chi.Router) {
			log.Println("🔒 Setting up protected routes group")
			r.Use(authMiddleware.ValidateJWT()) // Protect only this group

			// Add middleware to log protected route access
			r.Use(func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					log.Printf("🔒 PROTECTED ROUTE ACCESS: %s %s", r.Method, r.URL.Path)
					next.ServeHTTP(w, r)
				})
			})

			RegisterUserRoutes(r, userHandler, workoutHandler, exerciseHandler, conversationHandler, messageHandler)
		})
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("❌ Route not found: %s %s", r.Method, r.URL.Path)
		http.Error(w, "Route not found", http.StatusNotFound)
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("❌ Method not allowed: %s %s", r.Method, r.URL.Path)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("[%s] %s", method, route)
		return nil
	})

	log.Println("✅ Router setup completed successfully")
	return r
}
