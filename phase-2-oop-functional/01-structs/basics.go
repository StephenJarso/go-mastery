package main

import (
	"fmt"
	"unsafe"
)

// Person is a struct that represents a person.
// It groups related data together.
type Person struct {
	// Name is the person's name
	Name string
	// Age is the person's age
	Age int
	// Email is the person's email address
	Email string
}

// Company is another struct to demonstrate multiple structs
type Company struct {
	Name     string
	Location string
	Employees int
}

// StructBasics demonstrates the basics of struct creation and usage
func StructBasics() {
	fmt.Println("=== Struct Basics ===")

	// Method 1: Create struct using named fields
	// This is the recommended way as it's clear and maintainable
	alice := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}
	fmt.Printf("Person 1: %+v\n", alice)

	// Method 2: Create struct using positional arguments
	// Not recommended for larger structs as it's unclear
	bob := Person{"Bob", 25, "bob@example.com"}
	fmt.Printf("Person 2: %+v\n", bob)

	// Method 3: Using new() function
	// Returns a pointer to zero-initialized struct
	charlie := new(Person)
	charlie.Name = "Charlie"
	charlie.Age = 35
	charlie.Email = "charlie@example.com"
	fmt.Printf("Person 3: %+v\n", charlie)

	// Method 4: Create and immediately assign
	diana := Person{Name: "Diana"} // Age and Email are zero values
	fmt.Printf("Person 4 (partial): %+v\n", diana)
}

// ZeroValues demonstrates Go's zero values concept
func ZeroValues() {
	fmt.Println("\n=== Zero Values ===")

	// When you declare a struct without initializing fields,
	// Go assigns zero values to each field:
	// - string: empty string ""
	// - int: 0
	// - bool: false
	// - pointers: nil
	var uninitialized Person
	fmt.Printf("Uninitialized struct: %+v\n", uninitialized)
	fmt.Printf("Name field: %q\n", uninitialized.Name)
	fmt.Printf("Age field: %d\n", uninitialized.Age)
}

// AccessingFields demonstrates how to access and modify struct fields
func AccessingFields() {
	fmt.Println("\n=== Accessing and Modifying Fields ===")

	person := Person{Name: "Eve", Age: 28, Email: "eve@example.com"}

	// Access fields using dot notation
	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)

	// Modify fields
	person.Age = 29
	fmt.Printf("After birthday: %+v\n", person)

	// Note: You can only modify if you have a non-pointer
	// or a pointer to the struct
}

// PointersToStructs demonstrates working with pointers to structs
func PointersToStructs() {
	fmt.Println("\n=== Pointers to Structs ===")

	original := Person{Name: "Frank", Age: 40, Email: "frank@example.com"}
	fmt.Printf("Original: %+v\n", original)

	// Create a pointer to the struct
	pointer := &original
	fmt.Printf("Pointer: %p\n", pointer)

	// Modify through pointer (Go handles dereferencing automatically for fields)
	pointer.Age = 41
	fmt.Printf("After modification: %+v\n", original)

	// Explicit dereferencing (also works)
	(*pointer).Name = "Frederick"
	fmt.Printf("After name change: %+v\n", original)

	// Important: When you modify a pointer, the original is modified
	// This is different from value receivers
}

// AnonymousStructs demonstrates inline struct definitions (rarely used)
func AnonymousStructs() {
	fmt.Println("\n=== Anonymous Structs ===")

	// You can define structs without a type name
	// This is useful for quick, one-off data structures
	config := struct {
		Host string
		Port int
		SSL  bool
	}{
		Host: "localhost",
		Port: 8080,
		SSL:  false,
	}

	fmt.Printf("Config: %+v\n", config)
	fmt.Printf("Host: %s, Port: %d\n", config.Host, config.Port)
}

// ComparingStructs demonstrates struct comparison
func ComparingStructs() {
	fmt.Println("\n=== Comparing Structs ===")

	person1 := Person{Name: "Grace", Age: 30, Email: "grace@example.com"}
	person2 := Person{Name: "Grace", Age: 30, Email: "grace@example.com"}
	person3 := Person{Name: "Helen", Age: 30, Email: "helen@example.com"}

	// Structs can be compared with == if all fields are comparable
	fmt.Printf("person1 == person2: %v\n", person1 == person2) // true
	fmt.Printf("person1 == person3: %v\n", person1 == person3) // false
}

// CopyingStructs demonstrates value semantics of structs
func CopyingStructs() {
	fmt.Println("\n=== Copying Structs ===")

	original := Person{Name: "Ivan", Age: 50, Email: "ivan@example.com"}

	// Structs are values, so assignment creates a copy
	copy := original
	copy.Name = "Ivan Jr."

	fmt.Printf("Original: %+v\n", original)
	fmt.Printf("Copy: %+v\n", copy)

	// Note: original is unchanged because copy is a separate instance
	// This is different from pointers
}

// StructSize demonstrates that Go tracks memory efficiently
func StructSize() {
	fmt.Println("\n=== Struct Size ===")

	//var p Person
	//fmt.Printf("Size of Person struct: %d bytes\n", unsafe(len([]byte{}))) // Would use unsafe.Sizeof in production
	fmt.Printf("Size of string field: %d bytes\n", 16)                    // strings are 16 bytes
	fmt.Printf("Size of int field: %d bytes\n", 8)                       // int64 is 8 bytes on 64-bit systems
}

// Using unsafe to show actual sizes


func StructSizeWithUnsafe() {
	fmt.Println("\n=== Struct Size (with unsafe) ===")

	var p Person
	fmt.Printf("Size of Person struct: %d bytes\n", unsafe.Sizeof(p))
	fmt.Printf("Size of Name field: %d bytes\n", unsafe.Sizeof(p.Name))
	fmt.Printf("Size of Age field: %d bytes\n", unsafe.Sizeof(p.Age))
	fmt.Printf("Size of Email field: %d bytes\n", unsafe.Sizeof(p.Email))
}
//To understand more about value Receiver and pointer receiver,here is a small code to explain the difference
//value receiver- method gets a copy of the struct
//Pointer receiver- the method gets the actual struct(can modify it)

// 1.Value Receiver
type Person1 struct{
	Name string
	Age int
}
// Her  p is a copy.If you change p.Nmae inside Greet()the original is untouched.
func(p Person1)Greet()string{
return "Hi,I'm "+p.Name
}
func(p Person1) Birthday(){
	p.Age++
}
func(p *Person1) Birthday1(){
	p.Age++
}
func main() {
	// Run all examples
	StructBasics()
	ZeroValues()
	AccessingFields()
	PointersToStructs()
	AnonymousStructs()
	ComparingStructs()
	CopyingStructs()
	StructSizeWithUnsafe()

	fmt.Println("\n=== All Examples Completed ===")

	bob := Person1{
		Name: "bob",
		Age: 25,
	}
	bob.Birthday()
	fmt.Println(bob.Age)// 25 -nothing changed

	bob.Birthday1()
	fmt.Println(bob.Age)// 26 - changed
}


