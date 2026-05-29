package fileio

import (
	"os"
	"bufio"
	"strings"
	"path/filepath"
)


var _ = os.Open
var _ = bufio.NewScanner
var _ = strings.Contains
var _ = filepath.WalkDir

// Exercise 1: Filter Log File
// Read from srcPath and write only lines containing keyword to destPath.
func FilterLogs(srcPath, destPath, keyword string) error {
	// TODO: Implement
	return nil
}

// Exercise 2: Count Words in File
// Return the total word count in a text file.
func CountWords(filePath string) (int, error) {
	// TODO: Implement
	return 0, nil
}

// Exercise 3: Find Files by Extension
// Recursively walk dirPath and return absolute paths of all files matching the given extension (e.g. ".txt").
func FindFiles(dirPath, ext string) ([]string, error) {
	// TODO: Implement
	return nil, nil
}
