package kata

import "strings"
import "strconv"
import "fmt"

func HighAndLow(n string) string {
	parsed := strings.Split(n, " ")
	high, _ := strconv.Atoi(parsed[0])
	low, _ := strconv.Atoi(parsed[0])
	var v int
	for _, val := range parsed {
		v, _ = strconv.Atoi(val)
		if v > high {
			high = v
		}
		if v < low {
			low = v
		}
	}
	return fmt.Sprint(high) + " " + fmt.Sprint(low)
}
