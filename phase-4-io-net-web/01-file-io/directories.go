package fileio

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// In Go, operations on file systems and directory paths should use the
// path/filepath package instead of the path package. path/filepath is
// designed to handle operating system specific path separators (e.g., / on Unix, \ on Windows).

// ManageDirectories demonstrates creating, reading, and removing directories.
func ManageDirectories(dirName string) error {
	// 1. Create a single directory.
	// 0755: Owner can read/write/execute, group & others can read/execute.
	err := os.Mkdir(dirName, 0755)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// 2. Create nested directories (e.g., dirName/subdir1/subdir2).
	// MkdirAll acts like 'mkdir -p', creating all parent directories as needed.
	nestedDir := filepath.Join(dirName, "subdir1", "subdir2")
	err = os.MkdirAll(nestedDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create nested directories: %w", err)
	}

	// Create a dummy file in the nested directory to test listing.
	dummyFile := filepath.Join(nestedDir, "test.txt")
	err = os.WriteFile(dummyFile, []byte("hello"), 0644)
	if err != nil {
		return fmt.Errorf("failed to write dummy file: %w", err)
	}

	return nil
}

// ListDirectoryContents lists the files and subdirectories directly inside a directory.
func ListDirectoryContents(dirPath string) ([]string, error) {
	// os.ReadDir reads the directory and returns a sorted slice of fs.DirEntry.
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var results []string
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		// entry.Name() returns just the base name.
		// info.IsDir() checks if it's a directory.
		typeStr := "File"
		if entry.IsDir() {
			typeStr = "Dir"
		}
		results = append(results, fmt.Sprintf("%s (%s, %d bytes)", entry.Name(), typeStr, info.Size()))
	}

	return results, nil
}

// WalkDirectoryTree recursively walks through the directory tree starting from root.
func WalkDirectoryTree(root string) ([]string, error) {
	var files []string

	// filepath.WalkDir walks the file tree rooted at root, calling fn for each file
	// or directory in the tree, including root itself.
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// If there's an error accessing the path, return it to stop walking.
			// Or return nil to skip it and continue.
			return err
		}

		// Rel returns a relative path representation of the walk path relative to root.
		relPath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		if d.IsDir() {
			files = append(files, fmt.Sprintf("[DIR] %s", relPath))
		} else {
			files = append(files, fmt.Sprintf("[FILE] %s", relPath))
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking directory: %w", err)
	}

	return files, nil
}

// PathManipulation demonstrates helpful path utility functions.
func PathManipulation() {
	// Always use filepath.Join to combine path segments. It handles separators correctly.
	p := filepath.Join("home", "user", "documents", "project.go")
	fmt.Printf("Joined path: %s\n", p)

	// filepath.Base returns the last element of path.
	fmt.Printf("Base name: %s\n", filepath.Base(p))

	// filepath.Dir returns all but the last element of path (the directory).
	fmt.Printf("Directory name: %s\n", filepath.Dir(p))

	// filepath.Ext returns the file name extension used by path.
	fmt.Printf("Extension: %s\n", filepath.Ext(p))

	// filepath.Clean returns the shortest path name equivalent to path by purely lexical processing.
	dirtyPath := "home/user/../user/documents/./project.go"
	fmt.Printf("Cleaned path: %s\n", filepath.Clean(dirtyPath))

	// filepath.IsAbs reports whether the path is absolute.
	fmt.Printf("Is path absolute? %v\n", filepath.IsAbs(p))
}
