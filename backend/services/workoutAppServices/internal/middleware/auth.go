package middleware

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// UserContext represents the authenticated user
type UserContext struct {
	UserID string `json:"sub"`
	Email  string `json:"email"`
	Name   string `json:"name"`
}

// AuthMiddleware handles authentication
type AuthMiddleware struct{}

// NewAuthMiddleware creates a new auth middleware instance
func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

// GoogleJWTClaims represents the claims in a Google JWT token
type GoogleJWTClaims struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Iss           string `json:"iss"`
	Aud           string `json:"aud"`
	Iat           int64  `json:"iat"`
	Exp           int64  `json:"exp"`
}

// ValidateJWT middleware - validates Google JWT tokens
func (a *AuthMiddleware) ValidateJWT() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check for Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header required", http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				http.Error(w, "Bearer token required", http.StatusUnauthorized)
				return
			}

			// Validate Google JWT token
			claims, err := a.validateGoogleJWT(tokenString)
			if err != nil {
				http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
				return
			}

			// Create user context from Google claims
			userCtx := &UserContext{
				UserID: claims.Sub,
				Email:  claims.Email,
				Name:   claims.Name,
			}

			// Add user context to request
			ctx := context.WithValue(r.Context(), "user", userCtx)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// validateGoogleJWT validates a Google JWT token by parsing its payload
func (a *AuthMiddleware) validateGoogleJWT(tokenString string) (*GoogleJWTClaims, error) {
	// Split the JWT token into parts
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid JWT format")
	}

	// Decode the payload (second part)
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("failed to decode JWT payload: %v", err)
	}

	// Parse the claims
	var claims GoogleJWTClaims
	if err := json.Unmarshal(payload, &claims); err != nil {
		return nil, fmt.Errorf("failed to parse JWT claims: %v", err)
	}

	// Basic validation
	if claims.Iss != "accounts.google.com" && claims.Iss != "https://accounts.google.com" {
		return nil, errors.New("invalid issuer")
	}

	// Check if token is expired
	now := time.Now().Unix()
	if claims.Exp > 0 && claims.Exp < now {
		return nil, errors.New("token expired")
	}

	// Check if token is not yet valid
	if claims.Iat > 0 && claims.Iat > now+300 { // Allow 5 minutes clock skew
		return nil, errors.New("token not yet valid")
	}

	return &claims, nil
}

// GetUserFromContext extracts user context from request context
func GetUserFromContext(ctx context.Context) (*UserContext, bool) {
	user, ok := ctx.Value("user").(*UserContext)
	return user, ok
}
