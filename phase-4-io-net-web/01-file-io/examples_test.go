package fileio

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestBasicFileOperations(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "basic_file_test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	filename := filepath.Join(tempDir, "test.txt")
	content := "Line 1: Go Mastery\nLine 2: File I/O Basics"

	// 1. Write file simple
	if err := WriteFileSimple(filename, content); err != nil {
		t.Fatalf("WriteFileSimple failed: %v", err)
	}

	// 2. Read file simple
	readContent, err := ReadFileSimple(filename)
	if err != nil {
		t.Fatalf("ReadFileSimple failed: %v", err)
	}
	if readContent != content {
		t.Errorf("expected read content %q, got %q", content, readContent)
	}

	// 3. Write advanced (append mode)
	appendContent := "\nLine 3: Appended Content"
	if err := WriteFileAdvanced(filename, appendContent, true); err != nil {
		t.Fatalf("WriteFileAdvanced (append) failed: %v", err)
	}

	// 4. Read chunked
	chunkedContent, err := ReadFileChunked(filename, 10)
	if err != nil {
		t.Fatalf("ReadFileChunked failed: %v", err)
	}
	expectedFullContent := content + appendContent
	if chunkedContent != expectedFullContent {
		t.Errorf("expected chunked content %q, got %q", expectedFullContent, chunkedContent)
	}
}

func TestBufferedFileOperations(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "buffered_file_test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	filename := filepath.Join(tempDir, "buffered.txt")
	lines := []string{"Line A", "Line B", "Line C"}

	// 1. Write buffered
	if err := WriteBuffered(filename, lines); err != nil {
		t.Fatalf("WriteBuffered failed: %v", err)
	}

	// 2. Read line-by-line buffered
	readLines, err := ReadBufferedLineByLine(filename)
	if err != nil {
		t.Fatalf("ReadBufferedLineByLine failed: %v", err)
	}
	expectedLinesWithNewline := []string{"Line A\n", "Line B\n", "Line C\n"}
	if !reflect.DeepEqual(readLines, expectedLinesWithNewline) {
		t.Errorf("expected read lines %v, got %v", expectedLinesWithNewline, readLines)
	}

	// 3. Read with scanner
	scannerLines, err := ReadWithScanner(filename)
	if err != nil {
		t.Fatalf("ReadWithScanner failed: %v", err)
	}
	if !reflect.DeepEqual(scannerLines, lines) {
		t.Errorf("expected scanner lines %v, got %v", lines, scannerLines)
	}
}

func TestDirectoryOperations(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "directory_test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	testDir := filepath.Join(tempDir, "test_structure")

	// 1. Manage directories
	if err := ManageDirectories(testDir); err != nil {
		t.Fatalf("ManageDirectories failed: %v", err)
	}

	// 2. List directory contents
	contents, err := ListDirectoryContents(filepath.Join(testDir, "subdir1"))
	if err != nil {
		t.Fatalf("ListDirectoryContents failed: %v", err)
	}
	if len(contents) == 0 {
		t.Error("expected directory contents to contain subdir2")
	}

	// 3. Walk directory tree
	tree, err := WalkDirectoryTree(testDir)
	if err != nil {
		t.Fatalf("WalkDirectoryTree failed: %v", err)
	}

	// There should be a DIR called subdir1, a DIR called subdir1/subdir2, and a FILE subdir1/subdir2/test.txt
	foundFile := false
	for _, entry := range tree {
		if entry == "[FILE] subdir1/subdir2/test.txt" {
			foundFile = true
			break
		}
	}
	if !foundFile {
		t.Errorf("expected walk tree to contain [FILE] subdir1/subdir2/test.txt, tree was: %v", tree)
	}
}

func TestStreamingOperations(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "streaming_test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	src := filepath.Join(tempDir, "source.bin")
	dst1 := filepath.Join(tempDir, "dest1.bin")
	dst2 := filepath.Join(tempDir, "dest2.bin")

	// Create 1MB of pseudo-random data
	data := make([]byte, 1024*1024)
	for i := range data {
		data[i] = byte(i % 256)
	}
	if err := os.WriteFile(src, data, 0644); err != nil {
		t.Fatalf("failed to create source binary file: %v", err)
	}

	// 1. StreamCopyFile
	written1, err := StreamCopyFile(src, dst1)
	if err != nil {
		t.Fatalf("StreamCopyFile failed: %v", err)
	}
	if written1 != int64(len(data)) {
		t.Errorf("expected copied bytes %d, got %d", len(data), written1)
	}

	// Compare files
	copiedData1, err := os.ReadFile(dst1)
	if err != nil || !reflect.DeepEqual(copiedData1, data) {
		t.Errorf("StreamCopyFile did not copy content correctly")
	}

	// 2. CustomChunkStreamCopy
	written2, err := CustomChunkStreamCopy(src, dst2, 4096) // 4KB chunks
	if err != nil {
		t.Fatalf("CustomChunkStreamCopy failed: %v", err)
	}
	if written2 != int64(len(data)) {
		t.Errorf("expected custom copied bytes %d, got %d", len(data), written2)
	}

	copiedData2, err := os.ReadFile(dst2)
	if err != nil || !reflect.DeepEqual(copiedData2, data) {
		t.Errorf("CustomChunkStreamCopy did not copy content correctly")
	}
}
