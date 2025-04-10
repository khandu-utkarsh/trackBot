package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"workout_app_backend/internal/models"
	"workout_app_backend/internal/testutils"
)

func TestUserHandler_ListUsers(t *testing.T) {
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()
	userModel := models.GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	// Create test users
	testUsers := []*models.User{
		{Email: "user1@example.com"},
		{Email: "user2@example.com"},
	}
	for _, user := range testUsers {
		_, err := userModel.Create(ctx, user)
		if err != nil {
			t.Fatalf("Failed to create test user: %v", err)
		}
	}

	handler := GetUserHandlerInstance(userModel)

	req := httptest.NewRequest("GET", "/api/users", nil)
	w := httptest.NewRecorder()

	handler.ListUsers(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response []models.User
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if len(response) != len(testUsers) {
		t.Errorf("Expected %d users, got %d", len(testUsers), len(response))
	}
}

func TestUserHandler_CreateUser(t *testing.T) {
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()
	userModel := models.GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	handler := GetUserHandlerInstance(userModel)

	tests := []struct {
		name       string
		user       *models.User
		wantStatus int
	}{
		{
			name: "Valid user",
			user: &models.User{
				Email: "test@example.com",
			},
			wantStatus: http.StatusCreated,
		},
		{
			name: "Empty email",
			user: &models.User{
				Email: "",
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.user)
			if err != nil {
				t.Fatalf("Failed to marshal user: %v", err)
			}

			req := httptest.NewRequest("POST", "/api/users", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			handler.CreateUser(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusCreated {
				var response models.User
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if response.Email != tt.user.Email {
					t.Errorf("Expected email %s, got %s", tt.user.Email, response.Email)
				}
			}
		})
	}
}

func TestUserHandler_GetUser(t *testing.T) {
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()
	userModel := models.GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	// Create a test user
	testUser := &models.User{
		Email: "test@example.com",
	}
	id, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}
	testUser.ID = id

	handler := GetUserHandlerInstance(userModel)

	tests := []struct {
		name       string
		userID     int64
		wantStatus int
	}{
		{
			name:       "Existing user",
			userID:     testUser.ID,
			wantStatus: http.StatusOK,
		},
		{
			name:       "Non-existent user",
			userID:     99999,
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/api/users/"+strconv.FormatInt(tt.userID, 10), nil)
			w := httptest.NewRecorder()

			handler.GetUser(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusOK {
				var response models.User
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if response.ID != tt.userID {
					t.Errorf("Expected user ID %d, got %d", tt.userID, response.ID)
				}
			}
		})
	}
}

func TestUserHandler_UpdateUser(t *testing.T) {
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()
	userModel := models.GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	// Create a test user
	testUser := &models.User{
		Email: "test@example.com",
	}
	id, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}
	testUser.ID = id

	handler := GetUserHandlerInstance(userModel)

	tests := []struct {
		name       string
		userID     int64
		updateUser *models.User
		wantStatus int
	}{
		{
			name:   "Valid update",
			userID: testUser.ID,
			updateUser: &models.User{
				Email: "updated@example.com",
			},
			wantStatus: http.StatusOK,
		},
		{
			name:   "Non-existent user",
			userID: 99999,
			updateUser: &models.User{
				Email: "updated@example.com",
			},
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.updateUser)
			if err != nil {
				t.Fatalf("Failed to marshal user: %v", err)
			}

			req := httptest.NewRequest("PUT", "/api/users/"+strconv.FormatInt(tt.userID, 10), bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			handler.UpdateUser(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusOK {
				var response models.User
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if response.Email != tt.updateUser.Email {
					t.Errorf("Expected email %s, got %s", tt.updateUser.Email, response.Email)
				}
			}
		})
	}
}

func TestUserHandler_DeleteUser(t *testing.T) {
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()
	userModel := models.GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	// Create a test user
	testUser := &models.User{
		Email: "test@example.com",
	}
	id, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}
	testUser.ID = id

	handler := GetUserHandlerInstance(userModel)

	tests := []struct {
		name       string
		userID     int64
		wantStatus int
	}{
		{
			name:       "Existing user",
			userID:     testUser.ID,
			wantStatus: http.StatusNoContent,
		},
		{
			name:       "Non-existent user",
			userID:     99999,
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("DELETE", "/api/users/"+strconv.FormatInt(tt.userID, 10), nil)
			w := httptest.NewRecorder()

			handler.DeleteUser(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusNoContent {
				// Verify the user was actually deleted
				_, err := userModel.Get(ctx, tt.userID)
				if err == nil {
					t.Error("Expected user to be deleted")
				}
			}
		})
	}
}
