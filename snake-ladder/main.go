package main

import (
	"bufio"
	"fmt"
	"os"
	"snake-ladder/game"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the number of players")
	scanner.Scan()
	playerCountStr := scanner.Text()
	playerCount, err := strconv.Atoi(strings.TrimSpace(playerCountStr))
	if err != nil || playerCount <= 0 {
		fmt.Println("Invalid number of players")
		return
	}
	names := make([]string, 0)
	for i := 0; i < playerCount; i++ {
		fmt.Printf("Enter name for player %d: ", i+1)
		scanner.Scan()
		names[i] = strings.TrimSpace(scanner.Text())
	}
	g := game.NewGame(playerCount, names)

}
