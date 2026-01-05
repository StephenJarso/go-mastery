package testing_basics

import (
	"strings"
	"unicode"
)

// IsPalindrome checks if a string reads the same forward and backward.
// It ignores casing and non-alphanumeric characters.
func IsPalindrome(s string) bool {
	var clean []rune
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			clean = append(clean, r)
		}
	}

	n := len(clean)
	for i := 0; i < n/2; i++ {
		if clean[i] != clean[n-1-i] {
			return false
		}
	}
	return true
}

// Factorial calculates the factorial of n.
// Returns -1 for negative inputs (invalid).
func Factorial(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
