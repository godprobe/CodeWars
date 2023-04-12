package kata

import "math"

func TwiceAsOld(ageOfFather, ageOfSon int) int {
	return int(math.Abs(float64(ageOfFather - 2*ageOfSon)))
}
