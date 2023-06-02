package kata

func Rps(player1, player2 string) string {
	if player1 == player2 {
		return "Draw!"
	}

	switch player1 + player2 {
	case "rockpaper", "paperscissors", "scissorsrock":
		return "Player 2 won!"
	}

	return "Player 1 won!"
}
