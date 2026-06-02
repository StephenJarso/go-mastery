package solutions

import (
	"strings"
)


func ConcatenateStringsPlus(strs []string) string {
	res := ""
	for _, s := range strs {
		res += s
	}
	return res
}

func ConcatenateStringsBuilder(strs []string) string {
	var sb strings.Builder
	for _, s := range strs {
		sb.WriteString(s)
	}
	return sb.String()
}
