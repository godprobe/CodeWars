package kata

import "testing"

/*
Is the string uppercase?
Task
Create a method to see whether the string is ALL CAPS.

Examples (input, output},)
"c", False},
"C", True},
"hello I AM DONALD", False},
"HELLO I AM DONALD", True},
"ACSKLDFJSgSKLDFJSKLDFJ", False},
"ACSKLDFJSGSKLDFJSKLDFJ", True},
In this Kata, a string is said to be in ALL CAPS whenever it does not contain any lowercase letter so any string containing no letters at all is trivially considered to be in ALL CAPS.
*/

var (
	got, want bool
	cases     = []struct {
		input       MyString
		isUppercase bool
	}{
		{"c", false},
		{"C", true},
		{"hello I AM DONALD", false},
		{"HELLO I AM DONALD", true},
		{"ACSKLDFJSgSKLDFJSKLDFJ", false},
		{"ACSKLDFJSGSKLDFJSKLDFJ", true},
		{"a", false},
		{"b", false},
		{"c", false},
		{"d", false},
		{"e", false},
		{"f", false},
		{"g", false},
		{"h", false},
		{"i", false},
		{"j", false},
		{"k", false},
		{"l", false},
		{"m", false},
		{"n", false},
		{"o", false},
		{"p", false},
		{"q", false},
		{"r", false},
		{"s", false},
		{"t", false},
		{"u", false},
		{"v", false},
		{"w", false},
		{"x", false},
		{"y", false},
		{"z", false},
		{"abcdefghijklmnopqrstuvwxyz", false},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYz", false},
		{"false", false},
		{"true", false},
		{"False", false},
		{"True", false},
		{"WHAT DOES THE FOX SAY", true},
		{"HTML CSS JAVASCRIPT PYTHON C PERL LISP JAVA GO RUBY NODEJS RUST SCALA", true},
	}
)

func TestIsUpperCase(t *testing.T) {
	for _, val := range cases {
		got = val.input.IsUpperCase()
		want = val.isUppercase
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
