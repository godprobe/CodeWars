package kata

import "testing"

/*
DESCRIPTION:
Given an array of integers as strings and numbers, return the sum of the array values as if all were numbers.

Return your answer as a number.

*/

var (
	got, want int
	cases     = []struct {
		data   []any
		result int
	}{
		{[]any{9, 3, "7", "3"}, 22},
		{[]any{"5", "0", 9, 3, 2, 1, "9", 6, 7}, 42},
		{[]any{"3", 6, 6, 0, "5", 8, 5, "6", 2, "0"}, 41},
		{[]any{"1", "5", "8", 8, 9, 9, 2, "3"}, 45},
		{[]any{8, 0, 0, 8, 5, 7, 2, 3, 7, 8, 6, 7}, 61},
	}
)

func TestSumMix(t *testing.T) {
	for _, val := range cases {
		got = SumMix(val.data)
		want = val.result
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
