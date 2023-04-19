package kata

import (
	"strconv"
	"strings"
)

func Points(pts []string) int {
	var score int
	for _, val := range pts {
		scores := strings.Split(val, ":")
		x, _ := strconv.Atoi(scores[0])
		y, _ := strconv.Atoi(scores[1])
		if x > y {
			score += 3
		} else if x == y {
			score += 1
		}
	}
	return score
}
