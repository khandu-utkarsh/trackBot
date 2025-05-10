package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

// Logger returns the standard Chi request logger middleware.
// This provides a consistent way to access the logger middleware
// and allows for easier customization in the future if needed.
func Logger() func(http.Handler) http.Handler {
	return middleware.Logger
}
