package httpweb

import (
	"encoding/json"
	"net/http"
)

// Go's net/http package is extremely powerful and forms the foundation
// of almost all Go web applications. Go 1.22 introduced enhanced routing
// in http.ServeMux, allowing developers to match HTTP methods and path variables.

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// SetupRoutes registers handlers on a ServeMux and returns it.
func SetupRoutes() *http.ServeMux {
	// http.NewServeMux creates a new HTTP request multiplexer.
	mux := http.NewServeMux()

	// 1. Basic handler using http.HandlerFunc
	mux.HandleFunc("GET /", handleHome)

	// 2. JSON handler demonstrating response formatting
	mux.HandleFunc("GET /api/users", handleGetUsers)

	// 3. Path parameter handler (Go 1.22+ wildcard matching)
	// The {id} segment matches any value and is accessible via r.PathValue("id")
	mux.HandleFunc("GET /api/users/{id}", handleGetUserByID)

	// 4. POST request handling
	mux.HandleFunc("POST /api/users", handleCreateUser)

	return mux
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	// Set status header and write a plain text response.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to Go Web Development!"))
}

func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: "1", Name: "StephenJarso"},
		{ID: "2", Name: "Jacob"},
	}

	// Set Content-Type header to let the client know we are returning JSON.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Marshal users slice into JSON and write to http.ResponseWriter
	json.NewEncoder(w).Encode(users)
}

func handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	// Go 1.22+ PathValue retrieves the value of wildcards in the route pattern.
	id := r.PathValue("id")

	if id != "1" && id != "2" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
		return
	}

	user := User{ID: id, Name: "User " + id}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User

	// Decode the JSON request body into the User struct.
	// Always limit request body size in production to avoid DDoS!
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}
	defer r.Body.Close()

	// Respond with the created user and HTTP 201 Created status
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
