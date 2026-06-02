package testing_basics

import (
	"testing"
)


type MockNotifier struct {
	SentMsg string
	SendErr error
}

func (m *MockNotifier) Send(msg string) error {
	m.SentMsg = msg
	return m.SendErr
}

func TestSumOfEvens(t *testing.T) {
	res := SumOfEvens([]int{1, 2, 3, 4, 5, 6})
	if res != 12 {
		t.Errorf("expected 12, got %d", res)
	}
}

func TestNotifyUser(t *testing.T) {
	mock := &MockNotifier{}
	err := NotifyUser(mock, "hello")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if mock.SentMsg != "hello" {
		t.Errorf("expected mock sent message to be 'hello', got %q", mock.SentMsg)
	}

	err = NotifyUser(mock, "")
	if err == nil {
		t.Error("expected error for empty message")
	}
}
