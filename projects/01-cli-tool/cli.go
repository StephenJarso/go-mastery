package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"go-mastery/projects"
)

type CLIClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewCLIClient(baseURL string) *CLIClient {
	return &CLIClient{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

func main() {
	baseURL := os.Getenv("TASK_API_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	client := NewCLIClient(baseURL)

	if len(os.Args) < 2 {
		client.printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		title := addCmd.String("title", "", "Task title")
		desc := addCmd.String("desc", "", "Task description")
		payload := addCmd.String("payload", "", "Task payload")
		addCmd.Parse(os.Args[2:])

		if *title == "" {
			fmt.Println("Error: -title flag is required")
			addCmd.Usage()
			os.Exit(1)
		}

		client.addTask(*title, *desc, *payload)

	case "list":
		client.listTasks()

	case "get":
		getCmd := flag.NewFlagSet("get", flag.ExitOnError)
		id := getCmd.String("id", "", "Task ID")
		getCmd.Parse(os.Args[2:])

		if *id == "" {
			fmt.Println("Error: -id flag is required")
			getCmd.Usage()
			os.Exit(1)
		}

		client.getTask(*id)

	case "process":
		processCmd := flag.NewFlagSet("process", flag.ExitOnError)
		id := processCmd.String("id", "", "Task ID")
		processCmd.Parse(os.Args[2:])

		if *id == "" {
			fmt.Println("Error: -id flag is required")
			processCmd.Usage()
			os.Exit(1)
		}

		client.processTask(*id)

	default:
		fmt.Printf("Unknown command %q\n", command)
		client.printUsage()
		os.Exit(1)
	}
}

func (c *CLIClient) printUsage() {
	fmt.Println("Usage: task-cli <command> [flags]")
	fmt.Println("\nCommands:")
	fmt.Println("  add -title <title> [-desc <desc>] [-payload <payload>]")
	fmt.Println("  list")
	fmt.Println("  get -id <task-id>")
	fmt.Println("  process -id <task-id>")
}

func (c *CLIClient) addTask(title, desc, payload string) {
	data := map[string]string{
		"title":       title,
		"description": desc,
		"payload":     payload,
	}
	body, _ := json.Marshal(data)

	resp, err := c.HTTPClient.Post(c.BaseURL+"/tasks", "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("HTTP request failed: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		respBody, _ := io.ReadAll(resp.Body)
		fmt.Printf("Failed to add task: %s (Status: %d)\n", string(respBody), resp.StatusCode)
		return
	}

	var task projects.Task
	json.NewDecoder(resp.Body).Decode(&task)
	fmt.Printf("Task created successfully: ID=%s, Title=%s, Status=%s\n", task.ID, task.Title, task.Status)
}

func (c *CLIClient) listTasks() {
	resp, err := c.HTTPClient.Get(c.BaseURL + "/tasks")
	if err != nil {
		fmt.Printf("HTTP request failed: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to list tasks (Status: %d)\n", resp.StatusCode)
		return
	}

	var tasks []projects.Task
	json.NewDecoder(resp.Body).Decode(&tasks)

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Printf("%-10s %-20s %-12s %-20s\n", "ID", "TITLE", "STATUS", "RESULT")
	fmt.Println(strings.Repeat("-", 68))
	for _, t := range tasks {
		fmt.Printf("%-10s %-20s %-12s %-20s\n", t.ID, t.Title, t.Status, t.Result)
	}
}

func (c *CLIClient) getTask(id string) {
	resp, err := c.HTTPClient.Get(c.BaseURL + "/tasks/" + id)
	if err != nil {
		fmt.Printf("HTTP request failed: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("Task with ID %q not found.\n", id)
		return
	} else if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to get task (Status: %d)\n", resp.StatusCode)
		return
	}

	var t projects.Task
	json.NewDecoder(resp.Body).Decode(&t)

	fmt.Printf("Task Details:\n")
	fmt.Printf("  ID:          %s\n", t.ID)
	fmt.Printf("  Title:       %s\n", t.Title)
	fmt.Printf("  Description: %s\n", t.Description)
	fmt.Printf("  Payload:     %s\n", t.Payload)
	fmt.Printf("  Status:      %s\n", t.Status)
	fmt.Printf("  Created At:  %s\n", t.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("  Updated At:  %s\n", t.UpdatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("  Result:      %s\n", t.Result)
}

func (c *CLIClient) processTask(id string) {
	resp, err := c.HTTPClient.Post(c.BaseURL+"/tasks/"+id+"/process", "application/json", nil)
	if err != nil {
		fmt.Printf("HTTP request failed: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("Task with ID %q not found.\n", id)
		return
	} else if resp.StatusCode == http.StatusConflict {
		fmt.Printf("Task with ID %q is already being processed.\n", id)
		return
	} else if resp.StatusCode != http.StatusAccepted {
		respBody, _ := io.ReadAll(resp.Body)
		fmt.Printf("Failed to submit task: %s (Status: %d)\n", string(respBody), resp.StatusCode)
		return
	}

	fmt.Printf("Task %q submitted successfully for background processing.\n", id)
}
