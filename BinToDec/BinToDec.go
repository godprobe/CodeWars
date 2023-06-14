package kata

import (
	"math"
	"strconv"
	"strings"
)

func BinToDec(bin string) int {
	var total, iter, isOne int
	runes := strings.Split(bin, "")
	for i := len(runes); i > 0; i-- {
		isOne, _ = strconv.Atoi(runes[i-1])
		total += int(math.Pow(float64(2), float64(iter))) * isOne
		iter++
	}
	return total
}
