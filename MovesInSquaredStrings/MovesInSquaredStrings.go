package kata

import (
	"strings"
)

func VertMirror(s string) string {
	split := strings.Split(s, "\n")
	split = reverseSlice(split)
	return strings.Join(split, "\n")
}

func HorMirror(s string) string {
	s = reverseString(s)
	split := strings.Split(s, "\n")
	split = reverseSlice(split)
	return strings.Join(split, "\n")
}

type FParam func(string) string

func Oper(f FParam, x string) string {
	return f(x)
}

func reverseString(s string) string {
	var (
		temp rune
	)
	runes := []rune(s)
	last := len(runes) - 1
	for i := 0; i <= last/2; i++ {
		temp = runes[i]
		runes[i] = runes[last-i]
		runes[last-i] = temp
	}
	return string(runes)
}

func reverseSlice(st []string) []string {
	for i := 0; i < len(st); i++ {
		st[i] = reverseString(st[i])
	}
	return st
}
