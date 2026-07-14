# Phase 2: Object-Oriented & Functional Concepts in Go

Welcome to **Phase 2** of your Go journey! In this phase, we dive into how Go approaches structure, behavior, reuse, packaging, and dependency management. 

If you are coming from an Object-Oriented Programming (OOP) background in languages like **Java, C++, C#, or Python**, you will notice that Go does things differently. Go is not a traditional class-based OOP language. Instead, it uses **composition over inheritance**, **implicit interfaces (duck typing)**, and **packages** to achieve encapsulation, polymorphism, and code reuse.

This guide provides a deep explanation of these concepts, side-by-side comparisons with traditional OOP, and detailed breakdowns of each topic.

---

## 🎯 Learning Objectives

After completing this phase, you will understand:
1. **Structs**: How Go represents state, and how it differs from class-based state.
2. **Methods**: How Go attaches behavior to data, including value vs. pointer receivers.
3. **Interfaces**: Go's implicit polymorphism and why it enables loose coupling.
4. **Packages & Visibility**: How Go manages namespaces and encapsulation.
5. **Modules**: Go's modern dependency management system.

---

## 📚 Detailed Topic Explanations & OOP Comparisons

### 1. Structs: Defining State (No Classes!)

In traditional OOP languages like Java or Python, the fundamental building block is the **Class**. A class is a blueprint that combines both **state** (fields/properties) and **behavior** (methods) into a single entity.

Go does not have classes or objects. Instead, Go separates state and behavior:
* **Structs** define the state (data structure) only.
* **Methods** are declared separately and bound to a struct type.

#### ⚖️ Class vs. Struct Comparison

| Concept | Traditional OOP (Java/Python) | Go |
| :--- | :--- | :--- |
| **Blueprint** | `class Person { String name; int age; }` | `type Person struct { Name string; Age int }` |
| **Encapsulation** | `private`, `protected`, `public` keywords. | Capitalization rules (Uppercase = Public/Exported, Lowercase = Private/Unexported). |
| **Inheritance** | Class inheritance (`class Student extends Person`). | Struct Embedding / Composition (`type Student struct { Person; School string }`). |
| **Instantiation** | `Person p = new Person("Alice", 30);` | `p := Person{Name: "Alice", Age: 30}` |

#### 🔄 Composition (Embedding) vs. Inheritance

In Java, if `Student extends Person`, a `Student` object *is* a `Person` object (is-a relationship). You can assign a `Student` to a variable of type `Person`.

In Go, inheritance is replaced by **composition** via **Struct Embedding**:

```go
type Person struct {
    Name string
    Age  int
}

type Student struct {
    Person // Embedded struct (anonymous field)
    School string
}
```

* **Field Promotion**: Because `Person` is embedded without a field name, its fields (`Name` and `Age`) are **promoted** to `Student`. This means you can access them directly:
  ```go
  s := Student{Person: Person{Name: "Bob", Age: 20}, School: "MIT"}
  fmt.Println(s.Name) // Promoted: Bob (shortcut for s.Person.Name)
  ```
* **No Subtyping**: Even though `Student` embeds `Person`, a `Student` is **not** a `Person`. You cannot pass a `Student` to a function that expects a `Person` parameter. This prevents inheritance hierarchies and their associated fragilities.

#### 🏷️ Struct Tags
Go structs can have metadata tags attached to their fields. These tags are represented as string literals and are used by libraries (like the standard `encoding/json` library or database drivers) to control serialization.

```go
type User struct {
    Username string `json:"username" db:"user_name"`
    Password string `json:"-"` // Ignored by JSON serializer
}
```
*Compare to*: Java annotations (`@JsonProperty("username")`) or Python decorators.

---

### 2. Methods: Attaching Behavior

In Go, methods are just functions with a special argument called a **receiver**. The receiver binds the function to a specific type, making it a method of that type.

```go
type Circle struct {
    Radius float64
}

// Receiver: (c Circle)
func (c Circle) Area() float64 {
    return 3.14159 * c.Radius * c.Radius
}
```

#### 📌 Value Receivers vs. Pointer Receivers

Go methods can have either a **value receiver** or a **pointer receiver**. Understanding the difference is crucial for writing efficient and correct Go code.

##### Value Receiver `(c Circle)`
* **How it works**: The method receives a **copy** of the struct.
* **Side effects**: Any changes made to the receiver inside the method do **not** affect the original struct.
* **When to use**:
  * For small, immutable structs (e.g., a `Point` or `Time` struct).
  * When you only need to read data and do not need to modify the state.

