package kata

import "math"

func PowersOfTwo(n int) []uint64 {
	var (
		powers []uint64
	)
	for i := 0; i <= n; i++ {
		powers = append(powers, uint64(math.Pow(2, float64(i))))
	}
	return powers
}
