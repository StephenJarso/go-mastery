package advanced_testing

import (
	"net/http"
	"io"
	"errors"
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
