package errors_advanced

import (
	"fmt"
	"testing"
)

func TestValidationError(t *testing.T) {
	e := ValidationError{Field: "Email", Msg: "invalid format"}
	if e.Error() == "" {
		t.Skip("Exercise ValidationError not implemented yet")
	}
	expected := "validation failed on field Email: invalid format"
	if e.Error() != expected {
		t.Errorf("expected %q, got %q", expected, e.Error())
	}
}

func TestCheckDatabaseError(t *testing.T) {
	wrapped := fmt.Errorf("db operation failed: %w", ErrConnection)
	if !CheckDatabaseError(wrapped) {
		t.Skip("Exercise CheckDatabaseError not implemented yet")
	}
}

func TestSafeExecute(t *testing.T) {
	err := SafeExecute(func() {
		panic("something went wrong")
	})
	if err == nil {
		t.Skip("Exercise SafeExecute not implemented yet")
	}
	if err.Error() != "recovered panic: something went wrong" {
		t.Errorf("unexpected recover error: %v", err)
	}
}
