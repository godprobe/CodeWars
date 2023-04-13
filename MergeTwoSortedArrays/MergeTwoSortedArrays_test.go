package kata

import (
	"reflect"
	"testing"
)

/*
DESCRIPTION:
You are given two sorted arrays that both only contain integers. Your task is to find a way to merge them into a single one, sorted in asc order. Complete the function mergeArrays(arr1, arr2), where arr1 and arr2 are the original sorted arrays.

You don't need to worry about validation, since arr1 and arr2 must be arrays with 0 or more Integers. If both arr1 and arr2 are empty, then just return an empty array.

Note: arr1 and arr2 may be sorted in different orders. Also arr1 and arr2 may have same integers. Remove duplicated in the returned result.

Examples (input -> output)
* [1, 2, 3, 4, 5], [6, 7, 8, 9, 10] -> [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

* [1, 3, 5, 7, 9], [10, 8, 6, 4, 2] -> [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

* [1, 3, 5, 7, 9, 11, 12], [1, 2, 3, 4, 5, 10, 12] -> [1, 2, 3, 4, 5, 7, 9, 10, 11, 12]
Happy coding!
*/

var (
	got, want []int
	cases     = []struct {
		array1 []int
		array2 []int
		result []int
	}{
		{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8}},
		{[]int{1, 3, 5, 7, 9}, []int{10, 8, 6, 4, 2}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{1, 3, 5, 7, 9, 11, 12}, []int{1, 2, 3, 4, 5, 10, 12}, []int{1, 2, 3, 4, 5, 7, 9, 10, 11, 12}},
	}
)

func TestMergeArrays(t *testing.T) {
	for _, val := range cases {
		got := MergeArrays(val.array1, val.array2)
		want := val.result
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
