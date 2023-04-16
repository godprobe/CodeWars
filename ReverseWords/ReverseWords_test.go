package kata

import "testing"

/*
DESCRIPTION:
Complete the function that accepts a string parameter, and reverses each word in the string. All spaces in the string should be retained.

Examples
"This is an example!" ==> "sihT si na !elpmaxe"
"double  spaces"      ==> "elbuod  secaps"
*/

var (
	got, want string
	cases     = []struct {
		input  string
		output string
	}{
		{"This is an example!", "sihT si na !elpmaxe"},
		{"double  spaces", "elbuod  secaps"},
		{"The quick brown fox jumps over the lazy dog.", "ehT kciuq nworb xof spmuj revo eht yzal .god"},
		{"apple", "elppa"},
		{"a b c d", "a b c d"},
		{"double  spaced  words", "elbuod  decaps  sdrow"},
		{"stressed desserts", "desserts stressed"},
		{"hello hello", "olleh olleh"},
	}
)

func TestReverseWords(t *testing.T) {
	for _, val := range cases {
		got = ReverseWords(val.input)
		want = val.output
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
}
