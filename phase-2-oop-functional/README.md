# Phase 2: Object-Oriented & Functional Concepts

**Duration:** 2-3 weeks  
**Prerequisites:** Phase 1 (Go fundamentals)  
**Focus:** Structs, Methods, Interfaces, Packages, and Modules

## 🎯 Learning Objectives

After completing this phase, you'll be able to:

- ✅ Define and work with structs effectively
- ✅ Create methods with value and pointer receivers
- ✅ Design and implement interfaces
- ✅ Understand Go's approach to OOP (composition over inheritance)
- ✅ Work with the standard library packages
- ✅ Manage project dependencies with Go Modules

## 📚 Topics Covered

### 1. Structs (1-2 days)
**File:** `01-structs/`

Structs are the foundation of Go's data organization. Unlike traditional OOP languages with classes, Go uses structs combined with methods.

**Key Concepts:**
- Struct definition and field declaration
- Creating instances
- Accessing fields
- Struct embedding (Go's replacement for inheritance)
- Struct tags (metadata for serialization)
- Zero values and initialization

**Files:**
- `basics.go` - Basic struct definition and usage
- `embedding.go` - Struct composition patterns
- `tags.go` - Struct tags for JSON, XML, validation
- `examples_test.go` - Test examples

**Example:**
```go
type Person struct {
    Name string
    Age  int
    Email string
}

p := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}
```

### 2. Methods (1-2 days)
**File:** `02-methods/`

Methods are functions associated with a type. They're Go's way of adding behavior to data.

**Key Concepts:**
- Method declaration and receivers
- Value receivers vs. pointer receivers
- When to use each
- Pointer semantics
- Method sets
- Receiver patterns

**Files:**
- `value-receivers.go` - Methods with value receivers
- `pointer-receivers.go` - Methods with pointer receivers
- `method-chaining.go` - Building fluent interfaces
- `examples_test.go` - Usage examples

**Example:**
```go
func (p *Person) SetAge(age int) {
    p.Age = age
}

func (p Person) GetDescription() string {
    return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}
```

### 3. Interfaces (2-3 days)
**File:** `03-interfaces/`

Interfaces define a contract - "if it can do this, it's this type." This is Go's way of polymorphism and is implicit (no explicit implementation declaration).

**Key Concepts:**
- Interface definition
- Implicit implementation (duck typing)
- Empty interface `interface{}`
- Type assertions and type switches
- Common interfaces (Reader, Writer, Stringer)
- Interface composition
- Practical interface design

**Files:**
- `basics.go` - Basic interface definition
- `duck-typing.go` - Implicit implementation
- `empty-interface.go` - Working with any type
- `type-assertions.go` - Type checks and conversions
- `io-interfaces.go` - Standard library interfaces
- `examples_test.go` - Practical usage

**Example:**
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Any type that implements Read() automatically satisfies Reader
```

### 4. Packages & Standard Library (2-3 days)
**File:** `04-packages/`

Packages are the way Go organizes code. Understanding the standard library is crucial.

**Key Concepts:**
- Package organization
- Exported vs. unexported identifiers
- Creating custom packages
- Import statements
- Package-level variables and functions
- Standard library overview

**Covered Packages:**
- `fmt` - Formatted I/O
- `strings` - String operations
- `strconv` - String conversions
- `io` and `ioutil` - I/O operations
- `os` - OS interaction
- `time` - Time handling
- `math` - Math functions
- `encoding` - Encoding/decoding

**Files:**
- `standard-library-overview.md` - Documentation
- `fmt-strings.go` - fmt and strings packages
- `strconv-conversions.go` - Type conversions
- `io-operations.go` - I/O operations
- `os-interaction.go` - OS operations
- `time-duration.go` - Time handling
- `examples_test.go` - Usage examples

**Example:**
```go
import (
    "fmt"
    "strings"
)

name := "Alice"
upper := strings.ToUpper(name)  // "ALICE"
fmt.Println(upper)               // Formatted output
```

### 5. Modules & Dependency Management (1-2 days)
**File:** `05-modules-dependency/`

Go Modules manage project dependencies and versioning.

**Key Concepts:**
- `go.mod` and `go.sum` files
- Initializing modules
- Adding dependencies
- Semantic versioning
- Updating packages
- Vendoring
- Go module commands

**Files:**
- `dependency-management.md` - Complete guide
- `examples.go` - Examples using external packages
- `go.mod` - Example module file

**Commands:**
```bash
go mod init github.com/username/project  # Initialize module
go get github.com/lib/pq                 # Add dependency
go mod tidy                              # Clean up dependencies
go mod vendor                            # Create vendor directory
```

## 🗂️ Directory Structure

```
phase-2-oop-functional/
├── README.md                          # This file
├── 01-structs/
│   ├── basics.go
│   ├── embedding.go
│   ├── tags.go
│   └── examples_test.go
│
├── 02-methods/
│   ├── value-receivers.go
│   ├── pointer-receivers.go
│   ├── method-chaining.go
│   └── examples_test.go
│
├── 03-interfaces/
│   ├── basics.go
│   ├── duck-typing.go
│   ├── empty-interface.go
│   ├── type-assertions.go
│   ├── io-interfaces.go
│   └── examples_test.go
│
├── 04-packages/
│   ├── standard-library-overview.md
│   ├── fmt-strings.go
│   ├── strconv-conversions.go
│   ├── io-operations.go
│   ├── os-interaction.go
│   ├── time-duration.go
│   └── examples_test.go
│
└── 05-modules-dependency/
    ├── dependency-management.md
    ├── go.mod
    └── examples.go
```

## 🚀 How to Use This Phase

### 1. **Start with Structs**
```bash
cd phase-2-oop-functional/01-structs
go run basics.go
```

### 2. **Read the Comments**
Each file is heavily commented. Read through them to understand concepts.

### 3. **Study the Tests**
Tests show practical usage. Read `*_test.go` files to see how code is used.

### 4. **Run Tests**
```bash
go test -v ./...
```

### 5. **Experiment**
Modify examples and run them to deepen understanding.

### 6. **Progress Sequentially**
- Day 1-2: Structs
- Day 3-4: Methods
- Day 5-7: Interfaces
- Day 8-9: Packages
- Day 10-11: Modules

## 📖 Key Concepts Summary

### Structs
- Group related data together
- Can be embedded in other structs (composition)
- Tags provide metadata for serialization

### Methods
- Functions with receivers
- Value receiver: works on copy
- Pointer receiver: works on original
- Choose receiver type based on whether you need to modify

### Interfaces
- Define method sets (contracts)
- Implicit implementation (duck typing)
- Enable polymorphism
- Use empty interface sparingly

### Packages
- Organize code logically
- Exported identifiers start with capital letter
- Standard library provides rich functionality

### Modules
- Manage dependencies
- Semantic versioning
- Reproducible builds

## 💡 Learning Tips

1. **Read Comments First** - Each file explains concepts
2. **Run Code** - Don't just read, execute
3. **Modify Examples** - Change values and see results
4. **Study Tests** - Tests show proper usage
5. **Build Small Programs** - Combine concepts in mini-projects
6. **Type Along** - Don't copy-paste, type code yourself
7. **Ask Questions** - Go community is helpful

## 🎯 Practice Exercises

### Exercise 1: Person Struct
```go
// Define a Person struct with Name, Age, Email
// Create a method Birthday() that increments age
// Create a method Describe() that returns formatted string
// Test your implementation
```

### Exercise 2: Interface Implementation
```go
// Define a Shape interface with Area() and Perimeter() methods
// Implement it with Rectangle and Circle types
// Create a function that works with any Shape
```

### Exercise 3: Package Creation
```go
// Create a calculator package
// Implement Add, Subtract, Multiply, Divide functions
// Use it from a main program
```

## 🔗 Related Resources

- [Structs Documentation](https://golang.org/doc/effective_go#embedding)
- [Methods and Receivers](https://golang.org/doc/effective_go#methods)
- [Interfaces](https://golang.org/doc/effective_go#interfaces_and_other_types)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

## ✅ Completion Checklist

- [ ] Read and understand struct basics
- [ ] Create and use structs in code
- [ ] Understand value vs. pointer receivers
- [ ] Implement multiple interfaces
- [ ] Work with interface{}
- [ ] Use type assertions
- [ ] Explore standard library packages
- [ ] Initialize a Go module
- [ ] Add and manage dependencies
- [ ] Complete practice exercises
- [ ] Write your own examples

## 🎓 What's Next?

Once you've completed Phase 2:
1. Make sure you understand structs and methods deeply
2. Be comfortable with interfaces and duck typing
3. Know how to organize code in packages
4. Understand dependency management

**Move on to Phase 3:** Advanced Language Features (Concurrency, Reflection, Generics)

---

**Ready to start?** Begin with `01-structs/basics.go` and work through each topic!
