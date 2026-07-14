# Phase 3: Advanced Language Features

Welcome to **Phase 3** of your Go journey! In this phase, we tackle advanced features that give Go its reputation for high performance, ease of concurrent programming, and robust systems development. 

This phase focuses on:
1. **Concurrency**: Understanding goroutines, channels, synchronization, and context.
2. **Reflection**: Meta-programming, dynamic typing, and runtime inspection.
3. **Generics**: Writing reusable, type-safe data structures and algorithms.
4. **Advanced Error Handling**: Moving beyond simple errors to nested wrapping, assertions, and recovery.

---

## 🎯 Learning Objectives

By the end of this phase, you will understand:
* **The CSP Concurrency Model**: Why Go handles concurrency differently than OS-thread languages.
* **Goroutines & Channels**: How to build lock-free, concurrent, passing-message systems.
* **Low-level Sync**: When and how to use Mutexes, WaitGroups, and Once.
* **Reflection & Generics**: How to write generic code and when to use runtime inspection.
* **Resilient Error Control**: Implementing wrapped error chains, custom errors, and clean recovery from panics.

---

## 📚 Detailed Topic Explanations & Language Comparisons

### 1. Concurrency: Goroutines & Channels (The CSP Model)

Traditional languages like **Java, C++, and Python** traditionally use **OS-level threads** for concurrency:
* Each thread maps directly to an OS thread.
* Stack sizes are large (typically 1MB to 2MB).
* Context switching requires kernel transitions, which are slow and expensive.
* Thread communication is usually done by **sharing memory** (shared variables), protected by locks (mutexes). This is notoriously prone to deadlocks, race conditions, and synchronization bugs.

Go implements **Communicating Sequential Processes (CSP)**. The core philosophy is:
> *"Do not communicate by sharing memory; instead, share memory by communicating."*

#### ⚖️ Concurrency Comparison

| Concept | Traditional Threads (Java/C++) | Go Goroutines |
| :--- | :--- | :--- |
| **Model** | OS Threads (1:1 mapping). | M:N Multiplexing (user-space scheduler). |
| **Stack Size** | Fixed, large (~1MB - 2MB). | Dynamic, starts very small (~2KB) and grows/shrinks. |
| **Startup Cost** | Slow (requires system calls). | Fast (nanoseconds, user-space allocation). |
| **Communication** | Shared memory (volatile fields, concurrent collections). | Channels (type-safe pipes). |
| **Locks** | ReentrantLocks, synchronized blocks. | Mutexes (only when needed) or Channel-based sync. |

#### 🌀 Goroutines vs. Threads
Go's runtime schedules **Goroutines** onto a pool of OS threads. Because goroutines have dynamic stacks and context switches happen in user space (without kernel overhead), you can easily run **hundreds of thousands** of goroutines simultaneously on a single laptop without exhausting memory.

#### 📣 Channels: Unbuffered vs. Buffered
Channels are strongly-typed, thread-safe pipes used to send and receive values between goroutines.

* **Unbuffered Channels (`ch := make(chan int)`)**:
  * The sender blocks until the receiver reads the value.
  * The receiver blocks until the sender writes a value.
  * **Use case**: Guaranteed synchronization / handoffs.
* **Buffered Channels (`ch := make(chan int, 100)`)**:
  * The sender can write values without blocking, up to the buffer size.
  * The sender only blocks when the buffer is full.
  * The receiver only blocks when the buffer is empty.
  * **Use case**: Producer-consumer pipelines, handling bursty traffic.

#### 🧳 The Context Package (`context`)
In Go, network requests and background operations span multiple goroutines. The `context` package provides a standardized way to propagate:
* **Cancellation signals**: Tell goroutines to stop work early (e.g., if a user closes a browser connection).
* **Timeouts / Deadlines**: Automatically cancel operations that take too long.
* **Request-scoped values**: Pass data like Trace IDs or Auth tokens across boundaries.

---

