package kata

func Digitize(n int) []int {
	var arr []int

	for {
		arr = append(arr, n%10)
		n /= 10
		if n == 0 {
			return arr
		}
	}
}
