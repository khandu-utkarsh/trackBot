package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	models "workout_app_backend/internal/models"
	testutils "workout_app_backend/internal/testutils"

	"github.com/go-chi/chi/v5"
)

func setupTestConversationHandler(t *testing.T) (*ConversationHandler, *models.User, context.Context, func()) {
	db, cleanup := testutils.SetupTestDB(t)
	ctx := context.Background()

	// Initialize user model and create test user
	userModel := models.GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	testUser := &models.User{
		Email: "test@example.com",
	}
	userID, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}
	testUser.ID = userID

	// Initialize conversation model
	conversationModel := models.GetConversationModelInstance(db, "conversations_test", "users_test")
	if err := conversationModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize conversations table: %v", err)
	}

	handler := NewConversationHandler(conversationModel, userModel)

	return handler, testUser, ctx, cleanup
}

func TestConversationHandler_ListConversationsByUser(t *testing.T) {
	handler, testUser, ctx, cleanup := setupTestConversationHandler(t)
	defer cleanup()

	// Create test conversations
	testConversations := []*models.Conversation{
		{
			UserID:   testUser.ID,
			Title:    "First Conversation",
			IsActive: true,
		},
		{
			UserID:   testUser.ID,
			Title:    "Second Conversation",
			IsActive: false,
		},
	}

	for _, conversation := range testConversations {
		_, err := handler.conversationModel.Create(ctx, conversation)
		if err != nil {
			t.Fatalf("Failed to create test conversation: %v", err)
		}
	}

	tests := []struct {
		name       string
		method     string
		userID     int64
		wantStatus int
		wantError  bool
		wantCount  int
	}{
		{
			name:       "Valid user with conversations",
			method:     http.MethodGet,
			userID:     testUser.ID,
			wantStatus: http.StatusOK,
			wantError:  false,
			wantCount:  2,
		},
		{
			name:       "Non-existent user",
			method:     http.MethodGet,
			userID:     99999,
			wantStatus: http.StatusNotFound,
			wantError:  true,
		},
		{
			name:       "Invalid user ID",
			method:     http.MethodGet,
			userID:     0,
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
		{
			name:       "Invalid method",
			method:     http.MethodPost,
			userID:     testUser.ID,
			wantStatus: http.StatusMethodNotAllowed,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := chi.NewRouter()
			r.Get("/api/users/{userID}/conversations", handler.ListConversationsByUser)

			req := httptest.NewRequest(tt.method, fmt.Sprintf("/api/users/%d/conversations", tt.userID), nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusOK {
				var response []models.Conversation
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if len(response) != tt.wantCount {
					t.Errorf("Expected %d conversations, got %d", tt.wantCount, len(response))
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

func TestConversationHandler_CreateConversation(t *testing.T) {
	handler, testUser, _, cleanup := setupTestConversationHandler(t)
	defer cleanup()

	tests := []struct {
		name         string
		method       string
		userID       int64
		conversation *models.Conversation
		wantStatus   int
		wantError    bool
	}{
		{
			name:   "Valid conversation",
			method: http.MethodPost,
			userID: testUser.ID,
			conversation: &models.Conversation{
				Title: "Test Conversation",
			},
			wantStatus: http.StatusCreated,
			wantError:  false,
		},
		{
			name:   "Empty title",
			method: http.MethodPost,
			userID: testUser.ID,
			conversation: &models.Conversation{
				Title: "",
			},
			wantStatus: http.StatusInternalServerError,
			wantError:  true,
		},
		{
			name:   "Non-existent user",
			method: http.MethodPost,
			userID: 99999,
			conversation: &models.Conversation{
				Title: "Test Conversation",
			},
			wantStatus: http.StatusNotFound,
			wantError:  true,
		},
		{
			name:   "Invalid user ID",
			method: http.MethodPost,
			userID: 0,
			conversation: &models.Conversation{
				Title: "Test Conversation",
			},
			wantStatus: http.StatusBadRequest,
			wantError:  true,
		},
		{
			name:   "Invalid method",
			method: http.MethodGet,
			userID: testUser.ID,
			conversation: &models.Conversation{
				Title: "Test Conversation",
			},
			wantStatus: http.StatusMethodNotAllowed,
			wantError:  true,
		},
		{
			name:         "Invalid JSON",
			method:       http.MethodPost,
			userID:       testUser.ID,
			conversation: nil,
			wantStatus:   http.StatusBadRequest,
			wantError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body []byte
			var err error
			if tt.conversation != nil {
				body, err = json.Marshal(tt.conversation)
				if err != nil {
					t.Fatalf("Failed to marshal conversation: %v", err)
				}
			} else {
				body = []byte("invalid json")
			}

			r := chi.NewRouter()
			r.Post("/api/users/{userID}/conversations", handler.CreateConversation)

			req := httptest.NewRequest(tt.method, fmt.Sprintf("/api/users/%d/conversations", tt.userID), bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusCreated {
				var response models.Conversation
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if response.Title != tt.conversation.Title {
					t.Errorf("Expected title %s, got %s", tt.conversation.Title, response.Title)
				}
				if response.UserID != tt.userID {
					t.Errorf("Expected user ID %d, got %d", tt.userID, response.UserID)
				}
				if !response.IsActive {
					t.Error("Expected conversation to be active")
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

func TestConversationHandler_GetConversation(t *testing.T) {
	handler, testUser, ctx, cleanup := setupTestConversationHandler(t)
	defer cleanup()

	// Create a test conversation
	testConversation := &models.Conversation{
		UserID:   testUser.ID,
		Title:    "Test Conversation",
		IsActive: true,
	}
	conversationID, err := handler.conversationModel.Create(ctx, testConversation)
	if err != nil {
		t.Fatalf("Failed to create test conversation: %v", err)
	}

	tests := []struct {
		name           string
		method         string
		userID         int64
		conversationID int64
		wantStatus     int
		wantError      bool
	}{
		{
			name:           "Valid conversation",
			method:         http.MethodGet,
			userID:         testUser.ID,
			conversationID: conversationID,
			wantStatus:     http.StatusOK,
			wantError:      false,
		},
		{
			name:           "Non-existent conversation",
			method:         http.MethodGet,
			userID:         testUser.ID,
			conversationID: 99999,
			wantStatus:     http.StatusNotFound,
			wantError:      true,
		},
		{
			name:           "Wrong user",
			method:         http.MethodGet,
			userID:         99999,
			conversationID: conversationID,
			wantStatus:     http.StatusNotFound,
			wantError:      true,
		},
		{
			name:           "Invalid user ID",
			method:         http.MethodGet,
			userID:         0,
			conversationID: conversationID,
			wantStatus:     http.StatusBadRequest,
			wantError:      true,
		},
		{
			name:           "Invalid conversation ID",
			method:         http.MethodGet,
			userID:         testUser.ID,
			conversationID: 0,
			wantStatus:     http.StatusBadRequest,
			wantError:      true,
		},
		{
			name:           "Invalid method",
			method:         http.MethodPost,
			userID:         testUser.ID,
			conversationID: conversationID,
			wantStatus:     http.StatusMethodNotAllowed,
			wantError:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := chi.NewRouter()
			r.Get("/api/users/{userID}/conversations/{conversationID}", handler.GetConversation)

			req := httptest.NewRequest(tt.method, fmt.Sprintf("/api/users/%d/conversations/%d", tt.userID, tt.conversationID), nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusOK {
				var response models.Conversation
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if response.ID != tt.conversationID {
					t.Errorf("Expected conversation ID %d, got %d", tt.conversationID, response.ID)
				}
				if response.UserID != tt.userID {
					t.Errorf("Expected user ID %d, got %d", tt.userID, response.UserID)
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

func TestConversationHandler_UpdateConversation(t *testing.T) {
	handler, testUser, ctx, cleanup := setupTestConversationHandler(t)
	defer cleanup()

	// Create a test conversation
	testConversation := &models.Conversation{
		UserID:   testUser.ID,
		Title:    "Original Title",
		IsActive: true,
	}
	conversationID, err := handler.conversationModel.Create(ctx, testConversation)
	if err != nil {
		t.Fatalf("Failed to create test conversation: %v", err)
	}

	tests := []struct {
		name           string
		method         string
		userID         int64
		conversationID int64
		updateData     *models.Conversation
		wantStatus     int
		wantError      bool
	}{
		{
			name:           "Valid update",
			method:         http.MethodPut,
			userID:         testUser.ID,
			conversationID: conversationID,
			updateData: &models.Conversation{
				Title:    "Updated Title",
				IsActive: false,
			},
			wantStatus: http.StatusOK,
			wantError:  false,
		},
		{
			name:           "Empty title",
			method:         http.MethodPut,
			userID:         testUser.ID,
			conversationID: conversationID,
			updateData: &models.Conversation{
				Title:    "",
				IsActive: true,
			},
			wantStatus: http.StatusInternalServerError,
			wantError:  true,
		},
		{
			name:           "Non-existent conversation",
			method:         http.MethodPut,
			userID:         testUser.ID,
			conversationID: 99999,
			updateData: &models.Conversation{
				Title:    "Updated Title",
				IsActive: true,
			},
			wantStatus: http.StatusNotFound,
			wantError:  true,
		},
		{
			name:           "Wrong user",
			method:         http.MethodPut,
			userID:         99999,
			conversationID: conversationID,
			updateData: &models.Conversation{
				Title:    "Updated Title",
				IsActive: true,
			},
			wantStatus: http.StatusNotFound,
			wantError:  true,
		},
		{
			name:           "Invalid method",
			method:         http.MethodGet,
			userID:         testUser.ID,
			conversationID: conversationID,
			updateData: &models.Conversation{
				Title:    "Updated Title",
				IsActive: true,
			},
			wantStatus: http.StatusMethodNotAllowed,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.updateData)
			if err != nil {
				t.Fatalf("Failed to marshal conversation: %v", err)
			}

			r := chi.NewRouter()
			r.Put("/api/users/{userID}/conversations/{conversationID}", handler.UpdateConversation)

			req := httptest.NewRequest(tt.method, fmt.Sprintf("/api/users/%d/conversations/%d", tt.userID, tt.conversationID), bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusOK {
				var response models.Conversation
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if response.Title != tt.updateData.Title {
					t.Errorf("Expected title %s, got %s", tt.updateData.Title, response.Title)
				}
				if response.IsActive != tt.updateData.IsActive {
					t.Errorf("Expected IsActive %v, got %v", tt.updateData.IsActive, response.IsActive)
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

func TestConversationHandler_DeleteConversation(t *testing.T) {
	handler, testUser, ctx, cleanup := setupTestConversationHandler(t)
	defer cleanup()

	// Create a test conversation
	testConversation := &models.Conversation{
		UserID:   testUser.ID,
		Title:    "Test Conversation",
		IsActive: true,
	}
	conversationID, err := handler.conversationModel.Create(ctx, testConversation)
	if err != nil {
		t.Fatalf("Failed to create test conversation: %v", err)
	}

	tests := []struct {
		name           string
		method         string
		userID         int64
		conversationID int64
		wantStatus     int
		wantError      bool
	}{
		{
			name:           "Valid deletion",
			method:         http.MethodDelete,
			userID:         testUser.ID,
			conversationID: conversationID,
			wantStatus:     http.StatusOK,
			wantError:      false,
		},
		{
			name:           "Non-existent conversation",
			method:         http.MethodDelete,
			userID:         testUser.ID,
			conversationID: 99999,
			wantStatus:     http.StatusNotFound,
			wantError:      true,
		},
		{
			name:           "Wrong user",
			method:         http.MethodDelete,
			userID:         99999,
			conversationID: conversationID,
			wantStatus:     http.StatusNotFound,
			wantError:      true,
		},
		{
			name:           "Invalid method",
			method:         http.MethodGet,
			userID:         testUser.ID,
			conversationID: conversationID,
			wantStatus:     http.StatusMethodNotAllowed,
			wantError:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := chi.NewRouter()
			r.Delete("/api/users/{userID}/conversations/{conversationID}", handler.DeleteConversation)

			req := httptest.NewRequest(tt.method, fmt.Sprintf("/api/users/%d/conversations/%d", tt.userID, tt.conversationID), nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusOK {
				var response map[string]int64
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if response["deleted_id"] != tt.conversationID {
					t.Errorf("Expected deleted_id %d, got %d", tt.conversationID, response["deleted_id"])
				}

				// Verify the conversation was actually deleted
				_, err := handler.conversationModel.Get(ctx, tt.conversationID)
				if err == nil {
					t.Error("Expected conversation to be deleted")
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
