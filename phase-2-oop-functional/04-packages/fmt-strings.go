package phase2packages

import (
	"fmt"
	"strings"
)

// fmt-strings.go demonstrates the fmt and strings packages
// These are fundamental Go packages used extensively in programs

// ===== FMT PACKAGE =====

// fmtBasics demonstrates basic fmt functionality
func FmtBasics() {
	fmt.Println("\n=== FMT Package Basics ===")

	// Print - outputs to stdout with spaces between args
	fmt.Print("Hello", " ", "World\n")

	// Println - adds newline
	fmt.Println("Line 1")
	fmt.Println("Line 2")

	// Printf - formatted printing
	name := "Alice"
	age := 30
	fmt.Printf("%s is %d years old\n", name, age)

	// Sprintf - format to string
	message := fmt.Sprintf("%s will be %d next year", name, age+1)
	fmt.Println("Message:", message)
}

// fmtFormatVerbs demonstrates formatting verbs
func FmtFormatVerbs() {
	fmt.Println("\n=== Format Verbs ===")

	// General verbs
	fmt.Printf("%%v (value): %v\n", 42)        // value in default format
	fmt.Printf("%%T (type): %T\n", 42)         // type of value
	fmt.Printf("%%#v (repr): %#v\n", "test")   // Go syntax representation

	// String verbs
	str := "hello"
	fmt.Printf("%%s (string): %s\n", str)
	fmt.Printf("%%q (quoted): %q\n", str)
	fmt.Printf("%%x (hex): %x\n", str)

	// Integer verbs
	num := 255
	fmt.Printf("%%d (decimal): %d\n", num)
	fmt.Printf("%%o (octal): %o\n", num)
	fmt.Printf("%%x (hex): %x\n", num)
	fmt.Printf("%%b (binary): %b\n", num)

	// Float verbs
	float := 3.14159265
	fmt.Printf("%%f (float): %f\n", float)
	fmt.Printf("%%.2f (2 decimals): %.2f\n", float)
	fmt.Printf("%%e (scientific): %e\n", float)
	fmt.Printf("%%g (general): %g\n", float)
}

// fmtWidth demonstrates width and padding
func FmtWidth() {
	fmt.Println("\n=== Width and Padding ===")

	values := []interface{}{"hello", 42, 3.14}

	// Left align (default right)
	fmt.Printf("|%5v|\n", values[0])  // right aligned
	fmt.Printf("|%-5v|\n", values[0]) // left aligned

	// Zero padding for numbers
	fmt.Printf("|%05d|\n", 42)
	fmt.Printf("|%+d|\n", 42)  // show sign
	fmt.Printf("|% d|\n", 42)  // space for positive
}

// ===== STRINGS PACKAGE =====

// stringsBasics demonstrates basic strings operations
func StringsBasics() {
	fmt.Println("\n=== Strings Package Basics ===")

	text := "Hello, World!"

	// Length
	fmt.Printf("Length of '%s': %d\n", text, len(text))

	// Case conversion
	fmt.Println("Upper:", strings.ToUpper(text))
	fmt.Println("Lower:", strings.ToLower(text))
	fmt.Println("Title:", strings.Title(text))

	// Contains and finding
	fmt.Println("Contains 'World':", strings.Contains(text, "World"))
	fmt.Println("Contains 'xyz':", strings.Contains(text, "xyz"))
	fmt.Println("Index of 'World':", strings.Index(text, "World"))
	fmt.Println("Index of 'xyz':", strings.Index(text, "xyz"))
}

// stringsTrim demonstrates trimming operations
func StringsTrim() {
	fmt.Println("\n=== Trimming Operations ===")

	text := "  hello world  "

	// TrimSpace - remove leading/trailing whitespace
	fmt.Printf("Original: '%s'\n", text)
	fmt.Printf("Trimmed: '%s'\n", strings.TrimSpace(text))

	// Trim specific characters
	text2 := "xxxHello Worldxxx"
	fmt.Printf("Trim 'x': '%s'\n", strings.Trim(text2, "x"))
	fmt.Printf("Trim prefix: '%s'\n", strings.TrimPrefix(text2, "xxx"))
	fmt.Printf("Trim suffix: '%s'\n", strings.TrimSuffix(text2, "xxx"))
}

// stringsReplace demonstrates replacement operations
func StringsReplace() {
	fmt.Println("\n=== Replacement Operations ===")

	text := "hello hello hello"

	// Replace first n occurrences
	fmt.Println("Replace first 1:", strings.Replace(text, "hello", "hi", 1))
	fmt.Println("Replace first 2:", strings.Replace(text, "hello", "hi", 2))
	fmt.Println("Replace all (-1):", strings.Replace(text, "hello", "hi", -1))

	// ReplaceAll replaces all occurrences
	fmt.Println("ReplaceAll:", strings.ReplaceAll(text, "hello", "hey"))
}

// stringsSplit demonstrates splitting and joining
func StringsSplit() {
	fmt.Println("\n=== Split and Join ===")

	// Split
	csv := "apple,banana,cherry,date"
	parts := strings.Split(csv, ",")
	fmt.Println("Split:", parts)

	// Fields - split on whitespace
	text := "one  two   three    four"
	fields := strings.Fields(text)
	fmt.Println("Fields:", fields)

	// Join
	joined := strings.Join(parts, " | ")
	fmt.Println("Joined:", joined)
}

// stringsPrefix demonstrates prefix/suffix operations
func StringsPrefix() {
	fmt.Println("\n=== Prefix and Suffix ===")

	text := "golang.org"

	fmt.Println("HasPrefix 'go':", strings.HasPrefix(text, "go"))
	fmt.Println("HasPrefix 'rust':", strings.HasPrefix(text, "rust"))
	fmt.Println("HasSuffix '.org':", strings.HasSuffix(text, ".org"))
	fmt.Println("HasSuffix '.com':", strings.HasSuffix(text, ".com"))
}

// stringsRepeat demonstrates repeat operation
func StringsRepeat() {
	fmt.Println("\n=== Repeat ===")

	fmt.Println(strings.Repeat("ab", 3))
	fmt.Println(strings.Repeat("*", 10))
}

// PackageFmtStringsPlayground runs all package examples
func PackageFmtStringsPlayground() {
	fmt.Println("\n========== FMT & STRINGS PACKAGES ==========")
	FmtBasics()
	FmtFormatVerbs()
	FmtWidth()
	StringsBasics()
	StringsTrim()
	StringsReplace()
	StringsSplit()
	StringsPrefix()
	StringsRepeat()
}
