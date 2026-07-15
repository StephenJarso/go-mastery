package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"go-mastery/projects"
	"go-mastery/projects/03-concurrent-system"
)

// TaskServer manages tasks and coordinates with the worker pool
type TaskServer struct {
	mu         sync.RWMutex
	tasks      map[string]*projects.Task
	workerPool *concurrent_system.WorkerPool
}

func NewTaskServer() *TaskServer {
	s := &TaskServer{
		tasks: make(map[string]*projects.Task),
	}

	// Initialize the worker pool with 3 workers and a queue size of 100.
	// When a worker finishes processing, the callback updates the task status in memory.
	s.workerPool = concurrent_system.NewWorkerPool(3, 100, func(t *projects.Task) {
		s.mu.Lock()
		defer s.mu.Unlock()
		if existing, ok := s.tasks[t.ID]; ok {
			existing.Status = t.Status
			existing.Result = t.Result
			existing.UpdatedAt = t.UpdatedAt
			log.Printf("[WorkerCallback] Task %s processed with status %s", t.ID, t.Status)
		}
	})
	s.workerPool.Start()

	return s
}

func (s *TaskServer) Close() {
	s.workerPool.Shutdown()
}

func (s *TaskServer) HandleTasks(w http.ResponseWriter, r *http.Header) {
	// To be registered with router
}

func main() {
	server := NewTaskServer()
	defer server.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks", server.listTasksHandler)
	mux.HandleFunc("POST /tasks", server.createTaskHandler)
	mux.HandleFunc("GET /tasks/{id}", server.getTaskHandler)
	mux.HandleFunc("POST /tasks/{id}/process", server.processTaskHandler)

	addr := ":8080"
	log.Printf("Starting Task REST API Server on %s...", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

// createTaskHandler handles POST /tasks
func (s *TaskServer) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Payload     string `json:"payload"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if input.Title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	taskID := fmt.Sprintf("task-%d", len(s.tasks)+1)
	task := &projects.Task{
		ID:          taskID,
		Title:       input.Title,
		Description: input.Description,
		Payload:     input.Payload,
		Status:      projects.StatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.tasks[taskID] = task

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// listTasksHandler handles GET /tasks
func (s *TaskServer) listTasksHandler(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	list := make([]*projects.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		list = append(list, task)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

// getTaskHandler handles GET /tasks/{id}
func (s *TaskServer) getTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	s.mu.RLock()
	task, exists := s.tasks[id]
	s.mu.RUnlock()

	if !exists {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// processTaskHandler handles POST /tasks/{id}/process
func (s *TaskServer) processTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	s.mu.Lock()
	task, exists := s.tasks[id]
	if !exists {
		s.mu.Unlock()
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	if task.Status == projects.StatusProcessing {
		s.mu.Unlock()
		http.Error(w, "task is already processing", http.StatusConflict)
		return
	}

	// Submit task to worker pool
	submitted := s.workerPool.Submit(task)
	s.mu.Unlock()

	if !submitted {
		http.Error(w, "worker pool queue is full", http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "task submitted for concurrent processing",
		"id":      id,
	})
}
