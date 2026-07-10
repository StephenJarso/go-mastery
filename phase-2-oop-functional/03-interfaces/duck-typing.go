package phase2oopfunctional

import "fmt"

// duck-typing.go demonstrates Go's implicit interface implementation (duck typing)
// "If it walks like a duck and quacks like a duck, then it is a duck"
// In Go, you don't explicitly declare that a type implements an interface
// If a type has all the methods of an interface, it automatically satisfies it

// Writer defines the interface for anything that can write
type Writer interface {
	Write(data string) error
}

// ConsoleWriter writes to console
type ConsoleWriter struct {
	Name string
}

func (cw ConsoleWriter) Write(data string) error {
	fmt.Printf("[%s] Writing to console: %s\n", cw.Name, data)
	return nil
}

// FileWriter writes to file (simulated)
type FileWriter struct {
	Filename string
}

func (fw FileWriter) Write(data string) error {
	fmt.Printf("[%s] Writing to file: %s\n", fw.Filename, data)
	return nil
}

// NetworkWriter writes to network (simulated)
type NetworkWriter struct {
	Address string
}

func (nw NetworkWriter) Write(data string) error {
	fmt.Printf("[%s] Writing to network: %s\n", nw.Address, data)
	return nil
}

// Notice: We never explicitly said "ConsoleWriter implements Writer"
// Go automatically recognizes it because it has a Write method with the right signature

// ProcessData takes any Writer and uses it without knowing the concrete type
// This is the power of duck typing: loose coupling
func ProcessData(writer Writer, data string) error {
	fmt.Println("Processing data...")
	return writer.Write(data)
}

// DuckTypingExample shows how different types satisfying the same interface
// can be used interchangeably without explicit declaration
func DuckTypingExample() {
	fmt.Println("\n=== Duck Typing (Implicit Implementation) ===")

	console := ConsoleWriter{Name: "Console1"}
	file := FileWriter{Filename: "output.txt"}
	network := NetworkWriter{Address: "192.168.1.1"}

	// All three types satisfy Writer interface, but there's no explicit "implements" declaration
	writers := []Writer{console, file, network}

	data := "Important information"

	for _, w := range writers {
		ProcessData(w, data)
	}

	fmt.Println("\nKey Point: We never said these types 'implement' Writer")
	fmt.Println("But Go knows they do because they have the Write method!")
}

// ---- Advanced Duck Typing Example ----

// SaverAndLoader is a composed interface for types that can save and load
type Saver interface {
	Save() error
}

type Loader interface {
	Load() error
}

type SaverAndLoader interface {
	Saver
	Loader
}

// Database implements both Saver and Loader (and thus SaverAndLoader)
type Database struct {
	Name string
}

func (db Database) Save() error {
	fmt.Printf("Saving to database: %s\n", db.Name)
	return nil
}

func (db Database) Load() error {
	fmt.Printf("Loading from database: %s\n", db.Name)
	return nil
}

// FileStorage implements both Saver and Loader
type FileStorage struct {
	Path string
}

func (fs FileStorage) Save() error {
	fmt.Printf("Saving to file: %s\n", fs.Path)
	return nil
}

func (fs FileStorage) Load() error {
	fmt.Printf("Loading from file: %s\n", fs.Path)
	return nil
}

// Backup works with any type that implements SaverAndLoader
func Backup(storage SaverAndLoader) error {
	fmt.Println("Starting backup...")
	if err := storage.Load(); err != nil {
		return err
	}
	if err := storage.Save(); err != nil {
		return err
	}
	fmt.Println("Backup complete!")
	return nil
}

// ComposedInterfaceExample shows using multiple interfaces together
func ComposedInterfaceExample() {
	fmt.Println("\n=== Composed Interface Example ===")

	db := Database{Name: "ProductDB"}
	fs := FileStorage{Path: "/backups/storage.bak"}

	// Both types satisfy SaverAndLoader because they implement Save and Load
	Backup(db)
	fmt.Println()
	Backup(fs)
}

// ---- Interface Satisfaction Satisfaction ----

// Calculator interface defines mathematical operations
type Calculator interface {
	Add(a, b int) int
	Subtract(a, b int) int
	Multiply(a, b int) int
}

// SimpleCalc implements Calculator
type SimpleCalc struct{}

func (sc SimpleCalc) Add(a, b int) int {
	return a + b
}

func (sc SimpleCalc) Subtract(a, b int) int {
	return a - b
}

func (sc SimpleCalc) Multiply(a, b int) int {
	return a * b
}

// ScientificCalc also implements Calculator (plus more)
type ScientificCalc struct{}

func (sc ScientificCalc) Add(a, b int) int {
	return a + b
}

func (sc ScientificCalc) Subtract(a, b int) int {
	return a - b
}

func (sc ScientificCalc) Multiply(a, b int) int {
	return a * b
}

// ScientificCalc has extra methods not in the interface
func (sc ScientificCalc) Power(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

// PerformCalculation works with any Calculator
// The fact that ScientificCalc has Power() doesn't matter here
func PerformCalculation(calc Calculator, a, b int) {
	fmt.Printf("Add: %d, Subtract: %d, Multiply: %d\n",
		calc.Add(a, b),
		calc.Subtract(a, b),
		calc.Multiply(a, b),
	)
}

// DuckTypingWithExtendedInterface shows that types can have more methods
// than required by the interface they satisfy
func DuckTypingWithExtendedInterface() {
	fmt.Println("\n=== Duck Typing with Extended Interface ===")

	simple := SimpleCalc{}
	scientific := ScientificCalc{}

	fmt.Println("Simple Calculator:")
	PerformCalculation(simple, 10, 5)

	fmt.Println("\nScientific Calculator (when used as Calculator):")
	PerformCalculation(scientific, 10, 5)

	fmt.Println("\nBut ScientificCalc has extra methods!")
	fmt.Printf("Power: 2^3 = %d\n", scientific.Power(2, 3))
}

// DuckTypingPlayground demonstrates the flexibility of duck typing
func DuckTypingPlayground() {
	fmt.Println("\n========== DUCK TYPING EXAMPLES ==========")
	DuckTypingExample()
	ComposedInterfaceExample()
	DuckTypingWithExtendedInterface()
}
