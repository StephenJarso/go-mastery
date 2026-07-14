package benchmarking_basics

import (
	"reflect"
	"testing"
)

func TestFilterAndSquareCorrectness(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{4, 16, 36, 64, 100}

	gotSlow := FilterAndSquareSlow(input)
	gotFast := FilterAndSquareFast(input)

	if !reflect.DeepEqual(gotSlow, expected) {
		t.Errorf("FilterAndSquareSlow() = %v; expected %v", gotSlow, expected)
	}

	if !reflect.DeepEqual(gotFast, expected) {
		t.Errorf("FilterAndSquareFast() = %v; expected %v", gotFast, expected)
	}
}

func BenchmarkFilterAndSquare(b *testing.B) {
	// Generate a large slice of integers to make allocation costs noticeable
	input := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		input[i] = i
	}

	b.Run("Slow", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = FilterAndSquareSlow(input)
		}
	})

	b.Run("Fast", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = FilterAndSquareFast(input)
		}
	})
}
