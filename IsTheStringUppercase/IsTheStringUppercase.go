package kata

import (
	"fmt"
	"strings"
)

type MyString string

func (s MyString) IsUpperCase() bool {
	return strings.ToUpper(fmt.Sprint(s)) == fmt.Sprint(s)
}
