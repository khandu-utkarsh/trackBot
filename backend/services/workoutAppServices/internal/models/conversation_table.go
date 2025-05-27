package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	database_utilities "workout_app_backend/internal/database/utils"
)

// Common errors
var (
	ErrConversationNotFound = errors.New("conversation not found")
)

// Conversation represents a chat conversation
type Conversation struct {
	ID          int64     `json:"id" db:"id"`
	UserID      int64     `json:"user_id" db:"user_id"`
	Title       string    `json:"title" db:"title"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	LastMessage *string   `json:"last_message,omitempty" db:"last_message"`
}

// ConversationModel handles conversation-related database operations
type ConversationModel struct {
	db         *sql.DB
	name       string
	foreignKey string
}

// GetConversationModelInstance creates a new ConversationModel instance
func GetConversationModelInstance(db *sql.DB, name string, foreignKey string) *ConversationModel {
	return &ConversationModel{db: db, name: name, foreignKey: foreignKey}
}

// Initialize creates the conversations table if it doesn't exist
func (m *ConversationModel) Initialize(ctx context.Context) error {
	schema := fmt.Sprintf(`
		id SERIAL PRIMARY KEY,
		user_id BIGINT NOT NULL REFERENCES %s(id) ON DELETE CASCADE,
		title VARCHAR(255) NOT NULL,
		is_active BOOLEAN DEFAULT true,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	`, m.foreignKey)

	return database_utilities.CreateTable(m.db, ctx, m.name, schema)
}

// validateConversation checks if the conversation data is valid
func (m *ConversationModel) validateConversation(conversation *Conversation) error {
	if conversation == nil {
		return fmt.Errorf("%w: conversation cannot be nil", ErrInvalidInput)
	}
	if conversation.UserID <= 0 {
		return fmt.Errorf("%w: invalid user ID", ErrInvalidInput)
	}
	if conversation.Title == "" {
		return fmt.Errorf("%w: conversation title cannot be empty", ErrInvalidInput)
	}
	return nil
}

// scanConversation scans a database row into a Conversation struct
func (m *ConversationModel) scanConversation(row *sql.Row) (*Conversation, error) {
	var conversation Conversation
	err := row.Scan(&conversation.ID, &conversation.UserID, &conversation.Title, &conversation.IsActive, &conversation.CreatedAt, &conversation.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, ErrConversationNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("error scanning conversation: %w", err)
	}
	return &conversation, nil
}

// Create creates a new conversation
func (m *ConversationModel) Create(ctx context.Context, conversation *Conversation) (int64, error) {
	if err := m.validateConversation(conversation); err != nil {
		return 0, err
	}

	now := time.Now()
	query := fmt.Sprintf("INSERT INTO %s (user_id, title, is_active, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id", m.name)

	var id int64
	err := m.db.QueryRowContext(ctx, query, conversation.UserID, conversation.Title, conversation.IsActive, now, now).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error creating conversation: %w", err)
	}

	return id, nil
}

// Get retrieves a conversation by ID
func (m *ConversationModel) Get(ctx context.Context, id int64) (*Conversation, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: invalid conversation ID", ErrInvalidInput)
	}

	query := fmt.Sprintf(`
		SELECT id, user_id, title, is_active, created_at, updated_at
		FROM %s
		WHERE id = $1
	`, m.name)

	return m.scanConversation(m.db.QueryRowContext(ctx, query, id))
}

// ListByUser retrieves all conversations for a user
func (m *ConversationModel) ListByUser(ctx context.Context, userID int64) ([]*Conversation, error) {
	if userID <= 0 {
		return nil, fmt.Errorf("%w: invalid user ID", ErrInvalidInput)
	}

	query := fmt.Sprintf(`
		SELECT id, user_id, title, is_active, created_at, updated_at
		FROM %s
		WHERE user_id = $1
		ORDER BY updated_at DESC
	`, m.name)

	rows, err := m.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("error querying conversations: %w", err)
	}
	defer rows.Close()

	conversations := make([]*Conversation, 0)
	for rows.Next() {
		var conversation Conversation
		if err := rows.Scan(&conversation.ID, &conversation.UserID, &conversation.Title, &conversation.IsActive, &conversation.CreatedAt, &conversation.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning conversation row: %w", err)
		}
		conversations = append(conversations, &conversation)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating conversation rows: %w", err)
	}

	return conversations, nil
}

// Update updates an existing conversation
func (m *ConversationModel) Update(ctx context.Context, conversation *Conversation) error {
	if err := m.validateConversation(conversation); err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET title = $1, is_active = $2, updated_at = $3 WHERE id = $4", m.name)
	result, err := m.db.ExecContext(ctx, query, conversation.Title, conversation.IsActive, time.Now(), conversation.ID)
	if err != nil {
		return fmt.Errorf("error updating conversation: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return ErrConversationNotFound
	}

	return nil
}

// Delete removes a conversation from the database
func (m *ConversationModel) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("%w: invalid conversation ID", ErrInvalidInput)
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", m.name)
	result, err := m.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting conversation: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return ErrConversationNotFound
	}

	return nil
}
