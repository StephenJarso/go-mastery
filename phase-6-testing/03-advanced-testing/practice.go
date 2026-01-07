package advanced_testing

import (
	"errors"
	"fmt"
)

// NotificationService defines the channel for alerts.
type NotificationService interface {
	SendNotification(userID int, message string) error
}

// Order represents a sales checkout.
type Order struct {
	OrderID int
	UserID  int
	Amount  float64
}

// OrderProcessor manages the checkout pipeline.
type OrderProcessor struct {
	Store    UserStore
	Notifier NotificationService
}

// ProcessOrder verifies the user exists and sends a billing notification.
func (op *OrderProcessor) ProcessOrder(order *Order) error {
	if order == nil {
		return errors.New("order cannot be nil")
	}

	// 1. Verify User Exists
	_, err := op.Store.Get(order.UserID)
	if err != nil {
		return fmt.Errorf("user check failed: %w", err)
	}

	// 2. Send Notification
	msg := fmt.Sprintf("Order %d processed for %.2f", order.OrderID, order.Amount)
	err = op.Notifier.SendNotification(order.UserID, msg)
	if err != nil {
		return fmt.Errorf("notification failed: %w", err)
	}

	return nil
}
