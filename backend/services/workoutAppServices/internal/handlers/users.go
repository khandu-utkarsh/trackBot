package handlers

import (
	"errors"
	"net/http"
	"strings"
	api_models "workout_app_backend/internal/generated"
	models "workout_app_backend/internal/models"

	"github.com/go-chi/chi/v5"
)

//!Functions implemented (per OpenAPI spec)
//CreateUser
//GetUserByEmail

// UserHandler handles HTTP requests related to users.
type UserHandler struct {
	userModel *models.UserModel
}

// GetUserHandlerInstance creates a new UserHandler instance.
func GetUserHandlerInstance(userModel *models.UserModel) *UserHandler {
	return &UserHandler{userModel: userModel}
}

func (h *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	logRequest("GetUserByEmail")

	if err := validateHTTPMethod(r, http.MethodGet); err != nil {
		handleHTTPError(w, err)
		return
	}

	// Extract email from URL path parameter
	email := chi.URLParam(r, "email")
	if email == "" {
		respondWithError(w, "Email parameter is required", http.StatusBadRequest)
		return
	}

	// Get user from database
	ctx := r.Context()
	user, err := h.userModel.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			respondWithError(w, "User not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	// Convert to API response model
	userResponse := &api_models.User{
		Id:    user.ID,
		Email: user.Email,
	}
	userResponse.SetCreatedAt(user.CreatedAt)

	respondWithJSON(w, http.StatusOK, userResponse)
}

// Helper functions

func validateUserInput(user *models.User) error {
	if user.Email == "" {
		return errors.New("email is required")
	}
	if !strings.Contains(user.Email, "@") {
		return errors.New("invalid email format")
	}
	return nil
}
