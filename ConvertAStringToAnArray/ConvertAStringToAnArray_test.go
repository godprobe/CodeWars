package kata

import (
	"reflect"
	"testing"
)

/*
DESCRIPTION:
Write a function to split a string and convert it into an array of words.

Examples (Input ==> Output):
"Robin Singh" ==> ["Robin", "Singh"]

"I love arrays they are my favorite" ==> ["I", "love", "arrays", "they", "are", "my", "favorite"]
*/

var (
	got, want []string
	cases     = []struct {
		input  string
		output []string
	}{
		{"Robin Singh", []string{"Robin", "Singh"}},
		{"CodeWars", []string{"CodeWars"}},
		{"I love arrays they are my favorite", []string{"I", "love", "arrays", "they", "are", "my", "favorite"}},
		{"1 2 3", []string{"1", "2", "3"}},
	}
)

func TestStringToArray(t *testing.T) {
	for _, val := range cases {
		got = StringToArray(val.input)
		want = val.output
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
