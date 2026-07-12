package networking

import (
	"bufio"
	"net"
	"testing"
	"time"
)

func TestKVStoreServer(t *testing.T) {
	address := "127.0.0.1:18081"
	store := NewKVStore()

	readyChan := make(chan struct{})
	stopChan := make(chan struct{})

	// Start server in background
	go func() {
		err := StartKVStoreServer(address, store, readyChan, stopChan)
		if err != nil {
			t.Errorf("Server error: %v", err)
		}
	}()

	// Wait for server to be ready
	<-readyChan

	// Helper to send a command and get response
	sendCommand := func(cmd string) string {
		conn, err := net.Dial("tcp", address)
		if err != nil {
			t.Fatalf("failed to dial KV server: %v", err)
		}
		defer conn.Close()

		_, err = conn.Write([]byte(cmd + "\n"))
		if err != nil {
			t.Fatalf("failed to write command: %v", err)
		}

		reader := bufio.NewReader(conn)
		res, err := reader.ReadString('\n')
		if err != nil {
			t.Fatalf("failed to read response: %v", err)
		}
		return res
	}

	// 1. Get non-existent key
	res := sendCommand("GET username")
	if res != "ERR: Key not found\n" {
		t.Errorf("expected 'ERR: Key not found\n', got %q", res)
	}

	// 2. Set key
	res = sendCommand("SET username StephenJarso")
	if res != "OK\n" {
		t.Errorf("expected 'OK\n', got %q", res)
	}

	// 3. Get key
	res = sendCommand("GET username")
	if res != "StephenJarso\n" {
		t.Errorf("expected 'StephenJarso\n', got %q", res)
	}

	// 4. Delete key
	res = sendCommand("DELETE username")
	if res != "OK\n" {
		t.Errorf("expected 'OK\n', got %q", res)
	}

	// 5. Get key after delete
	res = sendCommand("GET username")
	if res != "ERR: Key not found\n" {
		t.Errorf("expected 'ERR: Key not found\n', got %q", res)
	}

	// Clean up
	close(stopChan)
	// Give server time to shutdown
	time.Sleep(100 * time.Millisecond)
}

func TestHeartbeatMonitor(t *testing.T) {
	address := "127.0.0.1:18082"
	monitor := NewHeartbeatMonitor(address)

	readyChan := make(chan struct{})
	stopChan := make(chan struct{})

	// Start monitor server
	go func() {
		err := monitor.StartMonitorServer(readyChan, stopChan)
		if err != nil {
			t.Errorf("Monitor server error: %v", err)
		}
	}()

	<-readyChan

	// Dial UDP to send heartbeats
	serverAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		t.Fatalf("failed to resolve UDP address: %v", err)
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		t.Fatalf("failed to dial UDP monitor: %v", err)
	}
	defer conn.Close()

	// Send heartbeats for two clients
	_, err = conn.Write([]byte("HEARTBEAT client-A"))
	if err != nil {
		t.Errorf("failed to send heartbeat: %v", err)
	}

	_, err = conn.Write([]byte("HEARTBEAT client-B"))
	if err != nil {
		t.Errorf("failed to send heartbeat: %v", err)
	}

	// Give the UDP server a brief moment to process the packets
	time.Sleep(100 * time.Millisecond)

	// Verify heartbeats were registered
	tA, okA := monitor.GetLastSeen("client-A")
	if !okA {
		t.Error("expected client-A heartbeat to be registered")
	}
	if time.Since(tA) > 1*time.Second {
		t.Errorf("registered time %v is too old", tA)
	}

	_, okC := monitor.GetLastSeen("client-C")
	if okC {
		t.Error("expected client-C heartbeat to NOT be registered")
	}

	// Clean up
	close(stopChan)
	time.Sleep(100 * time.Millisecond)
}
