package kata

import "testing"

/*
Complete the method that takes a boolean value and return a "Yes" string for true, or a "No" string for false.
*/

func TestBoolToWord(t *testing.T) {
	var (
		got, want string
		cases     = []struct {
			boolean bool
			word    string
		}{
			{true, "Yes"},
			{false, "No"},
		}
	)
	for _, val := range cases {
		got = BoolToWord(val.boolean)
		want = val.word
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
