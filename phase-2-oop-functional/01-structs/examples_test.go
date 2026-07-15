//go:build ignore

package main

import (
	"encoding/json"
	"testing"
)

// Test suite for struct examples

// TestStructCreation tests different ways to create structs
func TestStructCreation(t *testing.T) {
	tests := []struct {
		name     string
		person   Person
		expected string
	}{
		{
			name:     "named fields",
			person:   Person{Name: "Alice", Age: 30, Email: "alice@example.com"},
			expected: "Alice",
		},
		{
			name:     "positional",
			person:   Person{"Bob", 25, "bob@example.com"},
			expected: "Bob",
		},
		{
			name:     "partial",
			person:   Person{Name: "Charlie"},
			expected: "Charlie",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.person.Name != tt.expected {
				t.Errorf("got %s, want %s", tt.person.Name, tt.expected)
			}
		})
	}
}

// TestZeroValues tests that uninitialized fields have zero values
func TestZeroValues(t *testing.T) {
	var p Person

	if p.Name != "" {
		t.Errorf("expected empty string, got %q", p.Name)
	}

	if p.Age != 0 {
		t.Errorf("expected 0, got %d", p.Age)
	}

	if p.Email != "" {
		t.Errorf("expected empty string, got %q", p.Email)
	}
}

// TestFieldAccess tests accessing and modifying fields
func TestFieldAccess(t *testing.T) {
	p := Person{Name: "Diana", Age: 28, Email: "diana@example.com"}

	// Test access
	if p.Name != "Diana" {
		t.Errorf("expected Diana, got %s", p.Name)
	}

	// Test modification
	p.Age = 29
	if p.Age != 29 {
		t.Errorf("expected 29, got %d", p.Age)
	}
}

// TestPointersToStructs tests pointer behavior
func TestPointersToStructs(t *testing.T) {
	original := Person{Name: "Eve", Age: 30, Email: "eve@example.com"}
	pointer := &original

	// Modify through pointer
	pointer.Age = 31

	// Original should be modified
	if original.Age != 31 {
		t.Errorf("expected 31, got %d", original.Age)
	}
}

// TestStructComparison tests comparing structs
func TestStructComparison(t *testing.T) {
	p1 := Person{Name: "Frank", Age: 40, Email: "frank@example.com"}
	p2 := Person{Name: "Frank", Age: 40, Email: "frank@example.com"}
	p3 := Person{Name: "Grace", Age: 40, Email: "grace@example.com"}

	if p1 != p2 {
		t.Error("expected p1 == p2")
	}

	if p1 == p3 {
		t.Error("expected p1 != p3")
	}
}

// TestJSONMarshaling tests JSON marshaling with tags
func TestJSONMarshaling(t *testing.T) {
	p := Person{
		ID:        1,
		FirstName: "Henry",
		LastName:  "Johnson",
		Email:     "henry@example.com",
	}

	data, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var decoded Person
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if decoded.FirstName != p.FirstName {
		t.Errorf("expected %s, got %s", p.FirstName, decoded.FirstName)
	}
}

// TestEmbedding tests struct embedding
func TestEmbedding(t *testing.T) {
	s := Student{
		ID:   1,
		Name: "Iris",
		Address: Address{
			Street: "789 Pine Rd",
			City:   "Boston",
			State:  "MA",
			Zip:    "02101",
		},
		GPA: 3.8,
	}

	// Test that promoted fields are accessible
	if s.Street != "789 Pine Rd" {
		t.Errorf("expected promoted field access")
	}

	if s.City != "Boston" {
		t.Errorf("expected promoted field access")
	}
}

// BenchmarkStructCreation benchmarks struct creation
func BenchmarkStructCreation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Person{
			Name:  "Jack",
			Age:   25,
			Email: "jack@example.com",
		}
	}
}

// BenchmarkFieldAccess benchmarks field access
func BenchmarkFieldAccess(b *testing.B) {
	p := Person{Name: "Kate", Age: 30, Email: "kate@example.com"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.Name
		_ = p.Age
		_ = p.Email
	}
}

// BenchmarkJSONMarshaling benchmarks JSON marshaling
func BenchmarkJSONMarshaling(b *testing.B) {
	p := Person{
		ID:        1,
		FirstName: "Leo",
		LastName:  "Smith",
		Email:     "leo@example.com",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(p)
	}
}
