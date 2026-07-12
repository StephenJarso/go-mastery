package fileio

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// PRACTICE EXERCISE #1: Word Counter
// Implement a function that reads a file and counts the total number of words in it.
// To do this efficiently, use a bufio.Scanner and configure it to split by words
// using scanner.Split(bufio.ScanWords).
func CountWordsInFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// bufio.ScanWords is a built-in split function that splits input into space-separated words.
	scanner.Split(bufio.ScanWords)

	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error scanning file: %w", err)
	}

	return wordCount, nil
}

// PRACTICE EXERCISE #2: Log Filter
// Implement a function that reads a log file line-by-line, filters for lines
// that contain the substring `filterWord`, and writes only those matching lines
// to a new destination file. Use bufio.Scanner for reading and bufio.Writer for writing
// to ensure it is highly efficient and safe for large log files.
func FilterLogFile(srcLog, dstLog, filterWord string) (int, error) {
	inputFile, err := os.Open(srcLog)
	if err != nil {
		return 0, fmt.Errorf("failed to open source log: %w", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(dstLog)
	if err != nil {
		return 0, fmt.Errorf("failed to create destination log: %w", err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	matchedLines := 0

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, filterWord) {
			_, err := writer.WriteString(line + "\n")
			if err != nil {
				return matchedLines, fmt.Errorf("failed to write to buffer: %w", err)
			}
			matchedLines++
		}
	}

	if err := scanner.Err(); err != nil {
		return matchedLines, fmt.Errorf("scanner encountered error: %w", err)
	}

	// CRITICAL: Remember to flush the buffered writer to ensure all filtered
	// logs are actually written to the destination file.
	err = writer.Flush()
	if err != nil {
		return matchedLines, fmt.Errorf("failed to flush output file: %w", err)
	}

	return matchedLines, nil
}

// PRACTICE EXERCISE #3: Smart File Syncer
// Implement a function that compares a source file and a destination file.
// - If the destination file does not exist, copy the source file to the destination.
// - If it does exist, compare their file sizes using os.Stat. If they have the same size, do nothing.
// - If their sizes differ, overwrite the destination file with the source file using streaming copy.
// Return true if a copy operation was performed, false otherwise.
func SmartFileSync(src, dst string) (bool, error) {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return false, fmt.Errorf("failed to stat source file: %w", err)
	}

	dstInfo, err := os.Stat(dst)
	if err != nil {
		if os.IsNotExist(err) {
			// Destination file doesn't exist, we must copy it.
			_, copyErr := streamCopy(src, dst)
			if copyErr != nil {
				return false, copyErr
			}
			return true, nil
		}
		return false, fmt.Errorf("failed to stat destination file: %w", err)
	}

	// If destination file exists, compare sizes.
	if srcInfo.Size() == dstInfo.Size() {
		// File sizes are equal, assume they are synced (no action needed).
		return false, nil
	}

	// Sizes are different, overwrite destination with source.
	_, copyErr := streamCopy(src, dst)
	if copyErr != nil {
		return false, copyErr
	}
	return true, nil
}

// Helper function to stream copy
func streamCopy(src, dst string) (int64, error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()

	n, err := io.Copy(dstFile, srcFile)
	if err != nil {
		return 0, err
	}

	err = dstFile.Sync()
	return n, err
}
