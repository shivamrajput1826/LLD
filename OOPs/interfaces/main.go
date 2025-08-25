package main

import "fmt"

type Vehicle interface {
	Start()
	Stop()
}

type Car struct {
	Name string
}

func (c *Car) Start() {
	fmt.Println("The car is starting")
}
func (c *Car) Stop() {
	fmt.Println("The car is stopped")
}

func main() {
	car1 := Car{"Maruti"}
	car1.Start()
	car1.Stop()
}

// now interface{} => any
// one type can implement multiple interfaces
type Flyable interface {
	Fly()
}

// Second interface
type Drivable interface {
	Drive()
}

// Implementing multiple interfaces
type FlyingCar struct{}

func (f FlyingCar) Fly() {
	fmt.Println("FlyingCar is flying...")
}

func (f FlyingCar) Drive() {
	fmt.Println("FlyingCar is driving...")
}

// interfaces can be composite together can be used in
// case one type wants a subpart of type to be implememented

type Engine interface {
	Start()
	Stop()
}

type Transmission interface {
	ShiftGear(gear int)
}

type CarInterface interface {
	Engine
	Transmission
}
