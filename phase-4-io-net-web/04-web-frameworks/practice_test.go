package webframeworks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUnifiedAPIRoutes(t *testing.T) {
	engines := []string{"gin", "echo"}

	for _, engine := range engines {
		t.Run("Engine_"+engine, func(t *testing.T) {
			handler, err := SetupUnifiedAPIRoutes(engine)
			if err != nil {
				t.Fatalf("failed to setup unified routes: %v", err)
			}

			// 1. Test GET /ping
			reqPing := httptest.NewRequest("GET", "/ping", nil)
			rrPing := httptest.NewRecorder()
			handler.ServeHTTP(rrPing, reqPing)

			if rrPing.Code != http.StatusOK {
				t.Errorf("expected GET /ping status 200, got %d", rrPing.Code)
			}

			var pingRes map[string]string
			json.Unmarshal(rrPing.Body.Bytes(), &pingRes)
			if pingRes["message"] != "pong" {
				t.Errorf("expected pong response, got %v", pingRes)
			}

			// 2. Test GET /users/:id
			reqUser := httptest.NewRequest("GET", "/users/stephen", nil)
			rrUser := httptest.NewRecorder()
			handler.ServeHTTP(rrUser, reqUser)

			if rrUser.Code != http.StatusOK {
				t.Errorf("expected GET /users/stephen status 200, got %d", rrUser.Code)
			}

			var userRes map[string]string
			json.Unmarshal(rrUser.Body.Bytes(), &userRes)
			if userRes["user_id"] != "stephen" {
				t.Errorf("expected user_id 'stephen', got %v", userRes)
			}

			// 3. Test POST /users (Success)
			payload := `{"username": "StephenJarso"}`
			reqPostOk := httptest.NewRequest("POST", "/users", bytes.NewBufferString(payload))
			reqPostOk.Header.Set("Content-Type", "application/json")
			rrPostOk := httptest.NewRecorder()
			handler.ServeHTTP(rrPostOk, reqPostOk)

			if rrPostOk.Code != http.StatusCreated {
				t.Errorf("expected POST /users status 201, got %d", rrPostOk.Code)
			}

			var postOkRes map[string]string
			json.Unmarshal(rrPostOk.Body.Bytes(), &postOkRes)
			if postOkRes["status"] != "created" || postOkRes["username"] != "StephenJarso" {
				t.Errorf("unexpected response content: %v", postOkRes)
			}

			// 4. Test POST /users (Failure)
			payloadBad := `{"username": ""}`
			reqPostBad := httptest.NewRequest("POST", "/users", bytes.NewBufferString(payloadBad))
			reqPostBad.Header.Set("Content-Type", "application/json")
			rrPostBad := httptest.NewRecorder()
			handler.ServeHTTP(rrPostBad, reqPostBad)

			if rrPostBad.Code != http.StatusBadRequest {
				t.Errorf("expected POST /users status 400, got %d", rrPostBad.Code)
			}
		})
	}
}
