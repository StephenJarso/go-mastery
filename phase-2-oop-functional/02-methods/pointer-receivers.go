package main

import (
	"fmt"
)

// PointerReceiverExamples demonstrates methods with POINTER receivers.
// Pointer receivers work on the ORIGINAL struct.
// Modifications affect the original.

// Account2 is our example struct for pointer receiver methods
type Account2 struct {
	ID      int
	Balance float64
	Owner   string
}

// Deposit is a method with a POINTER receiver.
// It modifies the receiver (adds money to account).
func (a *Account2) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("deposit amount must be positive")
	}
	a.Balance += amount
	return nil
}

// Withdraw is a method with a POINTER receiver.
// It modifies the receiver (removes money from account).
func (a *Account2) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("withdraw amount must be positive")
	}
	if a.Balance < amount {
		return fmt.Errorf("insufficient funds")
	}
	a.Balance -= amount
	return nil
}

// Transfer transfers money to another account
// Uses pointer receiver because it modifies this account
func (a *Account2) Transfer(other *Account2, amount float64) error {
	if err := a.Withdraw(amount); err != nil {
		return err
	}
	if err := other.Deposit(amount); err != nil {
		// Withdraw was already done, need to undo it
		a.Deposit(amount)
		return err
	}
	return nil
}

// BasicPointerReceiver shows how pointer receivers work
func BasicPointerReceiver() {
	fmt.Println("=== Basic Pointer Receiver ===")

	account := Account2{
		ID:      1001,
		Balance: 1000.00,
		Owner:   "Alice",
	}

	fmt.Printf("Initial account: %+v\n", account)
	fmt.Printf("Balance: $%.2f\n", account.Balance)

	// Call method with pointer receiver
	account.Deposit(500)
	fmt.Printf("After $500 deposit: $%.2f\n", account.Balance)

	// The original account IS modified
	fmt.Printf("Final account: %+v\n", account)
	fmt.Println("Original modified by pointer receiver method")
}

// PointerReceiverModification shows modifications persist
func PointerReceiverModification() {
	fmt.Println("\n=== Pointer Receiver Modifications ===")

	account := Account2{
		ID:      2002,
		Balance: 500.00,
		Owner:   "Bob",
	}

	fmt.Printf("Initial: $%.2f\n", account.Balance)

	// Multiple modifications
	account.Deposit(200)   // 500 + 200 = 700
	format.Printf("After deposit: $%.2f\n", account.Balance)

	account.Withdraw(150)  // 700 - 150 = 550
	fmt.Printf("After withdrawal: $%.2f\n", account.Balance)

	account.Deposit(100)   // 550 + 100 = 650
	fmt.Printf("After second deposit: $%.2f\n", account.Balance)
}

// PointerReceiverErrorHandling shows error handling with pointer receivers
func PointerReceiverErrorHandling() {
	fmt.Println("\n=== Error Handling with Pointer Receivers ===")

	account := Account2{
		ID:      3003,
		Balance: 300.00,
		Owner:   "Charlie",
	}

	fmt.Printf("Initial balance: $%.2f\n", account.Balance)

	// Try to withdraw more than balance
	err := account.Withdraw(500)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Balance unchanged: $%.2f\n", account.Balance)

	// Try to deposit negative amount
	err = account.Deposit(-100)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Balance unchanged: $%.2f\n", account.Balance)

	// Successful withdrawal
	err = account.Withdraw(100)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Successfully withdrew $100: $%.2f\n", account.Balance)
	}
}

