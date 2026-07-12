package serialization

import (
	"fmt"
	"time"

	"github.com/vmihailenco/msgpack/v5"
)

// MessagePack is an efficient binary serialization format.
// It lets you exchange data like JSON but is much faster and produces significantly smaller payloads.
// It is widely used in high-performance cache caches (like Redis) and RPC communication.

type LogEntry struct {
	Level     string    `msgpack:"level"`     // Msgpack tag mapping
	Message   string    `msgpack:"message"`
	Timestamp time.Time `msgpack:"timestamp"`
	Host      string    `msgpack:"host,omitempty"`
}

// MarshalLogEntry converts a LogEntry struct to MessagePack binary format.
func MarshalLogEntry(entry LogEntry) ([]byte, error) {
	data, err := msgpack.Marshal(entry)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal msgpack: %w", err)
	}
	return data, nil
}

// UnmarshalLogEntry parses MessagePack binary bytes back to a LogEntry struct.
func UnmarshalLogEntry(data []byte) (LogEntry, error) {
	var entry LogEntry
	err := msgpack.Unmarshal(data, &entry)
	if err != nil {
		return LogEntry{}, fmt.Errorf("failed to unmarshal msgpack: %w", err)
	}
	return entry, nil
}
