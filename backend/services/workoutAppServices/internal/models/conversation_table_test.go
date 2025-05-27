package models

import (
	"context"
	"strings"
	"testing"
	testutils "workout_app_backend/internal/testutils"
)

func setupTestConversationModel(t *testing.T) (*ConversationModel, *UserModel, *User, context.Context, func()) {
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
	id, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}
	testUser.ID = id

	// Initialize conversation model
	conversationModel := GetConversationModelInstance(db, "conversations_test", "users_test")
	if err := conversationModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize conversations table: %v", err)
	}

	return conversationModel, userModel, testUser, ctx, cleanup
}

func TestConversationModel_Create(t *testing.T) {
	conversationModel, _, testUser, ctx, cleanup := setupTestConversationModel(t)
	defer cleanup()

	tests := []struct {
		name         string
		conversation *Conversation
		wantErr      bool
		errContains  string
	}{
		{
			name: "valid conversation",
			conversation: &Conversation{
				UserID:   testUser.ID,
				Title:    "Test Conversation",
				IsActive: true,
			},
			wantErr: false,
		},
		{
			name: "conversation with empty title",
			conversation: &Conversation{
				UserID:   testUser.ID,
				Title:    "",
				IsActive: true,
			},
			wantErr:     true,
			errContains: "conversation title cannot be empty",
		},
		{
			name: "conversation with invalid user ID",
			conversation: &Conversation{
				UserID:   0,
				Title:    "Test Conversation",
				IsActive: true,
			},
			wantErr:     true,
			errContains: "invalid user ID",
		},
		{
			name:         "nil conversation",
			conversation: nil,
			wantErr:      true,
			errContains:  "conversation cannot be nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := conversationModel.Create(ctx, tt.conversation)

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

			// Verify the conversation was created
			created, err := conversationModel.Get(ctx, id)
			if err != nil {
				t.Errorf("Failed to get created conversation: %v", err)
				return
			}

			if created.UserID != tt.conversation.UserID {
				t.Errorf("Created conversation UserID = %v, want %v", created.UserID, tt.conversation.UserID)
			}
			if created.Title != tt.conversation.Title {
				t.Errorf("Created conversation Title = %v, want %v", created.Title, tt.conversation.Title)
			}
			if created.IsActive != tt.conversation.IsActive {
				t.Errorf("Created conversation IsActive = %v, want %v", created.IsActive, tt.conversation.IsActive)
			}
		})
	}
}

func TestConversationModel_Get(t *testing.T) {
	conversationModel, _, testUser, ctx, cleanup := setupTestConversationModel(t)
	defer cleanup()

	// Create a test conversation
	testConversation := &Conversation{
		UserID:   testUser.ID,
		Title:    "Test Conversation",
		IsActive: true,
	}
	id, err := conversationModel.Create(ctx, testConversation)
	if err != nil {
		t.Fatalf("Failed to create test conversation: %v", err)
	}

	tests := []struct {
		name        string
		id          int64
		wantErr     bool
		errContains string
	}{
		{
			name:    "valid conversation ID",
			id:      id,
			wantErr: false,
		},
		{
			name:        "invalid conversation ID",
			id:          0,
			wantErr:     true,
			errContains: "invalid conversation ID",
		},
		{
			name:        "non-existent conversation ID",
			id:          99999,
			wantErr:     true,
			errContains: "conversation not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conversation, err := conversationModel.Get(ctx, tt.id)

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

			if conversation.ID != tt.id {
				t.Errorf("Get() conversation ID = %v, want %v", conversation.ID, tt.id)
			}
			if conversation.UserID != testUser.ID {
				t.Errorf("Get() conversation UserID = %v, want %v", conversation.UserID, testUser.ID)
			}
			if conversation.Title != testConversation.Title {
				t.Errorf("Get() conversation Title = %v, want %v", conversation.Title, testConversation.Title)
			}
		})
	}
}

