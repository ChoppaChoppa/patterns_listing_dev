package main

import "fmt"

func main() {
	auto := Car{4, "Volvo", PetrolMove{}}
	auto.Move()
	auto.Movable = ElectricMove{}
	auto.Move()
}

type IMovable interface {
	Move()
}

type PetrolMove struct {
	IMovable
}

func (p PetrolMove) Move() {
	fmt.Println("перемещение на бензине")
}

type ElectricMove struct {
	IMovable
}

func (e ElectricMove) Move() {
	fmt.Println("перемещение на электричестве")
}

type Car struct {
	Passengers int
	Model      string
	Movable    IMovable
}

func (c *Car) Move() {
	c.Movable.Move()
}
