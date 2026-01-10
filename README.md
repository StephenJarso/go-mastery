# Go Mastery 🚀

A comprehensive, hands-on learning repository for mastering Go programming language from the fundamentals onwards. This repo is structured as a progressive journey from Object-Oriented & Functional concepts through Advanced Features, with working code examples, tests, and real-world projects.

## 📚 What's Inside

This repository covers:

- **Phase 1:** Go Fundamentals (Variables, Control Flow, Functions, Pointers & Collections)
- **Phase 2:** Object-Oriented & Functional Concepts (Structs, Methods, Interfaces, Packages)
- **Phase 3:** Advanced Language Features (Concurrency, Reflection, Generics, Error Handling)
- **Phase 4:** I/O, Networking & Web Development
- **Phase 5:** Data & Persistence (JSON, Databases, ORMs)
- **Phase 6:** Testing & Quality Assurance
- **Projects:** Real-world applications tying everything together

## 🎯 Learning Path

Each phase builds on the previous one. Start with Phase 1 and progress sequentially:

```
Phase 1: Variables → Control Flow → Functions → Pointers
    ↓
Phase 2: Structs → Methods → Interfaces → Packages → Modules
    ↓
Phase 3: Goroutines → Channels → Patterns → Reflection → Generics
    ↓
Phase 4: File I/O → Networking → HTTP → Frameworks → REST APIs
    ↓
Phase 5: JSON → Databases → ORMs → Data Serialization
    ↓
Phase 6: Unit Testing → Benchmarking → Integration Tests → Quality
    ↓
Projects: Apply everything in real-world scenarios
```

## 🗂️ Repository Structure

```
go-mastery/
├── README.md                          # This file
├── ROADMAP.md                         # Phase progression tracker
├── resources.md                       # Learning resources & references
├── .gitignore
│
├── phase-1-fundamentals/              # Variables, Loops, Functions, Pointers
│   ├── 01-variables-types/
│   ├── 02-control-flow/
│   ├── 03-functions/
│   ├── 04-pointers/
│   └── README.md
│
├── phase-2-oop-functional/            # Structs, Methods, Interfaces
│   ├── 01-structs/
│   ├── 02-methods/
│   ├── 03-interfaces/
│   ├── 04-packages/
│   ├── 05-modules-dependency/
│   └── README.md
│
├── phase-3-advanced/                  # Concurrency, Reflection, Generics
│   ├── 01-concurrency/
│   ├── 02-reflection/
│   ├── 03-generics/
│   ├── 04-error-handling-advanced/
│   └── README.md
│
├── phase-4-io-networking/             # File I/O, HTTP, Web Development
│   ├── 01-file-io/
│   ├── 02-networking-basics/
│   ├── 03-http-web/
│   ├── 04-frameworks/
│   ├── 05-rest-api/
│   └── README.md
│
├── phase-5-data-persistence/          # JSON, Databases, ORMs
│   ├── 01-json/
│   ├── 02-database-sql/
│   ├── 03-orms/
│   └── README.md
│
├── phase-6-testing/                   # Testing & Quality
│   ├── 01-unit-testing/
│   ├── 02-benchmarking/
│   ├── 03-advanced-testing/
│   └── README.md
│
└── projects/                          # End-to-end projects
    ├── 01-cli-tool/
    ├── 02-rest-api/
    ├── 03-concurrent-system/
    └── README.md
```

## 🚀 How to Use This Repo

### 1. **Clone the Repository**
```bash
git clone https://github.com/StephenJarso/go-mastery.git
cd go-mastery
```

### 2. **Start with Phase 2**
Each phase has a dedicated folder with organized topics:
```bash
cd phase-2-oop-functional/01-structs
```

### 3. **Read the Code**
- Each file has **detailed comments** explaining concepts
- Code is organized progressively from simple to complex
- Examples demonstrate practical usage

### 4. **Run the Examples**
```bash
go run basics.go          # Run a single example
go test -v               # Run all tests in directory
go test -v ./...         # Run all tests recursively
```

### 5. **Study the Tests**
- Tests show how code is used correctly
- Table-driven tests demonstrate edge cases
- Use `*_test.go` files as learning resources

### 6. **Track Your Progress**
- Check off items in `ROADMAP.md` as you complete them
- Use the checklist to ensure comprehensive learning

## 📖 Phase Overview

### Phase 2: Object-Oriented & Functional Concepts
**Duration:** 2-3 weeks  
**Focus:** Structs, methods, interfaces, package management