func TestConversationModel_ListByUser(t *testing.T) {
	conversationModel, _, testUser, ctx, cleanup := setupTestConversationModel(t)
	defer cleanup()

	// Create test conversations
	testConversations := []*Conversation{
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
		_, err := conversationModel.Create(ctx, conversation)
		if err != nil {
			t.Fatalf("Failed to create test conversation: %v", err)
		}
	}

	tests := []struct {
		name        string
		userID      int64
		wantCount   int
		wantErr     bool
		errContains string
	}{
		{
			name:      "valid user ID",
			userID:    testUser.ID,
			wantCount: 2,
			wantErr:   false,
		},
		{
			name:        "invalid user ID",
			userID:      0,
			wantErr:     true,
			errContains: "invalid user ID",
		},
		{
			name:      "non-existent user ID",
			userID:    99999,
			wantCount: 0,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conversations, err := conversationModel.ListByUser(ctx, tt.userID)

			if tt.wantErr {
				if err == nil {
					t.Errorf("ListByUser() error = nil, wantErr %v", tt.wantErr)
					return
				}
				if !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("ListByUser() error = %v, want error containing %v", err, tt.errContains)
				}
				return
			}

			if err != nil {
				t.Errorf("ListByUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(conversations) != tt.wantCount {
				t.Errorf("ListByUser() returned %v conversations, want %v", len(conversations), tt.wantCount)
			}

			// Verify conversations belong to the user
			for _, conversation := range conversations {
				if conversation.UserID != tt.userID {
					t.Errorf("ListByUser() conversation UserID = %v, want %v", conversation.UserID, tt.userID)
				}
			}
		})
	}
}

func TestConversationModel_Update(t *testing.T) {
	conversationModel, _, testUser, ctx, cleanup := setupTestConversationModel(t)
	defer cleanup()

	// Create a test conversation
	testConversation := &Conversation{
		UserID:   testUser.ID,
		Title:    "Original Title",
		IsActive: true,
	}
	id, err := conversationModel.Create(ctx, testConversation)
	if err != nil {
		t.Fatalf("Failed to create test conversation: %v", err)
	}
	testConversation.ID = id

	tests := []struct {
		name         string
		conversation *Conversation
		wantErr      bool
		errContains  string
	}{
		{
			name: "valid update",
			conversation: &Conversation{
				ID:       id,
				UserID:   testUser.ID,
				Title:    "Updated Title",
				IsActive: false,
			},
			wantErr: false,
		},
		{
			name: "update with empty title",
			conversation: &Conversation{
				ID:       id,
				UserID:   testUser.ID,
				Title:    "",
				IsActive: true,
			},
			wantErr:     true,
			errContains: "conversation title cannot be empty",
		},
		{
			name: "update non-existent conversation",
			conversation: &Conversation{
				ID:       99999,
				UserID:   testUser.ID,
				Title:    "Updated Title",
				IsActive: true,
			},
			wantErr:     true,
			errContains: "conversation not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := conversationModel.Update(ctx, tt.conversation)

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

			// Verify the conversation was updated
			updated, err := conversationModel.Get(ctx, tt.conversation.ID)
			if err != nil {
				t.Errorf("Failed to get updated conversation: %v", err)
				return
			}

			if updated.Title != tt.conversation.Title {
				t.Errorf("Updated conversation Title = %v, want %v", updated.Title, tt.conversation.Title)
			}
			if updated.IsActive != tt.conversation.IsActive {
				t.Errorf("Updated conversation IsActive = %v, want %v", updated.IsActive, tt.conversation.IsActive)
			}
		})
	}
}

func TestConversationModel_Delete(t *testing.T) {
	conversationModel, _, testUser, ctx, cleanup := setupTestConversationModel(t)
	defer cleanup()

	// Create a test conversation
	testConversation := &Conversation{
		UserID:   testUser.ID,
		Title:    "Test Conversation",
		IsActive: true,
	}
	id, err := conversationModel.Create(ctx, testConversation)
	if err != nil {
		t.Fatalf("Failed to create test conversation: %v", err)
	}

	tests := []struct {
		name        string
		id          int64
		wantErr     bool
		errContains string
	}{
		{
			name:    "valid conversation ID",
			id:      id,
			wantErr: false,
		},
		{
			name:        "invalid conversation ID",
			id:          0,
			wantErr:     true,
			errContains: "invalid conversation ID",
		},
		{
			name:        "non-existent conversation ID",
			id:          99999,
			wantErr:     true,
			errContains: "conversation not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := conversationModel.Delete(ctx, tt.id)

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

			// Verify the conversation was deleted
			_, err = conversationModel.Get(ctx, tt.id)
			if err == nil {
				t.Errorf("Delete() conversation still exists after deletion")
			}
		})
	}
}
