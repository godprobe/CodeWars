package kata

/*
Messi goals function
Messi is a soccer player with goals in three leagues:

LaLiga
Copa del Rey
Champions
Complete the function to return his total number of goals in all three leagues.

Note: the input will always be valid.

For example:

5, 10, 2  -->  17
*/

import "testing"

func TestMessi(t *testing.T) {
	got := MessiGoals([]int{5, 10, 2})
	want := 17
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
