package game

type Game struct {
	Players []*Player
	Board   *Board
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
	}
}

func (g *Game) Winner() *Player {
	for _, p := range g.Players {
		if p.CurrentPos == 99 {
			return p
		}
	}
	return nil
}
