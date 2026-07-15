package phase2oopfunctional

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// type-assertions.go demonstrates type assertions and type switches
// Type assertions are used to extract the concrete value from an interface{}
// or to safely check if an interface value holds a specific type

// ---- Basic Type Assertions ----

// BasicTypeAssertion shows the syntax: value.(Type)
// Returns two values: (value, ok) where ok is a boolean indicating success
func BasicTypeAssertion() {
	fmt.Println("\n=== Basic Type Assertions ===")

	var i interface{} = "hello"

	// Type assertion: extract string from interface{}
	s, ok := i.(string)
	if ok {
		fmt.Printf("String value: %q\n", s)
	}

	// Type assertion fails gracefully
	num, ok := i.(int)
	if !ok {
		fmt.Printf("Failed to assert %v as int\n", i)
	}

	fmt.Printf("Extracted int: %d (success: %v)\n", num, ok)
}

// ---- Panic on Failed Assertion ----

// PanicOnFailedAssertion demonstrates that assertions without checking panic
func PanicOnFailedAssertion() {
	fmt.Println("\n=== Panic on Failed Assertion ===")

	var i interface{} = 42

	// This will panic because i is an int, not a string
	// DON'T do this without checking:
	// s := i.(string)  // PANIC!

	// Instead, check first:
	if s, ok := i.(string); ok {
		fmt.Printf("Got string: %s\n", s)
	} else {
		fmt.Printf("Value is not a string, it's %T\n", i)
	}
}

// ---- Type Switches ----

// Describe uses a type switch to handle different types
func Describe(value interface{}) string {
	switch v := value.(type) {
	case nil:
		return "nil"
	case bool:
		return fmt.Sprintf("bool: %v", v)
	case int:
		return fmt.Sprintf("int: %d", v)
	case int64:
		return fmt.Sprintf("int64: %d", v)
	case float64:
		return fmt.Sprintf("float64: %f", v)
	case string:
		return fmt.Sprintf("string: %q", v)
	case []int:
		return fmt.Sprintf("[]int with %d elements", len(v))
	case map[string]interface{}:
		return fmt.Sprintf("map with %d keys", len(v))
	default:
		return fmt.Sprintf("unknown type: %T", v)
	}
}

// AssertionTypeSwitchExample shows how to handle multiple types elegantly
func AssertionTypeSwitchExample() {
	fmt.Println("\n=== Type Switch Pattern ===")

	values := []interface{}{
		nil,
		true,
		42,
		int64(9223372036854775807),
		3.14159,
		"hello world",
		[]int{1, 2, 3, 4, 5},
		map[string]interface{}{"name": "Alice", "age": 30},
	}

	for _, v := range values {
		fmt.Println(Describe(v))
	}
}

// ---- Type Assertions with Interfaces ----

// Reader is an interface for reading
type Reader interface {
	Read() string
}

// SimpleStringReader reads from a string
type SimpleStringReader struct {
	content string
}

func (sr SimpleStringReader) Read() string {
	return sr.content
}

// FileReader reads from a file (simulated)
type FileReader struct {
	filename string
}

func (fr FileReader) Read() string {
	return fmt.Sprintf("Content of %s", fr.filename)
}

// GetReaderType uses type assertion to identify the concrete type
func GetReaderType(r Reader) string {
	switch r.(type) {
	case SimpleStringReader:
		return "SimpleStringReader"
	case FileReader:
		return "FileReader"
	default:
		return "Unknown Reader"
	}
}

// ExtractIfSimpleStringReader safely extracts a SimpleStringReader
func ExtractIfSimpleStringReader(r Reader) (SimpleStringReader, bool) {
	sr, ok := r.(SimpleStringReader)
	return sr, ok
}

// InterfaceTypeAssertionExample shows type assertions on interfaces
func InterfaceTypeAssertionExample() {
	fmt.Println("\n=== Type Assertions on Interfaces ===")

	readers := []Reader{
		SimpleStringReader{content: "Hello from string"},
		FileReader{filename: "data.txt"},
		SimpleStringReader{content: "Another string"},
	}

	for i, r := range readers {
		fmt.Printf("[%d] Type: %s\n", i, GetReaderType(r))
		fmt.Printf("      Content: %s\n", r.Read())

		// Extract if it's a SimpleStringReader
		if sr, ok := ExtractIfSimpleStringReader(r); ok {
			fmt.Printf("      Confirmed StringReader with: %q\n", sr.content)
		}
		fmt.Println()
	}
}

// ---- Practical Type Assertion Pattern ----

// Logger interface for logging
type Logger interface {
	Log(msg string)
}

// ConsoleLogger logs to console
type ConsoleLogger struct {
	prefix string
}

func (cl ConsoleLogger) Log(msg string) {
	fmt.Printf("[%s] %s\n", cl.prefix, msg)
}

// Handler processes requests with a logger
type Handler struct {
	logger Logger
}

// ProcessRequest shows how to use type assertions to add type-specific behavior
func (h *Handler) ProcessRequest(request string) {
	h.logger.Log(fmt.Sprintf("Processing: %s", request))

	// Type assertion: if logger is ConsoleLogger, customize behavior
	if cl, ok := h.logger.(ConsoleLogger); ok {
		fmt.Printf("  (Using ConsoleLogger with prefix: %q)\n", cl.prefix)
	}

	// Could have different behavior for different logger types
	h.logger.Log("Request complete")
}

