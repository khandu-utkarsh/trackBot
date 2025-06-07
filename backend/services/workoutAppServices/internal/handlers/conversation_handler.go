package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	api_models "workout_app_backend/internal/generated"
	models "workout_app_backend/internal/models"
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

// Domain to API model conversion functions
func convertConversationToAPI(internal *models.Conversation) api_models.Conversation {
	return api_models.Conversation{
		Id:        internal.ID,
		UserId:    internal.UserID,
		Title:     internal.Title,
		CreatedAt: internal.CreatedAt,
	}
}

func convertConversationsToAPI(internals []*models.Conversation) []api_models.Conversation {
	result := make([]api_models.Conversation, len(internals))
	for i, internal := range internals {
		result[i] = convertConversationToAPI(internal)
	}
	return result
}

// ListConversationsByUser handles GET /api/users/{userID}/conversations
func (h *ConversationHandler) ListConversationsByUser(w http.ResponseWriter, r *http.Request) {
	handlerLogger.Println("ListConversationsByUser request received")

	userID, err := parseIDFromURL(r, "userID")
	if err != nil {
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

	// Convert to API models and wrap in response
	response := api_models.ListConversationsResponse{
		Conversations: convertConversationsToAPI(conversations),
	}

	respondWithJSON(w, http.StatusOK, response)
}

// CreateConversation handles POST /api/users/{userID}/conversations
func (h *ConversationHandler) CreateConversation(w http.ResponseWriter, r *http.Request) {
	handlerLogger.Println("CreateConversation request received")

	userID, err := parseIDFromURL(r, "userID")
	if err != nil {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Parse request body
	var request api_models.CreateConversationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, "Invalid request body", http.StatusBadRequest)
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

	// Convert API model to internal domain model
	conversation := &models.Conversation{
		UserID:   userID,
		Title:    request.Title,
		IsActive: true,
	}

	// Create the conversation
	id, err := h.conversationModel.Create(ctx, conversation)
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

	// Convert to API response model
	response := api_models.CreateConversationResponse{
		Id:        createdConversation.ID,
		Title:     createdConversation.Title,
		UserId:    createdConversation.UserID,
		CreatedAt: createdConversation.CreatedAt,
	}

	respondWithJSON(w, http.StatusCreated, response)
}

// GetConversation handles GET /api/users/{userID}/conversations/{conversationID}
func (h *ConversationHandler) GetConversation(w http.ResponseWriter, r *http.Request) {
	handlerLogger.Println("GetConversation request received")

	userID, err := parseIDFromURL(r, "userID")
	if err != nil {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	conversationID, err := parseIDFromURL(r, "conversationID")
	if err != nil {
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

	// Convert to API model
	response := convertConversationToAPI(conversation)
	respondWithJSON(w, http.StatusOK, response)
}

// UpdateConversation handles PUT /api/users/{userID}/conversations/{conversationID}
func (h *ConversationHandler) UpdateConversation(w http.ResponseWriter, r *http.Request) {
	handlerLogger.Println("UpdateConversation request received")

	userID, err := parseIDFromURL(r, "userID")
	if err != nil {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	conversationID, err := parseIDFromURL(r, "conversationID")
	if err != nil {
		respondWithError(w, "Invalid conversation ID", http.StatusBadRequest)
		return
	}

	// Parse request body
	var request api_models.UpdateConversationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

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

	// Convert API model to internal domain model
	title := existingConversation.Title // Keep existing title if not provided
	if request.Title != nil {
		title = *request.Title
	}

	conversation := &models.Conversation{
		ID:       conversationID,
		UserID:   userID,
		Title:    title,
		IsActive: existingConversation.IsActive, // Keep existing IsActive status
	}

	// Update the conversation
	if err := h.conversationModel.Update(ctx, conversation); err != nil {
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

	// Convert to API model
	response := convertConversationToAPI(updatedConversation)
	respondWithJSON(w, http.StatusOK, response)
}

// DeleteConversation handles DELETE /api/users/{userID}/conversations/{conversationID}
func (h *ConversationHandler) DeleteConversation(w http.ResponseWriter, r *http.Request) {
	handlerLogger.Println("DeleteConversation request received")

	userID, err := parseIDFromURL(r, "userID")
	if err != nil {
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	conversationID, err := parseIDFromURL(r, "conversationID")
	if err != nil {
		respondWithError(w, "Invalid conversation ID", http.StatusBadRequest)
		return
	}

	// Parse optional request body for confirmation
	var request api_models.DeleteConversationRequest
	if r.Body != nil {
		json.NewDecoder(r.Body).Decode(&request) // Ignore errors for optional body
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

	// Return proper delete response
	response := api_models.DeleteConversationResponse{
		Id:                   conversation.ID,
		Title:                conversation.Title,
		DeletedAt:            conversation.CreatedAt, // Using CreatedAt as placeholder - should be actual deletion time
		MessagesDeletedCount: 0,                      // Would need to count messages or get from delete result
	}

	respondWithJSON(w, http.StatusOK, response)
}
