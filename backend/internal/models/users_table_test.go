package models

import (
	"context"
	"fmt"
	"testing"
	"workout_app_backend/internal/testutils"
)

func TestUserModel_Create(t *testing.T) {
	//	t.Skip("Skipping this test temporarily")
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()
	fmt.Println("db", db)

	ctx := context.Background()
	userModel := GetUserModelInstance(db, "users_test")

	// Initialize the users table
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	// Test cases
	tests := []struct {
		name    string
		user    *User
		wantErr bool
	}{
		{
			name: "Valid user",
			user: &User{
				Email: "test@example.com",
			},
			wantErr: false,
		},
		{
			name: "Duplicate email",
			user: &User{
				Email: "test@example.com",
			},
			wantErr: true,
		},
		{
			name: "Empty email",
			user: &User{
				Email: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := userModel.Create(ctx, tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserModel.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserModel_Get(t *testing.T) {
	//t.Skip("Skipping this test temporarily")
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()
	userModel := GetUserModelInstance(db, "users_test")

	// Initialize the users table
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	// Create a test user
	testUser := &User{
		Email: "test@example.com",
	}
	id, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}
	testUser.ID = id

	// Test cases
	tests := []struct {
		name    string
		id      int64
		wantErr bool
	}{
		{
			name:    "Existing user",
			id:      testUser.ID,
			wantErr: false,
		},
		{
			name:    "Non-existent user",
			id:      999,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := userModel.Get(ctx, tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserModel.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && user == nil {
				t.Error("UserModel.Get() returned nil user when error was expected")
			}
		})
	}
}

func TestUserModel_List(t *testing.T) {
	db, cleanup := testutils.SetupTestDB(t)
	defer cleanup()

	ctx := context.Background()
	userModel := GetUserModelInstance(db, "users_test")

	// Initialize the users table
	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	// Create multiple test users
	users := []*User{
		{Email: "user1@example.com"},
		{Email: "user2@example.com"},
		{Email: "user3@example.com"},
	}

	for _, user := range users {
		_, err := userModel.Create(ctx, user)
		if err != nil {
			t.Fatalf("Failed to create test user: %v", err)
		}
	}

	// Test listing users
	list, err := userModel.List(ctx)
	if err != nil {
		t.Fatalf("UserModel.List() error = %v", err)
	}

	if len(list) != len(users) {
		t.Errorf("UserModel.List() returned %d users, want %d", len(list), len(users))
	}
}
