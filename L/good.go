package main

import "fmt"

type (
	Aircraft interface {
		Fly()
	}

	Boeing747 struct {
		Aircraft
	}

	Spitfire struct {
		Aircraft
	}
)

func (a *Boeing747) Fly() {
	fmt.Println("Boeing can fly.")
}

func (a *Spitfire) Fly() {
	fmt.Println("Spitfire can fly.")
}

func (a *Spitfire) Fire() {
	fmt.Println("Spitfire can fire.")
}

func main() {
	boeing747 := &Boeing747{}
	spitfire := &Spitfire{}

	boeing747.Fly()

	spitfire.Fire()
	spitfire.Fly()
}