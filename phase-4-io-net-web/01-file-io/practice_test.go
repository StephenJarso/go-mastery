package fileio

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCountWordsInFile(t *testing.T) {
	// Create a temporary directory for testing.
	tempDir, err := os.MkdirTemp("", "wordcount_test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	tempFile := filepath.Join(tempDir, "sample.txt")
	content := "Hello world! This is a simple test of word counting in Go."
	err = os.WriteFile(tempFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to write temp file: %v", err)
	}

	count, err := CountWordsInFile(tempFile)
	if err != nil {
		t.Fatalf("CountWordsInFile failed: %v", err)
	}

	expectedWords := 12 // Hello(1) world!(2) This(3) is(4) a(5) simple(6) test(7) of(8) word(9) counting(10) in(11) Go.(12)
	if count != expectedWords {
		t.Errorf("expected %d words, got %d", expectedWords, count)
	}
}

func TestFilterLogFile(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "logfilter_test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	srcLog := filepath.Join(tempDir, "app.log")
	dstLog := filepath.Join(tempDir, "error.log")

	logContent := `[INFO] Application starting
[WARN] Configuration parameter missing, using default
[ERROR] Database connection failed
[INFO] Attempting reconnect
[ERROR] Authentication service offline`

	err = os.WriteFile(srcLog, []byte(logContent), 0644)
	if err != nil {
		t.Fatalf("failed to write source log file: %v", err)
	}

	count, err := FilterLogFile(srcLog, dstLog, "[ERROR]")
	if err != nil {
		t.Fatalf("FilterLogFile failed: %v", err)
	}

	if count != 2 {
		t.Errorf("expected 2 matched lines, got %d", count)
	}

	filteredBytes, err := os.ReadFile(dstLog)
	if err != nil {
		t.Fatalf("failed to read filtered log: %v", err)
	}

	expectedOutput := "[ERROR] Database connection failed\n[ERROR] Authentication service offline\n"
	if string(filteredBytes) != expectedOutput {
		t.Errorf("expected filtered content:\n%q\nbut got:\n%q", expectedOutput, string(filteredBytes))
	}
}

func TestSmartFileSync(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "filesync_test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	srcFile := filepath.Join(tempDir, "source.txt")
	dstFile := filepath.Join(tempDir, "dest.txt")

	// 1. Write source file
	srcContent := "Hello synchronization!"
	err = os.WriteFile(srcFile, []byte(srcContent), 0644)
	if err != nil {
		t.Fatalf("failed to write source file: %v", err)
	}

	// Case 1: Destination does not exist -> should copy.
	copied, err := SmartFileSync(srcFile, dstFile)
	if err != nil {
		t.Fatalf("SmartFileSync failed (Case 1): %v", err)
	}
	if !copied {
		t.Error("expected copied to be true when destination does not exist")
	}

	// Verify dest has correct content
	dstBytes, err := os.ReadFile(dstFile)
	if err != nil || string(dstBytes) != srcContent {
		t.Errorf("expected dest content %q, got %q (err: %v)", srcContent, string(dstBytes), err)
	}

	// Case 2: Destination exists and has the same size -> should NOT copy.
	copied, err = SmartFileSync(srcFile, dstFile)
	if err != nil {
		t.Fatalf("SmartFileSync failed (Case 2): %v", err)
	}
	if copied {
		t.Error("expected copied to be false when destination exists and is same size")
	}

	// Case 3: Destination exists but has a different size -> should copy/overwrite.
	err = os.WriteFile(dstFile, []byte("short"), 0644)
	if err != nil {
		t.Fatalf("failed to rewrite destination file with different size: %v", err)
	}

	copied, err = SmartFileSync(srcFile, dstFile)
	if err != nil {
		t.Fatalf("SmartFileSync failed (Case 3): %v", err)
	}
	if !copied {
		t.Error("expected copied to be true when destination exists but has different size")
	}

	dstBytes, err = os.ReadFile(dstFile)
	if err != nil || string(dstBytes) != srcContent {
		t.Errorf("expected dest to be updated to %q, got %q (err: %v)", srcContent, string(dstBytes), err)
	}
}
