# Phase 2.4: Packages & Standard Library

**Duration:** 2-3 days  
**Prerequisites:** Phase 2.1-2.3 (Structs, Methods, Interfaces)  
**Focus:** Package organization, standard library overview, and practical usage

## 🎯 Learning Objectives

After completing this phase, you'll be able to:

- ✅ Understand Go package structure and organization
- ✅ Know the difference between exported and unexported identifiers
- ✅ Use critical standard library packages effectively
- ✅ Import and organize packages correctly
- ✅ Work with fmt for formatting and printing
- ✅ Manipulate strings with the strings package
- ✅ Convert between types with strconv
- ✅ Handle time and dates with the time package
- ✅ Organize your own packages

## 📚 Topics Covered

### 1. Package Fundamentals (1 day)
**Concept Files:** Standard library overview

**Key Concepts:**
- Package declaration and imports
- Exported vs. unexported identifiers (capitalization)
- Package-level variables and functions
- Blank imports and side effects
- Package naming conventions
- The "main" package

**Key Rules:**
```go
package main  // Package name must match directory

import (      // Import other packages
    "fmt"
    "strings"
)

// Exported (starts with capital)
func HelperFunction() {}
var PublicVariable = 10

// Unexported (starts with lowercase)
func helperFunction() {}
var privateVariable = 10
```

### 2. FMT Package (1 day)
**File:** `fmt-strings.go`

Formatting, printing, and scanning.

**Key Functions:**
- `Print()`, `Println()`, `Printf()`
- `Sprintf()` - format to string
- `Errorf()` - format error
- `Scan()`, `Scanf()`, `Scanln()` - reading input

**Common Format Verbs:**
- `%v` - value in default format
- `%T` - type of value
- `%s` - string
- `%d` - decimal integer
- `%x` - hex integer
- `%f` - floating point
- `%t` - boolean

**Example:**
```go
name := "Alice"
age := 30
fmt.Printf("%s is %d years old\n", name, age)
fmt.Sprintf("%d + %d = %d", 2, 3, 5)
```

### 3. Strings Package (1 day)
**File:** `fmt-strings.go`

String manipulation operations.

**Key Functions:**
- `ToUpper()`, `ToLower()`, `Title()`
- `Contains()`, `Index()`, `LastIndex()`
- `Split()`, `Join()`, `Fields()`
- `Replace()`, `ReplaceAll()`
- `TrimSpace()`, `Trim()`, `TrimPrefix()`, `TrimSuffix()`
- `HasPrefix()`, `HasSuffix()`
- `Repeat()`

**Example:**
```go
str := "hello world"
strings.ToUpper(str)           // "HELLO WORLD"
strings.Contains(str, "world")  // true
strings.Split(str, " ")        // ["hello" "world"]
strings.Replace(str, "o", "0", -1) // "hell0 w0rld"
```

### 4. Strconv Package (1 day)
**File:** `strconv-conversions.go`

Type conversions between strings and other types.

**Key Functions:**
- `ParseInt()`, `ParseFloat()`, `ParseBool()`
- `FormatInt()`, `FormatFloat()`, `FormatBool()`
- `Itoa()` - integer to string
- `Atoi()` - string to integer
- `Quote()`, `Unquote()` - string quoting

**Example:**
```go
num, err := strconv.ParseInt("123", 10, 64)
str := strconv.Itoa(456)
float, err := strconv.ParseFloat("3.14", 64)
bool, err := strconv.ParseBool("true")
```

### 5. Time Package (1 day)
**File:** `time-duration.go`

Time and duration handling.

**Key Types:**
- `time.Time` - point in time
- `time.Duration` - elapsed time

**Key Functions:**
- `Now()` - current time
- `Parse()`, `Format()` - string conversions
- `Sleep()` - pause execution
- `Unix()`, `UnixNano()` - Unix timestamps
- `Add()`, `Sub()` - time arithmetic
- `Before()`, `After()`, `Equal()` - comparisons

**Example:**
```go
now := time.Now()
future := now.Add(24 * time.Hour)
duration := time.Second * 5
time.Sleep(duration)
formatted := now.Format("2006-01-02 15:04:05")
```

## 📁 Directory Structure

```
phase-2-oop-functional/04-packages/
├── fmt-strings.go       # fmt and strings packages
├── strconv-conversions.go  # Type conversions
├── time-duration.go     # Time and duration handling
├── io-operations.go     # File I/O (io, os, bufio)
├── os-interaction.go    # OS operations
├── examples_test.go     # Tests
└── README.md           # This file
```

