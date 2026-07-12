package httpweb

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPServerAndClient(t *testing.T) {
	// 1. Setup routes
	mux := SetupRoutes()

	// 2. Wrap routes with LoggingMiddleware and RecoveryMiddleware
	// This mirrors how a production server is started.
	wrappedHandler := RecoveryMiddleware(LoggingMiddleware(mux))

	// 3. httptest.NewServer spins up a local server on a random loopback port.
	// This is perfect for integration testing HTTP servers and clients.
	server := httptest.NewServer(wrappedHandler)
	defer server.Close() // Ensure the test server is shut down afterwards

	// Test GET Home using client
	homeContent, err := FetchHome(server.URL + "/")
	if err != nil {
		t.Fatalf("FetchHome failed: %v", err)
	}
	expectedHome := "Welcome to Go Web Development!"
	if homeContent != expectedHome {
		t.Errorf("expected home content %q, got %q", expectedHome, homeContent)
	}

	// Test GET Users
	client := NewHTTPClient()
	resp, err := client.Get(server.URL + "/api/users")
	if err != nil {
		t.Fatalf("failed to GET users: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 OK, got %s", resp.Status)
	}

	// Test GET Wildcard User
	resp2, err := client.Get(server.URL + "/api/users/1")
	if err != nil {
		t.Fatalf("failed to GET user 1: %v", err)
	}
	defer resp2.Body.Close()
	if resp2.StatusCode != http.StatusOK {
		t.Errorf("expected 200 OK, got %s", resp2.Status)
	}

	// Test POST CreateUser using client
	createdUser, err := CreateUserRequest(server.URL+"/api/users", "StephenJarso")
	if err != nil {
		t.Fatalf("CreateUserRequest failed: %v", err)
	}
	if createdUser.Name != "StephenJarso" {
		t.Errorf("expected created user name 'StephenJarso', got %q", createdUser.Name)
	}
}

func TestRecoveryMiddleware(t *testing.T) {
	// Handler that panics
	panickyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("something went horribly wrong")
	})

	recoveredHandler := RecoveryMiddleware(panickyHandler)

	req := httptest.NewRequest("GET", "/panic", nil)
	rr := httptest.NewRecorder()

	// Call the handler. It should NOT crash the test suite!
	recoveredHandler.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500 Internal Server Error, got %d", rr.Code)
	}
}
