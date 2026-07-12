package fileio

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Buffered I/O is used to minimize the number of system calls.
// Instead of talking directly to the disk for every single byte,
// the bufio package keeps an in-memory buffer (usually 4KB or 8KB).
// Reading/writing from the buffer is extremely fast.

// WriteBuffered demonstrates how to write data using bufio.Writer.
func WriteBuffered(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Create a new buffered writer wrapping the file.
	writer := bufio.NewWriter(file)

	for _, line := range lines {
		// WriteString writes the string to the internal buffer.
		// No data is actually written to the disk at this point!
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("failed to write to buffer: %w", err)
		}
	}

	// CRITICAL: You MUST call Flush() to write any remaining buffered data
	// to the underlying file descriptor. If you forget this, some or all of
	// your data will be lost when the program exits!
	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("failed to flush buffer to disk: %w", err)
	}

	return nil
}

// ReadBufferedLineByLine demonstrates reading a file line-by-line using bufio.Reader.
func ReadBufferedLineByLine(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create a new buffered reader.
	reader := bufio.NewReader(file)
	var lines []string

	for {
		// ReadString reads until the first occurrence of the delimiter (here, '\n').
		// It returns the string containing the delimiter.
		line, err := reader.ReadString('\n')
		if len(line) > 0 {
			lines = append(lines, line)
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading line: %w", err)
		}
	}

	return lines, nil
}

// ReadWithScanner demonstrates using bufio.Scanner.
// This is the most idiomatic way in Go to read line-oriented text files.
// It handles splitting lines and stripping the trailing newline character automatically.
func ReadWithScanner(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create a new scanner.
	scanner := bufio.NewScanner(file)
	var lines []string

	// Scan() advances the scanner to the next token (by default, line).
	// It returns false when there are no more tokens or an error occurred.
	for scanner.Scan() {
		// Text() returns the most recent token (line) as a string, excluding the newline.
		lines = append(lines, scanner.Text())
	}

	// Err() returns the first non-EOF error that was encountered by the scanner.
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner encountered error: %w", err)
	}

	return lines, nil
}
