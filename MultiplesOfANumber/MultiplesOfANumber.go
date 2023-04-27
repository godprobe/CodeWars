package kata

func FindMultiples(integer, limit int) []int {
	multiples := []int{integer}
	for i := 2 * integer; i <= limit; i += integer {
		multiples = append(multiples, i)
	}
	return multiples
}
