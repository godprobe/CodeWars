package kata

import (
	"strings"
)

func VertMirror(s string) string {
	split := strings.Split(s, "\n")
	for i := 0; i < len(split); i++ {
		split[i] = reverse(split[i])
	}
	return strings.Join(split, "\n")
}

func HorMirror(s string) string {
	s = reverse(s)
	split := strings.Split(s, "\n")
	for i := 0; i < len(split); i++ {
		split[i] = reverse(split[i])
	}
	return strings.Join(split, "\n")
}

type FParam func(string) string

func Oper(f FParam, x string) string {
	return f(x)
}

func reverse(s string) string {
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
