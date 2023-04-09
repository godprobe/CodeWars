package kata

func MessiGoals(goals []int) int {
	var (
		total int
	)
	for _, val := range goals {
		total += val
	}
	return total
}
