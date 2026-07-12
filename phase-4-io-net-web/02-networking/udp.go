package networking

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// UDP (User Datagram Protocol) is a connectionless, unreliable transport protocol.
// There is no connection establishment (handshake) or byte stream.
// Instead, UDP operates on discrete "packets" (datagrams).
// Since it is connectionless, there is no listener.Accept() loop.
// The server simply listens for packets from any sender and replies to them by IP.

// StartUDPServer runs a UDP echo server that listens on the given port.
func StartUDPServer(address string, readyChan chan<- struct{}, stopChan <-chan struct{}) error {
	// 1. Resolve the UDP address
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return fmt.Errorf("failed to resolve UDP address: %w", err)
	}

	// 2. Listen on UDP port.
	// net.ListenUDP returns a *net.UDPConn.
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen on UDP: %w", err)
	}
	defer conn.Close()

	// Signal server is ready.
	close(readyChan)

	// Goroutine to monitor stopChan and close the connection early if signaled.
	go func() {
		<-stopChan
		conn.Close()
	}()

	buffer := make([]byte, 2048)

	for {
		// Set a read deadline so ReadFromUDP doesn't block indefinitely
		// if the server is trying to shut down.
		conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))

		// 3. ReadFromUDP blocks until a UDP packet is received.
		// It returns:
		//  - n: number of bytes read
		//  - remoteAddr: the *net.UDPAddr of the sender (so we know where to reply!)
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			select {
			case <-stopChan:
				// Graceful shutdown
				return nil
			default:
				// If it was just a read timeout, continue the loop to check stopChan.
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					continue
				}
				return fmt.Errorf("error reading from UDP socket: %w", err)
			}
		}

		message := string(buffer[:n])
		response := "UDP_ECHO: " + message

		// 4. Send a reply back to the remote address.
		// Since UDP is connectionless, we must specify the target address for every write.
		_, err = conn.WriteToUDP([]byte(response), remoteAddr)
		if err != nil {
			fmt.Printf("failed to write response to %v: %v\n", remoteAddr, err)
		}
	}
}

// RunUDPClient sends a datagram to a UDP server and waits for a response.
func RunUDPClient(serverAddress string, message string) (string, error) {
	// 1. Resolve the server address.
	serverAddr, err := net.ResolveUDPAddr("udp", serverAddress)
	if err != nil {
		return "", fmt.Errorf("failed to resolve server address: %w", err)
	}

	// 2. Dial UDP.
	// Note: Because UDP is connectionless, net.DialUDP doesn't actually connect to the server.
	// It simply sets the default remote address for future Write and Read calls.
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		return "", fmt.Errorf("failed to dial UDP server: %w", err)
	}
	defer conn.Close()

	// 3. Write message.
	_, err = conn.Write([]byte(message))
	if err != nil {
		return "", fmt.Errorf("failed to write UDP packet: %w", err)
	}

	// 4. Read response. Set a timeout so we don't block forever if a packet is lost.
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	buffer := make([]byte, 2048)
	n, err := conn.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("failed to read UDP packet: %w", err)
	}

	return strings.TrimSpace(string(buffer[:n])), nil
}
