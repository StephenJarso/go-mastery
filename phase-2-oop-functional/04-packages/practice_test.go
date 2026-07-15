package phase2packages

import (
	"testing"
	"time"
)

func TestFormatUser(t *testing.T) {
	u := User{Name: "john doe", Age: 30}
	expected := "User: JOHN DOE (Age: 30)"
	if res := FormatUser(u); res != expected {
		t.Errorf("FormatUser(%v) = %q, want %q", u, res, expected)
	}
}

func TestParseAndSum(t *testing.T) {
	sum, err := ParseAndSum("15", "25")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if sum != 40 {
		t.Errorf("expected 40, got %d", sum)
	}

	_, err = ParseAndSum("abc", "25")
	if err == nil {
		t.Error("expected error parsing 'abc'")
	}
}

func TestDaysUntilBirthday(t *testing.T) {
	ref := time.Date(2026, 7, 16, 0, 0, 0, 0, time.UTC)
	
	// Birthday in 4 days
	days, err := DaysUntilBirthday(ref, "1999-07-20")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if days != 4 {
		t.Errorf("expected 4 days, got %d", days)
	}

	// Birthday today
	daysToday, err := DaysUntilBirthday(ref, "1999-07-16")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if daysToday != 0 {
		t.Errorf("expected 0 days, got %d", daysToday)
	}

	// Birthday already passed this year (June 10)
	daysPassed, err := DaysUntilBirthday(ref, "1999-06-10")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// June 10, 2027 is 329 days from July 16, 2026
	expectedNextYear := int(time.Date(2027, 6, 10, 0, 0, 0, 0, time.UTC).Sub(ref).Hours() / 24)
	if daysPassed != expectedNextYear {
		t.Errorf("expected %d days, got %d", expectedNextYear, daysPassed)
	}
}
