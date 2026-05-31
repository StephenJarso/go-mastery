package restapi

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
)


func TestItemRouter(t *testing.T) {
	r := SetupItemRouter()

	// Post Item
	payload := `{"id":1,"name":"pencil"}`
	req := httptest.NewRequest("POST", "/items", bytes.NewBufferString(payload))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusCreated {
		t.Errorf("expected 201 Created, got %d", rr.Code)
	}

	// Get Item
	req = httptest.NewRequest("GET", "/items/1", nil)
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", rr.Code)
	}
}
