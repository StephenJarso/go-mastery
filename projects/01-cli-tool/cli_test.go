package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go-mastery/projects"
)

func TestCLIClient_Methods(t *testing.T) {
	// Create a mock server that simulates our REST API
	var mockedTasks = []*projects.Task{
		{
			ID:          "task-1",
			Title:       "Existing Task",
			Description: "Simulated",
			Status:      projects.StatusPending,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockedTasks)
	})

	mux.HandleFunc("POST /tasks", func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Payload     string `json:"payload"`
		}
		json.NewDecoder(r.Body).Decode(&input)

		newTask := &projects.Task{
			ID:          "task-2",
			Title:       input.Title,
			Description: input.Description,
			Payload:     input.Payload,
			Status:      projects.StatusPending,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		mockedTasks = append(mockedTasks, newTask)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTask)
	})

	mux.HandleFunc("GET /tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		for _, t := range mockedTasks {
			if t.ID == id {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(t)
				return
			}
		}
		http.Error(w, "not found", http.StatusNotFound)
	})

	mux.HandleFunc("POST /tasks/{id}/process", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		for _, t := range mockedTasks {
			if t.ID == id {
				t.Status = projects.StatusProcessing
				w.WriteHeader(http.StatusAccepted)
				w.Write([]byte(`{"message": "processing"}`))
				return
			}
		}
		http.Error(w, "not found", http.StatusNotFound)
	})

	mockServer := httptest.NewServer(mux)
	defer mockServer.Close()

	cli := NewCLIClient(mockServer.URL)

	// Test list
	// Capture output redirecting to stdout isn't strictly necessary; we test that the API calls go through successfully without panicking.
	cli.listTasks()

	// Test add
	cli.addTask("New Task", "Description", "Payload")
	if len(mockedTasks) != 2 {
		t.Errorf("Expected 2 tasks after addTask, got %d", len(mockedTasks))
	}
	if mockedTasks[1].Title != "New Task" {
		t.Errorf("Expected added task title to be 'New Task', got %s", mockedTasks[1].Title)
	}

	// Test get
	cli.getTask("task-1")

	// Test process
	cli.processTask("task-1")
	if mockedTasks[0].Status != projects.StatusProcessing {
		t.Errorf("Expected task-1 status to be processing, got %s", mockedTasks[0].Status)
	}
}
