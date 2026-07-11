package errors_advanced

import (
	"errors"
	"testing"
)

func TestValidationError(t *testing.T) {
	ve := &ValidationError{
		Errors: map[string]string{
			"email": "invalid email format",
			"age":   "must be at least 18",
		},
	}

	errStr := ve.Error()
	if !errors.As(ve, &ve) {
		t.Error("ValidationError should satisfy errors.As pointer check")
	}

	if !testing.Short() {
		t.Logf("ValidationError output: %s", errStr)
	}
}

func TestRetryOperation(t *testing.T) {
	// 1. Success on first try
	tries := 0
	fSuccess := func() error {
		tries++
		return nil
	}
	err := RetryOperation(fSuccess, 3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if tries != 1 {
		t.Errorf("Expected 1 try, got %d", tries)
	}

	// 2. Failure then success
	tries = 0
	fRecover := func() error {
		tries++
		if tries < 3 {
			return errors.New("temporary error")
		}
		return nil
	}
	err = RetryOperation(fRecover, 3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if tries != 3 {
		t.Errorf("Expected 3 tries, got %d", tries)
	}

	// 3. Complete failure wrapping
	tries = 0
	finalErr := errors.New("permanent failure")
	fFail := func() error {
		tries++
		return finalErr
	}
	err = RetryOperation(fFail, 2)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if tries != 3 { // first try + 2 retries
		t.Errorf("Expected 3 tries, got %d", tries)
	}
	if !errors.Is(err, finalErr) {
		t.Errorf("Expected wrapped error to contain finalErr")
	}
}

func TestSafeRun(t *testing.T) {
	// 1. Safe run without panic
	err := SafeRun(func() {
		// Do nothing
	})
	if err != nil {
		t.Errorf("Unexpected error from non-panicking function: %v", err)
	}

	// 2. Safe run with panic
	err = SafeRun(func() {
		panic("boom")
	})
	if err == nil {
		t.Fatal("Expected error from panicking function, got nil")
	}
	if err.Error() != "panic occurred: boom" {
		t.Errorf("Expected 'panic occurred: boom', got %q", err.Error())
	}
}
