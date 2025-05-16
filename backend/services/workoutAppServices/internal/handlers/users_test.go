package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	models "workout_app_backend/internal/models"
	testutils "workout_app_backend/internal/testutils"
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

	tests := []struct {
		name       string
		method     string
		wantStatus int
	}{
		{
			name:       "Valid GET request",
			method:     http.MethodGet,
			wantStatus: http.StatusOK,
		},
		{
			name:       "Invalid method",
			method:     http.MethodPost,
			wantStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/api/users", nil)
			w := httptest.NewRecorder()

			handler.ListUsers(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusOK {
				var response []models.User
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if len(response) != len(testUsers) {
					t.Errorf("Expected %d users, got %d", len(testUsers), len(response))
				}
			}
		})
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
		method     string
		user       *models.User
		wantStatus int
		wantError  bool
	}{
		{
			name:   "Valid user",
			method: http.MethodPost,
			user: &models.User{
				Email: "test@example.com",
			},
			wantStatus: http.StatusCreated,
			wantError:  false,
		},
		{
			name:   "Empty email",
			method: http.MethodPost,
			user: &models.User{
				Email: "",
			},
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
		{
			name:   "Invalid email format",
			method: http.MethodPost,
			user: &models.User{
				Email: "invalid-email",
			},
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
		{
			name:   "Invalid method",
			method: http.MethodGet,
			user: &models.User{
				Email: "test@example.com",
			},
			wantStatus: http.StatusMethodNotAllowed,
			wantError:  true,
		},
		{
			name:       "Invalid JSON",
			method:     http.MethodPost,
			user:       nil,
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body []byte
			var err error
			if tt.user != nil {
				body, err = json.Marshal(tt.user)
				if err != nil {
					t.Fatalf("Failed to marshal user: %v", err)
				}
			} else {
				body = []byte("invalid json")
			}

			req := httptest.NewRequest(tt.method, "/api/users", bytes.NewBuffer(body))
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

			if tt.wantError {
				var errorResponse map[string]string
				if err := json.NewDecoder(w.Body).Decode(&errorResponse); err != nil {
					t.Fatalf("Failed to decode error response: %v", err)
				}
				if _, ok := errorResponse["error"]; !ok {
					t.Error("Expected error response to contain 'error' field")
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
		method     string
		userID     int64
		wantStatus int
		wantError  bool
	}{
		{
			name:       "Existing user",
			method:     http.MethodGet,
			userID:     testUser.ID,
			wantStatus: http.StatusOK,
			wantError:  false,
		},
		{
			name:       "Non-existent user",
			method:     http.MethodGet,
			userID:     99999,
			wantStatus: http.StatusNotFound,
			wantError:  true,
		},
		{
			name:       "Invalid method",
			method:     http.MethodPost,
			userID:     testUser.ID,
			wantStatus: http.StatusMethodNotAllowed,
			wantError:  true,
		},
		{
			name:       "Invalid user ID format",
			method:     http.MethodGet,
			userID:     -1,
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/api/users/"+strconv.FormatInt(tt.userID, 10), nil)
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

			if tt.wantError {
				var errorResponse map[string]string
				if err := json.NewDecoder(w.Body).Decode(&errorResponse); err != nil {
					t.Fatalf("Failed to decode error response: %v", err)
				}
				if _, ok := errorResponse["error"]; !ok {
					t.Error("Expected error response to contain 'error' field")
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
		method     string
		userID     int64
		updateUser *models.User
		wantStatus int
		wantError  bool
	}{
		{
			name:   "Valid update",
			method: http.MethodPut,
			userID: testUser.ID,
			updateUser: &models.User{
				Email: "updated@example.com",
			},
			wantStatus: http.StatusOK,
			wantError:  false,
		},
		{
			name:   "Non-existent user",
			method: http.MethodPut,
			userID: 99999,
			updateUser: &models.User{
				Email: "updated@example.com",
			},
			wantStatus: http.StatusNotFound,
			wantError:  true,
		},
		{
			name:   "Invalid method",
			method: http.MethodPost,
			userID: testUser.ID,
			updateUser: &models.User{
				Email: "updated@example.com",
			},
			wantStatus: http.StatusMethodNotAllowed,
			wantError:  true,
		},
		{
			name:   "Invalid email format",
			method: http.MethodPut,
			userID: testUser.ID,
			updateUser: &models.User{
				Email: "invalid-email",
			},
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
		{
			name:   "Empty email",
			method: http.MethodPut,
			userID: testUser.ID,
			updateUser: &models.User{
				Email: "",
			},
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.updateUser)
			if err != nil {
				t.Fatalf("Failed to marshal user: %v", err)
			}

			req := httptest.NewRequest(tt.method, "/api/users/"+strconv.FormatInt(tt.userID, 10), bytes.NewBuffer(body))
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

			if tt.wantError {
				var errorResponse map[string]string
				if err := json.NewDecoder(w.Body).Decode(&errorResponse); err != nil {
					t.Fatalf("Failed to decode error response: %v", err)
				}
				if _, ok := errorResponse["error"]; !ok {
					t.Error("Expected error response to contain 'error' field")
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
		method     string
		userID     int64
		wantStatus int
		wantError  bool
	}{
		{
			name:       "Existing user",
			method:     http.MethodDelete,
			userID:     testUser.ID,
			wantStatus: http.StatusOK,
			wantError:  false,
		},
		{
			name:       "Non-existent user",
			method:     http.MethodDelete,
			userID:     99999,
			wantStatus: http.StatusNotFound,
			wantError:  true,
		},
		{
			name:       "Invalid method",
			method:     http.MethodPost,
			userID:     testUser.ID,
			wantStatus: http.StatusMethodNotAllowed,
			wantError:  true,
		},
		{
			name:       "Invalid user ID format",
			method:     http.MethodDelete,
			userID:     -1,
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/api/users/"+strconv.FormatInt(tt.userID, 10), nil)
			w := httptest.NewRecorder()

			handler.DeleteUser(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusOK {
				var response map[string]int64
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if response["deleted_id"] != tt.userID {
					t.Errorf("Expected deleted_id %d, got %d", tt.userID, response["deleted_id"])
				}

				// Verify the user was actually deleted
				_, err := userModel.Get(ctx, tt.userID)
				if err == nil {
					t.Error("Expected user to be deleted")
				}
			}

			if tt.wantError {
				var errorResponse map[string]string
				if err := json.NewDecoder(w.Body).Decode(&errorResponse); err != nil {
					t.Fatalf("Failed to decode error response: %v", err)
				}
				if _, ok := errorResponse["error"]; !ok {
					t.Error("Expected error response to contain 'error' field")
				}
			}
		})
	}
}
