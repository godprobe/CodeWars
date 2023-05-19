package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func NbDig(n, d int) int {
	var allnums []int
	for i := 0; i <= n; i++ {
		allnums = append(allnums, i*i)
	}

	return len(strings.Split(fmt.Sprint(allnums), strconv.Itoa(d))) - 1
}
