package kata

func CountPositivesSumNegatives(numbers []int) []int {
	var (
		posCount, negSum int
	)
	for _, val := range numbers {
		switch {
		case val > 0:
			posCount += 1
		default:
			negSum += val
		}
	}
	return []int{posCount, negSum}
}
