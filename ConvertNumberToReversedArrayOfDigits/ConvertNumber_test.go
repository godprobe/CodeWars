package kata

import (
	"reflect"
	"testing"
)

/*
DESCRIPTION:
Convert number to reversed array of digits
Given a random non-negative number, you have to return the digits of this number within an array in reverse order.

Example(Input => Output):
35231 => [1,3,2,5,3]
0 => [0]
*/

var (
	got, want []int
	cases     = []struct {
		n   int
		arr []int
	}{
		{35231, []int{1, 3, 2, 5, 3}},
		{0, []int{0}},
	}
)

func TestDigitize(t *testing.T) {
	for _, val := range cases {
		got = Digitize(val.n)
		want = val.arr
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
