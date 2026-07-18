package mapsdemo

import (
	"reflect"
	"sort"
	"testing"
)

func TestWordFrequency(t *testing.T) {
	tests := []struct {
		input    []string
		expected map[string]int
	}{
		{
			input:    []string{"apple", "banana", "apple", "cherry", "banana", "apple"},
			expected: map[string]int{"apple": 3, "banana": 2, "cherry": 1},
		},
		{
			input:    []string{},
			expected: map[string]int{},
		},
		{
			input:    []string{"hello"},
			expected: map[string]int{"hello": 1},
		},
	}

	for _, tt := range tests {
		res := WordFrequency(tt.input)
		if !reflect.DeepEqual(res, tt.expected) {
			t.Errorf("WordFrequency(%v) = %v; expected %v", tt.input, res, tt.expected)
		}
	}
}

func TestGroupByGrade(t *testing.T) {
	tests := []struct {
		input    map[string]string
		expected map[string][]string
	}{
		{
			input: map[string]string{
				"Alice":   "A",
				"Bob":     "B",
				"Charlie": "A",
				"David":   "C",
				"Eve":     "B",
			},
			expected: map[string][]string{
				"A": {"Alice", "Charlie"},
				"B": {"Bob", "Eve"},
				"C": {"David"},
			},
		},
		{
			input:    map[string]string{},
			expected: map[string][]string{},
		},
	}

	for _, tt := range tests {
		res := GroupByGrade(tt.input)
		// Since map iteration is random, the slice values could be in any order.
		// Sort slices in both res and expected to ensure consistent comparison.
		for k, v := range res {
			sort.Strings(v)
			res[k] = v
		}
		for k, v := range tt.expected {
			sort.Strings(v)
			tt.expected[k] = v
		}
		if !reflect.DeepEqual(res, tt.expected) {
			t.Errorf("GroupByGrade(%v) = %v; expected %v", tt.input, res, tt.expected)
		}
	}
}

func TestMergeMaps(t *testing.T) {
	a := map[string]int{"apple": 5, "banana": 10}
	b := map[string]int{"banana": 5, "cherry": 15}
	expected := map[string]int{"apple": 5, "banana": 15, "cherry": 15}

	res := MergeMaps(a, b)
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("MergeMaps(%v, %v) = %v; expected %v", a, b, res, expected)
	}

	// Verify original maps were not modified
	if a["banana"] != 10 {
		t.Errorf("Original map 'a' was modified: a['banana'] = %d", a["banana"])
	}
	if b["banana"] != 5 {
		t.Errorf("Original map 'b' was modified: b['banana'] = %d", b["banana"])
	}
}