### 2. Reflection: Runtime Meta-Programming

Reflection allows a program to inspect and manipulate its own types, variables, and methods at runtime.

#### ⚖️ Reflection Comparison
In dynamic languages like Python, reflection is natural because everything is inspected at runtime. In static languages like Java, reflection uses the `Class` and `Field` APIs. In Go, the `reflect` package provides this functionality.

Go reflection is built around two primary types:
1. `reflect.Type`: Represents the Go type (e.g., `string`, `int`, `main.User`).
2. `reflect.Value`: Represents the concrete value inside a variable.

```go
var x float64 = 3.4
t := reflect.TypeOf(x)  // t represents float64
v := reflect.ValueOf(x) // v represents 3.4
```

#### ⚠️ Rules of Reflection
Reflection is powerful but comes with major drawbacks:
1. **No Compile-time Safety**: Code that uses reflection can easily panic at runtime if you pass the wrong type.
2. **Performance Overhead**: Reflection bypasses compiler optimizations and requires dynamic allocations, making it significantly slower.
3. **Complexity**: Code using reflection is harder to read, maintain, and document.
*Rule of thumb*: Avoid reflection unless writing highly generalized libraries (e.g., JSON encoders, ORMs, dependency injection containers).

---

### 3. Generics: Type-Safe Reusability

Introduced in Go 1.18, Generics allow you to write functions, structs, and interfaces with **type parameters**.

#### ⚖️ Generics Comparison

* **Java Generics**: Implemented using **Type Erasure**. The compiler checks types, but then removes the type parameter and replaces it with `Object` at runtime. This causes runtime type casting under the hood.
* **C++ Templates**: Implemented using **Monomorphization**. The compiler duplicates the template code for every concrete type used. This yields peak performance but results in larger executable binaries (code bloat).
* **Go Generics**: Uses a hybrid approach (GC-shape stenciling and dictionary passing) that preserves concrete types at runtime while avoiding excessive code duplication.

#### 📦 Syntax and Type Constraints
Go generics use square brackets `[...]` to declare type parameters, and **constraints** to restrict what types can be passed.

```go
// T is a type parameter constrained by the 'comparable' interface
func Index[T comparable](slice []T, target T) int {
    for i, v := range slice {
        if v == target { // 'comparable' allows using the == operator
            return i
        }
    }
    return -1
}
```

* **`any`**: Constraint allowing any type (alias for `interface{}`).
* **`comparable`**: Built-in constraint matching any type that supports `==` and `!=`.
* **Custom Constraints**: Defined using interfaces with type approximations (e.g., `~int` matches `int` and any custom type defined as `type MyInt int`).

---

### 4. Advanced Error Handling

Go is unique in how it handles errors. It has no `try-catch-finally` block. Instead, Go treats **errors as values** returned from functions alongside the actual results.

#### ⚖️ Exceptions vs. Errors-as-Values

* **Java/Python Exceptions**: If an error occurs, an exception is thrown, breaking the normal control flow and bubbling up the call stack until caught. This can make code control paths difficult to trace and leads to silent failures if exceptions are caught blindly.
* **Go Errors**: Returning `(Value, error)` forces the developer to handle the error immediately or explicitly propagate it. This encourages robust, readable, and predictable control flow.

#### 🔗 Error Wrapping (Go 1.13+)
When propagating an error up the stack, it is helpful to add context. In Go, you do this by **wrapping** the error using `%w` inside `fmt.Errorf`:

```go
func ReadConfig() ([]byte, error) {
    data, err := os.ReadFile("config.json")
    if err != nil {
        // %w wraps the original error 'err'
        return nil, fmt.Errorf("failed to read config: %w", err)
    }
    return data, nil
}
```

#### 🔍 Querying Wrapped Errors

If you wrap an error, you can inspect the error chain using the `errors` package:

* **`errors.Is(err, targetErr)`**: Checks if any error in the chain matches a specific sentinel error (similar to `err == targetErr`).
  ```go
  if errors.Is(err, os.ErrNotExist) {
      fmt.Println("File does not exist")
  }
  ```
