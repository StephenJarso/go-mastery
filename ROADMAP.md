# Go Mastery Learning Roadmap

Use this document to track your progress through each phase. Check off items as you complete them!

## Phase 2: Object-Oriented & Functional Concepts
**Estimated Duration:** 2-3 weeks  
**Status:** 🟢 Active

### Structs
- [ ] Struct definition and field declaration
- [ ] Creating struct instances (with and without field names)
- [ ] Accessing and modifying fields
- [ ] Struct embedding (composition over inheritance)
- [ ] Struct tags (JSON, XML, validation)
- [ ] Anonymous structs
- [ ] Nested structs
- [ ] Completed: `phase-2-oop-functional/01-structs/`

### Methods
- [ ] Method definition with receivers
- [ ] Value receivers vs. pointer receivers
- [ ] When to use each receiver type
- [ ] Method chaining patterns
- [ ] Methods on interface types
- [ ] Practical method patterns
- [ ] Completed: `phase-2-oop-functional/02-methods/`

### Interfaces
- [ ] Interface declaration and definition
- [ ] Implicit interface implementation (duck typing)
- [ ] Empty interface `interface{}`
- [ ] Type assertions and type switches
- [ ] Common interfaces (io.Reader, io.Writer, fmt.Stringer)
- [ ] Interface composition
- [ ] Interface best practices
- [ ] Completed: `phase-2-oop-functional/03-interfaces/`

### Packages & Standard Library
- [ ] Package organization and naming
- [ ] Exported vs. unexported identifiers
- [ ] `fmt` package (formatting, printing)
- [ ] `strings` package (string operations)
- [ ] `strconv` package (string conversions)
- [ ] `io` and `ioutil` packages
- [ ] `os` package (OS interaction)
- [ ] `time` package (time and duration)
- [ ] Creating custom packages
- [ ] Completed: `phase-2-oop-functional/04-packages/`

### Modules & Dependency Management
- [ ] Understanding Go Modules
- [ ] `go mod init` command
- [ ] `go.mod` and `go.sum` files
- [ ] Adding dependencies with `go get`
- [ ] `go mod tidy` and dependency cleanup
- [ ] Vendoring with `go mod vendor`
- [ ] Version management and constraints
- [ ] Completed: `phase-2-oop-functional/05-modules-dependency/`

**Phase 2 Completion:** All items checked ✓

---

## Phase 3: Advanced Language Features
**Estimated Duration:** 3-4 weeks  
**Status:** 🟢 Active

### Concurrency Fundamentals
- [x] Goroutines and their lifecycle
- [x] Launching goroutines with `go` keyword
- [x] WaitGroup for synchronization
- [x] Channels (unbuffered)
- [x] Buffered channels
- [x] Channel operations (send, receive, close)
- [x] Completed: `phase-3-advanced/01-concurrency/`

### Concurrency Primitives & Synchronization
- [x] `sync.Mutex` for mutual exclusion
- [x] `sync.RWMutex` for read-write locks
- [x] `sync.Cond` for condition variables
- [x] `sync.Once` for one-time initialization
- [x] Race conditions and detection
- [x] Deadlock prevention
- [x] Completed: `phase-3-advanced/01-concurrency/`

### Concurrency Patterns
- [x] `select` statement for multiplexing
- [x] Worker pool pattern
- [x] Fan-out/fan-in pattern
- [x] Pipeline pattern
- [x] Context package for cancellation
- [x] Timeouts with context
- [x] Completed: `phase-3-advanced/01-concurrency/patterns/`

### Reflection
- [ ] `reflect` package basics
- [ ] `reflect.Type` and `reflect.Value`
- [ ] Inspecting struct fields
- [ ] Type reflection
- [ ] Value manipulation through reflection
- [ ] Method invocation via reflection
- [ ] When to use reflection
- [ ] Completed: `phase-3-advanced/02-reflection/`

### Generics
- [ ] Generic type parameters (Go 1.18+)
- [ ] Generic functions
- [ ] Type constraints and bounds
- [ ] Practical generic patterns
- [ ] When generics help and hurt
- [ ] Built-in generic constraints
- [ ] Completed: `phase-3-advanced/03-generics/`

### Advanced Error Handling
- [ ] Error wrapping with `fmt.Errorf` and `%w`
- [ ] Error unwrapping with `errors.Is()`
- [ ] Error type checking with `errors.As()`
- [ ] Custom error types with methods
- [ ] Error handling strategies
- [ ] Panic and recover patterns
- [ ] Completed: `phase-3-advanced/04-error-handling-advanced/`

