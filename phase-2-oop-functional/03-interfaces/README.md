# Phase 2.3: Interfaces & Duck Typing

**Duration:** 2-3 days  
**Prerequisites:** Phase 2.1 (Structs) and Phase 2.2 (Methods)  
**Focus:** Interface design, implementation, and practical patterns

## 🎯 Learning Objectives

After completing this phase, you'll be able to:

- ✅ Define and implement interfaces correctly
- ✅ Understand Go's implicit interface implementation (duck typing)
- ✅ Use the empty interface `interface{}` appropriately
- ✅ Perform safe type assertions and type switches
- ✅ Work with standard library interfaces (io.Reader, io.Writer, etc.)
- ✅ Design interfaces that promote loose coupling
- ✅ Compose multiple interfaces together
- ✅ Write polymorphic code using interfaces

## 📚 Topics Covered

### 1. Basic Interfaces (1 day)
**File:** `basics.go`

Interfaces define contracts—method sets that types can satisfy implicitly.

**Key Concepts:**
- Interface declaration syntax
- Implicit implementation (no "implements" keyword)
- Method sets and satisfaction rules
- Polymorphism through interfaces
- Treating different types uniformly

**Example:**
```go
type Speaker interface {
    speak() string
}

type Dog struct{ Name string }
func (d Dog) speak() string { return d.Name + " says woof!" }

type Cat struct{ Name string }
func (c Cat) speak() string { return c.Name + " says meow!" }

// Both Dog and Cat automatically satisfy Speaker
```

### 2. Duck Typing (1 day)
**File:** `duck-typing.go`

"If it walks like a duck and quacks like a duck, then it is a duck."

**Key Concepts:**
- Implicit interface satisfaction
- No explicit "implements" declarations
- Decoupled design patterns
- Interface composition
- Types with extra methods

**Example:**
```go
// Define what you need, not what exists
type Reader interface {
    Read() string
}

// Any type with Read() satisfies Reader
type ConsoleReader struct{}
func (c ConsoleReader) Read() string { return "..." }

type FileReader struct{}
func (f FileReader) Read() string { return "..." }
```

### 3. Empty Interface (1 day)
**File:** `empty-interface.go`

The empty interface `interface{}` accepts any type since all types have zero methods.

**Key Concepts:**
- `interface{}` represents any type
- Type assertions to extract values
- Type switches for multiple types
- When to use (and avoid) empty interface
- Generic containers and functions

**Example:**
```go
func LogValue(label string, value interface{}) {
    fmt.Printf("%s: %v\n", label, value)
}

// Accepts anything!
LogValue("name", "Alice")
LogValue("age", 30)
```

### 4. Type Assertions & Type Switches (1 day)
**File:** `type-assertions.go`

Safely extract concrete values from interface types.

**Key Concepts:**
- Type assertion syntax: `value.(Type)`
- Comma-ok idiom: `value, ok := i.(string)`
- Type switches for multiple cases
- Panic-free checking patterns
- Assertion chains

**Example:**
```go
var i interface{} = "hello"

// Safe assertion with comma-ok
s, ok := i.(string)
if ok {
    fmt.Println("String:", s)
}

// Type switch (idiomatic)
switch v := i.(type) {
case string:
    fmt.Println("String:", v)
case int:
    fmt.Println("Int:", v)
default:
    fmt.Println("Unknown")
}
```

### 5. Standard Library Interfaces (1 day)
**File:** `io-interfaces.go`

Go's standard library is built on powerful interfaces.

**Key Concepts:**
- `io.Reader` - Read from any source
- `io.Writer` - Write to any destination
- `fmt.Stringer` - Custom string representation
- `io.Closer` - Resource cleanup
- Composed interfaces
- `io.Copy` - Universal copying
- Real-world file I/O patterns

**Key Interfaces:**
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

type Stringer interface {
    String() string
}
```

## 📁 Directory Structure

```
phase-2-oop-functional/03-interfaces/
├── basics.go              # Basic interface definition and usage
├── duck-typing.go         # Implicit implementation patterns
├── empty-interface.go     # interface{} and type assertions
├── type-assertions.go     # Type switches and assertions
├── io-interfaces.go       # Standard library interfaces
├── examples_test.go       # Comprehensive test suite
└── README.md             # This file
```

## 🚀 How to Use This Phase

### Step 1: Start with Basics
```bash
cd phase-2-oop-functional/03-interfaces
go run basics.go
```

### Step 2: Understand Duck Typing
```bash
go run duck-typing.go
```

### Step 3: Work with Empty Interface
```bash
go run empty-interface.go
```

### Step 4: Master Type Assertions
```bash
go run type-assertions.go
```

### Step 5: Explore Standard Library
```bash
go run io-interfaces.go
```

### Step 6: Run Tests
```bash
go test -v                  # Run all tests
go test -v -run TestBasic   # Run specific tests
go test -bench=.            # Run benchmarks
```

## 🔑 Key Concepts

### Interface Definition
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

An interface defines a method set. A type satisfies an interface if it implements all methods.

### Implicit Implementation
```go
// No "implements Reader" needed!
type MyReader struct{}
func (m MyReader) Read(p []byte) (int, error) { ... }

// MyReader automatically satisfies Reader
var r Reader = MyReader{}
```

### Polymorphism
```go
func ProcessData(r Reader) {
    // Works with ANY Reader: strings, files, network, etc.
}
```

### Type Assertion
```go
// Extract concrete value (safe)
if str, ok := value.(string); ok {
    fmt.Println("Got string:", str)
}

