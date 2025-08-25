package main

import "fmt"

type Car struct {
	Name  string
	Color string
}

func CarProvider() *Car {
	return &Car{}
}

func (c *Car) DisplayInfo() {
	fmt.Printf("this car name is %s,having color %s", c.Name, c.Color)
}

func main() {
	car1 := Car{"Maruti", "white"}
	car1.DisplayInfo()
}
