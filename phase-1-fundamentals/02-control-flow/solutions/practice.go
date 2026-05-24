package solutions

import (
	"strconv"
)


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

func FindPrimes(limit int) []int {
	var primes []int
	for i := 2; i <= limit; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes
}

func CountCharFrequency(s string) map[rune]int {
	freq := make(map[rune]int)
	for _, r := range s {
		freq[r]++
	}
	return freq
}
