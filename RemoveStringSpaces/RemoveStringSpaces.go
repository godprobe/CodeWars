package kata

import "strings"

func NoSpace(n string) string {
	return strings.ReplaceAll(n, " ", "")
}
