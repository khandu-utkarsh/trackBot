package models

import (
	"context"
	"time"
	"workout_app_backend/internal/database"
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
	db database.Database
}

// GetUserModelInstance creates a new UserModel instance.
func GetUserModelInstance(db database.Database) *UserModel {
	return &UserModel{db: db}
}

// Initialize creates the users table if it doesn't exist.
func (m *UserModel) Initialize(ctx context.Context) error {
	schema := `
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	`

	return m.db.CreateTable(ctx, "users", schema)
}

// Create inserts a new user into the database.
func (m *UserModel) Create(ctx context.Context, user *User) error {
	now := time.Now()
	_, err := m.db.ExecContext(ctx,
		"INSERT INTO users (email, created_at, updated_at) VALUES ($1, $2, $3)",
		user.Email, now, now)
	return err
}

// Get retrieves a user by ID.
func (m *UserModel) Get(ctx context.Context, id int64) (*User, error) {
	var user User
	err := m.db.QueryRowContext(ctx, "SELECT id, email, created_at, updated_at FROM users WHERE id = $1", id).Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}

// GetByEmail retrieves a user by email.
func (m *UserModel) GetByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err := m.db.QueryRowContext(ctx, "SELECT id, email, created_at, updated_at FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}

// List retrieves all users (consider pagination later).
func (m *UserModel) List(ctx context.Context) ([]*User, error) {
	rows, err := m.db.QueryContext(ctx, "SELECT id, email, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (m *UserModel) Update(ctx context.Context, user *User) error {
	_, err := m.db.ExecContext(ctx, "UPDATE users SET email = $1, updated_at = $2 WHERE id = $3", user.Email, time.Now(), user.ID)
	return err
}

func (m *UserModel) Delete(ctx context.Context, id int64) error {
	_, err := m.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", id)
	return err
}
