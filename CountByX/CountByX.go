package kata

func CountBy(x, n int) []int {
	arr := make([]int, n)
	for i := 1; i <= n; i++ {
		arr[i-1] = i * x
	}
	return arr
}