* **`errors.As(err, &targetStruct)`**: Checks if any error in the chain matches a specific concrete error type and extracts it (similar to casting).
  ```go
  var pathErr *os.PathError
  if errors.As(err, &pathErr) {
      fmt.Println("Failed path:", pathErr.Path)
  }
  ```

#### 🚨 Panic & Recover
* **Panic**: Go's mechanism for fatal runtime failures (like dividing by zero or nil pointer dereferencing). Panics should **never** be used for normal error reporting.
* **Recover**: Used within a `defer` block to stop a panic from crashing the program. Useful in web servers to prevent a single failing request from crashing the entire server.

---

## 🗂️ Phase 3 Code Directory Structure

* **[01-concurrency/](file:///home/sjarso/go-mastery/phase-3-advanced/01-concurrency)**: Concurrent systems.
  * [basics.go](file:///home/sjarso/go-mastery/phase-3-advanced/01-concurrency/basics.go): Goroutines, `sync.WaitGroup`, and runtime scheduling.
  * [channels.go](file:///home/sjarso/go-mastery/phase-3-advanced/01-concurrency/channels.go): Unbuffered/buffered channels, select statement, close.
  * [sync-primitives.go](file:///home/sjarso/go-mastery/phase-3-advanced/01-concurrency/sync-primitives.go): Mutex, RWMutex, Once, and Cond operations.
  * **[patterns/](file:///home/sjarso/go-mastery/phase-3-advanced/01-concurrency/patterns)**: Concurrency blueprints.
    * [worker_pool.go](file:///home/sjarso/go-mastery/phase-3-advanced/01-concurrency/patterns/worker_pool.go): Distributing tasks among worker routines.
    * [pipeline.go](file:///home/sjarso/go-mastery/phase-3-advanced/01-concurrency/patterns/pipeline.go): Sequential processing stages.
    * [context.go](file:///home/sjarso/go-mastery/phase-3-advanced/01-concurrency/patterns/context.go): Timeout, cancellation, metadata passing.
* **[02-reflection/](file:///home/sjarso/go-mastery/phase-3-advanced/02-reflection)**: Type introspection.
  * [basics.go](file:///home/sjarso/go-mastery/phase-3-advanced/02-reflection/basics.go): Extracting types, values, and parsing struct tags.
  * [manipulation.go](file:///home/sjarso/go-mastery/phase-3-advanced/02-reflection/manipulation.go): Dynamic struct modification and method invocations.
* **[03-generics/](file:///home/sjarso/go-mastery/phase-3-advanced/03-generics)**: Reusable structures.
  * [basics.go](file:///home/sjarso/go-mastery/phase-3-advanced/03-generics/basics.go): Type parameters, custom constraints, `comparable`.
  * [patterns.go](file:///home/sjarso/go-mastery/phase-3-advanced/03-generics/patterns.go): Generic stacks and generic slice helper routines.
* **[04-error-handling-advanced/](file:///home/sjarso/go-mastery/phase-3-advanced/04-error-handling-advanced)**: Resilient architectures.
  * [wrapping.go](file:///home/sjarso/go-mastery/phase-3-advanced/04-error-handling-advanced/wrapping.go): Error chains, wrapping, checking via `Is` and `As`.
  * [panic-recover.go](file:///home/sjarso/go-mastery/phase-3-advanced/04-error-handling-advanced/panic-recover.go): Catching runtime crashes safely.

---

## 🚀 How to Run the Code

To run example programs, navigate into a topic and run the go files directly:

```bash
cd phase-3-advanced/01-concurrency
go run basics.go
```

To run all tests in Phase 3, ensuring you check for data races:
```bash
cd phase-3-advanced
go test -v -race ./...
```
*(The `-race` flag compiles Go code with race detector tooling enabled to flag concurrent read/write access to shared memory.)*
