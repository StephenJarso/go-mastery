# Phase 6: Testing & Quality

Welcome to **Phase 6** of your Go journey! In this phase, we master the tools, practices, and philosophies that keep Go codebases reliable, performant, and clean.

In Go, testing and tooling are not afterthoughts or external plug-ins. They are first-class citizens built directly into the language syntax and the `go` command-line tool. You don't need to choose between testing libraries, runners, benchmark suites, or linters; Go provides them all out of the box.

In this guide, we cover:
1. **Unit Testing**: Writing idiomatic, robust test suites and table-driven tests.
2. **Benchmarking**: Measuring execution speed and memory allocations scientifically.
3. **Advanced Testing**: Manual mocking via interfaces, and testing HTTP handlers using `httptest`.
4. **Code Quality**: Linting, static analysis, formatting, and writing idiomatic Go.

---

## 🎯 Learning Objectives

By the end of this phase, you will understand:
* **The Go Test Philosophy**: Why Go prefers standard `if` statements over assertion libraries.
* **Table-Driven Tests**: How to write clean, maintainable, and dry test suites.
* **Micro-benchmarking**: Measuring nanoseconds per operation and memory allocations.
* **Test Isolation (Mocking)**: Writing test doubles without dynamic bytecode/reflection mocking libraries.
* **HTTP Testing**: Isolating handlers and testing network clients using `net/http/httptest`.

---

## 📚 Detailed Topic Explanations & Language Comparisons

### 1. Unit Testing: Built-in and Idiomatic

In Java, you use JUnit. In Python, you use Pytest or Unittest. In JavaScript, you use Jest. 

In Go, you use the built-in `testing` package and the `go test` command. 

#### ⚖️ Testing Framework Comparison

| Feature | Pytest / JUnit / Jest | Go `testing` |
| :--- | :--- | :--- |
| **Runner** | Needs third-party libraries/installers. | Built-in (`go test`). |
| **Assertions** | `assertEquals(a, b)`, `expect(x).toBe(y)`. | Standard Go `if` statements. |
| **Parameterization** | `@pytest.mark.parametrize`, `@ParameterizedTest`. | **Table-Driven Tests** (slice of structs in a loop). |
| **Subtests** | Nested describes/context blocks. | `t.Run(name, func)`. |

#### 🚫 No Assertion Libraries
Go does not have assert functions. Instead, you write normal Go control flow:
```go
// Go style
if result != expected {
    t.Errorf("Expected %d, got %d", expected, result)
}
```
**Why?**
1. **No New DSL**: You don't have to learn a custom domain-specific language for asserting. If you know Go, you know how to write tests.
2. **Clear Error Context**: You write custom error messages explaining *why* the test failed, which is far more helpful than `expected true, got false`.
3. **Control Flow**: Traditional asserts throw exceptions to stop execution. Go allows tests to continue reporting multiple failures (`t.Errorf`) or stop immediately (`t.Fatalf`) depending on context.

#### 📊 Table-Driven Tests
Go developers write tests using **table-driven tests**. This pattern groups test cases as a slice of anonymous structs, running each case in a loop using `t.Run()`:

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -2, -3, -5},
        {"zero case", 0, 0, 0},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            result := Add(tc.a, tc.b)
            if result != tc.expected {
                t.Errorf("Add(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
            }
        })
    }
}
```

---

### 2. Benchmarking: Built-in Performance Profiling

In other languages, benchmarking is done via external libraries (like JMH in Java) or custom wrapper loops. In Go, benchmarking is built into the `testing` package.

#### 📈 Benchmark Syntax
A benchmark function starts with `Benchmark...` and accepts `*testing.B`. The runner manages the loop variable `b.N` dynamically until it achieves a statistically stable result:

```go
func BenchmarkStringConcat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = "hello" + " " + "world"
    }
}
```

#### 🛠️ Core Benchmark Tools
* **`b.ResetTimer()`**: Stops the timer, resets the elapsed time, and restarts it. Useful for excluding expensive setup operations from benchmark results.
* **`b.ReportAllocs()`**: Automatically calculates memory allocations (`B/op` - bytes per operation) and allocs count (`allocs/op`).
* **Running benchmarks**:
  ```bash
  go test -bench=. -benchmem
  ```

---

### 3. Advanced Testing & Test Doubles (Mocking)

#### 🛡️ Idiomatic Mocking (Mocking without Libraries)
In Java, mocking frameworks like Mockito use dynamic bytecode manipulation to override class behaviors at runtime.

In Go, you mock behavior using **Interfaces** and **Manual Test Doubles**. If a function depends on an interface, you can pass a custom struct in tests that implements that interface:

```go
// Production interface
type EmailSender interface {
    SendEmail(to, body string) error
}

// Test Mock (Test Double)
type MockEmailSender struct {
    LastTo   string
    LastBody string
    MockErr  error
}

