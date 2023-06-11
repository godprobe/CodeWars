package kata

import "fmt"

func MultiTable(number int) string {
	result := fmt.Sprintf("1 * %d = %d", number, number)
	for i := 2; i <= 10; i++ {
		result += fmt.Sprintf("\n%d * %d = %d", i, number, i*number)
	}
	return result
}
