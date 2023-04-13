package kata

func Past(h, m, s int) int {
	return 1000 * (s + 60*m + 3600*h)
}
