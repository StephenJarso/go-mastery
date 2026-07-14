package advanced_testing

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserJSONHandler(t *testing.T) {
	t.Run("successful retrieval", func(t *testing.T) {
		mockUser := &User{ID: 42, Name: "Bob", Email: "bob@example.com"}
		mockStore := &MockUserStore{GetResult: mockUser}
		handler := UserJSONHandler(mockStore)

		// Create request: GET /user?id=42
		req := httptest.NewRequest("GET", "/user?id=42", nil)
		w := httptest.NewRecorder()

		// Call the handler directly
		handler.ServeHTTP(w, req)

		resp := w.Result()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status 200, got %d", resp.StatusCode)
		}

		if resp.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected Content-Type application/json, got %q", resp.Header.Get("Content-Type"))
		}

		var gotUser User
		err := json.NewDecoder(resp.Body).Decode(&gotUser)
		if err != nil {
			t.Fatalf("failed to decode response body: %v", err)
		}

		if gotUser.ID != 42 || gotUser.Name != "Bob" {
			t.Errorf("retrieved user mismatch: %+v", gotUser)
		}

		if mockStore.GetCalls != 1 {
			t.Errorf("expected 1 Get call, got %d", mockStore.GetCalls)
		}
	})

	t.Run("missing parameter", func(t *testing.T) {
		mockStore := &MockUserStore{}
		handler := UserJSONHandler(mockStore)

		req := httptest.NewRequest("GET", "/user", nil)
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		resp := w.Result()
		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", resp.StatusCode)
		}
	})

	t.Run("invalid parameter type", func(t *testing.T) {
		mockStore := &MockUserStore{}
		handler := UserJSONHandler(mockStore)

		req := httptest.NewRequest("GET", "/user?id=abc", nil)
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		resp := w.Result()
		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status 400, got %d", resp.StatusCode)
		}
	})

	t.Run("user not found", func(t *testing.T) {
		mockStore := &MockUserStore{GetError: errors.New("not found")}
		handler := UserJSONHandler(mockStore)

		req := httptest.NewRequest("GET", "/user?id=99", nil)
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		resp := w.Result()
		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("expected status 404, got %d", resp.StatusCode)
		}
	})
}