##### Pointer Receiver `(c *Circle)`
* **How it works**: The method receives a **pointer** (reference) to the struct.
* **Side effects**: Changes made to the receiver inside the method **will mutate** the original struct.
* **When to use**:
  * When the method needs to modify (mutate) the state of the receiver.
  * When the struct is large, to avoid the CPU/memory cost of copying the entire struct on every method call.
  * For consistency: if *any* method on a struct requires a pointer receiver, *all* methods should use pointer receivers.

##### ⚖️ Comparison with `this` / `self`

In Java or Python, the method receiver is implicit:
* Java: `this` refers to the current instance (always a reference).
* Python: `self` is explicitly passed as the first argument, representing the instance.

In Go:
* You choose the name of the receiver (e.g., `c` instead of `this` or `self`).
* You explicitly choose whether it behaves like a value copy (value receiver) or a reference (pointer receiver).

---

### 3. Interfaces: Polymorphism & Implicit Implementation

Interfaces are Go's mechanism for polymorphism. An interface defines a contract: a set of method signatures. Any concrete type that implements those methods satisfies the interface.

#### 🦆 Duck Typing (Implicit Implementation)

In languages like Java, a class must explicitly state that it implements an interface:
```java
// Java
public class ConsoleLogger implements Logger {
    public void log(String msg) { System.out.println(msg); }
}
```

In Go, implementation is **implicit**. There is no `implements` keyword. If a type defines the methods in an interface, it implements the interface automatically.

```go
// Go
type Logger interface {
    Log(msg string)
}

type ConsoleLogger struct{}

// Satisfies the Logger interface implicitly
func (cl ConsoleLogger) Log(msg string) {
    fmt.Println(msg)
}
```

* **The Duck Test**: "If it walks like a duck and quacks like a duck, it's a duck."
* **Why this is powerful**:
  * **Decoupling**: The consumer of a type can define the interface it needs, rather than the producer. If you import a library with a struct, you can write your own interface matching its methods and mock it in tests, even though the library author never defined that interface!
  * **Minimal Interfaces**: Go interfaces are usually very small (often 1 or 2 methods, e.g., `io.Reader` and `io.Writer`).

#### 📦 The Empty Interface (`interface{}` or `any`)

An interface with zero methods is implemented by all types. In Go, `interface{}` (aliased as `any` in Go 1.18+) can hold values of any type.
* **Comparison**: Equivalent to `Object` in Java, `object` in C#, or `any` in TypeScript.
* **Use case**: When you need to handle arbitrary values of unknown types (e.g., JSON parsing or generic logging).

#### 🛠️ Type Assertions and Type Switches

Since an interface variable hides the concrete type, you need a way to extract the concrete type or check if it implements another interface.

##### Type Assertion
```go
// Check if interface value 'i' contains a string
s, ok := i.(string)
if ok {
    fmt.Println("It's a string:", s)
}
```

##### Type Switch
```go
switch v := i.(type) {
case int:
    fmt.Printf("Integer: %d\n", v)
case string:
    fmt.Printf("String: %s\n", v)
default:
    fmt.Printf("Unknown type\n")
}
```

---

### 4. Packages & Encapsulation

Go organizes code into packages. A package is a collection of source files in the same directory that are compiled together.

#### 👁️ Visibility (The Capitalization Rule)

Go does not have access modifiers like `public`, `private`, or `protected`. Instead, visibility is determined purely by the **capitalization of the first letter** of the identifier (variable, function, struct, interface, field, etc.).

* **Exported (Public)**: Starts with an **uppercase** letter. Visible to code outside the package.
  ```go
  package mathutils
  
  const Pi = 3.14159 // Visible outside the package
  
  func Add(a, b int) int { return a + b } // Visible
  ```
* **Unexported (Private/Package-Private)**: Starts with a **lowercase** letter. Only visible within the same package.
  ```go
  package mathutils
  
  const secretKey = "12345" // Private to mathutils package
  
  func subtract(a, b int) int { return a - b } // Private
  ```

---

### 5. Go Modules & Dependency Management

A **Go Module** is a collection of Go packages stored in a file tree with a `go.mod` file at its root. The `go.mod` file defines the module's path (its import path prefix) and its dependency requirements.

#### ⚖️ Package Manager Comparison

| Feature | Go Modules | npm (Node.js) | pip (Python) | Maven (Java) |
| :--- | :--- | :--- | :--- | :--- |
| **Config File** | `go.mod` | `package.json` | `requirements.txt` / `pyproject.toml` | `pom.xml` |
| **Lock File** | `go.sum` | `package-lock.json` | `poetry.lock` / `Pipfile.lock` | N/A (local repo caching) |
| **Commands** | `go mod tidy`, `go get` | `npm install` | `pip install -r` | `mvn install` |

