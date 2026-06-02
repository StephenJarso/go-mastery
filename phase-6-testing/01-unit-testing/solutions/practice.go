package solutions



func SumOfEvens(nums []int) int {
	sum := 0
	for _, n := range nums {
		if n%2 == 0 {
			sum += n
		}
	}
	return sum
}
