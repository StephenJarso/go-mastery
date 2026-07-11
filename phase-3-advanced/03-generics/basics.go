package generics

import (
	"fmt"
)

// Go Generics (introduced in Go 1.18) allow writing code with type parameters,
// enabling type-safe code reuse without the performance cost or type-safety issues
// of reflection or empty interfaces (interface{}).

// 1. Generic Functions with Built-in Constraints
// The 'any' constraint is an alias for the empty interface 'interface{}'.
// It means the type parameter can be any type.

// PrintSlice prints elements of a slice of any type.
func PrintSlice[T any](slice []T) {
	for _, v := range slice {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

// The 'comparable' constraint is a built-in constraint that permits any type
// whose values can be compared using == or !=.
// This includes numbers, strings, pointers, channels, and structs/arrays of comparable types.

// FindIndex returns the index of the target in a slice, or -1 if not found.
func FindIndex[T comparable](slice []T, target T) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

// 2. Custom Type Constraints
// We define custom constraints using interfaces. We can restrict types to a specific
// set of underlying types using the type union operator '|'.

// Number constraint permits only numeric types.
type Number interface {
	int | int32 | int64 | float32 | float64
}

// Sum sums the elements of a slice containing numeric values.
func Sum[T Number](slice []T) T {
	var total T
	for _, v := range slice {
		total += v
	}
	return total
}

// 3. Underling Type Approximation (~)
// If we want a constraint to match custom types whose underlying type matches,
// we use the tilde symbol '~'. For example, '~int' matches 'type CustomInt int'.

type CustomInt int

type Integer interface {
	~int | ~int64
}

// Max returns the maximum of two integers, including custom types based on int.
func Max[T Integer](a, b T) T {
	if a > b {
		return a
	}
	return b
}
