package main

import "fmt"

// empty-interface.go demonstrates interface{}, Go's way of representing any type
// interface{} (empty interface) has zero methods, so every type satisfies it
// It's useful for generic functions but requires type assertions to be useful

// ---- Basic Empty Interface Usage ----

// LogValue logs any value to the console
// Using interface{} means it can accept any type
func LogValue(label string, value interface{}) {
	fmt.Printf("%s: %v (type: %T)\n", label, value, value)
}

// PrintAnything shows that interface{} can hold anything
func PrintAnything(things ...interface{}) {
	fmt.Println("Printing anything:")
	for i, thing := range things {
		fmt.Printf("  [%d] %v (type: %T)\n", i, thing, thing)
	}
}

// BasicEmptyInterfaceExample demonstrates basic usage of interface{}
func BasicEmptyInterfaceExample() {
	fmt.Println("\n=== Basic Empty Interface Usage ===")

	LogValue("String", "hello")
	LogValue("Integer", 42)
	LogValue("Float", 3.14)
	LogValue("Boolean", true)
	LogValue("Slice", []int{1, 2, 3})
	LogValue("Map", map[string]int{"a": 1, "b": 2})

	fmt.Println()
	PrintAnything("text", 123, 45.67, true, []string{"a", "b"})
}

// ---- Type Assertions with Empty Interface ----

// GetType returns a description of what type the value is
func GetType(value interface{}) string {
	// Type assertion: v.(Type) - returns (value, ok) or panics
	switch v := value.(type) {
	case string:
		return fmt.Sprintf("String: %q", v)
	case int:
		return fmt.Sprintf("Integer: %d", v)
	case float64:
		return fmt.Sprintf("Float: %f", v)
	case bool:
		return fmt.Sprintf("Boolean: %v", v)
	case []int:
		return fmt.Sprintf("Int Slice: %v", v)
	case map[string]string:
		return fmt.Sprintf("String Map: %v", v)
	default:
		return fmt.Sprintf("Unknown type: %T", v)
	}
}

// TypeAssertionExample shows how to work with interface{} values
func TypeAssertionExample() {
	fmt.Println("\n=== Type Assertions with Empty Interface ===")

	values := []interface{}{
		"Hello",
		42,
		3.14,
		true,
		[]int{1, 2, 3},
		map[string]string{"key": "value"},
	}

	for _, val := range values {
		fmt.Println(GetType(val))
	}
}

// ---- Converting interface{} values ----

// SumNumbers attempts to sum a slice of interface{} values
// This shows how to safely extract and use values from interface{}
func SumNumbers(numbers ...interface{}) (int64, error) {
	var sum int64

	for i, num := range numbers {
		// Type assertion with error checking
		intVal, ok := num.(int)
		if !ok {
			return 0, fmt.Errorf("value at index %d is not an int: %T", i, num)
		}
		sum += int64(intVal)
	}

	return sum, nil
}

// StringifyValues converts interface{} values to strings
func StringifyValues(values ...interface{}) []string {
	result := make([]string, len(values))
	for i, v := range values {
		// Direct type assertion without checking - will panic if wrong type
		// In production, use ok check like: v.(string)
		result[i] = fmt.Sprintf("%v", v)
	}
	return result
}

// ConversionExample shows converting interface{} values
func ConversionExample() {
	fmt.Println("\n=== Converting interface{} Values ===")

	// Example 1: Sum numbers
	fmt.Println("Summing numbers:")
	sum, err := SumNumbers(10, 20, 30, 40)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Sum: %d\n", sum)
	}

	// Example 2: Sum with error
	fmt.Println("\nTrying to sum with non-integer:")
	sum, err = SumNumbers(10, 20, "thirty", 40)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Example 3: Stringify
	fmt.Println("\nStringifying values:")
	stringified := StringifyValues("hello", 123, 45.6, true)
	fmt.Println("Stringified:", stringified)
}

// ---- Type Switch Pattern ----

