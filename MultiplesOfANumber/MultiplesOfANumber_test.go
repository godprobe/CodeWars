package kata

import (
	"reflect"
	"testing"
)

/*DESCRIPTION:
In this simple exercise, you will build a program that takes a value, integer , and returns a list of its multiples up to another value, limit . If limit is a multiple of integer, it should be included as well. There will only ever be positive integers passed into the function, not consisting of 0. The limit will always be higher than the base.

For example, if the parameters passed are (2, 6), the function should return [2, 4, 6] as 2, 4, and 6 are the multiples of 2 up to 6.
*/

var (
	got, want []int
	cases     = []struct {
		integer, limit int
		multiples      []int
	}{
		{5, 25, []int{5, 10, 15, 20, 25}},
		{1, 2, []int{1, 2}},
		{2, 6, []int{2, 4, 6}},
	}
)

func TestFindMultiples(t *testing.T) {
	for _, val := range cases {
		got = FindMultiples(val.integer, val.limit)
		want = val.multiples
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
