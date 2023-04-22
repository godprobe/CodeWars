package kata

import "testing"

/*
In this little assignment you are given a string of space separated numbers, and have to return the highest and lowest number.

Examples
HighAndLow("1 2 3 4 5")  // return "5 1"
HighAndLow("1 2 -3 4 5") // return "5 -3"
HighAndLow("1 9 3 4 -5") // return "9 -5"
Notes
All numbers are valid Int32, no need to validate them.
There will always be at least one number in the input string.
Output string must be two numbers separated by a single space, and highest number is first.
*/

var (
	got, want string
	cases     = []struct {
		input, output string
	}{
		{"1 2 3 4 5", "5 1"},
		{"1 2 -3 4 5", "5 -3"},
		{"1 9 3 4 -5", "9 -5"},
		{"8 3 -5 42 -1 0 0 -9 4 7 4 -4", "42 -9"},
		{"1 2 3", "3 1"},
	}
)

func TestHighAndLow(t *testing.T) {
	for _, val := range cases {
		got = HighAndLow(val.input)
		want = val.output

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
}
