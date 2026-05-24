package solutions

import "errors"



func Filter(numbers []int, predicate func(int) bool) []int {
	var result []int
	for _, n := range numbers {
		if predicate(n) {
			result = append(result, n)
		}
	}
	return result
}

func NewCalculator(op string) (func(int, int) int, error) {
	switch op {
	case "+":
		return func(a, b int) int { return a + b }, nil
	case "-":
		return func(a, b int) int { return a - b }, nil
	case "*":
		return func(a, b int) int { return a * b }, nil
	default:
		return nil, errors.New("unsupported operation")
	}
}

func SafeDivide(dividend, divisor int) (quotient int, remainder int, err error) {
	if divisor == 0 {
		err = errors.New("cannot divide by zero")
		return
	}
	quotient = dividend / divisor
	remainder = dividend % divisor
	return
}