// Type switch (idiomatic)
switch v := value.(type) {
case string:
    fmt.Println("String")
case int:
    fmt.Println("Int")
}
```

## 📋 Learning Path

### Day 1: Basics & Duck Typing
- [ ] Read `basics.go` completely
- [ ] Run it and understand polymorphism
- [ ] Read `duck-typing.go`
- [ ] See how types satisfy interfaces implicitly
- [ ] Study the examples closely

### Day 2: Empty Interface
- [ ] Read `empty-interface.go`
- [ ] Understand when to use `interface{}`
- [ ] Learn type assertion patterns
- [ ] Practice the comma-ok idiom

### Day 3: Type Assertions & Standard Library
- [ ] Read `type-assertions.go`
- [ ] Understand type switches
- [ ] Read `io-interfaces.go`
- [ ] Study Reader/Writer patterns
- [ ] Run all tests: `go test -v`

### Day 4: Mastery
- [ ] Run benchmarks: `go test -bench=.`
- [ ] Complete practice exercises
- [ ] Write custom interface implementations
- [ ] Experiment with composition

## 💡 Best Practices

### ✅ DO:
1. **Design small interfaces** - Single method preferred
2. **Depend on interfaces** - Not concrete types
3. **Return concrete types** - Accept interfaces
4. **Use type switches** - Better than multiple assertions
5. **Compose interfaces** - Combine small interfaces
6. **Check errors safely** - Use comma-ok idiom

### ❌ DON'T:
1. **Use `interface{}` for everything** - Only when necessary
2. **Create overly broad interfaces** - Keep them focused
3. **Ignore type assertion errors** - Always check
4. **Panic on bad assertions** - Use comma-ok instead
5. **Violate the interface contract** - Implement correctly

## 🎯 Practice Exercises

### Exercise 1: Implement Reader
```go
// Create a custom Reader that returns bytes from a source
// Implement io.Reader interface
// Test with io.Copy
```

### Exercise 2: Implement Writer
```go
// Create a custom Writer that collects bytes
// Implement io.Writer interface  
// Use with fmt.Fprintf
```

### Exercise 3: Type Switch Practice
```go
// Create a function that accepts interface{}
// Use type switch to handle: string, int, bool, []int
// Return appropriate string representation
```

### Exercise 4: Compose Interfaces
```go
// Create ReadWriter combining Reader and Writer
// Implement both io.Reader and io.Writer
// Test with bytes.Buffer
```

### Exercise 5: Custom Interface Design
```go
// Design your own interface for a logger
// Implement with Console, File, and Network writers
// Use in an application
```

## 📊 Comparison: Go vs Other Languages

| Aspect | Go | Java | Python |
|--------|----|----|--------|
| **Interface Declaration** | Implicit | Explicit | Duck typing |
| **Implementation** | Implicit | Explicit | Implicit |
| **Coupling** | Loose | Tight | N/A |
| **Multiple Interfaces** | Yes | Yes (interfaces) | Yes |
| **Type Assertions** | Safe | Casting | Dynamic |

## 🔗 Related Resources

- [Effective Go - Interfaces](https://golang.org/doc/effective_go#interfaces_and_other_types)
- [Go by Example - Interfaces](https://gobyexample.com/interfaces)
- [io Package Documentation](https://golang.org/pkg/io/)
- [Go Spec - Interface types](https://golang.org/ref/spec#Interface_types)
- [Go Spec - Type assertions](https://golang.org/ref/spec#Type_assertions)

## ✅ Completion Checklist

- [ ] Read and run `basics.go`
- [ ] Understand implicit implementation
- [ ] Read and run `duck-typing.go`
- [ ] Understand loose coupling through interfaces
- [ ] Read and run `empty-interface.go`
- [ ] Understand when to use `interface{}`
- [ ] Read and run `type-assertions.go`
- [ ] Master type switches
- [ ] Read and run `io-interfaces.go`
- [ ] Understand standard library patterns
- [ ] Run all tests: `go test -v`
- [ ] Run benchmarks: `go test -bench=.`
- [ ] Complete all practice exercises
- [ ] Write custom implementations
- [ ] Understand interface composition
- [ ] Know common interfaces (Reader, Writer, Stringer, Closer)

## 🎬 What's Next?

Once Phase 2.3 is complete:
1. You understand interfaces deeply
2. You can design interfaces for your code
3. You're comfortable with duck typing
4. You can work with standard library interfaces
5. You can write polymorphic code

**Move on to Phase 2.4:** Packages & Standard Library

Complete Phase 2 includes:
- Phase 2.1: Structs ✅
- Phase 2.2: Methods ✅
- Phase 2.3: Interfaces ✅ (YOU ARE HERE)
- Phase 2.4: Packages & Standard Library
- Phase 2.5: Modules & Dependency Management

---

## 💡 Quick Reference

### Defining an Interface
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

### Implementing an Interface
```go
type MyReader struct{}
func (m MyReader) Read(p []byte) (int, error) {
    // Implementation
    return len(p), nil
}
```

### Using an Interface
```go
var r Reader = MyReader{}
data := make([]byte, 100)
n, err := r.Read(data)
```

### Type Assertion
```go
value := interface{}("hello")
str, ok := value.(string)
```

### Type Switch
```go
switch v := value.(type) {
case string:
    fmt.Println("String:", v)
case int:
    fmt.Println("Int:", v)
default:
    fmt.Println("Unknown:", v)
}
```

---

**Ready to start?** Run `go run basics.go` and let's master interfaces! 🚀
