package models

import (
	"context"
	"strings"
	"testing"
	testutils "workout_app_backend/internal/testutils"
)

func setupTestMessageModel(t *testing.T) (*MessageModel, *ConversationModel, *UserModel, *User, *Conversation, context.Context, func()) {
	db, cleanup := testutils.SetupTestDB(t)
	ctx := context.Background()

	// Initialize user model and create test user
	userModel := GetUserModelInstance(db, "users_test")
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	testUser := &User{
		Email: "test@example.com",
	}
	userID, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}
	testUser.ID = userID

	// Initialize conversation model and create test conversation
	conversationModel := GetConversationModelInstance(db, "conversations_test", "users_test")
	if err := conversationModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize conversations table: %v", err)
	}

	testConversation := &Conversation{
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
	messageModel := GetMessageModelInstance(db, "messages_test", "conversations_test")
	if err := messageModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize messages table: %v", err)
	}

	return messageModel, conversationModel, userModel, testUser, testConversation, ctx, cleanup
}

func TestMessageModel_Create(t *testing.T) {
	messageModel, _, _, testUser, testConversation, ctx, cleanup := setupTestMessageModel(t)
	defer cleanup()

	tests := []struct {
		name        string
		message     *Message
		wantErr     bool
		errContains string
	}{
		{
			name: "valid user message",
			message: &Message{
				ConversationID: testConversation.ID,
				UserID:         testUser.ID,
				Content:        "Hello, this is a test message",
				MessageType:    MessageTypeUser,
			},
			wantErr: false,
		},
		{
			name: "valid assistant message",
			message: &Message{
				ConversationID: testConversation.ID,
				UserID:         testUser.ID,
				Content:        "Hello! How can I help you today?",
				MessageType:    MessageTypeAssistant,
			},
			wantErr: false,
		},
		{
			name: "valid system message",
			message: &Message{
				ConversationID: testConversation.ID,
				UserID:         testUser.ID,
				Content:        "System notification",
				MessageType:    MessageTypeSystem,
			},
			wantErr: false,
		},
		{
			name: "message with empty content",
			message: &Message{
				ConversationID: testConversation.ID,
				UserID:         testUser.ID,
				Content:        "",
				MessageType:    MessageTypeUser,
			},
			wantErr:     true,
			errContains: "message content cannot be empty",
		},
		{
			name: "message with invalid conversation ID",
			message: &Message{
				ConversationID: 0,
				UserID:         testUser.ID,
				Content:        "Test message",
				MessageType:    MessageTypeUser,
			},
			wantErr:     true,
			errContains: "invalid conversation ID",
		},
		{
			name: "message with invalid user ID",
			message: &Message{
				ConversationID: testConversation.ID,
				UserID:         0,
				Content:        "Test message",
				MessageType:    MessageTypeUser,
			},
			wantErr:     true,
			errContains: "invalid user ID",
		},
		{
			name: "message with invalid message type",
			message: &Message{
				ConversationID: testConversation.ID,
				UserID:         testUser.ID,
				Content:        "Test message",
				MessageType:    "invalid",
			},
			wantErr:     true,
			errContains: "invalid message type",
		},
		{
			name:        "nil message",
			message:     nil,
			wantErr:     true,
			errContains: "message cannot be nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := messageModel.Create(ctx, tt.message)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Create() error = nil, wantErr %v", tt.wantErr)
					return
				}
				if !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("Create() error = %v, want error containing %v", err, tt.errContains)
				}
				return
			}

			if err != nil {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if id <= 0 {
				t.Errorf("Create() returned invalid ID = %v", id)
			}

			// Verify the message was created
			created, err := messageModel.Get(ctx, id)
			if err != nil {
				t.Errorf("Failed to get created message: %v", err)
				return
			}

			if created.ConversationID != tt.message.ConversationID {
				t.Errorf("Created message ConversationID = %v, want %v", created.ConversationID, tt.message.ConversationID)
			}
			if created.UserID != tt.message.UserID {
				t.Errorf("Created message UserID = %v, want %v", created.UserID, tt.message.UserID)
			}
			if created.Content != tt.message.Content {
				t.Errorf("Created message Content = %v, want %v", created.Content, tt.message.Content)
			}
			if created.MessageType != tt.message.MessageType {
				t.Errorf("Created message MessageType = %v, want %v", created.MessageType, tt.message.MessageType)
			}
		})
	}
}

