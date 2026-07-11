# Phase 3: Advanced Language Features

**Duration:** 3-4 weeks  
**Prerequisites:** Phase 2 (Object-Oriented & Functional Concepts)  
**Focus:** Concurrency, Reflection, Generics, and Advanced Error Handling

---

## 🎯 Learning Objectives

After completing this phase, you'll be able to:

- ✅ Write highly concurrent, thread-safe Go code using goroutines and channels
- ✅ Avoid race conditions and deadlocks using `sync` package primitives
- ✅ Apply idiomatic concurrency patterns (worker pools, pipelines, fan-in/fan-out)
- ✅ Use the `context` package to manage timeouts and cancellations
- ✅ Understand and use reflection when necessary, while knowing its drawbacks
- ✅ Write reusable generic data structures and functions using Go Generics
- ✅ Implement robust error handling strategies using error wrapping, unwrapping, type switches, and panics/recoveries safely

---

## 📚 Topics Covered

### 1. Concurrency (1-2 weeks)
**File:** `01-concurrency/`

Go's concurrency model is based on CSP (Communicating Sequential Processes) and is built directly into the language syntax and runtime.

**Key Concepts:**
- **Goroutines**: Lightweight, green threads managed by the Go runtime.
- **Channels**: Strongly-typed pipes for passing data and synchronizing execution between goroutines.
- **Sync Package**: Low-level synchronization primitives like Mutex, RWMutex, Once, and Cond.
- **Context**: Standard package for propagating cancellation, deadlines, and request-scoped values across API boundaries.
- **Patterns**: Multiplexing (`select`), Worker Pools, Fan-in/Fan-out, Pipelines.

**Files:**
- `01-concurrency/basics.go` - Launching goroutines and using `sync.WaitGroup`.
- `01-concurrency/channels.go` - Unbuffered, buffered channels, range, select, and close operations.
- `01-concurrency/sync-primitives.go` - Safe state sharing with `sync.Mutex`, `sync.RWMutex`, `sync.Once`, and `sync.Cond`.
- `01-concurrency/patterns/` - Advanced patterns (Worker Pools, Pipelines, Context propagation).
- `01-concurrency/examples_test.go` - Practical examples and test validation.

---

### 2. Reflection (2-3 days)
**File:** `02-reflection/`

Reflection allows a program to inspect and manipulate its own structure and behavior at runtime.

**Key Concepts:**
- `reflect.Type` vs. `reflect.Value`.
- Inspecting struct tags and fields.
- Dynamic value modification and method invocation.
- Performance implications and type-safety trade-offs.

**Files:**
- `02-reflection/basics.go` - Inspecting types, fields, and tags.
- `02-reflection/manipulation.go` - Modifying values dynamically and invoking methods.
- `02-reflection/examples_test.go` - Verification of dynamic operations.

---

### 3. Generics (3-4 days)
**File:** `03-generics/`

Generics (introduced in Go 1.18) allow writing code with type parameters, enabling type-safe code reuse without reflection.

**Key Concepts:**
- Type parameters, constraints, and bounds.
- Custom interfaces as constraints.
- Generic collections (e.g., Maps, Slices, Sets).
- Built-in `comparable` constraint and `golang.org/x/exp/constraints`.

**Files:**
- `03-generics/basics.go` - Declaring type parameters, generic functions, and basic constraints.
- `03-generics/patterns.go` - Reusable data structures (generic stack, generic slice operations).
- `03-generics/examples_test.go` - Test validation of generic constructs.

---

### 4. Advanced Error Handling (3-4 days)
**File:** `04-error-handling-advanced/`

Go treats errors as values. Go 1.13 introduced standard error wrapping utilities to make error inspections more robust.

**Key Concepts:**
- Error wrapping with `fmt.Errorf("... %w", err)`.
- Error unwrapping and querying with `errors.Is` and `errors.As`.
- Creating custom error types with rich contextual fields.
- Idiomatic panic and recover patterns (panic vs. returning errors).

**Files:**
- `04-error-handling-advanced/wrapping.go` - Wrapping, unwrapping, and type asserting errors.
- `04-error-handling-advanced/panic-recover.go` - Recovering from panics safely.
- `04-error-handling-advanced/examples_test.go` - Tests demonstrating robust error assertions.

---

## 🚀 How to Use This Phase

### 1. Run Examples
Navigate to the directory of a specific topic and run the go files directly:
```bash
cd phase-3-advanced/01-concurrency
go run basics.go
```

### 2. Run Tests with Race Detector
For concurrency, it is critical to always test with the race detector enabled:
```bash
go test -v -race ./...
```

---

## ✅ Completion Checklist

- [ ] Run and understand basic goroutines and wait groups
- [ ] Understand blocking vs non-blocking channel operations
- [ ] Implement a worker pool using channels and select
- [ ] Prevent data races using `sync.Mutex` and test with `-race`
- [ ] Use `context.WithTimeout` to cancel slow network/DB calls
- [ ] Write a generic function to filter/map slices
- [ ] Wrap errors and check them using `errors.Is` and `errors.As`
- [ ] Handle a panic gracefully using `recover` in a deferred function
