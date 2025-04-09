package middleware

import (
	"net/http"

	"github.com/go-chi/cors"
)

// CorsConfigured returns a configured CORS handler.
func CorsConfigured() func(http.Handler) http.Handler {
	// Configure CORS options
	corsMiddleware := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin domains
		AllowedOrigins: []string{"http://localhost:3000", "http://127.0.0.1:3000"}, // Allow frontend dev server
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true }, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	return corsMiddleware.Handler
}
