package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
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

// Common helper functions to reduce code duplication

// validateHTTPMethod checks if the request method matches the expected method
func validateHTTPMethod(r *http.Request, expectedMethod string) error {
	if r.Method != expectedMethod {
		return &HTTPError{
			Message: "Method not allowed",
			Code:    http.StatusMethodNotAllowed,
		}
	}
	return nil
}

// parseIDFromURL extracts and validates an ID parameter from URL
func parseIDFromURL(r *http.Request, paramName string) (int64, error) {
	idStr := chi.URLParam(r, paramName)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		return 0, &HTTPError{
			Message: "Invalid " + paramName,
			Code:    http.StatusBadRequest,
		}
	}
	return id, nil
}

// decodeJSONBody decodes JSON request body into the provided interface
func decodeJSONBody(r *http.Request, dest interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return &HTTPError{
			Message: "Invalid request body",
			Code:    http.StatusBadRequest,
		}
	}
	return nil
}

// HTTPError represents an HTTP error with message and status code
type HTTPError struct {
	Message string
	Code    int
}

func (e *HTTPError) Error() string {
	return e.Message
}

// handleHTTPError processes HTTPError and responds appropriately
func handleHTTPError(w http.ResponseWriter, err error) bool {
	if httpErr, ok := err.(*HTTPError); ok {
		respondWithError(w, httpErr.Message, httpErr.Code)
		return true
	}
	return false
}

// logRequest logs incoming requests with handler name
func logRequest(handlerName string) {
	handlerLogger.Printf("%s request received", handlerName)
}

// parseUserAndWorkoutIDs extracts and validates userID and workoutID from URL
func parseUserAndWorkoutIDs(r *http.Request) (userID, workoutID int64, err error) {
	userID, err = parseIDFromURL(r, "userID")
	if err != nil {
		return 0, 0, err
	}

	workoutID, err = parseIDFromURL(r, "workoutID")
	if err != nil {
		return 0, 0, err
	}

	return userID, workoutID, nil
}

// parseUserWorkoutAndExerciseIDs extracts and validates userID, workoutID, and exerciseID from URL
func parseUserWorkoutAndExerciseIDs(r *http.Request) (userID, workoutID, exerciseID int64, err error) {
	userID, workoutID, err = parseUserAndWorkoutIDs(r)
	if err != nil {
		return 0, 0, 0, err
	}

	exerciseID, err = parseIDFromURL(r, "exerciseID")
	if err != nil {
		return 0, 0, 0, err
	}

	return userID, workoutID, exerciseID, nil
}
