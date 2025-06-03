package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

var handlerLogger *log.Logger

func init() {
	handlerLogger = log.New(os.Stdout, "Handler: ", log.LstdFlags)
}

func respondWithError(w http.ResponseWriter, message string, code int) {
	handlerLogger.Println("Responding with error: ", message) //! Logging the error.
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	handlerLogger.Println("Responding with JSON: ", payload) //! Logging the payload.
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Failed to marshal response"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
