package kata

import (
	"math"
	"strings"
)

func BinToDec(bin string) int {
	var total int
	l := len(bin) - 1
	m := map[string]int{"0": 0, "1": 1}
	runes := strings.Split(bin, "")
	for i, val := range runes {
		total += int(math.Pow(2.0, float64(l-i))) * m[val]
	}
	return total
}
