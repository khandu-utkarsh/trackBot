package main

import (
	"log"
	"net/http"
	"os"
	"workout_app_backend/internal/database"
	"workout_app_backend/internal/handlers"
	"workout_app_backend/internal/models"
	"workout_app_backend/internal/routes"
	"workout_app_backend/internal/utils"
)

func main() {
	utils.LoadEnv()

	db, err := database.GetInstance()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}
	log.Println("Database connected successfully")

	// Initialize models
	log.Println("Initializing models")
	if err := models.InitializeModels(db); err != nil {
		log.Fatalf("Failed to initialize models: %v", err)
	}

	// Setup Router using the routes package
	log.Println("Setting up router...")
	r := routes.SetupRouter(handlers.GetUserHandlerInstance(models.GetUserModelInstance(db, "users")),
		handlers.GetWorkoutHandlerInstance(models.GetWorkoutModelInstance(db, "workouts", "users")),
		handlers.GetExerciseHandlerInstance(models.GetExerciseModelInstance(db, "exercises", "workouts")))

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
