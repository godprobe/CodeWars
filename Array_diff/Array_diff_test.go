package kata

import (
	"reflect"
	"testing"
)

/*
Your goal in this kata is to implement a difference function,
which subtracts one list from another and returns the result.

It should remove all values from list a, which are present in list b keeping their order.

array_diff({1, 2}, 2, {1}, 1, *z) == {2} (z == 1)
If a value is present in b, all of its occurrences must be removed from the other:

array_diff({1, 2, 2, 2, 3}, 5, {2}, 1, *z) == {1, 3} (z == 2)
*/

var (
	got, want, emptyArr []int
	cases               = []struct {
		arrA, arrB, arrDiff []int
	}{
		{[]int{1, 2}, []int{1}, []int{2}},
		{[]int{1, 2, 2}, []int{1}, []int{2, 2}},
		{[]int{1, 2, 2}, []int{2}, []int{1}},
		{[]int{1, 2, 2}, emptyArr, []int{1, 2, 2}},
		{emptyArr, []int{1, 2}, emptyArr},
		{[]int{1, 2, 3}, []int{1, 2}, []int{3}},
	}
)

func TestArrDiff(t *testing.T) {
	for _, val := range cases {
		got = ArrayDiff(val.arrA, val.arrB)
		want = val.arrDiff
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
