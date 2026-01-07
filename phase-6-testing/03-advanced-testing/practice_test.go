package advanced_testing

import (
	"errors"
	"testing"
)

// MockNotificationService is a test double for NotificationService.
type MockNotificationService struct {
	LastUserID  int
	LastMessage string
	SendError   error
	SendCalls   int
}

func (m *MockNotificationService) SendNotification(userID int, message string) error {
	m.SendCalls++
	m.LastUserID = userID
	m.LastMessage = message
	return m.SendError
}

func TestProcessOrder(t *testing.T) {
	t.Run("successful processing and notification", func(t *testing.T) {
		mockUser := &User{ID: 1, Name: "Alice"}
		mockStore := &MockUserStore{GetResult: mockUser}
		mockNotifier := &MockNotificationService{}

		processor := &OrderProcessor{
			Store:    mockStore,
			Notifier: mockNotifier,
		}

		order := &Order{OrderID: 123, UserID: 1, Amount: 99.95}
		err := processor.ProcessOrder(order)
		if err != nil {
			t.Fatalf("unexpected processing error: %v", err)
		}

		if mockStore.GetCalls != 1 {
			t.Errorf("expected 1 Get call, got %d", mockStore.GetCalls)
		}

		if mockNotifier.SendCalls != 1 {
			t.Errorf("expected 1 Notification send call, got %d", mockNotifier.SendCalls)
		}

		expectedMsg := "Order 123 processed for 99.95"
		if mockNotifier.LastMessage != expectedMsg {
			t.Errorf("expected msg %q, got %q", expectedMsg, mockNotifier.LastMessage)
		}

		if mockNotifier.LastUserID != 1 {
			t.Errorf("expected notification sent to User 1, got %d", mockNotifier.LastUserID)
		}
	})

	t.Run("non-existent user", func(t *testing.T) {
		mockStore := &MockUserStore{GetError: errors.New("user not found")}
		mockNotifier := &MockNotificationService{}

		processor := &OrderProcessor{
			Store:    mockStore,
			Notifier: mockNotifier,
		}

		order := &Order{OrderID: 123, UserID: 999, Amount: 50.00}
		err := processor.ProcessOrder(order)
		if err == nil {
			t.Fatal("expected error due to missing user, got nil")
		}

		if mockNotifier.SendCalls != 0 {
			t.Errorf("expected 0 notifications sent for non-existent user, got %d", mockNotifier.SendCalls)
		}
	})

	t.Run("notification failure", func(t *testing.T) {
		mockUser := &User{ID: 1}
		mockStore := &MockUserStore{GetResult: mockUser}
		notifyErr := errors.New("SMS gateway down")
		mockNotifier := &MockNotificationService{SendError: notifyErr}

		processor := &OrderProcessor{
			Store:    mockStore,
			Notifier: mockNotifier,
		}

		order := &Order{OrderID: 123, UserID: 1, Amount: 15.00}
		err := processor.ProcessOrder(order)
		if !errors.Is(err, notifyErr) {
			t.Fatalf("expected error wrapping notification failure, got: %v", err)
		}

		if mockNotifier.SendCalls != 1 {
			t.Errorf("expected 1 Notification send call attempt, got %d", mockNotifier.SendCalls)
		}
	})
}
