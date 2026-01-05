package controlflow

// PRACTICE EXERCISE: FizzBuzz
// Implement a function that processes numbers from start to end (inclusive)
// and returns a slice of strings with the following replacements:
// - If the number is divisible by 3, replace it with "Fizz".
// - If the number is divisible by 5, replace it with "Buzz".
// - If the number is divisible by both 3 and 5, replace it with "FizzBuzz".
// - Otherwise, convert the number to its string representation (e.g. "4").

import "strconv"

func FizzBuzz(start, end int) []string {
	result := make([]string, 0, end-start+1)
	for i := start; i <= end; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}
