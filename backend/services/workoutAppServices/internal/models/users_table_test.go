package models

import (
	"context"
	"strings"
	"testing"
	"workout_app_backend/services/workoutAppServices/internal/testutils"
)

func setupTestUserModel(t *testing.T) (*UserModel, context.Context, func()) {
	db, cleanup := testutils.SetupTestDB(t)
	ctx := context.Background()
	userModel := GetUserModelInstance(db, "users_test")

	if err := userModel.Initialize(ctx); err != nil {
		t.Fatalf("Failed to initialize users table: %v", err)
	}

	return userModel, ctx, cleanup
}

func TestUserModel_Create(t *testing.T) {
	userModel, ctx, cleanup := setupTestUserModel(t)
	defer cleanup()

	tests := []struct {
		name    string
		user    *User
		wantErr bool
		errMsg  string
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
			errMsg:  "duplicate key",
		},
		{
			name: "Empty email",
			user: &User{
				Email: "",
			},
			wantErr: true,
			errMsg:  "email cannot be empty",
		},
		{
			name:    "Nil user",
			user:    nil,
			wantErr: true,
			errMsg:  "user cannot be nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := userModel.Create(ctx, tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserModel.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if id <= 0 {
					t.Error("UserModel.Create() returned invalid ID")
				}
				// Verify the user was created correctly
				user, err := userModel.Get(ctx, id)
				if err != nil {
					t.Errorf("Failed to get created user: %v", err)
				}
				if user.Email != tt.user.Email {
					t.Errorf("Created user email = %v, want %v", user.Email, tt.user.Email)
				}
			} else if tt.errMsg != "" && err != nil {
				// Check if the error message contains the expected message
				if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("UserModel.Create() error message = %v, want to contain %v", err, tt.errMsg)
				}
			}
		})
	}
}

func TestUserModel_Get(t *testing.T) {
	userModel, ctx, cleanup := setupTestUserModel(t)
	defer cleanup()

	// Create a test user
	testUser := &User{
		Email: "test@example.com",
	}
	id, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}
	testUser.ID = id

	tests := []struct {
		name    string
		id      int64
		want    *User
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Existing user",
			id:      testUser.ID,
			want:    testUser,
			wantErr: false,
		},
		{
			name:    "Non-existent user",
			id:      999,
			wantErr: true,
			errMsg:  "user not found",
		},
		{
			name:    "Invalid ID",
			id:      0,
			wantErr: true,
			errMsg:  "invalid user ID",
		},
		{
			name:    "Negative ID",
			id:      -1,
			wantErr: true,
			errMsg:  "invalid user ID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := userModel.Get(ctx, tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserModel.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if user == nil {
					t.Error("UserModel.Get() returned nil user")
					return
				}
				if user.ID != tt.want.ID {
					t.Errorf("UserModel.Get() ID = %v, want %v", user.ID, tt.want.ID)
				}
				if user.Email != tt.want.Email {
					t.Errorf("UserModel.Get() Email = %v, want %v", user.Email, tt.want.Email)
				}
			} else if tt.errMsg != "" && err != nil {
				// Check if the error message contains the expected message
				if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("UserModel.Get() error message = %v, want to contain %v", err, tt.errMsg)
				}
			}
		})
	}
}

func TestUserModel_GetByEmail(t *testing.T) {
	userModel, ctx, cleanup := setupTestUserModel(t)
	defer cleanup()

	// Create a test user
	testUser := &User{
		Email: "test@example.com",
	}
	id, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}
	testUser.ID = id

	tests := []struct {
		name    string
		email   string
		want    *User
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Existing email",
			email:   testUser.Email,
			want:    testUser,
			wantErr: false,
		},
		{
			name:    "Non-existent email",
			email:   "nonexistent@example.com",
			wantErr: true,
			errMsg:  "user not found",
		},
		{
			name:    "Empty email",
			email:   "",
			wantErr: true,
			errMsg:  "email cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := userModel.GetByEmail(ctx, tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserModel.GetByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if user == nil {
					t.Error("UserModel.GetByEmail() returned nil user")
					return
				}
				if user.ID != tt.want.ID {
					t.Errorf("UserModel.GetByEmail() ID = %v, want %v", user.ID, tt.want.ID)
				}
				if user.Email != tt.want.Email {
					t.Errorf("UserModel.GetByEmail() Email = %v, want %v", user.Email, tt.want.Email)
				}
			} else if tt.errMsg != "" && err != nil {
				// Check if the error message contains the expected message
				if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("UserModel.GetByEmail() error message = %v, want to contain %v", err, tt.errMsg)
				}
			}
		})
	}
}

