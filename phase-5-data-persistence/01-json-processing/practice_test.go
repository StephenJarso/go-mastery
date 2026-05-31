package jsonprocessing

import (
	"testing"
	"time"
	"encoding/json"
)


func TestParseConfig(t *testing.T) {
	payload := `{"app_name":"GoApp","port":8080}`
	c, err := ParseConfig([]byte(payload))
	if err != nil || c.AppName != "GoApp" || c.Port != 8080 {
		t.Errorf("failed to parse config: %+v, %v", c, err)
	}

	_, err = ParseConfig([]byte(`{"port":8080}`))
	if err == nil {
		t.Error("expected error for empty app_name")
	}
}

func TestCustomDateUnmarshal(t *testing.T) {
	payload := `{"created":"25-12-2026"}`
	var c PracticeCustomDateConfig
	err := json.Unmarshal([]byte(payload), &c)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}
	expected := time.Date(2026, 12, 25, 0, 0, 0, 0, time.UTC)
	if !c.Created.Time.Equal(expected) {
		t.Errorf("expected %v, got %v", expected, c.Created.Time)
	}
}
