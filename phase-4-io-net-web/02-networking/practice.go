package networking

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

// PRACTICE EXERCISE #1: TCP-Based Key-Value Store
// Build a simple key-value store server that communicates over TCP.
// The store itself should be thread-safe (use sync.RWMutex).
// The protocol is simple text commands separated by newlines:
//   - "SET <key> <value>" -> sets the key, returns "OK\n"
//   - "GET <key>"         -> returns "<value>\n" or "ERR: Key not found\n"
//   - "DELETE <key>"      -> deletes the key, returns "OK\n"
//   - Any other input    -> returns "ERR: Invalid command\n"

type KVStore struct {
	mu    sync.RWMutex
	store map[string]string
}

func NewKVStore() *KVStore {
	return &KVStore{
		store: make(map[string]string),
	}
}

func (k *KVStore) Set(key, val string) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.store[key] = val
}

func (k *KVStore) Get(key string) (string, bool) {
	k.mu.RLock()
	defer k.mu.RUnlock()
	val, ok := k.store[key]
	return val, ok
}

func (k *KVStore) Delete(key string) {
	k.mu.Lock()
	defer k.mu.Unlock()
	delete(k.store, key)
}

// StartKVStoreServer starts listening on the address and handles commands.
func StartKVStoreServer(address string, store *KVStore, readyChan chan<- struct{}, stopChan <-chan struct{}) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to start KV store listener: %w", err)
	}
	defer listener.Close()

	close(readyChan)

	// Keep track of client connections for cleanup.
	activeConns := make(chan net.Conn, 100)

	go func() {
		<-stopChan
		listener.Close()
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
		conn, err := listener.Accept()
		if err != nil {
			select {
			case <-stopChan:
				return nil
			default:
				return fmt.Errorf("error accepting client: %w", err)
			}
		}

		activeConns <- conn

		// Handle client KV commands in a separate goroutine
		go func(c net.Conn) {
			defer c.Close()
			scanner := bufio.NewScanner(c)
			for scanner.Scan() {
				input := scanner.Text()
				parts := strings.Fields(input)
				if len(parts) == 0 {
					continue
				}

				cmd := strings.ToUpper(parts[0])
				var response string

				switch cmd {
				case "SET":
					if len(parts) < 3 {
						response = "ERR: Usage: SET <key> <value>\n"
					} else {
						// Value might contain spaces, join remaining fields.
						val := strings.Join(parts[2:], " ")
						store.Set(parts[1], val)
						response = "OK\n"
					}
				case "GET":
					if len(parts) != 2 {
						response = "ERR: Usage: GET <key>\n"
					} else {
						val, ok := store.Get(parts[1])
						if ok {
							response = val + "\n"
						} else {
							response = "ERR: Key not found\n"
						}
					}
				case "DELETE":
					if len(parts) != 2 {
						response = "ERR: Usage: DELETE <key>\n"
					} else {
						store.Delete(parts[1])
						response = "OK\n"
					}
				default:
					response = "ERR: Invalid command\n"
				}

				_, err := c.Write([]byte(response))
				if err != nil {
					break
				}
			}
		}(conn)
	}
}

// PRACTICE EXERCISE #2: UDP Heartbeat Monitor
// Build a UDP server that monitors incoming client heartbeat packets.
// Clients send UDP packets containing text like: "HEARTBEAT <client_id>".
// The server stores the last-seen timestamp for each client ID.
// This is typical in gaming or cluster-node health checks where TCP overhead is undesired.

type HeartbeatMonitor struct {
	mu         sync.RWMutex
	lastSeen   map[string]time.Time
	listenPort string
}

func NewHeartbeatMonitor(port string) *HeartbeatMonitor {
	return &HeartbeatMonitor{
		lastSeen:   make(map[string]time.Time),
		listenPort: port,
	}
}

func (hm *HeartbeatMonitor) RegisterHeartbeat(clientID string) {
	hm.mu.Lock()
	defer hm.mu.Unlock()
	hm.lastSeen[clientID] = time.Now()
}

func (hm *HeartbeatMonitor) GetLastSeen(clientID string) (time.Time, bool) {
	hm.mu.RLock()
	defer hm.mu.RUnlock()
	t, ok := hm.lastSeen[clientID]
	return t, ok
}

// StartMonitorServer starts listening on UDP and registers incoming heartbeats.
func (hm *HeartbeatMonitor) StartMonitorServer(readyChan chan<- struct{}, stopChan <-chan struct{}) error {
	addr, err := net.ResolveUDPAddr("udp", hm.listenPort)
	if err != nil {
		return fmt.Errorf("failed to resolve UDP address: %w", err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return fmt.Errorf("failed to start UDP heartbeat server: %w", err)
	}
	defer conn.Close()

	close(readyChan)

	go func() {
		<-stopChan
		conn.Close()
	}()

	buffer := make([]byte, 1024)

	for {
		conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			select {
			case <-stopChan:
				return nil
			default:
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					continue
				}
				return fmt.Errorf("error reading heartbeat UDP packet: %w", err)
			}
		}

		payload := string(buffer[:n])
		parts := strings.Fields(payload)

		// Expecting: "HEARTBEAT <client_id>"
		if len(parts) == 2 && strings.ToUpper(parts[0]) == "HEARTBEAT" {
			clientID := parts[1]
			hm.RegisterHeartbeat(clientID)
		}
	}
}
