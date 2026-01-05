package testing_basics

import (
	"errors"
	"strings"
)

var (
	ErrEmailEmpty      = errors.New("email cannot be empty")
	ErrEmailMissingAt  = errors.New("email is missing '@'")
	ErrEmailMultipleAt = errors.New("email contains multiple '@'")
	ErrEmailEmptyLocal = errors.New("email local part cannot be empty")
	ErrEmailInvalidDom = errors.New("email domain part is invalid")
)

// SumOfEvens sums all even integers in a slice.
func SumOfEvens(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		if n%2 == 0 {
			sum += n
		}
	}
	return sum
}

// ValidateEmail performs basic validation on email address formats.
// Returns specific sentinel errors for different invalid states.
func ValidateEmail(email string) error {
	if email == "" {
		return ErrEmailEmpty
	}

	parts := strings.Split(email, "@")
	if len(parts) < 2 {
		return ErrEmailMissingAt
	}
	if len(parts) > 2 {
		return ErrEmailMultipleAt
	}

	local := parts[0]
	domain := parts[1]

	if local == "" {
		return ErrEmailEmptyLocal
	}

	// Domain must contain a dot and have letters before/after it
	dotIndex := strings.Index(domain, ".")
	if dotIndex <= 0 || dotIndex == len(domain)-1 {
		return ErrEmailInvalidDom
	}

	return nil
}
