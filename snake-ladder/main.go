package main

import (
	"snake-ladder/game"
	"sync"
)

func main() {
	manager := game.GetNewGameManager()
	var wg sync.WaitGroup

	games := []*game.Game{
		manager.NewGame(2, []string{"Alice", "Bob"}),
		manager.NewGame(3, []string{"X", "Y", "Z"}),
		manager.NewGame(2, []string{"Mohan", "Sohan"}),
	}

	wg.Add(len(games))

	// Run them concurrently
	for _, g := range games {
		go func(g *game.Game) {
			defer wg.Done()
			g.Play()
		}(g)
	}

	wg.Wait()
}
