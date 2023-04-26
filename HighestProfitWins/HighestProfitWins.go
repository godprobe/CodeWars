package kata

func MinMax(prices []int) [2]int {
	result := [2]int{prices[0], prices[0]}
	for _, val := range prices {
		if val < result[0] {
			result[0] = val
		}
		if val > result[1] {
			result[1] = val
		}
	}
	return result
}
