package main

import "tic-tac/game"

func main() {
	player1 := game.NewPlayer("Player 1", 'X')
	player2 := game.NewPlayer("Player 2", 'O')

	game := game.NewGame(player1, player2)
	game.Play()
}
