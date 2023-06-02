package kata

func Rps(player1, player2 string) string {
	rock := "rock"
	paper := "paper"
	scissors := "scissors"

	if player1 == player2 {
		return "Draw!"
	}

	switch player1 + player2 {
	case rock + paper, paper + scissors, scissors + rock:
		return "Player 2 won!"
	}

	return "Player 1 won!"
}