func TestUserModel_List(t *testing.T) {
	userModel, ctx, cleanup := setupTestUserModel(t)
	defer cleanup()

	// Create multiple test users
	users := []*User{
		{Email: "user1@example.com"},
		{Email: "user2@example.com"},
		{Email: "user3@example.com"},
	}

	for _, user := range users {
		id, err := userModel.Create(ctx, user)
		if err != nil {
			t.Fatalf("Failed to create test user: %v", err)
		}
		user.ID = id
	}

	// Test listing users
	list, err := userModel.List(ctx)
	if err != nil {
		t.Fatalf("UserModel.List() error = %v", err)
	}

	if len(list) != len(users) {
		t.Errorf("UserModel.List() returned %d users, want %d", len(list), len(users))
	}

	// Verify each user's data
	userMap := make(map[string]*User)
	for _, user := range users {
		userMap[user.Email] = user
	}

	for _, user := range list {
		original, exists := userMap[user.Email]
		if !exists {
			t.Errorf("UserModel.List() returned unexpected user with email %s", user.Email)
			continue
		}
		if user.ID != original.ID {
			t.Errorf("UserModel.List() user ID = %v, want %v", user.ID, original.ID)
		}
	}
}

func TestUserModel_Update(t *testing.T) {
	userModel, ctx, cleanup := setupTestUserModel(t)
	defer cleanup()

	// Create a test user
	testUser := &User{
		Email: "test@example.com",
	}
	id, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}
	testUser.ID = id

	tests := []struct {
		name    string
		user    *User
		wantErr bool
		errMsg  string
	}{
		{
			name: "Valid update",
			user: &User{
				ID:    testUser.ID,
				Email: "updated@example.com",
			},
			wantErr: false,
		},
		{
			name: "Non-existent user",
			user: &User{
				ID:    999,
				Email: "nonexistent@example.com",
			},
			wantErr: true,
			errMsg:  "user not found",
		},
		{
			name: "Empty email",
			user: &User{
				ID:    testUser.ID,
				Email: "",
			},
			wantErr: true,
			errMsg:  "email cannot be empty",
		},
		{
			name:    "Nil user",
			user:    nil,
			wantErr: true,
			errMsg:  "user cannot be nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := userModel.Update(ctx, tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserModel.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// Verify the update
				updated, err := userModel.Get(ctx, tt.user.ID)
				if err != nil {
					t.Errorf("Failed to get updated user: %v", err)
					return
				}
				if updated.Email != tt.user.Email {
					t.Errorf("Updated user email = %v, want %v", updated.Email, tt.user.Email)
				}
			} else if tt.errMsg != "" && err != nil {
				// Check if the error message contains the expected message
				if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("UserModel.Update() error message = %v, want to contain %v", err, tt.errMsg)
				}
			}
		})
	}
}

func TestUserModel_Delete(t *testing.T) {
	userModel, ctx, cleanup := setupTestUserModel(t)
	defer cleanup()

	// Create a test user
	testUser := &User{
		Email: "test@example.com",
	}
	id, err := userModel.Create(ctx, testUser)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	tests := []struct {
		name    string
		id      int64
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Existing user",
			id:      id,
			wantErr: false,
		},
		{
			name:    "Non-existent user",
			id:      999,
			wantErr: true,
			errMsg:  "user not found",
		},
		{
			name:    "Invalid ID",
			id:      0,
			wantErr: true,
			errMsg:  "invalid user ID",
		},
		{
			name:    "Negative ID",
			id:      -1,
			wantErr: true,
			errMsg:  "invalid user ID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := userModel.Delete(ctx, tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserModel.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// Verify the deletion
				_, err := userModel.Get(ctx, tt.id)
				if err == nil {
					t.Error("UserModel.Delete() did not delete the user")
				} else if err != ErrUserNotFound {
					t.Errorf("Unexpected error after deletion: %v", err)
				}
			} else if tt.errMsg != "" && err != nil {
				// Check if the error message contains the expected message
				if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("UserModel.Delete() error message = %v, want to contain %v", err, tt.errMsg)
				}
			}
		})
	}
}
