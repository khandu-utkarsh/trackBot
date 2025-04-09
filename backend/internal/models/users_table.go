package models

import (
	"context"
	"fmt"
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
	db   database.Database
	name string
}

// GetUserModelInstance creates a new UserModel instance.
func GetUserModelInstance(db database.Database, name string) *UserModel {
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

	return m.db.CreateTable(ctx, m.name, schema)
}

// Create inserts a new user into the database.
func (m *UserModel) Create(ctx context.Context, user *User) (int64, error) {
	now := time.Now()

	query := fmt.Sprintf("INSERT INTO %s (email, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id", m.name)
	var id int64
	err := m.db.QueryRowContext(ctx, query, user.Email, now, now).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Get retrieves a user by ID.
func (m *UserModel) Get(ctx context.Context, id int64) (*User, error) {
	var user User
	query := fmt.Sprintf("SELECT id, email, created_at, updated_at FROM %s WHERE id = $1", m.name)
	err := m.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}

// GetByEmail retrieves a user by email.
func (m *UserModel) GetByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	query := fmt.Sprintf("SELECT id, email, created_at, updated_at FROM %s WHERE email = $1", m.name)
	err := m.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}

// List retrieves all users (consider pagination later).
func (m *UserModel) List(ctx context.Context) ([]*User, error) {
	query := fmt.Sprintf("SELECT id, email, created_at, updated_at FROM %s", m.name)
	rows, err := m.db.QueryContext(ctx, query)
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
	query := fmt.Sprintf("UPDATE %s SET email = $1, updated_at = $2 WHERE id = $3", m.name)
	_, err := m.db.ExecContext(ctx, query, user.Email, time.Now(), user.ID)
	return err
}

func (m *UserModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", m.name)
	_, err := m.db.ExecContext(ctx, query, id)
	return err
}
