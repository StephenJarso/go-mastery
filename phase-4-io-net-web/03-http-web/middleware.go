package httpweb

import (
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

// Middleware in Go is a design pattern that wraps an http.Handler with another
// http.Handler to execute logic before and/or after the wrapped handler runs.
// It is commonly used for logging, authentication, rate limiting, and panic recovery.

// LoggingMiddleware logs details about every incoming HTTP request.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Serve the request by calling the next handler in the chain.
		next.ServeHTTP(w, r)

		// Record the execution duration.
		duration := time.Since(start)
		log.Printf("[HTTP] %s %s %s", r.Method, r.URL.Path, duration)
	})
}

// RecoveryMiddleware recovers from any panics thrown inside handlers
// to prevent the entire web server from crashing.
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// defer runs when this handler function returns (or during a panic).
		defer func() {
			if err := recover(); err != nil {
				// Log the panic error details and stack trace.
				log.Printf("[PANIC RECOVER] %v\n%s", err, debug.Stack())

				// Return a generic HTTP 500 Internal Server Error to the client.
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "Internal Server Error"}`))
			}
		}()

		// Call the next handler in the chain.
		next.ServeHTTP(w, r)
	})
}
