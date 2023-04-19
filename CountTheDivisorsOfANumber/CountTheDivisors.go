package kata

import "math"

func Divisors(n int) int {
	var count = 1
	for f := 0.0; f <= float64(n)/2; f++ {
		if math.Mod(float64(n), f) == 0 {
			count++
		}
	}
	return count
}
