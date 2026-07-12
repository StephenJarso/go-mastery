package fileio

import (
	"fmt"
	"io"
	"os"
)

// When processing large files (Gigabytes or Terabytes), loading the entire file
// into memory using os.ReadFile or even bufio.Scanner can cause the program
// to run out of memory (OOM panic).
// To prevent this, we stream files in small chunks or pipe them directly
// from an io.Reader to an io.Writer.

// StreamCopyFile duplicates a file by streaming its contents.
// It uses io.Copy, which internally allocates a small buffer (typically 32KB)
// and pipes data from the source reader to the destination writer.
func StreamCopyFile(src, dst string) (int64, error) {
	sourceFile, err := os.Open(src)
	if err != nil {
		return 0, fmt.Errorf("failed to open source: %w", err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return 0, fmt.Errorf("failed to create destination: %w", err)
	}
	defer destFile.Close()

	// io.Copy streams bytes from sourceFile (io.Reader) to destFile (io.Writer).
	// It reads until EOF on sourceFile and returns the number of bytes copied.
	bytesWritten, err := io.Copy(destFile, sourceFile)
	if err != nil {
		return 0, fmt.Errorf("error during streaming copy: %w", err)
	}

	// Flush and sync to commit write changes to stable storage.
	err = destFile.Sync()
	if err != nil {
		return 0, fmt.Errorf("error syncing destination file: %w", err)
	}

	return bytesWritten, nil
}

// CustomChunkStreamCopy demonstrates how io.Copy works under the hood
// by reading and writing using a custom buffer size.
func CustomChunkStreamCopy(src, dst string, bufferSize int) (int64, error) {
	if bufferSize <= 0 {
		bufferSize = 32 * 1024 // 32KB default
	}

	sourceFile, err := os.Open(src)
	if err != nil {
		return 0, fmt.Errorf("failed to open source file: %w", err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return 0, fmt.Errorf("failed to create destination file: %w", err)
	}
	defer destFile.Close()

	buffer := make([]byte, bufferSize)
	var totalBytes int64

	for {
		// Read a chunk from the source file.
		n, readErr := sourceFile.Read(buffer)
		if n > 0 {
			// Write the chunk to the destination file.
			_, writeErr := destFile.Write(buffer[:n])
			if writeErr != nil {
				return totalBytes, fmt.Errorf("failed to write chunk: %w", writeErr)
			}
			totalBytes += int64(n)
		}

		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			return totalBytes, fmt.Errorf("failed to read chunk: %w", readErr)
		}
	}

	// Sync to ensure all data is written to disk.
	if err := destFile.Sync(); err != nil {
		return totalBytes, fmt.Errorf("failed to sync destination file: %w", err)
	}

	return totalBytes, nil
}
