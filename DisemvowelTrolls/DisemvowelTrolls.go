package kata

import "regexp"

func Disemvowel(comment string) string {
	re := regexp.MustCompile(`(?i)[aeiou]`)
	return re.ReplaceAllString(comment, "")
}
