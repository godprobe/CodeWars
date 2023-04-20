package kata

func GetGrade(a, b, c int) rune {
	var finalGrade rune

	switch avg := (a + b + c) / 3; {
	case avg > 89:
		finalGrade = 'A'
	case avg > 79:
		finalGrade = 'B'
	case avg > 69:
		finalGrade = 'C'
	case avg > 59:
		finalGrade = 'D'
	default:
		finalGrade = 'F'
	}

	return finalGrade
}
