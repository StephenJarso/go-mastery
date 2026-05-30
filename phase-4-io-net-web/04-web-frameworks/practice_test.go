package webframeworks

import (
	"testing"
	"net/http"
	"net/http/httptest"
)


func TestGinPing(t *testing.T) {
	r := SetupGinPing()
	req := httptest.NewRequest("GET", "/ping", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}
	if rr.Body.String() != `{"message":"pong"}` {
		t.Errorf("expected pong json, got %s", rr.Body.String())
	}
}

func TestEchoPing(t *testing.T) {
	e := SetupEchoPing()
	req := httptest.NewRequest("GET", "/ping", nil)
	rr := httptest.NewRecorder()
	e.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}
	if rr.Body.String() != `{"message":"pong"}
` { // Echo appends a newline to JSON output
		t.Errorf("expected pong json, got %s", rr.Body.String())
	}
}