* **Deterministic Builds**: The `go.sum` file contains cryptographic hashes of the dependencies, ensuring that the exact same code is compiled by everyone, preventing security risks and dependency drift.

---

## 🗂️ Phase 2 Code Directory Structure

Each sub-directory inside `phase-2-oop-functional/` provides structured, runnable files that demonstrate these concepts.

* **[01-structs/](file:///home/sjarso/go-mastery/phase-2-oop-functional/01-structs)**: Learn how structs group data.
  * [basics.go](file:///home/sjarso/go-mastery/phase-2-oop-functional/01-structs/basics.go): Instantiation, pointers, zero values.
  * [embedding.go](file:///home/sjarso/go-mastery/phase-2-oop-functional/01-structs/embedding.go): Composition over inheritance.
  * [tags.go](file:///home/sjarso/go-mastery/phase-2-oop-functional/01-structs/tags.go): Serializing structs to JSON/XML using struct tags.
* **[02-methods/](file:///home/sjarso/go-mastery/phase-2-oop-functional/02-methods)**: Bind functions to structures.
  * [value-receivers.go](file:///home/sjarso/go-mastery/phase-2-oop-functional/02-methods/value-receivers.go): Read-only method copies.
  * [pointer-receivers.go](file:///home/sjarso/go-mastery/phase-2-oop-functional/02-methods/pointer-receivers.go): Mutator methods and reference passing.
  * [method-chaining.go](file:///home/sjarso/go-mastery/phase-2-oop-functional/02-methods/method-chaining.go): Creating fluent builder patterns in Go.
* **[03-interfaces/](file:///home/sjarso/go-mastery/phase-2-oop-functional/03-interfaces)**: Polymorphic interactions.
  * [basics.go](file:///home/sjarso/go-mastery/phase-2-oop-functional/03-interfaces/basics.go): Interface syntax and contracts.
  * [duck-typing.go](file:///home/sjarso/go-mastery/phase-2-oop-functional/03-interfaces/duck-typing.go): Implicit structural typing in action.
  * [empty-interface.go](file:///home/sjarso/go-mastery/phase-2-oop-functional/03-interfaces/empty-interface.go): Dynamic variables and handling any type.
  * [type-assertions.go](file:///home/sjarso/go-mastery/phase-2-oop-functional/03-interfaces/type-assertions.go): Safely extracting concrete structures.
* **[04-packages/](file:///home/sjarso/go-mastery/phase-2-oop-functional/04-packages)**: Namespacing and standard utilities.
  * [fmt-strings.go](file:///home/sjarso/go-mastery/phase-2-oop-functional/04-packages/fmt-strings.go): Formatted input/output and string manipulation.
  * [strconv-conversions.go](file:///home/sjarso/go-mastery/phase-2-oop-functional/04-packages/strconv-conversions.go): Safe parsing of string numbers and booleans.
  * [time-duration.go](file:///home/sjarso/go-mastery/phase-2-oop-functional/04-packages/time-duration.go): Working with durations, timestamps, and formatting.
* **[05-modules-dependency/](file:///home/sjarso/go-mastery/phase-2-oop-functional/05-modules-dependency)**: Working with the package ecosystem.
  * [examples.go](file:///home/sjarso/go-mastery/phase-2-oop-functional/05-modules-dependency/examples.go): Importing external modules and handling path dependencies.

---

## 🚀 How to Run the Code

Start by navigating into any topic folder, and run the main entry file or tests:

```bash
# Go to Structs folder
cd phase-2-oop-functional/01-structs

# Run the basic struct example
go run basics.go

# Run the test assertions for the exercises
go test -v
```

To run all tests across Phase 2:
```bash
cd phase-2-oop-functional
go test -v ./...
```

---

## 🎯 Practice Exercises

To solidify your learning, complete the exercises specified in each folder:

1. **Exercise 1 (Structs & Methods)**: Create a `BankAccount` struct with fields for `OwnerName` and `Balance`. Define a pointer receiver method `Deposit(amount float64)` and a value receiver method `GetSummary() string`.
2. **Exercise 2 (Interfaces)**: Define a `PaymentGateway` interface with a `Charge(amount float64) error` method. Implement the interface for `StripeGateway` and `PayPalGateway`. Create a function `ProcessCheckout(gateway PaymentGateway, cartTotal float64)` to test both.
3. **Exercise 3 (Packages)**: Create a package called `validator` that contains unexported functions for regex pattern matching, and an exported function `IsValidEmail(email string) bool`. Import and use it in your main package.
