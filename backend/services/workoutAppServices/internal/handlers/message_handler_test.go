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

func setupTestMessageHandler(t *testing.T) (*MessageHandler, *models.User, *models.Conversation, context.Context, func()) {
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

	// Initialize conversation model and create test conversation
	conversationModel := models.GetConversationModelInstance(db, "conversations_test", "users_test")
	if err := conversationModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize conversations table: %v", err)
	}

	testConversation := &models.Conversation{
		UserID:   testUser.ID,
		Title:    "Test Conversation",
		IsActive: true,
	}
	conversationID, err := conversationModel.Create(ctx, testConversation)
	if err != nil {
		t.Fatalf("Failed to create test conversation: %v", err)
	}
	testConversation.ID = conversationID

	// Initialize message model
	messageModel := models.GetMessageModelInstance(db, "messages_test", "conversations_test")
	if err := messageModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize messages table: %v", err)
	}

	handler := NewMessageHandler(messageModel, conversationModel, nil)

	return handler, testUser, testConversation, ctx, cleanup
}

func TestMessageHandler_ListMessagesByConversation(t *testing.T) {
	handler, testUser, testConversation, ctx, cleanup := setupTestMessageHandler(t)
	defer cleanup()

	// Create test messages
	testMessages := []*models.Message{
		{
			ConversationID: testConversation.ID,
			UserID:         testUser.ID,
			Content:        "First message",
			MessageType:    models.MessageTypeUser,
		},
		{
			ConversationID: testConversation.ID,
			UserID:         testUser.ID,
			Content:        "Second message",
			MessageType:    models.MessageTypeAssistant,
		},
	}

	for _, message := range testMessages {
		_, err := handler.messageModel.Create(ctx, message)
		if err != nil {
			t.Fatalf("Failed to create test message: %v", err)
		}
	}

	tests := []struct {
		name           string
		method         string
		userID         int64
		conversationID int64
		wantStatus     int
		wantError      bool
		wantCount      int
	}{
		{
			name:           "Valid conversation with messages",
			method:         http.MethodGet,
			userID:         testUser.ID,
			conversationID: testConversation.ID,
			wantStatus:     http.StatusOK,
			wantError:      false,
			wantCount:      2,
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
			conversationID: testConversation.ID,
			wantStatus:     http.StatusNotFound,
			wantError:      true,
		},
		{
			name:           "Invalid user ID",
			method:         http.MethodGet,
			userID:         0,
			conversationID: testConversation.ID,
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
			conversationID: testConversation.ID,
			wantStatus:     http.StatusMethodNotAllowed,
			wantError:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := chi.NewRouter()
			r.Get("/api/users/{userID}/conversations/{conversationID}/messages", handler.ListMessagesByConversation)

			req := httptest.NewRequest(tt.method, fmt.Sprintf("/api/users/%d/conversations/%d/messages", tt.userID, tt.conversationID), nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusOK {
				var response []models.Message
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if len(response) != tt.wantCount {
					t.Errorf("Expected %d messages, got %d", tt.wantCount, len(response))
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

func TestMessageHandler_CreateMessage(t *testing.T) {
	handler, testUser, testConversation, _, cleanup := setupTestMessageHandler(t)
	defer cleanup()

	tests := []struct {
		name           string
		method         string
		userID         int64
		conversationID int64
		message        *models.Message
		wantStatus     int
		wantError      bool
	}{
		{
			name:           "Valid user message",
			method:         http.MethodPost,
			userID:         testUser.ID,
			conversationID: testConversation.ID,
			message: &models.Message{
				Content:     "Hello, this is a test message",
				MessageType: models.MessageTypeUser,
			},
			wantStatus: http.StatusCreated,
			wantError:  false,
		},
		{
			name:           "Valid assistant message",
			method:         http.MethodPost,
			userID:         testUser.ID,
			conversationID: testConversation.ID,
			message: &models.Message{
				Content:     "Hello! How can I help you?",
				MessageType: models.MessageTypeAssistant,
			},
			wantStatus: http.StatusCreated,
			wantError:  false,
		},
		{
			name:           "Empty content",
			method:         http.MethodPost,
			userID:         testUser.ID,
			conversationID: testConversation.ID,
			message: &models.Message{
				Content:     "",
				MessageType: models.MessageTypeUser,
			},
			wantStatus: http.StatusInternalServerError,
			wantError:  true,
		},
		{
			name:           "Invalid message type",
			method:         http.MethodPost,
			userID:         testUser.ID,
			conversationID: testConversation.ID,
			message: &models.Message{
				Content:     "Test message",
				MessageType: "invalid",
			},
			wantStatus: http.StatusInternalServerError,
			wantError:  true,
		},
		{
			name:           "Non-existent conversation",
			method:         http.MethodPost,
			userID:         testUser.ID,
			conversationID: 99999,
			message: &models.Message{
				Content:     "Test message",
				MessageType: models.MessageTypeUser,
			},
			wantStatus: http.StatusNotFound,
			wantError:  true,
		},
		{
			name:           "Wrong user",
			method:         http.MethodPost,
			userID:         99999,
			conversationID: testConversation.ID,
			message: &models.Message{
				Content:     "Test message",
				MessageType: models.MessageTypeUser,
			},
			wantStatus: http.StatusNotFound,
			wantError:  true,
		},
		{
			name:           "Invalid method",
			method:         http.MethodGet,
			userID:         testUser.ID,
			conversationID: testConversation.ID,
			message: &models.Message{
				Content:     "Test message",
				MessageType: models.MessageTypeUser,
			},
			wantStatus: http.StatusMethodNotAllowed,
			wantError:  true,
		},
		{
			name:           "Invalid JSON",
			method:         http.MethodPost,
			userID:         testUser.ID,
			conversationID: testConversation.ID,
			message:        nil,
			wantStatus:     http.StatusBadRequest,
			wantError:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body []byte
			var err error
			if tt.message != nil {
				body, err = json.Marshal(tt.message)
				if err != nil {
					t.Fatalf("Failed to marshal message: %v", err)
				}
			} else {
				body = []byte("invalid json")
			}

			r := chi.NewRouter()
			r.Post("/api/users/{userID}/conversations/{conversationID}/messages", handler.CreateMessage)

			req := httptest.NewRequest(tt.method, fmt.Sprintf("/api/users/%d/conversations/%d/messages", tt.userID, tt.conversationID), bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusCreated {
				var response models.Message
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if response.Content != tt.message.Content {
					t.Errorf("Expected content %s, got %s", tt.message.Content, response.Content)
				}
				if response.MessageType != tt.message.MessageType {
					t.Errorf("Expected message type %s, got %s", tt.message.MessageType, response.MessageType)
				}
				if response.UserID != tt.userID {
					t.Errorf("Expected user ID %d, got %d", tt.userID, response.UserID)
				}
				if response.ConversationID != tt.conversationID {
					t.Errorf("Expected conversation ID %d, got %d", tt.conversationID, response.ConversationID)
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

func TestMessageHandler_GetMessage(t *testing.T) {
	handler, testUser, testConversation, ctx, cleanup := setupTestMessageHandler(t)
	defer cleanup()

	// Create a test message
	testMessage := &models.Message{
		ConversationID: testConversation.ID,
		UserID:         testUser.ID,
		Content:        "Test message content",
		MessageType:    models.MessageTypeUser,
	}
	messageID, err := handler.messageModel.Create(ctx, testMessage)
	if err != nil {
		t.Fatalf("Failed to create test message: %v", err)
	}

	tests := []struct {
		name           string
		method         string
		userID         int64
		conversationID int64
		messageID      int64
		wantStatus     int
		wantError      bool
	}{
		{
			name:           "Valid message",
			method:         http.MethodGet,
			userID:         testUser.ID,
			conversationID: testConversation.ID,
			messageID:      messageID,
			wantStatus:     http.StatusOK,
			wantError:      false,
		},
		{
			name:           "Non-existent message",
			method:         http.MethodGet,
			userID:         testUser.ID,
			conversationID: testConversation.ID,
			messageID:      99999,
			wantStatus:     http.StatusNotFound,
			wantError:      true,
		},
		{
			name:           "Wrong conversation",
			method:         http.MethodGet,
			userID:         testUser.ID,
			conversationID: 99999,
			messageID:      messageID,
			wantStatus:     http.StatusNotFound,
			wantError:      true,
		},
		{
			name:           "Wrong user",
			method:         http.MethodGet,
			userID:         99999,
			conversationID: testConversation.ID,
			messageID:      messageID,
			wantStatus:     http.StatusNotFound,
			wantError:      true,
		},
		{
			name:           "Invalid method",
			method:         http.MethodPost,
			userID:         testUser.ID,
			conversationID: testConversation.ID,
			messageID:      messageID,
			wantStatus:     http.StatusMethodNotAllowed,
			wantError:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := chi.NewRouter()
			r.Get("/api/users/{userID}/conversations/{conversationID}/messages/{messageID}", handler.GetMessage)

			req := httptest.NewRequest(tt.method, fmt.Sprintf("/api/users/%d/conversations/%d/messages/%d", tt.userID, tt.conversationID, tt.messageID), nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusOK {
				var response models.Message
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if response.ID != tt.messageID {
					t.Errorf("Expected message ID %d, got %d", tt.messageID, response.ID)
				}
				if response.ConversationID != tt.conversationID {
					t.Errorf("Expected conversation ID %d, got %d", tt.conversationID, response.ConversationID)
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

func TestMessageHandler_UpdateMessage(t *testing.T) {
	handler, testUser, testConversation, ctx, cleanup := setupTestMessageHandler(t)
	defer cleanup()

	// Create a test message
	testMessage := &models.Message{
		ConversationID: testConversation.ID,
		UserID:         testUser.ID,
		Content:        "Original content",
		MessageType:    models.MessageTypeUser,
	}
	messageID, err := handler.messageModel.Create(ctx, testMessage)
	if err != nil {
		t.Fatalf("Failed to create test message: %v", err)
	}

	tests := []struct {
		name           string
		method         string
		userID         int64
		conversationID int64
		messageID      int64
		updateData     *models.Message
		wantStatus     int
		wantError      bool
	}{
		{
			name:           "Valid update",
			method:         http.MethodPut,
			userID:         testUser.ID,
			conversationID: testConversation.ID,
			messageID:      messageID,
			updateData: &models.Message{
				Content:     "Updated content",
				MessageType: models.MessageTypeUser,
			},
			wantStatus: http.StatusOK,
			wantError:  false,
		},
		{
			name:           "Empty content",
			method:         http.MethodPut,
			userID:         testUser.ID,
			conversationID: testConversation.ID,
			messageID:      messageID,
			updateData: &models.Message{
				Content:     "",
				MessageType: models.MessageTypeUser,
			},
			wantStatus: http.StatusInternalServerError,
			wantError:  true,
		},
		{
			name:           "Non-existent message",
			method:         http.MethodPut,
			userID:         testUser.ID,
			conversationID: testConversation.ID,
			messageID:      99999,
			updateData: &models.Message{
				Content:     "Updated content",
				MessageType: models.MessageTypeUser,
			},
			wantStatus: http.StatusNotFound,
			wantError:  true,
		},
		{
			name:           "Wrong conversation",
			method:         http.MethodPut,
			userID:         testUser.ID,
			conversationID: 99999,
			messageID:      messageID,
			updateData: &models.Message{
				Content:     "Updated content",
				MessageType: models.MessageTypeUser,
			},
			wantStatus: http.StatusNotFound,
			wantError:  true,
		},
		{
			name:           "Invalid method",
			method:         http.MethodGet,
			userID:         testUser.ID,
			conversationID: testConversation.ID,
			messageID:      messageID,
			updateData: &models.Message{
				Content:     "Updated content",
				MessageType: models.MessageTypeUser,
			},
			wantStatus: http.StatusMethodNotAllowed,
			wantError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.updateData)
			if err != nil {
				t.Fatalf("Failed to marshal message: %v", err)
			}

			r := chi.NewRouter()
			r.Put("/api/users/{userID}/conversations/{conversationID}/messages/{messageID}", handler.UpdateMessage)

			req := httptest.NewRequest(tt.method, fmt.Sprintf("/api/users/%d/conversations/%d/messages/%d", tt.userID, tt.conversationID, tt.messageID), bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Expected status code %d, got %d", tt.wantStatus, w.Code)
			}

			if tt.wantStatus == http.StatusOK {
				var response models.Message
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if response.Content != tt.updateData.Content {
					t.Errorf("Expected content %s, got %s", tt.updateData.Content, response.Content)
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

func TestMessageHandler_DeleteMessage(t *testing.T) {
	handler, testUser, testConversation, ctx, cleanup := setupTestMessageHandler(t)
	defer cleanup()

	// Create a test message
	testMessage := &models.Message{
		ConversationID: testConversation.ID,
		UserID:         testUser.ID,
		Content:        "Test message",
		MessageType:    models.MessageTypeUser,
	}
	messageID, err := handler.messageModel.Create(ctx, testMessage)
	if err != nil {
		t.Fatalf("Failed to create test message: %v", err)
	}

	tests := []struct {
		name           string
		method         string
		userID         int64
		conversationID int64
		messageID      int64
		wantStatus     int
		wantError      bool
	}{
		{
			name:           "Valid deletion",
			method:         http.MethodDelete,
			userID:         testUser.ID,
			conversationID: testConversation.ID,
			messageID:      messageID,
			wantStatus:     http.StatusOK,
			wantError:      false,
		},
		{
			name:           "Non-existent message",
			method:         http.MethodDelete,
			userID:         testUser.ID,
			conversationID: testConversation.ID,
			messageID:      99999,
			wantStatus:     http.StatusNotFound,
			wantError:      true,
		},
		{
			name:           "Wrong conversation",
			method:         http.MethodDelete,
			userID:         testUser.ID,
			conversationID: 99999,
			messageID:      messageID,
			wantStatus:     http.StatusNotFound,
			wantError:      true,
		},
		{
			name:           "Invalid method",
			method:         http.MethodGet,
			userID:         testUser.ID,
			conversationID: testConversation.ID,
			messageID:      messageID,
			wantStatus:     http.StatusMethodNotAllowed,
			wantError:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := chi.NewRouter()
			r.Delete("/api/users/{userID}/conversations/{conversationID}/messages/{messageID}", handler.DeleteMessage)

			req := httptest.NewRequest(tt.method, fmt.Sprintf("/api/users/%d/conversations/%d/messages/%d", tt.userID, tt.conversationID, tt.messageID), nil)
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
				if response["deleted_id"] != tt.messageID {
					t.Errorf("Expected deleted_id %d, got %d", tt.messageID, response["deleted_id"])
				}

				// Verify the message was actually deleted
				_, err := handler.messageModel.Get(ctx, tt.messageID)
				if err == nil {
					t.Error("Expected message to be deleted")
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
