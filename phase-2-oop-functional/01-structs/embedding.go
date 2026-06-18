package main

import "fmt"

// Embedding demonstrates Go's approach to composition.
// Go doesn't have inheritance, but uses embedding (struct composition)
// to build complex types from simpler ones.

// Address represents a physical address
type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

// Employee demonstrates embedding an Address struct
// This is Go's way of "inheriting" fields from Address
type Employee struct {
	ID        int
	Name      string
	Address   Address // Embedded struct (not anonymous)
	Department string
}

// Student demonstrates anonymous struct embedding
// When you embed a struct without a field name, it's called anonymous embedding
// The fields of the embedded struct become directly accessible
type Student struct {
	ID    int
	Name  string
	Address // Anonymous embedding - fields become accessible directly
	GPA   float64
}

// Company demonstrates multiple embedding levels
type Company struct {
	Name string
}

type Manager struct {
	Employee // Embedding Employee
	TeamSize int
	Company  Company // Embedding another struct
}

// EmbeddingBasics shows the difference between named and anonymous embedding
func EmbeddingBasics() {
	fmt.Println("=== Embedding Basics ===")

	// Named embedding - you access nested struct through field name
	emp := Employee{
		ID:   1,
		Name: "Alice",
		Address: Address{
			Street: "123 Main St",
			City:   "Boston",
			State:  "MA",
			Zip:    "02101",
		},
		Department: "Engineering",
	}

	fmt.Printf("Employee: %+v\n", emp)
	fmt.Printf("Address via named field: %s, %s %s\n",
		emp.Address.Street, emp.Address.City, emp.Address.State)
}

// AnonymousEmbedding shows the power of anonymous embedding
func AnonymousEmbedding() {
	fmt.Println("\n=== Anonymous Embedding ===")

	// With anonymous embedding, Address fields are promoted to Student
	student := Student{
		ID:   1001,
		Name: "Bob",
		Address: Address{
			Street: "456 Oak Ave",
			City:   "Cambridge",
			State:  "MA",
			Zip:    "02138",
		},
		GPA: 3.8,
	}

	fmt.Printf("Student: %+v\n", student)

	// With anonymous embedding, you can access Address fields directly!
	fmt.Printf("Street (direct): %s\n", student.Street)
	fmt.Printf("City (direct): %s\n", student.City)

	// You can still access through the struct name
	fmt.Printf("Street (via Address): %s\n", student.Address.Street)
}

// PromotedFields shows how embedded struct fields are promoted
func PromotedFields() {
	fmt.Println("\n=== Promoted Fields ===")

	student := Student{
		ID:   1002,
		Name: "Charlie",
		Address: Address{
			Street: "789 Pine Rd",
			City:   "Harvard",
			State:  "MA",
			Zip:    "02134",
		},
		GPA: 3.9,
	}

	// Access promoted fields directly (as if they belong to Student)
	fmt.Printf("ID: %d\n", student.ID)
	fmt.Printf("Name: %s\n", student.Name)
	fmt.Printf("Street: %s (promoted field from Address)\n", student.Street)
	fmt.Printf("City: %s (promoted field from Address)\n", student.City)
	fmt.Printf("GPA: %.2f\n", student.GPA)

	// This is very convenient for composition!
}

// MethodPromotion shows that embedded struct methods are also promoted
type Vehicle struct {
	Brand string
	Model string
}

// Start is a method on Vehicle
func (v Vehicle) Start() string {
	return fmt.Sprintf("%s %s is starting...", v.Brand, v.Model)
}

// Stop is another method on Vehicle
func (v Vehicle) Stop() string {
	return fmt.Sprintf("%s %s has stopped", v.Brand, v.Model)
}

// Car embeds Vehicle, so it gets Start() and Stop() methods automatically
type Car struct {
	Vehicle // Embedded struct
	Doors   int
	Trunk   bool
}