func TestMessageModel_Get(t *testing.T) {
	messageModel, _, _, testUser, testConversation, ctx, cleanup := setupTestMessageModel(t)
	defer cleanup()

	// Create a test message
	testMessage := &Message{
		ConversationID: testConversation.ID,
		UserID:         testUser.ID,
		Content:        "Test message content",
		MessageType:    MessageTypeUser,
	}
	id, err := messageModel.Create(ctx, testMessage)
	if err != nil {
		t.Fatalf("Failed to create test message: %v", err)
	}

	tests := []struct {
		name        string
		id          int64
		wantErr     bool
		errContains string
	}{
		{
			name:    "valid message ID",
			id:      id,
			wantErr: false,
		},
		{
			name:        "invalid message ID",
			id:          0,
			wantErr:     true,
			errContains: "invalid message ID",
		},
		{
			name:        "non-existent message ID",
			id:          99999,
			wantErr:     true,
			errContains: "message not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message, err := messageModel.Get(ctx, tt.id)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Get() error = nil, wantErr %v", tt.wantErr)
					return
				}
				if !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("Get() error = %v, want error containing %v", err, tt.errContains)
				}
				return
			}

			if err != nil {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if message.ID != tt.id {
				t.Errorf("Get() message ID = %v, want %v", message.ID, tt.id)
			}
			if message.ConversationID != testConversation.ID {
				t.Errorf("Get() message ConversationID = %v, want %v", message.ConversationID, testConversation.ID)
			}
			if message.UserID != testUser.ID {
				t.Errorf("Get() message UserID = %v, want %v", message.UserID, testUser.ID)
			}
			if message.Content != testMessage.Content {
				t.Errorf("Get() message Content = %v, want %v", message.Content, testMessage.Content)
			}
		})
	}
}

func TestMessageModel_ListByConversation(t *testing.T) {
	messageModel, _, _, testUser, testConversation, ctx, cleanup := setupTestMessageModel(t)
	defer cleanup()

	// Create test messages
	testMessages := []*Message{
		{
			ConversationID: testConversation.ID,
			UserID:         testUser.ID,
			Content:        "First message",
			MessageType:    MessageTypeUser,
		},
		{
			ConversationID: testConversation.ID,
			UserID:         testUser.ID,
			Content:        "Second message",
			MessageType:    MessageTypeAssistant,
		},
		{
			ConversationID: testConversation.ID,
			UserID:         testUser.ID,
			Content:        "Third message",
			MessageType:    MessageTypeUser,
		},
	}

	for _, message := range testMessages {
		_, err := messageModel.Create(ctx, message)
		if err != nil {
			t.Fatalf("Failed to create test message: %v", err)
		}
	}

	tests := []struct {
		name           string
		conversationID int64
		wantCount      int
		wantErr        bool
		errContains    string
	}{
		{
			name:           "valid conversation ID",
			conversationID: testConversation.ID,
			wantCount:      3,
			wantErr:        false,
		},
		{
			name:           "invalid conversation ID",
			conversationID: 0,
			wantErr:        true,
			errContains:    "invalid conversation ID",
		},
		{
			name:           "non-existent conversation ID",
			conversationID: 99999,
			wantCount:      0,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			messages, err := messageModel.ListByConversation(ctx, tt.conversationID)

			if tt.wantErr {
				if err == nil {
					t.Errorf("ListByConversation() error = nil, wantErr %v", tt.wantErr)
					return
				}
				if !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("ListByConversation() error = %v, want error containing %v", err, tt.errContains)
				}
				return
			}

			if err != nil {
				t.Errorf("ListByConversation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(messages) != tt.wantCount {
				t.Errorf("ListByConversation() returned %v messages, want %v", len(messages), tt.wantCount)
			}

			// Verify messages belong to the conversation and are ordered by created_at
			for i, message := range messages {
				if message.ConversationID != tt.conversationID {
					t.Errorf("ListByConversation() message ConversationID = %v, want %v", message.ConversationID, tt.conversationID)
				}

				// Check ordering (should be ascending by created_at)
				if i > 0 && message.CreatedAt.Before(messages[i-1].CreatedAt) {
					t.Errorf("ListByConversation() messages not ordered correctly by created_at")
				}
			}
		})
	}
}

