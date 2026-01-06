package benchmarking_basics

// FilterAndSquareSlow filters even numbers and squares them.
// This is slower because it starts with zero capacity, causing the Go runtime
// to continuously allocate and copy backing arrays as new elements are appended.
func FilterAndSquareSlow(input []int) []int {
	var result []int // zero capacity
	for _, v := range input {
		if v%2 == 0 {
			result = append(result, v*v)
		}
	}
	return result
}

// FilterAndSquareFast optimizes the operation by pre-allocating the maximum possible
// capacity in the slice initialization, preventing runtime memory re-allocations.
func FilterAndSquareFast(input []int) []int {
	// Since we are filtering, the final slice will have at most len(input) items.
	// Pre-allocating capacity avoids growing the underlying array in loops.
	result := make([]int, 0, len(input))
	for _, v := range input {
		if v%2 == 0 {
			result = append(result, v*v)
		}
	}
	return result
}
