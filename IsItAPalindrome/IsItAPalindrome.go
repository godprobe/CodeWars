package kata

import (
	"strings"
)

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	// misses an opportunity to abort early by using the more generic function below
	z := reverseString(s)
	return z == s
}

func reverseString(s string) string {
	l := len(s) - 1
	if l <= -1 {
		return s
	}
	r := []rune(s)
	for i := 0; i <= l/2; i++ {
		temp := r[i]
		r[i] = r[l-i]
		r[l-i] = temp
	}
	return string(r)
}
