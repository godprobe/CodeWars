package kata

import (
	"strconv"
	"strings"
)

func NbDig(n, d int) int {
	set := make([]int, 0)
	for i := 0; i <= n; i++ {
		set = append(set, i*i)
	}
	count := 0
	var allnums string
	for _, val := range set {
		allnums += strconv.Itoa(val)
	}
	count = len(strings.Split(allnums, strconv.Itoa(d))) - 1
	return count
}
