package kata

import "sort"

func MergeArrays(arr1, arr2 []int) []int {
	combinedArray := append(arr1, arr2...)
	sort.Ints(combinedArray)
	finalArray := []int{combinedArray[0]}
	for _, val := range combinedArray {
		if finalArray[len(finalArray)-1] != val {
			finalArray = append(finalArray, val)
		}
	}
	return finalArray
}
