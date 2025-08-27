package game

import "fmt"

type Game struct {
	Players       []*Player
	Board         *Board
	CurrentPlayer *Player
}

func NewGame(playerNumber int, names []string) *Game {
	game := Game{
		Board:   NewBoard(),
		Players: make([]*Player, 0, playerNumber),
	}
	game.InitializeGame(playerNumber, names)
	return &game
}

func (g *Game) InitializeGame(playerNumber int, names []string) {
	for i := 0; i < playerNumber; i++ {
		newPlayer := NewPlayer(names[i], i)
		g.Players = append(g.Players, newPlayer)
		g.CurrentPlayer = g.Players[0]
	}
}

func (g *Game) SwitchPlayer() {
	currentPlayer := g.CurrentPlayer
	noOfPlayers := len(g.Players)
	nextPlayerIndex := (currentPlayer.Index + 1) % noOfPlayers
	g.CurrentPlayer = g.Players[nextPlayerIndex]
}

func (g *Game) Play() {
	for {
		fmt.Printf("Player %s turn\n", g.CurrentPlayer.Name)
		g.CurrentPlayer.Move(g.Board)
		if g.CurrentPlayer.CurrentPos == 99 {
			fmt.Printf("Player %s has won the game!\n", g.CurrentPlayer.Name)
			break
		}
		g.SwitchPlayer()
	}
}