**Phase 3 Completion:** All items checked ✓

---

## Phase 4: I/O, Networking & Web Development
**Estimated Duration:** 4-5 weeks  
**Status:** 🔴 Not Started

### File I/O Operations
- [ ] Reading files with `os.Open`
- [ ] Writing files
- [ ] Buffered I/O with `bufio`
- [ ] Working with directories
- [ ] File paths and the `filepath` package
- [ ] Streaming large files
- [ ] File permissions and attributes

### Networking Basics
- [ ] TCP socket programming
- [ ] UDP socket programming
- [ ] Connection handling
- [ ] Port listening and accepting connections
- [ ] Client-server architecture

### HTTP & Web Development
- [ ] `net/http` package fundamentals
- [ ] Building HTTP servers
- [ ] HTTP handlers and multiplexing
- [ ] Making HTTP requests with HTTP client
- [ ] Request/response handling
- [ ] Status codes and headers
- [ ] File serving and static content
- [ ] Middleware patterns
- [ ] Basic routing

### Web Frameworks
- [ ] Gin framework basics
- [ ] Echo framework basics
- [ ] Chi router
- [ ] Framework comparison
- [ ] Building apps with frameworks

### REST API Development
- [ ] Designing REST endpoints
- [ ] CRUD operations
- [ ] Request validation
- [ ] Response formatting
- [ ] Error responses
- [ ] HTTP status codes
- [ ] API versioning

**Phase 4 Completion:** All items checked ✓

---

## Phase 5: Data & Persistence
**Estimated Duration:** 3-4 weeks  
**Status:** 🔴 Not Started

### JSON Processing
- [ ] Marshaling Go types to JSON
- [ ] Unmarshaling JSON to Go types
- [ ] Struct tags for JSON mapping
- [ ] Custom marshaling/unmarshaling
- [ ] Handling nested JSON
- [ ] JSON performance considerations

### Database Fundamentals
- [ ] `database/sql` package
- [ ] Connecting to databases
- [ ] Executing queries
- [ ] Prepared statements
- [ ] Result handling and scanning
- [ ] Connection pooling
- [ ] Transactions

### ORMs
- [ ] GORM basics
- [ ] Model definition
- [ ] CRUD operations with ORM
- [ ] Query building
- [ ] Associations (One-to-Many, Many-to-Many)
- [ ] Migrations
- [ ] Transactions with ORM

### Data Serialization
- [ ] Protocol Buffers (protobuf)
- [ ] MessagePack
- [ ] XML encoding/decoding

**Phase 5 Completion:** All items checked ✓

---

## Phase 6: Testing & Quality
**Estimated Duration:** 2-3 weeks  
**Status:** 🔴 Not Started

### Unit Testing
- [ ] `testing` package basics
- [ ] Writing test functions
- [ ] Table-driven tests
- [ ] Test organization
- [ ] Coverage analysis (`go cover`)
- [ ] Test utilities and helpers

### Benchmarking
- [ ] Benchmark functions
- [ ] Performance measurement
- [ ] Running benchmarks
- [ ] Analyzing results

### Advanced Testing
- [ ] Mocking and test doubles
- [ ] Integration testing
- [ ] End-to-end testing
- [ ] Test fixtures

### Code Quality
- [ ] `gofmt` and formatting
- [ ] `go vet` static analysis
- [ ] `golint` linting
- [ ] Code review practices
- [ ] Idiomatic Go patterns

**Phase 6 Completion:** All items checked ✓

---

## Projects
**Status:** 🔴 Not Started

### Project 1: CLI Tool
- [ ] Design
- [ ] Implementation
- [ ] Testing
- [ ] Documentation

### Project 2: REST API
- [ ] Design
- [ ] Implementation
- [ ] Testing
- [ ] Deployment

### Project 3: Concurrent System
- [ ] Design
- [ ] Implementation
- [ ] Testing
- [ ] Performance optimization

---

## Summary Statistics

**Total Items to Complete:** TBD
**Items Completed:** 0
**Completion Percentage:** 0%

**Phases Completed:** 0/6
**Estimated Total Time:** 15-20 weeks

---

## Tips for Using This Roadmap

1. **Start at Phase 2** - You likely know Phase 1 (fundamentals)
2. **Check off as you go** - Celebrate small wins
3. **Don't skip items** - Each builds on previous knowledge
4. **Do the code exercises** - Reading isn't enough
5. **Build projects** - Apply learning in real scenarios
6. **Review periodically** - Reinforce concepts
7. **Update as you progress** - Keep this current

**Remember:** Mastery takes time. Focus on understanding, not speed!
