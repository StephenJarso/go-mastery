package phase2packages

import (
	"fmt"
	"strconv"
)

// strconv-conversions.go demonstrates the strconv package
// Used for converting strings to/from other types

// ===== STRING TO INTEGER =====

// stringToIntegerExample shows converting strings to integers
func StringToIntegerExample() {
	fmt.Println("\n=== String to Integer ===")

	// ParseInt - parse string as integer
	str := "123"
	num, err := strconv.ParseInt(str, 10, 64) // base 10, 64-bit
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Parsed '%s' as int64: %d\n", str, num)
	}

	// Atoi - convenience function for decimal integers (same as ParseInt base 10)
	num2, err := strconv.Atoi("456")
	if err == nil {
		fmt.Printf("Atoi('456'): %d\n", num2)
	}

	// Different bases
	hex, _ := strconv.ParseInt("FF", 16, 64) // hex
	bin, _ := strconv.ParseInt("1010", 2, 64) // binary
	oct, _ := strconv.ParseInt("777", 8, 64) // octal

	fmt.Printf("Hex FF: %d\n", hex)
	fmt.Printf("Binary 1010: %d\n", bin)
	fmt.Printf("Octal 777: %d\n", oct)

	// Error handling
	_, err = strconv.ParseInt("abc", 10, 64)
	if err != nil {
		fmt.Printf("Error parsing 'abc': %v\n", err)
	}
}

// ===== INTEGER TO STRING =====

// integerToStringExample shows converting integers to strings
func IntegerToStringExample() {
	fmt.Println("\n=== Integer to String ===")

	// FormatInt - format integer as string
	num := int64(255)

	str10 := strconv.FormatInt(num, 10) // decimal
	str16 := strconv.FormatInt(num, 16) // hex
	str2 := strconv.FormatInt(num, 2)   // binary
	str8 := strconv.FormatInt(num, 8)   // octal

	fmt.Printf("%d in base 10: %s\n", num, str10)
	fmt.Printf("%d in base 16: %s\n", num, str16)
	fmt.Printf("%d in base 2: %s\n", num, str2)
	fmt.Printf("%d in base 8: %s\n", num, str8)

	// Itoa - convenience function (same as FormatInt base 10)
	str := strconv.Itoa(789)
	fmt.Printf("Itoa(789): %s\n", str)
}

// ===== STRING TO FLOAT =====

// stringToFloatExample shows converting strings to floats
func StringToFloatExample() {
	fmt.Println("\n=== String to Float ===")

	// ParseFloat
	str := "3.14159"
	num, err := strconv.ParseFloat(str, 64) // 64-bit precision
	if err == nil {
		fmt.Printf("ParseFloat('%s'): %f\n", str, num)
	}

	// Scientific notation
	scientific := "1.23e-4"
	num2, _ := strconv.ParseFloat(scientific, 64)
	fmt.Printf("ParseFloat('%s'): %f\n", scientific, num2)

	// Negative numbers
	negative := "-42.5"
	num3, _ := strconv.ParseFloat(negative, 64)
	fmt.Printf("ParseFloat('%s'): %f\n", negative, num3)
}

// ===== FLOAT TO STRING =====

// floatToStringExample shows converting floats to strings
func FloatToStringExample() {
	fmt.Println("\n=== Float to String ===")

	num := 3.14159265

	// FormatFloat - format float as string
	// syntax: FormatFloat(f float64, fmt byte, prec, bitSize int) string
	// fmt: 'f' (decimal), 'e' (scientific), 'g' (general)

	decimal := strconv.FormatFloat(num, 'f', 2, 64) // 2 decimal places
	scientific := strconv.FormatFloat(num, 'e', 4, 64) // scientific notation
	general := strconv.FormatFloat(num, 'g', 4, 64) // general format

	fmt.Printf("Decimal ('f', 2): %s\n", decimal)
	fmt.Printf("Scientific ('e', 4): %s\n", scientific)
	fmt.Printf("General ('g', 4): %s\n", general)
}

// ===== STRING TO BOOLEAN =====

// stringToBoolExample shows converting strings to booleans
func StringToBoolExample() {
	fmt.Println("\n=== String to Boolean ===")

	// ParseBool accepts: "1", "t", "T", "true", "TRUE", "True",
	//                    "0", "f", "F", "false", "FALSE", "False"

	truthyValues := []string{"true", "True", "TRUE", "t", "T", "1"}
	for _, str := range truthyValues {
		bool, _ := strconv.ParseBool(str)
		fmt.Printf("ParseBool('%s'): %v\n", str, bool)
	}

	fmt.Println()

	falsyValues := []string{"false", "False", "FALSE", "f", "F", "0"}
	for _, str := range falsyValues {
		bool, _ := strconv.ParseBool(str)
		fmt.Printf("ParseBool('%s'): %v\n", str, bool)
	}
}

// ===== BOOLEAN TO STRING =====

// boolToStringExample shows converting booleans to strings
func BoolToStringExample() {
	fmt.Println("\n=== Boolean to String ===")

	// FormatBool
	true := strconv.FormatBool(true)
	false := strconv.FormatBool(false)

	fmt.Printf("FormatBool(true): %s\n", true)
	fmt.Printf("FormatBool(false): %s\n", false)
}

// ===== QUOTE =====

// quoteExample shows quoting strings
func QuoteExample() {
	fmt.Println("\n=== Quoting ===")

	str := "Hello\nWorld"

	// Quote - with double quotes
	quoted := strconv.Quote(str)
	fmt.Printf("Quote: %s\n", quoted)

	// QuoteToASCII - escape non-ASCII
	ascii := strconv.QuoteToASCII("Hello 世界")
	fmt.Printf("QuoteToASCII: %s\n", ascii)

	// Unquote - reverse operation
	unquoted, _ := strconv.Unquote(quoted)
	fmt.Printf("Unquote: %s\n", unquoted)
}

// PackageStrconvPlayground runs all strconv examples
func PackageStrconvPlayground() {
	fmt.Println("\n========== STRCONV PACKAGE ==========")
	StringToIntegerExample()
	IntegerToStringExample()
	StringToFloatExample()
	FloatToStringExample()
	StringToBoolExample()
	BoolToStringExample()
	QuoteExample()
}
