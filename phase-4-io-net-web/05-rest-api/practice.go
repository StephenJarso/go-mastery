package restapi

import (
	"net/http"
	"github.com/go-chi/chi/v5"
)


var _ = http.StatusOK
var _ = chi.NewRouter

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Exercise 1: REST API Routing
// Create REST endpoints using chi router:
// GET /items/{id} -> retrieves Item with matching id.
// POST /items -> creates a new Item.
func SetupItemRouter() *chi.Mux {
	// TODO: Implement
	return nil
}
