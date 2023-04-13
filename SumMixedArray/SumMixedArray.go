package kata

import "strconv"

func SumMix(data []any) int {
	sum := 0
	for _, val := range data {
		iVal, ok := val.(int)
		if ok {
			sum += iVal
		}
		sVal, ok := val.(string)
		if ok {
			sVal, _ := strconv.Atoi(sVal)
			sum += sVal
		}
	}
	return sum
}
