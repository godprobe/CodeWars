package kata

func MxDifLg(a1 []string, a2 []string) (maxDiff int) {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	a1min, a1max := getMinMaxStrLens(a1)
	a2min, a2max := getMinMaxStrLens(a2)
	if a1min != 0 && a2min != 0 {
		if a1max-a2min > a2max-a2min {
			maxDiff = a1max - a2min
		} else {
			maxDiff = a2max - a1min
		}
	}
	return
}

func getMinMaxStrLens(ax []string) (min, max int) {
	min = len(ax[0])
	for _, val := range ax {
		x := len(val)
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}
	return
}
