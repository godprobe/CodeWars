package kata_test

import (
	"reflect"
	"testing"
)

/*
DESCRIPTION:
Create a function with two arguments that will return an array of the first n multiples of x.

Assume both the given number and the number of times to count will be positive numbers greater than 0.

Return the results as an array or list ( depending on language ).

Examples
countBy(1,10)  should return  {1,2,3,4,5,6,7,8,9,10}
countBy(2,5)  should return {2,4,6,8,10}
*/

var (
	got, want []int
	cases     = []struct {
		x, n   int
		result []int
	}{
		{1, 10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{2, 5, []int{2, 4, 6, 8, 10}},
		{1, 5, []int{1, 2, 3, 4, 5}},
		{2, 5, []int{2, 4, 6, 8, 10}},
		{3, 5, []int{3, 6, 9, 12, 15}},
		{50, 5, []int{50, 100, 150, 200, 250}},
		{100, 5, []int{100, 200, 300, 400, 500}},
	}
)

func TestCountBy(t *testing.T) {
	for _, val := range cases {
		got = CountBy(val.x, val.n)
		want = val.result
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

}
