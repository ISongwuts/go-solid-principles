package main

import "fmt"

type (
	Employee struct {}
	Displayer struct {}
)

func (e *Employee) Work() {
	fmt.Println("Employee is working.")
}

func (e *Employee) OverTimeWork() {
	fmt.Println("Employee is overtime Working.")
}

func (d *Displayer) DisplayTimeWorking() {
	fmt.Println("10 Hours of working.")
}

func main() {
	em := &Employee{}
	display := &Displayer{}
	em.Work()
	em.OverTimeWork()
	display.DisplayTimeWorking()
}