- Understand struct definition and composition
- Learn method receivers (value vs. pointer)
- Master interfaces and duck typing
- Work with standard library packages
- Manage dependencies with Go Modules

**Topics:**
1. Structs (basics, embedding, tags)
2. Methods (value receivers, pointer receivers)
3. Interfaces (implementation, type assertions)
4. Packages (standard library overview)
5. Modules (dependency management)

### Phase 3: Advanced Language Features
**Duration:** 3-4 weeks  
**Focus:** Concurrency, reflection, generics, advanced error handling

- Launch and manage goroutines
- Work with channels and synchronization primitives
- Apply concurrent design patterns
- Use reflection for meta-programming
- Leverage generics for type-safe code
- Master advanced error handling

**Topics:**
1. Goroutines & Channels
2. Concurrency Patterns (worker pool, fan-out/fan-in, pipelines)
3. Reflection
4. Generics
5. Advanced Error Handling

### Phase 4: I/O, Networking & Web Development
**Duration:** 4-5 weeks  
**Focus:** File I/O, HTTP, web frameworks, REST APIs

- Read and write files efficiently
- Build HTTP servers and clients
- Create REST APIs
- Learn web frameworks (Gin, Echo)
- Handle middleware and routing

**Topics:**
1. File I/O Operations
2. Networking Basics (TCP/UDP)
3. HTTP & Web Development
4. Web Frameworks
5. REST API Development

### Phase 5: Data & Persistence
**Duration:** 3-4 weeks  
**Focus:** JSON, databases, ORMs

- Encode/decode JSON
- Work with SQL databases
- Use ORMs like GORM
- Handle transactions and migrations

**Topics:**
1. JSON Processing
2. Database/SQL
3. ORMs (GORM)
4. Data Serialization

### Phase 6: Testing & Quality
**Duration:** 2-3 weeks  
**Focus:** Testing, benchmarking, profiling

- Write unit tests
- Create benchmarks
- Perform profiling
- Understand code quality tools

**Topics:**
1. Unit Testing
2. Benchmarking
3. Advanced Testing
4. Code Quality

## 💡 Learning Tips

1. **Read the comments first** - Each file has detailed explanations
2. **Run the code** - Don't just read, execute and experiment
3. **Modify examples** - Change code to see what happens
4. **Write your own tests** - Solidify understanding by testing
5. **Build projects** - Apply learning in real-world scenarios
6. **Read the tests** - Tests show proper usage patterns
7. **Take breaks** - Don't rush; absorb concepts gradually

## 🔗 Quick Links

- [ROADMAP.md](./ROADMAP.md) - Phase progression checklist
- [resources.md](./resources.md) - External learning resources
- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com)
- [Effective Go](https://golang.org/doc/effective_go)

## 📊 Progression Timeline

| Phase | Duration | Status |
|-------|----------|--------|
| Phase 1 | 1-2 weeks | 🟢 Completed |
| Phase 2 | 2-3 weeks | 🟢 Completed |
| Phase 3 | 3-4 weeks | 📚 Coming Soon |
| Phase 4 | 4-5 weeks | 📚 Coming Soon |
| Phase 5 | 3-4 weeks | 📚 Coming Soon |
| Phase 6 | 2-3 weeks | 📚 Coming Soon |
| Projects | Ongoing | 🚀 Coming Soon |

**Total Time to Advanced:** ~15-20 weeks (3-5 months) of consistent study

## 🎓 What You'll Learn

After completing this repo, you'll be able to:

✅ Design and implement Go applications with clean architecture  
✅ Write concurrent, high-performance Go code  
✅ Build REST APIs and web services  
✅ Work with databases and persistence layers  
✅ Write comprehensive tests and benchmarks  
✅ Follow Go idioms and best practices  
✅ Contribute to real Go projects  
✅ Debug and optimize Go applications  

## 🤝 Contributing

This is your personal learning repository, but feel free to:
- Add your own notes and examples
- Create variations of examples
- Build additional projects
- Share insights and discoveries

## 📝 License

MIT License - Feel free to use this repository for learning and teaching.

## 🎯 Next Steps

1. ✅ You're reading this file
2. → Start with `phase-1-fundamentals/README.md`
3. → Work through `01-variables-types/` examples
4. → Progress through each topic sequentially
5. → Check off items in `ROADMAP.md`
6. → Build projects to apply learning

---

**Happy Learning! 🚀**

*Last Updated: July 2026*