## 🚀 How to Use This Phase

### Step 1: Understand Packages
Read about Go package structure and conventions.

### Step 2: Explore fmt and strings
```bash
go run fmt-strings.go
```

### Step 3: Learn strconv
```bash
go run strconv-conversions.go
```

### Step 4: Master time
```bash
go run time-duration.go
```

### Step 5: Run Tests
```bash
go test -v
```

## 🔑 Key Concepts

### Package Organization
```go
// Package declaration
package main

// Imports
import (
    "fmt"       // Single import
    "strings"   // Another import
)

// Exported (visible outside package)
func PublicFunction() {}
var PublicVariable = 10

// Unexported (private to package)
func privateFunction() {}
var privateVariable = 10
```

### Formatting
```go
fmt.Printf("%v %T\n", 42, 42)      // General formatting
fmt.Printf("%d %x %b\n", 255, 255, 255) // Different bases
fmt.Printf("%.2f\n", 3.14159)      // Precision
fmt.Printf("|%10s|\n", "hello")    // Width and alignment
```

### Type Conversions
```go
// String to Integer
num, err := strconv.Atoi("42")
num64, err := strconv.ParseInt("42", 10, 64)

// Integer to String
str := strconv.Itoa(42)
str := strconv.FormatInt(255, 16)

// String to Float
float, err := strconv.ParseFloat("3.14", 64)

// Error handling is important!
if err != nil {
    log.Fatal(err)
}
```

### Time Operations
```go
now := time.Now()
formatted := now.Format("2006-01-02 15:04:05")
parsed, _ := time.Parse("2006-01-02", "2024-07-10")
future := now.Add(24 * time.Hour)
duration := future.Sub(now)
```

## 📋 Learning Path

### Day 1: Package Fundamentals & fmt
- [ ] Read about Go packages
- [ ] Understand exported vs. unexported
- [ ] Read and run `fmt-strings.go` (fmt section)
- [ ] Experiment with different format verbs
- [ ] Study the fmt examples

### Day 2: Strings & Strconv
- [ ] Read and run `fmt-strings.go` (strings section)
- [ ] Practice string operations
- [ ] Read and run `strconv-conversions.go`
- [ ] Understand error handling in conversions
- [ ] Practice type conversions

### Day 3: Time & Integration
- [ ] Read and run `time-duration.go`
- [ ] Understand time formatting (reference time!)
- [ ] Practice time arithmetic
- [ ] Run all tests
- [ ] Combine packages in small programs

## 💡 Tips

1. **Capitalization matters** - Determines export
2. **Reference time is 2006-01-02 15:04:05** - Use for time formatting
3. **Always check errors** - Parsing can fail
4. **Use constants** - `time.Hour`, `time.Second`, etc.
5. **Read godoc** - `go doc package/function`
6. **fmt has many verbs** - Learn the common ones
7. **Strings are immutable** - Operations return new strings

## ✅ Completion Checklist

- [ ] Understand package structure
- [ ] Know exported vs. unexported rules
- [ ] Can use fmt for printing
- [ ] Can format strings with fmt.Sprintf
- [ ] Can manipulate strings
- [ ] Can parse command line arguments
- [ ] Can convert types with strconv
- [ ] Can handle time.Time values
- [ ] Can format and parse dates
- [ ] Can work with durations
- [ ] Run all tests
- [ ] Write custom packages
- [ ] Read godoc for packages
- [ ] Use multiple packages together

## 🎯 Practice Exercises

### Exercise 1: String Formatter
```go
// Create a function that formats user data
type User struct {
    Name string
    Age int
}
// Use strings and fmt to display nicely
```

### Exercise 2: Type Converter
```go
// Create a program that converts between types
// Read from command line
// Use strconv for conversions
// Handle errors gracefully
```

### Exercise 3: Time Calculator
```go
// Create a function that calculates days until birthday
// Use time.Now() and time.Parse
// Return formatted string
```

## 🔗 Related Resources

- [fmt Package](https://golang.org/pkg/fmt/)
- [strings Package](https://golang.org/pkg/strings/)
- [strconv Package](https://golang.org/pkg/strconv/)
- [time Package](https://golang.org/pkg/time/)
- [Effective Go - Package Names](https://golang.org/doc/effective_go#package_names)

---

**Ready to start?** Run the example files and explore the standard library! 🚀
