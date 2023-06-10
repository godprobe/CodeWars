package kata

import "testing"

/*
DESCRIPTION:
Write a function that checks if a given string (case insensitive) is a palindrome.
A palindrome is a word, number, phrase, or other sequence of symbols that reads the same
backwards as forwards, such as madam or racecar, the date and time 12/21/33 12:21,
and the sentence: "A man, a plan, a canal – Panama".
*/

var (
	got, want bool
	cases     = []struct {
		input  string
		result bool
	}{
		{"1234567", false},
		{"a", true},
		{"aba", true},
		{"Abba", true},
		{"hello", false},
		{"", true},
	}
)

func TestIsPalindrome(t *testing.T) {
	for _, val := range cases {
		got = IsPalindrome(val.input)
		want = val.result
		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	}
}
