package httpweb

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

// ContextKey is a custom type for context keys to avoid collisions.
type ContextKey string

const RequestIDKey ContextKey = "request_id"

// PRACTICE EXERCISE #1: API Token Authentication Middleware
// Implement a middleware that validates an API Token passed in the Authorization header.
// - It should check for the "Authorization" header in the format: "Bearer <token>".
// - If the token matches targetToken, allow the request to proceed.
// - If the header is missing, incorrectly formatted, or the token is incorrect,
//   return StatusUnauthorized (401) with a JSON error payload:
//   {"error": "Unauthorized: invalid token"}
func TokenAuthMiddleware(targetToken string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			// Check format "Bearer <token>"
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized: invalid token"})
				return
			}

			// Extract token
			token := strings.TrimPrefix(authHeader, "Bearer ")
			if token != targetToken {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized: invalid token"})
				return
			}

			// Valid token, proceed.
			next.ServeHTTP(w, r)
		})
	}
}

// PRACTICE EXERCISE #2: Request ID Injector Middleware
// Implement a middleware that injects a unique Request ID into:
// 1. The response headers as "X-Request-ID".
// 2. The request Context under the key RequestIDKey (so subsequent handlers can retrieve it).
// Accept a generator function that returns a string Request ID.
func RequestIDMiddleware(idGenerator func() string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Generate the request ID
			reqID := idGenerator()

			// 1. Inject into response headers
			w.Header().Set("X-Request-ID", reqID)

			// 2. Inject into request context
			ctx := context.WithValue(r.Context(), RequestIDKey, reqID)

			// 3. Pass request with updated context to the next handler
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
