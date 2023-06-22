package kata

import "regexp"

func Disemvowel(comment string) string {
	re := regexp.MustCompile(`[aeiouAEIOU]`)
	return re.ReplaceAllString(comment, "")
}
