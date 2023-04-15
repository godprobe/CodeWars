package kata

import "sort"

func MergeArraysMap(arr1, arr2 []int) []int {
	combined := append(arr1, arr2...)
	keys := make(map[int]bool) // keys of a map are guaranteed to be unique
	final := []int{}
	for _, val := range combined {
		if _, key := keys[val]; !key {
			keys[val] = true
			final = append(final, val)
		}
	}
	sort.Ints(final)
	return final
}
