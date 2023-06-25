package kata

import "math"

func MxDifLg(a1 []string, a2 []string) int {
	maxDiff := -1
	for _, val1 := range a1 {
		for _, val2 := range a2 {
			x := int(math.Abs(float64(len(val1) - len(val2))))
			if maxDiff < x {
				maxDiff = x
			}
		}
	}
	return maxDiff
}
