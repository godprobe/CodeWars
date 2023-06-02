package kata

import "testing"

/*
DESCRIPTION:
Rock Paper Scissors
Let's play! You have to return which player won! In case of a draw return Draw!.

Examples(Input1, Input2 --> Output):

"scissors", "paper" --> "Player 1 won!"
"scissors", "rock" --> "Player 2 won!"
"paper", "paper" --> "Draw!"
*/

var (
	got, want string
	cases     = []struct {
		player1, player2, outcome string
	}{
		{"scissors", "paper", "Player 1 won!"},
		{"scissors", "rock", "Player 2 won!"},
		{"paper", "paper", "Draw!"},
	}
)

func TestRps(t *testing.T) {
	for _, val := range cases {
		got = Rps(val.player1, val.player2)
		want = val.outcome
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
}
