package main

import (
	"database/sql"
	database "workout_app_backend/internal/database/init"
	models "workout_app_backend/internal/models"
)

func SetupDB() (*sql.DB, error) {

	db, err := database.GetInstance()
	if err != nil {
		mainLogger.Fatalf("Failed to get database instance: %v", err)
	}

	if err := db.Ping(); err != nil {
		mainLogger.Fatalf("Database ping failed: %v", err)
	}
	mainLogger.Println("Database connected successfully")

	//Check and initalize databse models if needed

	// Initialize models
	mainLogger.Println("Initializing models")
	if err := models.InitializeModels(db); err != nil {
		mainLogger.Fatalf("Failed to initialize models: %v", err)
	}

	return db, nil
}
