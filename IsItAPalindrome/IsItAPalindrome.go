package kata

import (
	"strings"
)

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	// misses an opportunity to abort early in the next step, opting to use a more generic function
	z := reverseString(s)
	return z == s
}

func reverseString(s string) string {
	l := len(s) - 1
	if l <= -1 {
		return s
	}
	r := []rune(s)
	for i, j := 0, l; i < j; i, j = i+1, j-1 {
		temp := r[i]
		r[i] = r[j]
		r[j] = temp
	}
	return string(r)
}
