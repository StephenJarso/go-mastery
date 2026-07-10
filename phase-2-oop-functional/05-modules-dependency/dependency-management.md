# Go Modules & Dependency Management - Examples

This file demonstrates concepts related to Go Modules.

## Commands Overview

### Initialize a Module
```bash
go mod init github.com/username/projectname
```

This creates `go.mod` file:
```
module github.com/username/projectname
go 1.21
```

### Add a Dependency
```bash
go get github.com/gorilla/mux
```

Updates `go.mod`:
```
require github.com/gorilla/mux v1.8.1
```

Creates/updates `go.sum` with checksums.

### Update Dependencies

**Update all to latest:**
```bash
go get -u
```

**Update specific package:**
```bash
go get -u github.com/gorilla/mux
```

**Update only patch versions (safer):**
```bash
go get -u=patch github.com/gorilla/mux
```

### Clean Up

**Remove unused, add missing:**
```bash
go mod tidy
```

**Verify integrity:**
```bash
go mod verify
```

### View Dependencies

**List all direct and indirect:**
```bash
go list -m all
```

**Show dependency graph:**
```bash
go mod graph
```

**Show why a package is needed:**
```bash
go mod why github.com/gorilla/mux
```

### Vendoring

**Copy dependencies into vendor/:**
```bash
go mod vendor
```

**Build using vendor:**
```bash
go build -mod=vendor
```

## Semantic Versioning

Go versions follow semver: `vMAJOR.MINOR.PATCH`

- `v0.0.1` - Pre-release, use carefully
- `v1.0.0` - Stable first release
- `v1.2.3` - Major.Minor.Patch
  - Major: Breaking changes
  - Minor: New features (backward compatible)
  - Patch: Bug fixes (backward compatible)

**Version constraints:**
- `@v1.2.3` - Exact version
- `@v1.2` - Latest v1.2.x
- `@v1` - Latest v1.x.x
- `@latest` - Latest version
- `@>= v1.2.0` - Range constraints

## File Structure Example

```
myproject/
├── go.mod                 # Module definition
├── go.sum                 # Checksums for reproducibility
├── main.go                # Your code
├── handlers/
│   └── user.go
├── models/
│   └── user.go
└── vendor/                # Optional: local copy of dependencies
    └── github.com/
        └── gorilla/
            └── mux/
```

## go.mod Structure

```
module github.com/username/myapp     # Your module path (import path)

go 1.21                              # Minimum Go version

require (                            # Direct dependencies
    github.com/gorilla/mux v1.8.1
    github.com/lib/pq v1.10.9
)

require (                            # Indirect dependencies (managed by Go)
    github.com/some/package v1.2.3 // indirect
)

exclude (                            # Versions to skip
    github.com/broken/package v1.2.3
)

replace (                            # Path replacements
    github.com/original/lib => github.com/fork/lib v1.0.0
    github.com/local/lib => ../local/lib
)
```

## Common Workflows

### Starting a New Project
```bash
$ mkdir myapp
$ cd myapp
$ go mod init github.com/alice/myapp
$ go get github.com/gorilla/mux
$ go get github.com/lib/pq
$ go mod tidy
$ git add go.mod go.sum
$ git commit -m "Initial setup"
```

### Collaborating on Project
```bash
# Team member clones repo
$ git clone https://github.com/alice/myapp.git
$ cd myapp
# go.mod and go.sum already define dependencies
$ go mod download  # Download all dependencies
$ go test ./...
```

### Updating for Bug Fix
```bash
# Update all patch versions (safe)
$ go get -u=patch
$ go mod tidy
$ go test ./...
$ git add go.mod go.sum
$ git commit -m "Update dependencies (patches)"
```

### Adding New Dependency
```bash
$ go get github.com/newlib/package
# Code imports the package
$ go mod tidy
$ go test ./...
$ git add go.mod go.sum
$ git commit -m "Add newlib/package"
```

### Local Development with Fork
```
# go.mod
module github.com/alice/myapp

require github.com/bob/lib v1.0.0

replace github.com/bob/lib => ../bob-lib-fork
```

## go.sum Purpose

Ensures reproducible builds and security:

```
github.com/gorilla/mux v1.8.1 h1:cLLB7ypYfwvjG5Rz/dDZchTSL9fUzL6mKPxWqFv7g0=
github.com/gorilla/mux v1.8.1/go.mod h1:3LBFNW88K7Fn7g1t/91a0gS/tXhK5C2aK2p4pxP9NA=
```

- First line: Hash of module contents
- Second line: Hash of go.mod of that module
- Verified against Go's checksum database
- Detects tampering and corruption
- ALWAYS commit to version control

## Best Practices

1. **Commit go.mod and go.sum** - Ensures everyone uses same versions
2. **Use `go get -u=patch` by default** - Safer than major updates
3. **Review updates** - Check changelogs before updating
4. **Run tests after updates** - Verify compatibility
5. **Use `go mod tidy` regularly** - Keep dependencies clean
6. **Use named versions** - Not "latest" or floating versions
7. **Keep go.sum clean** - Let Go manage it
8. **Vendor for offline/air-gapped** - But usually unnecessary

## Troubleshooting

### Module not found
```bash
$ go get module/path@version
# Check network connectivity
# Check module name spelling
```

### Conflicting versions
```bash
# Check dependency graph
$ go mod graph | grep package
# May need to update or replace
```

### go.sum mismatch
```bash
# Verify integrity
$ go mod verify
# May indicate network issues or tampering
```

### Build fails with new dependency
```bash
# Ensure all dependencies are available
$ go mod download
$ go mod tidy
# Try building again
```

## Resources

- [Go Modules Documentation](https://golang.org/doc/modules)
- [Go Blog: Using Go Modules](https://blog.golang.org/using-go-modules)
- [Semantic Versioning](https://semver.org/)
- [Go FAQ: Modules](https://golang.org/doc/faq#modules)

---

For more details, run: `go help mod`
