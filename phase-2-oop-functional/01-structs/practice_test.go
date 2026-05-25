package main

import (
	"testing"
	"encoding/json"
)


func TestBorrowBook(t *testing.T) {
	book := Book{ID: 1, Title: "Go Book", Author: "Author", Available: true}
	msg, err := BorrowBook(&book)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if book.Available {
		t.Error("expected book to be unavailable")
	}
	if msg != `You have successfully borrowed "Go Book".` {
		t.Errorf("unexpected message: %q", msg)
	}

	_, err = BorrowBook(&book)
	if err == nil {
		t.Error("expected error when borrowing already borrowed book")
	}
}

func TestNewEmployee(t *testing.T) {
	emp := NewEmployee("John", 30, "Boston", "MA", 101, "Developer")
	if emp.Name != "John" || emp.Age != 30 || emp.City != "Boston" || emp.State != "MA" || emp.ID != 101 || emp.Position != "Developer" {
		t.Errorf("unexpected employee values: %+v", emp)
	}
}

func TestToJSON(t *testing.T) {
	p := Product{ID: 1, Name: "Laptop", Price: 999.99}
	js, err := ToJSON(p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var decoded Product
	err = json.Unmarshal([]byte(js), &decoded)
	if err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}
	if decoded.ID != p.ID || decoded.Name != p.Name || decoded.Price != p.Price {
		t.Errorf("decoded struct does not match: %+v", decoded)
	}
}
