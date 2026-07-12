package networking

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

// TCP (Transmission Control Protocol) is a connection-oriented, reliable
// transport protocol. Go's 'net' package provides TCP socket programming
// APIs that make handling connections simple and concurrent.

// StartTCPServer sets up a TCP server that listens on the given address.
// It accepts connections and processes them concurrently in goroutines.
// The closeChan can be used to stop the server from the outside.
func StartTCPServer(address string, readyChan chan<- struct{}, stopChan <-chan struct{}) error {
	// 1. Listen on a TCP port.
	// net.Listen returns a Listener interface that listens for incoming connections.
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", address, err)
	}
	// Always ensure the listener is closed when the function finishes.
	defer listener.Close()

	// Signal that the server is listening and ready for connections.
	close(readyChan)

	// A channel to track active connection goroutines, so we can clean up if needed.
	activeConns := make(chan net.Conn, 100)

	// Goroutine to monitor stopChan and close the listener early if signaled.
	go func() {
		<-stopChan
		listener.Close()
		// Close all active client connections as well.
		for {
			select {
			case conn := <-activeConns:
				if conn != nil {
					conn.Close()
				}
			default:
				return
			}
		}
	}()

	for {
		// 2. Accept blocks until an incoming connection is made.
		// It returns a net.Conn (Connection) which implements io.Reader and io.Writer.
		conn, err := listener.Accept()
		if err != nil {
			// If the listener was closed from stopChan, Accept returns an error.
			// We handle this gracefully.
			select {
			case <-stopChan:
				return nil
			default:
				return fmt.Errorf("error accepting connection: %w", err)
			}
		}

		// Keep track of the connection.
		activeConns <- conn

		// 3. Handle connection concurrently in a separate goroutine.
		// This is vital so the main loop can immediately accept other incoming connections.
		go handleTCPClient(conn)
	}
}

// handleTCPClient reads data from a client connection and echoes it back.
func handleTCPClient(conn net.Conn) {
	// Always close the client connection to release resources (port/sockets).
	defer conn.Close()

	// Set a deadline for I/O operations to prevent stale connections from hanging.
	// This connection will timeout after 5 minutes of inactivity.
	conn.SetDeadline(time.Now().Add(5 * time.Minute))

	// We use bufio.NewScanner to read lines from the connection easily.
	// Since conn implements io.Reader, we can pass it directly.
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		text := scanner.Text()

		// Command processing
		if strings.TrimSpace(text) == "EXIT" {
			conn.Write([]byte("Goodbye!\n"))
			break
		}

		// Process input and send response back to client (conn implements io.Writer).
		response := fmt.Sprintf("ECHO: %s\n", text)
		_, err := conn.Write([]byte(response))
		if err != nil {
			fmt.Printf("Error writing to client: %v\n", err)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner read error: %v\n", err)
	}
}

// RunTCPClient connects to a TCP server, sends a message, and reads the response.
func RunTCPClient(address string, message string) (string, error) {
	// 1. Dial connects to the TCP server at the specified address.
	// We set a timeout for the connection attempt.
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		return "", fmt.Errorf("failed to dial server: %w", err)
	}
	defer conn.Close()

	// 2. Write the message to the server, appending a newline to signal end of command.
	_, err = conn.Write([]byte(message + "\n"))
	if err != nil {
		return "", fmt.Errorf("failed to send message: %w", err)
	}

	// 3. Read the response. We wrap the conn in a bufio.Reader.
	reader := bufio.NewReader(conn)
	// Read until a newline is reached.
	response, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	return strings.TrimSpace(response), nil
}
