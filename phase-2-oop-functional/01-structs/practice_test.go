package main

import (
	"testing"
)

// TestBookCreation tests creating Book structs
func TestBookCreation(t *testing.T) {
	book := Book{
		ID:        1,
		Title:     "The Go Programming Language",
		Author:    "Donovan & Kernighan",
		Pages:     380,
		Published: 2015,
		Available: true,
	}

	if book.Title != "The Go Programming Language" {
		t.Errorf("expected title to be 'The Go Programming Language', got %s", book.Title)
	}

	if book.Pages != 380 {
		t.Errorf("expected 380 pages, got %d", book.Pages)
	}
}

// TestLibraryCreation tests creating Library structs
func TestLibraryCreation(t *testing.T) {
	library := Library{
		Name:      "City Library",
		Location:  "Boston, MA",
		BookCount: 10,
	}

	if library.Name != "City Library" {
		t.Errorf("expected name 'City Library', got %s", library.Name)
	}

	if library.BookCount != 10 {
		t.Errorf("expected 10 books, got %d", library.BookCount)
	}
}

// TestBookModification tests modifying book fields
func TestBookModification(t *testing.T) {
	book := Book{
		ID:        1,
		Title:     "Go in Action",
		Available: true,
	}

	// Test modification
	book.Available = false

	if book.Available != false {
		t.Error("expected book to be unavailable")
	}
}

// TestBookComparison tests comparing books
func TestBookComparison(t *testing.T) {
	book1 := Book{ID: 1, Title: "Book A", Pages: 300}
	book2 := Book{ID: 1, Title: "Book A", Pages: 300}
	book3 := Book{ID: 2, Title: "Book B", Pages: 250}

	if book1 != book2 {
		t.Error("expected book1 == book2")
	}

	if book1 == book3 {
		t.Error("expected book1 != book3")
	}
}
