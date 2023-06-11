package kata

import "fmt"

func MultiTable(number int) string {
	var result string
	result += "1 * " + fmt.Sprint(number) + " = " + fmt.Sprint(number)
	for i := 2; i <= 10; i++ {
		result += "\n" +
			fmt.Sprint(i) +
			" * " +
			fmt.Sprint(number) +
			" = " +
			fmt.Sprint(i*number)
	}
	return result
}
