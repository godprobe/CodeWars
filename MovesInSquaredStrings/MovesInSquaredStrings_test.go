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
	vertFuncString = "VertMirror"
	horFuncString  = "HorMirror"
	got, want      string
	cases          = []struct {
		whichFunc     string
		input, output string
	}{
		{vertFuncString, "hSgdHQ\nHnDMao\nClNNxX\niRvxxH\nbqTVvA\nwvSyRu", "QHdgSh\noaMDnH\nXxNNlC\nHxxvRi\nAvVTqb\nuRySvw"},
		{horFuncString, "lVHt\nJVhv\nCSbg\nyeCt", "yeCt\nCSbg\nJVhv\nlVHt"},
	}
)

func TestVertMirror(t *testing.T) {
	for _, val := range cases {
		if val.whichFunc == vertFuncString {
			got = VertMirror(val.input)
			want = val.output
		}
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
}

func TestHorMirror(t *testing.T) {
	for _, val := range cases {
		if val.whichFunc == horFuncString {
			got = HorMirror(val.input)
			want = val.output
		}
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
}

func TestOper(t *testing.T) {
	for _, val := range cases {
		want = val.output
		switch val.whichFunc {
		case vertFuncString:
			got = VertMirror(val.input)
		case horFuncString:
			got = HorMirror(val.input)
		default:
			t.Errorf("Incompatible func parameter.")
		}
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
}
