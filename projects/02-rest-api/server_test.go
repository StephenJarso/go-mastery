package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go-mastery/projects"
)

func TestTaskServer(t *testing.T) {
	server := NewTaskServer()
	defer server.Close()

	// 1. Create a task
	payload := `{"title": "Test Task", "description": "Verify API", "payload": "run unit test"}`
	req := httptest.NewRequest("POST", "/tasks", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	server.createTaskHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status 201 Created, got %d", resp.StatusCode)
	}

	var task projects.Task
	if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if task.Title != "Test Task" || task.ID != "task-1" {
		t.Errorf("Unexpected created task values: %+v", task)
	}

	// 2. List tasks
	reqList := httptest.NewRequest("GET", "/tasks", nil)
	wList := httptest.NewRecorder()
	server.listTasksHandler(wList, reqList)

	respList := wList.Result()
	if respList.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", respList.StatusCode)
	}

	var tasks []projects.Task
	if err := json.NewDecoder(respList.Body).Decode(&tasks); err != nil {
		t.Fatalf("Failed to decode list response: %v", err)
	}

	if len(tasks) != 1 {
		t.Errorf("Expected 1 task in list, got %d", len(tasks))
	}

	// 3. Process the task
	reqProc := httptest.NewRequest("POST", "/tasks/task-1/process", nil)
	// Add mock path value for standard library path matcher since we bypass mux in direct handler call
	reqProc.SetPathValue("id", "task-1")
	wProc := httptest.NewRecorder()
	server.processTaskHandler(wProc, reqProc)

	respProc := wProc.Result()
	if respProc.StatusCode != http.StatusAccepted {
		t.Errorf("Expected status 202 Accepted, got %d", respProc.StatusCode)
	}

	// Wait briefly for background worker callback to execute
	time.Sleep(300 * time.Millisecond)

	// 4. Retrieve task details and verify status is completed/processing
	reqGet := httptest.NewRequest("GET", "/tasks/task-1", nil)
	reqGet.SetPathValue("id", "task-1")
	wGet := httptest.NewRecorder()
	server.getTaskHandler(wGet, reqGet)

	respGet := wGet.Result()
	if respGet.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", respGet.StatusCode)
	}

	var updatedTask projects.Task
	if err := json.NewDecoder(respGet.Body).Decode(&updatedTask); err != nil {
		t.Fatalf("Failed to decode get response: %v", err)
	}

	if updatedTask.Status != projects.StatusCompleted {
		t.Errorf("Expected task status to be COMPLETED after background processing, got %s. Result: %s", updatedTask.Status, updatedTask.Result)
	}
}
