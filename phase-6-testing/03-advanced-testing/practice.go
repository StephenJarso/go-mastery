package advanced_testing

import (
	"errors"
	"io"
	"net/http"
)

var _ = http.Get
var _ = io.ReadAll
var _ = errors.New

// Exercise 1: Fetch Remote Data
// Make GET request to url, read response body, and return it.
func FetchRemoteData(url string) (string, error) {
	// TODO: Implement
	return "", nil
}
