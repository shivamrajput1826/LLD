package game

import (
	"fmt"
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
	fmt.Printf("Player %s turn\n", p.Name)
	diceValue := rollDice()
	fmt.Printf("Player %s has thrown %d\n", p.Name, diceValue)
	newPos := b.GetNewPosition(p.CurrentPos, diceValue)
	fmt.Printf("Player %s new Position is: %d\n", p.Name, newPos)
	p.CurrentPos = newPos
	return diceValue
}
