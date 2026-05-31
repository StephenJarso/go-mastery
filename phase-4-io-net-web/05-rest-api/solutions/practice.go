package solutions

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"encoding/json"
	"strconv"
)


type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var items = make(map[int]Item)

func SetupItemRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/items/{id}", func(w http.ResponseWriter, req *http.Request) {
		idStr := chi.URLParam(req, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		item, ok := items[id]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(item)
	})

	r.Post("/items", func(w http.ResponseWriter, req *http.Request) {
		var item Item
		err := json.NewDecoder(req.Body).Decode(&item)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		items[item.ID] = item
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(item)
	})
	return r
}
