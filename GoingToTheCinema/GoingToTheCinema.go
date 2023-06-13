package kata

import (
	"fmt"
	"math"
)

func Movie(card, ticket int, perc float64) int {
	i := 0
	p := perc
	SystemACost := 0
	SystemBCost := float64(card)
	for i = 1; float64(SystemACost) <= math.Ceil(SystemBCost); i++ {
		SystemACost += ticket
		SystemBCost += float64(ticket) * p
		p *= perc
		fmt.Printf("A: %d, \tB: %v, \tTickets: %d\n", SystemACost, SystemBCost, i)
	}
	return i - 1
}
