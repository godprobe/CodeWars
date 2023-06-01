package kata

func CountBy(x, n int) []int {
	result := []int{}
	for i := 1; i <= x; i++ {
		result = append(result, n*i)
	}
	return result
}
