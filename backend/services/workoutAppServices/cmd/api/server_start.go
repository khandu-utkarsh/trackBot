package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	handlers "workout_app_backend/internal/handlers"
	models "workout_app_backend/internal/models"
	routes "workout_app_backend/internal/routes"
	services "workout_app_backend/internal/services"
	utils "workout_app_backend/internal/utils"
)

func main() {

	//!Loading all the needed environment variables
	utils.LoadEnv()

	db, err := SetupDB()
	if err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}

	// Setup models
	userModel := models.GetUserModelInstance(db, "users")
	workoutModel := models.GetWorkoutModelInstance(db, "workouts", "users")
	exerciseModel := models.GetExerciseModelInstance(db, "exercises", "workouts")
	conversationModel := models.GetConversationModelInstance(db, "conversations", "users")
	messageModel := models.GetMessageModelInstance(db, "messages", "conversations")

	// Setup LLM client
	llmServiceURL := os.Getenv("LLM_SERVICE_URL")
	if llmServiceURL == "" {
		llmServiceURL = "http://localhost:8081" // Default LLM service URL
	}
	llmClient := services.NewLLMClient(llmServiceURL)

	// Setup handlers
	userHandler := handlers.GetUserHandlerInstance(userModel)
	workoutHandler := handlers.GetWorkoutHandlerInstance(workoutModel, userModel)
	exerciseHandler := handlers.GetExerciseHandlerInstance(exerciseModel, workoutModel)
	conversationHandler := handlers.NewConversationHandler(conversationModel, userModel)
	messageHandler := handlers.NewMessageHandler(messageModel, conversationModel, llmClient)

	// Setup Router using the routes package
	r := routes.SetupRouter(userHandler, workoutHandler, exerciseHandler, conversationHandler, messageHandler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create a new server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
