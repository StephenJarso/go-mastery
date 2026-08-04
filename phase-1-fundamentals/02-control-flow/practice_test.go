package controlflow

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	res := FizzBuzz(1, 15)
	if res == nil {
		t.Skip("Exercise FizzBuzz not implemented yet")
	}
	expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	if len(res) != len(expected) {
		t.Fatalf("length mismatch: got %d, expected %d", len(res), len(expected))
	}
	for i, v := range res {
		if v != expected[i] {
			t.Errorf("at index %d: got %s, expected %s", i, v, expected[i])
		}
	}
}

func TestFindPrimes(t *testing.T) {
	res := FindPrimes(20)
	if res == nil {
		t.Skip("Exercise FindPrimes not implemented yet")
	}
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19}
	if len(res) != len(expected) {
		t.Fatalf("length mismatch: got %d, expected %d", len(res), len(expected))
	}
	for i, v := range res {
		if v != expected[i] {
			t.Errorf("at index %d: got %d, expected %d", i, v, expected[i])
		}
	}
}

func TestCountCharFrequency(t *testing.T) {
	res := CountCharFrequency("hello")
	if res == nil {
		t.Skip("Exercise CountCharFrequency not implemented yet")
	}
	expected := map[rune]int{'h': 1, 'e': 1, 'l': 2, 'o': 1}
	if len(res) != len(expected) {
		t.Errorf("map size mismatch: got %d, expected %d", len(res), len(expected))
	}
	for k, v := range expected {
		if res[k] != v {
			t.Errorf("for key %c: got %d, expected %d", k, res[k], v)
		}
	}
}
