package kata

import "testing"

/*
DESCRIPTION:
Clock shows h hours, m minutes and s seconds after midnight.

Your task is to write a function which returns the time since midnight in milliseconds.

Example:
h = 0
m = 1
s = 1

result = 61000
Input constraints:

0 <= h <= 23
0 <= m <= 59
0 <= s <= 59
*/

var (
	got, want int
	cases     = []struct {
		h      int
		m      int
		s      int
		result int
	}{
		{0, 1, 1, 61000},
		{1, 1, 1, 3661000},
		{0, 0, 0, 0},
		{1, 0, 1, 3601000},
		{1, 0, 0, 3600000},
	}
)

func TestPast(t *testing.T) {
	for _, time := range cases {
		got = Past(time.h, time.m, time.s)
		want = time.result
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
