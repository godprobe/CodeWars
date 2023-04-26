package kata

func MinMax(prices []int) [2]int {
	min, max := prices[0], prices[0]
	for _, val := range prices[1:] {
		if val < min {
			min = val
		} else if val > max {
			max = val
		}
	}
	return [2]int{min, max}
}
