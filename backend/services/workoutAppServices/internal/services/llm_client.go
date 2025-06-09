package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	api_models "workout_app_backend/internal/generated"
)

// LLMClient handles communication with the LLM service
type LLMClient struct {
	baseURL    string
	httpClient *http.Client
}

// NewLLMClient creates a new LLM client
func NewLLMClient(baseURL string) *LLMClient {
	return &LLMClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 120 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 100,
				IdleConnTimeout:     90 * time.Second,
			},
		},
	}
}

// ProcessChatMessage sends a chat message to the LLM service and returns the response
func (c *LLMClient) ProcessChatMessage(ctx context.Context, messages []api_models.Message, userID, conversationID int64, context map[string]interface{}) (*api_models.LLMServiceMessageResponse, error) {

	// Create the request
	request := api_models.LLMServiceMessageRequest{
		Messages:       messages,
		UserId:         userID,
		ConversationId: conversationID,
	}

	// Marshal the request
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	url := fmt.Sprintf("%s/api/v1/process_conversation", c.baseURL)
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	fmt.Printf("Sending request to LLM service at %s\n", url)
	fmt.Printf("Request body: %s\n", string(requestBody))

	// Send the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		if err.Error() == "context deadline exceeded" {
			return nil, fmt.Errorf("request to LLM service timed out after 120 seconds. Please check if the service is running and accessible at %s", url)
		}
		return nil, fmt.Errorf("failed to send request to LLM service: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("LLM service returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	fmt.Printf("Received response from LLM service: %s\n", string(bodyBytes))

	// Parse the response
	var chatResponse api_models.LLMServiceMessageResponse
	if err := json.Unmarshal(bodyBytes, &chatResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &chatResponse, nil
}
