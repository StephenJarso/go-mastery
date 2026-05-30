package solutions

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "secret-token" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func QueryParamsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Create JSON manually to avoid other package dependencies
	var parts []string
	for k, v := range query {
		parts = append(parts, fmt.Sprintf("%q:%q", k, v[0]))
	}
	if len(parts) == 0 {
		w.Write([]byte("{}"))
		return
	}
	w.Write([]byte("{" + fmt.Sprintf("%s", parts[0]) + "}")) // Simple JSON format logic
}
