package kata

import (
	"math"
)

func Digitize(n int) []int {
	arr := []int{}

	for i := 0; i < lenOfInt(n); i++ {
		// ((n / 10^i) % 10)
		arr = append(arr, int(math.Mod(math.Floor(float64(n)/math.Pow(10, float64(i))), 10)))
	}
	return arr
}

func lenOfInt(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}
