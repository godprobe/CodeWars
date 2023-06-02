package kata

func Rps(player1, player2 string) string {
	if player1 == player2 {
		return "Draw!"
	}

	switch player1 {
	case "rock":
		if player2 == "paper" {
			return "Player 2 won!"
		}
	case "paper":
		if player2 == "scissors" {
			return "Player 2 won!"
		}
	case "scissors":
		if player2 == "rock" {
			return "Player 2 won!"
		}
	}
	return "Player 1 won!"
}
