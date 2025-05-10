package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"workout_app_backend/services/workoutAppServices/internal/handlers"
	"workout_app_backend/services/workoutAppServices/internal/models"
	"workout_app_backend/services/workoutAppServices/internal/routes"
	"workout_app_backend/services/workoutAppServices/internal/utils"
)

func main() {

	//!Loading all the needed environment variables
	utils.LoadEnv()

	db, err := SetupDB()
	if err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}

	// Setup Router using the routes package

	userModel := models.GetUserModelInstance(db, "users")
	workoutModel := models.GetWorkoutModelInstance(db, "workouts", "users")
	exerciseModel := models.GetExerciseModelInstance(db, "exercises", "workouts")

	r := routes.SetupRouter(handlers.GetUserHandlerInstance(userModel),
		handlers.GetWorkoutHandlerInstance(workoutModel, userModel),
		handlers.GetExerciseHandlerInstance(exerciseModel, workoutModel))

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
