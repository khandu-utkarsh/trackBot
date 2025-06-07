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

// CreateUser handles POST /api/users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	logRequest("CreateUser")

	if err := validateHTTPMethod(r, http.MethodPost); err != nil {
		handleHTTPError(w, err)
		return
	}

	// 1. Parse request using generated API model
	var createUserReq api_models.CreateUserRequest
	if err := decodeJSONBody(r, &createUserReq); err != nil {
		handleHTTPError(w, err)
		return
	}

	// 2. Convert API model to internal domain model
	user := &models.User{
		Email: createUserReq.Email,
	}

	// 3. Validate using internal model
	if err := validateUserInput(user); err != nil {
		respondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 4. Business logic using internal models
	ctx := r.Context()
	id, err := h.userModel.Create(ctx, user)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			// Get existing user and return 200 (as per OpenAPI spec)
			existingUser, getErr := h.userModel.GetByEmail(ctx, createUserReq.Email)
			if getErr != nil {
				respondWithError(w, "Failed to get existing user", http.StatusInternalServerError)
				return
			}

			// Convert to full User schema for 200 response
			userResponse := &api_models.User{
				Id:    existingUser.ID,
				Email: existingUser.Email,
			}
			userResponse.SetCreatedAt(existingUser.CreatedAt)

			respondWithJSON(w, http.StatusOK, userResponse)
			return
		}
		respondWithError(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// 5. Convert to API response model (201 Created)
	response := &api_models.CreateUserResponse{
		Id:    id,
		Email: createUserReq.Email,
	}

	// 6. Return API response
	respondWithJSON(w, http.StatusCreated, response)
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
