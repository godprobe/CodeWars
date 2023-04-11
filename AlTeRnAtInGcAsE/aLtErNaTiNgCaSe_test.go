package kata

/*
DESCRIPTION:
altERnaTIng cAsE <=> ALTerNAtiNG CaSe
Define String.prototype.toAlternatingCase (or a similar function/method such as to_alternating_case/toAlternatingCase/ToAlternatingCase in your selected language; see the initial solution for details) such that each lowercase letter becomes uppercase and each uppercase letter becomes lowercase. For example:

ToAlternatingCase("hello world"); // => "HELLO WORLD"
ToAlternatingCase("HELLO WORLD"); // => "hello world"
ToAlternatingCase("hello WORLD"); // => "HELLO world"
ToAlternatingCase("HeLLo WoRLD"); // => "hEllO wOrld"
ToAlternativeCase("12345"); // => "12345" (Non-alphabetical characters are unaffected)
ToAlternativeCase("1a2b3c4d5e"); // => "1A2B3C4D5E"
ToAlternativeCase("String.prototype.toAlternatingCase"); // => "sTRING.PROTOTYPE.TOaLTERNATINGcASE"
As usual, your function/method should be pure, i.e. it should not mutate the original string.
*/

import "testing"

var (
	cases = []struct {
		input  string
		output string
	}{
		{"hello world", "HELLO WORLD"},
		{"HELLO WORLD", "hello world"},
		{"hello WORLD", "HELLO world"},
		{"HeLLo WoRLD", "hEllO wOrld"},
		{"12345", "12345"},
		{"1a2b3c4d5e", "1A2B3C4D5E"},
		{"String.prototype.toAlternatingCase", "sTRING.PROTOTYPE.TOaLTERNATINGcASE"},
	}
)

func TestAlternatingCase(t *testing.T) {
	for _, test := range cases {
		got := ToAlternatingCase(test.input)
		want := test.output
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
