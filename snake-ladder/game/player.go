package game

import (
	"math/rand"
)

type Player struct {
	Name       string
	Index      int
	CurrentPos int
}

func NewPlayer(name string, index int) *Player {
	return &Player{
		Name:       name,
		Index:      index,
		CurrentPos: 0,
	}
}

func rollDice() int {
	return rand.Intn(6) + 1
}

func (p *Player) Move(b *Board) int {
	diceValue := rollDice()
	newPos := b.GetNewPosition(p.CurrentPos, diceValue)
	p.CurrentPos = newPos
	return diceValue
}
