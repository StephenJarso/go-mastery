# Phase 5: Data & Persistence

Welcome to **Phase 5** of your Go journey! In this phase, we explore how Go applications store, serialize, and query data. 

In any real-world application, managing data efficiently is critical. Go does not dictate a single database engine or serialization standard. Instead, it provides clean, pluggable interfaces like `database/sql` and `encoding/json` that let you hook in SQLite, PostgreSQL, Redis, or binary serialization tools seamlessly.

In this guide, we cover:
1. **JSON Processing**: Encoding, decoding, custom formats, and stream processing.
2. **Database SQL Fundamentals**: Managing connection pools, running queries safely, and handling transactions.
3. **ORMs (GORM)**: Easing development with object-relation mapping, migrations, and associations.
4. **Data Serialization**: Writing XML, and comparing binary protocols like MessagePack and Protocol Buffers (Protobuf).

---

## 🎯 Learning Objectives

By the end of this phase, you will understand:
* **Dynamic vs. Static JSON**: Marshaling fixed structures vs. parsing arbitrary payloads.
* **Database Pool Safety**: Tuning SQL pools and managing safe transactions with deferred rollbacks.
* **ORMs vs. Raw SQL**: Selecting the right strategy and avoiding common issues like the N+1 query problem.
* **Binary Formats**: Why and when to choose Protobuf or MessagePack over JSON for microservices and caching.

---

## 📚 Detailed Topic Explanations & Language Comparisons

### 1. JSON Processing: Serialization & Deserialization

JSON is the lingua franca of web communication. 

In dynamic languages like Python, parsing JSON yields an instant, schema-less dictionary (`json.loads()`). In Java, libraries like Jackson use complex runtime reflection and annotations. In Go, the standard `encoding/json` package uses struct tags and reflection to map JSON fields to strongly-typed structs.

#### 👁️ Struct Tags for JSON Mapping
Because Go field visibility depends on capitalization, JSON keys (which are usually camelCase or snake_case) would conflict with Go conventions without tags:
```go
type User struct {
    FirstName string `json:"first_name"` // Maps JSON "first_name" to FirstName
    Age       int    `json:"age"`
    Password  string `json:"-"`          // Excludes from JSON output
}
```

#### 📌 Unmarshal vs. NewDecoder
Go offers two ways to read JSON:
* **`json.Unmarshal(data []byte, &v)`**:
  * Reads from a byte slice already in memory.
  * **When to use**: You already have the entire dataset loaded in memory (e.g., from a database query result or a local cache).
* **`json.NewDecoder(r io.Reader).Decode(&v)`**:
  * Streams JSON data directly from an input stream (`io.Reader`).
  * **When to use**: Reading from network sockets, files, or HTTP request bodies. It processes data as it arrives, avoiding loading the entire payload into RAM at once.

#### 🔄 Custom Marshaling
By implementing the `json.Marshaler` and `json.Unmarshaler` interfaces, you can override how specific types are read or written. This is commonly used to parse custom date-time strings (like `"YYYY-MM-DD"`) that don't match the RFC3339 default.

---

### 2. Database Fundamentals: SQLite & raw SQL

Go's database connectivity is built on the standard `database/sql` package. 

#### 🔌 Driver Abstraction (Compare to JDBC/DB-API)
In Java, you use JDBC; in Python, you use PEP-249 DB-API. 
In Go, `database/sql` is the general abstraction. To use a specific database (like SQLite, PostgreSQL, or MySQL), you import a driver anonymously:
```go
import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3" // Registers driver with database/sql
)
```
Your code interacts only with `*sql.DB` and `*sql.Rows`, making it easy to swap databases without rewriting core queries.

#### 🏊 Connection Pooling
Go manages the database connection pool automatically. You do not need to create or destroy physical connections manually. You should configure the pool limits to suit your system load:
* `SetMaxOpenConns(n)`: Limits active connections to protect database RAM.
* `SetMaxIdleConns(n)`: Keeps idle connections alive to avoid handshaking overhead on incoming requests.
* `SetConnMaxLifetime(d)`: Recycles connections to prevent stale connections.

#### 🛡️ Preventing SQL Injection
**Never** build queries using string formatting (`fmt.Sprintf`). Go uses parameterized placeholders (e.g., `?` in SQLite/MySQL, `$1`, `$2` in PostgreSQL):
```go
// SAFE: Parameterized query executed as a prepared statement by the driver
row := db.QueryRow("SELECT id FROM users WHERE email = ?", inputEmail)
```

#### 🔄 Transactions with Safe Rollbacks
A database transaction executes multiple operations atomically. In Go, you begin a transaction with `db.Begin()`. To guarantee that a transaction doesn't remain dangling under an error or panic, use the deferred rollback pattern:
```go
tx, err := db.Begin()
if err != nil {
    return err
}
defer tx.Rollback() // Safe: does nothing if tx.Commit() succeeds first

// ... Run database queries using tx ...

if err := tx.Commit(); err != nil {
    return err
}
```

