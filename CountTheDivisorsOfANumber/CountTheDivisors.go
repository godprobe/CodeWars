package kata

func Divisors(n int) int {
	var count = 1
	for f := 1; f <= n/2; f++ {
		if n%f == 0 {
			count++
		}
	}
	return count
}
