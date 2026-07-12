package webframeworks

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Chi is a lightweight, idiomatic routing library.
// It is designed to be 100% compatible with net/http handlers.
// You don't need a special Context type; it uses standard http.ResponseWriter and *http.Request.

func SetupChiRouter() http.Handler {
	r := chi.NewRouter()

	// 1. Chi comes with excellent built-in middlewares.
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// 2. Basic GET request
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Chi router!"))
	})

	// 3. Route grouping
	r.Route("/api", func(r chi.Router) {
		// Define route with path parameters
		r.Get("/items/{id}", handleChiGetItem)
	})

	return r
}

func handleChiGetItem(w http.ResponseWriter, r *http.Request) {
	// 4. Retrieve path parameters using chi.URLParam
	itemID := chi.URLParam(r, "id")

	response := map[string]string{
		"router":  "chi",
		"item_id": itemID,
		"status":  "success",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
