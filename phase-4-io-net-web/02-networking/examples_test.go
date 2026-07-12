package networking

import (
	"testing"
	"time"
)

func TestBasicTCPConnection(t *testing.T) {
	address := "127.0.0.1:18083"
	readyChan := make(chan struct{})
	stopChan := make(chan struct{})

	// Start basic TCP server in the background
	go func() {
		err := StartTCPServer(address, readyChan, stopChan)
		if err != nil {
			t.Errorf("TCP server error: %v", err)
		}
	}()

	// Wait for server to listen
	<-readyChan

	// Run client
	msg := "Hello TCP Server!"
	response, err := RunTCPClient(address, msg)
	if err != nil {
		t.Fatalf("RunTCPClient failed: %v", err)
	}

	expectedResponse := "ECHO: " + msg
	if response != expectedResponse {
		t.Errorf("expected response %q, got %q", expectedResponse, response)
	}

	// Try exiting
	exitResponse, err := RunTCPClient(address, "EXIT")
	if err != nil {
		t.Fatalf("RunTCPClient EXIT failed: %v", err)
	}
	if exitResponse != "Goodbye!" {
		t.Errorf("expected 'Goodbye!', got %q", exitResponse)
	}

	// Clean up server
	close(stopChan)
	time.Sleep(100 * time.Millisecond)
}

func TestBasicUDPConnection(t *testing.T) {
	address := "127.0.0.1:18084"
	readyChan := make(chan struct{})
	stopChan := make(chan struct{})

	// Start basic UDP server in the background
	go func() {
		err := StartUDPServer(address, readyChan, stopChan)
		if err != nil {
			t.Errorf("UDP server error: %v", err)
		}
	}()

	// Wait for server to listen
	<-readyChan

	// Run client
	msg := "Hello UDP Server!"
	response, err := RunUDPClient(address, msg)
	if err != nil {
		t.Fatalf("RunUDPClient failed: %v", err)
	}

	expectedResponse := "UDP_ECHO: " + msg
	if response != expectedResponse {
		t.Errorf("expected response %q, got %q", expectedResponse, response)
	}

	// Clean up server
	close(stopChan)
	time.Sleep(100 * time.Millisecond)
}
