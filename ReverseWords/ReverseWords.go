package kata

import (
	"regexp"
)

func ReverseWords(str string) string {
	// a := regexp.MustCompile(`\b`) // fails on punctuation
	a := regexp.MustCompile(`\b`) // fails on punctuation
	sliced := a.Split(str, -1)
	// sliced := strings.Split(str, ` `)
	var result string
	for _, val := range sliced {
		result += ReverseWord(val)
	}
	return result
}

func ReverseWord(str string) string {
	var res string
	for i := len(str); i > 0; i-- {
		res = res + string(str[i-1])
	}
	return res
}
