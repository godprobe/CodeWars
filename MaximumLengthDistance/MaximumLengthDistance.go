package kata

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}

	a1min, a1max := getMinMaxStrLens(a1)
	a2min, a2max := getMinMaxStrLens(a2)

	if a1max-a2min > a2max-a1min {
		return a1max - a2min
	}
	return a2max - a1min
}

func getMinMaxStrLens(ax []string) (min, max int) {
	min, max = len(ax[0]), len(ax[0])

	for i := 1; i < len(ax); i++ {
		x := len(ax[i])
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}
	return
}
