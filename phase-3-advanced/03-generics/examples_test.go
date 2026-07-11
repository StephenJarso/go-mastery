package generics

import (
	"strconv"
	"testing"
)

func TestFindIndex(t *testing.T) {
	// 1. Comparable integers
	ints := []int{10, 20, 30, 40, 50}
	if idx := FindIndex(ints, 30); idx != 2 {
		t.Errorf("Expected index 2, got %d", idx)
	}
	if idx := FindIndex(ints, 100); idx != -1 {
		t.Errorf("Expected index -1, got %d", idx)
	}

	// 2. Comparable strings
	strings := []string{"apple", "banana", "cherry"}
	if idx := FindIndex(strings, "banana"); idx != 1 {
		t.Errorf("Expected index 1, got %d", idx)
	}
}

func TestSum(t *testing.T) {
	// 1. Sum integers
	ints := []int{1, 2, 3, 4}
	if got := Sum(ints); got != 10 {
		t.Errorf("Expected sum 10, got %d", got)
	}

	// 2. Sum floats
	floats := []float64{1.5, 2.5, 3.0}
	if got := Sum(floats); got != 7.0 {
		t.Errorf("Expected sum 7.0, got %f", got)
	}
}

func TestMax(t *testing.T) {
	// 1. Regular integers
	if got := Max(10, 20); got != 20 {
		t.Errorf("Expected Max(10, 20) = 20, got %d", got)
	}

	// 2. CustomInt approximation
	var val1 CustomInt = 50
	var val2 CustomInt = 30
	if got := Max(val1, val2); got != 50 {
		t.Errorf("Expected Max(50, 30) = 50, got %d", got)
	}
}

func TestMapAndFilter(t *testing.T) {
	// Test Map: double integers and convert to strings
	ints := []int{1, 2, 3}
	doubled := Map(ints, func(x int) int { return x * 2 })
	if doubled[0] != 2 || doubled[1] != 4 || doubled[2] != 6 {
		t.Errorf("Expected Map doubling to be [2, 4, 6], got %v", doubled)
	}

	stringified := Map(ints, func(x int) string { return strconv.Itoa(x) })
	if stringified[0] != "1" || stringified[1] != "2" || stringified[2] != "3" {
		t.Errorf("Expected Map string conversion to be ['1', '2', '3'], got %v", stringified)
	}

	// Test Filter: extract even numbers
	nums := []int{1, 2, 3, 4, 5, 6}
	evens := Filter(nums, func(x int) bool { return x%2 == 0 })
	if len(evens) != 3 || evens[0] != 2 || evens[1] != 4 || evens[2] != 6 {
		t.Errorf("Expected Filter evens to be [2, 4, 6], got %v", evens)
	}
}

func TestStack(t *testing.T) {
	// 1. String Stack
	stack := NewStack[string]()
	if stack.Size() != 0 {
		t.Errorf("Expected stack size 0, got %d", stack.Size())
	}

	stack.Push("first")
	stack.Push("second")
	if stack.Size() != 2 {
		t.Errorf("Expected stack size 2, got %d", stack.Size())
	}

	val, ok := stack.Pop()
	if !ok || val != "second" {
		t.Errorf("Expected Pop() to return 'second', true; got %q, %t", val, ok)
	}

	val, ok = stack.Pop()
	if !ok || val != "first" {
		t.Errorf("Expected Pop() to return 'first', true; got %q, %t", val, ok)
	}

	_, ok = stack.Pop()
	if ok {
		t.Error("Expected third Pop() to be empty (false), got true")
	}
}
