package errors_advanced

import (
	"testing"
	"fmt"
)


func TestValidationError(t *testing.T) {
	e := ValidationError{Field: "Email", Msg: "invalid format"}
	expected := "validation failed on field Email: invalid format"
	if e.Error() != expected {
		t.Errorf("expected %q, got %q", expected, e.Error())
	}
}

func TestCheckDatabaseError(t *testing.T) {
	wrapped := fmt.Errorf("db operation failed: %w", ErrConnection)
	if !CheckDatabaseError(wrapped) {
		t.Error("expected CheckDatabaseError to be true for wrapped ErrConnection")
	}
}

func TestSafeExecute(t *testing.T) {
	err := SafeExecute(func() {
		panic("something went wrong")
	})
	if err == nil || err.Error() != "recovered panic: something went wrong" {
		t.Errorf("unexpected recover error: %v", err)
	}
}
