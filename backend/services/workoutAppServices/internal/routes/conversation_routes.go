package routes

import (
	handlers "workout_app_backend/internal/handlers"

	"github.com/go-chi/chi/v5"
)

// RegisterConversationRoutes sets up the RESTful routes for conversation-related actions.
func RegisterConversationRoutes(r chi.Router, conversationHandler *handlers.ConversationHandler, messageHandler *handlers.MessageHandler) {
	r.Route("/users/{userID}/conversations", func(r chi.Router) {
		// List all conversations for a user
		r.Get("/", conversationHandler.ListConversationsByUser) // GET /api/users/{userID}/conversations

		// Create a new conversation
		r.Post("/", conversationHandler.CreateConversation) // POST /api/users/{userID}/conversations

		// Routes for a specific conversation
		r.Route("/{conversationID}", func(r chi.Router) {
			r.Get("/", conversationHandler.GetConversation)       // GET /api/users/{userID}/conversations/{conversationID}
			r.Put("/", conversationHandler.UpdateConversation)    // PUT /api/users/{userID}/conversations/{conversationID}
			r.Delete("/", conversationHandler.DeleteConversation) // DELETE /api/users/{userID}/conversations/{conversationID}

			// Message routes within a conversation
			RegisterMessageRoutes(r, messageHandler)
		})
	})
}

// RegisterMessageRoutes sets up routes for messages within a conversation
func RegisterMessageRoutes(r chi.Router, messageHandler *handlers.MessageHandler) {
	r.Route("/messages", func(r chi.Router) {
		// List all messages in a conversation
		r.Get("/", messageHandler.ListMessagesByConversation) // GET /api/users/{userID}/conversations/{conversationID}/messages

		// Create a new message in a conversation
		r.Post("/", messageHandler.CreateMessage) // POST /api/users/{userID}/conversations/{conversationID}/messages

		// Routes for a specific message
		r.Route("/{messageID}", func(r chi.Router) {
			r.Get("/", messageHandler.GetMessage)       // GET /api/users/{userID}/conversations/{conversationID}/messages/{messageID}
			r.Put("/", messageHandler.UpdateMessage)    // PUT /api/users/{userID}/conversations/{conversationID}/messages/{messageID}
			r.Delete("/", messageHandler.DeleteMessage) // DELETE /api/users/{userID}/conversations/{conversationID}/messages/{messageID}
		})
	})
}
