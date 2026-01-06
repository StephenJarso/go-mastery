package functions

import (
	"errors"
)

// Go functions can return multiple values, including errors.
// They also support named return parameters and closures.

// Divide returns the quotient and remainder of division.
// Demonstrates returning multiple values.
func Divide(dividend, divisor int) (int, int, error) {
	if divisor == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return dividend / divisor, dividend % divisor, nil
}

// CalculateStats returns stats about a slice.
// Demonstrates named return parameters.
func CalculateStats(nums []int) (min int, max int) {
	if len(nums) == 0 {
		return 0, 0
	}
	min = nums[0]
	max = nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	// A bare return statement returns the current values of the named return variables.
	return
}

// MakeMultiplier returns a function closure that multiplies an integer input by factor.
func MakeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
