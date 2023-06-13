package kata

import "math"

func Movie(card, ticket int, perc float64) int {
	i := 0
	p := perc
	CostA := 0
	CostB := float64(card)
	for i = 1; float64(CostA) <= math.Ceil(CostB); i++ {
		CostA += ticket
		CostB += float64(ticket) * p
		p *= perc
	}
	return i - 1
}
