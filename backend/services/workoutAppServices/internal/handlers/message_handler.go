package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	api_models "workout_app_backend/internal/generated"
	models "workout_app_backend/internal/models"
	services "workout_app_backend/internal/services"
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

// Domain to API model conversion functions
func convertMessageToAPI(internal *models.Message) api_models.Message {
	return api_models.Message{
		Id:             &internal.ID,
		ConversationId: &internal.ConversationID,
		UserId:         &internal.UserID,
		Content:        internal.Content,
		MessageType:    api_models.MessageType(internal.MessageType),
		CreatedAt:      &internal.CreatedAt,
	}
}

func convertMessagesToAPI(internals []*models.Message) []api_models.Message {
	result := make([]api_models.Message, len(internals))
	for i, internal := range internals {
		result[i] = convertMessageToAPI(internal)
	}
	return result
}

func convertAPIMessageType(apiType api_models.MessageType) models.MessageType {
	switch apiType {
	case api_models.USER:
		return models.MessageTypeUser
	case api_models.ASSISTANT:
		return models.MessageTypeAssistant
	case api_models.SYSTEM:
		return models.MessageTypeSystem
	default:
		return models.MessageTypeUser // Default fallback
	}
}

// ListMessagesByConversation handles GET /api/users/{userID}/conversations/{conversationID}/messages
func (h *MessageHandler) ListMessagesByConversation(w http.ResponseWriter, r *http.Request) {
	handlerLogger.Println("ListMessagesByConversation request received")

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

	// Convert to API models and wrap in response
	response := api_models.ListMessagesResponse{
		Messages: convertMessagesToAPI(messages),
	}

	respondWithJSON(w, http.StatusOK, response)
}

// CreateMessage handles POST /api/users/{userID}/conversations/{conversationID}/messages
func (h *MessageHandler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	handlerLogger.Println("CreateMessage request received")
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
	var request api_models.CreateMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	handlerLogger.Println("📝 CHECKPOINT 6: Create Message request received:", request)
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

	// Convert API model to internal domain model
	message := &models.Message{
		ConversationID: conversationID,
		UserID:         userID,
		Content:        request.Content,
		MessageType:    convertAPIMessageType(request.MessageType),
	}

	// Create the user message
	id, err := h.messageModel.Create(ctx, message)
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

	handlerLogger.Println("Checking if the message is a user message and if the LLM client is not nil")

	// If this is a user message, trigger AI response
	if message.MessageType == models.MessageTypeUser && h.llmClient != nil {
		handlerLogger.Println("Message is a user message and LLM client is not nil, so we will process the AI response")
		go h.processAIResponse(context.Background(), userID, conversationID)
	}

	// Convert to API model
	response := convertMessageToAPI(createdMessage)
	respondWithJSON(w, http.StatusCreated, response)
}

// processAIResponse handles the AI response generation in a separate goroutine
func (h *MessageHandler) processAIResponse(ctx context.Context, userID, conversationID int64) {
	handlerLogger.Println("processAIResponse request received")
	// Get conversation history
	messagePointers, err := h.messageModel.ListByConversation(ctx, conversationID)

	// Convert []*Message to []Message
	messages := make([]models.Message, len(messagePointers))
	for i, msgPtr := range messagePointers {
		messages[i] = *msgPtr
	}

	handlerLogger.Println("Conversation history retrieved", messages)
	if err != nil {
		handlerLogger.Println("Error getting conversation history", err)
		// Log error but don't fail the original request
		return
	}

	// Call LLM service
	llmResponse, err := h.llmClient.ProcessChatMessage(ctx, convertMessagesToAPI(messagePointers), userID, conversationID, nil)
	if err != nil {
		handlerLogger.Println("Error processing AI response", err)
		// Log error but don't fail the original request
		return
	}

	handlerLogger.Println("LLM response: ", llmResponse)

	// Create assistant message with the AI response
	assistantMessage := &models.Message{
		ConversationID: conversationID,
		UserID:         userID,
		Content:        llmResponse.Message.Content,
		MessageType:    models.MessageTypeAssistant,
	}

	handlerLogger.Println("Assistant message from the handler: ", assistantMessage)

	// Save the assistant message
	_, err = h.messageModel.Create(ctx, assistantMessage)
	if err != nil {
		handlerLogger.Println("Error creating assistant message", err)
		// Log error but don't fail the original request
		return
	}

	// TODO: You could emit a WebSocket event here to notify the frontend
	// that a new AI response is available
}

// GetMessage handles GET /api/users/{userID}/conversations/{conversationID}/messages/{messageID}
func (h *MessageHandler) GetMessage(w http.ResponseWriter, r *http.Request) {
	handlerLogger.Println("GetMessage request received")

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

	messageID, err := parseIDFromURL(r, "messageID")
	if err != nil {
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

	// Convert to API model
	response := convertMessageToAPI(message)
	respondWithJSON(w, http.StatusOK, response)
}

// UpdateMessage handles PUT /api/users/{userID}/conversations/{conversationID}/messages/{messageID}
func (h *MessageHandler) UpdateMessage(w http.ResponseWriter, r *http.Request) {
	handlerLogger.Println("UpdateMessage request received")

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

	messageID, err := parseIDFromURL(r, "messageID")
	if err != nil {
		respondWithError(w, "Invalid message ID", http.StatusBadRequest)
		return
	}

	// Parse request body
	var request api_models.Message // Assuming we use the full Message model for updates
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respondWithError(w, "Invalid request body", http.StatusBadRequest)
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

	// Convert API model to internal domain model
	message := &models.Message{
		ID:             messageID,
		ConversationID: conversationID,
		UserID:         userID,
		Content:        request.Content,
		MessageType:    convertAPIMessageType(request.MessageType),
	}

	// Update the message
	if err := h.messageModel.Update(ctx, message); err != nil {
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

	// Convert to API model
	response := convertMessageToAPI(updatedMessage)
	respondWithJSON(w, http.StatusOK, response)
}

// DeleteMessage handles DELETE /api/users/{userID}/conversations/{conversationID}/messages/{messageID}
func (h *MessageHandler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	handlerLogger.Println("DeleteMessage request received")

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

	messageID, err := parseIDFromURL(r, "messageID")
	if err != nil {
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
