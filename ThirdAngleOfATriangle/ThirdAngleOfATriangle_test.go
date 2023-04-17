package kata

import "testing"

/*
DESCRIPTION:
You are given two interior angles (in degrees) of a triangle.

Write a function to return the 3rd.

Note: only positive integers will be tested.

https://en.wikipedia.org/wiki/Triangle
*/

var (
	got, want int
	cases     = []struct {
		angleA, angleB, angleC int
	}{
		{30, 60, 90},
		{60, 60, 60},
		{43, 78, 59},
		{10, 20, 150},
	}
)

func TestOtherAngle(t *testing.T) {
	for _, val := range cases {
		got = OtherAngle(val.angleA, val.angleB)
		want = val.angleC
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
