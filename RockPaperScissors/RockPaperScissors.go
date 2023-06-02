package kata

func Rps(player1, player2 string) string {
	rock := "rock"
	paper := "paper"
	scissors := "scissors"

	outcome := "Draw!"
	if player1 == player2 {
		return outcome
	}

	outcome = "Player 2 won!"
	switch player1 + player2 {
	case rock + paper:
		return outcome
	case paper + scissors:
		return outcome
	case scissors + rock:
		return outcome
	}
	outcome = "Player 1 won!"
	return outcome
}
