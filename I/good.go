package main

import "fmt"

type (
	Vehicle interface {
		Move()
	}

	Car interface {
		Vehicle
		Aircondition()
	}

	Motorcycle interface {
		Vehicle
	}

	GTR35 struct{}
	R15   struct{}
)

func (c *GTR35) Move() {
	fmt.Println("GTR35 is moving.")
}

func (c *GTR35) Aircondition() {
	fmt.Println("GTR35 have airconditioner.")
}

func (m *R15) Move() {
	fmt.Println("R15 is moving.")
}

func main() {
	car := &GTR35{}
	Motorcycle := &R15{}

	car.Move()
	car.Aircondition()

	Motorcycle.Move()
}