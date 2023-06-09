package kata

import "testing"

/*
This kata is the first of a sequence of four about "Squared Strings".

You are given a string of n lines, each substring being n characters long: For example:

s = "abcd\nefgh\nijkl\nmnop"

We will study some transformations of this square of strings.

Vertical mirror: vert_mirror (or vertMirror or vert-mirror)
vert_mirror(s) => "dcba\nhgfe\nlkji\nponm"
Horizontal mirror: hor_mirror (or horMirror or hor-mirror)
 hor_mirror(s) => "mnop\nijkl\nefgh\nabcd"
or printed:

vertical mirror   |horizontal mirror
abcd --> dcba     |abcd --> mnop
efgh     hgfe     |efgh     ijkl
ijkl     lkji     |ijkl     efgh
mnop     ponm     |mnop     abcd
Task:
Write these two functions
and

high-order function oper(fct, s) where

fct is the function of one variable f to apply to the string s (fct will be one of vertMirror, horMirror)

Examples:
s = "abcd\nefgh\nijkl\nmnop"
oper(vert_mirror, s) => "dcba\nhgfe\nlkji\nponm"
oper(hor_mirror, s) => "mnop\nijkl\nefgh\nabcd"
Note:
The form of the parameter fct in oper changes according to the language. You can see each form according to the language in "Sample Tests".

Bash Note:
The input strings are separated by , instead of \n. The output strings should be separated by \r instead of \n. See "Sample Tests".
*/

var (
	vertFunc  FParam = VertMirror // VertMirror(string) string
	horFunc   FParam = HorMirror  // HorMirror(string) string
	got, want string
	cases     = []struct {
		whichFunc     FParam
		input, output string
	}{
		{vertFunc, "abcd\nefgh\nijkl\nmnop", "dcba\nhgfe\nlkji\nponm"},
		{horFunc, "abcd\nefgh\nijkl\nmnop", "mnop\nijkl\nefgh\nabcd"},
		{vertFunc, "hSgdHQ\nHnDMao\nClNNxX\niRvxxH\nbqTVvA\nwvSyRu", "QHdgSh\noaMDnH\nXxNNlC\nHxxvRi\nAvVTqb\nuRySvw"},
		{horFunc, "lVHt\nJVhv\nCSbg\nyeCt", "yeCt\nCSbg\nJVhv\nlVHt"},
	}
)

func TestOper(t *testing.T) {
	for _, val := range cases {
		want = val.output
		got = val.whichFunc(val.input)
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
}
