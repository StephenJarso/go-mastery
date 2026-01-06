package benchmarking_basics

import "strings"

// ConcatNaive concatenates a slice of strings using the simple + operator in a loop.
// This is inefficient because strings are immutable in Go, so each addition
// creates a brand-new string in memory and copies the previous content.
func ConcatNaive(parts []string) string {
	res := ""
	for _, p := range parts {
		res += p
	}
	return res
}

// ConcatBuilder concatenates a slice of strings using strings.Builder.
// This is highly efficient because it uses an internal byte buffer that grows dynamically.
// Pre-allocating buffer memory using Grow() minimizes allocations.
func ConcatBuilder(parts []string) string {
	var builder strings.Builder

	// Calculate total length to allocate once
	totalLen := 0
	for _, p := range parts {
		totalLen += len(p)
	}

	builder.Grow(totalLen) // Allocate memory once
	for _, p := range parts {
		builder.WriteString(p)
	}
	return builder.String()
}
