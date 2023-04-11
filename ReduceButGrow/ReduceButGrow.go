package kata

func Grow(arr []int) int {
	result := 1
	for _, val := range arr {
		result *= val
	}
	return result
}