// Transfer demonstrates more complex pointer receiver interactions
func TransferExample() {
	fmt.Println("\n=== Transfer Between Accounts ===")

	alice := Account2{
		ID:      1,
		Balance: 1000.00,
		Owner:   "Alice",
	}

	bob := Account2{
		ID:      2,
		Balance: 500.00,
		Owner:   "Bob",
	}

	fmt.Printf("Before transfer:\n")
	fmt.Printf("  Alice: $%.2f\n", alice.Balance)
	fmt.Printf("  Bob: $%.2f\n", bob.Balance)

	// Transfer $300 from Alice to Bob
	err := alice.Transfer(&bob, 300)
	if err != nil {
		fmt.Printf("Transfer error: %v\n", err)
		return
	}

	fmt.Printf("\nAfter $300 transfer from Alice to Bob:\n")
	fmt.Printf("  Alice: $%.2f\n", alice.Balance)
	fmt.Printf("  Bob: $%.2f\n", bob.Balance)

	// Try to transfer more than available
	fmt.Printf("\nTrying to transfer $2000 from Alice to Bob...\n")
	err = alice.Transfer(&bob, 2000)
	if err != nil {
		fmt.Printf("Transfer error: %v\n", err)
	}

	fmt.Printf("\nAfter failed transfer:\n")
	fmt.Printf("  Alice: $%.2f (unchanged)\n", alice.Balance)
	fmt.Printf("  Bob: $%.2f (unchanged)\n", bob.Balance)
}

// Person2 for demonstrating state modification
type Person2 struct {
	Name string
	Age  int
}

// HaveBirthday increments age (modifies receiver)
func (p *Person2) HaveBirthday() {
	p.Age++
}

// Rename changes the person's name
func (p *Person2) Rename(newName string) {
	if newName != "" {
		p.Name = newName
	}
}

func StateModificationExample() {
	fmt.Println("\n=== State Modification Example ===")

	person := Person2{Name: "Diana", Age: 30}

	fmt.Printf("Initial: %s, age %d\n", person.Name, person.Age)

	person.HaveBirthday()
	fmt.Printf("After birthday: %s, age %d\n", person.Name, person.Age)

	person.Rename("Diane")
	fmt.Printf("After rename: %s, age %d\n", person.Name, person.Age)

	person.HaveBirthday()
	person.HaveBirthday()
	fmt.Printf("After 2 more birthdays: %s, age %d\n", person.Name, person.Age)
}

// WhenToUsePointerReceivers explains the decision logic
func WhenToUsePointerReceivers() {
	fmt.Println("\n=== When to Use Pointer Receivers ===")
	fmt.Println(`
Use pointer receivers when:

1. Method modifies the receiver
   - Deposit/Withdraw from account
   - Update state
   - Modify fields

2. Receiver is large
   - Avoid copying overhead
   - Array, large struct, etc
   - Performance matters

3. Consistency
   - If ANY method uses pointer receiver
   - Use pointer receiver for ALL methods on type
   - Consistent API

4. Implementing interfaces
   - Method needs to modify to satisfy interface
   - Must use pointer receiver

Advantages:
- Modify original receiver
- Efficient for large structs
- Clear intent: "this changes state"

Disadvantages:
- Not safe for concurrent use
- Can't call on non-addressable values
- Less explicit (need to remember it modifies)
`)
}

// ReceiverConsistency shows why consistency matters
func ReceiverConsistency() {
	fmt.Println("\n=== Receiver Consistency ===")
	fmt.Println(`
Go convention:
If ANY method on a type uses a pointer receiver,
use pointer receiver for ALL methods.

Why?
- Consistency
- Predictability
- Method sets are clearer

Example:
  type Account struct { ... }
  
  func (a *Account) Deposit(x float64) { ... }  // pointer
  func (a *Account) Withdraw(x float64) { ... } // pointer
  func (a *Account) GetBalance() { ... }        // pointer (not value)
  
Not:
  func (a Account) GetBalance() { ... }         // inconsistent!
`)
}

func main() {
	BasicPointerReceiver()
	PointerReceiverModification()
	PointerReceiverErrorHandling()
	TransferExample()
	StateModificationExample()
	WhenToUsePointerReceivers()
	ReceiverConsistency()

	fmt.Println("\n=== Pointer Receiver Examples Complete ===")
}
