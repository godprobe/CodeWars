package kata

import (
	"reflect"
	"testing"
)

/*
DESCRIPTION:
Story
Ben has a very simple idea to make some profit: he buys something and sells it again. Of course, this wouldn't give him any profit at all if he was simply to buy and sell it at the same price. Instead, he's going to buy it for the lowest possible price and sell it at the highest.

Task
Write a function that returns both the minimum and maximum number of the given list/array.

Examples (Input --> Output)
[1,2,3,4,5] --> [1,5]
[2334454,5] --> [5,2334454]
[1]         --> [1,1]
Remarks
All arrays or lists will always have at least one element, so you don't need to check the length. Also, your function will always get an array or a list, you don't have to check for null, undefined or similar.
*/

var (
	got, want [2]int
	cases     = []struct {
		prices []int
		minmax [2]int
	}{
		{[]int{1, 2, 3, 4, 5}, [2]int{1, 5}},
		{[]int{2334454, 5}, [2]int{5, 2334454}},
		{[]int{1}, [2]int{1, 1}},
	}
)

func TestMinMax(t *testing.T) {
	for _, val := range cases {
		got = MinMax(val.prices)
		want = val.minmax
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
