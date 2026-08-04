package phase2packages

import (
	"testing"
	"time"
)

func TestFormatURLPath(t *testing.T) {
	res := FormatURLPath("Hello, World! Go 101")
	if res == "" {
		t.Skip("Exercise FormatURLPath not implemented yet")
	}
	expected := "hello-world-go-101"
	if res != expected {
		t.Errorf("expected %q, got %q", expected, res)
	}
}

func TestParseAndSumCSV(t *testing.T) {
	sum, err := ParseAndSumCSV(" 5, 10, 15 ")
	if sum == 0 && err == nil {
		t.Skip("Exercise ParseAndSumCSV not implemented yet")
	}
	if err != nil || sum != 30 {
		t.Errorf("expected 30, got %d, err: %v", sum, err)
	}
}

func TestDaysSinceBirth(t *testing.T) {
	birth := "2000-01-01"
	now := time.Date(2000, 1, 11, 12, 0, 0, 0, time.UTC)
	days, err := DaysSinceBirth(birth, now)
	if days == 0 && err == nil {
		t.Skip("Exercise DaysSinceBirth not implemented yet")
	}
	if err != nil || days != 10 {
		t.Errorf("expected 10 days, got %d, err: %v", days, err)
	}
}
