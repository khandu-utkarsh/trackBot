package handlers

import (
	"errors"
	"net/http"
	"strings"
	models "workout_app_backend/internal/models"
)

//!Functions needed
//ListUsers
//CreateUser
//GetUser
//UpdateUser
//DeleteUser

// UserHandler handles HTTP requests related to users.
type UserHandler struct {
	userModel *models.UserModel
}

// GetUserHandlerInstance creates a new UserHandler instance.
func GetUserHandlerInstance(userModel *models.UserModel) *UserHandler {
	return &UserHandler{userModel: userModel}
}

// ListUsers handles GET /api/users
func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	logRequest("ListUsers")

	if err := validateHTTPMethod(r, http.MethodGet); err != nil {
		handleHTTPError(w, err)
		return
	}

	ctx := r.Context()
	users, err := h.userModel.List(ctx)
	if err != nil {
		respondWithError(w, "Failed to list users", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}

// CreateUser handles POST /api/users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	logRequest("CreateUser")

	if err := validateHTTPMethod(r, http.MethodPost); err != nil {
		handleHTTPError(w, err)
		return
	}

	var user models.User
	if err := decodeJSONBody(r, &user); err != nil {
		handleHTTPError(w, err)
		return
	}

	if err := validateUserInput(&user); err != nil {
		respondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	id, err := h.userModel.Create(ctx, &user)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			respondWithError(w, "Email already exists", http.StatusConflict)
			return
		}
		respondWithError(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Get the created user to return complete data
	createdUser, err := h.userModel.Get(ctx, id)
	if err != nil {
		respondWithError(w, "Failed to get created user", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, createdUser)
}

// GetUser handles GET /api/users/{userID}
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	logRequest("GetUser")

	if err := validateHTTPMethod(r, http.MethodGet); err != nil {
		handleHTTPError(w, err)
		return
	}

	userID, err := parseIDFromURL(r, "userID")
	if err != nil {
		handleHTTPError(w, err)
		return
	}

	ctx := r.Context()
	user, err := h.userModel.Get(ctx, userID)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			respondWithError(w, "User not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

// UpdateUser handles PUT /api/users/{userID}
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	logRequest("UpdateUser")

	if err := validateHTTPMethod(r, http.MethodPut); err != nil {
		handleHTTPError(w, err)
		return
	}

	userID, err := parseIDFromURL(r, "userID")
	if err != nil {
		handleHTTPError(w, err)
		return
	}

	var user models.User
	if err := decodeJSONBody(r, &user); err != nil {
		handleHTTPError(w, err)
		return
	}

	if err := validateUserInput(&user); err != nil {
		respondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID = userID
	ctx := r.Context()
	if err := h.userModel.Update(ctx, &user); err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			respondWithError(w, "User not found", http.StatusNotFound)
			return
		}
		if errors.Is(err, models.ErrDuplicateEmail) {
			respondWithError(w, "Email already exists", http.StatusConflict)
			return
		}
		respondWithError(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// Get the updated user to return complete data
	updatedUser, err := h.userModel.Get(ctx, userID)
	if err != nil {
		respondWithError(w, "Failed to get updated user", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedUser)
}

// DeleteUser handles DELETE /api/users/{userID}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	logRequest("DeleteUser")

	if err := validateHTTPMethod(r, http.MethodDelete); err != nil {
		handleHTTPError(w, err)
		return
	}

	userID, err := parseIDFromURL(r, "userID")
	if err != nil {
		handleHTTPError(w, err)
		return
	}

	ctx := r.Context()
	if err := h.userModel.Delete(ctx, userID); err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			respondWithError(w, "User not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]int64{"deleted_id": userID})
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
