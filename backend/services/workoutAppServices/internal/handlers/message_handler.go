package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	models "workout_app_backend/internal/models"
	services "workout_app_backend/internal/services"

	"github.com/go-chi/chi/v5"
)

// MessageHandler handles message-related HTTP requests
type MessageHandler struct {
	messageModel      *models.MessageModel
	conversationModel *models.ConversationModel
	llmClient         *services.LLMClient
}

// NewMessageHandler creates a new message handler
func NewMessageHandler(messageModel *models.MessageModel, conversationModel *models.ConversationModel, llmClient *services.LLMClient) *MessageHandler {
	return &MessageHandler{
		messageModel:      messageModel,
		conversationModel: conversationModel,
		llmClient:         llmClient,
	}
}

// ListMessagesByConversation handles GET /api/users/{userID}/conversations/{conversationID}/messages
func (h *MessageHandler) ListMessagesByConversation(w http.ResponseWriter, r *http.Request) {
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

	// Get messages for the conversation
	messages, err := h.messageModel.ListByConversation(ctx, conversationID)
	if err != nil {
		respondWithError(w, "Failed to list messages", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, messages)
}

// CreateMessage handles POST /api/users/{userID}/conversations/{conversationID}/messages
func (h *MessageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
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

	conversationIDStr := chi.URLParam(r, "conversationID")
	conversationID, err := strconv.ParseInt(conversationIDStr, 10, 64)
	if err != nil || conversationID <= 0 {
		respondWithError(w, "Invalid conversation ID", http.StatusBadRequest)
		return
	}

	var message models.Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		respondWithError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set the IDs from the URL path
	message.ConversationID = conversationID
	message.UserID = userID

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

	// Create the user message
	id, err := h.messageModel.Create(ctx, &message)
	if err != nil {
		respondWithError(w, "Failed to create message", http.StatusInternalServerError)
		return
	}

	// Get the created message to return complete data
	createdMessage, err := h.messageModel.Get(ctx, id)
	if err != nil {
		respondWithError(w, "Failed to get created message", http.StatusInternalServerError)
		return
	}

	// If this is a user message, trigger AI response
	if message.MessageType == models.MessageTypeUser && h.llmClient != nil {
		go h.processAIResponse(ctx, userID, conversationID)
	}

	respondWithJSON(w, http.StatusCreated, createdMessage)
}

// processAIResponse handles the AI response generation in a separate goroutine
func (h *MessageHandler) processAIResponse(ctx context.Context, userID, conversationID int64) {
	// Get conversation history
	messagePointers, err := h.messageModel.ListByConversation(ctx, conversationID)
	if err != nil {
		// Log error but don't fail the original request
		return
	}

	// Convert []*Message to []Message
	messages := make([]models.Message, len(messagePointers))
	for i, msgPtr := range messagePointers {
		messages[i] = *msgPtr
	}

	// Call LLM service
	llmResponse, err := h.llmClient.ProcessChatMessage(ctx, messages, userID, conversationID, nil)
	if err != nil {
		// Log error but don't fail the original request
		return
	}

	// Create assistant message with the AI response
	assistantMessage := &models.Message{
		ConversationID: conversationID,
		UserID:         userID,
		Content:        llmResponse.Message,
		MessageType:    models.MessageTypeAssistant,
	}

	// Save the assistant message
	_, err = h.messageModel.Create(ctx, assistantMessage)
	if err != nil {
		// Log error but don't fail the original request
		return
	}

	// TODO: You could emit a WebSocket event here to notify the frontend
	// that a new AI response is available
}

// GetMessage handles GET /api/users/{userID}/conversations/{conversationID}/messages/{messageID}
func (h *MessageHandler) GetMessage(w http.ResponseWriter, r *http.Request) {
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

	messageIDStr := chi.URLParam(r, "messageID")
	messageID, err := strconv.ParseInt(messageIDStr, 10, 64)
	if err != nil || messageID <= 0 {
		respondWithError(w, "Invalid message ID", http.StatusBadRequest)
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

	// Get the message
	message, err := h.messageModel.Get(ctx, messageID)
	if err != nil {
		if errors.Is(err, models.ErrMessageNotFound) {
			respondWithError(w, "Message not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to get message", http.StatusInternalServerError)
		return
	}

	// Verify the message belongs to the conversation
	if message.ConversationID != conversationID {
		respondWithError(w, "Message not found", http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, message)
}

// UpdateMessage handles PUT /api/users/{userID}/conversations/{conversationID}/messages/{messageID}
func (h *MessageHandler) UpdateMessage(w http.ResponseWriter, r *http.Request) {
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

	messageIDStr := chi.URLParam(r, "messageID")
	messageID, err := strconv.ParseInt(messageIDStr, 10, 64)
	if err != nil || messageID <= 0 {
		respondWithError(w, "Invalid message ID", http.StatusBadRequest)
		return
	}

	var message models.Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		respondWithError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set the IDs from the URL path
	message.ID = messageID
	message.ConversationID = conversationID
	message.UserID = userID

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

	// Verify the message exists and belongs to the conversation
	existingMessage, err := h.messageModel.Get(ctx, messageID)
	if err != nil {
		if errors.Is(err, models.ErrMessageNotFound) {
			respondWithError(w, "Message not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to get message", http.StatusInternalServerError)
		return
	}

	if existingMessage.ConversationID != conversationID {
		respondWithError(w, "Message not found", http.StatusNotFound)
		return
	}

	// Update the message
	if err := h.messageModel.Update(ctx, &message); err != nil {
		if errors.Is(err, models.ErrMessageNotFound) {
			respondWithError(w, "Message not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to update message", http.StatusInternalServerError)
		return
	}

	// Get the updated message to return complete data
	updatedMessage, err := h.messageModel.Get(ctx, messageID)
	if err != nil {
		respondWithError(w, "Failed to get updated message", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedMessage)
}

// DeleteMessage handles DELETE /api/users/{userID}/conversations/{conversationID}/messages/{messageID}
func (h *MessageHandler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
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

	messageIDStr := chi.URLParam(r, "messageID")
	messageID, err := strconv.ParseInt(messageIDStr, 10, 64)
	if err != nil || messageID <= 0 {
		respondWithError(w, "Invalid message ID", http.StatusBadRequest)
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

	// Verify the message exists and belongs to the conversation
	message, err := h.messageModel.Get(ctx, messageID)
	if err != nil {
		if errors.Is(err, models.ErrMessageNotFound) {
			respondWithError(w, "Message not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to get message", http.StatusInternalServerError)
		return
	}

	if message.ConversationID != conversationID {
		respondWithError(w, "Message not found", http.StatusNotFound)
		return
	}

	// Delete the message
	if err := h.messageModel.Delete(ctx, messageID); err != nil {
		if errors.Is(err, models.ErrMessageNotFound) {
			respondWithError(w, "Message not found", http.StatusNotFound)
			return
		}
		respondWithError(w, "Failed to delete message", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]int64{"deleted_id": messageID})
}
