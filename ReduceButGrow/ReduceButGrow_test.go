package kata

import "testing"

/*
DESCRIPTION:
Given a non-empty array of integers, return the result of multiplying the values together in order. Example:

[1, 2, 3, 4] => 1 * 2 * 3 * 4 = 24
*/

func TestGrow(t *testing.T) {
	got := Grow([]int{1, 2, 3, 4})
	want := 24
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
