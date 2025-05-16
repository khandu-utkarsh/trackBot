package main

import (
	"database/sql"
	"log"
	database "workout_app_backend/internal/database/init"
	models "workout_app_backend/internal/models"
)

func SetupDB() (*sql.DB, error) {

	db, err := database.GetInstance()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}
	log.Println("Database connected successfully")

	//Check and initalize databse models if needed

	// Initialize models
	log.Println("Initializing models")
	if err := models.InitializeModels(db); err != nil {
		log.Fatalf("Failed to initialize models: %v", err)
	}

	return db, nil
}
