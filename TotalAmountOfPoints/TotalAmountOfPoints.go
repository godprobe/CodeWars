package kata

import (
	"strings"
)

func Points(pts []string) int {
	var score int
	for _, val := range pts {
		scores := strings.Split(val, ":")
		x := scores[0]
		y := scores[1]
		switch {
		case x > y:
			score += 3
		case x == y:
			score += 1
		}
	}
	return score
}
