//go:build ignore

package main

import (
	"testing"
)

// TestValueReceiver tests value receiver behavior
func TestValueReceiver(t *testing.T) {
	p := Person{Name: "TestUser", Age: 30}

	// Call value receiver method
	description := p.Describe()

	if description != "TestUser is 30 years old" {
		t.Errorf("expected 'TestUser is 30 years old', got %q", description)
	}
}

// TestValueReceiverIsAdult tests the IsAdult method
func TestValueReceiverIsAdult(t *testing.T) {
	tests := []struct {
		name     string
		age      int
		expected bool
	}{
		{"Adult", 30, true},
		{"Minor", 15, false},
		{"Exactly 18", 18, true},
		{"Just under 18", 17, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Person{Name: "Test", Age: tt.age}
			if p.IsAdult() != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, p.IsAdult())
			}
		})
	}
}

// TestPointerReceiverDeposit tests deposit functionality
func TestPointerReceiverDeposit(t *testing.T) {
	account := &Account2{
		ID:      1,
		Balance: 100.00,
		Owner:   "Test",
	}

	err := account.Deposit(50)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if account.Balance != 150.00 {
		t.Errorf("expected 150.00, got %.2f", account.Balance)
	}
}

// TestPointerReceiverWithdraw tests withdraw functionality
func TestPointerReceiverWithdraw(t *testing.T) {
	account := &Account2{
		ID:      2,
		Balance: 100.00,
		Owner:   "Test",
	}

	err := account.Withdraw(30)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if account.Balance != 70.00 {
		t.Errorf("expected 70.00, got %.2f", account.Balance)
	}
}

// TestPointerReceiverWithdrawInsufficientFunds tests withdrawal error
func TestPointerReceiverWithdrawInsufficientFunds(t *testing.T) {
	account := &Account2{
		ID:      3,
		Balance: 50.00,
		Owner:   "Test",
	}

	err := account.Withdraw(100)
	if err == nil {
		t.Error("expected error for insufficient funds")
	}

	if account.Balance != 50.00 {
		t.Errorf("balance should not change: expected 50.00, got %.2f", account.Balance)
	}
}

// TestMethodChaining tests the query builder chaining
func TestMethodChaining(t *testing.T) {
	query := NewQueryBuilder().
		From("users").
		Select("id", "name").
		Where("age > 18").
		Limit(10).
		Build()

	if query != "SELECT id, name FROM users WHERE age > 18 LIMIT 10" {
		t.Errorf("unexpected query: %s", query)
	}
}

// BenchmarkValueReceiverMethod benchmarks value receiver performance
func BenchmarkValueReceiverMethod(b *testing.B) {
	p := Person{Name: "Benchmark", Age: 30}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = p.Describe()
	}
}

// BenchmarkPointerReceiverMethod benchmarks pointer receiver performance
func BenchmarkPointerReceiverMethod(b *testing.B) {
	account := &Account2{
		ID:      1,
		Balance: 1000.00,
		Owner:   "Benchmark",
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = account.Deposit(100)
		account.Withdraw(100) // Reset
	}
}
