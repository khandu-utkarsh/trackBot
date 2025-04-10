package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"workout_app_backend/internal/models"
	// Add other necessary imports like "encoding/json", "log", "strconv"
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
	ctx := r.Context()
	users, err := h.userModel.List(ctx)
	if err != nil {
		http.Error(w, "Failed to list users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// CreateUser handles POST /api/users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if user.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	id, err := h.userModel.Create(ctx, &user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	user.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// GetUser handles GET /api/users/{userID}
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Path[len("/api/users/"):]
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	user, err := h.userModel.Get(ctx, userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// UpdateUser handles PUT /api/users/{userID}
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Path[len("/api/users/"):]
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if user.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	existingUser, err := h.userModel.Get(ctx, userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	existingUser.Email = user.Email
	if err := h.userModel.Update(ctx, existingUser); err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(existingUser); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// DeleteUser handles DELETE /api/users/{userID}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Path[len("/api/users/"):]
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	deletedID, err := h.userModel.Delete(ctx, userID)
	if deletedID == -1 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	response := map[string]int64{"deleted_id": deletedID}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
