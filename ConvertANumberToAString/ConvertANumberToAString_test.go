package kata

import "testing"

/*
DESCRIPTION:
We need a function that can transform a number (integer) into a string.

What ways of achieving this do you know?

Examples (input --> output):
123  --> "123"
999  --> "999"
-100 --> "-100"
*/

var (
	got, want string
	cases     = []struct {
		input  int
		output string
	}{
		{67, "67"},
		{79585, "79585"},
		{79585, "79585"},
		{3, "3"},
		{-1, "-1"},
	}
)

func TestNumberToString(t *testing.T) {
	for _, val := range cases {
		got = NumberToString(val.input)
		want = val.output
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
}
