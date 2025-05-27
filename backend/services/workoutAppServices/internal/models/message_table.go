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
	ErrMessageNotFound = errors.New("message not found")
)

// MessageType represents the type of message
type MessageType string

const (
	MessageTypeUser      MessageType = "user"
	MessageTypeAssistant MessageType = "assistant"
	MessageTypeSystem    MessageType = "system"
)

// Message represents a single message in a conversation
type Message struct {
	ID             int64       `json:"id" db:"id"`
	ConversationID int64       `json:"conversation_id" db:"conversation_id"`
	UserID         int64       `json:"user_id" db:"user_id"`
	Content        string      `json:"content" db:"content"`
	MessageType    MessageType `json:"message_type" db:"message_type"`
	CreatedAt      time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at" db:"updated_at"`
}

// MessageModel handles message-related database operations
type MessageModel struct {
	db         *sql.DB
	name       string
	foreignKey string
}

// GetMessageModelInstance creates a new MessageModel instance
func GetMessageModelInstance(db *sql.DB, name string, foreignKey string) *MessageModel {
	return &MessageModel{db: db, name: name, foreignKey: foreignKey}
}

// Initialize creates the messages table if it doesn't exist
func (m *MessageModel) Initialize(ctx context.Context) error {
	schema := fmt.Sprintf(`
		id SERIAL PRIMARY KEY,
		conversation_id BIGINT NOT NULL REFERENCES %s(id) ON DELETE CASCADE,
		user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
		content TEXT NOT NULL,
		message_type VARCHAR(20) NOT NULL CHECK (message_type IN ('user', 'assistant', 'system')),
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	`, m.foreignKey)

	return database_utilities.CreateTable(m.db, ctx, m.name, schema)
}

// validateMessage checks if the message data is valid
func (m *MessageModel) validateMessage(message *Message) error {
	if message == nil {
		return fmt.Errorf("%w: message cannot be nil", ErrInvalidInput)
	}
	if message.ConversationID <= 0 {
		return fmt.Errorf("%w: invalid conversation ID", ErrInvalidInput)
	}
	if message.UserID <= 0 {
		return fmt.Errorf("%w: invalid user ID", ErrInvalidInput)
	}
	if message.Content == "" {
		return fmt.Errorf("%w: message content cannot be empty", ErrInvalidInput)
	}
	if message.MessageType != MessageTypeUser && message.MessageType != MessageTypeAssistant && message.MessageType != MessageTypeSystem {
		return fmt.Errorf("%w: invalid message type", ErrInvalidInput)
	}
	return nil
}

// scanMessage scans a database row into a Message struct
func (m *MessageModel) scanMessage(row *sql.Row) (*Message, error) {
	var message Message
	err := row.Scan(&message.ID, &message.ConversationID, &message.UserID, &message.Content, &message.MessageType, &message.CreatedAt, &message.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, ErrMessageNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("error scanning message: %w", err)
	}
	return &message, nil
}

// Create creates a new message
func (m *MessageModel) Create(ctx context.Context, message *Message) (int64, error) {
	if err := m.validateMessage(message); err != nil {
		return 0, err
	}

	now := time.Now()
	query := fmt.Sprintf("INSERT INTO %s (conversation_id, user_id, content, message_type, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", m.name)

	var id int64
	err := m.db.QueryRowContext(ctx, query, message.ConversationID, message.UserID, message.Content, message.MessageType, now, now).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error creating message: %w", err)
	}

	return id, nil
}

// Get retrieves a message by ID
func (m *MessageModel) Get(ctx context.Context, id int64) (*Message, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: invalid message ID", ErrInvalidInput)
	}

	query := fmt.Sprintf(`
		SELECT id, conversation_id, user_id, content, message_type, created_at, updated_at
		FROM %s
		WHERE id = $1
	`, m.name)

	return m.scanMessage(m.db.QueryRowContext(ctx, query, id))
}

// ListByConversation retrieves all messages for a conversation
func (m *MessageModel) ListByConversation(ctx context.Context, conversationID int64) ([]*Message, error) {
	if conversationID <= 0 {
		return nil, fmt.Errorf("%w: invalid conversation ID", ErrInvalidInput)
	}

	query := fmt.Sprintf(`
		SELECT id, conversation_id, user_id, content, message_type, created_at, updated_at
		FROM %s
		WHERE conversation_id = $1
		ORDER BY created_at ASC
	`, m.name)

	rows, err := m.db.QueryContext(ctx, query, conversationID)
	if err != nil {
		return nil, fmt.Errorf("error querying messages: %w", err)
	}
	defer rows.Close()

	messages := make([]*Message, 0)
	for rows.Next() {
		var message Message
		if err := rows.Scan(&message.ID, &message.ConversationID, &message.UserID, &message.Content, &message.MessageType, &message.CreatedAt, &message.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning message row: %w", err)
		}
		messages = append(messages, &message)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating message rows: %w", err)
	}

	return messages, nil
}

// Update updates an existing message
func (m *MessageModel) Update(ctx context.Context, message *Message) error {
	if err := m.validateMessage(message); err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET content = $1, updated_at = $2 WHERE id = $3", m.name)
	result, err := m.db.ExecContext(ctx, query, message.Content, time.Now(), message.ID)
	if err != nil {
		return fmt.Errorf("error updating message: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return ErrMessageNotFound
	}

	return nil
}

// Delete removes a message from the database
func (m *MessageModel) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("%w: invalid message ID", ErrInvalidInput)
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", m.name)
	result, err := m.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting message: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return ErrMessageNotFound
	}

	return nil
}
