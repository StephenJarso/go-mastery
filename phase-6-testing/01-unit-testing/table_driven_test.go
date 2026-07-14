package testing_basics

import "testing"

// TestCalculateTax demonstrates the idiomatic table-driven test pattern in Go.
func TestCalculateTax(t *testing.T) {
	// Define the test table (slice of structs)
	tests := []struct {
		name     string
		income   float64
		expected float64
	}{
		{
			name:     "Zero Income",
			income:   0,
			expected: 0,
		},
		{
			name:     "Negative Income",
			income:   -500,
			expected: 0,
		},
		{
			name:     "First Bracket Limit",
			income:   10000,
			expected: 1000, // 10,000 * 0.10
		},
		{
			name:     "Second Bracket Midpoint",
			income:   30000,
			expected: 4000, // 1000 + (20,000 * 0.15)
		},
		{
			name:     "Second Bracket Limit",
			income:   50000,
			expected: 7000, // 1000 + 6000
		},
		{
			name:     "Third Bracket",
			income:   100000,
			expected: 19500, // 1000 + 6000 + (50,000 * 0.25)
		},
	}

	// Loop over each test case in the table
	for _, tc := range tests {
		// Run each test case in a separate subtest context
		t.Run(tc.name, func(t *testing.T) {
			got := CalculateTax(tc.income)
			if got != tc.expected {
				t.Errorf("CalculateTax(%.2f) = %.2f; expected %.2f", tc.income, got, tc.expected)
			}
		})
	}
}
