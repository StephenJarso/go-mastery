package jsonprocessing

import (
	"encoding/json"
	"testing"
)

func TestParseConfig(t *testing.T) {
	// Case 1: Config with empty values -> should apply defaults
	cfg1, err := ParseConfig(`{}`)
	if err != nil {
		t.Fatalf("ParseConfig failed: %v", err)
	}
	if cfg1.Port != 8080 {
		t.Errorf("expected default port 8080, got %d", cfg1.Port)
	}
	if cfg1.Host != "localhost" {
		t.Errorf("expected default host 'localhost', got %q", cfg1.Host)
	}

	// Case 2: Config with partial values
	cfg2, err := ParseConfig(`{"port": 9090, "features": {"metrics": true}}`)
	if err != nil {
		t.Fatalf("ParseConfig failed: %v", err)
	}
	if cfg2.Port != 9090 {
		t.Errorf("expected port 9090, got %d", cfg2.Port)
	}
	if cfg2.Host != "localhost" {
		t.Errorf("expected default host 'localhost', got %q", cfg2.Host)
	}
	if !cfg2.Features["metrics"] {
		t.Error("expected features['metrics'] to be true")
	}
}

func TestTransactionCustomUnmarshal(t *testing.T) {
	jsonData := `{"id": "tx-999", "amount": "$1250.75"}`
	
	var tx Transaction
	err := json.Unmarshal([]byte(jsonData), &tx)
	if err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if tx.ID != "tx-999" {
		t.Errorf("expected transaction ID 'tx-999', got %q", tx.ID)
	}
	if tx.Amount != 1250.75 {
		t.Errorf("expected transaction amount 1250.75, got %f", tx.Amount)
	}

	// Test error cases
	badData := `{"id": "tx-abc", "amount": "invalid-money"}`
	var txBad Transaction
	err = json.Unmarshal([]byte(badData), &txBad)
	if err == nil {
		t.Error("expected unmarshal to fail for invalid currency formats")
	}
}
