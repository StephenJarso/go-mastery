# Go Modules & Dependency Management

**Duration:** 1-2 days  
**Prerequisites:** Phase 2.1-2.4 (all prior Phase 2 topics)  
**Focus:** Managing dependencies, versioning, and project structure

## 🎯 Learning Objectives

After completing this phase, you'll be able to:

- ✅ Understand Go Modules and their purpose
- ✅ Initialize a Go module in your projects
- ✅ Add and manage dependencies
- ✅ Understand semantic versioning
- ✅ Update and remove dependencies
- ✅ Work with go.mod and go.sum files
- ✅ Handle transitive dependencies
- ✅ Use vendor directories
- ✅ Manage versions effectively

## 📚 Topics Covered

### 1. What are Go Modules?

Go Modules are Go's dependency management system (since Go 1.11).

**Before Go Modules:**
- Dependencies stored in `$GOPATH`
- Hard to manage versions
- Reproducibility issues

**With Go Modules:**
- Dependencies per project
- Version control
- Reproducible builds
- Clear dependency graph

### 2. Go Module Basics

**Files:**
- `go.mod` - Defines module and dependencies
- `go.sum` - Hash verification for security

**Basic Commands:**
```bash
go mod init github.com/username/project    # Initialize module
go get github.com/lib/pq                   # Add dependency
go get -u github.com/lib/pq                # Update dependency
go mod tidy                                # Remove unused dependencies
go list -m all                             # List all dependencies
go mod graph                               # Show dependency graph
go mod vendor                              # Create vendor directory
```

### 3. go.mod File Structure

```
module github.com/username/project

go 1.21

require (
    github.com/lib/pq v1.10.9
    github.com/joho/godotenv v1.5.1
)

exclude (
    github.com/broken/package v1.2.3
)

replace (
    github.com/old/path => github.com/new/path v1.2.3
)
```

**Sections:**
- `module` - Your module's name (import path)
- `go` - Go version required
- `require` - Dependencies you need
- `exclude` - Versions to skip
- `replace` - Map one import path to another

### 4. Semantic Versioning

Go uses semantic versioning: `v1.2.3`

- **v1** - Major version (breaking changes)
- **.2** - Minor version (new features, backward compatible)
- **.3** - Patch version (bug fixes, backward compatible)

**Rules:**
- `v1.2.3` - Exact version
- `v1.2` - Latest patch in v1.2.x
- `v1` - Latest in v1.x.x
- `v0` - Pre-release, use with caution

### 5. Adding Dependencies

**Explicit adding:**
```bash
go get github.com/gorilla/mux
go get github.com/gorilla/mux@v1.8.0      # Specific version
go get github.com/gorilla/mux@latest      # Latest version
go get github.com/gorilla/mux@v1.8        # Latest v1.8.x
```

**Automatic adding:**
When you import a package and run `go build` or `go test`, it's automatically added.

### 6. Updating Dependencies

```bash
go get -u                          # Update all dependencies
go get -u github.com/lib/pq        # Update specific dependency
go get -u=patch github.com/lib/pq  # Update only patch versions
```

### 7. Removing Unused Dependencies

```bash
go mod tidy  # Remove unused dependencies, add missing ones
```

### 8. Common Workflows

**Starting a new project:**
```bash
mkdir myproject
cd myproject
go mod init github.com/username/myproject
go get github.com/some/package
# Start coding
```

**Adding to existing project:**
```bash
go mod init github.com/username/project
go mod tidy
```

**Working with private repositories:**
```bash
# Set Git credentials
git config --global url."https://<token>@github.com/".insteadOf "https://github.com/"

# Go will handle the rest
go get github.com/mycompany/private-package
```

### 9. go.sum File

Ensures reproducible builds and detects tampering.

```
github.com/lib/pq v1.10.9 h1:...
github.com/lib/pq v1.10.9/go.mod h1:...
```

**Security:**
- Cryptographic hashes of module contents
- Verified against Go's checksum database
- Commit to version control
- Never edit manually

### 10. Vendoring

Copy dependencies into your repository.

```bash
go mod vendor  # Create vendor/ directory
```

**When to vendor:**
- Working offline
- Private dependencies
- Reproducibility requirements
- Air-gapped deployments

### 11. Replacing Modules

Use locally for development:

```
replace github.com/some/package => ../local/package
```

Or different versions:

```
replace github.com/old/package => github.com/new/package v1.2.3
```

### 12. Working with Multiple Modules

**Workspaces (Go 1.18+):**

```bash
go work init
go work use ./module1 ./module2 ./module3
```

