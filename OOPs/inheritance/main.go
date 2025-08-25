package main

import "fmt"

// go does not provide inheritance (classical).
// it does so using stuct embedding

type Animal struct {
	Name string
}

func (a *Animal) Eat() {
	fmt.Printf("a %s eat well\n", a.Name)
}

// a child struct can override the parent struct method
func (a *Animal) Speak() {
	fmt.Printf("a %s speak", a.Name)
}

type Dog struct {
	Animal
}

func (d *Dog) Bark() {
	fmt.Printf("d %s is barking\n", d.Name)
}
func (d *Dog) Speak() {
	fmt.Printf("a %s bark", d.Name)
}
func main() {
	myDog := Dog{Animal{Name: "Buddy"}}
	myDog.Eat()
	myDog.Bark()
	myDog.Speak()
}

// multiple Inheritance in struct
// First embedded struct
type Engine struct {
	Horsepower int
}

// Second embedded struct
type Wheels struct {
	Count int
}

// Car struct embedding Engine and Wheels
type Car struct {
	Engine
	Wheels
}

func New() {
	myCar := Car{Engine{200}, Wheels{4}}
	fmt.Println("Horsepower:", myCar.Horsepower)
	fmt.Println("Wheels:", myCar.Count)
}

// In go interfaces allow struct based polymorphism
type Janwar interface {
	Speak()
}

type Doggy struct {
	Name string
}

func (d *Doggy) Speak() {
	fmt.Println(d.Name, "barks")
}

type Employee struct {
	Name   string
	Salary float64
}

func (e Employee) DisplayDetails() {
	fmt.Printf("Employee: %s, Salary: %.2f\n", e.Name, e.Salary)
}

type Manager struct {
	Employee
	Bonus float64
}

func (m Manager) DisplayDetails() {
	m.Employee.DisplayDetails()
	fmt.Printf("Bonus: %.2f\n", m.Bonus)
}

func EmployeeNew() {
	manager := Manager{Employee: Employee{Name: "Alice", Salary: 70000}, Bonus: 10000}
	manager.DisplayDetails()
}
