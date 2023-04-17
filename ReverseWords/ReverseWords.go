package kata

func ReverseWords(str string) string {
	var result string
	var word string
	for _, val := range str {
		switch string(val) {
		case " ":
			result += word + " "
			word = ""
		default:
			word = string(val) + word
		}
	}
	result += word
	return result
}