## 📁 Directory Structure

```
myproject/
├── go.mod              # Module definition
├── go.sum              # Checksums
├── main.go             # Your code
├── package1/
│   └── file.go
└── vendor/             # Optional: vendored dependencies
    └── github.com/
```

## 🔑 Important Commands Reference

| Command | Purpose |
|---------|----------|
| `go mod init` | Initialize module |
| `go get` | Add dependency |
| `go get -u` | Update all dependencies |
| `go mod tidy` | Clean up dependencies |
| `go list -m all` | List all dependencies |
| `go mod graph` | Show dependency tree |
| `go mod vendor` | Create vendor directory |
| `go mod download` | Download modules |
| `go mod verify` | Verify integrity |

## 💡 Best Practices

### ✅ DO:
1. **Commit go.mod and go.sum** - Ensures reproducibility
2. **Use specific versions** - Don't use floating versions
3. **Update regularly** - But test thoroughly
4. **Check breaking changes** - Before major upgrades
5. **Use go get -u for patches** - Only patch updates for stability
6. **Run go mod tidy** - Regularly clean up
7. **Test thoroughly** - After dependency updates

### ❌ DON'T:
1. **Edit go.mod manually** - Use go get/mod commands
2. **Ignore security updates** - Keep dependencies current
3. **Use v0 in production** - Pre-release, unstable
4. **Commit vendor carelessly** - It's large
5. **Ignore go.sum changes** - Verify before committing

## 📊 Module Lifecycle

```
1. go mod init
   ↓
2. Add dependencies (go get)
   ↓
3. Write code and import packages
   ↓
4. go mod tidy (clean up)
   ↓
5. Commit go.mod and go.sum
   ↓
6. Update when needed (go get -u)
   ↓
7. Test and commit
```

## 🎯 Common Scenarios

### Scenario 1: Starting a New Project
```bash
$ go mod init github.com/alice/awesome-app
$ go get github.com/gorilla/mux
$ go get github.com/lib/pq
$ go mod tidy
$ git add go.mod go.sum
$ git commit -m "Initial dependencies"
```

### Scenario 2: Adding a Dependency
```bash
$ go get github.com/newpackage/lib
$ go mod tidy
$ go test ./...
$ git add go.mod go.sum
$ git commit -m "Add newpackage/lib"
```

### Scenario 3: Updating All Dependencies
```bash
$ go get -u
$ go mod tidy
$ go test ./...
$ git add go.mod go.sum
$ git commit -m "Update all dependencies"
```

### Scenario 4: Using Local Package During Development
```
// go.mod
module github.com/alice/myapp

require github.com/alice/lib v1.0.0

replace github.com/alice/lib => ../lib
```

## 🔗 Related Resources

- [Go Modules Documentation](https://golang.org/doc/modules)
- [Semantic Versioning](https://semver.org/)
- [Go Wiki: Modules](https://github.com/golang/go/wiki/Modules)
- [Using Go Modules](https://blog.golang.org/using-go-modules)

## ✅ Completion Checklist

- [ ] Understand what Go Modules are
- [ ] Know when to use modules (always in Go 1.11+)
- [ ] Can initialize a module
- [ ] Can add dependencies
- [ ] Can update dependencies
- [ ] Understand go.mod structure
- [ ] Understand go.sum purpose
- [ ] Know semantic versioning
- [ ] Can use go mod tidy
- [ ] Understand vendor directory
- [ ] Can list dependencies
- [ ] Know when to replace modules
- [ ] Understand module caching
- [ ] Can troubleshoot dependency issues

## 🎬 What's Next?

Once Phase 2.5 is complete:
1. You can manage dependencies effectively
2. You understand versioning
3. You can set up new projects
4. You know best practices
5. You're ready for larger projects

**You've completed Phase 2: Object-Oriented & Functional Concepts! 🎉**

Next: Phase 3 - Advanced Language Features (Concurrency, Reflection, Generics)

---

## 📝 Quick Reference: Module Commands

```bash
# Initialize
go mod init github.com/user/project

# Add dependency
go get github.com/gorilla/mux
go get github.com/gorilla/mux@v1.8.0

# Update
go get -u                      # Update all
go get -u github.com/lib/pq    # Update specific
go get -u=patch github.com/lib/pq  # Patch only

# Clean
go mod tidy

# View
go list -m all
go mod graph

# Vendor
go mod vendor

# Verify
go mod verify

# Download
go mod download
```

---

**Phase 2 Complete!** You've mastered structs, methods, interfaces, packages, and dependency management. You're ready for advanced features! 🚀
