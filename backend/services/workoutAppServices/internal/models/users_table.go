package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	database_utilities "workout_app_backend/services/workoutAppServices/internal/database/utils"
)

// Common errors
var (
	ErrUserNotFound   = errors.New("user not found")
	ErrInvalidInput   = errors.New("invalid input")
	ErrDuplicateEmail = errors.New("email already exists")
)

// User represents a user in the system.
// Matches the TypeScript interface.
type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserModel handles user-related database operations.
type UserModel struct {
	db   *sql.DB
	name string
}

// GetUserModelInstance creates a new UserModel instance.
func GetUserModelInstance(db *sql.DB, name string) *UserModel {
	return &UserModel{db: db, name: name}
}

// Initialize creates the users table if it doesn't exist.
func (m *UserModel) Initialize(ctx context.Context) error {
	schema := `
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL CHECK (email <> ''),
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	`
	return database_utilities.CreateTable(m.db, ctx, m.name, schema)
}

// validateUser checks if the user data is valid
func (m *UserModel) validateUser(user *User) error {
	if user == nil {
		return fmt.Errorf("%w: user cannot be nil", ErrInvalidInput)
	}
	if user.Email == "" {
		return fmt.Errorf("%w: email cannot be empty", ErrInvalidInput)
	}
	return nil
}

// scanUser scans a database row into a User struct
func (m *UserModel) scanUser(row *sql.Row) (*User, error) {
	var user User
	err := row.Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("error scanning user: %w", err)
	}
	return &user, nil
}

// Create inserts a new user into the database.
func (m *UserModel) Create(ctx context.Context, user *User) (int64, error) {
	if err := m.validateUser(user); err != nil {
		return 0, err
	}

	now := time.Now()
	query := fmt.Sprintf("INSERT INTO %s (email, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id", m.name)

	var id int64
	err := m.db.QueryRowContext(ctx, query, user.Email, now, now).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error creating user: %w", err)
	}

	return id, nil
}

// Get retrieves a user by ID.
func (m *UserModel) Get(ctx context.Context, id int64) (*User, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: invalid user ID", ErrInvalidInput)
	}

	query := fmt.Sprintf("SELECT id, email, created_at, updated_at FROM %s WHERE id = $1", m.name)
	return m.scanUser(m.db.QueryRowContext(ctx, query, id))
}

// GetByEmail retrieves a user by email.
func (m *UserModel) GetByEmail(ctx context.Context, email string) (*User, error) {
	if email == "" {
		return nil, fmt.Errorf("%w: email cannot be empty", ErrInvalidInput)
	}

	query := fmt.Sprintf("SELECT id, email, created_at, updated_at FROM %s WHERE email = $1", m.name)
	return m.scanUser(m.db.QueryRowContext(ctx, query, email))
}

// List retrieves all users.
func (m *UserModel) List(ctx context.Context) ([]*User, error) {
	query := fmt.Sprintf("SELECT id, email, created_at, updated_at FROM %s", m.name)
	rows, err := m.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying users: %w", err)
	}
	defer rows.Close()

	users := make([]*User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning user row: %w", err)
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating user rows: %w", err)
	}

	return users, nil
}

// Update updates an existing user.
func (m *UserModel) Update(ctx context.Context, user *User) error {
	if err := m.validateUser(user); err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET email = $1, updated_at = $2 WHERE id = $3", m.name)
	result, err := m.db.ExecContext(ctx, query, user.Email, time.Now(), user.ID)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}

// Delete removes a user from the database.
func (m *UserModel) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("%w: invalid user ID", ErrInvalidInput)
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", m.name)
	result, err := m.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}
