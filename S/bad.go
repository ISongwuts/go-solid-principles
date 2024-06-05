package main

import "fmt"

type EmployeeInfo struct{}

func (e *EmployeeInfo) Work() {
	fmt.Println("Employee is working.")
}

func (e *EmployeeInfo) OverTimeWork() {
	fmt.Println("Employee is overtime Working.")
}

func (e *EmployeeInfo) DisplayTimeWorking() {
	fmt.Println("10 Hours of working.")
}

// func main() {
// 	em := &EmployeeInfo{}
// 	em.Work()
// 	em.OverTimeWork()
// 	em.DisplayTimeWorking()
// }