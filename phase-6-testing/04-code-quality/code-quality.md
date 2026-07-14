# Code Quality & Static Analysis in Go

Writing tests is only half the battle. To ensure a codebase is maintainable, readable, and free of common traps, Go provides standard toolsets for formatting, linting, and static analysis.

Go has a strong culture of **convention over configuration**. Instead of debating style choices, the ecosystem conforms to standardized styling and patterns enforced directly by the Go CLI.

---

## 1. Code Formatting (`gofmt` & `go fmt`)

In other languages, team members often debate style rules (e.g., tabs vs. spaces, brace placement, line lengths). Go eliminates these discussions entirely via `gofmt`.

* **`gofmt`**: The command-line tool that parses Go source files and formats them according to a single, standard specification.
* **`go fmt`**: A wrapper command that runs `gofmt` on all packages in the current project.

### ⚙️ Command Usage
```bash
# Print formatting diffs for the current project without writing changes
gofmt -d .

# Format and write changes back to files in-place
go fmt ./...
```
*Note: Almost all Go IDEs (VS Code, GoLand) run `gofmt` automatically upon saving a file.*

---

## 2. Static Analysis (`go vet`)

The `go vet` command is a compiler-assistant static analysis tool that scans your codebase for suspicious constructs that compile successfully but represent likely runtime bugs.

### 🔍 Issues Checked by `go vet`
* **Mismatched Printf Verb Formats**: E.g., using `%d` to format a string variable:
  ```go
  fmt.Printf("User: %d", "Alice") // flagged by go vet
  ```
* **Unreachable Code**: E.g., code blocks positioned directly after a return statement.
* **Shallow Copy of Locks**: E.g., copying a struct containing a `sync.Mutex` by value (which copies the lock state, leading to deadlocks).
* **Incorrect Struct Tag Syntaxes**: E.g., typos in JSON tags like `json: "name"`.

### ⚙️ Command Usage
```bash
# Vet all packages in the current module
go vet ./...
```

---

## 3. Linting (`staticcheck`)

While `go vet` targets definitive bugs, linters focus on style, performance optimizations, and code simplifications. 

Historically, developers used `golint`. Today, **`staticcheck`** is the official, industry-standard linter for Go.

### 🔍 Issues Checked by `staticcheck`
* **Unused Codes**: Unused constants, fields, or private functions.
* **Simplification Suggestions**: E.g., replacing `select {}` loops or converting complex slice expressions to simple built-ins.
* **Performance Enhancements**: E.g., flagging instances where a string allocation is made unnecessarily inside a tight loop.

### ⚙️ Installation & Usage
```bash
# Install staticcheck
go install honnef.co/go/tools/cmd/staticcheck@latest

# Run on all packages
staticcheck ./...
```

---

## 4. Idiomatic Go Coding Patterns

When reviewing Go code, look out for the following idiomatic patterns:

### 1. Line of Sight (Shallow Nesting)
Keep the "happy path" aligned to the left of the screen by handling errors and edge cases early, returning quickly.

```go
// ❌ UNIDIOMATIC (Deep nesting)
func Process(u *User) error {
    if u != nil {
        if u.Active {
            err := u.Save()
            if err == nil {
                return nil
            } else {
                return err
            }
        }
        return ErrInactive
    }
    return ErrNilUser
}

//  IDIOMATIC (Shallow nesting)
func Process(u *User) error {
    if u == nil {
        return ErrNilUser
    }
    if !u.Active {
        return ErrInactive
    }
    if err := u.Save(); err != nil {
        return fmt.Errorf("process failed: %w", err)
    }
    return nil
}
```

### 2. Return Errors Immediately
Do not swallow errors or delay their propagation. Check them as soon as they are returned.

### 3. Keep Interfaces Small
"The bigger the interface, the weaker the abstraction." Limit interfaces to 1-3 methods. Larger interfaces are difficult to implement and mock, indicating that the interface is trying to do too much.
