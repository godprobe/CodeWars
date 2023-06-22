package kata

func BouncingBall(h, bounce, window float64) int {
	if isValidInput(h, bounce, window) {
		sightings := 1
		for currentMax := h * bounce; window < currentMax; currentMax *= bounce {
			sightings += 2
		}
		return sightings
	} else {
		return -1
	}
}

func isValidInput(h, bounce, window float64) bool {
	return h > 0 && h > window && bounce > 0 && 1 > bounce
}
