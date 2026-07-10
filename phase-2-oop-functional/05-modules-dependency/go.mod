// Example go.mod file for a Go project
// This file defines the module and its dependencies

module github.com/StephenJarso/go-mastery

go 1.21

// require section lists all dependencies
// Each line is: import_path version
// Versions follow semantic versioning: vMAJOR.MINOR.PATCH

require (
    // Example dependencies (these would be actual packages in a real project)
    // github.com/gorilla/mux v1.8.1
    // github.com/lib/pq v1.10.9
    // github.com/joho/godotenv v1.5.1
)

// indirect comments appear when a dependency is required by another dependency
// Go manages these automatically

// exclude prevents specific versions from being used
// Uncomment if needed:
// exclude (
//     github.com/broken/package v1.2.3
// )

// replace allows mapping one import path to another
// Useful for:
// 1. Local development: replace github.com/alice/lib => ../lib
// 2. Fork usage: replace github.com/original/lib => github.com/fork/lib v1.0.0
// Uncomment if needed:
// replace (
//     github.com/somepackage => ../local/path
// )
