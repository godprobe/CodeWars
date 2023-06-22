package kata

import "testing"

/*
Trolls are attacking your comment section!

A common way to deal with this situation is to remove all of the vowels from the trolls' comments, neutralizing the threat.

Your task is to write a function that takes a string and return a new string with all vowels removed.

For example, the string "This website is for losers LOL!" would become "Ths wbst s fr lsrs LL!".

Note: for this kata y isn't considered a vowel.
*/

var (
	got, want string
	cases = []struct {
		troll, trll string
	} {
		{"test", "tst"},
		{"This website is for losers LOL!", "Ths wbst s fr lsrs LL!"},
	}
)
func TestDisemvowel(t *testing.T) {
	for _, val := range cases {
		got = Disemvowel(val.troll)
		want = val.trll
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
}
