package kata

import "fmt"

func MultiTable(number int) string {
	var result string
	for i := 1; i <= 10; i++ {
		result += fmt.Sprintf("%d * %d = %d\n", i, number, i*number)
	}
	return result[:len(result)-1]
}
