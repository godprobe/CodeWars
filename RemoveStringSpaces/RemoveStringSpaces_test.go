package kata

import "testing"

/*
Write a function that removes the spaces from the string, then return the resultant string.

Examples:

Input -> Output
"8 j 8   mBliB8g  imjB8B8  jl  B" -> "8j8mBliB8gimjB8B8jlB"
"8 8 Bi fk8h B 8 BB8B B B  B888 c hl8 BhB fd" -> "88Bifk8hB8BB8BBBB888chl8BhBfd"
"8aaaaa dddd r     " -> "8aaaaaddddr"
*/

var (
	got, want string
	cases     = []struct {
		in  string
		out string
	}{
		{"8 j 8   mBliB8g  imjB8B8  jl  B", "8j8mBliB8gimjB8B8jlB"},
		{"8 8 Bi fk8h B 8 BB8B B B  B888 c hl8 BhB fd", "88Bifk8hB8BB8BBBB888chl8BhBfd"},
		{"8aaaaa dddd r     ", "8aaaaaddddr"},
		{"jfBm  gk lf8hg  88lbe8 ", "jfBmgklf8hg88lbe8"},
		{"8j aam", "8jaam"},
	}
)

func TestNoSpace(t *testing.T) {
	for _, val := range cases {
		got = NoSpace(val.in)
		want = val.out
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
}
