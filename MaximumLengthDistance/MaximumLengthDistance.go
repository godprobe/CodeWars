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
	min = len(ax[0])
	max = len(ax[0])
	for i := 1; i < len(ax); i++ {
		switch x := len(ax[i]); {
		case x < min:
			min = x
		case x > max:
			max = x
		}
	}
	return
}
