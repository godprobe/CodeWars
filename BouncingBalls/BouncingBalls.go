package kata

func BouncingBall(h, bounce, window float64) int {
	if AssertValidInput(h, bounce, window) {
		var sightings int
		for currentMax := h; currentMax > window; currentMax *= bounce {
			if currentMax > window {
				sightings++
			} else {
				return sightings
			}
			if currentMax*bounce > window {
				sightings++
			}
		}
		return sightings
	} else {
		return -1
	}
}

func AssertValidInput(h, bounce, window float64) bool {
	return h > 0 && bounce > 0 && bounce < 1 && window < h
}
