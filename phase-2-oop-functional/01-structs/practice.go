package main

import "fmt"

// PRACTICE EXERCISE #1: Define Custom Structs
// Complete this exercise to solidify your understanding of structs.

// Book represents a book in a library
type Book struct {
	ID        int
	Title     string
	Author    string
	Pages     int
	Published int // Year published
	Available bool
}

// Library represents a library with a collection of books
type Library struct {
	Name      string
	Location  string
	BookCount int
}

// PracticeExercise1 demonstrates using your custom structs
func PracticeExercise1() {
	fmt.Println("=== Practice Exercise #1: Custom Structs ===")

	// Create a book instance
	book1 := Book{
		ID:        1,
		Title:     "The Go Programming Language",
		Author:    "Donovan & Kernighan",
		Pages:     380,
		Published: 2015,
		Available: true,
	}

	// Create another book
	book2 := Book{
		ID:        2,
		Title:     "Go in Action",
		Author:    "Manning",
		Pages:     457,
		Published: 2015,
		Available: false,
	}

	// Create a library
	library := Library{
		Name:      "City Library",
		Location:  "Boston, MA",
		BookCount: 2,
	}

	// Print the data
	fmt.Printf("\nLibrary: %+v\n", library)
	fmt.Printf("Book 1: %+v\n", book1)
	fmt.Printf("Book 2: %+v\n", book2)

	// Access individual fields
	fmt.Printf("\nBook 1 Title: %s by %s\n", book1.Title, book1.Author)
	fmt.Printf("Book 1 Pages: %d\n", book1.Pages)
	fmt.Printf("Book 1 Published: %d\n", book1.Published)
	fmt.Printf("Book 1 Available: %v\n", book1.Available)

	// Modify fields
	book2.Available = true // Book is now available
	fmt.Printf("\nAfter lending - Book 2 Available: %v\n", book2.Available)

	// Using pointer
	bookPtr := &book1
	bookPtr.Pages = 382 // Update pages
	fmt.Printf("Book 1 Pages (via pointer): %d\n", book1.Pages)
}

// CHALLENGE EXERCISE: Add more methods
// 1. BorrowBook(book *Book) - sets Available to false and returns success message
// 2. ReturnBook(book *Book) - sets Available to true and returns success message
// 3. BookInfo(book Book) string - returns formatted book information

func BorrowBook(book *Book) string {
	book.Available = false
	return fmt.Sprintf("You have successfully borrowed %q.", book.Title)
}

func ReturnBook(book *Book) string {
	book.Available = true
	return fmt.Sprintf("You have successfully returned %q.", book.Title)
}

func BookInfo(book Book) string {
	status := "available"
	if !book.Available {
		status = "borrowed"
	}
	return fmt.Sprintf("%q by %s (%d pages) - %s", book.Title, book.Author, book.Pages, status)
}

func main() {
	PracticeExercise1()
	
	fmt.Println("\n=== Challenge Exercises ===")
	book := Book{
		ID:        3,
		Title:     "Designing Data-Intensive Applications",
		Author:    "Martin Kleppmann",
		Pages:     612,
		Published: 2017,
		Available: true,
	}
	fmt.Println(BookInfo(book))
	fmt.Println(BorrowBook(&book))
	fmt.Println(BookInfo(book))
	fmt.Println(ReturnBook(&book))
	fmt.Println(BookInfo(book))

	fmt.Println("\n=== Practice Exercise Completed ===")
}
