package controlflow

// Go has only one looping construct: the "for" loop.
// Conditional statements are written using "if", "else if", and "else".
// Switch statements simplify nested conditionals.

// SumUpTo returns the sum of all integers from 1 up to limit.
func SumUpTo(limit int) int {
	sum := 0
	// Standard three-component for loop
	for i := 1; i <= limit; i++ {
		sum += i
	}
	return sum
}

// GetGradeName returns a text description of a score.
func GetGradeName(score int) string {
	// If statement with an initializer statement (common Go pattern)
	if isBonus := score > 100; isBonus {
		return "Superb"
	}

	// Switch case
	switch {
	case score >= 90:
		return "Excellent"
	case score >= 80:
		return "Very Good"
	case score >= 70:
		return "Good"
	case score >= 60:
		return "Pass"
	default:
		return "Fail"
	}
}
