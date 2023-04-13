package kata

import "strconv"

func SumMix(data []any) int {
	sum := 0
	for _, val := range data {
		switch val := val.(type) {
		case int:
			sum += val
		case string:
			s, _ := strconv.Atoi(val)
			sum += s
		}
	}
	return sum
}
