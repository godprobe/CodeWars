package kata

import "testing"

/*
DESCRIPTION:
Your function takes two arguments:

current father's age (years)
current age of his son (years)
Сalculate how many years ago the father was twice as old as his son (or in how many years he will be twice as old). The answer is always greater or equal to 0, no matter if it was in the past or it is in the future.
*/

func TestTwiceAsOld(t *testing.T) {
	var (
		got, want int
		cases     = []struct {
			ageOfFather int
			ageOfSon    int
			years       int
		}{
			{36, 7, 22},
			{55, 30, 5},
			{42, 21, 0},
			{22, 1, 20},
			{29, 0, 29},
		}
	)
	for _, val := range cases {
		got = TwiceAsOld(val.ageOfFather, val.ageOfSon)
		want = val.years
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
