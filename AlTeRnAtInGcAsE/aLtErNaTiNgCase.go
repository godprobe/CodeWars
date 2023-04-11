package kata

import (
	"strings"
	"unicode"
)

func ToAlternatingCase(input string) string {
	result := ""
	for _, char := range input {
		if unicode.IsLower(char) {
			result += strings.ToUpper(string(char))
		} else {
			result += strings.ToLower(string(char))
		}
	}
	return result
}
