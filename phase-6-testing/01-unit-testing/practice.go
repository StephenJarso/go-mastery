package testing_basics

import (
	"errors"
)


var _ = errors.New

// Exercise 1: Sum of Evens
// Return the sum of all even numbers in a slice.
func SumOfEvens(nums []int) int {
	// TODO: Implement
	return 0
}

// Exercise 2: Mock dependency testing
type Notifier interface {
	Send(msg string) error
}

func NotifyUser(n Notifier, msg string) error {
	if len(msg) == 0 {
		return errors.New("empty message")
	}
	return n.Send(msg)
}
