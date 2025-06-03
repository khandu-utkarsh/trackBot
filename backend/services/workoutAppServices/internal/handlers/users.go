package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	models "workout_app_backend/internal/models"

	"github.com/go-chi/chi/v5"
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
	handlerLogger.Println("ListUsers request received") //! Logging the request.
	if r.Method != http.MethodGet {
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
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
	handlerLogger.Println("CreateUser request received") //! Logging the request.
	if r.Method != http.MethodPost {
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, "Invalid request body", http.StatusBadRequest)
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
	handlerLogger.Println("GetUser request received") //! Logging the request.
	if r.Method != http.MethodGet {
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
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
	handlerLogger.Println("UpdateUser request received") //! Logging the request.
	if r.Method != http.MethodPut {
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, "Invalid request body", http.StatusBadRequest)
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
	handlerLogger.Println("DeleteUser request received") //! Logging the request.
	if r.Method != http.MethodDelete {
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
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