func TestMessageModel_Update(t *testing.T) {
	messageModel, _, _, testUser, testConversation, ctx, cleanup := setupTestMessageModel(t)
	defer cleanup()

	// Create a test message
	testMessage := &Message{
		ConversationID: testConversation.ID,
		UserID:         testUser.ID,
		Content:        "Original content",
		MessageType:    MessageTypeUser,
	}
	id, err := messageModel.Create(ctx, testMessage)
	if err != nil {
		t.Fatalf("Failed to create test message: %v", err)
	}
	testMessage.ID = id

	tests := []struct {
		name        string
		message     *Message
		wantErr     bool
		errContains string
	}{
		{
			name: "valid update",
			message: &Message{
				ID:             id,
				ConversationID: testConversation.ID,
				UserID:         testUser.ID,
				Content:        "Updated content",
				MessageType:    MessageTypeUser,
			},
			wantErr: false,
		},
		{
			name: "update with empty content",
			message: &Message{
				ID:             id,
				ConversationID: testConversation.ID,
				UserID:         testUser.ID,
				Content:        "",
				MessageType:    MessageTypeUser,
			},
			wantErr:     true,
			errContains: "message content cannot be empty",
		},
		{
			name: "update non-existent message",
			message: &Message{
				ID:             99999,
				ConversationID: testConversation.ID,
				UserID:         testUser.ID,
				Content:        "Updated content",
				MessageType:    MessageTypeUser,
			},
			wantErr:     true,
			errContains: "message not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := messageModel.Update(ctx, tt.message)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Update() error = nil, wantErr %v", tt.wantErr)
					return
				}
				if !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("Update() error = %v, want error containing %v", err, tt.errContains)
				}
				return
			}

			if err != nil {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Verify the message was updated
			updated, err := messageModel.Get(ctx, tt.message.ID)
			if err != nil {
				t.Errorf("Failed to get updated message: %v", err)
				return
			}

			if updated.Content != tt.message.Content {
				t.Errorf("Updated message Content = %v, want %v", updated.Content, tt.message.Content)
			}
		})
	}
}

func TestMessageModel_Delete(t *testing.T) {
	messageModel, _, _, testUser, testConversation, ctx, cleanup := setupTestMessageModel(t)
	defer cleanup()

	// Create a test message
	testMessage := &Message{
		ConversationID: testConversation.ID,
		UserID:         testUser.ID,
		Content:        "Test message",
		MessageType:    MessageTypeUser,
	}
	id, err := messageModel.Create(ctx, testMessage)
	if err != nil {
		t.Fatalf("Failed to create test message: %v", err)
	}

	tests := []struct {
		name        string
		id          int64
		wantErr     bool
		errContains string
	}{
		{
			name:    "valid message ID",
			id:      id,
			wantErr: false,
		},
		{
			name:        "invalid message ID",
			id:          0,
			wantErr:     true,
			errContains: "invalid message ID",
		},
		{
			name:        "non-existent message ID",
			id:          99999,
			wantErr:     true,
			errContains: "message not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := messageModel.Delete(ctx, tt.id)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Delete() error = nil, wantErr %v", tt.wantErr)
					return
				}
				if !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("Delete() error = %v, want error containing %v", err, tt.errContains)
				}
				return
			}

			if err != nil {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Verify the message was deleted
			_, err = messageModel.Get(ctx, tt.id)
			if err == nil {
				t.Errorf("Delete() message still exists after deletion")
			}
		})
	}
}

func TestMessageModel_CascadeDelete(t *testing.T) {
	messageModel, conversationModel, _, testUser, testConversation, ctx, cleanup := setupTestMessageModel(t)
	defer cleanup()

	// Create test messages
	testMessages := []*Message{
		{
			ConversationID: testConversation.ID,
			UserID:         testUser.ID,
			Content:        "First message",
			MessageType:    MessageTypeUser,
		},
		{
			ConversationID: testConversation.ID,
			UserID:         testUser.ID,
			Content:        "Second message",
			MessageType:    MessageTypeAssistant,
		},
	}

	messageIDs := make([]int64, 0)
	for _, message := range testMessages {
		id, err := messageModel.Create(ctx, message)
		if err != nil {
			t.Fatalf("Failed to create test message: %v", err)
		}
		messageIDs = append(messageIDs, id)
	}

	// Verify messages exist
	messages, err := messageModel.ListByConversation(ctx, testConversation.ID)
	if err != nil {
		t.Fatalf("Failed to list messages: %v", err)
	}
	if len(messages) != 2 {
		t.Fatalf("Expected 2 messages, got %v", len(messages))
	}

	// Delete the conversation (should cascade delete messages)
	err = conversationModel.Delete(ctx, testConversation.ID)
	if err != nil {
		t.Fatalf("Failed to delete conversation: %v", err)
	}

	// Verify messages were cascade deleted
	for _, messageID := range messageIDs {
		_, err = messageModel.Get(ctx, messageID)
		if err == nil {
			t.Errorf("Message %v still exists after conversation deletion", messageID)
		}
	}
}