func (m *MockEmailSender) SendEmail(to, body string) error {
    m.LastTo = to
    m.LastBody = body
    return m.MockErr
}
```
This is fully type-safe, compile-time verified, and requires no reflection magic.

#### 🌐 Testing HTTP Handlers (`net/http/httptest`)
Go provides the `net/http/httptest` package to test web servers and clients without opening actual TCP network sockets.

##### 1. Testing Handlers with `ResponseRecorder`
`httptest.NewRecorder()` acts as a fake browser connection. It records response codes, headers, and body writes:
```go
req := httptest.NewRequest("GET", "/hello", nil)
w := httptest.NewRecorder()

HelloHandler(w, req) // Call HTTP handler directly

resp := w.Result()
if resp.StatusCode != http.StatusOK {
    t.Errorf("Expected 200, got %d", resp.StatusCode)
}
```

##### 2. Testing HTTP Clients with `httptest.NewServer`
`httptest.NewServer` spins up a real local test server on a random loopback port. You feed this test server's URL to your HTTP client config, allowing you to test real HTTP client exchanges (including connection handshakes and query parameters) in isolation.

---

### 4. Code Quality & Static Analysis

Go enforces code quality and style directly in the toolchain.

* **Formatting (`gofmt`)**: Go has a single standard code layout. Running `gofmt -w file.go` automatically formats your code, ending all style arguments (braces placement, spacing, etc.).
* **Static Analysis (`go vet`)**: Inspects code for suspicious structures that compilation might miss, such as mismatched printf tags or unreachable code.
* **Linting (`staticcheck`)**: The industry standard for linting Go code, catching performance issues, bugs, and stylistic violations.

---

## 🗂️ Phase 6 Code Directory Structure

* **[01-unit-testing/](file:///home/sjarso/go-mastery/phase-6-testing/01-unit-testing)**: Learn standard testing patterns.
  * [basics.go](file:///home/sjarso/go-mastery/phase-6-testing/01-unit-testing/basics.go) & [basics_test.go](file:///home/sjarso/go-mastery/phase-6-testing/01-unit-testing/basics_test.go): Basic test definitions and testing helpers.
  * [table_driven.go](file:///home/sjarso/go-mastery/phase-6-testing/01-unit-testing/table_driven.go) & [table_driven_test.go](file:///home/sjarso/go-mastery/phase-6-testing/01-unit-testing/table_driven_test.go): Creating parameter-driven tests and subtests.
  * [practice.go](file:///home/sjarso/go-mastery/phase-6-testing/01-unit-testing/practice.go) & [practice_test.go](file:///home/sjarso/go-mastery/phase-6-testing/01-unit-testing/practice_test.go): Exercises covering slice aggregations and validations.
* **[02-benchmarking/](file:///home/sjarso/go-mastery/phase-6-testing/02-benchmarking)**: Performance measurements.
  * [concat.go](file:///home/sjarso/go-mastery/phase-6-testing/02-benchmarking/concat.go) & [concat_test.go](file:///home/sjarso/go-mastery/phase-6-testing/02-benchmarking/concat_test.go): Benchmarking raw string concat vs `strings.Builder`.
  * [practice.go](file:///home/sjarso/go-mastery/phase-6-testing/02-benchmarking/practice.go) & [practice_test.go](file:///home/sjarso/go-mastery/phase-6-testing/02-benchmarking/practice_test.go): Optimizing slice filtering structures.
* **[03-advanced-testing/](file:///home/sjarso/go-mastery/phase-6-testing/03-advanced-testing)**: Mocking and network tests.
  * [mocking.go](file:///home/sjarso/go-mastery/phase-6-testing/03-advanced-testing/mocking.go) & [mocking_test.go](file:///home/sjarso/go-mastery/phase-6-testing/03-advanced-testing/mocking_test.go): Designing interface-based test doubles.
  * [http_handler.go](file:///home/sjarso/go-mastery/phase-6-testing/03-advanced-testing/http_handler.go) & [http_handler_test.go](file:///home/sjarso/go-mastery/phase-6-testing/03-advanced-testing/http_handler_test.go): Testing standard HTTP endpoints via `httptest.ResponseRecorder`.
  * [practice.go](file:///home/sjarso/go-mastery/phase-6-testing/03-advanced-testing/practice.go) & [practice_test.go](file:///home/sjarso/go-mastery/phase-6-testing/03-advanced-testing/practice_test.go): Mocking user storage databases and verifying API payloads.

---

## 🚀 How to Run the Code

To run unit tests in a specific folder:
```bash
cd phase-6-testing/01-unit-testing
go test -v
```

To run benchmarks in a folder:
```bash
cd phase-6-testing/02-benchmarking
go test -bench=. -benchmem
```

To run a test coverage check:
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```
