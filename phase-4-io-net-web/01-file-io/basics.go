package fileio

import (
	"fmt"
	"io"
	"os"
)

// FileIOBasics demonstrates basic file read and write operations.
// In Go, files are represented by the *os.File struct, which implements
// basic interfaces like io.Reader, io.Writer, io.Closer, etc.

// WriteFileSimple writes data to a file in a single operation.
// This is suitable for small files where you don't need streaming.
func WriteFileSimple(filename string, content string) error {
	// os.WriteFile creates or overwrites the file with the specified content.
	// The third parameter '0644' represents Unix file permissions:
	//   - '0' prefix: octal notation
	//   - '6' (owner): read (4) + write (2) = 6
	//   - '4' (group): read (4) = 4
	//   - '4' (others): read (4) = 4
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}

// ReadFileSimple reads the entire file into memory at once.
// Suitable for small files.
func ReadFileSimple(filename string) (string, error) {
	// os.ReadFile reads the whole file and returns its byte contents.
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}
	return string(data), nil
}

// WriteFileAdvanced demonstrates fine-grained file creation/writing using os.OpenFile.
func WriteFileAdvanced(filename string, content string, appendMode bool) error {
	// Flags define how the file should be opened:
	//   - os.O_CREATE: create the file if it doesn't exist
	//   - os.O_WRONLY: open the file for writing only
	//   - os.O_APPEND: append data to the file (instead of truncating/overwriting)
	//   - os.O_TRUNC: truncate (empty) the file if it exists (ignored if O_APPEND is set)
	flags := os.O_CREATE | os.O_WRONLY
	if appendMode {
		flags |= os.O_APPEND
	} else {
		flags |= os.O_TRUNC
	}

	// Open the file with the flags and perm 0644.
	file, err := os.OpenFile(filename, flags, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	// CRITICAL: Always defer closing the file to release system resources.
	// In production code, we should also handle the error returned by Close()
	// when writing, but a basic defer guarantees cleanup.
	defer file.Close()

	// Write content as a byte slice.
	_, err = file.Write([]byte(content))
	if err != nil {
		return fmt.Errorf("failed to write to file handle: %w", err)
	}

	return nil
}

// ReadFileChunked demonstrates reading a file in chunks using a fixed-size buffer.
// This is essential for reading files that are too large to fit in memory all at once.
func ReadFileChunked(filename string, chunkSize int) (string, error) {
	// os.Open opens a file for reading only.
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file for chunked read: %w", err)
	}
	defer file.Close()

	// Create a buffer of the specified size.
	buffer := make([]byte, chunkSize)
	var result []byte

	for {
		// Read reads up to len(buffer) bytes from the file into buffer.
		// It returns the number of bytes read (n) and any error encountered.
		n, err := file.Read(buffer)
		if n > 0 {
			// Append only the bytes that were actually read.
			result = append(result, buffer[:n]...)
		}

		if err == io.EOF {
			// io.EOF (End Of File) indicates we have reached the end of the file.
			break
		}
		if err != nil {
			// Any other error is a real read failure.
			return "", fmt.Errorf("error reading chunk: %w", err)
		}
	}

	return string(result), nil
}
