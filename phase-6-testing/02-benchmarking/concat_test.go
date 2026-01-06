package benchmarking_basics

import (
	"fmt"
	"testing"
)

// generateStrings generates a slice of count strings, each of length length.
func generateStrings(count, length int) []string {
	res := make([]string, count)
	str := ""
	for i := 0; i < length; i++ {
		str += "a"
	}
	for i := 0; i < count; i++ {
		res[i] = str
	}
	return res
}

// TestConcatCorrectness validates both implementations yield the same result.
func TestConcatCorrectness(t *testing.T) {
	input := []string{"hello", " ", "world", "!"}
	expected := "hello world!"

	if got := ConcatNaive(input); got != expected {
		t.Errorf("ConcatNaive() = %q; expected %q", got, expected)
	}
	if got := ConcatBuilder(input); got != expected {
		t.Errorf("ConcatBuilder() = %q; expected %q", got, expected)
	}
}

// BenchmarkConcat runs a suite of sub-benchmarks using different input sizes.
func BenchmarkConcat(b *testing.B) {
	sizes := []int{10, 100, 1000}

	for _, size := range sizes {
		input := generateStrings(size, 10) // size strings, each 10 bytes long

		b.Run(fmt.Sprintf("Naive-%d", size), func(b *testing.B) {
			b.ReportAllocs() // Automatically reports allocation counts and bytes
			b.ResetTimer()   // Exclude generation and setup time
			for i := 0; i < b.N; i++ {
				_ = ConcatNaive(input)
			}
		})

		b.Run(fmt.Sprintf("Builder-%d", size), func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = ConcatBuilder(input)
			}
		})
	}
}
