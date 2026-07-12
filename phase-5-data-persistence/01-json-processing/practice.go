package jsonprocessing

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// PRACTICE EXERCISE #1: Configuration Parser
// Implement a function that parses a server configuration JSON.
// The JSON has some optional fields.
// Struct rules:
// - "port" (int): if not specified (or 0), default to 8080.
// - "host" (string): if not specified (or ""), default to "localhost".
// - "features" (map[string]bool): holds dynamic flag states.
// Return the parsed Config struct.

type Config struct {
	Port     int             `json:"port,omitempty"`
	Host     string          `json:"host,omitempty"`
	Features map[string]bool `json:"features"`
}

func ParseConfig(jsonData string) (Config, error) {
	var cfg Config
	err := json.Unmarshal([]byte(jsonData), &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Apply default values
	if cfg.Port == 0 {
		cfg.Port = 8080
	}
	if cfg.Host == "" {
		cfg.Host = "localhost"
	}
	if cfg.Features == nil {
		cfg.Features = make(map[string]bool)
	}

	return cfg, nil
}

// PRACTICE EXERCISE #2: Currency Transaction Deserializer
// In financial systems, transaction values are sometimes formatted as strings with currency symbols.
// E.g., JSON payload: {"id": "tx-123", "amount": "$150.50"}
// Parse this JSON into the Transaction struct, converting the string amount to a float64.
// Implement custom unmarshaling for the Transaction struct to handle the currency parsing.

type Transaction struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
}

// UnmarshalJSON implements custom unmarshaling for Transaction.
// It parses the ID normally, but strips any "$" prefix from the amount string and converts it to float64.
func (t *Transaction) UnmarshalJSON(data []byte) error {
	// To avoid infinite recursion, we unmarshal into a temporary alias struct or a map.
	type Alias Transaction
	
	// Create an intermediate structure to capture the raw amount string.
	var raw struct {
		ID     string `json:"id"`
		Amount string `json:"amount"`
	}

	err := json.Unmarshal(data, &raw)
	if err != nil {
		return fmt.Errorf("failed to parse raw transaction: %w", err)
	}

	t.ID = raw.ID

	// Strip currency symbols (e.g. "$", "€", "£")
	cleanAmountStr := strings.TrimLeft(raw.Amount, "$€£ ")
	
	amountVal, err := strconv.ParseFloat(cleanAmountStr, 64)
	if err != nil {
		return fmt.Errorf("invalid currency format %q: %w", raw.Amount, err)
	}

	t.Amount = amountVal
	return nil
}
