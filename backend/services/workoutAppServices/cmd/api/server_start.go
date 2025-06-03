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
	middleware "workout_app_backend/internal/middleware"
	models "workout_app_backend/internal/models"
	routes "workout_app_backend/internal/routes"
	services "workout_app_backend/internal/services"
	utils "workout_app_backend/internal/utils"
)

var mainLogger *log.Logger

func init() {
	utils.LoadEnv()
	mainLogger = log.New(os.Stdout, "Main: ", log.LstdFlags)
	mainLogger.Println("Main package initialized") //! Logging the initialization.
}

func main() {

	db, err := SetupDB()
	if err != nil {
		mainLogger.Println("Failed to setup database: ", err) //! Logging the error.
	}

	// Setup models
	userModel := models.GetUserModelInstance(db, "users")
	workoutModel := models.GetWorkoutModelInstance(db, "workouts", "users")
	exerciseModel := models.GetExerciseModelInstance(db, "exercises", "workouts")
	conversationModel := models.GetConversationModelInstance(db, "conversations", "users")
	messageModel := models.GetMessageModelInstance(db, "messages", "conversations")

	mainLogger.Println("Database setup complete") //! Logging the initialization.

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
	authMiddleware := middleware.NewAuthMiddleware()
	authHandler := handlers.GetAuthHandlerInstance(authMiddleware, userModel)

	mainLogger.Println("Handlers setup complete") //! Logging the initialization.

	// Setup Router using the routes package
	r := routes.SetupRouter(userHandler, workoutHandler, exerciseHandler, conversationHandler, messageHandler, authHandler)

	mainLogger.Println("Router setup complete") //! Logging the initialization.

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
		mainLogger.Printf("Server starting on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			mainLogger.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	mainLogger.Println("Shutting down server...")

	// Create a deadline for server shutdown once the quit signal is received
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() //! Defering the cancel.

	// Attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		mainLogger.Fatal("Server forced to shutdown:", err)
	}

	mainLogger.Println("Server exiting")
}
