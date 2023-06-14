package kata

import "testing"

/*
DESCRIPTION:
Complete the function which converts a binary number (given as a string) to a decimal number.
*/

var (
	got, want int
	cases     = []struct {
		input  string
		result int
	}{
		{"0", 0},
		{"1", 1},
		{"10", 2},
		{"11", 3},
		{"101010", 42},
		{"1001001", 73},
	}
)

func TestBinToDec(t *testing.T) {
	for _, val := range cases {
		got = BinToDec(val.input)
		want = val.result
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
}
