package webframeworks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestChiRouter(t *testing.T) {
	router := SetupChiRouter()

	// Test GET /
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", rr.Code)
	}
	if rr.Body.String() != "Hello from Chi router!" {
		t.Errorf("unexpected body: %q", rr.Body.String())
	}

	// Test GET /api/items/123
	req2 := httptest.NewRequest("GET", "/api/items/123", nil)
	rr2 := httptest.NewRecorder()
	router.ServeHTTP(rr2, req2)

	if rr2.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", rr2.Code)
	}

	var res map[string]string
	json.Unmarshal(rr2.Body.Bytes(), &res)
	if res["item_id"] != "123" || res["router"] != "chi" {
		t.Errorf("unexpected response: %v", res)
	}
}

func TestGinRouter(t *testing.T) {
	router := SetupGinRouter()

	// Test GET /
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", rr.Code)
	}
	if rr.Body.String() != "Hello from Gin!" {
		t.Errorf("unexpected body: %q", rr.Body.String())
	}

	// Test GET /api/items/42?details=true
	req2 := httptest.NewRequest("GET", "/api/items/42?details=true", nil)
	rr2 := httptest.NewRecorder()
	router.ServeHTTP(rr2, req2)

	if rr2.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", rr2.Code)
	}

	var res2 map[string]interface{}
	json.Unmarshal(rr2.Body.Bytes(), &res2)
	if res2["item_id"] != "42" || res2["details"] != "true" || res2["router"] != "gin" {
		t.Errorf("unexpected response: %v", res2)
	}

	// Test POST /api/items
	payload := `{"name": "Gopher Plush", "price": 24.99}`
	req3 := httptest.NewRequest("POST", "/api/items", bytes.NewBufferString(payload))
	req3.Header.Set("Content-Type", "application/json")
	rr3 := httptest.NewRecorder()
	router.ServeHTTP(rr3, req3)

	if rr3.Code != http.StatusCreated {
		t.Errorf("expected 201 Created, got %d", rr3.Code)
	}
}

func TestEchoRouter(t *testing.T) {
	router := SetupEchoRouter()

	// Test GET /
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", rr.Code)
	}
	if rr.Body.String() != "Hello from Echo!" {
		t.Errorf("unexpected body: %q", rr.Body.String())
	}

	// Test GET /api/items/99?details=true
	req2 := httptest.NewRequest("GET", "/api/items/99?details=true", nil)
	rr2 := httptest.NewRecorder()
	router.ServeHTTP(rr2, req2)

	if rr2.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", rr2.Code)
	}

	var res2 map[string]string
	json.Unmarshal(rr2.Body.Bytes(), &res2)
	if res2["item_id"] != "99" || res2["details"] != "true" || res2["router"] != "echo" {
		t.Errorf("unexpected response: %v", res2)
	}
}
