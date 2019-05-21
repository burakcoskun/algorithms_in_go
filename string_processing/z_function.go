package string_processing

import "unicode/utf8"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Z_function(s string) []int {
	length := utf8.RuneCountInString(s)
	z := make([]int, length)

	var left, right int
	for i := 1; i < length; i++ {
		if i < right {
			z[i] = min(z[i-left], right-i+1)
		}
		for i+z[i] < length && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > right {
			left = i
			right = i + z[i] - 1
		}
	}
	return z
}
