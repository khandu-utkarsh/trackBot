package routes

import (
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

	routesLogger.Println("Setting up router...") //! Logging the request.

	r := chi.NewRouter()

	// Add debug middleware to log all incoming requests
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			routesLogger.Printf("INCOMING REQUEST: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr) //! Logging the request.
			next.ServeHTTP(w, r)
			routesLogger.Printf("COMPLETED REQUEST: %s %s", r.Method, r.URL.Path) //! Logging the request.
		})
	})

	// Public middleware
	r.Use(middleware.CorsConfigured())
	r.Use(middleware.Recovery())

	//!Additional Middleware
	r.Use(chimiddleware.Timeout(60 * time.Second))
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.Compress(5))

	// Health check (public)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		routesLogger.Println("Health check endpoint hit") //! Logging the request.
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Setup auth middleware
	authMiddleware := middleware.NewAuthMiddleware()

	r.Route("/api", func(r chi.Router) {
		routesLogger.Println("🔶 Setting up /api routes") //! Logging the request.

		routesLogger.Println("Setting up auth routes") //! Logging the request.
		// Unprotected routes first
		r.Post("/auth/google", func(w http.ResponseWriter, r *http.Request) {
			authHandler.GoogleLogin(w, r)
		})
		r.Post("/auth/logout", func(w http.ResponseWriter, r *http.Request) {
			authHandler.Logout(w, r)
		})

		// Protected routes
		r.Group(func(r chi.Router) {
			routesLogger.Println("🔒 Setting up protected routes group") //! Logging the request.
			r.Use(authMiddleware.ValidateJWT())                         // Protect only this group

			// Add middleware to log protected route access
			r.Use(func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					routesLogger.Printf("🔒 PROTECTED ROUTE ACCESS: %s %s", r.Method, r.URL.Path) //! Logging the request.
					next.ServeHTTP(w, r)
				})
			})

			// Auth me endpoint (protected)
			r.Get("/auth/me", func(w http.ResponseWriter, r *http.Request) {
				routesLogger.Println("🔑 Auth me endpoint hit") //! Logging the request.
				authHandler.Me(w, r)
			})

			RegisterUserRoutes(r, userHandler, workoutHandler, exerciseHandler, conversationHandler, messageHandler)
		})
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		routesLogger.Printf("Route not found: %s %s", r.Method, r.URL.Path) //! Logging the request.
		http.Error(w, "Route not found", http.StatusNotFound)
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		routesLogger.Printf("Method not allowed: %s %s", r.Method, r.URL.Path) //! Logging the request.
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	})

	chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		routesLogger.Printf("[%s] %s", method, route) //! Logging the request.
		return nil
	})

	routesLogger.Println("Router setup completed successfully") //! Logging the request.
	return r
}