// PracticalTypeAssertionExample shows real-world usage
func PracticalTypeAssertionExample() {
	fmt.Println("\n=== Practical Type Assertion ===")

	handler := &Handler{
		logger: ConsoleLogger{prefix: "APP"},
	}

	handler.ProcessRequest("GET /api/users")
	fmt.Println()
	handler.ProcessRequest("POST /api/users")
}

// ---- Common Pattern: Comma OK Idiom ----

// Process demonstrates the "comma ok" pattern
// This is idiomatic Go for safe type checking
func Process(value interface{}) {
	// Pattern 1: Check and use
	if str, ok := value.(string); ok {
		fmt.Printf("String: %s\n", str)
		return
	}

	// Pattern 2: Check and use with different variable
	if num, ok := value.(int); ok {
		fmt.Printf("Int: %d\n", num)
		return
	}

	// Pattern 3: Default case
	fmt.Printf("Other type: %T\n", value)
}

// CommaOKExample shows the standard Go pattern for type checking
func CommaOKExample() {
	fmt.Println("\n=== Comma OK Idiom ===")

	Process("hello")
	Process(42)
	Process(3.14)
	Process([]int{1, 2, 3})
}

// ---- Type Assertion in Collections ----

// MixedCollection holds different types
type MixedCollection []interface{}

// CountTypes counts occurrences of each type
func (mc MixedCollection) CountTypes() map[string]int {
	counts := make(map[string]int)

	for _, item := range mc {
		switch item.(type) {
		case string:
			counts["string"]++
		case int:
			counts["int"]++
		case float64:
			counts["float64"]++
		case bool:
			counts["bool"]++
		default:
			counts["other"]++
		}
	}

	return counts
}

// GetStrings extracts all string values
func (mc MixedCollection) GetStrings() []string {
	var strings []string

	for _, item := range mc {
		if str, ok := item.(string); ok {
			strings = append(strings, str)
		}
	}

	return strings
}

// CollectionTypeAssertionExample shows type assertions in collections
func CollectionTypeAssertionExample() {
	fmt.Println("\n=== Type Assertions in Collections ===")

	collection := MixedCollection{
		"hello",
		42,
		3.14,
		"world",
		true,
		99,
		"test",
	}

	counts := collection.CountTypes()
	fmt.Println("Type counts:")
	for typ, count := range counts {
		fmt.Printf("  %s: %d\n", typ, count)
	}

	fmt.Println("\nString values:")
	for _, s := range collection.GetStrings() {
		fmt.Printf("  %q\n", s)
	}
}

// ---- Assertion Chain Pattern ----

// AssertReader tries to assert value as Reader, then specific types
func AssertReader(value interface{}) {
	// First assert it's a Reader
	reader, ok := value.(Reader)
	if !ok {
		fmt.Printf("%v is not a Reader\n", value)
		return
	}

	fmt.Printf("Value is a Reader: %s\n", reader.Read())

	// Then check specific type
	if sr, ok := value.(SimpleStringReader); ok {
		fmt.Printf("  Specifically a SimpleStringReader: %q\n", sr.content)
	} else if fr, ok := value.(FileReader); ok {
		fmt.Printf("  Specifically a FileReader: %s\n", fr.filename)
	}
}

// AssertionChainExample shows chaining assertions
func AssertionChainExample() {
	fmt.Println("\n=== Assertion Chain ===")

	sr := SimpleStringReader{content: "test content"}
	fr := FileReader{filename: "important.txt"}
	var notReader interface{} = 42

	AssertReader(sr)
	fmt.Println()
	AssertReader(fr)
	fmt.Println()
	AssertReader(notReader)
}

// ---- io.Reader Example (Standard Library) ----

// DemonstrationOfIOReader shows type assertions with standard library interfaces
func DemonstrationOfIOReader() {
	fmt.Println("\n=== Standard Library: io.Reader ===")

	// Create different readers
	strReader := strings.NewReader("Hello from strings")
	fileReader, _ := os.Open("type-assertions.go") // May not exist, but that's OK for demo

	// Use them as io.Reader interface
	readers := []io.Reader{strReader}
	if fileReader != nil {
		readers = append(readers, fileReader)
		defer fileReader.Close()
	}

	for i, r := range readers {
		// Type assertion to identify concrete type
		switch r.(type) {
		case *strings.Reader:
			fmt.Printf("[%d] This is a *strings.Reader\n", i)
		case *os.File:
			fmt.Printf("[%d] This is a *os.File\n", i)
		default:
			fmt.Printf("[%d] Unknown reader type: %T\n", i, r)
		}
	}
}

// TypeAssertionPlayground runs all type assertion examples
func TypeAssertionPlayground() {
	fmt.Println("\n========== TYPE ASSERTIONS & TYPE SWITCHES ==========")
	BasicTypeAssertion()
	PanicOnFailedAssertion()
	AssertionTypeSwitchExample()
	InterfaceTypeAssertionExample()
	PracticalTypeAssertionExample()
	CommaOKExample()
	CollectionTypeAssertionExample()
	AssertionChainExample()
	DemonstrationOfIOReader()
}
