package httpweb

import (
	"testing"
	"net/http"
	"net/http/httptest"
)


func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	HelloHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}
	if rr.Body.String() != "Hello, World!" {
		t.Errorf("expected Hello, World!, got %s", rr.Body.String())
	}
}

func TestAuthMiddleware(t *testing.T) {
	handler := AuthMiddleware(http.HandlerFunc(HelloHandler))
	
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", rr.Code)
	}

	req.Header.Set("Authorization", "secret-token")
	rr2 := httptest.NewRecorder()
	handler.ServeHTTP(rr2, req)
	if rr2.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr2.Code)
	}
}
