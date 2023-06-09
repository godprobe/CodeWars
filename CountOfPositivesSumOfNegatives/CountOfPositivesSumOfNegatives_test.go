package kata

import (
	"reflect"
	"testing"
)

/*
DESCRIPTION:
Given an array of integers.

Return an array, where the first element is the count of positives numbers and the second element is sum of negative numbers. 0 is neither positive nor negative.

If the input is an empty array or is null, return an empty array.

Example
For input [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15], you should return [10, -65].
*/

var (
	got, want []int
	cases     = []struct {
		input, output []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}, []int{10, -65}},
	}
)

func TestCountPositivesSumNegatives(t *testing.T) {
	for _, val := range cases {
		got = CountPositivesSumNegatives(val.input)
		want = val.output
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