func MethodPromotion() {
	fmt.Println("\n=== Method Promotion ===")

	car := Car{
		Vehicle: Vehicle{
			Brand: "Tesla",
			Model: "Model 3",
		},
		Doors: 4,
		Trunk: true,
	}

	// Car doesn't define Start() or Stop(), but inherits them from Vehicle
	fmt.Println(car.Start())  // Promoted method
	fmt.Println(car.Stop())   // Promoted method

	// This is a form of method inheritance through embedding
}

// ShadowingFields shows what happens when parent and child have same field name
type Animal struct {
	Name string
}

type Dog struct {
	Animal
	Name string // This shadows the embedded Animal.Name field
}

func ShadowingFields() {
	fmt.Println("\n=== Field Shadowing ===")

	dog := Dog{
		Animal: Animal{Name: "Animal Name"},
		Name:   "Dog Name",
	}

	fmt.Printf("dog.Name (Dog's own): %s\n", dog.Name)
	fmt.Printf("dog.Animal.Name (embedded): %s\n", dog.Animal.Name)

	// When there's a conflict, the outer (child) field takes precedence
	// This is called field shadowing or hiding
}

// ComplexEmbedding shows multiple levels of embedding
type Address2 struct {
	Street string
	City   string
}

type Person struct {
	Name    string
	Address2 // Anonymous embedding
}

type Company2 struct {
	Name     string
	CEO      Person // Embedding a struct that embeds another
	Location Address2
}

func ComplexEmbedding() {
	fmt.Println("\n=== Complex Multi-Level Embedding ===")


company := Company2{
		Name: "Tech Corp",
		CEO: Person{
			Name: "Diana",
			Address2: Address2{
				Street: "100 Tech Ave",
				City:   "Silicon Valley",
			},
		},
		Location: Address2{
			Street: "200 Business Blvd",
			City:   "San Francisco",
		},
	}

	fmt.Printf("Company: %+v\n", company)
	fmt.Printf("CEO Name: %s\n", company.CEO.Name)
	fmt.Printf("CEO City (promoted): %s\n", company.CEO.City)
}

// EmbeddingInterface shows embedding interfaces
// An interface can embed other interfaces to combine their methods
type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string) error
}

// ReadWriter embeds both Reader and Writer
// Any type implementing ReadWriter must implement both Read() and Write()
type ReadWriter interface {
	Reader
	Writer
}

type File struct {
	name string
	content string
}

func (f File) Read() string {
	return f.content
}

func (f *File) Write(data string) error {
	f.content = data
	return nil
}

func InterfaceEmbedding() {
	fmt.Println("\n=== Interface Embedding ===")

	file := &File{name: "test.txt", content: "Hello"}

	// file now satisfies ReadWriter interface
	fmt.Printf("Read: %s\n", file.Read())
	file.Write("World")
	fmt.Printf("Write completed. New content: %s\n", file.Read())
}

// BestPractices provides guidelines on using embedding
func BestPractices() {
	fmt.Println("\n=== Best Practices for Embedding ===")

	fmt.Println(`
When to use embedding:
1. Composition makes sense semantically ("is-a" or "has-a")
2. You want to reuse fields and methods
3. The relationship is clear and intuitive

When NOT to use embedding:
1. Just to avoid typing (this is lazy)
2. Multiple fields with same type (confusing)
3. Unclear relationships
4. If field shadowing would be problematic

Guidelines:
- Use anonymous embedding for simple, clear composition
- Use named embedding when the relationship needs to be explicit
- Be careful with field/method name conflicts
- Document complex embedding structures
- Prefer explicit composition if it makes code clearer
`)
}

func main() {
	EmbeddingBasics()
	AnonymousEmbedding()
	PromotedFields()
	MethodPromotion()
	ShadowingFields()
	ComplexEmbedding()
	InterfaceEmbedding()
	BestPractices()

	fmt.Println("\n=== All Embedding Examples Completed ===")
}
