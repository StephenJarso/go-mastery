package httpweb

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTokenAuthMiddleware(t *testing.T) {
	secret := "super-secret-token"
	authMiddleware := TokenAuthMiddleware(secret)

	// Create a dummy final handler to verify authorization succeeded
	finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Authorized Success"))
	})

	handlerToTest := authMiddleware(finalHandler)

	// Case 1: No Authorization Header
	req1 := httptest.NewRequest("GET", "/api/data", nil)
	rr1 := httptest.NewRecorder()
	handlerToTest.ServeHTTP(rr1, req1)

	if rr1.Code != http.StatusUnauthorized {
		t.Errorf("expected status %d, got %d", http.StatusUnauthorized, rr1.Code)
	}

	var errResp1 map[string]string
	json.NewDecoder(rr1.Body).Decode(&errResp1)
	if errResp1["error"] != "Unauthorized: invalid token" {
		t.Errorf("unexpected error message: %v", errResp1)
	}

	// Case 2: Wrong Token
	req2 := httptest.NewRequest("GET", "/api/data", nil)
	req2.Header.Set("Authorization", "Bearer wrong-token")
	rr2 := httptest.NewRecorder()
	handlerToTest.ServeHTTP(rr2, req2)

	if rr2.Code != http.StatusUnauthorized {
		t.Errorf("expected status %d, got %d", http.StatusUnauthorized, rr2.Code)
	}

	// Case 3: Correct Token
	req3 := httptest.NewRequest("GET", "/api/data", nil)
	req3.Header.Set("Authorization", "Bearer "+secret)
	rr3 := httptest.NewRecorder()
	handlerToTest.ServeHTTP(rr3, req3)

	if rr3.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr3.Code)
	}
	if rr3.Body.String() != "Authorized Success" {
		t.Errorf("expected body 'Authorized Success', got %q", rr3.Body.String())
	}
}

func TestRequestIDMiddleware(t *testing.T) {
	mockID := "req-12345"
	generator := func() string { return mockID }

	requestIDMiddleware := RequestIDMiddleware(generator)

	var retrievedID string
	finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the value from context
		val, ok := r.Context().Value(RequestIDKey).(string)
		if ok {
			retrievedID = val
		}
		w.WriteHeader(http.StatusOK)
	})

	handlerToTest := requestIDMiddleware(finalHandler)

	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	handlerToTest.ServeHTTP(rr, req)

	// Check response header
	respHeaderID := rr.Header().Get("X-Request-ID")
	if respHeaderID != mockID {
		t.Errorf("expected X-Request-ID response header to be %q, got %q", mockID, respHeaderID)
	}

	// Check context retrieval
	if retrievedID != mockID {
		t.Errorf("expected request context ID to be %q, got %q", mockID, retrievedID)
	}
}
