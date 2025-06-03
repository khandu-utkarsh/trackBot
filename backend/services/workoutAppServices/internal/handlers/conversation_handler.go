package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	models "workout_app_backend/internal/models"

	"github.com/go-chi/chi/v5"
)

// ConversationHandler handles conversation-related HTTP requests
type ConversationHandler struct {
	conversationModel *models.ConversationModel
	userModel         *models.UserModel
}

// NewConversationHandler creates a new conversation handler
func NewConversationHandler(conversationModel *models.ConversationModel, userModel *models.UserModel) *ConversationHandler {
	return &ConversationHandler{
		conversationModel: conversationModel,
		userModel:         userModel,
	}
}

// ListConversationsByUser handles GET /api/users/{userID}/conversations
func (h *ConversationHandler) ListConversationsByUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ListConversationsByUser called")
	if r.Method != http.MethodGet {
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	// Verify user exists
	_, err = h.userModel.Get(ctx, userID)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			respondWithError(w, "User not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to verify user", http.StatusInternalServerError)
		return
	}

	// Get conversations for the user
	conversations, err := h.conversationModel.ListByUser(ctx, userID)
	if err != nil {
		respondWithError(w, "Failed to list conversations", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, conversations)
}

// CreateConversation handles POST /api/users/{userID}/conversations
func (h *ConversationHandler) CreateConversation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var conversation models.Conversation
	if err := json.NewDecoder(r.Body).Decode(&conversation); err != nil {
		respondWithError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set the user ID from the URL path
	conversation.UserID = userID

	// Set default values
	conversation.IsActive = true

	ctx := r.Context()

	// Verify user exists
	_, err = h.userModel.Get(ctx, userID)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			respondWithError(w, "User not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to verify user", http.StatusInternalServerError)
		return
	}

	// Create the conversation
	id, err := h.conversationModel.Create(ctx, &conversation)
	if err != nil {
		respondWithError(w, "Failed to create conversation", http.StatusInternalServerError)
		return
	}

	// Get the created conversation to return complete data
	createdConversation, err := h.conversationModel.Get(ctx, id)
	if err != nil {
		respondWithError(w, "Failed to get created conversation", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, createdConversation)
}

// GetConversation handles GET /api/users/{userID}/conversations/{conversationID}
func (h *ConversationHandler) GetConversation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	conversationIDStr := chi.URLParam(r, "conversationID")
	conversationID, err := strconv.ParseInt(conversationIDStr, 10, 64)
	if err != nil || conversationID <= 0 {
		respondWithError(w, "Invalid conversation ID", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	// Get the conversation
	conversation, err := h.conversationModel.Get(ctx, conversationID)
	if err != nil {
		if errors.Is(err, models.ErrConversationNotFound) {
			respondWithError(w, "Conversation not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to get conversation", http.StatusInternalServerError)
		return
	}

	// Verify the conversation belongs to the user
	if conversation.UserID != userID {
		respondWithError(w, "Conversation not found", http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, conversation)
}

// UpdateConversation handles PUT /api/users/{userID}/conversations/{conversationID}
func (h *ConversationHandler) UpdateConversation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	conversationIDStr := chi.URLParam(r, "conversationID")
	conversationID, err := strconv.ParseInt(conversationIDStr, 10, 64)
	if err != nil || conversationID <= 0 {
		respondWithError(w, "Invalid conversation ID", http.StatusBadRequest)
		return
	}

	var conversation models.Conversation
	if err := json.NewDecoder(r.Body).Decode(&conversation); err != nil {
		respondWithError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set the IDs from the URL path
	conversation.ID = conversationID
	conversation.UserID = userID

	ctx := r.Context()

	// Verify the conversation exists and belongs to the user
	existingConversation, err := h.conversationModel.Get(ctx, conversationID)
	if err != nil {
		if errors.Is(err, models.ErrConversationNotFound) {
			respondWithError(w, "Conversation not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to get conversation", http.StatusInternalServerError)
		return
	}

	if existingConversation.UserID != userID {
		respondWithError(w, "Conversation not found", http.StatusNotFound)
		return
	}

	// Update the conversation
	if err := h.conversationModel.Update(ctx, &conversation); err != nil {
		if errors.Is(err, models.ErrConversationNotFound) {
			respondWithError(w, "Conversation not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to update conversation", http.StatusInternalServerError)
		return
	}

	// Get the updated conversation to return complete data
	updatedConversation, err := h.conversationModel.Get(ctx, conversationID)
	if err != nil {
		respondWithError(w, "Failed to get updated conversation", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedConversation)
}

// DeleteConversation handles DELETE /api/users/{userID}/conversations/{conversationID}
func (h *ConversationHandler) DeleteConversation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userIDStr := chi.URLParam(r, "userID")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	conversationIDStr := chi.URLParam(r, "conversationID")
	conversationID, err := strconv.ParseInt(conversationIDStr, 10, 64)
	if err != nil || conversationID <= 0 {
		respondWithError(w, "Invalid conversation ID", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	// Verify the conversation exists and belongs to the user
	conversation, err := h.conversationModel.Get(ctx, conversationID)
	if err != nil {
		if errors.Is(err, models.ErrConversationNotFound) {
			respondWithError(w, "Conversation not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to get conversation", http.StatusInternalServerError)
		return
	}

	if conversation.UserID != userID {
		respondWithError(w, "Conversation not found", http.StatusNotFound)
		return
	}

	// Delete the conversation (cascade will delete messages)
	if err := h.conversationModel.Delete(ctx, conversationID); err != nil {
		if errors.Is(err, models.ErrConversationNotFound) {
			respondWithError(w, "Conversation not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to delete conversation", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]int64{"deleted_id": conversationID})
}
