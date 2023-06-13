package kata

func Movie(card, ticket int, perc float64) int {
	var (
		SystemACost int
		SystemBCost float64
		i           int
	)
	SystemACost = ticket
	SystemBCost = float64(card) + float64(ticket)*perc
	for i = 2; float64(SystemACost) < SystemBCost; i++ {
		SystemACost += ticket
		SystemBCost += float64(ticket) * (float64(i-1) * perc)
		perc *= perc
		//		fmt.Printf("System A: %v, System B: %v after %d tickets.\n", SystemACost, SystemBCost, i)
	}
	return i
}
