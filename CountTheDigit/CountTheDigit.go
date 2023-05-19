package kata

import (
	"strconv"
	"strings"
)

func NbDig(n, d int) int {
	var allnums string
	for i := 0; i <= n; i++ {
		allnums += strconv.Itoa(i * i)
	}
	return len(strings.Split(allnums, strconv.Itoa(d))) - 1
}