// ProcessValue demonstrates the type switch pattern for interface{}
// This is the idiomatic Go way to handle multiple possible types
func ProcessValue(value interface{}) string {
	switch v := value.(type) {
	case nil:
		return "value is nil"
	case bool:
		return fmt.Sprintf("boolean: %v", v)
	case int:
		return fmt.Sprintf("integer: %d", v)
	case float64:
		return fmt.Sprintf("float: %f", v)
	case string:
		return fmt.Sprintf("string: %q", v)
	case []interface{}:
		return fmt.Sprintf("slice with %d elements", len(v))
	case map[string]interface{}:
		return fmt.Sprintf("map with %d keys", len(v))
	default:
		return fmt.Sprintf("unknown type: %T", v)
	}
}

// TypeSwitchExample shows the idiomatic pattern for handling interface{}
func TypeSwitchExample() {
	fmt.Println("\n=== Type Switch Pattern ===")

	testValues := []interface{}{
		nil,
		true,
		42,
		3.14,
		"hello",
		[]interface{}{1, 2, 3},
		map[string]interface{}{"a": 1, "b": 2},
	}

	for _, val := range testValues {
		fmt.Println(ProcessValue(val))
	}
}

// ---- When NOT to use interface{} ----

// AVOID: Generic function that's too vague
func BadGenericFunction(data interface{}, operation interface{}) interface{} {
	// What do data and operation actually contain?
	// This is confusing and error-prone
	return nil
}

// GOOD: Use specific interface or type parameter
type Processor interface {
	Process(input string) string
}

func GoodGenericFunction(data string, processor Processor) string {
	return processor.Process(data)
}

// WarningsAndBestPractices demonstrates when to use (and not use) interface{}
func WarningsAndBestPractices() {
	fmt.Println("\n=== When to Use interface{} ===")

	fmt.Println(`
GOOD uses of interface{}:
  1. Functions that truly accept any type (logging, printing)
  2. Data structures that need to store mixed types (slices, maps)
  3. APIs where you can't know the type ahead of time (JSON unmarshaling)
  4. Transitioning from untyped code

AVOID using interface{} for:
  1. Function parameters with specific operations needed
  2. Return types when specific type is known
  3. When you immediately need to do type assertions
  4. Creating confusion about what types are actually allowed

BETTER alternatives:
  1. Use specific types or interfaces
  2. Use type parameters/generics (Go 1.18+)
  3. Use concrete types and polymorphism via interfaces
	`)
}

// ---- Practical Example: A Generic Container ----

// Container is a generic container using interface{}
// It can store any type but retrieves as interface{}
type Container struct {
	data map[string]interface{}
}

func NewContainer() *Container {
	return &Container{data: make(map[string]interface{})}
}

func (c *Container) Set(key string, value interface{}) {
	c.data[key] = value
}

func (c *Container) Get(key string) (interface{}, bool) {
	val, ok := c.data[key]
	return val, ok
}

func (c *Container) GetString(key string) (string, error) {
	val, ok := c.Get(key)
	if !ok {
		return "", fmt.Errorf("key not found: %s", key)
	}
	str, ok := val.(string)
	if !ok {
		return "", fmt.Errorf("value for key %q is not a string: %T", key, val)
	}
	return str, nil
}

func (c *Container) GetInt(key string) (int, error) {
	val, ok := c.Get(key)
	if !ok {
		return 0, fmt.Errorf("key not found: %s", key)
	}
	intVal, ok := val.(int)
	if !ok {
		return 0, fmt.Errorf("value for key %q is not an int: %T", key, val)
	}
	return intVal, nil
}

// PracticalContainerExample shows a real use case for interface{}
func PracticalContainerExample() {
	fmt.Println("\n=== Practical: Generic Container ===")

	container := NewContainer()
	container.Set("name", "Alice")
	container.Set("age", 30)
	container.Set("email", "alice@example.com")

	// Retrieving with type assertions
	if name, err := container.GetString("name"); err == nil {
		fmt.Printf("Name: %s\n", name)
	}

	if age, err := container.GetInt("age"); err == nil {
		fmt.Printf("Age: %d\n", age)
	}

	// Trying to get with wrong type
	if _, err := container.GetInt("name"); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

// EmptyInterfacePlayground runs all empty interface examples
func main() {
	fmt.Println("\n========== EMPTY INTERFACE (interface{}) EXAMPLES ==========")
	BasicEmptyInterfaceExample()
	TypeAssertionExample()
	ConversionExample()
	TypeSwitchExample()
	WarningsAndBestPractices()
	PracticalContainerExample()
}
