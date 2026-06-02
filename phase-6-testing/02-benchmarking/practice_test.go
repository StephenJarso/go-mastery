package benchmarking_basics

import (
	"testing"
)


func TestConcatenate(t *testing.T) {
	strs := []string{"a", "b", "c"}
	if ConcatenateStringsPlus(strs) != "abc" {
		t.Error("plus concatenation failed")
	}
	if ConcatenateStringsBuilder(strs) != "abc" {
		t.Error("builder concatenation failed")
	}
}

func BenchmarkConcatenatePlus(b *testing.B) {
	strs := make([]string, 100)
	for i := 0; i < 100; i++ {
		strs[i] = "hello"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcatenateStringsPlus(strs)
	}
}

func BenchmarkConcatenateBuilder(b *testing.B) {
	strs := make([]string, 100)
	for i := 0; i < 100; i++ {
		strs[i] = "hello"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcatenateStringsBuilder(strs)
	}
}
