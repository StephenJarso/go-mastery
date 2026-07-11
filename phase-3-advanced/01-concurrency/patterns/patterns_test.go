package patterns

import (
	"context"
	"testing"
	"time"
)

// TestPipeline verifies the pipeline stages produce expected output.
func TestPipeline(t *testing.T) {
	intsChan := generate(1, 2, 3)
	doubledChan := double(intsChan)
	stringifiedChan := stringify(doubledChan)

	expected := []string{"Number: 2", "Number: 4", "Number: 6"}
	idx := 0

	for result := range stringifiedChan {
		if idx >= len(expected) {
			t.Fatalf("Got more results than expected: %q", result)
		}
		if result != expected[idx] {
			t.Errorf("Expected %q, got %q", expected[idx], result)
		}
		idx++
	}

	if idx != len(expected) {
		t.Errorf("Expected %d results, got %d", len(expected), idx)
	}
}

// TestFanInFanOut verifies the merge function correctly merges channels.
func TestFanInFanOut(t *testing.T) {
	in := generator(1, 2, 3)
	w1 := squareWorker(in)
	w2 := squareWorker(in)

	results := merge(w1, w2)
	sum := 0

	for res := range results {
		sum += res
	}

	// 1*1 + 2*2 + 3*3 = 1 + 4 + 9 = 14
	expected := 14
	if sum != expected {
		t.Errorf("Expected sum of squared values to be %d, got %d", expected, sum)
	}
}

// TestContextTimeout verifies custom timeouts work.
func TestContextTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	// Slow process
	select {
	case <-time.After(50 * time.Millisecond):
		t.Error("Expected timeout, but process completed")
	case <-ctx.Done():
		// Success: context fended off long duration
	}
}
