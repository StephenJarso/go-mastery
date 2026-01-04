package variablestypes

import (
	"fmt"
	"strconv"
)

// In Go, variables are statically typed, meaning their type is determined at compile time.
// Variables can be declared in two main ways:
// 1. Using the "var" keyword (with or without an explicit type).
// 2. Using the short variable declaration operator ":=" (only inside functions).

// Global declarations must use "var"
var GlobalConfig = "active"

func ShowDeclarations() string {
	// 1. Explicit declaration with var
	var username string = "stephenjarso"

	// 2. Implicit type deduction with var
	var age = 30

	// 3. Short declaration syntax
	email := "stephenjacob815@gmail.com"

	// 4. Constants
	const AppName = "GoMastery"

	return fmt.Sprintf("App: %s, User: %s, Age: %d, Email: %s", AppName, username, age, email)
}

// ConvertStringToInt demonstrates explicit type conversions.
// Go requires explicit conversions; there is no implicit type casting.
func ConvertStringToInt(valStr string) (int, error) {
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return 0, fmt.Errorf("failed to convert type: %w", err)
	}
	return val, nil
}
