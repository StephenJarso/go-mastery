package main

import (
	"fmt"
)

// ValueReceiverExamples demonstrates methods with VALUE receivers.
// Value receivers work on a COPY of the struct.
// The original struct is NOT modified.

// Person is a struct for examples
type Person struct {
	Name string
	Age  int
}

// Describe is a method with a VALUE receiver.
// It only reads data, doesn't modify.
func (p Person) Describe() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

// IsAdult checks if person is 18+
// Value receiver - we only read, don't modify
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// GetInitials returns the person's initials
// Value receiver - reading operation
func (p Person) GetInitials() string {
	if len(p.Name) == 0 {
		return ""
	}
	return string(p.Name[0]) + "."
}

// BasicValueReceiver shows how value receivers work
func BasicValueReceiver() {
	fmt.Println("=== Basic Value Receiver ===")

	person := Person{Name: "Alice", Age: 30}
	fmt.Printf("Person: %+v\n", person)

	// Call method with value receiver
	description := person.Describe()
	fmt.Printf("Description: %s\n", description)

	// The original person is unchanged (as expected)
	fmt.Printf("Person after method call: %+v\n", person)
}

// ValueReceiverImmutability demonstrates that value receivers can't modify
func ValueReceiverImmutability() {
	fmt.Println("\n=== Value Receiver Immutability ===")

	person := Person{Name: "Bob", Age: 25}

	// If we try to modify in the method, it only modifies the COPY
	// The original is safe and unchanged
	fmt.Printf("Before: %+v\n", person)
	fmt.Printf("IsAdult: %v\n", person.IsAdult())
	fmt.Printf("After: %+v\n", person)

	fmt.Println("The original person is unchanged!")
}

// MultipleValueReceiverCalls shows calling multiple value receiver methods
func MultipleValueReceiverCalls() {
	fmt.Println("\n=== Multiple Value Receiver Calls ===")

	people := []Person{
		{Name: "Charlie", Age: 17},
		{Name: "Diana", Age: 30},
		{Name: "Eve", Age: 45},
	}

	for _, p := range people {
		fmt.Printf("%s - %s - Adult: %v - Initials: %s\n",
			p.Describe(),
			p.GetInitials(),
			p.IsAdult(),
		)
	}
}

// ValueReceiverCopyingBehavior shows the copying semantics
func ValueReceiverCopyingBehavior() {
	fmt.Println("\n=== Value Receiver Copying ===")

	original := Person{Name: "Frank", Age: 35}

	// Value receivers work on copies, so large structs are copied
	// This means:
	// 1. Performance cost for large structs (data gets copied)
	// 2. Immutability guarantee (original can't be modified)

	describe := original.Describe()
	fmt.Printf("Describe called\n")
	fmt.Printf("Original: %+v\n", original)
	fmt.Printf("Description: %s\n", describe)
	fmt.Println("Original is unchanged - method worked on a copy")
}

// Account is another example struct
type Account struct {
	AccountID int
	Balance   float64
	Owner     string
}

// GetBalance returns the current balance
// Value receiver - read-only operation
func (a Account) GetBalance() float64 {
	return a.Balance
}

// IsOverdrawn checks if account is in negative
// Value receiver - read-only
func (a Account) IsOverdrawn() bool {
	return a.Balance < 0
}

// FormatBalance returns formatted balance string
// Value receiver - read-only
func (a Account) FormatBalance() string {
	return fmt.Sprintf("$%.2f", a.Balance)
}

// GetAccountInfo returns account information
// Value receiver - read-only
func (a Account) GetAccountInfo() string {
	return fmt.Sprintf("Account %d (%s): %s", a.AccountID, a.Owner, a.FormatBalance())
}

func RealWorldExample() {
	fmt.Println("\n=== Real World: Account Example ===")

	account := Account{
		AccountID: 12345,
		Balance:   1500.50,
		Owner:     "Grace",
	}

	fmt.Printf("Account Info: %s\n", account.GetAccountInfo())
	fmt.Printf("Balance: %s\n", account.FormatBalance())
	fmt.Printf("Overdrawn: %v\n", account.IsOverdrawn())

	// All these methods are read-only, so value receivers make sense
	// They guarantee the account data won't be accidentally modified
}

// WhenToUseValueReceivers explains the decision logic
func WhenToUseValueReceivers() {
	fmt.Println("\n=== When to Use Value Receivers ===")
	fmt.Println(`
Use value receivers when:

1. Method doesn't modify the receiver
   - Reading or calculating based on data
   - Transforming or formatting data
   - Querying state

2. Want immutability guarantee
   - Ensures receiver can't be modified
   - Good for concurrent code
   - Clear intent: "this doesn't change anything"

3. Receiver is small
   - Cost of copying is negligible
   - String, int, small struct: fine
   - But large struct: consider pointer receiver instead

Advantages of value receivers:
- Immutability guarantee
- Safe for concurrent use
- Clear API contract

Disadvantages:
- Copying cost for large structs
- Can't modify original
`)
}

// ComparisonWithPointerReceivers shows the difference
func ComparisonWithPointerReceivers() {
	fmt.Println("\n=== Comparison: Value vs Pointer ===")

	p := Person{Name: "Henry", Age: 28}

	// Value receiver method - works on copy
	fmt.Printf("Value receiver result: %s\n", p.Describe())
	fmt.Printf("After value receiver method: %+v\n", p)
	fmt.Println("Original unchanged")

	fmt.Println("\nWith pointer receiver methods (will show in next file)")
}

func main() {
	BasicValueReceiver()
	ValueReceiverImmutability()
	MultipleValueReceiverCalls()
	ValueReceiverCopyingBehavior()
	RealWorldExample()
	WhenToUseValueReceivers()
	ComparisonWithPointerReceivers()

	fmt.Println("\n=== Value Receiver Examples Complete ===")
}
