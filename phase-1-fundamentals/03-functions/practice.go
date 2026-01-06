package functions

// PRACTICE EXERCISE: Collection Filter
// Write a function that filters a slice of integers based on a predicate callback.
// The function signature should accept:
// - a slice of integers
// - a callback function: func(int) bool
// Returns a new slice of integers containing only elements for which the predicate returns true.

func Filter(numbers []int, predicate func(int) bool) []int {
	result := make([]int, 0)
	for _, n := range numbers {
		if predicate(n) {
			result = append(result, n)
		}
	}
	return result
}
