package testing_basics

import "testing"

// TestIsPalindrome demonstrates basic testing with if-checks and subtests.
func TestIsPalindrome(t *testing.T) {
	// Subtest 1: Simple palindrome
	t.Run("simple palindrome", func(t *testing.T) {
		if !IsPalindrome("racecar") {
			t.Error("Expected 'racecar' to be a palindrome")
		}
	})

	// Subtest 2: Palindrome with capital letters and spaces
	t.Run("mixed case and spacing", func(t *testing.T) {
		if !IsPalindrome("A man a plan a canal Panama") {
			t.Error("Expected 'A man a plan a canal Panama' to be a palindrome")
		}
	})

	// Subtest 3: Non-palindrome
	t.Run("non-palindrome", func(t *testing.T) {
		if IsPalindrome("hello") {
			t.Error("Expected 'hello' NOT to be a palindrome")
		}
	})
}

// assertEqual is a test helper that compares two integers.
// By calling t.Helper(), any test failure will report the line number of
// the calling function (like TestFactorial) rather than the line number inside assertEqual.
func assertEqual(t *testing.T, got, expected int) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

// TestFactorial demonstrates using a test helper function.
func TestFactorial(t *testing.T) {
	t.Run("zero case", func(t *testing.T) {
		assertEqual(t, Factorial(0), 1)
	})

	t.Run("positive case", func(t *testing.T) {
		assertEqual(t, Factorial(5), 120)
	})

	t.Run("negative case", func(t *testing.T) {
		assertEqual(t, Factorial(-3), -1)
	})
}
