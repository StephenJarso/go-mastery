package networking

import (
	"net"
	"errors"
	"time"
)


var _ = net.Dial
var _ = errors.New
var _ = time.Now

// Exercise 1: TCP Server
// Start a TCP listener on addr. Accept connections, read string payload, and write back "ECHO: <payload>".
// Accept a stop channel to terminate the listener.
func StartEchoTCPServer(addr string, stopChan chan struct{}) error {
	// TODO: Implement
	return nil
}

// Exercise 2: TCP Client
// Dial TCP server at addr, write message, and read response. Return response string.
func SendTCPMessage(addr, message string) (string, error) {
	// TODO: Implement
	return "", nil
}
