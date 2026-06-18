# Phase 2.2: Methods & Receivers

**Duration:** 2-3 days  
**Prerequisites:** Phase 2.1 (Structs)
**Focus:** Methods, Receivers, and Method Patterns

## 🎯 Learning Objectives

After completing this phase, you'll be able to:

- ✅ Define methods on struct types
- ✅ Understand value receivers vs. pointer receivers
- ✅ Know when to use each receiver type
- ✅ Design method sets for clean APIs
- ✅ Use method chaining for fluent interfaces
- ✅ Apply receiver patterns in real-world scenarios

## 📚 Topics Covered

### 1. Methods Basics (1 day)
**File:** `value-receivers.go`

Methods are functions with a receiver. They add behavior to types.

**Key Concepts:**
- Method declaration syntax
- Receiver types (value or pointer)
- Method sets
- Calling methods
- Method scope (exported vs unexported)

### 2. Value Receivers (1 day)
**File:** `value-receivers.go`

Value receivers work on a copy of the struct.

**Key Concepts:**
- When to use value receivers
- Immutability guarantees
- Performance considerations
- Reading-only operations

**Example:**
```go
func (p Person) Describe() string {
    return fmt.Sprintf("%s is %d", p.Name, p.Age)
}
```

### 3. Pointer Receivers (1 day)
**File:** `pointer-receivers.go`

Pointer receivers modify the original struct.

**Key Concepts:**
- When to use pointer receivers
- Modifying state
- Efficiency with large structs
- Consistency in method sets

**Example:**
```go
func (p *Person) HaveBirthday() {
    p.Age++
}
```

### 4. Method Chaining (1 day)
**File:** `method-chaining.go`

Return the receiver to enable method chaining (fluent interface).

**Key Concepts:**
- Builder pattern
- Fluent interfaces
- Readable API design
- Chaining syntax

**Example:**
```go
config.WithHost("localhost").
    WithPort(8080).
    WithTimeout(30).
    Build()
```

## 📁 Directory Structure

```
phase-2-oop-functional/02-methods/
├── value-receivers.go       # Methods with value receivers
├── pointer-receivers.go     # Methods with pointer receivers
├── method-chaining.go       # Builder pattern and fluent interfaces
├── examples_test.go         # Comprehensive tests
├── practice.go              # Practice exercises
├── practice_test.go         # Tests for practice exercises
├── go.mod                   # Module definition
└── README.md                # This file
```

## 🚀 How to Use This Phase

### Step 1: Study Methods Basics
```bash
cd phase-2-oop-functional/02-methods
go run value-receivers.go
```

### Step 2: Understand Pointer Receivers
```bash
go run pointer-receivers.go
```

### Step 3: Learn Method Chaining
```bash
go run method-chaining.go
```

### Step 4: Run Tests
```bash
go test -v
go test -bench=.
```

### Step 5: Practice
```bash
go run practice.go
```

## 🔑 Key Concepts

### Receiver Syntax
```go
// Method with value receiver (works on copy)
func (p Person) Method() {
    // p is a copy, modifications don't affect original
}

// Method with pointer receiver (works on original)
func (p *Person) Method() {
    // p points to original, modifications affect it
}
```

### Choosing Receiver Type

**Use Value Receiver when:**
- Method doesn't need to modify receiver
- Receiver is small and cheap to copy
- Want to guarantee immutability
- Implementing concurrent-safe code

**Use Pointer Receiver when:**
- Method needs to modify receiver
- Receiver is large (avoid copying overhead)
- Consistency: if ANY method uses pointer receiver, use it for all
- Implementing interfaces that need modification

### Method Sets

```go
// For type T:
// - Can call methods with T receiver
// - Can call methods with *T receiver (through auto-dereferencing)

// For type *T:
// - Can call methods with *T receiver
// - Can call methods with T receiver (through auto-dereferencing)
```

## 📋 Learning Path

### Day 1: Value Receivers
- [ ] Read `value-receivers.go` completely
- [ ] Run it and see the output
- [ ] Understand immutability concept
- [ ] Study test examples

### Day 2: Pointer Receivers
- [ ] Read `pointer-receivers.go`
- [ ] Run it and see the output
- [ ] Understand when modifications work
- [ ] Compare with value receivers

### Day 3: Method Chaining & Practice
- [ ] Read `method-chaining.go`
- [ ] Run it and see fluent interface
- [ ] Complete practice exercises
- [ ] Write your own method chains

### Day 4: Mastery
- [ ] Run all tests: `go test -v`
- [ ] Run benchmarks: `go test -bench=.`
- [ ] Complete all practice exercises
- [ ] Modify examples and experiment

## 💡 Tips

1. **Value receivers = immutable** - Perfect for read-only operations
2. **Pointer receivers = mutable** - Use when you need to change state
3. **Be consistent** - Use same receiver type for all methods on a type
4. **Auto-dereferencing** - Go automatically handles conversion between `T` and `*T`
5. **Method chaining** - Return receiver for fluent API design

## ✨ Quick Reference

```go
// Define a type
type Person struct {
    Name string
    Age  int
}

// Value receiver method (read-only)
func (p Person) Describe() string {
    return fmt.Sprintf("%s is %d", p.Name, p.Age)
}

// Pointer receiver method (can modify)
func (p *Person) Birthday() {
    p.Age++
}

// Using the methods
p := Person{Name: "Alice", Age: 30}
fmt.Println(p.Describe())  // "Alice is 30"
p.Birthday()               // Age becomes 31
fmt.Println(p.Describe())  // "Alice is 31"

// Method chaining
config := NewConfig().
    WithName("myapp").
    WithPort(8080).
    WithTimeout(30).
    Build()
```

## 🎯 Practice Exercises

### Exercise 1: Value Receiver
Create methods with value receivers that don't modify the struct.

### Exercise 2: Pointer Receiver
Create methods with pointer receivers that modify the struct.

### Exercise 3: Method Chaining
Implement a builder pattern with method chaining.

### Exercise 4: Mixed Receivers
Create a type with both value and pointer receiver methods. Understand which to use when.

## 🔗 Related Resources

- [Effective Go - Methods](https://golang.org/doc/effective_go#methods)
- [Go by Example - Methods](https://gobyexample.com/methods)
- [Method Sets](https://golang.org/ref/spec#Method_sets)
- [Receiver Types](https://golang.org/doc/effective_go#pointers_vs_values)

## ✅ Completion Checklist

- [ ] Read and run `value-receivers.go`
- [ ] Read and run `pointer-receivers.go`
- [ ] Understand when to use each
- [ ] Read and run `method-chaining.go`
- [ ] Run all tests: `go test -v`
- [ ] Complete practice exercises
- [ ] Write your own method examples
- [ ] Understand receiver consistency
- [ ] Know auto-dereferencing rules
- [ ] Master method chaining patterns

## 🎬 What's Next?

Once Phase 2.2 is complete:
1. Ensure you understand value vs pointer receivers deeply
2. Know when to use method chaining
3. Understand receiver consistency principle

**Move on to Phase 2.3:** Interfaces & Duck Typing

---

**Ready to start?** Run `go run value-receivers.go` and let's go! 🚀