---

### 3. ORMs (Object-Relational Mapping): GORM

GORM is the most popular ORM for Go.

#### ⚖️ ORM Comparison

| Feature | GORM (Go) | Hibernate (Java) | SQLAlchemy (Python) |
| :--- | :--- | :--- | :--- |
| **Model Mapping** | Structs and struct tags. | XML / Java Class Annotations. | Declarative classes. |
| **Caching** | Basic query caching. | Complex 1st/2nd level cache. | Session caching. |
| **Query Style** | Method chaining. | JPQL / Criteria API. | Expression-based. |

#### 🔄 Eager Loading & The N+1 Problem
If you query a list of 100 users and then query their profile details individually in a loop, your application executes 1 + 100 = 101 database queries. This is the **N+1 query problem**.

GORM solves this via eager loading using **`Preload`**:
```go
// Executes only 2 queries: one for users, and one for profiles using an IN clause.
var users []User
db.Preload("Profile").Find(&users)
```

#### ⚖️ When to use ORM vs. Raw SQL
* **GORM (ORM)**: Excellent for CRUD operations, simple table associations, and quick prototyping.
* **Raw SQL (`database/sql`)**: Better for performance-critical queries, bulk operations, analytical queries, and when you need complete control over indexing and execution plans.

---

### 4. Data Serialization: Beyond JSON

| Format | Output | Size | Speed | Schema Requirement | Use Case |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **JSON** | Text | Medium | Medium | No | Web APIs, browser communication. |
| **XML** | Text | Large | Slow | No | Enterprise systems, SOAP APIs. |
| **MessagePack**| Binary | Small | Fast | No | Dynamic key-value caches (e.g., Redis payload optimization). |
| **Protobuf** | Binary | Extremely Small | Extremely Fast | Yes (`.proto` file) | Microservice IPC (gRPC), high-throughput feeds. |

#### 🕸️ Protocol Buffers (Protobuf)
Developed by Google, Protobuf encodes structured data in a highly compact binary format.
* **Varints**: Integers are encoded using variable length byte structures (e.g., the number `5` takes 1 byte, while `200000` takes more), saving massive storage.
* **No Field Names**: Rather than sending `"username": "alice"`, Protobuf maps values to numeric field IDs (e.g., `1: "alice"`).
* **Code Generation**: A schema compiler (`protoc`) generates static helper classes, eliminating the performance cost of runtime reflection.

---

## 🗂️ Phase 5 Code Directory Structure

* **[01-json-processing/](file:///home/sjarso/go-mastery/phase-5-data-persistence/01-json-processing)**: Reading/Writing JSON.
  * [basic_json.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/01-json-processing/basic_json.go): standard Marshalling and Tag definitions.
  * [custom_json.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/01-json-processing/custom_json.go): Implementing Marshaler interfaces for custom types.
  * [nested_json.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/01-json-processing/nested_json.go): Managing collections and unmarshaling maps.
  * [performance.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/01-json-processing/performance.go): Streaming Decoder utilities.
* **[02-database-fundamentals/](file:///home/sjarso/go-mastery/phase-5-data-persistence/02-database-fundamentals)**: Working with SQL engines.
  * [connection.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/02-database-fundamentals/connection.go): SQLite connections and pooling.
  * [queries.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/02-database-fundamentals/queries.go): Executing SQL safely with parameters.
  * [transactions.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/02-database-fundamentals/transactions.go): Multi-query transaction scoping.
* **[03-orm/](file:///home/sjarso/go-mastery/phase-5-data-persistence/03-orm)**: High-level database wrappers.
  * [models.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/03-orm/models.go): Schemas, relationships, foreign keys.
  * [connection_orm.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/03-orm/connection_orm.go): Running Migrations.
  * [crud.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/03-orm/crud.go): CRUD APIs and loading preloads.
  * [transactions_orm.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/03-orm/transactions_orm.go): Transaction blocks.
* **[04-serialization/](file:///home/sjarso/go-mastery/phase-5-data-persistence/04-serialization)**: Non-JSON standard serialization.
  * [xml.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/04-serialization/xml.go): Parsing structured XML tags.
  * [msgpack.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/04-serialization/msgpack.go): Binary serialization.
  * [protobuf.go](file:///home/sjarso/go-mastery/phase-5-data-persistence/04-serialization/protobuf.go): Parsing protobuf binary payloads.

---

## 🚀 How to Run the Code

Start by going to any serialization/JSON folder:

```bash
# Go to the database fundamentals folder
cd phase-5-data-persistence/02-database-fundamentals

# Run the queries example
go run connection.go queries.go
```

Verify practices and exercises via tests:
```bash
go test -v ./...
```
