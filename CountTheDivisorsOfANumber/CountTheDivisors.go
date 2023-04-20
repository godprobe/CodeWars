package kata

func Divisors(n int) int {
	var count = 1
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count
}
