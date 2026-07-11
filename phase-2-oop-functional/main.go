package phase2oopfunctional

import (
	"fmt"
)

// MainEntryPoint serves as a central location to run all Phase 2 examples
// Use this to explore all the concepts covered in Phase 2

func RunPhase2Examples() {
	fmt.Println("\n" + string(make([]byte, 80)))
	fmt.Println("=" + string(make([]byte, 78)) + "=")
	fmt.Println("|" + centerText("PHASE 2: OBJECT-ORIENTED & FUNCTIONAL CONCEPTS", 78) + "|")
	fmt.Println("|" + centerText("Comprehensive Go Learning Examples", 78) + "|")
	fmt.Println("=" + string(make([]byte, 78)) + "=")

	// Phase 2.1: Structs
	fmt.Println("\n" + repeatChar("=", 80))
	fmt.Println("PHASE 2.1: STRUCTS & COMPOSITION")
	fmt.Println(repeatChar("=", 80))
	fmt.Println("\nRun: go run basics.go (in 01-structs/)")

	// Phase 2.2: Methods
	fmt.Println("\n" + repeatChar("=", 80))
	fmt.Println("PHASE 2.2: METHODS & RECEIVERS")
	fmt.Println(repeatChar("=", 80))
	fmt.Println("\nKey Concepts:")
	fmt.Println("  - Value receivers: func (p Person) Method()")
	fmt.Println("  - Pointer receivers: func (p *Person) Method()")
	fmt.Println("  - Method chaining for fluent interfaces")
	fmt.Println("\nRun: go run value-receivers.go, pointer-receivers.go (in 02-methods/)")

	// Phase 2.3: Interfaces
	fmt.Println("\n" + repeatChar("=", 80))
	fmt.Println("PHASE 2.3: INTERFACES & DUCK TYPING")
	fmt.Println(repeatChar("=", 80))

	fmt.Println("\n1. BASIC INTERFACES")
	fmt.Println("   Implicit implementation - no 'implements' keyword")
	BasicInterfaceUsage()

	fmt.Println("\n2. DUCK TYPING")
	fmt.Println("   'If it walks and quacks like a duck...'")
	printSeparator()

	fmt.Println("\n3. EMPTY INTERFACE (interface{})")
	fmt.Println("   Accepts any type - use with type assertions")
	fmt.Println("   Example: ProcessValue(\"hello\") ->", ProcessValue("hello"))
	fmt.Println("   Example: ProcessValue(42) ->", ProcessValue(42))

	fmt.Println("\n4. TYPE ASSERTIONS")
	fmt.Println("   Safe extraction with comma-ok: value, ok := i.(string)")
	printSeparator()

	fmt.Println("\n5. STANDARD LIBRARY INTERFACES")
	fmt.Println("   io.Reader, io.Writer, fmt.Stringer, io.Closer")
	fmt.Println("   Example Stringer:")
	person := Person{"Alice", 30}
	fmt.Println("   ", person)

	// Phase 2.4: Packages
	fmt.Println("\n" + repeatChar("=", 80))
	fmt.Println("PHASE 2.4: PACKAGES & STANDARD LIBRARY")
	fmt.Println(repeatChar("=", 80))

	fmt.Println("\n1. FMT PACKAGE - Formatting and Printing")
	fmt.Println("   - Print, Println, Printf")
	fmt.Printf("   - Format verbs: %%v, %%T, %%d, %%s, %%f, etc.\n")
	fmt.Printf("   Example: %s is %d years old\\n", "Bob", 25)

	fmt.Println("\n2. STRINGS PACKAGE - String Manipulation")
	fmt.Println("   - ToUpper, ToLower, Split, Join")
	fmt.Println("   - Contains, HasPrefix, HasSuffix, Replace")
	fmt.Println("   - TrimSpace, Fields")

	fmt.Println("\n3. STRCONV PACKAGE - Type Conversions")
	fmt.Println("   - Atoi, Itoa (int <-> string)")
	fmt.Println("   - ParseInt, ParseFloat, ParseBool")
	fmt.Println("   - FormatInt, FormatFloat, FormatBool")

	fmt.Println("\n4. TIME PACKAGE - Time and Dates")
	fmt.Println("   - time.Now() for current time")
	fmt.Println("   - Parse and Format for string conversions")
	fmt.Println("   - Add, Sub for arithmetic")
	fmt.Println("   - Sleep for delays")

	// Phase 2.5: Modules
	fmt.Println("\n" + repeatChar("=", 80))
	fmt.Println("PHASE 2.5: MODULES & DEPENDENCY MANAGEMENT")
	fmt.Println(repeatChar("=", 80))

	fmt.Println("\nKey Concepts:")
	fmt.Println("  - go.mod: Defines your module and dependencies")
	fmt.Println("  - go.sum: Checksums for reproducibility")
	fmt.Println("  - Semantic Versioning: vMAJOR.MINOR.PATCH")
	fmt.Println("\nCommon Commands:")
	fmt.Println("  - go mod init github.com/user/project")
	fmt.Println("  - go get github.com/package/name")
	fmt.Println("  - go get -u (update all dependencies)")
	fmt.Println("  - go mod tidy (clean up dependencies)")
	fmt.Println("  - go list -m all (list all dependencies)")

	// Summary
	fmt.Println("\n" + repeatChar("=", 80))
	fmt.Println("PHASE 2 COMPLETION SUMMARY")
	fmt.Println(repeatChar("=", 80))

	summary := []struct {
		topic      string
		key        string
		status     string
	}{
		{"Structs", "01-structs", "✅"},
		{"Methods", "02-methods", "✅"},
		{"Interfaces", "03-interfaces", "✅"},
		{"Packages", "04-packages", "✅"},
		{"Modules", "05-modules-dependency", "✅"},
	}

	for _, s := range summary {
		fmt.Printf("  %s %-25s %s\n", s.status, s.topic, "("+s.key+")")
	}

	fmt.Println("\n" + repeatChar("=", 80))
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println(repeatChar("=", 80))

	takeaways := []string{
		"✓ Go uses composition over inheritance",
		"✓ Interfaces are implicit and promote loose coupling",
		"✓ Duck typing enables flexible design",
		"✓ The standard library is comprehensive and well-designed",
		"✓ Go Modules ensure reproducible, manageable projects",
		"✓ Methods add behavior to types",
		"✓ Packages organize code effectively",
	}

	for _, t := range takeaways {
		fmt.Println("  " + t)
	}

	fmt.Println("\n" + repeatChar("=", 80))
	fmt.Println("WHAT'S NEXT?")
	fmt.Println(repeatChar("=", 80))
	fmt.Println("\nPhase 3: Advanced Language Features")
	fmt.Println("  - Concurrency (goroutines, channels)")
	fmt.Println("  - Reflection")
	fmt.Println("  - Generics (Go 1.18+)")
	fmt.Println("  - Advanced error handling")

	fmt.Println("\n" + repeatChar("=", 80))
	fmt.Println("YOU'RE READY FOR ADVANCED TOPICS! 🚀")
	fmt.Println(repeatChar("=", 80) + "\n")
}

// Helper functions
func centerText(text string, width int) string {
	padding := (width - len(text)) / 2
	return repeatChar(" ", padding) + text + repeatChar(" ", width-len(text)-padding)
}

func repeatChar(char string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += char
	}
	return result
}

func printSeparator() {
	fmt.Println(repeatChar("-", 80))
}

// Stub definitions to allow compilation without crossing package/directory boundaries
type Person struct {
	Name string
	Age  int
}

func BasicInterfaceUsage() {
	fmt.Println("   [rex says woof!]")
	fmt.Println("   [whiskers says meow!]")
}

func ProcessValue(value interface{}) string {
	switch v := value.(type) {
	case string:
		return fmt.Sprintf("string: %q", v)
	case int:
		return fmt.Sprintf("int: %d", v)
	default:
		return fmt.Sprintf("unknown: %T", v)
	}
}

