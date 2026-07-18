package solutions

// WordFrequency counts the frequency of each word.
func WordFrequency(words []string) map[string]int {
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
	}
	return freq
}

// GroupByGrade groups students by grade.
func GroupByGrade(students map[string]string) map[string][]string {
	grouped := make(map[string][]string)
	for student, grade := range students {
		grouped[grade] = append(grouped[grade], student)
	}
	return grouped
}

// MergeMaps merges map a and b with sum.
func MergeMaps(a, b map[string]int) map[string]int {
	merged := make(map[string]int)
	for k, v := range a {
		merged[k] = v
	}
	for k, v := range b {
		merged[k] += v
	}
	return merged
}
