package networking

import (
	"testing"
	"time"
)


func TestTCPEcho(t *testing.T) {
	addr := "127.0.0.1:9099"
	stop := make(chan struct{})
	defer close(stop)

	go func() {
		StartEchoTCPServer(addr, stop)
	}()

	// Wait for server to start
	time.Sleep(100 * time.Millisecond)

	res, err := SendTCPMessage(addr, "Hello Go")
	if err != nil {
		t.Fatalf("SendTCPMessage failed: %v", err)
	}

	expected := "ECHO: Hello Go"
	if res != expected {
		t.Errorf("expected %q, got %q", expected, res)
	}
}
