package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"workout_app_backend/internal/middleware"
	"workout_app_backend/internal/models"
	"workout_app_backend/internal/testutils"
)

func TestAuthHandler_GoogleLogin(t *testing.T) {
	// Setup test database
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	// Initialize models
	userModel := models.GetUserModelInstance(db, "users")
	ctx := context.Background()
	err := userModel.Initialize(ctx)
	if err != nil {
		t.Fatalf("Failed to initialize user model: %v", err)
	}

	// Create auth middleware and handler
	authMiddleware := middleware.NewAuthMiddleware()
	authHandler := GetAuthHandlerInstance(authMiddleware, userModel)

	t.Run("missing google token", func(t *testing.T) {
		reqBody := map[string]string{"googleToken": ""}
		jsonBody, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/api/auth/google", bytes.NewReader(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		authHandler.GoogleLogin(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("invalid method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/auth/google", nil)
		w := httptest.NewRecorder()

		authHandler.GoogleLogin(w, req)

		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, w.Code)
		}
	})
}

func TestAuthHandler_Logout(t *testing.T) {
	// Setup test database
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	userModel := models.GetUserModelInstance(db, "users")
	authMiddleware := middleware.NewAuthMiddleware()
	authHandler := GetAuthHandlerInstance(authMiddleware, userModel)

	t.Run("successful logout", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/api/auth/logout", nil)
		w := httptest.NewRecorder()

		authHandler.Logout(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}

		// Check if auth cookie is cleared
		cookies := w.Result().Cookies()
		authCookieCleared := false
		for _, cookie := range cookies {
			if cookie.Name == "auth_token" && cookie.MaxAge == -1 {
				authCookieCleared = true
				break
			}
		}
		if !authCookieCleared {
			t.Error("Expected auth cookie to be cleared")
		}
	})
}

func TestAuthHandler_Me(t *testing.T) {
	// Setup test database
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	userModel := models.GetUserModelInstance(db, "users")
	authMiddleware := middleware.NewAuthMiddleware()
	authHandler := GetAuthHandlerInstance(authMiddleware, userModel)

	t.Run("unauthorized request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/auth/me", nil)
		w := httptest.NewRecorder()

		authHandler.Me(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, w.Code)
		}
	})
}
