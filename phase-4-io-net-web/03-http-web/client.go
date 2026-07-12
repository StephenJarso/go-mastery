package httpweb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Go's default http.Client is shared and safe for concurrent use.
// However, in production, NEVER use the default client (http.DefaultClient)
// because it has no timeout, meaning slow servers can hang your application forever.
// Always instantiate a custom client with a Timeout.

func NewHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second, // Configure request timeouts
	}
}

// FetchHome retrieves content from the home endpoint.
func FetchHome(url string) (string, error) {
	client := NewHTTPClient()

	// Perform a simple GET request
	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("GET request failed: %w", err)
	}
	// CRITICAL: Always close the response Body when you are done reading it.
	// Failing to close it will leak connections and file descriptors.
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status: %s", resp.Status)
	}

	// Read the entire body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(bodyBytes), nil
}

// CreateUserRequest demonstrates making a POST request with JSON payload and custom headers.
func CreateUserRequest(url string, name string) (*User, error) {
	client := NewHTTPClient()

	payload := User{ID: "10", Name: name}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user payload: %w", err)
	}

	// For custom methods, headers, or body inputs, use http.NewRequest.
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("failed to create POST request: %w", err)
	}

	// Set headers on the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Go-Mastery-Client/1.0")

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request execution failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create user, status: %s", resp.Status)
	}

	var createdUser User
	err = json.NewDecoder(resp.Body).Decode(&createdUser)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &createdUser, nil
}
