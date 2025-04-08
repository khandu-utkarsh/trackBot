package main

import (
	"log"
	"net/http"
	"os"
	"workout_app_backend/internal/database"
	"workout_app_backend/internal/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("No .env file found")
	}

	//!Connect to database
	db, err := database.GetInstance()
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to ensure connection
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connected successfully")
		// Initialize models
		log.Println("Initializing models")
		models.InitializeModels(db)

	}

	// Initialize router
	r := mux.NewRouter()

	// Add CORS middleware
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	// Health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
