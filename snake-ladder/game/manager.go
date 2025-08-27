package game

import "sync"

type GameManager struct {
	Games []*Game
	mu    sync.Mutex
}

var (
	gameManagerInstance *GameManager
	once                sync.Once
)

func GetNewGameManager() *GameManager {
	once.Do(func() {
		gameManagerInstance = &GameManager{
			Games: make([]*Game, 0),
		}
	})
	return gameManagerInstance

}

func (gm *GameManager) NewGame(playerCount int, names []string) *Game {
	g := NewGame(playerCount, names)
	gm.mu.Lock()
	gm.Games = append(gm.Games, g)
	gm.mu.Unlock()
	return g
}
