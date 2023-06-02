package kata

func Rps(player1, player2 string) string {
	outcome := "Draw!"
	rock := "rock"
	paper := "paper"
	scissors := "scissors"

	if player1 == player2 {
		return outcome
	}

	outcome = "Player 2 won!"
	switch player1 {
	case rock:
		if player2 == paper {
			return outcome
		}
	case paper:
		if player2 == scissors {
			return outcome
		}
	case scissors:
		if player2 == rock {
			return outcome
		}
	}
	outcome = "Player 1 won!"
	return outcome
}
