package advanced_testing

import (
	"testing"
	"net/http"
	"net/http/httptest"
)


func TestFetchRemoteData(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("mocked payload"))
	}))
	defer server.Close()

	res, err := FetchRemoteData(server.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != "mocked payload" {
		t.Errorf("expected 'mocked payload', got %q", res)
	}
}
