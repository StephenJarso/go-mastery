package fileio

import (
	"testing"
	"os"
	"path/filepath"
)


func TestFilterLogs(t *testing.T) {
	tmpDir := t.TempDir()
	src := filepath.Join(tmpDir, "src.log")
	dest := filepath.Join(tmpDir, "dest.log")

	content := "INFO: start\nERROR: crash\nINFO: progress\n"
	err := os.WriteFile(src, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	err = FilterLogs(src, dest, "ERROR")
	if err != nil {
		t.Fatalf("FilterLogs failed: %v", err)
	}

	res, err := os.ReadFile(dest)
	if err != nil {
		t.Fatalf("failed to read dest log: %v", err)
	}

	expected := "ERROR: crash\n"
	if string(res) != expected {
		t.Errorf("expected %q, got %q", expected, string(res))
	}
}

func TestCountWords(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "text.txt")
	err := os.WriteFile(path, []byte("hello world this is Go"), 0644)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	count, err := CountWords(path)
	if err != nil || count != 5 {
		t.Errorf("expected 5 words, got %d, err: %v", count, err)
	}
}

func TestFindFiles(t *testing.T) {
	tmpDir := t.TempDir()
	err := os.WriteFile(filepath.Join(tmpDir, "a.txt"), []byte("a"), 0644)
	if err != nil {
		t.Fatalf("failed to create temp files: %v", err)
	}
	err = os.WriteFile(filepath.Join(tmpDir, "b.log"), []byte("b"), 0644)
	if err != nil {
		t.Fatalf("failed to create temp files: %v", err)
	}

	matches, err := FindFiles(tmpDir, ".txt")
	if err != nil || len(matches) != 1 {
		t.Errorf("expected 1 match, got %v, err: %v", matches, err)
	}
}
